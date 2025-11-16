--- vendor/cloud.google.com/go/auth/grpctransport/grpctransport.go.orig	2023-10-31 21:28:14 UTC
+++ vendor/cloud.google.com/go/auth/grpctransport/grpctransport.go
@@ -23,7 +23,6 @@ import (
 	"cloud.google.com/go/auth"
 	"cloud.google.com/go/auth/detect"
 	"cloud.google.com/go/auth/internal/transport"
-	"go.opencensus.io/plugin/ocgrpc"
 	"google.golang.org/grpc"
 	"google.golang.org/grpc/credentials"
 	grpcinsecure "google.golang.org/grpc/credentials/insecure"
@@ -272,8 +271,5 @@ func addOCStatsHandler(dialOpts []grpc.DialOption, opt
 }
 
 func addOCStatsHandler(dialOpts []grpc.DialOption, opts *Options) []grpc.DialOption {
-	if opts.DisableTelemetry {
 		return dialOpts
-	}
-	return append(dialOpts, grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
 }
