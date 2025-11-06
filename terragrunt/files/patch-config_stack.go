--- config/stack.go.orig	2025-11-05 17:03:04 UTC
+++ config/stack.go
@@ -10,7 +10,6 @@ import (
 	"strings"
 
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/tf"
 
 	"github.com/gruntwork-io/terragrunt/internal/ctyhelper"
@@ -375,25 +374,8 @@ func StackOutput(ctx context.Context, l log.Logger, op
 		}
 
 		for _, unit := range stackFile.Units {
-			unitDir := getUnitDir(dir, unit)
-
 			var output map[string]cty.Value
 
-			telemetryErr := telemetry.TelemeterFromContext(ctx).Collect(ctx, "unit_output", map[string]any{
-				"unit_name":   unit.Name,
-				"unit_source": unit.Source,
-				"unit_path":   unit.Path,
-			}, func(ctx context.Context) error {
-				var outputErr error
-
-				output, outputErr = unit.ReadOutputs(ctx, l, opts, unitDir)
-
-				return outputErr
-			})
-			if telemetryErr != nil {
-				return cty.NilVal, errors.New(telemetryErr)
-			}
-
 			key := filepath.Join(targetDir, unit.Path)
 			declaredUnits[key] = unit
 			outputs[key] = output
@@ -571,14 +553,7 @@ func generateUnits(ctx context.Context, l log.Logger, 
 
 			l.Infof("Generating unit %s from %s", unit.Name, sourceFile)
 
-			return telemetry.TelemeterFromContext(ctx).Collect(ctx, "stack_generate_unit", map[string]any{
-				"stack_file":  sourceFile,
-				"unit_name":   unit.Name,
-				"unit_source": unit.Source,
-				"unit_path":   unit.Path,
-			}, func(ctx context.Context) error {
 				return generateComponent(ctx, l, opts, &item)
-			})
 		})
 	}
 
@@ -604,14 +579,7 @@ func generateStacks(ctx context.Context, l log.Logger,
 
 			l.Infof("Generating stack %s from %s", stack.Name, sourceFile)
 
-			return telemetry.TelemeterFromContext(ctx).Collect(ctx, "stack_generate_stack", map[string]any{
-				"stack_file":   sourceFile,
-				"stack_name":   stack.Name,
-				"stack_source": stack.Source,
-				"stack_path":   stack.Path,
-			}, func(ctx context.Context) error {
 				return generateComponent(ctx, l, opts, &item)
-			})
 		})
 	}
 
