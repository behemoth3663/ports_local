--- vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore/core.go.orig	2026-11-26 21:11:25 UTC
+++ vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore/core.go
@@ -8,11 +8,9 @@ import (
 	"sync"
 
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/exported"
-	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/sanitizer"
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
-	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
 )
 
 // AccessToken represents an Azure service bearer access token with expiry information.
@@ -109,10 +107,10 @@ type Client struct {
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
@@ -135,15 +133,9 @@ func NewClient(moduleName, moduleVersion string, plOpt
 
 	pl := runtime.NewPipeline(moduleName, moduleVersion, plOpts, options)
 
-	tr := options.TracingProvider.NewTracer(moduleName, moduleVersion)
-	if tr.Enabled() && plOpts.Tracing.Namespace != "" {
-		tr.SetAttributes(tracing.Attribute{Key: shared.TracingNamespaceAttrName, Value: plOpts.Tracing.Namespace})
-	}
-	tr = sanitizer.WrapTracer(tr)
-
 	return &Client{
 		pl:        pl,
-		tr:        tr,
+		tr:        nil,
 		tp:        options.TracingProvider,
 		modVer:    moduleVersion,
 		namespace: plOpts.Tracing.Namespace,
@@ -156,7 +148,7 @@ func (c *Client) Pipeline() runtime.Pipeline {
 }
 
 // Tracer returns the tracer for this client.
-func (c *Client) Tracer() tracing.Tracer {
+func (c *Client) Tracer() any {
 	return c.tr
 }
 
@@ -164,10 +156,5 @@ func (c *Client) WithClientName(clientName string) *Cl
 // Note that the values for module name and version will be preserved from the source Client.
 //   - clientName - the fully qualified name of the client ("package.Client"); this is used by the tracing provider when creating spans
 func (c *Client) WithClientName(clientName string) *Client {
-	tr := c.tp.NewTracer(clientName, c.modVer)
-	if tr.Enabled() && c.namespace != "" {
-		tr.SetAttributes(tracing.Attribute{Key: shared.TracingNamespaceAttrName, Value: c.namespace})
-	}
-	tr = sanitizer.WrapTracer(tr)
-	return &Client{pl: c.pl, tr: tr, tp: c.tp, modVer: c.modVer, namespace: c.namespace}
+	return &Client{pl: c.pl, tr: nil, tp: c.tp, modVer: c.modVer, namespace: c.namespace}
 }
