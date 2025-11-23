--- internal/command/init_run_experiment.go.orig	2025-11-19 13:33:06 UTC
+++ internal/command/init_run_experiment.go
@@ -16,9 +16,6 @@ import (
 	"github.com/hashicorp/terraform/internal/states"
 	"github.com/hashicorp/terraform/internal/terraform"
 	"github.com/hashicorp/terraform/internal/tfdiags"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // `runPssInit` is an altered version of the logic in `run` that contains changes
@@ -101,19 +98,12 @@ func (c *InitCommand) runPssInit(initArgs *arguments.I
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
