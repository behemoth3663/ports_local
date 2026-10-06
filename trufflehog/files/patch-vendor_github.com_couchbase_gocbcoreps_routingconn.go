--- vendor/github.com/couchbase/gocbcoreps/routingconn.go.orig	2026-10-06 09:46:09 UTC
+++ vendor/github.com/couchbase/gocbcoreps/routingconn.go
@@ -4,9 +4,6 @@ import (
 	"context"
 	"crypto/tls"
 	"crypto/x509"
-	"go.opentelemetry.io/otel/metric"
-	"go.opentelemetry.io/otel/propagation"
-	"go.opentelemetry.io/otel/trace"
 
 	"github.com/couchbase/goprotostellar/genproto/view_v1"
 
@@ -27,8 +24,6 @@ import (
 	"github.com/couchbase/goprotostellar/genproto/search_v1"
 	"google.golang.org/grpc"
 	"google.golang.org/grpc/credentials"
-
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
 )
 
 type routingConnOptions struct {
@@ -36,8 +31,8 @@ type routingConnOptions struct {
 	ClientCertificate  *x509.CertPool
 	Username           string
 	Password           string
-	TracerProvider     trace.TracerProvider
-	MeterProvider      metric.MeterProvider
+	TracerProvider     interface{}
+	MeterProvider      interface{}
 }
 
 type routingConn struct {
@@ -87,16 +82,6 @@ func dialRoutingConn(ctx context.Context, address stri
 		dialOpts = append(dialOpts, perRpcDialOpt)
 	}
 
-	clientOpts := []otelgrpc.Option{
-		otelgrpc.WithPropagators(propagation.TraceContext{}),
-	}
-	if opts.TracerProvider != nil {
-		clientOpts = append(clientOpts, otelgrpc.WithTracerProvider(opts.TracerProvider))
-	}
-	if opts.MeterProvider != nil {
-		clientOpts = append(clientOpts, otelgrpc.WithMeterProvider(opts.MeterProvider))
-	}
-	dialOpts = append(dialOpts, grpc.WithStatsHandler(otelgrpc.NewClientHandler(clientOpts...)))
 	dialOpts = append(dialOpts, grpc.WithDefaultCallOptions(grpc.MaxRecvMsgSizeCallOption{MaxRecvMsgSize: maxMsgSize}))
 
 	conn, err := grpc.DialContext(ctx, address, dialOpts...)
