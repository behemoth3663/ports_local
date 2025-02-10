--- internal/rpcapi/grpc_testing.go.orig	2024-10-16 12:28:59 UTC
+++ internal/rpcapi/grpc_testing.go
@@ -8,7 +8,6 @@ import (
 	"net"
 	"testing"
 
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
 	"google.golang.org/grpc"
 	"google.golang.org/grpc/test/bufconn"
 	"google.golang.org/protobuf/proto"
@@ -25,11 +24,7 @@ func grpcClientForTesting(ctx context.Context, t *test
 // server end of this fake connection.
 func grpcClientForTesting(ctx context.Context, t *testing.T, registerServices func(srv *grpc.Server)) (conn grpc.ClientConnInterface, close func()) {
 	fakeListener := bufconn.Listen(1024 /* buffer size */)
-	srv := grpc.NewServer(
-		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
-		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
-	)
-
+	srv := grpc.NewServer()
 	// Caller gets an opportunity to register specific services before
 	// we actually start "serving".
 	registerServices(srv)
@@ -49,8 +44,6 @@ func grpcClientForTesting(ctx context.Context, t *test
 		ctx, "testfake",
 		grpc.WithContextDialer(fakeDialer),
 		grpc.WithInsecure(),
-		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
-		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
 	)
 	if err != nil {
 		t.Fatalf("failed to connect to the fake server: %s", err)
