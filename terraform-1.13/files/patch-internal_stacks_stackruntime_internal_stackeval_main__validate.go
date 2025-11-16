--- internal/stacks/stackruntime/internal/stackeval/main_validate.go.orig	2025-06-11 10:22:08 UTC
+++ internal/stacks/stackruntime/internal/stackeval/main_validate.go
@@ -6,8 +6,6 @@ import (
 import (
 	"context"
 
-	"go.opentelemetry.io/otel/codes"
-
 	"github.com/hashicorp/terraform/internal/promising"
 	"github.com/hashicorp/terraform/internal/tfdiags"
 )
@@ -62,12 +60,7 @@ func (m *Main) walkValidateObject(ctx context.Context,
 // arrange for them to be validated by a separate call to this method.
 func (m *Main) walkValidateObject(ctx context.Context, ws *walkState, obj Validatable) {
 	ws.AsyncTask(ctx, func(ctx context.Context) {
-		ctx, span := tracer.Start(ctx, obj.tracingName()+" validation")
-		defer span.End()
 		diags := obj.Validate(ctx)
 		ws.AddDiags(diags)
-		if diags.HasErrors() {
-			span.SetStatus(codes.Error, "validation returned errors")
-		}
 	})
 }
