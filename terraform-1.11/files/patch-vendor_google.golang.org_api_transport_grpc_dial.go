--- vendor/google.golang.org/api/transport/grpc/dial.go.orig	2023-06-08 22:34:15 UTC
+++ vendor/google.golang.org/api/transport/grpc/dial.go
@@ -16,7 +16,6 @@ import (
 	"strings"
 
 	"cloud.google.com/go/compute/metadata"
-	"go.opencensus.io/plugin/ocgrpc"
 	"golang.org/x/oauth2"
 	"google.golang.org/api/internal"
 	"google.golang.org/api/option"
@@ -205,10 +204,7 @@ func addOCStatsHandler(opts []grpc.DialOption, setting
 }
 
 func addOCStatsHandler(opts []grpc.DialOption, settings *internal.DialSettings) []grpc.DialOption {
-	if settings.TelemetryDisabled {
-		return opts
-	}
-	return append(opts, grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
+	return opts
 }
 
 // grpcTokenSource supplies PerRPCCredentials from an oauth.TokenSource.
