--- cli/command/telemetry_docker.go.orig	2026-09-23 09:45:10 UTC
+++ cli/command/telemetry_docker.go
@@ -4,7 +4,6 @@ import (
 package command
 
 import (
-	"context"
 	"fmt"
 	"io/fs"
 	"net/url"
@@ -14,11 +13,7 @@ import (
 	"strings"
 	"unicode"
 
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
-	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
-	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
-	sdktrace "go.opentelemetry.io/otel/sdk/trace"
+	"github.com/docker/cli/cli/debug"
 )
 
 const (
@@ -32,7 +27,7 @@ func dockerExporterOTLPEndpoint(cli Cli) (endpoint str
 func dockerExporterOTLPEndpoint(cli Cli) (endpoint string, secure bool) {
 	meta, err := cli.ContextStore().GetMetadata(cli.CurrentContext())
 	if err != nil {
-		otel.Handle(err)
+		debug.LogTelemetryError(err)
 		return "", false
 	}
 
@@ -47,12 +42,13 @@ func dockerExporterOTLPEndpoint(cli Cli) (endpoint str
 	if otelCfg != nil {
 		otelMap, ok := otelCfg.(map[string]any)
 		if !ok {
-			otel.Handle(fmt.Errorf(
+			debug.LogTelemetryError(fmt.Errorf(
 				"unexpected type for field %q: %T (expected: %T)",
 				otelContextFieldName,
 				otelCfg,
 				otelMap,
 			))
+			return "", false
 		}
 		// keys from https://opentelemetry.io/docs/concepts/sdk-configuration/otlp-exporter-configuration/
 		endpoint, _ = otelMap[otelExporterOTLPEndpoint].(string)
@@ -75,7 +71,7 @@ func dockerExporterOTLPEndpoint(cli Cli) (endpoint str
 	// We pretend we're the same as the environment reader.
 	u, err := url.Parse(endpoint)
 	if err != nil {
-		otel.Handle(fmt.Errorf("docker otel endpoint is invalid: %s", err))
+		debug.LogTelemetryError(fmt.Errorf("docker otel endpoint is invalid: %s", err))
 		return "", false
 	}
 
@@ -91,46 +87,14 @@ func dockerExporterOTLPEndpoint(cli Cli) (endpoint str
 	return endpoint, secure
 }
 
-func dockerSpanExporter(ctx context.Context, cli Cli) []sdktrace.TracerProviderOption {
-	endpoint, secure := dockerExporterOTLPEndpoint(cli)
-	if endpoint == "" {
-		return nil
-	}
-
-	opts := []otlptracegrpc.Option{
-		otlptracegrpc.WithEndpoint(endpoint),
-	}
-	if !secure {
-		opts = append(opts, otlptracegrpc.WithInsecure())
-	}
-
-	exp, err := otlptracegrpc.New(ctx, opts...)
-	if err != nil {
-		otel.Handle(err)
-		return nil
-	}
-	return []sdktrace.TracerProviderOption{sdktrace.WithBatcher(exp, sdktrace.WithExportTimeout(exportTimeout))}
+func dockerSpanExporter(cli Cli) bool {
+	endpoint, _ := dockerExporterOTLPEndpoint(cli)
+	return endpoint != ""
 }
 
-func dockerMetricExporter(ctx context.Context, cli Cli) []sdkmetric.Option {
-	endpoint, secure := dockerExporterOTLPEndpoint(cli)
-	if endpoint == "" {
-		return nil
-	}
-
-	opts := []otlpmetricgrpc.Option{
-		otlpmetricgrpc.WithEndpoint(endpoint),
-	}
-	if !secure {
-		opts = append(opts, otlpmetricgrpc.WithInsecure())
-	}
-
-	exp, err := otlpmetricgrpc.New(ctx, opts...)
-	if err != nil {
-		otel.Handle(err)
-		return nil
-	}
-	return []sdkmetric.Option{sdkmetric.WithReader(newCLIReader(exp))}
+func dockerMetricExporter(cli Cli) bool {
+	endpoint, _ := dockerExporterOTLPEndpoint(cli)
+	return endpoint != ""
 }
 
 // unixSocketEndpoint converts the unix scheme from URL to
