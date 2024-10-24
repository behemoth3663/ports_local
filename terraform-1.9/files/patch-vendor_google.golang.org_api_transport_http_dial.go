--- vendor/google.golang.org/api/transport/http/dial.go.orig	2024-10-23 22:05:40 UTC
+++ vendor/google.golang.org/api/transport/http/dial.go
@@ -15,7 +15,6 @@ import (
 	"net/http"
 	"time"
 
-	"go.opencensus.io/plugin/ochttp"
 	"golang.org/x/net/http2"
 	"golang.org/x/oauth2"
 	"google.golang.org/api/googleapi/transport"
@@ -213,13 +212,7 @@ func addOCTransport(trans http.RoundTripper, settings 
 }
 
 func addOCTransport(trans http.RoundTripper, settings *internal.DialSettings) http.RoundTripper {
-	if settings.TelemetryDisabled {
-		return trans
-	}
-	return &ochttp.Transport{
-		Base:        trans,
-		Propagation: &propagation.HTTPFormat{},
-	}
+	return trans
 }
 
 // clonedTransport returns the given RoundTripper as a cloned *http.Transport.
