--- vendor/google.golang.org/api/transport/grpc/dial.go.orig	2026-10-02 15:35:43 UTC
+++ vendor/google.golang.org/api/transport/grpc/dial.go
@@ -21,7 +21,6 @@ import (
 	"cloud.google.com/go/auth/grpctransport"
 	"cloud.google.com/go/auth/oauth2adapt"
 	"cloud.google.com/go/compute/metadata"
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
 	"golang.org/x/oauth2"
 	"golang.org/x/time/rate"
 	"google.golang.org/api/internal"
@@ -378,10 +377,7 @@ func addOpenTelemetryStatsHandler(opts []grpc.DialOpti
 }
 
 func addOpenTelemetryStatsHandler(opts []grpc.DialOption, settings *internal.DialSettings) []grpc.DialOption {
-	if settings.TelemetryDisabled {
-		return opts
-	}
-	return append(opts, grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
+	return opts
 }
 
 // grpcTokenSource supplies PerRPCCredentials from an oauth.TokenSource.
@@ -395,7 +391,8 @@ func (ts grpcTokenSource) GetRequestMetadata(ctx conte
 
 // GetRequestMetadata gets the request metadata as a map from a grpcTokenSource.
 func (ts grpcTokenSource) GetRequestMetadata(ctx context.Context, uri ...string) (
-	map[string]string, error) {
+	map[string]string, error,
+) {
 	metadata, err := ts.TokenSource.GetRequestMetadata(ctx, uri...)
 	if err != nil {
 		return nil, err
@@ -421,7 +418,8 @@ func (ts grpcAPIKey) GetRequestMetadata(ctx context.Co
 
 // GetRequestMetadata gets the request metadata as a map from a grpcAPIKey.
 func (ts grpcAPIKey) GetRequestMetadata(ctx context.Context, uri ...string) (
-	map[string]string, error) {
+	map[string]string, error,
+) {
 	metadata := map[string]string{
 		"X-goog-api-key": ts.apiKey,
 	}
@@ -459,7 +457,6 @@ func isDirectPathXdsUsed(o *internal.DialSettings) boo
 		return true
 	}
 	return false
-
 }
 
 func isTokenSourceDirectPathCompatible(ts oauth2.TokenSource, o *internal.DialSettings) bool {
