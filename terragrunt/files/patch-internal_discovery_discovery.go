--- internal/discovery/discovery.go.orig	1979-11-29 21:00:00 UTC
+++ internal/discovery/discovery.go
@@ -2,16 +2,14 @@ import (
 
 import (
 	"context"
+	"errors"
 	"path/filepath"
 	"slices"
 	"sync"
 
-	"errors"
-
 	"github.com/gruntwork-io/terragrunt/internal/component"
 	"github.com/gruntwork-io/terragrunt/internal/filter"
 	"github.com/gruntwork-io/terragrunt/internal/git"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
@@ -92,18 +90,9 @@ func (d *Discovery) Discover(
 		withWorktree,
 	)
 
-	err = telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "discovery_phase_filesystem", map[string]any{
-			"num_workers":   d.numWorkers,
-			"with_worktree": withWorktree,
-		}, func(childCtx context.Context, l log.Logger) error {
-			var phaseErr error
+	_, _ = d.numWorkers, withWorktree
+	results, err = d.runFilesystemPhase(ctx, l, v, opts)
 
-			results, phaseErr = d.runFilesystemPhase(childCtx, l, v, opts)
-
-			return phaseErr
-		})
-
 	logPhaseComplete(l, "filesystem", results, err)
 
 	if err != nil && (!d.suppressParseErrors || errors.As(err, new(CoexistenceError))) {
@@ -125,24 +114,9 @@ func (d *Discovery) Discover(
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
-				var phaseErr error
+		_, _, _, _, _, _, _, _ = d.numWorkers, len(discovered), len(candidates), d.parseIncludes, d.parseExclude, d.readFiles, d.trackReads, reasonsStr
+		results, err = d.runParsePhase(ctx, l, v, opts, discovered, candidates)
 
-				results, phaseErr = d.runParsePhase(childCtx, l, v, opts, discovered, candidates)
-
-				return phaseErr
-			})
-
 		logPhaseComplete(l, "parse", results, err)
 
 		if err != nil && !d.suppressParseErrors {
@@ -170,21 +144,9 @@ func (d *Discovery) Discover(
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
-				var phaseErr error
+		_, _, _, _, _ = d.numWorkers, d.maxDependencyDepth, len(discovered), len(candidates), d.classifier.HasDependentFilters()
+		results, err = d.runGraphPhase(ctx, l, v, opts, discovered, candidates)
 
-				results, phaseErr = d.runGraphPhase(childCtx, l, v, opts, discovered, candidates)
-
-				return phaseErr
-			})
-
 		logPhaseComplete(l, "graph", results, err)
 
 		if err != nil && !d.suppressParseErrors {
@@ -200,19 +162,9 @@ func (d *Discovery) Discover(
 		l.Debugf("Discovery: starting relationship phase (components=%d, max_depth=%d)",
 			len(components), d.maxDependencyDepth)
 
-		err = telemetry.TelemeterFromContext(ctx).
-			Collect(ctx, l, "discovery_phase_relationship", map[string]any{
-				"num_workers":          d.numWorkers,
-				"max_dependency_depth": d.maxDependencyDepth,
-				"components_in":        len(components),
-			}, func(childCtx context.Context, l log.Logger) error {
-				var phaseErr error
+		_, _, _ = d.numWorkers, d.maxDependencyDepth, len(components)
+		components, err = d.runRelationshipPhase(ctx, l, v, opts, components)
 
-				components, phaseErr = d.runRelationshipPhase(childCtx, l, v, opts, components)
-
-				return phaseErr
-			})
-
 		l.Debugf(
 			"Discovery: relationship phase complete (components=%d, err=%v)",
 			len(components),
@@ -247,26 +199,24 @@ func (d *Discovery) Discover(
 
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
@@ -279,14 +229,11 @@ func (d *Discovery) Discover(
 	components = d.applyQueueFilters(opts, components)
 
 	if d.parseStackConfigs {
-		if err := telemetry.TelemeterFromContext(ctx).
-			Collect(ctx, l, "discovery_phase_stack_configs", map[string]any{
-				"components_in": len(components),
-			}, func(childCtx context.Context, childL log.Logger) error {
-				storeStackConfigs(childCtx, childL, v, opts, components)
-
-				return nil
-			}); err != nil {
+		if err := func() error {
+			_ = len(components)
+			storeStackConfigs(ctx, l, v, opts, components)
+			return nil
+		}(); err != nil {
 			return components, err
 		}
 	}
@@ -337,24 +284,20 @@ func (d *Discovery) runFilesystemPhase(
 
 		l.Debugf("Discovery: starting filesystem walk at %s", d.workingDir)
 
-		err := telemetry.TelemeterFromContext(ctx).
-			Collect(ctx, l, "discovery_filesystem_walk", map[string]any{
-				"num_workers": d.numWorkers,
-				"working_dir": d.workingDir,
-			}, func(childCtx context.Context, l log.Logger) error {
-				phase := NewFilesystemPhase(d.numWorkers)
+		_, _ = d.numWorkers, d.workingDir
+		phase := NewFilesystemPhase(d.numWorkers)
+		err := func() error {
+			var phaseErr error
 
-				var phaseErr error
-
-				result, phaseErr = phase.Run(childCtx, l, v, &PhaseInput{
-					Opts:       opts,
-					Classifier: d.classifier,
-					Discovery:  d,
-				})
-
-				return phaseErr
+			result, phaseErr = phase.Run(ctx, l, v, &PhaseInput{
+				Opts:       opts,
+				Classifier: d.classifier,
+				Discovery:  d,
 			})
 
+			return phaseErr
+		}()
+
 		logPhaseComplete(l, "filesystem walk", result, err)
 
 		mu.Lock()
@@ -382,24 +325,20 @@ func (d *Discovery) runFilesystemPhase(
 				len(d.gitExpressions),
 			)
 
-			err := telemetry.TelemeterFromContext(ctx).
-				Collect(ctx, l, "discovery_worktree_walk", map[string]any{
-					"num_workers":          d.numWorkers,
-					"git_expression_count": len(d.gitExpressions),
-				}, func(childCtx context.Context, l log.Logger) error {
-					phase := NewWorktreePhase(d.gitExpressions, d.numWorkers)
+			_, _ = d.numWorkers, len(d.gitExpressions)
+			phase := NewWorktreePhase(d.gitExpressions, d.numWorkers)
+			err := func() error {
+				var phaseErr error
 
-					var phaseErr error
-
-					result, phaseErr = phase.Run(childCtx, l, v, &PhaseInput{
-						Opts:       opts,
-						Classifier: d.classifier,
-						Discovery:  d,
-					})
-
-					return phaseErr
+				result, phaseErr = phase.Run(ctx, l, v, &PhaseInput{
+					Opts:       opts,
+					Classifier: d.classifier,
+					Discovery:  d,
 				})
 
+				return phaseErr
+			}()
+
 			logPhaseComplete(l, "worktree walk", result, err)
 
 			mu.Lock()
@@ -491,11 +430,7 @@ func (d *Discovery) runGraphPhase(
 
 		unparsed := d.potentialDependentsOutsideBoundary(l, v.FS, candidates)
 
-		buildErr := telemetry.TelemeterFromContext(ctx).Collect(
-			ctx, l, "discover_dependents", map[string]any{},
-			func(childCtx context.Context, l log.Logger) error {
-				return errors.Join(d.buildDependencyGraph(childCtx, l, v, opts, allComponents, unparsed)...)
-			})
+		buildErr := errors.Join(d.buildDependencyGraph(ctx, l, v, opts, allComponents, unparsed)...)
 
 		if buildErr != nil && !d.suppressParseErrors {
 			return &PhaseResults{
@@ -509,21 +444,19 @@ func (d *Discovery) runGraphPhase(
 
 	var result *PhaseResults
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(
-		ctx, l, "discover_dependencies", map[string]any{},
-		func(childCtx context.Context, l log.Logger) error {
-			var runErr error
+	err := func() error {
+		var runErr error
 
-			result, runErr = phase.Run(childCtx, l, v, &PhaseInput{
-				Opts:       opts,
-				Components: resultsToComponents(discovered),
-				Candidates: candidates,
-				Classifier: d.classifier,
-				Discovery:  d,
-			})
-
-			return runErr
+		result, runErr = phase.Run(ctx, l, v, &PhaseInput{
+			Opts:       opts,
+			Components: resultsToComponents(discovered),
+			Candidates: candidates,
+			Classifier: d.classifier,
+			Discovery:  d,
 		})
+
+		return runErr
+	}()
 
 	allDiscovered := discovered
 	if result != nil {
