--- vendor/cloud.google.com/go/storage/grpc_metrics.go.orig	2026-10-20 09:33:13 UTC
+++ vendor/cloud.google.com/go/storage/grpc_metrics.go
@@ -21,13 +21,7 @@ import (
 	"strings"
 	"time"
 
-	mexporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric"
 	"github.com/google/uuid"
-	"go.opentelemetry.io/contrib/detectors/gcp"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/sdk/metric"
-	"go.opentelemetry.io/otel/sdk/metric/metricdata"
-	"go.opentelemetry.io/otel/sdk/resource"
 	"google.golang.org/api/option"
 	"google.golang.org/grpc"
 	"google.golang.org/grpc/experimental/stats"
