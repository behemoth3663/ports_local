--- vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/client.go.orig	2026-11-15 15:04:53 UTC
+++ vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/client.go
@@ -10,7 +10,6 @@ import (
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
-	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
 )
 
 // ClientOptions contains configuration settings for a client's pipeline.
@@ -20,7 +19,7 @@ type Client struct {
 type Client struct {
 	ep string
 	pl runtime.Pipeline
-	tr tracing.Tracer
+	tr any
 }
 
 // NewClient creates a new Client instance with the provided values.
@@ -48,9 +47,7 @@ func NewClient(moduleName, moduleVersion string, cred 
 	if err != nil {
 		return nil, err
 	}
-
-	tr := options.TracingProvider.NewTracer(moduleName, moduleVersion)
-	return &Client{ep: ep, pl: pl, tr: tr}, nil
+	return &Client{ep: ep, pl: pl, tr: nil}, nil
 }
 
 // Endpoint returns the service's base URL for this client.
@@ -64,6 +61,6 @@ func (c *Client) Pipeline() runtime.Pipeline {
 }
 
 // Tracer returns the tracer for this client.
-func (c *Client) Tracer() tracing.Tracer {
+func (c *Client) Tracer() any {
 	return c.tr
 }
