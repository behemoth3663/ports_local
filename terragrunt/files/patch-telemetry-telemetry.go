--- telemetry/telemetry.go.orig	2024-02-22 10:11:49 UTC
+++ telemetry/telemetry.go
@@ -8,14 +8,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/options"
 
 	"github.com/pkg/errors"
-	otelmetric "go.opentelemetry.io/otel/metric"
-	"go.opentelemetry.io/otel/sdk/metric"
-
-	"go.opentelemetry.io/otel/attribute"
-
-	oteltrace "go.opentelemetry.io/otel/sdk/trace"
-	sdktrace "go.opentelemetry.io/otel/sdk/trace"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // TelemetryOptions - options for telemetry provider.
@@ -27,14 +19,6 @@ type TelemetryOptions struct {
 	ErrWriter  io.Writer
 }
 
-var spanExporter oteltrace.SpanExporter
-var traceProvider *sdktrace.TracerProvider
-var rootTracer trace.Tracer
-
-var meter otelmetric.Meter
-var metricProvider *metric.MeterProvider
-var metricExporter metric.Exporter
-
 // InitTelemetry - initialize the telemetry provider.
 func InitTelemetry(ctx context.Context, opts *TelemetryOptions) error {
 
@@ -51,18 +35,6 @@ func ShutdownTelemetry(ctx context.Context) error {
 
 // ShutdownTelemetry - shutdown the telemetry provider.
 func ShutdownTelemetry(ctx context.Context) error {
-	if traceProvider != nil {
-		if err := traceProvider.Shutdown(ctx); err != nil {
-			return errors.WithStack(err)
-		}
-		traceProvider = nil
-	}
-	if metricProvider != nil {
-		if err := metricProvider.Shutdown(ctx); err != nil {
-			return errors.WithStack(err)
-		}
-		metricProvider = nil
-	}
 	return nil
 }
 
@@ -73,28 +45,6 @@ func Telemetry(opts *options.TerragruntOptions, name s
 		opts.CtxTelemetryCtx = childCtx
 		return Time(opts, name, attrs, fn)
 	})
-}
-
-// mapToAttributes - convert map to attributes to pass to span.SetAttributes.
-func mapToAttributes(data map[string]interface{}) []attribute.KeyValue {
-	var attrs []attribute.KeyValue
-	for k, v := range data {
-		switch val := v.(type) {
-		case string:
-			attrs = append(attrs, attribute.String(k, val))
-		case int:
-			attrs = append(attrs, attribute.Int64(k, int64(val)))
-		case int64:
-			attrs = append(attrs, attribute.Int64(k, val))
-		case float64:
-			attrs = append(attrs, attribute.Float64(k, val))
-		case bool:
-			attrs = append(attrs, attribute.Bool(k, val))
-		default:
-			attrs = append(attrs, attribute.String(k, fmt.Sprintf("%v", val)))
-		}
-	}
-	return attrs
 }
 
 // ErrorMissingEnvVariable error for missing environment variable.
