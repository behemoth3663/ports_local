--- vendor/github.com/aws/aws-sdk-go-v2/service/route53/api_client.go.orig	2026-11-21 14:46:19 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/route53/api_client.go
@@ -22,7 +22,6 @@ import (
 	"github.com/aws/smithy-go/logging"
 	"github.com/aws/smithy-go/metrics"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net"
 	"net/http"
@@ -176,9 +175,6 @@ func getOperationMetrics(ctx context.Context) *operati
 	return nil
 }
 
-func operationTracer(p tracing.TracerProvider) tracing.Tracer {
-	return p.Tracer("github.com/aws/aws-sdk-go-v2/service/route53")
-}
 
 // Client provides the API client to make operations call for Amazon Route 53.
 type Client struct {
@@ -287,20 +283,9 @@ func (c *Client) invokeOperation(
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
+	
 	endTimer := startMetricTimer(ctx, "client.call.duration")
 	defer endTimer()
-	defer span.End()
 
 	handler := smithyhttp.NewClientHandlerWithOptions(options.HTTPClient, func(o *smithyhttp.ClientHandler) {
 		o.Meter = options.MeterProvider.Meter("github.com/aws/aws-sdk-go-v2/service/route53")
@@ -308,14 +293,9 @@ func (c *Client) invokeOperation(
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
@@ -325,11 +305,8 @@ func (c *Client) invokeOperation(
 		}
 	}
 
-	span.SetProperty("error", err != nil)
 	if err == nil {
-		span.SetStatus(tracing.SpanStatusOK)
 	} else {
-		span.SetStatus(tracing.SpanStatusError)
 	}
 
 	return result, metadata, err
@@ -796,7 +773,7 @@ func resolveTracerProvider(options *Options) {
 
 func resolveTracerProvider(options *Options) {
 	if options.TracerProvider == nil {
-		options.TracerProvider = &tracing.NopTracerProvider{}
+		options.TracerProvider = struct{}{}
 	}
 }
 
