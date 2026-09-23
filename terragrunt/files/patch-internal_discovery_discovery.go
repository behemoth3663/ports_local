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
@@ -76,18 +74,8 @@ func (d *Discovery) Discover(
 		withWorktree,
 	)
 
-	err = telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "discovery_phase_filesystem", map[string]any{
-			"num_workers":   d.numWorkers,
-			"with_worktree": withWorktree,
-		}, func(childCtx context.Context, l log.Logger) error {
-			var phaseErr error
+	results, err = d.runFilesystemPhase(ctx, l, v, opts)
 
-			results, phaseErr = d.runFilesystemPhase(childCtx, l, v, opts)
-
-			return phaseErr
-		})
-
 	logPhaseComplete(l, "filesystem", results, err)
 
 	if err != nil && (!d.suppressParseErrors || errors.As(err, new(CoexistenceError))) {
@@ -109,24 +97,8 @@ func (d *Discovery) Discover(
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
+		results, err = d.runParsePhase(ctx, l, v, opts, discovered, candidates)
 
-				results, phaseErr = d.runParsePhase(childCtx, l, v, opts, discovered, candidates)
-
-				return phaseErr
-			})
-
 		logPhaseComplete(l, "parse", results, err)
 
 		if err != nil && !d.suppressParseErrors {
@@ -154,21 +126,8 @@ func (d *Discovery) Discover(
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
+		results, err = d.runGraphPhase(ctx, l, v, opts, discovered, candidates)
 
-				results, phaseErr = d.runGraphPhase(childCtx, l, v, opts, discovered, candidates)
-
-				return phaseErr
-			})
-
 		logPhaseComplete(l, "graph", results, err)
 
 		if err != nil && !d.suppressParseErrors {
@@ -184,19 +143,8 @@ func (d *Discovery) Discover(
 		l.Debugf("Discovery: starting relationship phase (components=%d, max_depth=%d)",
 			len(components), d.maxDependencyDepth)
 
-		err = telemetry.TelemeterFromContext(ctx).
-			Collect(ctx, l, "discovery_phase_relationship", map[string]any{
-				"num_workers":          d.numWorkers,
-				"max_dependency_depth": d.maxDependencyDepth,
-				"components_in":        len(components),
-			}, func(childCtx context.Context, l log.Logger) error {
-				var phaseErr error
+		components, err = d.runRelationshipPhase(ctx, l, v, opts, components)
 
-				components, phaseErr = d.runRelationshipPhase(childCtx, l, v, opts, components)
-
-				return phaseErr
-			})
-
 		l.Debugf(
 			"Discovery: relationship phase complete (components=%d, err=%v)",
 			len(components),
@@ -231,27 +179,22 @@ func (d *Discovery) Discover(
 
 	components = d.dropOutsideBoundary(l, v.FS, components)
 
-	cycleCheckErr := telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "discovery_cycle_check", map[string]any{},
-			func(childCtx context.Context, l log.Logger) error {
-				if _, cycleErr := components.CycleCheck(); cycleErr != nil {
-					l.Debugf("Cycle: %v", cycleErr)
+	_, cycleCheckErr := components.CycleCheck()
+			if cycleCheckErr != nil {
+				l.Debugf("Cycle: %v", cycleCheckErr)
 
-					if d.breakCycles {
-						l.Warnf("Cycle detected in dependency graph, attempting removal of cycles.")
+				if d.breakCycles {
+					l.Warnf("Cycle detected in dependency graph, attempting removal of cycles.")
 
-						var removeErr error
+					var removeErr error
 
-						components, removeErr = removeCycles(components)
-						if removeErr != nil {
-							return removeErr
-						}
+					components, removeErr = removeCycles(components)
+					if removeErr != nil {
+						return components, removeErr
 					}
 				}
+			}
 
-				return nil
-			})
-
 	if cycleCheckErr != nil && !d.suppressParseErrors {
 		return components, cycleCheckErr
 	}
@@ -263,16 +206,7 @@ func (d *Discovery) Discover(
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
-			return components, err
-		}
+		storeStackConfigs(ctx, l, v, opts, components)
 	}
 
 	return components, nil
@@ -317,28 +251,16 @@ func (d *Discovery) runFilesystemPhase(
 	g.SetLimit(maxPhases)
 
 	g.Go(func() error {
-		var result *PhaseResults
-
 		l.Debugf("Discovery: starting filesystem walk at %s", d.workingDir)
 
-		err := telemetry.TelemeterFromContext(ctx).
-			Collect(ctx, l, "discovery_filesystem_walk", map[string]any{
-				"num_workers": d.numWorkers,
-				"working_dir": d.workingDir,
-			}, func(childCtx context.Context, l log.Logger) error {
-				phase := NewFilesystemPhase(d.numWorkers)
+		phase := NewFilesystemPhase(d.numWorkers)
 
-				var phaseErr error
+		result, err := phase.Run(ctx, l, v, &PhaseInput{
+			Opts:       opts,
+			Classifier: d.classifier,
+			Discovery:  d,
+		})
 
-				result, phaseErr = phase.Run(childCtx, l, v, &PhaseInput{
-					Opts:       opts,
-					Classifier: d.classifier,
-					Discovery:  d,
-				})
-
-				return phaseErr
-			})
-
 		logPhaseComplete(l, "filesystem walk", result, err)
 
 		mu.Lock()
@@ -359,31 +281,19 @@ func (d *Discovery) runFilesystemPhase(
 
 	if len(d.gitExpressions) > 0 && d.worktrees != nil {
 		g.Go(func() error {
-			var result *PhaseResults
-
 			l.Debugf(
 				"Discovery: starting worktree walk (git_expressions=%d)",
 				len(d.gitExpressions),
 			)
 
-			err := telemetry.TelemeterFromContext(ctx).
-				Collect(ctx, l, "discovery_worktree_walk", map[string]any{
-					"num_workers":          d.numWorkers,
-					"git_expression_count": len(d.gitExpressions),
-				}, func(childCtx context.Context, l log.Logger) error {
-					phase := NewWorktreePhase(d.gitExpressions, d.numWorkers)
+			phase := NewWorktreePhase(d.gitExpressions, d.numWorkers)
 
-					var phaseErr error
+			result, err := phase.Run(ctx, l, v, &PhaseInput{
+				Opts:       opts,
+				Classifier: d.classifier,
+				Discovery:  d,
+			})
 
-					result, phaseErr = phase.Run(childCtx, l, v, &PhaseInput{
-						Opts:       opts,
-						Classifier: d.classifier,
-						Discovery:  d,
-					})
-
-					return phaseErr
-				})
-
 			logPhaseComplete(l, "worktree walk", result, err)
 
 			mu.Lock()
@@ -473,11 +383,7 @@ func (d *Discovery) runGraphPhase(
 		allComponents := resultsToComponents(discovered)
 		allComponents = append(allComponents, resultsToComponents(candidates)...)
 
-		buildErr := telemetry.TelemeterFromContext(ctx).Collect(
-			ctx, l, "discover_dependents", map[string]any{},
-			func(childCtx context.Context, l log.Logger) error {
-				return errors.Join(d.buildDependencyGraph(childCtx, l, v, opts, allComponents)...)
-			})
+		buildErr := errors.Join(d.buildDependencyGraph(ctx, l, v, opts, allComponents)...)
 
 		if buildErr != nil && !d.suppressParseErrors {
 			return &PhaseResults{
@@ -489,23 +395,13 @@ func (d *Discovery) runGraphPhase(
 
 	phase := NewGraphPhase(d.numWorkers, d.maxDependencyDepth)
 
-	var result *PhaseResults
-
-	err := telemetry.TelemeterFromContext(ctx).Collect(
-		ctx, l, "discover_dependencies", map[string]any{},
-		func(childCtx context.Context, l log.Logger) error {
-			var runErr error
-
-			result, runErr = phase.Run(childCtx, l, v, &PhaseInput{
-				Opts:       opts,
-				Components: resultsToComponents(discovered),
-				Candidates: candidates,
-				Classifier: d.classifier,
-				Discovery:  d,
-			})
-
-			return runErr
-		})
+	result, err := phase.Run(ctx, l, v, &PhaseInput{
+		Opts:       opts,
+		Components: resultsToComponents(discovered),
+		Candidates: candidates,
+		Classifier: d.classifier,
+		Discovery:  d,
+	})
 
 	allDiscovered := discovered
 	if result != nil {
