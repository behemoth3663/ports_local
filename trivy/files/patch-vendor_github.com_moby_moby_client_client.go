--- vendor/github.com/moby/moby/client/client.go.orig	2026-10-02 12:18:10 UTC
+++ vendor/github.com/moby/moby/client/client.go
@@ -70,7 +70,6 @@ import (
 	"github.com/docker/go-connections/sockets"
 	"github.com/moby/moby/client/internal/mod"
 	"github.com/moby/moby/client/pkg/versions"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 )
 
 // DummyHost is a hostname used for local communication.
@@ -200,16 +199,12 @@ func New(ops ...Opt) (*Client, error) {
 	}
 	c := &Client{
 		clientConfig: clientConfig{
-			host:    DefaultDockerHost,
-			version: MaxAPIVersion,
-			client:  client,
-			proto:   hostURL.Scheme,
-			addr:    hostURL.Host,
-			traceOpts: []otelhttp.Option{
-				otelhttp.WithSpanNameFormatter(func(_ string, req *http.Request) string {
-					return req.Method + " " + req.URL.Path
-				}),
-			},
+			host:      DefaultDockerHost,
+			version:   MaxAPIVersion,
+			client:    client,
+			proto:     hostURL.Scheme,
+			addr:      hostURL.Host,
+			traceOpts: nil,
 		},
 	}
 	cfg := &c.clientConfig
@@ -248,7 +243,7 @@ func New(ops ...Opt) (*Client, error) {
 		}
 	}
 
-	c.client.Transport = otelhttp.NewTransport(c.client.Transport, c.traceOpts...)
+	c.client.Transport = c.client.Transport
 
 	if len(cfg.responseHooks) > 0 {
 		c.client.Transport = &responseHookTransport{
