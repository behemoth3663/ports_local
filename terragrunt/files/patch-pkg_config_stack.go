--- pkg/config/stack.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/config/stack.go
@@ -19,7 +19,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/spinner"
 	"github.com/gruntwork-io/terragrunt/internal/strict"
 	"github.com/gruntwork-io/terragrunt/internal/strict/controls"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tf"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
@@ -589,17 +588,9 @@ func generateUnits(
 				util.RelPathForLog(opts.rootWorkingDir, opts.sourceFile, opts.logShowAbsPaths),
 			)
 
-			return telemetry.TelemeterFromContext(ctx).
-				Collect(ctx, l, "stack_generate_unit", map[string]any{
-					"stack_file":  opts.sourceFile,
-					"unit_name":   unit.Name,
-					"unit_source": unit.Source,
-					"unit_path":   unit.Path,
-				}, func(ctx context.Context, l log.Logger) error {
-					return spinner.ShowAfter(ctx, l, generateMessages(opts, &item), func() error {
-						return generateComponent(ctx, l, v, opts, &item)
-					})
-				})
+			return spinner.ShowAfter(ctx, l, generateMessages(opts, &item), func() error {
+				return generateComponent(ctx, l, v, opts, &item)
+			})
 		})
 	}
 
@@ -644,17 +635,9 @@ func generateStacks(
 				util.RelPathForLog(opts.rootWorkingDir, opts.sourceFile, opts.logShowAbsPaths),
 			)
 
-			return telemetry.TelemeterFromContext(ctx).
-				Collect(ctx, l, "stack_generate_stack", map[string]any{
-					"stack_file":   opts.sourceFile,
-					"stack_name":   stack.Name,
-					"stack_source": stack.Source,
-					"stack_path":   stack.Path,
-				}, func(ctx context.Context, l log.Logger) error {
-					return spinner.ShowAfter(ctx, l, generateMessages(opts, &item), func() error {
-						return generateComponent(ctx, l, v, opts, &item)
-					})
-				})
+			return spinner.ShowAfter(ctx, l, generateMessages(opts, &item), func() error {
+				return generateComponent(ctx, l, v, opts, &item)
+			})
 		})
 	}
 
