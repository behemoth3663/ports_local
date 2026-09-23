diff --git a/go.mod b/go.mod
index 03f631fe8..b605fa112 100644
--- go.mod.orig
+++ go.mod
@@ -75,23 +75,6 @@ require (
 	github.com/xeipuuv/gojsonschema v1.2.0
 	github.com/yuin/goldmark v1.8.5
 	github.com/zclconf/go-cty v1.19.0
-	go.opentelemetry.io/contrib/bridges/otellogrus v0.20.1
-	go.opentelemetry.io/otel v1.46.0
-	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc v0.22.0
-	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.22.0
-	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.46.0
-	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.46.0
-	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.46.0
-	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.46.0
-	go.opentelemetry.io/otel/exporters/stdout/stdoutlog v0.22.0
-	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.46.0
-	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.46.0
-	go.opentelemetry.io/otel/log v0.22.0
-	go.opentelemetry.io/otel/metric v1.46.0
-	go.opentelemetry.io/otel/sdk v1.46.0
-	go.opentelemetry.io/otel/sdk/log v0.22.0
-	go.opentelemetry.io/otel/sdk/metric v1.46.0
-	go.opentelemetry.io/otel/trace v1.46.0
 	go.uber.org/mock v0.6.0
 	golang.org/x/crypto v0.56.0
 	golang.org/x/mod v0.40.0
@@ -126,9 +109,6 @@ require (
 	github.com/Azure/azure-sdk-for-go/sdk/security/keyvault/internal v1.2.0 // indirect
 	github.com/Azure/go-ansiterm v0.0.0-20250102033503-faa5f7b0171c // indirect
 	github.com/AzureAD/microsoft-authentication-library-for-go v1.7.2 // indirect
-	github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v1.33.0 // indirect
-	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v0.57.0 // indirect
-	github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping v0.57.0 // indirect
 	github.com/Masterminds/goutils v1.1.1 // indirect
 	github.com/Masterminds/sprig/v3 v3.3.0 // indirect
 	github.com/Microsoft/go-winio v0.6.2 // indirect
@@ -283,12 +263,6 @@ require (
 	github.com/yusufpapurcu/wmi v1.2.4 // indirect
 	github.com/zclconf/go-cty-yaml v1.1.0 // indirect
 	go.mongodb.org/mongo-driver v1.17.9 // indirect
-	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
-	go.opentelemetry.io/contrib/detectors/gcp v1.44.0 // indirect
-	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.70.0 // indirect
-	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.70.0 // indirect
-	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.46.0 // indirect
-	go.opentelemetry.io/proto/otlp v1.11.0 // indirect
 	go.yaml.in/yaml/v2 v2.4.3 // indirect
 	go.yaml.in/yaml/v3 v3.0.5 // indirect
 	go.yaml.in/yaml/v4 v4.0.0-rc.4 // indirect
