--- vendor/cloud.google.com/go/auth/grpctransport/grpctransport.go.orig	2026-10-06 09:46:09 UTC
+++ vendor/cloud.google.com/go/auth/grpctransport/grpctransport.go
@@ -24,7 +24,6 @@ import (
 	"log/slog"
 	"net/http"
 	"os"
-	"sync"
 
 	"cloud.google.com/go/auth"
 	"cloud.google.com/go/auth/credentials"
@@ -32,11 +31,9 @@ import (
 	"cloud.google.com/go/auth/internal/transport"
 	"cloud.google.com/go/auth/internal/transport/headers"
 	"github.com/googleapis/gax-go/v2/internallog"
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
 	"google.golang.org/grpc"
 	grpccreds "google.golang.org/grpc/credentials"
 	grpcinsecure "google.golang.org/grpc/credentials/insecure"
-	"google.golang.org/grpc/stats"
 )
 
 const (
@@ -49,32 +46,9 @@ const (
 	quotaProjectHeaderKey = "X-goog-user-project"
 )
 
-var (
-	// Set at init time by dial_socketopt.go. If nil, socketopt is not supported.
-	timeoutDialerOption grpc.DialOption
-)
+// Set at init time by dial_socketopt.go. If nil, socketopt is not supported.
+var timeoutDialerOption grpc.DialOption
 
-// otelStatsHandler is a singleton otelgrpc.clientHandler to be used across
-// all dial connections to avoid the memory leak documented in
-// https://github.com/open-telemetry/opentelemetry-go-contrib/issues/4226
-//
-// TODO: When this module depends on a version of otelgrpc containing the fix,
-// replace this singleton with inline usage for simplicity.
-// The fix should be in https://github.com/open-telemetry/opentelemetry-go/pull/5797.
-var (
-	initOtelStatsHandlerOnce sync.Once
-	otelStatsHandler         stats.Handler
-)
-
-// otelGRPCStatsHandler returns singleton otelStatsHandler for reuse across all
-// dial connections.
-func otelGRPCStatsHandler() stats.Handler {
-	initOtelStatsHandlerOnce.Do(func() {
-		otelStatsHandler = otelgrpc.NewClientHandler()
-	})
-	return otelStatsHandler
-}
-
 // ClientCertProvider is a function that returns a TLS client certificate to be
 // used when opening TLS connections. It follows the same semantics as
 // [crypto/tls.Config.GetClientCertificate].
@@ -444,5 +418,5 @@ func addOpenTelemetryStatsHandler(dialOpts []grpc.Dial
 	if opts.DisableTelemetry {
 		return dialOpts
 	}
-	return append(dialOpts, grpc.WithStatsHandler(otelGRPCStatsHandler()))
+	return dialOpts
 }
