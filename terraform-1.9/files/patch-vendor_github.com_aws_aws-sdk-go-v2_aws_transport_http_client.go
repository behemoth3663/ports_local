--- vendor/github.com/aws/aws-sdk-go-v2/aws/transport/http/client.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/aws/transport/http/client.go
@@ -10,7 +10,6 @@ import (
 	"time"
 
 	"github.com/aws/aws-sdk-go-v2/aws"
-	"github.com/aws/smithy-go/tracing"
 )
 
 // Defaults for the HTTPTransportBuilder.
@@ -200,30 +199,7 @@ func traceDialContext(dc dialContext) dialContext {
 type dialContext func(ctx context.Context, network, addr string) (net.Conn, error)
 
 func traceDialContext(dc dialContext) dialContext {
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
+	return dc
 }
 
 // shallowCopyStruct creates a shallow copy of the passed in source struct, and
