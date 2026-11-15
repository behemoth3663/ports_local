--- vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime/policy_trace_namespace.go.orig	2026-11-15 15:04:53 UTC
+++ vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime/policy_trace_namespace.go
@@ -6,22 +6,9 @@ import (
 import (
 	"net/http"
 
-	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/internal/resource"
-	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
-	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
 )
 
-// httpTraceNamespacePolicy is a policy that adds the az.namespace attribute to the current Span
 func httpTraceNamespacePolicy(req *policy.Request) (resp *http.Response, err error) {
-	rawTracer := req.Raw().Context().Value(shared.CtxWithTracingTracer{})
-	if tracer, ok := rawTracer.(tracing.Tracer); ok && tracer.Enabled() {
-		rt, err := resource.ParseResourceType(req.Raw().URL.Path)
-		if err == nil {
-			// add the namespace attribute to the current span
-			span := tracer.SpanFromContext(req.Raw().Context())
-			span.SetAttributes(tracing.Attribute{Key: shared.TracingNamespaceAttrName, Value: rt.Namespace})
-		}
-	}
 	return req.Next()
 }
