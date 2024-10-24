--- internal/rpcapi/server.go.orig	2024-10-16 12:28:59 UTC
+++ internal/rpcapi/server.go
@@ -9,7 +9,6 @@ import (
 	"os"
 
 	"github.com/hashicorp/go-plugin"
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
 	"google.golang.org/grpc"
 )
 
@@ -38,10 +37,7 @@ func ServePlugin(ctx context.Context, opts ServerOpts)
 			},
 		},
 		GRPCServer: func(opts []grpc.ServerOption) *grpc.Server {
-			fullOpts := []grpc.ServerOption{
-				grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
-				grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
-			}
+			fullOpts := []grpc.ServerOption{}
 			fullOpts = append(fullOpts, opts...)
 			server := grpc.NewServer(fullOpts...)
 			// We'll also monitor the given context for cancellation
