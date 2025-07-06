--- internal/rpcapi/server.go.orig	2025-04-09 13:10:05 UTC
+++ internal/rpcapi/server.go
@@ -9,7 +9,6 @@ import (
 	"os"
 
 	"github.com/hashicorp/go-plugin"
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
 	"google.golang.org/grpc"
 )
 
@@ -39,8 +38,6 @@ func ServePlugin(ctx context.Context, opts ServerOpts)
 		},
 		GRPCServer: func(opts []grpc.ServerOption) *grpc.Server {
 			fullOpts := []grpc.ServerOption{
-				grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
-				grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
 			}
 			fullOpts = append(fullOpts, opts...)
 			server := grpc.NewServer(fullOpts...)
