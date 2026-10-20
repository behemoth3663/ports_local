--- vendor/cloud.google.com/go/auth/grpctransport/grpctransport.go.orig	2026-10-20 09:33:13 UTC
+++ vendor/cloud.google.com/go/auth/grpctransport/grpctransport.go
@@ -31,7 +31,6 @@ import (
 	"cloud.google.com/go/auth/internal/transport"
 	"cloud.google.com/go/auth/internal/transport/headers"
 	"github.com/googleapis/gax-go/v2/internallog"
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
 	"google.golang.org/grpc"
 	grpccreds "google.golang.org/grpc/credentials"
 	grpcinsecure "google.golang.org/grpc/credentials/insecure"
