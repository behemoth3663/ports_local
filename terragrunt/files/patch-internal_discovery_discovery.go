--- internal/discovery/discovery.go.orig	1979-11-29 21:00:00 UTC
+++ internal/discovery/discovery.go
@@ -11,7 +11,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/component"
 	"github.com/gruntwork-io/terragrunt/internal/filter"
 	"github.com/gruntwork-io/terragrunt/internal/git"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
@@ -92,17 +91,13 @@ func (d *Discovery) Discover(
 		withWorktree,
 	)
 
-	err = telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "discovery_phase_filesystem", map[string]any{
-			"num_workers":   d.numWorkers,
-			"with_worktree": withWorktree,
-		}, func(childCtx context.Context, l log.Logger) error {
+	err = func(childCtx context.Context, l log.Logger) error {
 			var phaseErr error
 
 			results, phaseErr = d.runFilesystemPhase(childCtx, l, v, opts)
 
 			return phaseErr
-		})
+		}(ctx, l)
 
 	logPhaseComplete(l, "filesystem", results, err)
 
@@ -125,23 +120,13 @@ func (d *Discovery) Discover(
 		l.Debugf("Discovery: starting parse phase (discovered=%d, candidates=%d, reasons=%s)",
 			len(discovered), len(candidates), reasonsStr)
 
-		err = telemetry.TelemeterFromContext(ctx).
-			Collect(ctx, l, "discovery_phase_parse", map[string]any{
-				"num_workers":       d.numWorkers,
-				"discovered_in":     len(discovered),
-				"candidates_in":     len(candidates),
-				"parse_includes":    d.parseIncludes,
-				"parse_exclude":     d.parseExclude,
-				"read_files":        d.readFiles,
-				"track_reads":       d.trackReads,
-				"activation_reason": reasonsStr,
-			}, func(childCtx context.Context, l log.Logger) error {
+		err = func(childCtx context.Context, l log.Logger) error {
 				var phaseErr error
 
 				results, phaseErr = d.runParsePhase(childCtx, l, v, opts, discovered, candidates)
 
 				return phaseErr
-			})
+			}(ctx, l)
 
 		logPhaseComplete(l, "parse", results, err)
 
@@ -170,20 +155,13 @@ func (d *Discovery) Discover(
 			d.classifier.HasDependentFilters(),
 		)
 
-		err = telemetry.TelemeterFromContext(ctx).
-			Collect(ctx, l, "discovery_phase_graph", map[string]any{
-				"num_workers":           d.numWorkers,
-				"max_dependency_depth":  d.maxDependencyDepth,
-				"discovered_in":         len(discovered),
-				"candidates_in":         len(candidates),
-				"has_dependent_filters": d.classifier.HasDependentFilters(),
-			}, func(childCtx context.Context, l log.Logger) error {
+		err = func(childCtx context.Context, l log.Logger) error {
 				var phaseErr error
 
 				results, phaseErr = d.runGraphPhase(childCtx, l, v, opts, discovered, candidates)
 
 				return phaseErr
-			})
+			}(ctx, l)
 
 		logPhaseComplete(l, "graph", results, err)
 
@@ -200,18 +178,13 @@ func (d *Discovery) Discover(
 		l.Debugf("Discovery: starting relationship phase (components=%d, max_depth=%d)",
 			len(components), d.maxDependencyDepth)
 
-		err = telemetry.TelemeterFromContext(ctx).
-			Collect(ctx, l, "discovery_phase_relationship", map[string]any{
-				"num_workers":          d.numWorkers,
-				"max_dependency_depth": d.maxDependencyDepth,
-				"components_in":        len(components),
-			}, func(childCtx context.Context, l log.Logger) error {
+		err = func(childCtx context.Context, l log.Logger) error {
 				var phaseErr error
 
 				components, phaseErr = d.runRelationshipPhase(childCtx, l, v, opts, components)
 
 				return phaseErr
-			})
+			}(ctx, l)
 
 		l.Debugf(
 			"Discovery: relationship phase complete (components=%d, err=%v)",
@@ -247,9 +220,7 @@ func (d *Discovery) Discover(
 
 	components = d.dropOutsideBoundary(l, v.FS, components)
 
-	cycleCheckErr := telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "discovery_cycle_check", map[string]any{},
-			func(childCtx context.Context, l log.Logger) error {
+	cycleCheckErr := func(childCtx context.Context, l log.Logger) error {
 				if _, cycleErr := components.CycleCheck(); cycleErr != nil {
 					l.Debugf("Cycle: %v", cycleErr)
 
@@ -266,7 +237,7 @@ func (d *Discovery) Discover(
 				}
 
 				return nil
-			})
+			}(ctx, l)
 
 	if cycleCheckErr != nil && !d.suppressParseErrors {
 		return components, cycleCheckErr
@@ -279,14 +250,11 @@ func (d *Discovery) Discover(
 	components = d.applyQueueFilters(opts, components)
 
 	if d.parseStackConfigs {
-		if err := telemetry.TelemeterFromContext(ctx).
-			Collect(ctx, l, "discovery_phase_stack_configs", map[string]any{
-				"components_in": len(components),
-			}, func(childCtx context.Context, childL log.Logger) error {
+		if err := func(childCtx context.Context, childL log.Logger) error {
 				storeStackConfigs(childCtx, childL, v, opts, components)
 
 				return nil
-			}); err != nil {
+			}(ctx, l); err != nil {
 			return components, err
 		}
 	}
@@ -337,11 +305,7 @@ func (d *Discovery) runFilesystemPhase(
 
 		l.Debugf("Discovery: starting filesystem walk at %s", d.workingDir)
 
-		err := telemetry.TelemeterFromContext(ctx).
-			Collect(ctx, l, "discovery_filesystem_walk", map[string]any{
-				"num_workers": d.numWorkers,
-				"working_dir": d.workingDir,
-			}, func(childCtx context.Context, l log.Logger) error {
+		err := func(childCtx context.Context, l log.Logger) error {
 				phase := NewFilesystemPhase(d.numWorkers)
 
 				var phaseErr error
@@ -353,7 +317,7 @@ func (d *Discovery) runFilesystemPhase(
 				})
 
 				return phaseErr
-			})
+			}(ctx, l)
 
 		logPhaseComplete(l, "filesystem walk", result, err)
 
@@ -382,11 +346,7 @@ func (d *Discovery) runFilesystemPhase(
 				len(d.gitExpressions),
 			)
 
-			err := telemetry.TelemeterFromContext(ctx).
-				Collect(ctx, l, "discovery_worktree_walk", map[string]any{
-					"num_workers":          d.numWorkers,
-					"git_expression_count": len(d.gitExpressions),
-				}, func(childCtx context.Context, l log.Logger) error {
+			err := func(childCtx context.Context, l log.Logger) error {
 					phase := NewWorktreePhase(d.gitExpressions, d.numWorkers)
 
 					var phaseErr error
@@ -398,7 +358,7 @@ func (d *Discovery) runFilesystemPhase(
 					})
 
 					return phaseErr
-				})
+				}(ctx, l)
 
 			logPhaseComplete(l, "worktree walk", result, err)
 
@@ -491,11 +451,9 @@ func (d *Discovery) runGraphPhase(
 
 		unparsed := d.potentialDependentsOutsideBoundary(l, v.FS, candidates)
 
-		buildErr := telemetry.TelemeterFromContext(ctx).Collect(
-			ctx, l, "discover_dependents", map[string]any{},
-			func(childCtx context.Context, l log.Logger) error {
-				return errors.Join(d.buildDependencyGraph(childCtx, l, v, opts, allComponents, unparsed)...)
-			})
+		buildErr := func(childCtx context.Context, l log.Logger) error {
+			return errors.Join(d.buildDependencyGraph(childCtx, l, v, opts, allComponents, unparsed)...)
+		}(ctx, l)
 
 		if buildErr != nil && !d.suppressParseErrors {
 			return &PhaseResults{
@@ -509,9 +467,7 @@ func (d *Discovery) runGraphPhase(
 
 	var result *PhaseResults
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(
-		ctx, l, "discover_dependencies", map[string]any{},
-		func(childCtx context.Context, l log.Logger) error {
+	err := func(childCtx context.Context, l log.Logger) error {
 			var runErr error
 
 			result, runErr = phase.Run(childCtx, l, v, &PhaseInput{
@@ -523,7 +479,7 @@ func (d *Discovery) runGraphPhase(
 			})
 
 			return runErr
-		})
+		}(ctx, l)
 
 	allDiscovered := discovered
 	if result != nil {
