--- vendor/google.golang.org/api/transport/grpc/dial.go.orig	2025-08-12 20:04:23 UTC
+++ vendor/google.golang.org/api/transport/grpc/dial.go
@@ -22,7 +22,6 @@ import (
 	"cloud.google.com/go/auth/grpctransport"
 	"cloud.google.com/go/auth/oauth2adapt"
 	"cloud.google.com/go/compute/metadata"
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
 	"golang.org/x/oauth2"
 	"golang.org/x/time/rate"
 	"google.golang.org/api/internal"
@@ -70,7 +69,6 @@ func otelGRPCStatsHandler() stats.Handler {
 // dial connections.
 func otelGRPCStatsHandler() stats.Handler {
 	initOtelStatsHandlerOnce.Do(func() {
-		otelStatsHandler = otelgrpc.NewClientHandler()
 	})
 	return otelStatsHandler
 }
