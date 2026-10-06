--- vendor/github.com/couchbase/gocbcoreps/routingclient.go.orig	2026-10-06 09:46:09 UTC
+++ vendor/github.com/couchbase/gocbcoreps/routingclient.go
@@ -3,8 +3,6 @@ import (
 import (
 	"context"
 	"crypto/x509"
-	"go.opentelemetry.io/otel/metric"
-	"go.opentelemetry.io/otel/trace"
 	"net"
 	"sync"
 
@@ -49,8 +47,8 @@ type DialOptions struct {
 	Logger             *zap.Logger
 	InsecureSkipVerify bool
 	PoolSize           uint32
-	TracerProvider     trace.TracerProvider
-	MeterProvider      metric.MeterProvider
+	TracerProvider     interface{}
+	MeterProvider      interface{}
 }
 
 func Dial(target string, opts *DialOptions) (*RoutingClient, error) {
@@ -200,6 +198,7 @@ func (c *RoutingClient) CollectionV1() admin_collectio
 func (c *RoutingClient) CollectionV1() admin_collection_v1.CollectionAdminServiceClient {
 	return &routingImpl_CollectionV1{c}
 }
+
 func (c *RoutingClient) BucketV1() admin_bucket_v1.BucketAdminServiceClient {
 	return &routingImpl_BucketV1{c}
 }
