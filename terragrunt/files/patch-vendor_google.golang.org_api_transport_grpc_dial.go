diff --git a/vendor/google.golang.org/api/transport/grpc/dial.go b/vendor/google.golang.org/api/transport/grpc/dial.go
index 7c6bf402b..fecea40bc 100644
--- vendor/google.golang.org/api/transport/grpc/dial.go.orig
+++ vendor/google.golang.org/api/transport/grpc/dial.go
@@ -21,7 +21,6 @@ import (
 	"cloud.google.com/go/auth/grpctransport"
 	"cloud.google.com/go/auth/oauth2adapt"
 	"cloud.google.com/go/compute/metadata"
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
 	"golang.org/x/oauth2"
 	"golang.org/x/time/rate"
 	"google.golang.org/api/internal"
@@ -378,10 +377,8 @@ func dial(ctx context.Context, insecure bool, o *internal.DialSettings) (*grpc.C
 }
 
 func addOpenTelemetryStatsHandler(opts []grpc.DialOption, settings *internal.DialSettings) []grpc.DialOption {
-	if settings.TelemetryDisabled {
-		return opts
-	}
-	return append(opts, grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
+	_ = settings
+	return opts
 }
 
 // grpcTokenSource supplies PerRPCCredentials from an oauth.TokenSource.
