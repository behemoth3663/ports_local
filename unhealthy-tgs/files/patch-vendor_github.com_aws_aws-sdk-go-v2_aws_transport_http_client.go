--- vendor/github.com/aws/aws-sdk-go-v2/aws/transport/http/client.go.orig	2026-11-21 14:46:01 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/aws/transport/http/client.go
@@ -11,7 +11,6 @@ import (
 	"time"
 
 	"github.com/aws/aws-sdk-go-v2/aws"
-	"github.com/aws/smithy-go/tracing"
 )
 
 // Defaults for the HTTPTransportBuilder.
@@ -242,7 +241,7 @@ func defaultHTTPTransport() *http.Transport {
 
 	tr := &http.Transport{
 		Proxy:                 http.ProxyFromEnvironment,
-		DialContext:           traceDialContext(dialer.DialContext),
+		DialContext:           dialer.DialContext,
 		TLSHandshakeTimeout:   DefaultHTTPTransportTLSHandleshakeTimeout,
 		MaxIdleConns:          DefaultHTTPTransportMaxIdleConns,
 		MaxIdleConnsPerHost:   DefaultHTTPTransportMaxIdleConnsPerHost,
@@ -261,32 +260,6 @@ type dialContext func(ctx context.Context, network, ad
 
 type dialContext func(ctx context.Context, network, addr string) (net.Conn, error)
 
-func traceDialContext(dc dialContext) dialContext {
-	return func(ctx context.Context, network, addr string) (net.Conn, error) {
-		span, _ := tracing.GetSpan(ctx)
-		span.SetProperty("net.peer.name", addr)
-
-		conn, err := dc(ctx, network, addr)
-		if err != nil {
-			return conn, err
-		}
-
-		raddr := conn.RemoteAddr()
-		if raddr == nil {
-			return conn, err
-		}
-
-		host, port, err := net.SplitHostPort(raddr.String())
-		if err != nil { // don't blow up just because we couldn't parse
-			span.SetProperty("net.peer.addr", raddr.String())
-		} else {
-			span.SetProperty("net.peer.host", host)
-			span.SetProperty("net.peer.port", port)
-		}
-
-		return conn, err
-	}
-}
 
 // shallowCopyStruct creates a shallow copy of the passed in source struct, and
 // returns that copy of the same struct type.
