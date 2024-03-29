--- vendor/cloud.google.com/go/internal/trace/trace.go.orig	2024-02-26 15:48:16 UTC
+++ vendor/cloud.google.com/go/internal/trace/trace.go
@@ -17,16 +17,10 @@ import (
 import (
 	"context"
 	"errors"
-	"fmt"
 	"os"
 	"strings"
 	"sync"
 
-	"go.opencensus.io/trace"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	ottrace "go.opentelemetry.io/otel/trace"
 	"google.golang.org/api/googleapi"
 	"google.golang.org/genproto/googleapis/rpc/code"
 	"google.golang.org/grpc/status"
@@ -72,16 +66,14 @@ func IsOpenCensusTracingEnabled() bool {
 // GOOGLE_API_GO_EXPERIMENTAL_TELEMETRY_PLATFORM_TRACING is NOT set to the
 // case-insensitive value "opentelemetry".
 func IsOpenCensusTracingEnabled() bool {
-	return !IsOpenTelemetryTracingEnabled()
+	return false
 }
 
 // IsOpenTelemetryTracingEnabled returns true if the environment variable
 // GOOGLE_API_GO_EXPERIMENTAL_TELEMETRY_PLATFORM_TRACING is set to the
 // case-insensitive value "opentelemetry".
 func IsOpenTelemetryTracingEnabled() bool {
-	openTelemetryTracingEnabledMu.RLock()
-	defer openTelemetryTracingEnabledMu.RUnlock()
-	return openTelemetryTracingEnabled
+	return false
 }
 
 // StartSpan adds a span to the trace with the given name. If IsOpenCensusTracingEnabled
@@ -93,11 +85,6 @@ func StartSpan(ctx context.Context, name string) conte
 // switch to "opentelemetry" and explicitly setting the environment variable to
 // "opencensus" will be required to continue using OpenCensus tracing.
 func StartSpan(ctx context.Context, name string) context.Context {
-	if IsOpenTelemetryTracingEnabled() {
-		ctx, _ = otel.GetTracerProvider().Tracer(OpenTelemetryTracerName).Start(ctx, name)
-	} else {
-		ctx, _ = trace.StartSpan(ctx, name)
-	}
 	return ctx
 }
 
@@ -110,34 +97,8 @@ func EndSpan(ctx context.Context, err error) {
 // switch to "opentelemetry" and explicitly setting the environment variable to
 // "opencensus" will be required to continue using OpenCensus tracing.
 func EndSpan(ctx context.Context, err error) {
-	if IsOpenTelemetryTracingEnabled() {
-		span := ottrace.SpanFromContext(ctx)
-		if err != nil {
-			span.SetStatus(codes.Error, toOpenTelemetryStatusDescription(err))
-			span.RecordError(err)
-		}
-		span.End()
-	} else {
-		span := trace.FromContext(ctx)
-		if err != nil {
-			span.SetStatus(toStatus(err))
-		}
-		span.End()
-	}
 }
 
-// toStatus converts an error to an equivalent OpenCensus status.
-func toStatus(err error) trace.Status {
-	var err2 *googleapi.Error
-	if ok := errors.As(err, &err2); ok {
-		return trace.Status{Code: httpStatusCodeToOCCode(err2.Code), Message: err2.Message}
-	} else if s, ok := status.FromError(err); ok {
-		return trace.Status{Code: int32(s.Code()), Message: s.Message()}
-	} else {
-		return trace.Status{Code: int32(code.Code_UNKNOWN), Message: err.Error()}
-	}
-}
-
 // toOpenTelemetryStatus converts an error to an equivalent OpenTelemetry status description.
 func toOpenTelemetryStatusDescription(err error) string {
 	var err2 *googleapi.Error
@@ -196,58 +157,4 @@ func TracePrintf(ctx context.Context, attrMap map[stri
 // switch to "opentelemetry" and explicitly setting the environment variable to
 // "opencensus" will be required to continue using OpenCensus tracing.
 func TracePrintf(ctx context.Context, attrMap map[string]interface{}, format string, args ...interface{}) {
-	if IsOpenTelemetryTracingEnabled() {
-		attrs := otAttrs(attrMap)
-		ottrace.SpanFromContext(ctx).AddEvent(fmt.Sprintf(format, args...), ottrace.WithAttributes(attrs...))
-	} else {
-		attrs := ocAttrs(attrMap)
-		// TODO: (odeke-em): perhaps just pass around spans due to the cost
-		// incurred from using trace.FromContext(ctx) yet we could avoid
-		// throwing away the work done by ctx, span := trace.StartSpan.
-		trace.FromContext(ctx).Annotatef(attrs, format, args...)
-	}
-}
-
-// ocAttrs converts a generic map to OpenCensus attributes.
-func ocAttrs(attrMap map[string]interface{}) []trace.Attribute {
-	var attrs []trace.Attribute
-	for k, v := range attrMap {
-		var a trace.Attribute
-		switch v := v.(type) {
-		case string:
-			a = trace.StringAttribute(k, v)
-		case bool:
-			a = trace.BoolAttribute(k, v)
-		case int:
-			a = trace.Int64Attribute(k, int64(v))
-		case int64:
-			a = trace.Int64Attribute(k, v)
-		default:
-			a = trace.StringAttribute(k, fmt.Sprintf("%#v", v))
-		}
-		attrs = append(attrs, a)
-	}
-	return attrs
-}
-
-// otAttrs converts a generic map to OpenTelemetry attributes.
-func otAttrs(attrMap map[string]interface{}) []attribute.KeyValue {
-	var attrs []attribute.KeyValue
-	for k, v := range attrMap {
-		var a attribute.KeyValue
-		switch v := v.(type) {
-		case string:
-			a = attribute.Key(k).String(v)
-		case bool:
-			a = attribute.Key(k).Bool(v)
-		case int:
-			a = attribute.Key(k).Int(v)
-		case int64:
-			a = attribute.Key(k).Int64(v)
-		default:
-			a = attribute.Key(k).String(fmt.Sprintf("%#v", v))
-		}
-		attrs = append(attrs, a)
-	}
-	return attrs
 }
