--- vendor/cloud.google.com/go/auth/httptransport/transport.go.orig	2024-03-27 20:20:16 UTC
+++ vendor/cloud.google.com/go/auth/httptransport/transport.go
@@ -26,7 +26,6 @@ import (
 	"cloud.google.com/go/auth/internal"
 	"cloud.google.com/go/auth/internal/transport"
 	"cloud.google.com/go/auth/internal/transport/cert"
-	"go.opencensus.io/plugin/ochttp"
 	"golang.org/x/net/http2"
 )
 
@@ -154,13 +153,7 @@ func addOCTransport(trans http.RoundTripper, opts *Opt
 }
 
 func addOCTransport(trans http.RoundTripper, opts *Options) http.RoundTripper {
-	if opts.DisableTelemetry {
-		return trans
-	}
-	return &ochttp.Transport{
-		Base:        trans,
-		Propagation: &httpFormat{},
-	}
+	return trans
 }
 
 type authTransport struct {
