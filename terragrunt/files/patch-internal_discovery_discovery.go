--- internal/discovery/discovery.go.orig	1979-11-29 21:00:00 UTC
+++ internal/discovery/discovery.go
@@ -10,7 +10,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/component"
 	"github.com/gruntwork-io/terragrunt/internal/filter"
 	"github.com/gruntwork-io/terragrunt/internal/git"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tips"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
@@ -111,21 +110,8 @@ func (d *Discovery) Discover(
 
 		var results *foundResults
 
-		err = telemetry.TelemeterFromContext(ctx).
-			Collect(ctx, l, "discovery_phase_graph", map[string]any{
-				"num_workers":           d.numWorkers,
-				"max_dependency_depth":  d.maxDependencyDepth,
-				"discovered_in":         len(discovered),
-				"candidates_in":         len(candidates),
-				"has_dependent_filters": d.classifier.HasDependentFilters(),
-			}, func(childCtx context.Context, l log.Logger) error {
-				var phaseErr error
+				results, err = d.traverseGraph(ctx, l, v, opts, found)
 
-				results, phaseErr = d.traverseGraph(childCtx, l, v, opts, found)
-
-				return phaseErr
-			})
-
 		logComplete(l, "graph traversal", results, err)
 
 		if d.returnsPhaseError(err) {
@@ -141,14 +127,7 @@ func (d *Discovery) Discover(
 		l.Debugf("Discovery: starting relationships (components=%d, max_depth=%d)",
 			len(components), d.maxDependencyDepth)
 
-		err = telemetry.TelemeterFromContext(ctx).
-			Collect(ctx, l, "discovery_phase_relationship", map[string]any{
-				"num_workers":          d.numWorkers,
-				"max_dependency_depth": d.maxDependencyDepth,
-				"components_in":        len(components),
-			}, func(childCtx context.Context, l log.Logger) error {
-				return found.relations.link(childCtx, l, v, components)
-			})
+		err = found.relations.link(ctx, l, v, components)
 
 		l.Debugf(
 			"Discovery: relationships complete (components=%d, err=%v)",
@@ -184,26 +163,24 @@ func (d *Discovery) Discover(
 
 	components = d.dropOutsideBoundary(l, v.FS, components)
 
-	cycleCheckErr := telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "discovery_cycle_check", map[string]any{},
-			func(childCtx context.Context, l log.Logger) error {
-				if _, cycleErr := components.CycleCheck(); cycleErr != nil {
-					l.Debugf("Cycle: %v", cycleErr)
+	cycleCheckErr := func() error {
+		if _, cycleErr := components.CycleCheck(); cycleErr != nil {
+			l.Debugf("Cycle: %v", cycleErr)
 
-					if d.breakCycles {
-						l.Warnf("Cycle detected in dependency graph, attempting removal of cycles.")
+			if d.breakCycles {
+				l.Warnf("Cycle detected in dependency graph, attempting removal of cycles.")
 
-						var removeErr error
+				var removeErr error
 
-						components, removeErr = removeCycles(components)
-						if removeErr != nil {
-							return removeErr
-						}
-					}
+				components, removeErr = removeCycles(components)
+				if removeErr != nil {
+					return removeErr
 				}
+			}
+		}
 
-				return nil
-			})
+		return nil
+	}()
 
 	if cycleCheckErr != nil && !d.suppressParseErrors {
 		return components, cycleCheckErr
@@ -216,16 +193,7 @@ func (d *Discovery) Discover(
 	components = d.applyQueueFilters(opts, components)
 
 	if d.parseStackConfigs {
-		if err := telemetry.TelemeterFromContext(ctx).
-			Collect(ctx, l, "discovery_phase_stack_configs", map[string]any{
-				"components_in": len(components),
-			}, func(childCtx context.Context, childL log.Logger) error {
-				storeStackConfigs(childCtx, childL, v, s, opts, components)
-
-				return nil
-			}); err != nil {
-			return components, err
-		}
+				storeStackConfigs(ctx, l, v, s, opts, components)
 	}
 
 	return components, nil
@@ -259,24 +227,8 @@ func (d *Discovery) walkAndParse(
 
 	l.Debugf("Discovery: starting pipeline (workers=%d, parse_reasons=%s)", d.numWorkers, joinParseReasons(parseReasons))
 
-	var found walked
+			found, err := d.runPipeline(ctx, l, v, s, opts, len(parseReasons) > 0)
 
-	err := telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "discovery_pipeline", map[string]any{
-			"num_workers":       d.numWorkers,
-			"parse_includes":    d.parseIncludes,
-			"parse_exclude":     d.parseExclude,
-			"read_files":        d.readFiles,
-			"track_reads":       d.trackReads,
-			"activation_reason": joinParseReasons(parseReasons),
-		}, func(childCtx context.Context, l log.Logger) error {
-			var pipelineErr error
-
-			found, pipelineErr = d.runPipeline(childCtx, l, v, s, opts, len(parseReasons) > 0)
-
-			return pipelineErr
-		})
-
 	logComplete(l, "pipeline", &foundResults{discovered: found.discovered, candidates: found.candidates}, err)
 
 	return found, err
@@ -320,18 +272,8 @@ func (d *Discovery) runWorktreeWalk(
 		len(d.gitExpressions),
 	)
 
-	err := telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "discovery_worktree_walk", map[string]any{
-			"num_workers":          d.numWorkers,
-			"git_expression_count": len(d.gitExpressions),
-		}, func(childCtx context.Context, l log.Logger) error {
-			var walkErr error
+			result, err := d.discoverInWorktrees(ctx, l, v, s, opts)
 
-			result, walkErr = d.discoverInWorktrees(childCtx, l, v, s, opts)
-
-			return walkErr
-		})
-
 	logComplete(l, "worktree walk", result, err)
 
 	return result, err
@@ -354,11 +296,7 @@ func (d *Discovery) traverseGraph(
 
 		unparsed := d.potentialDependentsOutsideBoundary(l, v.FS, candidates)
 
-		buildErr := telemetry.TelemeterFromContext(ctx).Collect(
-			ctx, l, "discover_dependents", map[string]any{},
-			func(childCtx context.Context, l log.Logger) error {
-				return errors.Join(found.relations.linkForDependents(childCtx, l, v, allComponents, unparsed)...)
-			})
+		buildErr := errors.Join(found.relations.linkForDependents(ctx, l, v, allComponents, unparsed)...)
 
 		if d.returnsPhaseError(buildErr) {
 			return &foundResults{discovered: discovered, candidates: candidates}, buildErr
@@ -369,15 +307,13 @@ func (d *Discovery) traverseGraph(
 
 	var result *foundResults
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(
-		ctx, l, "discover_dependencies", map[string]any{},
-		func(childCtx context.Context, l log.Logger) error {
+	err := func(childCtx context.Context, l log.Logger) error {
 			var runErr error
 
 			result, runErr = graph.run(childCtx, l, v, opts, d, resultsToComponents(discovered), candidates)
 
 			return runErr
-		})
+		}(ctx, l)
 
 	return &foundResults{
 		discovered: deduplicateResults(append(discovered, result.discovered...)),
