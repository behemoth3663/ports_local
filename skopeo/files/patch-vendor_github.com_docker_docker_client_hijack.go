--- vendor/github.com/docker/docker/client/hijack.go.orig	2025-11-28 10:28:17 UTC
+++ vendor/github.com/docker/docker/client/hijack.go
@@ -12,7 +12,6 @@ import (
 	"github.com/docker/docker/api/types"
 	"github.com/docker/docker/api/types/versions"
 	"github.com/pkg/errors"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 )
 
 // postHijacked sends a POST request and hijacks the connection.
@@ -78,7 +77,7 @@ func setupHijackConn(dialer func(context.Context) (net
 	hc := &hijackedConn{conn, bufio.NewReader(conn)}
 
 	// Server hijacks the connection, error 'connection closed' expected
-	resp, err := otelhttp.NewTransport(hc).RoundTrip(req)
+	resp, err := hc.RoundTrip(req)
 	if err != nil {
 		return nil, "", err
 	}
