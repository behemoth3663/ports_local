--- vendor/github.com/moby/moby/client/hijack.go.orig	2026-03-05 13:48:11 UTC
+++ vendor/github.com/moby/moby/client/hijack.go
@@ -8,8 +8,6 @@ import (
 	"net/http"
 	"net/url"
 	"time"
-
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 )
 
 // postHijacked sends a POST request and hijacks the connection.
@@ -70,7 +68,7 @@ func setupHijackConn(dialer func(context.Context) (net
 	hc := &hijackedConn{conn, bufio.NewReader(conn)}
 
 	// Server hijacks the connection, error 'connection closed' expected
-	resp, err := otelhttp.NewTransport(hc).RoundTrip(req)
+	resp, err := hc.RoundTrip(req)
 	if err != nil {
 		return nil, "", err
 	}
