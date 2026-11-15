--- vendor/github.com/aws/aws-sdk-go-v2/aws/transport/http/client.go.orig	2026-11-15 15:04:53 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/aws/transport/http/client.go
@@ -11,7 +11,6 @@ import (
 	"time"
 
 	"github.com/aws/aws-sdk-go-v2/aws"
-	"github.com/aws/smithy-go/tracing"
 )
 
 // Defaults for the HTTPTransportBuilder.
@@ -263,10 +262,8 @@ func traceDialContext(dc dialContext) dialContext {
 
 func traceDialContext(dc dialContext) dialContext {
 	return func(ctx context.Context, network, addr string) (net.Conn, error) {
-		span, _ := tracing.GetSpan(ctx)
-		span.SetProperty("net.peer.name", addr)
-
-		conn, err := dc(ctx, network, addr)
+		_ , _ = 0, 0
+conn, err := dc(ctx, network, addr)
 		if err != nil {
 			return conn, err
 		}
@@ -276,12 +273,8 @@ func traceDialContext(dc dialContext) dialContext {
 			return conn, err
 		}
 
-		host, port, err := net.SplitHostPort(raddr.String())
+		_, _, err = net.SplitHostPort(raddr.String())
 		if err != nil { // don't blow up just because we couldn't parse
-			span.SetProperty("net.peer.addr", raddr.String())
-		} else {
-			span.SetProperty("net.peer.host", host)
-			span.SetProperty("net.peer.port", port)
 		}
 
 		return conn, err
