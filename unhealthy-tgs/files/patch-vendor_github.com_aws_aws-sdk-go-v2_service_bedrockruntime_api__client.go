--- vendor/github.com/aws/aws-sdk-go-v2/service/bedrockruntime/api_client.go.orig	2026-11-21 14:46:12 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/bedrockruntime/api_client.go
@@ -26,7 +26,6 @@ import (
 	"github.com/aws/smithy-go/metrics"
 	"github.com/aws/smithy-go/middleware"
 	smithyrand "github.com/aws/smithy-go/rand"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"github.com/aws/smithy-go/transport/http/protocol/restjson1"
 	"io"
@@ -183,9 +182,6 @@ func getOperationMetrics(ctx context.Context) *operati
 	return nil
 }
 
-func operationTracer(p tracing.TracerProvider) tracing.Tracer {
-	return p.Tracer("github.com/aws/aws-sdk-go-v2/service/bedrockruntime")
-}
 
 // Client provides the API client to make operations call for Amazon Bedrock
 // Runtime.
@@ -303,20 +299,9 @@ func (c *Client) invokeOperation(
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
 		o.Meter = options.MeterProvider.Meter("github.com/aws/aws-sdk-go-v2/service/bedrockruntime")
@@ -324,14 +309,9 @@ func (c *Client) invokeOperation(
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
@@ -341,11 +321,8 @@ func (c *Client) invokeOperation(
 		}
 	}
 
-	span.SetProperty("error", err != nil)
 	if err == nil {
-		span.SetStatus(tracing.SpanStatusOK)
 	} else {
-		span.SetStatus(tracing.SpanStatusError)
 	}
 
 	return result, metadata, err
@@ -396,20 +373,9 @@ func (c *Client) invokeEventStreamOperation(
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
 		o.Meter = options.MeterProvider.Meter("github.com/aws/aws-sdk-go-v2/service/bedrockruntime")
@@ -421,14 +387,9 @@ func (c *Client) invokeEventStreamOperation(
 	go func() {
 		_, _, asyncErr := decorated.Handle(ctx, params)
 		if asyncErr != nil {
-			span.SetProperty("exception.type", fmt.Sprintf("%T", asyncErr))
-			span.SetProperty("exception.message", asyncErr.Error())
 
 			var aerr smithy.APIError
 			if errors.As(asyncErr, &aerr) {
-				span.SetProperty("api.error_code", aerr.ErrorCode())
-				span.SetProperty("api.error_message", aerr.ErrorMessage())
-				span.SetProperty("api.error_fault", aerr.ErrorFault().String())
 			}
 
 			asyncErr = &smithy.OperationError{
@@ -437,11 +398,8 @@ func (c *Client) invokeEventStreamOperation(
 				Err:           asyncErr,
 			}
 		}
-		span.SetProperty("error", asyncErr != nil)
 		if asyncErr == nil {
-			span.SetStatus(tracing.SpanStatusOK)
 		} else {
-			span.SetStatus(tracing.SpanStatusError)
 		}
 	}()
 	res := <-results
@@ -708,13 +666,11 @@ func (m *serializeRequestMiddleware) HandleSerialize(
 		return middleware.SerializeOutput{}, middleware.Metadata{}, fmt.Errorf("input %T is not Serializable", in.Request)
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 
 	err := m.options.Protocol.SerializeRequest(ctx, m.operationSchema, input, req)
 
 	endTimer()
-	span.End()
 
 	if err != nil {
 		return middleware.SerializeOutput{}, middleware.Metadata{}, err
@@ -768,14 +724,12 @@ func (m *deserializeResponseMiddleware) HandleDeserial
 		return out, md, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 
 	err = m.options.Protocol.DeserializeResponse(ctx, m.operationSchema, TypeRegistry, resp, m.output)
 	out.Result = m.output
 
 	endTimer()
-	span.End()
 
 	return out, md, err
 }
@@ -1198,7 +1152,7 @@ func resolveTracerProvider(options *Options) {
 
 func resolveTracerProvider(options *Options) {
 	if options.TracerProvider == nil {
-		options.TracerProvider = &tracing.NopTracerProvider{}
+		options.TracerProvider = struct{}{}
 	}
 }
 
