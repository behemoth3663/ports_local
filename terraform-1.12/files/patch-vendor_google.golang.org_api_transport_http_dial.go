--- vendor/google.golang.org/api/transport/http/dial.go.orig	2024-01-04 19:45:26 UTC
+++ vendor/google.golang.org/api/transport/http/dial.go
@@ -15,15 +15,12 @@ import (
 	"net/http"
 	"time"
 
-	"go.opencensus.io/plugin/ochttp"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 	"golang.org/x/net/http2"
 	"golang.org/x/oauth2"
 	"google.golang.org/api/googleapi/transport"
 	"google.golang.org/api/internal"
 	"google.golang.org/api/internal/cert"
 	"google.golang.org/api/option"
-	"google.golang.org/api/transport/http/internal/propagation"
 )
 
 // NewClient returns an HTTP client for use communicating with a Google cloud
@@ -208,20 +205,11 @@ func addOpenTelemetryTransport(trans http.RoundTripper
 }
 
 func addOpenTelemetryTransport(trans http.RoundTripper, settings *internal.DialSettings) http.RoundTripper {
-	if settings.TelemetryDisabled {
 		return trans
-	}
-	return otelhttp.NewTransport(trans)
 }
 
 func addOCTransport(trans http.RoundTripper, settings *internal.DialSettings) http.RoundTripper {
-	if settings.TelemetryDisabled {
 		return trans
-	}
-	return &ochttp.Transport{
-		Base:        trans,
-		Propagation: &propagation.HTTPFormat{},
-	}
 }
 
 // clonedTransport returns the given RoundTripper as a cloned *http.Transport.
