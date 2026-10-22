--- vendor/cloud.google.com/go/kms/apiv1/autokey_client.go.orig	2026-10-22 19:52:07 UTC
+++ vendor/cloud.google.com/go/kms/apiv1/autokey_client.go
@@ -33,7 +33,6 @@ import (
 	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
 	gax "github.com/googleapis/gax-go/v2"
 	"github.com/googleapis/gax-go/v2/callctx"
-	trace "go.opentelemetry.io/otel/trace"
 	"google.golang.org/api/iterator"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
@@ -627,7 +626,7 @@ func (c *autokeyGRPCClient) CreateKeyHandle(ctx contex
 	}
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.CreateKeyHandleOperation")
 	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
+		lro.SetParentSpanContext(nil)
 	}
 	return &CreateKeyHandleOperation{
 		lro: lro,
@@ -945,7 +944,7 @@ func (c *autokeyRESTClient) CreateKeyHandle(ctx contex
 	override := fmt.Sprintf("/v1/%s", resp.GetName())
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.CreateKeyHandleOperation")
 	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
+		lro.SetParentSpanContext(nil)
 	}
 	return &CreateKeyHandleOperation{
 		lro:      lro,
