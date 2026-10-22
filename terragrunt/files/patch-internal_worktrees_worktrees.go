--- internal/worktrees/worktrees.go.orig	1979-11-29 21:00:00 UTC
+++ internal/worktrees/worktrees.go
@@ -4,6 +4,7 @@ import (
 
 import (
 	"context"
+	"errors"
 	"fmt"
 	"io/fs"
 	"os"
@@ -15,14 +16,11 @@ import (
 	"sync"
 	"time"
 
-	"errors"
-
 	"github.com/gruntwork-io/terragrunt/internal/component"
 	"github.com/gruntwork-io/terragrunt/internal/experiment"
 	"github.com/gruntwork-io/terragrunt/internal/filter"
 	"github.com/gruntwork-io/terragrunt/internal/git"
 	"github.com/gruntwork-io/terragrunt/internal/spinner"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
@@ -129,67 +127,54 @@ func (w *Worktrees) Cleanup(ctx context.Context, l log
 
 	gitRunner = gitRunner.WithWorkDir(w.OriginalWorkingDir)
 
-	return filter.TraceGitWorktreesCleanup(
-		ctx,
-		len(w.WorktreePairs),
-		gitRunner.GetRemoteURL(ctx),
-		func(ctx context.Context) error {
-			g, groupCtx := errgroup.WithContext(ctx)
-			// Every worktree was created under the same temporary directory, so
-			// the first one's filesystem sizes removal for all of them.
-			g.SetLimit(min(vfs.FSWorkersFor(v.FS, toRemove[0].Path), len(toRemove)))
+	_, _ = len(w.WorktreePairs), gitRunner.GetRemoteURL(ctx)
+	g, groupCtx := errgroup.WithContext(ctx)
+	// Every worktree was created under the same temporary directory, so
+	// the first one's filesystem sizes removal for all of them.
+	g.SetLimit(min(vfs.FSWorkersFor(v.FS, toRemove[0].Path), len(toRemove)))
 
-			for _, worktree := range toRemove {
-				g.Go(func() error {
-					// Skip removal if the worktree path doesn't exist (may have been cleaned up already)
-					if _, err := v.FS.Stat(worktree.Path); errors.Is(err, fs.ErrNotExist) {
-						l.Debugf(
-							"Worktree path %s already removed, skipping cleanup",
-							worktree.Path,
-						)
+	for _, worktree := range toRemove {
+		g.Go(func() error {
+			// Skip removal if the worktree path doesn't exist (may have been cleaned up already)
+			if _, err := v.FS.Stat(worktree.Path); errors.Is(err, fs.ErrNotExist) {
+				l.Debugf(
+					"Worktree path %s already removed, skipping cleanup",
+					worktree.Path,
+				)
 
-						return nil
-					}
+				return nil
+			}
 
-					err := filter.TraceGitWorktreeRemove(
-						groupCtx,
+			err := gitRunner.RemoveWorktree(groupCtx, worktree.Path)
+			if err != nil {
+				// If the error is due to the worktree not existing, log and continue
+				// This can happen during parallel test execution or if cleanup runs twice
+				errStr := err.Error()
+				if strings.Contains(errStr, "No such file or directory") ||
+					strings.Contains(errStr, "does not exist") ||
+					strings.Contains(errStr, "not a valid directory") {
+					l.Debugf(
+						"Worktree for reference %s already cleaned up: %v",
 						worktree.Ref,
-						worktree.Path,
-						func(ctx context.Context) error {
-							return gitRunner.RemoveWorktree(ctx, worktree.Path)
-						},
+						err,
 					)
-					if err != nil {
-						// If the error is due to the worktree not existing, log and continue
-						// This can happen during parallel test execution or if cleanup runs twice
-						errStr := err.Error()
-						if strings.Contains(errStr, "No such file or directory") ||
-							strings.Contains(errStr, "does not exist") ||
-							strings.Contains(errStr, "not a valid directory") {
-							l.Debugf(
-								"Worktree for reference %s already cleaned up: %v",
-								worktree.Ref,
-								err,
-							)
 
-							return nil
-						}
-
-						return fmt.Errorf(
-							"failed to remove Git worktree for reference %s (%s): %w",
-							worktree.Ref,
-							worktree.Path,
-							err,
-						)
-					}
-
 					return nil
-				})
+				}
+
+				return fmt.Errorf(
+					"failed to remove Git worktree for reference %s (%s): %w",
+					worktree.Ref,
+					worktree.Path,
+					err,
+				)
 			}
 
-			return g.Wait()
-		},
-	)
+			return nil
+		})
+	}
+
+	return g.Wait()
 }
 
 // worktreesToRemove returns the worktrees Cleanup has to delete, one entry per
@@ -655,16 +640,9 @@ func surveyGitExpressions(
 		g.Go(func() error {
 			var expressionDiffs *git.Diffs
 
-			diffErr := filter.TraceGitDiff(
-				gCtx, gitExpression.FromRef, gitExpression.ToRef, repoRemote,
-				func(ctx context.Context) error {
-					var err error
+			_, _, _ = gitExpression.FromRef, gitExpression.ToRef, repoRemote
+			expressionDiffs, diffErr := gitRunner.Diff(gCtx, gitExpression.FromRef, gitExpression.ToRef)
 
-					expressionDiffs, err = gitRunner.Diff(ctx, gitExpression.FromRef, gitExpression.ToRef)
-
-					return err
-				})
-
 			mu.Lock()
 			defer mu.Unlock()
 
@@ -754,62 +732,61 @@ func NewWorktrees(
 	repoBranch := gitRunner.GetCurrentBranch(ctx)
 	repoCommit := gitRunner.GetHeadCommit(ctx)
 
-	// Wrap entire worktree creation process with telemetry
-	traceErr := filter.TraceGitWorktreesCreate(
-		ctx, workingDir, len(gitRefs), repoRemote, repoBranch, repoCommit,
-		func(ctx context.Context) error {
-			// The survey decides which references get a worktree and how much
-			// of each one is checked out, so it runs before any of them are.
-			survey, err := surveyGitExpressions(ctx, v, gitRunner, gitExpressions, gitRefs, repoRemote)
-			if err != nil {
-				worktrees = newEmptyWorktrees(workingDir)
-				outerErr = err
+	_ = workingDir
+	_ = len(gitRefs)
+	traceErr := func() error {
+		// The survey decides which references get a worktree and how much
+		// of each one is checked out, so it runs before any of them are.
+		survey, err := surveyGitExpressions(ctx, v, gitRunner, gitExpressions, gitRefs, repoRemote)
+		if err != nil {
+			worktrees = newEmptyWorktrees(workingDir)
+			outerErr = err
 
-				return err
-			}
+			return err
+		}
 
-			// A symlink in a checked-out directory can point anywhere in the
-			// reference, so discovery that follows one walks out of the
-			// directories the filters name.
-			filteredPaths := opts.FilteredPathsOnly && !experiments.Evaluate(experiment.Symlinks)
+		// A symlink in a checked-out directory can point anywhere in the
+		// reference, so discovery that follows one walks out of the
+		// directories the filters name.
+		filteredPaths := opts.FilteredPathsOnly && !experiments.Evaluate(experiment.Symlinks)
 
-			pathspecs := survey.checkoutPathspecs(l, gitExpressions, filteredPaths)
+		pathspecs := survey.checkoutPathspecs(l, gitExpressions, filteredPaths)
 
-			refsToPaths, err := createGitWorktrees(
-				ctx,
-				l,
-				v,
-				gitRunner,
-				survey.refsToMaterialize(l, gitExpressions, gitRefs),
-				repoRemote,
-				repoBranch,
-				repoCommit,
-				experiments,
-				pathspecs,
-			)
+		refsToPaths, err := createGitWorktrees(
+			ctx,
+			l,
+			v,
+			gitRunner,
+			survey.refsToMaterialize(l, gitExpressions, gitRefs),
+			repoRemote,
+			repoBranch,
+			repoCommit,
+			experiments,
+			pathspecs,
+		)
 
-			worktrees = &Worktrees{
-				WorktreePairs:      survey.worktreePairs(gitExpressions, refsToPaths),
-				OriginalWorkingDir: workingDir,
-			}
+		worktrees = &Worktrees{
+			WorktreePairs:      survey.worktreePairs(gitExpressions, refsToPaths),
+			OriginalWorkingDir: workingDir,
+		}
 
-			if err != nil {
-				outerErr = err
+		if err != nil {
+			outerErr = err
 
-				return err
-			}
+			return err
+		}
 
-			for _, gitExpression := range gitExpressions {
-				if diffs := survey.expansions[gitExpression.String()].diffs; diffs != nil {
-					recordDiffTelemetry(ctx, diffs)
-				}
+		for _, gitExpression := range gitExpressions {
+			if diffs := survey.expansions[gitExpression.String()].diffs; diffs != nil {
+				recordDiffTelemetry(ctx, diffs)
 			}
+		}
 
-			return nil
-		})
+		return nil
+	}()
 
 	if traceErr != nil && outerErr == nil {
-		l.Warnf("telemetry trace error during worktree creation: %v", traceErr)
+		l.Warnf("worktree creation wrapper failed: %v", traceErr)
 	}
 
 	// cleanup worktrees
@@ -924,14 +901,7 @@ func recordDiffTelemetry(ctx context.Context, diffs *g
 
 // recordDiffTelemetry records telemetry metrics for git diff results.
 func recordDiffTelemetry(ctx context.Context, diffs *git.Diffs) {
-	telemeter := telemetry.TelemeterFromContext(ctx)
-	if telemeter == nil || telemeter.Meter == nil {
-		return
-	}
-
-	telemeter.Count(ctx, "git_diff_files_added", int64(len(diffs.Added)))
-	telemeter.Count(ctx, "git_diff_files_removed", int64(len(diffs.Removed)))
-	telemeter.Count(ctx, "git_diff_files_changed", int64(len(diffs.Changed)))
+	_, _ = ctx, diffs
 }
 
 // createGitWorktrees creates detached worktrees for each unique Git reference needed by filters.
@@ -1064,11 +1034,8 @@ func createGitWorktree(
 		return "", err
 	}
 
-	err = filter.TraceGitWorktreeCreate(
-		ctx, opts.ref, tmpDir, opts.repoRemote, opts.repoBranch, opts.repoCommit,
-		func(ctx context.Context) error {
-			return materializeGitWorktree(ctx, l, v, gitRunner, registerMu, tmpDir, opts)
-		})
+	_, _, _, _ = opts.ref, opts.repoRemote, opts.repoBranch, opts.repoCommit
+	err = materializeGitWorktree(ctx, l, v, gitRunner, registerMu, tmpDir, opts)
 	if err != nil {
 		if cleanErr := v.FS.RemoveAll(tmpDir); cleanErr != nil {
 			l.Warnf("failed to clean worktree directory %s: %v", tmpDir, cleanErr)
