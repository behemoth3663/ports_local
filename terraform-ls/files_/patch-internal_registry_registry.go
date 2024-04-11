--- internal/registry/registry.go.orig	2023-10-12 15:35:06 UTC
+++ internal/registry/registry.go
@@ -8,7 +8,6 @@ import (
 	"time"
 
 	"github.com/hashicorp/go-cleanhttp"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 )
 
 const (
@@ -27,7 +26,6 @@ type Client struct {
 func NewClient() Client {
 	client := cleanhttp.DefaultClient()
 	client.Timeout = defaultTimeout
-	client.Transport = otelhttp.NewTransport(client.Transport)
 
 	return Client{
 		BaseURL:          defaultBaseURL,
