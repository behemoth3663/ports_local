--- vendor/google.golang.org/api/transport/grpc/dial.go.orig	2026-01-14 16:38:08 UTC
+++ vendor/google.golang.org/api/transport/grpc/dial.go
@@ -21,7 +21,6 @@ import (
 	"cloud.google.com/go/auth/grpctransport"
 	"cloud.google.com/go/auth/oauth2adapt"
 	"cloud.google.com/go/compute/metadata"
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
 	"golang.org/x/oauth2"
 	"golang.org/x/time/rate"
 	"google.golang.org/api/internal"
@@ -377,10 +376,7 @@ func addOpenTelemetryStatsHandler(opts []grpc.DialOpti
 }
 
 func addOpenTelemetryStatsHandler(opts []grpc.DialOption, settings *internal.DialSettings) []grpc.DialOption {
-	if settings.TelemetryDisabled {
 		return opts
-	}
-	return append(opts, grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
 }
 
 // grpcTokenSource supplies PerRPCCredentials from an oauth.TokenSource.
