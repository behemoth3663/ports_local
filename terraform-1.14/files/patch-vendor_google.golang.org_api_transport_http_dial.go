--- vendor/google.golang.org/api/transport/http/dial.go.orig	2026-02-17 16:52:11 UTC
+++ vendor/google.golang.org/api/transport/http/dial.go
@@ -19,7 +19,6 @@ import (
 	"cloud.google.com/go/auth/credentials"
 	"cloud.google.com/go/auth/httptransport"
 	"cloud.google.com/go/auth/oauth2adapt"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 	"golang.org/x/net/http2"
 	"golang.org/x/oauth2"
 	"google.golang.org/api/googleapi/transport"
@@ -133,7 +132,6 @@ func newClientNewAuth(ctx context.Context, base http.R
 			DefaultMTLSEndpoint:     ds.DefaultMTLSEndpoint,
 			DefaultScopes:           ds.DefaultScopes,
 			SkipValidation:          skipValidation,
-			TelemetryAttributes:     ds.TelemetryAttributes,
 		},
 		UniverseDomain: ds.UniverseDomain,
 		Logger:         ds.Logger,
@@ -306,10 +304,7 @@ func addOpenTelemetryTransport(trans http.RoundTripper
 }
 
 func addOpenTelemetryTransport(trans http.RoundTripper, settings *internal.DialSettings) http.RoundTripper {
-	if settings.TelemetryDisabled {
 		return trans
-	}
-	return otelhttp.NewTransport(trans)
 }
 
 // clonedTransport returns the given RoundTripper as a cloned *http.Transport.
