--- vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/client.go.orig	2026-11-26 21:11:25 UTC
+++ vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/client.go
@@ -8,10 +8,8 @@ import (
 	armpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/policy"
 	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
-	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/sanitizer"
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
-	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
 )
 
 // ClientOptions contains configuration settings for a client's pipeline.
@@ -21,7 +19,7 @@ type Client struct {
 type Client struct {
 	ep string
 	pl runtime.Pipeline
-	tr tracing.Tracer
+	tr any
 }
 
 // NewClient creates a new Client instance with the provided values.
@@ -50,9 +48,7 @@ func NewClient(moduleName, moduleVersion string, cred 
 		return nil, err
 	}
 
-	tr := options.TracingProvider.NewTracer(moduleName, moduleVersion)
-	tr = sanitizer.WrapTracer(tr)
-	return &Client{ep: ep, pl: pl, tr: tr}, nil
+	return &Client{ep: ep, pl: pl, tr: nil}, nil
 }
 
 // Endpoint returns the service's base URL for this client.
@@ -66,6 +62,6 @@ func (c *Client) Pipeline() runtime.Pipeline {
 }
 
 // Tracer returns the tracer for this client.
-func (c *Client) Tracer() tracing.Tracer {
+func (c *Client) Tracer() any {
 	return c.tr
 }
