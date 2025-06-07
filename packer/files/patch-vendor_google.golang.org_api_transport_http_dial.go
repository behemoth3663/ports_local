--- vendor/google.golang.org/api/transport/http/dial.go.orig	2023-11-06 22:14:15 UTC
+++ vendor/google.golang.org/api/transport/http/dial.go
@@ -15,14 +15,12 @@ import (
 	"net/http"
 	"time"
 
-	"go.opencensus.io/plugin/ochttp"
 	"golang.org/x/net/http2"
 	"golang.org/x/oauth2"
 	"google.golang.org/api/googleapi/transport"
 	"google.golang.org/api/internal"
 	"google.golang.org/api/internal/cert"
 	"google.golang.org/api/option"
-	"google.golang.org/api/transport/http/internal/propagation"
 )
 
 // NewClient returns an HTTP client for use communicating with a Google cloud
@@ -204,13 +202,7 @@ func addOCTransport(trans http.RoundTripper, settings 
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
