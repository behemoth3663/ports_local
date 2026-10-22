--- vendor/cloud.google.com/go/kms/apiv1/hsm_management_client.go.orig	2026-10-22 19:52:07 UTC
+++ vendor/cloud.google.com/go/kms/apiv1/hsm_management_client.go
@@ -33,7 +33,6 @@ import (
 	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
 	gax "github.com/googleapis/gax-go/v2"
 	"github.com/googleapis/gax-go/v2/callctx"
-	trace "go.opentelemetry.io/otel/trace"
 	"google.golang.org/api/iterator"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
@@ -931,7 +930,7 @@ func (c *hsmManagementGRPCClient) CreateSingleTenantHs
 	}
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.CreateSingleTenantHsmInstanceOperation")
 	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
+		lro.SetParentSpanContext(nil)
 	}
 	return &CreateSingleTenantHsmInstanceOperation{
 		lro: lro,
@@ -961,7 +960,7 @@ func (c *hsmManagementGRPCClient) CreateSingleTenantHs
 	}
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.CreateSingleTenantHsmInstanceProposalOperation")
 	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
+		lro.SetParentSpanContext(nil)
 	}
 	return &CreateSingleTenantHsmInstanceProposalOperation{
 		lro: lro,
@@ -1015,7 +1014,7 @@ func (c *hsmManagementGRPCClient) ExecuteSingleTenantH
 	}
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.ExecuteSingleTenantHsmInstanceProposalOperation")
 	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
+		lro.SetParentSpanContext(nil)
 	}
 	return &ExecuteSingleTenantHsmInstanceProposalOperation{
 		lro: lro,
@@ -1497,7 +1496,7 @@ func (c *hsmManagementRESTClient) CreateSingleTenantHs
 	override := fmt.Sprintf("/v1/%s", resp.GetName())
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.CreateSingleTenantHsmInstanceOperation")
 	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
+		lro.SetParentSpanContext(nil)
 	}
 	return &CreateSingleTenantHsmInstanceOperation{
 		lro:      lro,
@@ -1574,7 +1573,7 @@ func (c *hsmManagementRESTClient) CreateSingleTenantHs
 	override := fmt.Sprintf("/v1/%s", resp.GetName())
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.CreateSingleTenantHsmInstanceProposalOperation")
 	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
+		lro.SetParentSpanContext(nil)
 	}
 	return &CreateSingleTenantHsmInstanceProposalOperation{
 		lro:      lro,
@@ -1719,7 +1718,7 @@ func (c *hsmManagementRESTClient) ExecuteSingleTenantH
 	override := fmt.Sprintf("/v1/%s", resp.GetName())
 	lro := longrunning.InternalNewOperationWithMetadata(*c.LROClient, resp, "*kms.ExecuteSingleTenantHsmInstanceProposalOperation")
 	if gax.IsFeatureEnabled("TRACING") {
-		lro.SetParentSpanContext(trace.SpanContextFromContext(ctx))
+		lro.SetParentSpanContext(nil)
 	}
 	return &ExecuteSingleTenantHsmInstanceProposalOperation{
 		lro:      lro,
