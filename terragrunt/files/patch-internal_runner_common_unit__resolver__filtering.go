--- internal/runner/common/unit_resolver_filtering.go.orig	2025-11-05 17:03:04 UTC
+++ internal/runner/common/unit_resolver_filtering.go
@@ -9,7 +9,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/report"
 	"github.com/gruntwork-io/terragrunt/options"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/util"
 )
 
@@ -114,12 +113,10 @@ func (r *UnitResolver) telemetryFlagIncludedDirs(ctx c
 func (r *UnitResolver) telemetryFlagIncludedDirs(ctx context.Context, l log.Logger, crossLinkedUnits Units) (Units, error) {
 	var withUnitsIncluded Units
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "flag_included_dirs", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
+	err := func(_ context.Context) error {
 		withUnitsIncluded = r.flagIncludedDirs(r.Stack.TerragruntOptions, l, crossLinkedUnits)
 		return nil
-	})
+	}(ctx)
 
 	return withUnitsIncluded, err
 }
@@ -161,9 +158,7 @@ func (r *UnitResolver) telemetryFlagUnitsThatAreInclud
 func (r *UnitResolver) telemetryFlagUnitsThatAreIncluded(ctx context.Context, withUnitsIncluded Units) (Units, error) {
 	var withUnitsThatAreIncludedByOthers Units
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "flag_units_that_are_included", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
+	err := func(_ context.Context) error {
 		result, err := r.flagUnitsThatAreIncluded(r.Stack.TerragruntOptions, withUnitsIncluded)
 		if err != nil {
 			return err
@@ -172,7 +167,7 @@ func (r *UnitResolver) telemetryFlagUnitsThatAreInclud
 		withUnitsThatAreIncludedByOthers = result
 
 		return nil
-	})
+	}(ctx)
 
 	return withUnitsThatAreIncludedByOthers, err
 }
@@ -266,12 +261,10 @@ func (r *UnitResolver) telemetryFlagUnitsThatRead(ctx 
 func (r *UnitResolver) telemetryFlagUnitsThatRead(ctx context.Context, withExcludedUnits Units) (Units, error) {
 	var withUnitsRead Units
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "flag_units_that_read", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
+	err := func(_ context.Context) error {
 		withUnitsRead = r.flagUnitsThatRead(r.Stack.TerragruntOptions, withExcludedUnits)
 		return nil
-	})
+	}(ctx)
 
 	return withUnitsRead, err
 }
@@ -304,12 +297,10 @@ func (r *UnitResolver) telemetryFlagExcludedDirs(ctx c
 func (r *UnitResolver) telemetryFlagExcludedDirs(ctx context.Context, l log.Logger, withUnitsRead Units) (Units, error) {
 	var withUnitsExcluded Units
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "flag_excluded_dirs", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
+	err := func(_ context.Context) error {
 		withUnitsExcluded = r.flagExcludedDirs(l, r.Stack.TerragruntOptions, r.Stack.Report, withUnitsRead)
 		return nil
-	})
+	}(ctx)
 
 	return withUnitsExcluded, err
 }
@@ -360,14 +351,12 @@ func (r *UnitResolver) telemetryFlagExcludedUnits(ctx 
 func (r *UnitResolver) telemetryFlagExcludedUnits(ctx context.Context, l log.Logger, withUnitsThatAreIncludedByOthers Units) (Units, error) {
 	var withExcludedUnits Units
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "flag_excluded_units", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
+	err := func(_ context.Context) error {
 		result := r.flagExcludedUnits(l, r.Stack.TerragruntOptions, r.Stack.Report, withUnitsThatAreIncludedByOthers)
 		withExcludedUnits = result
 
 		return nil
-	})
+	}(ctx)
 
 	return withExcludedUnits, err
 }
@@ -426,10 +415,7 @@ func (r *UnitResolver) telemetryApplyFilters(ctx conte
 
 	var filteredUnits Units
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "apply_unit_filters", map[string]any{
-		"working_dir":  r.Stack.TerragruntOptions.WorkingDir,
-		"filter_count": len(r.filters),
-	}, func(ctx context.Context) error {
+	err := func(ctx context.Context) error {
 		// Apply all filters in sequence
 		for _, filter := range r.filters {
 			if err := filter.Filter(ctx, units, r.Stack.TerragruntOptions); err != nil {
@@ -440,7 +426,7 @@ func (r *UnitResolver) telemetryApplyFilters(ctx conte
 		filteredUnits = units
 
 		return nil
-	})
+	}(ctx)
 
 	return filteredUnits, err
 }
