diff --git a/vendor/google.golang.org/api/transport/http/dial.go b/vendor/google.golang.org/api/transport/http/dial.go
index 02a8a7410..fdc0c362c 100644
--- vendor/google.golang.org/api/transport/http/dial.go.orig
+++ vendor/google.golang.org/api/transport/http/dial.go
@@ -19,7 +19,6 @@ import (
 	"cloud.google.com/go/auth/credentials"
 	"cloud.google.com/go/auth/httptransport"
 	"cloud.google.com/go/auth/oauth2adapt"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 	"golang.org/x/net/http2"
 	"golang.org/x/oauth2"
 	"google.golang.org/api/googleapi/transport"
@@ -306,10 +305,8 @@ func fallbackBaseTransport() *http.Transport {
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
