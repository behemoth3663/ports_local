--- internal/command/init_run.go.orig	2025-11-19 13:33:06 UTC
+++ internal/command/init_run.go
@@ -17,9 +17,6 @@ import (
 	"github.com/hashicorp/terraform/internal/states"
 	"github.com/hashicorp/terraform/internal/terraform"
 	"github.com/hashicorp/terraform/internal/tfdiags"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 )
 
 func (c *InitCommand) run(initArgs *arguments.Init, view views.Init) int {
@@ -99,19 +96,12 @@ func (c *InitCommand) run(initArgs *arguments.Init, vi
 			View:           view,
 		}
 
-		ctx, span := tracer.Start(ctx, "-from-module=...", trace.WithAttributes(
-			attribute.String("module_source", src),
-		))
-
 		initDirFromModuleAbort, initDirFromModuleDiags := c.initDirFromModule(ctx, path, src, hooks)
 		diags = diags.Append(initDirFromModuleDiags)
 		if initDirFromModuleAbort || initDirFromModuleDiags.HasErrors() {
 			view.Diagnostics(diags)
-			span.SetStatus(codes.Error, "module installation failed")
-			span.End()
 			return 1
 		}
-		span.End()
 
 		view.Output(views.EmptyMessage)
 	}
