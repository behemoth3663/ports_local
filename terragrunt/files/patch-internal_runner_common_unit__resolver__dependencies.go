--- internal/runner/common/unit_resolver_dependencies.go.orig	2025-11-05 17:03:04 UTC
+++ internal/runner/common/unit_resolver_dependencies.go
@@ -9,7 +9,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/options"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"github.com/gruntwork-io/terragrunt/shell"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/util"
 )
 
@@ -144,9 +143,7 @@ func (r *UnitResolver) telemetryCrossLinkDependencies(
 func (r *UnitResolver) telemetryCrossLinkDependencies(ctx context.Context, unitsMap, externalDependencies UnitsMap, canonicalTerragruntConfigPaths []string) (Units, error) {
 	var crossLinkedUnits Units
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "crosslink_dependencies", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
+	err := func(_ context.Context) error {
 		result, err := unitsMap.MergeMaps(externalDependencies).CrossLinkDependencies(canonicalTerragruntConfigPaths)
 		if err != nil {
 			return err
@@ -155,7 +152,7 @@ func (r *UnitResolver) telemetryCrossLinkDependencies(
 		crossLinkedUnits = result
 
 		return nil
-	})
+	}(ctx)
 
 	return crossLinkedUnits, err
 }
