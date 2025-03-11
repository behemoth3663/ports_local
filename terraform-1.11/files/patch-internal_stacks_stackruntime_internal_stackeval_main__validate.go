--- internal/stacks/stackruntime/internal/stackeval/main_validate.go.orig	2024-10-16 12:28:59 UTC
+++ internal/stacks/stackruntime/internal/stackeval/main_validate.go
@@ -8,7 +8,6 @@ import (
 
 	"github.com/hashicorp/terraform/internal/promising"
 	"github.com/hashicorp/terraform/internal/tfdiags"
-	"go.opentelemetry.io/otel/codes"
 )
 
 // ValidateAll checks the validation rules for declared in the configuration
@@ -61,12 +60,7 @@ func (m *Main) walkValidateObject(ctx context.Context,
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
