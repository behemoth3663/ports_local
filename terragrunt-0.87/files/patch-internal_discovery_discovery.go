--- internal/discovery/discovery.go.orig	2025-09-25 14:56:12 UTC
+++ internal/discovery/discovery.go
@@ -18,8 +18,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/experiment"
 	"github.com/gruntwork-io/terragrunt/util"
 
-	"github.com/gruntwork-io/terragrunt/telemetry"
-
 	"github.com/gruntwork-io/terragrunt/config"
 	"github.com/gruntwork-io/terragrunt/internal/errors"
 	"github.com/gruntwork-io/terragrunt/options"
@@ -886,12 +884,6 @@ func (d *Discovery) Discover(ctx context.Context, l lo
 	}
 
 	if d.discoverDependencies {
-		err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "discover_dependencies", map[string]any{
-			"working_dir":                    d.workingDir,
-			"config_count":                   len(cfgs),
-			"discover_external_dependencies": d.discoverExternalDependencies,
-			"max_dependency_depth":           d.maxDependencyDepth,
-		}, func(ctx context.Context) error {
 			dependencyDiscovery := NewDependencyDiscovery(cfgs, d.maxDependencyDepth)
 
 			if d.discoveryContext != nil {
@@ -924,16 +916,6 @@ func (d *Discovery) Discover(ctx context.Context, l lo
 
 			cfgs = dependencyDiscovery.cfgs
 
-			return nil
-		})
-		if err != nil {
-			return cfgs, errors.New(err)
-		}
-
-		err = telemetry.TelemeterFromContext(ctx).Collect(ctx, "discovery_cycle_check", map[string]any{
-			"working_dir":  d.workingDir,
-			"config_count": len(cfgs),
-		}, func(ctx context.Context) error {
 			if _, cycleErr := cfgs.CycleCheck(); cycleErr != nil {
 				l.Warnf("Cycle detected in dependency graph, attempting removal of cycles.")
 
@@ -946,12 +928,6 @@ func (d *Discovery) Discover(ctx context.Context, l lo
 					errs = append(errs, errors.New(removeErr))
 				}
 			}
-
-			return nil
-		})
-		if err != nil {
-			return cfgs, errors.New(err)
-		}
 	}
 
 	if len(errs) > 0 {
