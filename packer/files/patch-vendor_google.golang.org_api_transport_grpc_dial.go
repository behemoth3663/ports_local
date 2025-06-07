--- vendor/google.golang.org/api/transport/grpc/dial.go.orig	2025-07-13 15:46:27 UTC
+++ vendor/google.golang.org/api/transport/grpc/dial.go
@@ -17,7 +17,6 @@ import (
 	"time"
 
 	"cloud.google.com/go/compute/metadata"
-	"go.opencensus.io/plugin/ocgrpc"
 	"golang.org/x/oauth2"
 	"golang.org/x/time/rate"
 	"google.golang.org/api/internal"
@@ -204,10 +203,7 @@ func addOCStatsHandler(opts []grpc.DialOption, setting
 }
 
 func addOCStatsHandler(opts []grpc.DialOption, settings *internal.DialSettings) []grpc.DialOption {
-	if settings.TelemetryDisabled {
 		return opts
-	}
-	return append(opts, grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
 }
 
 // grpcTokenSource supplies PerRPCCredentials from an oauth.TokenSource.
