--- vendor/github.com/aws/aws-sdk-go-v2/service/ec2/api_client.go.orig	2026-11-21 14:46:10 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/ec2/api_client.go
@@ -26,7 +26,6 @@ import (
 	"github.com/aws/smithy-go/metrics"
 	"github.com/aws/smithy-go/middleware"
 	smithyrand "github.com/aws/smithy-go/rand"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net"
 	"net/http"
@@ -179,9 +178,6 @@ func getOperationMetrics(ctx context.Context) *operati
 	return nil
 }
 
-func operationTracer(p tracing.TracerProvider) tracing.Tracer {
-	return p.Tracer("github.com/aws/aws-sdk-go-v2/service/ec2")
-}
 
 // Client provides the API client to make operations call for Amazon Elastic
 // Compute Cloud.
@@ -293,20 +289,9 @@ func (c *Client) invokeOperation(
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
 		o.Meter = options.MeterProvider.Meter("github.com/aws/aws-sdk-go-v2/service/ec2")
@@ -314,14 +299,9 @@ func (c *Client) invokeOperation(
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
@@ -331,11 +311,8 @@ func (c *Client) invokeOperation(
 		}
 	}
 
-	span.SetProperty("error", err != nil)
 	if err == nil {
-		span.SetStatus(tracing.SpanStatusOK)
 	} else {
-		span.SetStatus(tracing.SpanStatusError)
 	}
 
 	return result, metadata, err
@@ -809,7 +786,7 @@ func resolveTracerProvider(options *Options) {
 
 func resolveTracerProvider(options *Options) {
 	if options.TracerProvider == nil {
-		options.TracerProvider = &tracing.NopTracerProvider{}
+		options.TracerProvider = struct{}{}
 	}
 }
 
