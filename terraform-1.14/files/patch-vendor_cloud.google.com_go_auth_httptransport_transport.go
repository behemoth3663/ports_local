--- vendor/cloud.google.com/go/auth/httptransport/transport.go.orig	2023-10-31 21:28:14 UTC
+++ vendor/cloud.google.com/go/auth/httptransport/transport.go
@@ -25,7 +25,6 @@ import (
 	"cloud.google.com/go/auth/detect"
 	"cloud.google.com/go/auth/internal"
 	"cloud.google.com/go/auth/internal/transport/cert"
-	"go.opencensus.io/plugin/ochttp"
 	"golang.org/x/net/http2"
 )
 
@@ -149,13 +148,7 @@ func addOCTransport(trans http.RoundTripper, opts *Opt
 }
 
 func addOCTransport(trans http.RoundTripper, opts *Options) http.RoundTripper {
-	if opts.DisableTelemetry {
 		return trans
-	}
-	return &ochttp.Transport{
-		Base:        trans,
-		Propagation: &httpFormat{},
-	}
 }
 
 type authTransport struct {
