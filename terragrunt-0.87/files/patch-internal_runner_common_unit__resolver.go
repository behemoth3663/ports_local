--- internal/runner/common/unit_resolver.go.orig	2025-09-25 14:56:12 UTC
+++ internal/runner/common/unit_resolver.go
@@ -17,7 +17,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/options"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"github.com/gruntwork-io/terragrunt/shell"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/util"
 )
 
@@ -70,189 +69,35 @@ func (r *UnitResolver) ResolveTerraformModules(ctx con
 		return nil, err
 	}
 
-	unitsMap, err := r.telemetryResolveUnits(ctx, l, canonicalTerragruntConfigPaths)
+	howThesePathsWereFound := "Terragrunt config file found in a subdirectory of " + r.Stack.TerragruntOptions.WorkingDir
+	unitsMap, err := r.resolveUnits(ctx, l, canonicalTerragruntConfigPaths, howThesePathsWereFound)
 	if err != nil {
 		return nil, err
 	}
 
-	externalDependencies, err := r.telemetryResolveExternalDependencies(ctx, l, unitsMap)
+	externalDependencies, err := r.resolveExternalDependenciesForUnits(ctx, l, unitsMap, UnitsMap{}, 0)
 	if err != nil {
 		return nil, err
 	}
 
-	crossLinkedUnits, err := r.telemetryCrossLinkDependencies(ctx, unitsMap, externalDependencies, canonicalTerragruntConfigPaths)
+	crossLinkedUnits, err := unitsMap.MergeMaps(externalDependencies).CrossLinkDependencies(canonicalTerragruntConfigPaths)
 	if err != nil {
 		return nil, err
 	}
 
-	withUnitsIncluded, err := r.telemetryFlagIncludedDirs(ctx, l, crossLinkedUnits)
+	withUnitsIncluded := r.flagIncludedDirs(r.Stack.TerragruntOptions, l, crossLinkedUnits)
+	withUnitsThatAreIncludedByOthers, err := r.flagUnitsThatAreIncluded(r.Stack.TerragruntOptions, withUnitsIncluded)
 	if err != nil {
 		return nil, err
 	}
 
-	withUnitsThatAreIncludedByOthers, err := r.telemetryFlagUnitsThatAreIncluded(ctx, withUnitsIncluded)
-	if err != nil {
-		return nil, err
-	}
+	withExcludedUnits := r.flagExcludedUnits(l, r.Stack.TerragruntOptions, withUnitsThatAreIncludedByOthers)
+	withUnitsRead := r.flagUnitsThatRead(r.Stack.TerragruntOptions, withExcludedUnits)
+	withUnitsExcluded := r.flagExcludedDirs(l, r.Stack.TerragruntOptions, r.Stack.Report, withUnitsRead)
 
-	withExcludedUnits, err := r.telemetryFlagExcludedUnits(ctx, l, withUnitsThatAreIncludedByOthers)
-	if err != nil {
-		return nil, err
-	}
-
-	withUnitsRead, err := r.telemetryFlagUnitsThatRead(ctx, withExcludedUnits)
-	if err != nil {
-		return nil, err
-	}
-
-	withUnitsExcluded, err := r.telemetryFlagExcludedDirs(ctx, l, withUnitsRead)
-	if err != nil {
-		return nil, err
-	}
-
 	return withUnitsExcluded, nil
 }
 
