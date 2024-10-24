--- internal/stacks/stackruntime/validate.go.orig	2024-10-16 12:28:59 UTC
+++ internal/stacks/stackruntime/validate.go
@@ -6,8 +6,6 @@ import (
 import (
 	"context"
 
-	"go.opentelemetry.io/otel/codes"
-
 	"github.com/hashicorp/terraform/internal/addrs"
 	"github.com/hashicorp/terraform/internal/providers"
 	"github.com/hashicorp/terraform/internal/stacks/stackconfig"
@@ -18,9 +16,6 @@ func Validate(ctx context.Context, req *ValidateReques
 // Validate performs static validation of a full stack configuration, returning
 // diagnostics in case of any detected problems.
 func Validate(ctx context.Context, req *ValidateRequest) tfdiags.Diagnostics {
-	ctx, span := tracer.Start(ctx, "validate stack configuration")
-	defer span.End()
-
 	main := stackeval.NewForValidating(req.Config, stackeval.ValidateOpts{
 		ProviderFactories: req.ProviderFactories,
 	})
@@ -29,9 +24,6 @@ func Validate(ctx context.Context, req *ValidateReques
 	diags = diags.Append(
 		main.DoCleanup(ctx),
 	)
-	if diags.HasErrors() {
-		span.SetStatus(codes.Error, "validation returned errors")
-	}
 	return diags
 }
 
