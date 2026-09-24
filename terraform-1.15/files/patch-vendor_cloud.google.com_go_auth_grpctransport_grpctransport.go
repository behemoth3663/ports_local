--- vendor/cloud.google.com/go/auth/grpctransport/grpctransport.go.orig	2025-09-22 16:26:20 UTC
+++ vendor/cloud.google.com/go/auth/grpctransport/grpctransport.go
@@ -32,7 +32,6 @@ import (
 	"cloud.google.com/go/auth/internal/transport"
 	"cloud.google.com/go/auth/internal/transport/headers"
 	"github.com/googleapis/gax-go/v2/internallog"
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
 	"google.golang.org/grpc"
 	grpccreds "google.golang.org/grpc/credentials"
 	grpcinsecure "google.golang.org/grpc/credentials/insecure"
@@ -68,12 +67,6 @@ var (
 
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
@@ -352,7 +345,6 @@ func dial(ctx context.Context, secure bool, opts *Opti
 	// Add tracing, but before the other options, so that clients can override the
 	// gRPC stats handler.
 	// This assumes that gRPC options are processed in order, left to right.
-	grpcOpts = addOpenTelemetryStatsHandler(grpcOpts, opts)
 	grpcOpts = append(grpcOpts, opts.GRPCDialOpts...)
 
 	return grpc.DialContext(ctx, transportCreds.Endpoint, grpcOpts...)
@@ -441,8 +433,5 @@ func addOpenTelemetryStatsHandler(dialOpts []grpc.Dial
 }
 
 func addOpenTelemetryStatsHandler(dialOpts []grpc.DialOption, opts *Options) []grpc.DialOption {
-	if opts.DisableTelemetry {
 		return dialOpts
-	}
-	return append(dialOpts, grpc.WithStatsHandler(otelGRPCStatsHandler()))
 }
