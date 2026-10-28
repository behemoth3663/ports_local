--- vendor/github.com/aws/aws-sdk-go-v2/aws/transport/http/client.go.orig	2026-10-28 11:06:54 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/aws/transport/http/client.go
@@ -11,7 +11,6 @@ import (
 	"time"
 
 	"github.com/aws/aws-sdk-go-v2/aws"
-	"github.com/aws/smithy-go/tracing"
 )
 
 // Defaults for the HTTPTransportBuilder.
@@ -263,8 +262,6 @@ func traceDialContext(dc dialContext) dialContext {
 
 func traceDialContext(dc dialContext) dialContext {
 	return func(ctx context.Context, network, addr string) (net.Conn, error) {
-		span, _ := tracing.GetSpan(ctx)
-		span.SetProperty("net.peer.name", addr)
 
 		conn, err := dc(ctx, network, addr)
 		if err != nil {
@@ -277,12 +274,9 @@ func traceDialContext(dc dialContext) dialContext {
 		}
 
 		host, port, err := net.SplitHostPort(raddr.String())
-		if err != nil { // don't blow up just because we couldn't parse
-			span.SetProperty("net.peer.addr", raddr.String())
-		} else {
-			span.SetProperty("net.peer.host", host)
-			span.SetProperty("net.peer.port", port)
-		}
+		_ = host
+		_ = port
+		_ = err
 
 		return conn, err
 	}
