--- internal/runner/common/unit_resolver.go.orig	2025-11-10 13:46:52 UTC
+++ internal/runner/common/unit_resolver.go
@@ -39,7 +39,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/errors"
 	"github.com/gruntwork-io/terragrunt/internal/strict/controls"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/util"
 )
 
@@ -126,15 +125,13 @@ func (r *UnitResolver) ResolveFromDiscovery(ctx contex
 	// Discovery found all dependencies as Component interfaces, but runner needs concrete *Unit pointers
 	var crossLinkedUnits Units
 
-	err = telemetry.TelemeterFromContext(ctx).Collect(ctx, "convert_discovery_to_runner", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
+	err = func(_ context.Context) error {
 		var linkErr error
 
 		crossLinkedUnits, linkErr = unitsMap.ConvertDiscoveryToRunner(canonicalTerragruntConfigPaths)
 
 		return linkErr
-	})
+	}(ctx)
 	if err != nil {
 		return nil, err
 	}
@@ -190,10 +187,7 @@ func (r *UnitResolver) telemetryBuildUnitsFromDiscover
 func (r *UnitResolver) telemetryBuildUnitsFromDiscovery(ctx context.Context, l log.Logger, discovered []component.Component) (UnitsMap, error) {
 	var unitsMap UnitsMap
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "build_units_from_discovery", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-		"unit_count":  len(discovered),
-	}, func(ctx context.Context) error {
+	err := func(ctx context.Context) error {
 		result, err := r.buildUnitsFromDiscovery(l, discovered)
 		if err != nil {
 			return err
@@ -202,7 +196,7 @@ func (r *UnitResolver) telemetryBuildUnitsFromDiscover
 		unitsMap = result
 
 		return nil
-	})
+	}(ctx)
 
 	return unitsMap, err
 }
@@ -326,9 +320,7 @@ func (r *UnitResolver) telemetryFlagExternalDependenci
 // telemetryFlagExternalDependencies flags external dependencies and prompts user for confirmation.
 // Discovery has already found and parsed external dependencies, so this only handles user prompts.
 func (r *UnitResolver) telemetryFlagExternalDependencies(ctx context.Context, l log.Logger, unitsMap UnitsMap) error {
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, "flag_external_dependencies", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(ctx context.Context) error {
+	return func(ctx context.Context) error {
 		return r.flagExternalDependencies(ctx, l, unitsMap)
-	})
+	}(ctx)
 }
