--- vendor/github.com/moby/moby/client/client.go.orig	2026-11-02 21:24:37 UTC
+++ vendor/github.com/moby/moby/client/client.go
@@ -69,7 +69,6 @@ import (
 	"github.com/docker/go-connections/sockets"
 	"github.com/moby/moby/client/internal/mod"
 	"github.com/moby/moby/client/pkg/versions"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 )
 
 // DummyHost is a hostname used for local communication.
@@ -204,11 +203,6 @@ func New(ops ...Opt) (*Client, error) {
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
@@ -246,8 +240,6 @@ func New(ops ...Opt) (*Client, error) {
 			c.scheme = "http"
 		}
 	}
-
-	c.client.Transport = otelhttp.NewTransport(c.client.Transport, c.traceOpts...)
 
 	if len(cfg.requestHooks) > 0 || len(cfg.responseHooks) > 0 {
 		c.client.Transport = &hookTransport{
