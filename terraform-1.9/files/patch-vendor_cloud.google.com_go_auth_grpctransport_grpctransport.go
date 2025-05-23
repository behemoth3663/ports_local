--- vendor/cloud.google.com/go/auth/grpctransport/grpctransport.go.orig	2025-04-14 20:28:21 UTC
+++ vendor/cloud.google.com/go/auth/grpctransport/grpctransport.go
@@ -31,7 +31,6 @@ import (
 	"cloud.google.com/go/auth/internal"
 	"cloud.google.com/go/auth/internal/transport"
 	"github.com/googleapis/gax-go/v2/internallog"
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
 	"google.golang.org/grpc"
 	grpccreds "google.golang.org/grpc/credentials"
 	grpcinsecure "google.golang.org/grpc/credentials/insecure"
@@ -67,12 +66,6 @@ var (
 
 // otelGRPCStatsHandler returns singleton otelStatsHandler for reuse across all
 // dial connections.
-func otelGRPCStatsHandler() stats.Handler {
-	initOtelStatsHandlerOnce.Do(func() {
-		otelStatsHandler = otelgrpc.NewClientHandler()
-	})
-	return otelStatsHandler
-}
 
 // ClientCertProvider is a function that returns a TLS client certificate to be
 // used when opening TLS connections. It follows the same semantics as
@@ -449,8 +442,5 @@ func addOpenTelemetryStatsHandler(dialOpts []grpc.Dial
 }
 
 func addOpenTelemetryStatsHandler(dialOpts []grpc.DialOption, opts *Options) []grpc.DialOption {
-	if opts.DisableTelemetry {
 		return dialOpts
-	}
-	return append(dialOpts, grpc.WithStatsHandler(otelGRPCStatsHandler()))
 }
