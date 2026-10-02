--- vendor/cloud.google.com/go/debug.md.orig	2026-10-02 16:09:03 UTC
+++ vendor/cloud.google.com/go/debug.md
@@ -99,7 +99,6 @@ for OpenCensus is now deprecated in the Google Cloud c
 OpenTelemetry, the successor project. The default experimental tracing support
 for OpenCensus is now deprecated in the Google Cloud client libraries for Go.
 
-Using the [OpenTelemetry-Go - OpenCensus Bridge](https://pkg.go.dev/go.opentelemetry.io/otel/bridge/opencensus), you can immediately begin exporting your traces with OpenTelemetry, even while
 dependencies of your application remain instrumented with OpenCensus. If you do
 not use the bridge, you will need to migrate your entire application and all of
 its instrumented dependencies at once.  For simple applications, this may be
@@ -123,7 +122,6 @@ Please refer to the following resources:
 Please refer to the following resources:
 
 * [Sunsetting OpenCensus](https://opentelemetry.io/blog/2023/sunsetting-opencensus/)
-* [OpenTelemetry-Go - OpenCensus Bridge](https://pkg.go.dev/go.opentelemetry.io/otel/bridge/opencensus)
 
 #### OpenTelemetry
 
@@ -157,14 +155,7 @@ import (
     "context"
     "log"
     "os"
-    texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
     octrace "go.opencensus.io/trace"
-    "go.opentelemetry.io/contrib/detectors/gcp"
-    "go.opentelemetry.io/otel"
-    "go.opentelemetry.io/otel/bridge/opencensus"
-    "go.opentelemetry.io/otel/sdk/resource"
-    sdktrace "go.opentelemetry.io/otel/sdk/trace"
-    semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
 )
 
 func main() {
