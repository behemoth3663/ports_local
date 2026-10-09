--- vendor/github.com/moby/moby/client/hijack.go.orig	2026-10-08 20:29:49 UTC
+++ vendor/github.com/moby/moby/client/hijack.go
@@ -8,8 +8,6 @@ import (
 	"net/http"
 	"net/url"
 	"time"
-
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 )
 
 // postHijacked sends a POST request and hijacks the connection.
@@ -84,19 +82,8 @@ func (cli *Client) setupHijackConn(req *http.Request, 
 		}
 	}()
 
-	cfg := &cli.clientConfig
-
-	var rt http.RoundTripper = otelhttp.NewTransport(hc, cli.traceOpts...)
-	if len(cfg.requestHooks) > 0 || len(cfg.responseHooks) > 0 {
-		rt = &hookTransport{
-			base:      rt,
-			reqHooks:  cfg.requestHooks,
-			respHooks: cfg.responseHooks,
-		}
-	}
-
 	// Server hijacks the connection, error 'connection closed' expected
-	resp, err := rt.RoundTrip(req)
+	resp, err := hc.RoundTrip(req)
 	if err != nil {
 		return nil, "", err
 	}
