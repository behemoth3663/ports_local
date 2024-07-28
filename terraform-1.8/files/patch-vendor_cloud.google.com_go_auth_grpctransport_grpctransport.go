--- vendor/cloud.google.com/go/auth/grpctransport/grpctransport.go.orig	2024-03-27 20:20:16 UTC
+++ vendor/cloud.google.com/go/auth/grpctransport/grpctransport.go
@@ -24,7 +24,6 @@ import (
 	"cloud.google.com/go/auth/credentials"
 	"cloud.google.com/go/auth/internal"
 	"cloud.google.com/go/auth/internal/transport"
-	"go.opencensus.io/plugin/ocgrpc"
 	"google.golang.org/grpc"
 	grpccreds "google.golang.org/grpc/credentials"
 	grpcinsecure "google.golang.org/grpc/credentials/insecure"
@@ -299,8 +298,5 @@ func addOCStatsHandler(dialOpts []grpc.DialOption, opt
 }
 
 func addOCStatsHandler(dialOpts []grpc.DialOption, opts *Options) []grpc.DialOption {
-	if opts.DisableTelemetry {
-		return dialOpts
-	}
-	return append(dialOpts, grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
+	return dialOpts
 }
