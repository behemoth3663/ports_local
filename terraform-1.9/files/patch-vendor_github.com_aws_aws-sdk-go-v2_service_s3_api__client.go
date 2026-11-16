--- vendor/github.com/aws/aws-sdk-go-v2/service/s3/api_client.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/s3/api_client.go
@@ -5,7 +5,6 @@ import (
 import (
 	"context"
 	cryptorand "crypto/rand"
-	"errors"
 	"fmt"
 	"github.com/aws/aws-sdk-go-v2/aws"
 	"github.com/aws/aws-sdk-go-v2/aws/defaults"
@@ -31,7 +30,6 @@ import (
 	"github.com/aws/smithy-go/metrics"
 	"github.com/aws/smithy-go/middleware"
 	smithyrand "github.com/aws/smithy-go/rand"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net"
 	"net/http"
@@ -165,10 +163,6 @@ func getOperationMetrics(ctx context.Context) *operati
 	return ctx.Value(operationMetricsKey{}).(*operationMetrics)
 }
 
-func operationTracer(p tracing.TracerProvider) tracing.Tracer {
-	return p.Tracer("github.com/aws/aws-sdk-go-v2/service/s3")
-}
-
 // Client provides the API client to make operations call for Amazon Simple
 // Storage Service.
 type Client struct {
@@ -283,20 +277,8 @@ func (c *Client) invokeOperation(
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
 		o.Meter = options.MeterProvider.Meter("github.com/aws/aws-sdk-go-v2/service/s3")
@@ -304,16 +286,6 @@ func (c *Client) invokeOperation(
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
@@ -321,13 +293,6 @@ func (c *Client) invokeOperation(
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
 
@@ -673,9 +638,6 @@ func (m *spanRetryLoop) HandleFinalize(
 ) (
 	middleware.FinalizeOutput, middleware.Metadata, error,
 ) {
-	tracer := operationTracer(m.options.TracerProvider)
-	ctx, span := tracer.StartSpan(ctx, "RetryLoop")
-	defer span.End()
 
 	return next.HandleFinalize(ctx, in)
 }
@@ -914,9 +876,6 @@ func resolveTracerProvider(options *Options) {
 }
 
 func resolveTracerProvider(options *Options) {
-	if options.TracerProvider == nil {
-		options.TracerProvider = &tracing.NopTracerProvider{}
-	}
 }
 
 func resolveMeterProvider(options *Options) {
@@ -1310,8 +1269,6 @@ func (m *spanInitializeStart) HandleInitialize(
 ) (
 	middleware.InitializeOutput, middleware.Metadata, error,
 ) {
-	ctx, _ = tracing.StartSpan(ctx, "Initialize")
-
 	return next.HandleInitialize(ctx, in)
 }
 
@@ -1327,9 +1284,6 @@ func (m *spanInitializeEnd) HandleInitialize(
 ) (
 	middleware.InitializeOutput, middleware.Metadata, error,
 ) {
-	ctx, span := tracing.PopSpan(ctx)
-	span.End()
-
 	return next.HandleInitialize(ctx, in)
 }
 
@@ -1345,8 +1299,6 @@ func (m *spanBuildRequestStart) HandleSerialize(
 ) (
 	middleware.SerializeOutput, middleware.Metadata, error,
 ) {
-	ctx, _ = tracing.StartSpan(ctx, "BuildRequest")
-
 	return next.HandleSerialize(ctx, in)
 }
 
@@ -1362,9 +1314,6 @@ func (m *spanBuildRequestEnd) HandleBuild(
 ) (
 	middleware.BuildOutput, middleware.Metadata, error,
 ) {
-	ctx, span := tracing.PopSpan(ctx)
-	span.End()
-
 	return next.HandleBuild(ctx, in)
 }
 
