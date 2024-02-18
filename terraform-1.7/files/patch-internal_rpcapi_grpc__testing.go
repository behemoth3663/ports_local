--- internal/rpcapi/grpc_testing.go.orig	2024-02-08 01:00:10 UTC
+++ internal/rpcapi/grpc_testing.go 
@@ -8,7 +8,6 @@ import (
 	"net"
 	"testing"
 
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
 	"google.golang.org/grpc"
 	"google.golang.org/grpc/test/bufconn"
 	"google.golang.org/protobuf/proto"
@@ -26,8 +25,6 @@ func grpcClientForTesting(ctx context.Context, t *test
 func grpcClientForTesting(ctx context.Context, t *testing.T, registerServices func(srv *grpc.Server)) (conn grpc.ClientConnInterface, close func()) {
 	fakeListener := bufconn.Listen(1024 /* buffer size */)
 	srv := grpc.NewServer(
-		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
-		grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
 	)
 
 	// Caller gets an opportunity to register specific services before
@@ -49,8 +46,6 @@ func grpcClientForTesting(ctx context.Context, t *test
 		ctx, "testfake",
 		grpc.WithContextDialer(fakeDialer),
 		grpc.WithInsecure(),
-		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
-		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
 	)
 	if err != nil {
 		t.Fatalf("failed to connect to the fake server: %s", err)
