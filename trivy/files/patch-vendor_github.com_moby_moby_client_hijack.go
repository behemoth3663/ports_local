--- vendor/github.com/moby/moby/client/hijack.go.orig	2026-11-02 21:24:37 UTC
+++ vendor/github.com/moby/moby/client/hijack.go
@@ -8,8 +8,6 @@ import (
 	"net/http"
 	"net/url"
 	"time"
-
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 )
 
 // postHijacked sends a POST request and hijacks the connection.
@@ -72,7 +70,7 @@ func (cli *Client) setupHijackConn(req *http.Request, 
 
 	cfg := &cli.clientConfig
 
-	var rt http.RoundTripper = otelhttp.NewTransport(hc, cli.traceOpts...)
+	var rt http.RoundTripper = hc
 	if len(cfg.requestHooks) > 0 || len(cfg.responseHooks) > 0 {
 		rt = &hookTransport{
 			base:      rt,
