--- vendor/google.golang.org/api/transport/grpc/dial.go.orig	2024-02-21 17:28:15 UTC
+++ vendor/google.golang.org/api/transport/grpc/dial.go
@@ -18,8 +18,6 @@ import (
 	"time"
 
 	"cloud.google.com/go/compute/metadata"
-	"go.opencensus.io/plugin/ocgrpc"
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
 	"golang.org/x/oauth2"
 	"golang.org/x/time/rate"
 	"google.golang.org/api/internal"
@@ -64,7 +62,6 @@ func otelGRPCStatsHandler() stats.Handler {
 // dial connections.
 func otelGRPCStatsHandler() stats.Handler {
 	initOtelStatsHandlerOnce.Do(func() {
-		otelStatsHandler = otelgrpc.NewClientHandler()
 	})
 	return otelStatsHandler
 }
@@ -246,17 +243,11 @@ func addOCStatsHandler(opts []grpc.DialOption, setting
 }
 
 func addOCStatsHandler(opts []grpc.DialOption, settings *internal.DialSettings) []grpc.DialOption {
-	if settings.TelemetryDisabled {
-		return opts
-	}
-	return append(opts, grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
+	return opts
 }
 
 func addOpenTelemetryStatsHandler(opts []grpc.DialOption, settings *internal.DialSettings) []grpc.DialOption {
-	if settings.TelemetryDisabled {
-		return opts
-	}
-	return append(opts, grpc.WithStatsHandler(otelGRPCStatsHandler()))
+	return opts
 }
 
 // grpcTokenSource supplies PerRPCCredentials from an oauth.TokenSource.
