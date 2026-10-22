--- go.mod.orig	1979-11-29 21:00:00 UTC
+++ go.mod
@@ -78,23 +78,6 @@ require (
 	github.com/yuin/goldmark v1.8.6
 	github.com/zclconf/go-cty v1.19.0
 	github.com/zclconf/go-cty-yaml v1.2.0
-	go.opentelemetry.io/contrib/bridges/otellogrus v0.21.0
-	go.opentelemetry.io/otel v1.47.0
-	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc v0.23.0
-	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.23.0
-	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.47.0
-	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.47.0
-	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.47.0
-	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.47.0
-	go.opentelemetry.io/otel/exporters/stdout/stdoutlog v0.23.0
-	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.47.0
-	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.47.0
-	go.opentelemetry.io/otel/log v1.47.0
-	go.opentelemetry.io/otel/metric v1.47.0
-	go.opentelemetry.io/otel/sdk v1.47.0
-	go.opentelemetry.io/otel/sdk/log v1.47.0
-	go.opentelemetry.io/otel/sdk/metric v1.47.0
-	go.opentelemetry.io/otel/trace v1.47.0
 	go.uber.org/mock v0.6.0
 	go.yaml.in/yaml/v3 v3.0.5
 	golang.org/x/crypto v0.57.0
@@ -163,7 +146,6 @@ require (
 	github.com/blang/semver v3.5.1+incompatible // indirect
 	github.com/buger/jsonparser v1.2.0 // indirect
 	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
-	github.com/cenkalti/backoff/v5 v5.0.3 // indirect
 	github.com/cespare/xxhash/v2 v2.3.0 // indirect
 	github.com/charmbracelet/colorprofile v0.4.3 // indirect
 	github.com/charmbracelet/ultraviolet v0.0.0-20260811164956-006e29f97886 // indirect
@@ -211,7 +193,6 @@ require (
 	github.com/googleapis/gax-go/v2 v2.26.2 // indirect
 	github.com/gorilla/css v1.0.1 // indirect
 	github.com/goware/prefixer v0.0.0-20160118172347-395022866408 // indirect
-	github.com/grpc-ecosystem/grpc-gateway/v2 v2.30.0 // indirect
 	github.com/hashicorp/aws-sdk-go-base/v2 v2.0.0-beta.74 // indirect
 	github.com/hashicorp/errwrap v1.1.0 // indirect
 	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
@@ -291,8 +272,13 @@ require (
 	go.opentelemetry.io/contrib/detectors/gcp v1.45.0 // indirect
 	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.70.0 // indirect
 	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.70.0 // indirect
-	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.47.0 // indirect
-	go.opentelemetry.io/proto/otlp v1.11.0 // indirect
+	go.opentelemetry.io/otel v1.47.0 // indirect
+	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.47.0 // indirect
+	go.opentelemetry.io/otel/log v1.47.0 // indirect
+	go.opentelemetry.io/otel/metric v1.47.0 // indirect
+	go.opentelemetry.io/otel/sdk v1.47.0 // indirect
+	go.opentelemetry.io/otel/sdk/metric v1.47.0 // indirect
+	go.opentelemetry.io/otel/trace v1.47.0 // indirect
 	go.yaml.in/yaml/v2 v2.4.3 // indirect
 	go.yaml.in/yaml/v4 v4.0.0-rc.4 // indirect
 	golang.org/x/exp v0.0.0-20260813180055-c1d0aacb2297 // indirect
