--- vendor/github.com/moby/moby/client/client.go.orig	2026-03-05 13:48:11 UTC
+++ vendor/github.com/moby/moby/client/client.go
@@ -68,7 +68,6 @@ import (
 	cerrdefs "github.com/containerd/errdefs"
 	"github.com/docker/go-connections/sockets"
 	"github.com/moby/moby/client/pkg/versions"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 )
 
 // DummyHost is a hostname used for local communication.
@@ -197,11 +196,6 @@ func New(ops ...Opt) (*Client, error) {
 			client:  client,
 			proto:   hostURL.Scheme,
 			addr:    hostURL.Host,
-			traceOpts: []otelhttp.Option{
-				otelhttp.WithSpanNameFormatter(func(_ string, req *http.Request) string {
-					return req.Method + " " + req.URL.Path
-				}),
-			},
 		},
 	}
 	cfg := &c.clientConfig
@@ -236,8 +230,6 @@ func New(ops ...Opt) (*Client, error) {
 			c.scheme = "http"
 		}
 	}
-
-	c.client.Transport = otelhttp.NewTransport(c.client.Transport, c.traceOpts...)
 
 	if len(cfg.responseHooks) > 0 {
 		c.client.Transport = &responseHookTransport{
