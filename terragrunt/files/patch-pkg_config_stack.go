--- pkg/config/stack.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/config/stack.go
@@ -15,7 +15,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/getter"
 	inthclparse "github.com/gruntwork-io/terragrunt/internal/hclparse"
 	"github.com/gruntwork-io/terragrunt/internal/strict"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tf"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
@@ -582,15 +581,7 @@ func generateUnits(
 				util.RelPathForLog(opts.rootWorkingDir, opts.sourceFile, opts.logShowAbsPaths),
 			)
 
-			return telemetry.TelemeterFromContext(ctx).
-				Collect(ctx, l, "stack_generate_unit", map[string]any{
-					"stack_file":  opts.sourceFile,
-					"unit_name":   unit.Name,
-					"unit_source": unit.Source,
-					"unit_path":   unit.Path,
-				}, func(ctx context.Context, l log.Logger) error {
-					return generateComponent(ctx, l, v, opts, &item)
-				})
+			return generateComponent(ctx, l, v, opts, &item)
 		})
 	}
 
@@ -635,15 +626,7 @@ func generateStacks(
 				util.RelPathForLog(opts.rootWorkingDir, opts.sourceFile, opts.logShowAbsPaths),
 			)
 
-			return telemetry.TelemeterFromContext(ctx).
-				Collect(ctx, l, "stack_generate_stack", map[string]any{
-					"stack_file":   opts.sourceFile,
-					"stack_name":   stack.Name,
-					"stack_source": stack.Source,
-					"stack_path":   stack.Path,
-				}, func(ctx context.Context, l log.Logger) error {
-					return generateComponent(ctx, l, v, opts, &item)
-				})
+			return generateComponent(ctx, l, v, opts, &item)
 		})
 	}
 
