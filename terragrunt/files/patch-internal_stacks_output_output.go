--- internal/stacks/output/output.go.orig	2025-11-10 20:41:38 UTC
+++ internal/stacks/output/output.go
@@ -12,7 +12,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/stacks/generate"
 	"github.com/gruntwork-io/terragrunt/options"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/zclconf/go-cty/cty"
 )
 
@@ -96,17 +95,13 @@ func StackOutput(ctx context.Context, l log.Logger, op
 
 			var output map[string]cty.Value
 
-			telemetryErr := telemetry.TelemeterFromContext(ctx).Collect(ctx, "unit_output", map[string]any{
-				"unit_name":   unit.Name,
-				"unit_source": unit.Source,
-				"unit_path":   unit.Path,
-			}, func(ctx context.Context) error {
+			telemetryErr := func(ctx context.Context) error {
 				var outputErr error
 
 				output, outputErr = unit.ReadOutputs(ctx, l, opts, unitDir)
 
 				return outputErr
-			})
+			}(ctx)
 			if telemetryErr != nil {
 				return cty.NilVal, errors.New(telemetryErr)
 			}
