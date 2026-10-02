--- internal/worktrees/worktrees.go.orig	1979-11-29 21:00:00 UTC
+++ internal/worktrees/worktrees.go
@@ -4,6 +4,7 @@ import (
 
 import (
 	"context"
+	"errors"
 	"fmt"
 	"io/fs"
 	"os"
@@ -15,13 +16,10 @@ import (
 	"sync"
 	"time"
 
-	"errors"
-
 	"github.com/gruntwork-io/terragrunt/internal/component"
 	"github.com/gruntwork-io/terragrunt/internal/experiment"
 	"github.com/gruntwork-io/terragrunt/internal/filter"
 	"github.com/gruntwork-io/terragrunt/internal/git"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
@@ -914,14 +912,8 @@ func recordDiffTelemetry(ctx context.Context, diffs *g
 
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
