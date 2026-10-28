--- vendor/github.com/aws/smithy-go/transport/http/client.go.orig	2026-10-28 11:06:54 UTC
+++ vendor/github.com/aws/smithy-go/transport/http/client.go
@@ -8,7 +8,6 @@ import (
 	smithy "github.com/aws/smithy-go"
 	"github.com/aws/smithy-go/metrics"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 )
 
 // ClientDo provides the interface for custom HTTP client implementations.
@@ -61,9 +60,6 @@ func (c ClientHandler) Handle(ctx context.Context, inp
 func (c ClientHandler) Handle(ctx context.Context, input interface{}) (
 	out interface{}, metadata middleware.Metadata, err error,
 ) {
-	ctx, span := tracing.StartSpan(ctx, "DoHTTPRequest")
-	defer span.End()
-
 	ctx, client, err := withMetrics(ctx, c.client, c.Meter)
 	if err != nil {
 		return nil, metadata, fmt.Errorf("instrument with HTTP metrics: %w", err)
@@ -79,15 +75,12 @@ func (c ClientHandler) Handle(ctx context.Context, inp
 		return nil, metadata, err
 	}
 
-	span.SetProperty("http.method", req.Method)
-	span.SetProperty("http.request_content_length", -1) // at least indicate unknown
 	length, ok, err := req.StreamLength()
 	if err != nil {
 		return nil, metadata, err
 	}
-	if ok {
-		span.SetProperty("http.request_content_length", length)
-	}
+	_ = length
+	_ = ok
 
 	resp, err := client.Do(builtRequest)
 	if resp == nil {
@@ -115,10 +108,6 @@ func (c ClientHandler) Handle(ctx context.Context, inp
 	if builtRequest.Body != nil {
 		_ = builtRequest.Body.Close()
 	}
-
-	span.SetProperty("net.protocol.version", fmt.Sprintf("%d.%d", resp.ProtoMajor, resp.ProtoMinor))
-	span.SetProperty("http.status_code", resp.StatusCode)
-	span.SetProperty("http.response_content_length", resp.ContentLength)
 
 	return &Response{Response: resp}, metadata, err
 }
