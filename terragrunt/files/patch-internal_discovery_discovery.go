--- internal/discovery/discovery.go.orig	2025-11-10 13:46:52 UTC
+++ internal/discovery/discovery.go
@@ -20,8 +20,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/shell"
 	"github.com/gruntwork-io/terragrunt/util"
 
-	"github.com/gruntwork-io/terragrunt/telemetry"
-
 	"github.com/gruntwork-io/terragrunt/config"
 	"github.com/gruntwork-io/terragrunt/internal/errors"
 	"github.com/gruntwork-io/terragrunt/options"
@@ -1099,13 +1097,7 @@ func (d *Discovery) Discover(
 
 		if shouldRunDependencyDiscovery {
 			g.Go(func() error {
-				return telemetry.TelemeterFromContext(ctx).Collect(ctx, "discover_dependencies", map[string]any{
-					"working_dir":                    d.workingDir,
-					"config_count":                   len(components),
-					"starting_component_count":       len(dependencyStartingComponents),
-					"discover_external_dependencies": d.discoverExternalDependencies,
-					"max_dependency_depth":           d.maxDependencyDepth,
-				}, func(ctx context.Context) error {
+				return func(ctx context.Context) error {
 					dependencyDiscovery := NewDependencyDiscovery(threadSafeComponents).
 						WithMaxDepth(d.maxDependencyDepth).
 						WithWorkingDir(d.workingDir).
@@ -1140,18 +1132,13 @@ func (d *Discovery) Discover(
 					}
 
 					return nil
-				})
+				}(ctx)
 			})
 		}
 
 		if shouldRunDependentDiscovery {
 			g.Go(func() error {
-				return telemetry.TelemeterFromContext(ctx).Collect(ctx, "discover_dependents", map[string]any{
-					"working_dir":              d.workingDir,
-					"config_count":             len(components),
-					"starting_component_count": len(dependentStartingComponents),
-					"max_dependency_depth":     d.maxDependencyDepth,
-				}, func(ctx context.Context) error {
+				return func(ctx context.Context) error {
 					dependentDiscovery := NewDependentDiscovery(threadSafeComponents).
 						WithMaxDepth(d.maxDependencyDepth).
 						WithWorkingDir(d.workingDir).
@@ -1201,7 +1188,7 @@ func (d *Discovery) Discover(
 					}
 
 					return nil
-				})
+				}(ctx)
 			})
 		}
 
@@ -1211,10 +1198,7 @@ func (d *Discovery) Discover(
 
 		components = threadSafeComponents.ToComponents()
 
-		err = telemetry.TelemeterFromContext(ctx).Collect(ctx, "discovery_cycle_check", map[string]any{
-			"working_dir":  d.workingDir,
-			"config_count": len(components),
-		}, func(ctx context.Context) error {
+		err = func(ctx context.Context) error {
 			if _, cycleErr := components.CycleCheck(); cycleErr != nil {
 				l.Warnf("Cycle detected in dependency graph, attempting removal of cycles.")
 
@@ -1231,7 +1215,7 @@ func (d *Discovery) Discover(
 			}
 
 			return nil
-		})
+		}(ctx)
 		if err != nil {
 			return components, errors.New(err)
 		}
