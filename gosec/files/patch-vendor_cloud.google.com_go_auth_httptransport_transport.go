--- vendor/cloud.google.com/go/auth/httptransport/transport.go.orig	2025-01-27 22:44:05 UTC
+++ vendor/cloud.google.com/go/auth/httptransport/transport.go
@@ -27,7 +27,6 @@ import (
 	"cloud.google.com/go/auth/internal"
 	"cloud.google.com/go/auth/internal/transport"
 	"cloud.google.com/go/auth/internal/transport/cert"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 	"golang.org/x/net/http2"
 )
 
@@ -169,10 +168,7 @@ func addOpenTelemetryTransport(trans http.RoundTripper
 }
 
 func addOpenTelemetryTransport(trans http.RoundTripper, opts *Options) http.RoundTripper {
-	if opts.DisableTelemetry {
-		return trans
-	}
-	return otelhttp.NewTransport(trans)
+	return trans
 }
 
 type authTransport struct {
