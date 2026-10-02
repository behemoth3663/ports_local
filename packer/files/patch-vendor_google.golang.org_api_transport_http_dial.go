--- vendor/google.golang.org/api/transport/http/dial.go.orig	2026-10-02 20:06:41 UTC
+++ vendor/google.golang.org/api/transport/http/dial.go
@@ -19,7 +19,7 @@ import (
 	"cloud.google.com/go/auth/credentials"
 	"cloud.google.com/go/auth/httptransport"
 	"cloud.google.com/go/auth/oauth2adapt"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
+
 	"golang.org/x/net/http2"
 	"golang.org/x/oauth2"
 	"google.golang.org/api/googleapi/transport"
@@ -306,10 +306,8 @@ func addOpenTelemetryTransport(trans http.RoundTripper
 }
 
 func addOpenTelemetryTransport(trans http.RoundTripper, settings *internal.DialSettings) http.RoundTripper {
-	if settings.TelemetryDisabled {
-		return trans
-	}
-	return otelhttp.NewTransport(trans)
+	_ = settings
+	return trans
 }
 
 // clonedTransport returns the given RoundTripper as a cloned *http.Transport.
