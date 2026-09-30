--- vendor/cloud.google.com/go/auth/httptransport/transport.go.orig	2025-09-22 16:26:20 UTC
+++ vendor/cloud.google.com/go/auth/httptransport/transport.go
@@ -28,7 +28,6 @@ import (
 	"cloud.google.com/go/auth/internal/transport"
 	"cloud.google.com/go/auth/internal/transport/cert"
 	"cloud.google.com/go/auth/internal/transport/headers"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 	"golang.org/x/net/http2"
 )
 
@@ -170,10 +169,7 @@ func addOpenTelemetryTransport(trans http.RoundTripper
 }
 
 func addOpenTelemetryTransport(trans http.RoundTripper, opts *Options) http.RoundTripper {
-	if opts.DisableTelemetry {
 		return trans
-	}
-	return otelhttp.NewTransport(trans)
 }
 
 type authTransport struct {
