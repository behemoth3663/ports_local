--- vendor/github.com/aws/aws-sdk-go-v2/service/sso/api_client.go.orig	2026-11-13 20:26:41 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/sso/api_client.go
@@ -22,7 +22,6 @@ import (
 	"github.com/aws/smithy-go/logging"
 	"github.com/aws/smithy-go/metrics"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"github.com/aws/smithy-go/transport/http/protocol/restjson1"
 	"net"
@@ -176,10 +175,6 @@ func getOperationMetrics(ctx context.Context) *operati
 	return nil
 }
 
-func operationTracer(p tracing.TracerProvider) tracing.Tracer {
-	return p.Tracer("github.com/aws/aws-sdk-go-v2/service/sso")
-}
-
 // Client provides the API client to make operations call for AWS Single Sign-On.
 type Client struct {
 	options Options
@@ -206,8 +201,6 @@ func New(options Options, optFns ...func(*Options)) *C
 
 	resolveEndpointResolverV2(&options)
 
-	resolveTracerProvider(&options)
-
 	resolveMeterProvider(&options)
 
 	resolveAuthSchemeResolver(&options)
@@ -289,20 +282,8 @@ func (c *Client) invokeOperation(
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
 		o.Meter = options.MeterProvider.Meter("github.com/aws/aws-sdk-go-v2/service/sso")
@@ -310,16 +291,6 @@ func (c *Client) invokeOperation(
 	decorated := middleware.DecorateHandler(handler, stack)
 	result, metadata, err = decorated.Handle(ctx, params)
 	if err != nil {
-		span.SetProperty("exception.type", fmt.Sprintf("%T", err))
-		span.SetProperty("exception.message", err.Error())
-
-		var aerr smithy.APIError
-		if errors.As(err, &aerr) {
-			span.SetProperty("api.error_code", aerr.ErrorCode())
-			span.SetProperty("api.error_message", aerr.ErrorMessage())
-			span.SetProperty("api.error_fault", aerr.ErrorFault().String())
-		}
-
 		err = &smithy.OperationError{
 			ServiceID:     ServiceID,
 			OperationName: opID,
@@ -327,13 +298,6 @@ func (c *Client) invokeOperation(
 		}
 	}
 
-	span.SetProperty("error", err != nil)
-	if err == nil {
-		span.SetStatus(tracing.SpanStatusOK)
-	} else {
-		span.SetStatus(tracing.SpanStatusError)
-	}
-
 	return result, metadata, err
 }
 
@@ -455,13 +419,11 @@ func (m *serializeRequestMiddleware) HandleSerialize(
 		return middleware.SerializeOutput{}, middleware.Metadata{}, fmt.Errorf("input %T is not Serializable", in.Request)
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 
 	err := m.options.Protocol.SerializeRequest(ctx, m.operationSchema, input, req)
 
 	endTimer()
-	span.End()
 
 	if err != nil {
 		return middleware.SerializeOutput{}, middleware.Metadata{}, err
@@ -515,14 +477,12 @@ func (m *deserializeResponseMiddleware) HandleDeserial
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
@@ -890,12 +850,6 @@ func addCredentialSource(stack *middleware.Stack, opti
 		ua.AddCredentialsSource(source)
 	}
 	return nil
-}
-
-func resolveTracerProvider(options *Options) {
-	if options.TracerProvider == nil {
-		options.TracerProvider = &tracing.NopTracerProvider{}
-	}
 }
 
 func resolveMeterProvider(options *Options) {
