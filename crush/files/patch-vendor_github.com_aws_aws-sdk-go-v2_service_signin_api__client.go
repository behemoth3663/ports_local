--- vendor/github.com/aws/aws-sdk-go-v2/service/signin/api_client.go.orig	2026-10-28 11:23:21 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/signin/api_client.go
@@ -23,7 +23,6 @@ import (
 	"github.com/aws/smithy-go/metrics"
 	"github.com/aws/smithy-go/middleware"
 	smithyrand "github.com/aws/smithy-go/rand"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net"
 	"net/http"
@@ -176,10 +175,6 @@ func getOperationMetrics(ctx context.Context) *operati
 	return nil
 }
 
-func operationTracer(p tracing.TracerProvider) tracing.Tracer {
-	return p.Tracer("github.com/aws/aws-sdk-go-v2/service/signin")
-}
-
 // Client provides the API client to make operations call for AWS Sign-In Service.
 type Client struct {
 	options Options
@@ -207,9 +202,6 @@ func New(options Options, optFns ...func(*Options)) *C
 	resolveIdempotencyTokenProvider(&options)
 
 	resolveEndpointResolverV2(&options)
-
-	resolveTracerProvider(&options)
-
 	resolveMeterProvider(&options)
 
 	resolveAuthSchemeResolver(&options)
@@ -289,20 +281,8 @@ func (c *Client) invokeOperation(
 		return nil, metadata, err
 	}
 
-	tracer := operationTracer(options.TracerProvider)
-	spanName := fmt.Sprintf("%s.%s", ServiceID, opID)
-
-	ctx = tracing.WithOperationTracer(ctx, tracer)
-
-	ctx, span := tracer.StartSpan(ctx, spanName, func(o *tracing.SpanOptions) {
-		o.Kind = tracing.SpanKindClient
-		o.Properties.Set("rpc.system", "aws-api")
-		o.Properties.Set("rpc.method", opID)
-		o.Properties.Set("rpc.service", ServiceID)
-	})
 	endTimer := startMetricTimer(ctx, "client.call.duration")
 	defer endTimer()
-	defer span.End()
 
 	handler := smithyhttp.NewClientHandlerWithOptions(options.HTTPClient, func(o *smithyhttp.ClientHandler) {
 		o.Meter = options.MeterProvider.Meter("github.com/aws/aws-sdk-go-v2/service/signin")
@@ -310,14 +290,9 @@ func (c *Client) invokeOperation(
 	decorated := middleware.DecorateHandler(handler, stack)
 	result, metadata, err = decorated.Handle(ctx, params)
 	if err != nil {
-		span.SetProperty("exception.type", fmt.Sprintf("%T", err))
-		span.SetProperty("exception.message", err.Error())
 
 		var aerr smithy.APIError
 		if errors.As(err, &aerr) {
-			span.SetProperty("api.error_code", aerr.ErrorCode())
-			span.SetProperty("api.error_message", aerr.ErrorMessage())
-			span.SetProperty("api.error_fault", aerr.ErrorFault().String())
 		}
 
 		err = &smithy.OperationError{
@@ -326,14 +301,6 @@ func (c *Client) invokeOperation(
 			Err:           err,
 		}
 	}
-
-	span.SetProperty("error", err != nil)
-	if err == nil {
-		span.SetStatus(tracing.SpanStatusOK)
-	} else {
-		span.SetStatus(tracing.SpanStatusError)
-	}
-
 	return result, metadata, err
 }
 
@@ -698,9 +665,6 @@ func (m *spanRetryLoop) HandleFinalize(
 ) (
 	middleware.FinalizeOutput, middleware.Metadata, error,
 ) {
-	tracer := operationTracer(m.options.TracerProvider)
-	ctx, span := tracer.StartSpan(ctx, "RetryLoop")
-	defer span.End()
 
 	return next.HandleFinalize(ctx, in)
 }
@@ -845,12 +809,6 @@ func addCredentialSource(stack *middleware.Stack, opti
 
 	mw := setCredentialSourceMiddleware{ua: ua, options: options}
 	return stack.Build.Insert(&mw, "UserAgent", middleware.Before)
-}
-
-func resolveTracerProvider(options *Options) {
-	if options.TracerProvider == nil {
-		options.TracerProvider = &tracing.NopTracerProvider{}
-	}
 }
 
 func resolveMeterProvider(options *Options) {
