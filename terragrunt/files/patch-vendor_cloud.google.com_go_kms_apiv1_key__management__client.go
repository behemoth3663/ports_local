--- vendor/cloud.google.com/go/kms/apiv1/key_management_client.go.orig	2026-10-22 19:52:07 UTC
+++ vendor/cloud.google.com/go/kms/apiv1/key_management_client.go
@@ -33,7 +33,6 @@ import (
 	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
 	gax "github.com/googleapis/gax-go/v2"
 	"github.com/googleapis/gax-go/v2/callctx"
-	trace "go.opentelemetry.io/otel/trace"
 	"google.golang.org/api/iterator"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
@@ -2152,7 +2151,7 @@ func (c *keyManagementGRPCClient) DeleteCryptoKey(ctx 
 	}
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.DeleteCryptoKeyOperation")
 	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
+		lro.SetParentSpanContext(nil)
 	}
 	return &DeleteCryptoKeyOperation{
 		lro: lro,
@@ -2182,7 +2181,7 @@ func (c *keyManagementGRPCClient) DeleteCryptoKeyVersi
 	}
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.DeleteCryptoKeyVersionOperation")
 	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
+		lro.SetParentSpanContext(nil)
 	}
 	return &DeleteCryptoKeyVersionOperation{
 		lro: lro,
@@ -3794,7 +3793,7 @@ func (c *keyManagementRESTClient) DeleteCryptoKey(ctx 
 	override := fmt.Sprintf("/v1/%s", resp.GetName())
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.DeleteCryptoKeyOperation")
 	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
+		lro.SetParentSpanContext(nil)
 	}
 	return &DeleteCryptoKeyOperation{
 		lro:      lro,
@@ -3868,7 +3867,7 @@ func (c *keyManagementRESTClient) DeleteCryptoKeyVersi
 	override := fmt.Sprintf("/v1/%s", resp.GetName())
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.DeleteCryptoKeyVersionOperation")
 	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
+		lro.SetParentSpanContext(nil)
 	}
 	return &DeleteCryptoKeyVersionOperation{
 		lro:      lro,
