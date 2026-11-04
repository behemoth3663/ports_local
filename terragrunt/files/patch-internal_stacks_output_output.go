--- internal/stacks/output/output.go.orig	1979-11-29 21:00:00 UTC
+++ internal/stacks/output/output.go
@@ -15,7 +15,6 @@ import (
 
 	"github.com/gruntwork-io/terragrunt/internal/configbridge"
 	"github.com/gruntwork-io/terragrunt/internal/stacks/generate"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/worker"
 	"github.com/gruntwork-io/terragrunt/internal/worktrees"
@@ -387,17 +386,7 @@ func readUnitOutput(
 ) (map[string]cty.Value, error) {
 	var output map[string]cty.Value
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "unit_output", map[string]any{
-		"unit_name":   unit.Name,
-		"unit_source": unit.Source,
-		"unit_path":   unit.Path,
-	}, func(ctx context.Context, l log.Logger) error {
-		var outputErr error
-
-		output, outputErr = unit.ReadOutputs(ctx, l, v, pctx, unitDir)
-
-		return outputErr
-	})
+	output, err := unit.ReadOutputs(ctx, l, v, pctx, unitDir)
 	if err != nil {
 		return nil, UnitOutputError{UnitName: unit.Name, UnitDir: unitDir, Err: err}
 	}
