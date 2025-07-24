--- vendor/google.golang.org/api/transport/grpc/dial.go.orig	2025-06-17 23:12:45 UTC
+++ vendor/google.golang.org/api/transport/grpc/dial.go
@@ -22,7 +22,6 @@ import (
 	"cloud.google.com/go/auth/grpctransport"
 	"cloud.google.com/go/auth/oauth2adapt"
 	"cloud.google.com/go/compute/metadata"
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
 	"golang.org/x/oauth2"
 	"golang.org/x/time/rate"
 	"google.golang.org/api/internal"
@@ -66,15 +65,6 @@ var (
 	otelStatsHandler         stats.Handler
 )
 
-// otelGRPCStatsHandler returns singleton otelStatsHandler for reuse across all
-// dial connections.
-func otelGRPCStatsHandler() stats.Handler {
-	initOtelStatsHandlerOnce.Do(func() {
-		otelStatsHandler = otelgrpc.NewClientHandler()
-	})
-	return otelStatsHandler
-}
-
 // Dial returns a GRPC connection for use communicating with a Google cloud
 // service, configured with the given ClientOptions.
 func Dial(ctx context.Context, opts ...option.ClientOption) (*grpc.ClientConn, error) {
@@ -397,10 +387,7 @@ func addOpenTelemetryStatsHandler(opts []grpc.DialOpti
 }
 
 func addOpenTelemetryStatsHandler(opts []grpc.DialOption, settings *internal.DialSettings) []grpc.DialOption {
-	if settings.TelemetryDisabled {
 		return opts
-	}
-	return append(opts, grpc.WithStatsHandler(otelGRPCStatsHandler()))
 }
 
 // grpcTokenSource supplies PerRPCCredentials from an oauth.TokenSource.
