diff --git a/vendor/github.com/moby/moby/client/hijack.go b/vendor/github.com/moby/moby/client/hijack.go
index 31c44e598..b3df732ad 100644
--- vendor/github.com/moby/moby/client/hijack.go.orig
+++ vendor/github.com/moby/moby/client/hijack.go
@@ -9,7 +9,6 @@ import (
 	"net/url"
 	"time"
 
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 )
 
 // postHijacked sends a POST request and hijacks the connection.
@@ -70,7 +69,7 @@ func setupHijackConn(dialer func(context.Context) (net.Conn, error), req *http.R
 	hc := &hijackedConn{conn, bufio.NewReader(conn)}
 
 	// Server hijacks the connection, error 'connection closed' expected
-	resp, err := otelhttp.NewTransport(hc).RoundTrip(req)
+	resp, err := hc.RoundTrip(req)
 	if err != nil {
 		return nil, "", err
 	}
