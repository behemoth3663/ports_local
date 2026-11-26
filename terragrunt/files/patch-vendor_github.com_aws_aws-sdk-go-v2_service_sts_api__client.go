--- vendor/github.com/aws/aws-sdk-go-v2/service/sts/api_client.go.orig	2026-11-26 21:11:26 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/sts/api_client.go
@@ -26,7 +26,6 @@ import (
 	"github.com/aws/smithy-go/logging"
 	"github.com/aws/smithy-go/metrics"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"github.com/aws/smithy-go/transport/http/protocol/awsquery"
 	"net"
@@ -180,10 +179,6 @@ func getOperationMetrics(ctx context.Context) *operati
 	return nil
 }
 
-func operationTracer(p tracing.TracerProvider) tracing.Tracer {
-	return p.Tracer("github.com/aws/aws-sdk-go-v2/service/sts")
-}
-
 // Client provides the API client to make operations call for AWS Security Token
 // Service.
 type Client struct {
@@ -298,36 +293,18 @@ func (c *Client) invokeOperation(
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
-
 	handler := smithyhttp.NewClientHandlerWithOptions(options.HTTPClient, func(o *smithyhttp.ClientHandler) {
 		o.Meter = options.MeterProvider.Meter("github.com/aws/aws-sdk-go-v2/service/sts")
 	})
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
-		}
+}
 
 		err = &smithy.OperationError{
 			ServiceID:     ServiceID,
@@ -335,14 +312,10 @@ func (c *Client) invokeOperation(
 			Err:           err,
 		}
 	}
+if err == nil {
+} else {
+}
 
-	span.SetProperty("error", err != nil)
-	if err == nil {
-		span.SetStatus(tracing.SpanStatusOK)
-	} else {
-		span.SetStatus(tracing.SpanStatusError)
-	}
-
 	return result, metadata, err
 }
 
@@ -469,13 +442,11 @@ func (m *serializeRequestMiddleware) HandleSerialize(
 		return middleware.SerializeOutput{}, middleware.Metadata{}, fmt.Errorf("input %T is not Serializable", in.Request)
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 
 	err := m.options.Protocol.SerializeRequest(ctx, m.operationSchema, input, req)
 
 	endTimer()
-	span.End()
 
 	if err != nil {
 		return middleware.SerializeOutput{}, middleware.Metadata{}, err
@@ -529,14 +500,12 @@ func (m *deserializeResponseMiddleware) HandleDeserial
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
@@ -932,7 +901,7 @@ func resolveTracerProvider(options *Options) {
 
 func resolveTracerProvider(options *Options) {
 	if options.TracerProvider == nil {
-		options.TracerProvider = &tracing.NopTracerProvider{}
+		options.TracerProvider = nil
 	}
 }
 
