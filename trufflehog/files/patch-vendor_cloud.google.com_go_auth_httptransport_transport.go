--- vendor/cloud.google.com/go/auth/httptransport/transport.go.orig	2026-10-06 09:46:09 UTC
+++ vendor/cloud.google.com/go/auth/httptransport/transport.go
@@ -28,7 +28,6 @@ import (
 	"cloud.google.com/go/auth/internal/transport"
 	"cloud.google.com/go/auth/internal/transport/cert"
 	"cloud.google.com/go/auth/internal/transport/headers"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 	"golang.org/x/net/http2"
 )
 
@@ -37,7 +36,7 @@ func newTransport(base http.RoundTripper, opts *Option
 )
 
 func newTransport(base http.RoundTripper, opts *Options) (http.RoundTripper, error) {
-	var headers = opts.Headers
+	headers := opts.Headers
 	ht := &headerTransport{
 		base:    base,
 		headers: headers,
@@ -173,7 +172,7 @@ func addOpenTelemetryTransport(trans http.RoundTripper
 	if opts.DisableTelemetry {
 		return trans
 	}
-	return otelhttp.NewTransport(trans)
+	return trans
 }
 
 type authTransport struct {
