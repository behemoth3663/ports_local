--- vendor/google.golang.org/api/transport/http/dial.go.orig	2021-12-13 21:18:59 UTC
+++ vendor/google.golang.org/api/transport/http/dial.go
@@ -15,13 +15,11 @@ import (
 	"net/http"
 	"time"
 
-	"go.opencensus.io/plugin/ochttp"
 	"golang.org/x/oauth2"
 	"google.golang.org/api/googleapi/transport"
 	"google.golang.org/api/internal"
 	"google.golang.org/api/option"
 	"google.golang.org/api/transport/cert"
-	"google.golang.org/api/transport/http/internal/propagation"
 	"google.golang.org/api/transport/internal/dca"
 )
 
@@ -202,11 +200,5 @@ func addOCTransport(trans http.RoundTripper, settings 
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
