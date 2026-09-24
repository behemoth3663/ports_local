--- internal/terraform/node_policy_eval.go.orig	2026-09-23 11:57:56 UTC
+++ internal/terraform/node_policy_eval.go
@@ -8,7 +8,6 @@ import (
 
 	"github.com/hashicorp/terraform/internal/dag"
 	"github.com/hashicorp/terraform/internal/tfdiags"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // nodePolicyEval is a node that completes the building of the policy graph,
@@ -35,12 +34,9 @@ func (n *nodePolicyEval) DynamicExpand(ctx EvalContext
 	ctx.Changes().Close()
 	ctx.State().Close()
 
-	_, span := tracer().Start(ctx.StopCtx(), "terraform.policy.evaluate")
-	policyGraph.span = span
-
 	// Add a finish node that depends on every policy node, so it runs last and
 	// ends the phase span once all policy evaluation has completed.
-	finish := &nodePolicyEvalFinish{span: policyGraph.span}
+	finish := &nodePolicyEvalFinish{}
 	policyGraph.graph.Add(finish)
 	for pn := range policyGraph.graph.VerticesSeq() {
 		if _, ok := pn.(*nodeResourcePolicy); !ok {
@@ -68,7 +64,6 @@ type nodePolicyEvalFinish struct {
 // must tolerate upstream failures so the span is still closed even if a policy
 // node returned error diagnostics.
 type nodePolicyEvalFinish struct {
-	span trace.Span
 }
 
 var _ GraphNodeExecutable = (*nodePolicyEvalFinish)(nil)
@@ -79,7 +74,6 @@ func (n *nodePolicyEvalFinish) Execute(ctx EvalContext
 }
 
 func (n *nodePolicyEvalFinish) Execute(ctx EvalContext, op walkOperation) tfdiags.Diagnostics {
-	n.span.End()
 	return nil
 }
 
