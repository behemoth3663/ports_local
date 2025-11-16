--- vendor/google.golang.org/api/transport/grpc/dial.go.orig	2024-01-04 19:45:26 UTC
+++ vendor/google.golang.org/api/transport/grpc/dial.go
@@ -18,8 +18,6 @@ import (
 	"time"
 
 	"cloud.google.com/go/compute/metadata"
-	"go.opencensus.io/plugin/ocgrpc"
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
 	"golang.org/x/oauth2"
 	"golang.org/x/time/rate"
 	"google.golang.org/api/internal"
@@ -62,12 +60,6 @@ var (
 
 // otelGRPCStatsHandler returns singleton otelStatsHandler for reuse across all
 // dial connections.
-func otelGRPCStatsHandler() stats.Handler {
-	initOtelStatsHandlerOnce.Do(func() {
-		otelStatsHandler = otelgrpc.NewClientHandler()
-	})
-	return otelStatsHandler
-}
 
 // Dial returns a GRPC connection for use communicating with a Google cloud
 // service, configured with the given ClientOptions.
@@ -235,17 +227,11 @@ func addOCStatsHandler(opts []grpc.DialOption, setting
 }
 
 func addOCStatsHandler(opts []grpc.DialOption, settings *internal.DialSettings) []grpc.DialOption {
-	if settings.TelemetryDisabled {
 		return opts
-	}
-	return append(opts, grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
 }
 
 func addOpenTelemetryStatsHandler(opts []grpc.DialOption, settings *internal.DialSettings) []grpc.DialOption {
-	if settings.TelemetryDisabled {
 		return opts
-	}
-	return append(opts, grpc.WithStatsHandler(otelGRPCStatsHandler()))
 }
 
 // grpcTokenSource supplies PerRPCCredentials from an oauth.TokenSource.
