--- vendor/github.com/docker/docker/client/client.go.orig	2025-11-28 10:28:17 UTC
+++ vendor/github.com/docker/docker/client/client.go
@@ -58,7 +58,6 @@ import (
 	"github.com/docker/docker/api/types/versions"
 	"github.com/docker/go-connections/sockets"
 	"github.com/pkg/errors"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 )
 
 // DummyHost is a hostname used for local communication.
@@ -140,8 +139,6 @@ type Client struct {
 	// negotiateLock is used to single-flight the version negotiation process
 	negotiateLock sync.Mutex
 
-	traceOpts []otelhttp.Option
-
 	// When the client transport is an *http.Transport (default) we need to do some extra things (like closing idle connections).
 	// Store the original transport as the http.Client transport will be wrapped with tracing libs.
 	baseTransport *http.Transport
@@ -202,12 +199,6 @@ func NewClientWithOpts(ops ...Opt) (*Client, error) {
 		client:  client,
 		proto:   hostURL.Scheme,
 		addr:    hostURL.Host,
-
-		traceOpts: []otelhttp.Option{
-			otelhttp.WithSpanNameFormatter(func(_ string, req *http.Request) string {
-				return req.Method + " " + req.URL.Path
-			}),
-		},
 	}
 
 	for _, op := range ops {
@@ -234,8 +225,6 @@ func NewClientWithOpts(ops ...Opt) (*Client, error) {
 			c.scheme = "http"
 		}
 	}
-
-	c.client.Transport = otelhttp.NewTransport(c.client.Transport, c.traceOpts...)
 
 	return c, nil
 }
