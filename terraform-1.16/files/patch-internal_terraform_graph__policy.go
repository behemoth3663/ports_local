--- internal/terraform/graph_policy.go.orig	2026-09-23 11:57:56 UTC
+++ internal/terraform/graph_policy.go
@@ -5,8 +5,6 @@ import (
 
 import (
 	"sync"
-
-	"go.opentelemetry.io/otel/trace"
 )
 
 // policySubgraph is a subgraph that stores resource policy nodes.
@@ -16,7 +14,6 @@ type policySubgraph struct {
 
 	// span carries the tracing information. We need the span itself so we can end it
 	// when the policy evaluation is finished
-	span trace.Span
 }
 
 func newPolicySubgraph() *policySubgraph {
