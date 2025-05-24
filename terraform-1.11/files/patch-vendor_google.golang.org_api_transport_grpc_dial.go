--- vendor/google.golang.org/api/transport/grpc/dial.go.orig	2025-05-06 20:54:19 UTC
+++ vendor/google.golang.org/api/transport/grpc/dial.go
@@ -22,7 +22,6 @@ import (
 	"cloud.google.com/go/auth/grpctransport"
 	"cloud.google.com/go/auth/oauth2adapt"
 	"cloud.google.com/go/compute/metadata"
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
 	"golang.org/x/oauth2"
 	"golang.org/x/time/rate"
 	"google.golang.org/api/internal"
@@ -68,12 +67,6 @@ var (
 
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
@@ -397,10 +390,7 @@ func addOpenTelemetryStatsHandler(opts []grpc.DialOpti
 }
 
 func addOpenTelemetryStatsHandler(opts []grpc.DialOption, settings *internal.DialSettings) []grpc.DialOption {
-	if settings.TelemetryDisabled {
 		return opts
-	}
-	return append(opts, grpc.WithStatsHandler(otelGRPCStatsHandler()))
 }
 
 // grpcTokenSource supplies PerRPCCredentials from an oauth.TokenSource.
