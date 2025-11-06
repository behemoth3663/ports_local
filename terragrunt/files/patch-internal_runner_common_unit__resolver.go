--- internal/runner/common/unit_resolver.go.orig	2025-11-05 17:03:04 UTC
+++ internal/runner/common/unit_resolver.go
@@ -39,7 +39,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/errors"
 	"github.com/gruntwork-io/terragrunt/internal/strict/controls"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/util"
 )
 
@@ -187,10 +186,7 @@ func (r *UnitResolver) telemetryBuildUnitsFromDiscover
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
@@ -199,7 +195,7 @@ func (r *UnitResolver) telemetryBuildUnitsFromDiscover
 		unitsMap = result
 
 		return nil
-	})
+	}(ctx)
 
 	return unitsMap, err
 }
@@ -314,9 +310,7 @@ func (r *UnitResolver) telemetryResolveExternalDepende
 func (r *UnitResolver) telemetryResolveExternalDependencies(ctx context.Context, l log.Logger, unitsMap UnitsMap) (UnitsMap, error) {
 	var externalDependencies UnitsMap
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "resolve_external_dependencies_for_units", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(ctx context.Context) error {
+	err := func(ctx context.Context) error {
 		result, err := r.resolveExternalDependenciesForUnits(ctx, l, unitsMap, UnitsMap{}, 0)
 		if err != nil {
 			return err
@@ -325,7 +319,7 @@ func (r *UnitResolver) telemetryResolveExternalDepende
 		externalDependencies = result
 
 		return nil
-	})
+	}(ctx)
 
 	return externalDependencies, err
 }