-// telemetryResolveUnits resolves Terraform units from the given Terragrunt configuration paths
-func (r *UnitResolver) telemetryResolveUnits(ctx context.Context, l log.Logger, canonicalTerragruntConfigPaths []string) (UnitsMap, error) {
-	var unitsMap UnitsMap
-
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "resolve_units", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(ctx context.Context) error {
-		howThesePathsWereFound := "Terragrunt config file found in a subdirectory of " + r.Stack.TerragruntOptions.WorkingDir
-
-		result, err := r.resolveUnits(ctx, l, canonicalTerragruntConfigPaths, howThesePathsWereFound)
-		if err != nil {
-			return err
-		}
-
-		unitsMap = result
-
-		return nil
-	})
-
-	return unitsMap, err
-}
-
-// telemetryResolveExternalDependencies resolves external dependencies for the given units
-func (r *UnitResolver) telemetryResolveExternalDependencies(ctx context.Context, l log.Logger, unitsMap UnitsMap) (UnitsMap, error) {
-	var externalDependencies UnitsMap
-
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "resolve_external_dependencies_for_units", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(ctx context.Context) error {
-		result, err := r.resolveExternalDependenciesForUnits(ctx, l, unitsMap, UnitsMap{}, 0)
-		if err != nil {
-			return err
-		}
-
-		externalDependencies = result
-
-		return nil
-	})
-
-	return externalDependencies, err
-}
-
-// telemetryCrossLinkDependencies cross-links dependencies between units
-func (r *UnitResolver) telemetryCrossLinkDependencies(ctx context.Context, unitsMap, externalDependencies UnitsMap, canonicalTerragruntConfigPaths []string) (Units, error) {
-	var crossLinkedUnits Units
-
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "crosslink_dependencies", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
-		result, err := unitsMap.MergeMaps(externalDependencies).CrossLinkDependencies(canonicalTerragruntConfigPaths)
-		if err != nil {
-			return err
-		}
-
-		crossLinkedUnits = result
-
-		return nil
-	})
-
-	return crossLinkedUnits, err
-}
-
-// telemetryFlagIncludedDirs flags directories that are included in the Terragrunt configuration
-func (r *UnitResolver) telemetryFlagIncludedDirs(ctx context.Context, l log.Logger, crossLinkedUnits Units) (Units, error) {
-	var withUnitsIncluded Units
-
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "flag_included_dirs", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
-		withUnitsIncluded = r.flagIncludedDirs(r.Stack.TerragruntOptions, l, crossLinkedUnits)
-		return nil
-	})
-
-	return withUnitsIncluded, err
-}
-
-// telemetryFlagUnitsThatAreIncluded flags units that are included in the Terragrunt configuration
-func (r *UnitResolver) telemetryFlagUnitsThatAreIncluded(ctx context.Context, withUnitsIncluded Units) (Units, error) {
-	var withUnitsThatAreIncludedByOthers Units
-
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "flag_units_that_are_included", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
-		result, err := r.flagUnitsThatAreIncluded(r.Stack.TerragruntOptions, withUnitsIncluded)
-		if err != nil {
-			return err
-		}
-
-		withUnitsThatAreIncludedByOthers = result
-
-		return nil
-	})
-
-	return withUnitsThatAreIncludedByOthers, err
-}
-
-// telemetryFlagExcludedUnits flags units that are excluded in the Terragrunt configuration
-func (r *UnitResolver) telemetryFlagExcludedUnits(ctx context.Context, l log.Logger, withUnitsThatAreIncludedByOthers Units) (Units, error) {
-	var withExcludedUnits Units
-
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "flag_excluded_units", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
-		result := r.flagExcludedUnits(l, r.Stack.TerragruntOptions, withUnitsThatAreIncludedByOthers)
-		withExcludedUnits = result
-
-		return nil
-	})
-
-	return withExcludedUnits, err
-}
-
-// telemetryFlagUnitsThatRead flags units that read files in the Terragrunt configuration
-func (r *UnitResolver) telemetryFlagUnitsThatRead(ctx context.Context, withExcludedUnits Units) (Units, error) {
-	var withUnitsRead Units
-
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "flag_units_that_read", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
-		withUnitsRead = r.flagUnitsThatRead(r.Stack.TerragruntOptions, withExcludedUnits)
-		return nil
-	})
-
-	return withUnitsRead, err
-}
-
-// telemetryFlagExcludedDirs flags directories that are excluded in the Terragrunt configuration
-func (r *UnitResolver) telemetryFlagExcludedDirs(ctx context.Context, l log.Logger, withUnitsRead Units) (Units, error) {
-	var withUnitsExcluded Units
-
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "flag_excluded_dirs", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
-		withUnitsExcluded = r.flagExcludedDirs(l, r.Stack.TerragruntOptions, r.Stack.Report, withUnitsRead)
-		return nil
-	})
-
-	return withUnitsExcluded, err
-}
-
 // Go through each of the given Terragrunt configuration files and resolve the unit that configuration file represents
 // into a Unit struct. Note that this method will NOT fill in the Dependencies field of the Unit
 // struct (see the crosslinkDependencies method for that). Return a map from unit path to Unit struct.
@@ -264,21 +109,7 @@ func (r *UnitResolver) resolveUnits(ctx context.Contex
 			return nil, ProcessingUnitError{UnderlyingError: os.ErrNotExist, UnitPath: terragruntConfigPath, HowThisUnitWasFound: howTheseUnitsWereFound}
 		}
 
-		var unit *Unit
-
-		err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "resolve_terraform_unit", map[string]any{
-			"config_path": terragruntConfigPath,
-			"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-		}, func(ctx context.Context) error {
-			m, err := r.resolveTerraformUnit(ctx, l, terragruntConfigPath, unitsMap, howTheseUnitsWereFound)
-			if err != nil {
-				return err
-			}
-
-			unit = m
-
-			return nil
-		})
+		unit, err := r.resolveTerraformUnit(ctx, l, terragruntConfigPath, unitsMap, howTheseUnitsWereFound)
 		if err != nil {
 			return unitsMap, err
 		}
@@ -286,22 +117,7 @@ func (r *UnitResolver) resolveUnits(ctx context.Contex
 		if unit != nil {
 			unitsMap[unit.Path] = unit
 
-			var dependencies UnitsMap
-
-			err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "resolve_dependencies_for_unit", map[string]any{
-				"config_path": terragruntConfigPath,
-				"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-				"unit_path":   unit.Path,
-			}, func(ctx context.Context) error {
-				deps, err := r.resolveDependenciesForUnit(ctx, l, unit, unitsMap, true)
-				if err != nil {
-					return err
-				}
-
-				dependencies = deps
-
-				return nil
-			})
+			dependencies, err := r.resolveDependenciesForUnit(ctx, l, unit, unitsMap, true)
 			if err != nil {
 				return unitsMap, err
 			}
@@ -803,7 +619,7 @@ func (r *UnitResolver) flagExcludedDirs(l log.Logger, 
 // flagExcludedDirs iterates over a unit slice and flags all entries as excluded listed in the queue-exclude-dir CLI flag.
 func (r *UnitResolver) flagExcludedDirs(l log.Logger, opts *options.TerragruntOptions, reportInstance *report.Report, units Units) Units {
 	// If we don't have any excludes, we don't need to do anything.
-	if (len(r.excludeGlobs) == 0 && r.doubleStarEnabled) || len(opts.ExcludeDirs) == 0 {
+	if (r.doubleStarEnabled && len(r.excludeGlobs) == 0) || (!r.doubleStarEnabled && len(opts.ExcludeDirs) == 0) {
 		return units
 	}
 
