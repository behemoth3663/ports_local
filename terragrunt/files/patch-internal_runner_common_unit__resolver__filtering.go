--- internal/runner/common/unit_resolver_filtering.go.orig	2025-11-10 13:46:52 UTC
+++ internal/runner/common/unit_resolver_filtering.go
@@ -9,7 +9,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/report"
 	"github.com/gruntwork-io/terragrunt/options"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/util"
 )
 
@@ -109,12 +108,10 @@ func (r *UnitResolver) telemetryApplyIncludeDirs(ctx c
 func (r *UnitResolver) telemetryApplyIncludeDirs(ctx context.Context, l log.Logger, crossLinkedUnits Units) (Units, error) {
 	var withUnitsIncluded Units
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "apply_include_dirs", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
+	err := func(_ context.Context) error {
 		withUnitsIncluded = r.applyIncludeDirs(r.Stack.TerragruntOptions, l, crossLinkedUnits)
 		return nil
-	})
+	}(ctx)
 
 	return withUnitsIncluded, err
 }
@@ -157,12 +154,10 @@ func (r *UnitResolver) telemetryFlagUnitsThatRead(ctx 
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
@@ -223,12 +218,10 @@ func (r *UnitResolver) telemetryApplyExcludeDirs(ctx c
 func (r *UnitResolver) telemetryApplyExcludeDirs(ctx context.Context, l log.Logger, withUnitsRead Units) (Units, error) {
 	var withUnitsExcluded Units
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "apply_exclude_dirs", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
+	err := func(_ context.Context) error {
 		withUnitsExcluded = r.applyExcludeDirs(l, r.Stack.TerragruntOptions, r.Stack.Report, withUnitsRead)
 		return nil
-	})
+	}(ctx)
 
 	return withUnitsExcluded, err
 }
@@ -279,14 +272,12 @@ func (r *UnitResolver) telemetryApplyExcludeModules(ct
 func (r *UnitResolver) telemetryApplyExcludeModules(ctx context.Context, l log.Logger, withUnitsThatAreIncludedByOthers Units) (Units, error) {
 	var withExcludedUnits Units
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "apply_exclude_modules", map[string]any{
-		"working_dir": r.Stack.TerragruntOptions.WorkingDir,
-	}, func(_ context.Context) error {
+	err := func(_ context.Context) error {
 		result := r.applyExcludeModules(l, r.Stack.TerragruntOptions, r.Stack.Report, withUnitsThatAreIncludedByOthers)
 		withExcludedUnits = result
 
 		return nil
-	})
+	}(ctx)
 
 	return withExcludedUnits, err
 }
@@ -345,10 +336,7 @@ func (r *UnitResolver) telemetryApplyFilters(ctx conte
 
 	var filteredUnits Units
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "apply_unit_filters", map[string]any{
-		"working_dir":  r.Stack.TerragruntOptions.WorkingDir,
-		"filter_count": len(r.filters),
-	}, func(ctx context.Context) error {
+	err := func(ctx context.Context) error {
 		// Apply all filters in sequence
 		for _, filter := range r.filters {
 			if err := filter.Filter(ctx, units, r.Stack.TerragruntOptions); err != nil {
@@ -359,7 +347,7 @@ func (r *UnitResolver) telemetryApplyFilters(ctx conte
 		filteredUnits = units
 
 		return nil
-	})
+	}(ctx)
 
 	return filteredUnits, err
 }
