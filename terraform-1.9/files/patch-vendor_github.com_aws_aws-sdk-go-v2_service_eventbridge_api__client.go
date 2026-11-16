--- vendor/github.com/aws/aws-sdk-go-v2/service/eventbridge/api_client.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/eventbridge/api_client.go
@@ -23,7 +23,6 @@ import (
 	"github.com/aws/smithy-go/logging"
 	"github.com/aws/smithy-go/metrics"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net"
 	"net/http"
@@ -157,9 +156,6 @@ func getOperationMetrics(ctx context.Context) *operati
 	return ctx.Value(operationMetricsKey{}).(*operationMetrics)
 }
 
-func operationTracer(p tracing.TracerProvider) tracing.Tracer {
-	return p.Tracer("github.com/aws/aws-sdk-go-v2/service/eventbridge")
-}
 
 // Client provides the API client to make operations call for Amazon EventBridge.
 type Client struct {
@@ -264,20 +260,10 @@ func (c *Client) invokeOperation(
 		return nil, metadata, err
 	}
 
-	tracer := operationTracer(options.TracerProvider)
-	spanName := fmt.Sprintf("%s.%s", ServiceID, opID)
 
-	ctx = tracing.WithOperationTracer(ctx, tracer)
 
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
 		o.Meter = options.MeterProvider.Meter("github.com/aws/aws-sdk-go-v2/service/eventbridge")
@@ -285,14 +271,9 @@ func (c *Client) invokeOperation(
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
@@ -302,11 +283,8 @@ func (c *Client) invokeOperation(
 		}
 	}
 
-	span.SetProperty("error", err != nil)
 	if err == nil {
-		span.SetStatus(tracing.SpanStatusOK)
 	} else {
-		span.SetStatus(tracing.SpanStatusError)
 	}
 
 	return result, metadata, err
@@ -643,9 +621,6 @@ func (m *spanRetryLoop) HandleFinalize(
 ) (
 	middleware.FinalizeOutput, middleware.Metadata, error,
 ) {
-	tracer := operationTracer(m.options.TracerProvider)
-	ctx, span := tracer.StartSpan(ctx, "RetryLoop")
-	defer span.End()
 
 	return next.HandleFinalize(ctx, in)
 }
@@ -824,7 +799,6 @@ func resolveTracerProvider(options *Options) {
 
 func resolveTracerProvider(options *Options) {
 	if options.TracerProvider == nil {
-		options.TracerProvider = &tracing.NopTracerProvider{}
 	}
 }
 
@@ -898,7 +872,6 @@ func (m *spanInitializeStart) HandleInitialize(
 ) (
 	middleware.InitializeOutput, middleware.Metadata, error,
 ) {
-	ctx, _ = tracing.StartSpan(ctx, "Initialize")
 
 	return next.HandleInitialize(ctx, in)
 }
@@ -915,8 +888,6 @@ func (m *spanInitializeEnd) HandleInitialize(
 ) (
 	middleware.InitializeOutput, middleware.Metadata, error,
 ) {
-	ctx, span := tracing.PopSpan(ctx)
-	span.End()
 
 	return next.HandleInitialize(ctx, in)
 }
@@ -933,7 +904,6 @@ func (m *spanBuildRequestStart) HandleSerialize(
 ) (
 	middleware.SerializeOutput, middleware.Metadata, error,
 ) {
-	ctx, _ = tracing.StartSpan(ctx, "BuildRequest")
 
 	return next.HandleSerialize(ctx, in)
 }
@@ -950,8 +920,6 @@ func (m *spanBuildRequestEnd) HandleBuild(
 ) (
 	middleware.BuildOutput, middleware.Metadata, error,
 ) {
-	ctx, span := tracing.PopSpan(ctx)
-	span.End()
 
 	return next.HandleBuild(ctx, in)
 }
