--- internal/worktrees/worktrees.go.orig	2025-12-17 15:14:20 UTC
+++ internal/worktrees/worktrees.go
@@ -17,7 +17,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/filter"
 	"github.com/gruntwork-io/terragrunt/internal/git"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"golang.org/x/sync/errgroup"
 )
 
@@ -479,14 +478,6 @@ func recordDiffTelemetry(ctx context.Context, diffs *g
 
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
 }
 
 // createGitWorktrees creates detached worktrees for each unique Git reference needed by filters.
