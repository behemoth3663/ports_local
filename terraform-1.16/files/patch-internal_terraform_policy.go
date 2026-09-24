--- internal/terraform/policy.go.orig	2026-09-23 11:57:56 UTC
+++ internal/terraform/policy.go
@@ -10,8 +10,6 @@ import (
 	"log"
 
 	"github.com/zclconf/go-cty/cty"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 
 	"github.com/hashicorp/terraform/internal/addrs"
 	"github.com/hashicorp/terraform/internal/configs"
@@ -28,10 +26,6 @@ func evaluatePolicies(ctx EvalContext, target addrs.Ab
 	// We want a per-resource parent span so we can reason about the evaluation of individual
 	// resources in the trace
 	evalCtx := ctx.StopCtx()
-	if phaseSpan := ctx.PolicyGraph().span; phaseSpan != nil {
-		evalCtx = trace.ContextWithSpan(evalCtx, phaseSpan)
-	}
-
 	result := ctx.PolicyClient().EvaluateResource(evalCtx, policy.EvaluationRequest[*proto.PolicyEvaluateResourceRequest_ResourceMetadata]{
 		Target:     target.Resource.Resource.Type,
 		Attrs:      policy.CtyToPolicyValue(attrs),
@@ -51,11 +45,6 @@ func getResourcesForPolicyCallback(ctx EvalContext, wa
 
 func getResourcesForPolicyCallback(ctx EvalContext, walkOperation walkOperation, provider providers.Interface, schema providers.GetProviderSchemaResponse, config *configs.Config) func(callbackCtx context.Context, target string, attrs cty.Value) ([]cty.Value, bool, error) {
 	return func(c context.Context, target string, attrs cty.Value) ([]cty.Value, bool, error) {
-		_, span := tracer().Start(c, "policy.callback.getResources", trace.WithAttributes(
-			attribute.String("policy.callback.getResources.type", target),
-		))
-		defer span.End()
-
 		found := make([]cty.Value, 0)
 		var filterMap map[string]cty.Value
 		if !attrs.IsNull() {
@@ -121,17 +110,12 @@ func getResourcesForPolicyCallback(ctx EvalContext, wa
 				}
 			}
 		})
-		span.SetAttributes(attribute.String("policy.callback.getResources.result_count", fmt.Sprintf("%d", len(found))))
 		return found, isPartialResult, nil
 	}
 }
 
 func getDataSourceForPolicyCallback(ctx EvalContext, provider providers.Interface, schema providers.GetProviderSchemaResponse) func(callbackCtx context.Context, datasource string, attrs cty.Value) (cty.Value, bool, error) {
 	return func(c context.Context, target string, attrs cty.Value) (cty.Value, bool, error) {
-		_, span := tracer().Start(c, "policy.callback.getDataSource", trace.WithAttributes(
-			attribute.String("policy.callback.getDataSource.type", target),
-		))
-		defer span.End()
 		if datasource, ok := schema.DataSources[target]; ok {
 			configVal, err := datasource.Body.CoerceValue(attrs)
 			if err != nil {
