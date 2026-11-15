--- vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore/core.go.orig	2026-11-15 15:04:53 UTC
+++ vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore/core.go
@@ -11,7 +11,6 @@ import (
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
-	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
 )
 
 // AccessToken represents an Azure service bearer access token with expiry information.
@@ -108,10 +107,10 @@ type Client struct {
 // Client is a basic HTTP client.  It consists of a pipeline and tracing provider.
 type Client struct {
 	pl runtime.Pipeline
-	tr tracing.Tracer
+	tr any
 
 	// cached on the client to support shallow copying with new values
-	tp        tracing.Provider
+	tp        any
 	modVer    string
 	namespace string
 }
@@ -133,12 +132,8 @@ func NewClient(moduleName, moduleVersion string, plOpt
 	}
 
 	pl := runtime.NewPipeline(moduleName, moduleVersion, plOpts, options)
+	tr := any(nil)
 
-	tr := options.TracingProvider.NewTracer(moduleName, moduleVersion)
-	if tr.Enabled() && plOpts.Tracing.Namespace != "" {
-		tr.SetAttributes(tracing.Attribute{Key: shared.TracingNamespaceAttrName, Value: plOpts.Tracing.Namespace})
-	}
-
 	return &Client{
 		pl:        pl,
 		tr:        tr,
@@ -154,7 +149,7 @@ func (c *Client) Pipeline() runtime.Pipeline {
 }
 
 // Tracer returns the tracer for this client.
-func (c *Client) Tracer() tracing.Tracer {
+func (c *Client) Tracer() any {
 	return c.tr
 }
 
@@ -162,9 +157,5 @@ func (c *Client) WithClientName(clientName string) *Cl
 // Note that the values for module name and version will be preserved from the source Client.
 //   - clientName - the fully qualified name of the client ("package.Client"); this is used by the tracing provider when creating spans
 func (c *Client) WithClientName(clientName string) *Client {
-	tr := c.tp.NewTracer(clientName, c.modVer)
-	if tr.Enabled() && c.namespace != "" {
-		tr.SetAttributes(tracing.Attribute{Key: shared.TracingNamespaceAttrName, Value: c.namespace})
-	}
-	return &Client{pl: c.pl, tr: tr, tp: c.tp, modVer: c.modVer, namespace: c.namespace}
+	return &Client{pl: c.pl, tr: nil, tp: c.tp, modVer: c.modVer, namespace: c.namespace}
 }
