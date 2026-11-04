--- internal/worktrees/worktrees.go.orig	1979-11-29 21:00:00 UTC
+++ internal/worktrees/worktrees.go
@@ -22,7 +22,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/filter"
 	"github.com/gruntwork-io/terragrunt/internal/git"
 	"github.com/gruntwork-io/terragrunt/internal/spinner"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
@@ -129,11 +128,7 @@ func (w *Worktrees) Cleanup(ctx context.Context, l log
 
 	gitRunner = gitRunner.WithWorkDir(w.OriginalWorkingDir)
 
-	return filter.TraceGitWorktreesCleanup(
-		ctx,
-		len(w.WorktreePairs),
-		gitRunner.GetRemoteURL(ctx),
-		func(ctx context.Context) error {
+	return func(ctx context.Context) error {
 			g, groupCtx := errgroup.WithContext(ctx)
 			// Every worktree was created under the same temporary directory, so
 			// the first one's filesystem sizes removal for all of them.
@@ -151,14 +146,7 @@ func (w *Worktrees) Cleanup(ctx context.Context, l log
 						return nil
 					}
 
-					err := filter.TraceGitWorktreeRemove(
-						groupCtx,
-						worktree.Ref,
-						worktree.Path,
-						func(ctx context.Context) error {
-							return gitRunner.RemoveWorktree(ctx, worktree.Path)
-						},
-					)
+					err := gitRunner.RemoveWorktree(groupCtx, worktree.Path)
 					if err != nil {
 						// If the error is due to the worktree not existing, log and continue
 						// This can happen during parallel test execution or if cleanup runs twice
@@ -188,8 +176,7 @@ func (w *Worktrees) Cleanup(ctx context.Context, l log
 			}
 
 			return g.Wait()
-		},
-	)
+		}(ctx)
 }
 
 // worktreesToRemove returns the worktrees Cleanup has to delete, one entry per
@@ -655,21 +642,13 @@ func surveyGitExpressions(
 		g.Go(func() error {
 			var expressionDiffs *git.Diffs
 
-			diffErr := filter.TraceGitDiff(
-				gCtx, gitExpression.FromRef, gitExpression.ToRef, repoRemote,
-				func(ctx context.Context) error {
-					var err error
+			expressionDiffs, err := gitRunner.Diff(gCtx, gitExpression.FromRef, gitExpression.ToRef)
 
-					expressionDiffs, err = gitRunner.Diff(ctx, gitExpression.FromRef, gitExpression.ToRef)
-
-					return err
-				})
-
 			mu.Lock()
 			defer mu.Unlock()
 
-			if diffErr != nil {
-				errs = append(errs, diffErr)
+			if err != nil {
+				errs = append(errs, err)
 
 				return nil
 			}
@@ -755,9 +734,7 @@ func NewWorktrees(
 	repoCommit := gitRunner.GetHeadCommit(ctx)
 
 	// Wrap entire worktree creation process with telemetry
-	traceErr := filter.TraceGitWorktreesCreate(
-		ctx, workingDir, len(gitRefs), repoRemote, repoBranch, repoCommit,
-		func(ctx context.Context) error {
+	traceErr := func(ctx context.Context) error {
 			// The survey decides which references get a worktree and how much
 			// of each one is checked out, so it runs before any of them are.
 			survey, err := surveyGitExpressions(ctx, v, gitRunner, gitExpressions, gitRefs, repoRemote)
@@ -806,7 +783,7 @@ func NewWorktrees(
 			}
 
 			return nil
-		})
+		}(ctx)
 
 	if traceErr != nil && outerErr == nil {
 		l.Warnf("telemetry trace error during worktree creation: %v", traceErr)
@@ -924,14 +901,8 @@ func recordDiffTelemetry(ctx context.Context, diffs *g
 
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
+	_ = ctx
+	_ = diffs
 }
 
 // createGitWorktrees creates detached worktrees for each unique Git reference needed by filters.
@@ -1064,11 +1035,7 @@ func createGitWorktree(
 		return "", err
 	}
 
-	err = filter.TraceGitWorktreeCreate(
-		ctx, opts.ref, tmpDir, opts.repoRemote, opts.repoBranch, opts.repoCommit,
-		func(ctx context.Context) error {
-			return materializeGitWorktree(ctx, l, v, gitRunner, registerMu, tmpDir, opts)
-		})
+	err = materializeGitWorktree(ctx, l, v, gitRunner, registerMu, tmpDir, opts)
 	if err != nil {
 		if cleanErr := v.FS.RemoveAll(tmpDir); cleanErr != nil {
 			l.Warnf("failed to clean worktree directory %s: %v", tmpDir, cleanErr)
