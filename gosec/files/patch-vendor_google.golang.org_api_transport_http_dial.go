--- vendor/google.golang.org/api/transport/http/dial.go.orig	2025-02-05 16:17:31 UTC
+++ vendor/google.golang.org/api/transport/http/dial.go
@@ -19,7 +19,6 @@ import (
 	"cloud.google.com/go/auth/credentials"
 	"cloud.google.com/go/auth/httptransport"
 	"cloud.google.com/go/auth/oauth2adapt"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 	"golang.org/x/net/http2"
 	"golang.org/x/oauth2"
 	"google.golang.org/api/googleapi/transport"
@@ -300,10 +299,7 @@ func addOpenTelemetryTransport(trans http.RoundTripper
 }
 
 func addOpenTelemetryTransport(trans http.RoundTripper, settings *internal.DialSettings) http.RoundTripper {
-	if settings.TelemetryDisabled {
-		return trans
-	}
-	return otelhttp.NewTransport(trans)
+	return trans
 }
 
 // clonedTransport returns the given RoundTripper as a cloned *http.Transport.
