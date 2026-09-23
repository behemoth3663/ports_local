diff --git a/vendor/github.com/moby/moby/client/client.go b/vendor/github.com/moby/moby/client/client.go
index 4b4ef976a..4b371347f 100644
--- vendor/github.com/moby/moby/client/client.go.orig
+++ vendor/github.com/moby/moby/client/client.go
@@ -70,7 +70,6 @@ import (
 	"github.com/docker/go-connections/sockets"
 	"github.com/moby/moby/client/internal/mod"
 	"github.com/moby/moby/client/pkg/versions"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 )
 
 // DummyHost is a hostname used for local communication.
@@ -205,11 +204,7 @@ func New(ops ...Opt) (*Client, error) {
 			client:  client,
 			proto:   hostURL.Scheme,
 			addr:    hostURL.Host,
-			traceOpts: []otelhttp.Option{
-				otelhttp.WithSpanNameFormatter(func(_ string, req *http.Request) string {
-					return req.Method + " " + req.URL.Path
-				}),
-			},
+			traceOpts: nil,
 		},
 	}
 	cfg := &c.clientConfig
@@ -248,8 +243,6 @@ func New(ops ...Opt) (*Client, error) {
 		}
 	}
 
-	c.client.Transport = otelhttp.NewTransport(c.client.Transport, c.traceOpts...)
-
 	if len(cfg.responseHooks) > 0 {
 		c.client.Transport = &responseHookTransport{
 			base:  c.client.Transport,
