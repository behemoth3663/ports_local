--- vendor/cloud.google.com/go/storage/storage.go.orig	2026-10-20 09:33:13 UTC
+++ vendor/cloud.google.com/go/storage/storage.go
@@ -44,9 +44,6 @@ import (
 	"cloud.google.com/go/storage/internal"
 	"cloud.google.com/go/storage/internal/apiv2/storagepb"
 	"github.com/googleapis/gax-go/v2"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/sdk/metric"
-	"go.opentelemetry.io/otel/sdk/metric/metricdata"
 	"google.golang.org/api/googleapi"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
