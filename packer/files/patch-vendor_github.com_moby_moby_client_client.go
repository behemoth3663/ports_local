--- vendor/github.com/moby/moby/client/client.go.orig	2026-10-02 20:06:40 UTC
+++ vendor/github.com/moby/moby/client/client.go
@@ -70,7 +70,6 @@ import (
 	"github.com/docker/go-connections/sockets"
 	"github.com/moby/moby/client/internal/mod"
 	"github.com/moby/moby/client/pkg/versions"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 )
 
 // DummyHost is a hostname used for local communication.
@@ -205,11 +204,6 @@ func New(ops ...Opt) (*Client, error) {
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
@@ -247,8 +241,6 @@ func New(ops ...Opt) (*Client, error) {
 			c.scheme = "http"
 		}
 	}
-
-	c.client.Transport = otelhttp.NewTransport(c.client.Transport, c.traceOpts...)
 
 	if len(cfg.responseHooks) > 0 {
 		c.client.Transport = &responseHookTransport{
