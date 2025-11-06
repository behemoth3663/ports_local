--- internal/discovery/discovery.go.orig	2025-11-05 17:03:04 UTC
+++ internal/discovery/discovery.go
@@ -19,8 +19,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/filter"
 	"github.com/gruntwork-io/terragrunt/util"
 
-	"github.com/gruntwork-io/terragrunt/telemetry"
-
 	"github.com/gruntwork-io/terragrunt/config"
 	"github.com/gruntwork-io/terragrunt/internal/errors"
 	"github.com/gruntwork-io/terragrunt/options"
@@ -987,12 +985,7 @@ func (d *Discovery) Discover(
 	}
 
 	if d.discoverDependencies {
-		err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "discover_dependencies", map[string]any{
-			"working_dir":                    d.workingDir,
-			"config_count":                   len(components),
-			"discover_external_dependencies": d.discoverExternalDependencies,
-			"max_dependency_depth":           d.maxDependencyDepth,
-		}, func(ctx context.Context) error {
+		err := func(ctx context.Context) error {
 			dependencyDiscovery := NewDependencyDiscovery(components, d.maxDependencyDepth)
 
 			if d.discoveryContext != nil {
@@ -1026,15 +1019,12 @@ func (d *Discovery) Discover(
 			components = dependencyDiscovery.components
 
 			return nil
-		})
+		}(ctx)
 		if err != nil {
 			return components, errors.New(err)
 		}
 
-		err = telemetry.TelemeterFromContext(ctx).Collect(ctx, "discovery_cycle_check", map[string]any{
-			"working_dir":  d.workingDir,
-			"config_count": len(components),
-		}, func(ctx context.Context) error {
+		err = func(ctx context.Context) error {
 			if _, cycleErr := components.CycleCheck(); cycleErr != nil {
 				l.Warnf("Cycle detected in dependency graph, attempting removal of cycles.")
 
@@ -1049,7 +1039,7 @@ func (d *Discovery) Discover(
 			}
 
 			return nil
-		})
+		}(ctx)
 		if err != nil {
 			return components, errors.New(err)
 		}
