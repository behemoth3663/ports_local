diff --git a/internal/filter/telemetry.go b/internal/filter/telemetry.go
index b70fa8394..4b4ed7817 100644
--- internal/filter/telemetry.go.orig
+++ internal/filter/telemetry.go
@@ -3,8 +3,6 @@ package filter
 
 import (
 	"context"
-
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 )
 
 // Telemetry operation names for git worktree and filter operations.
@@ -75,8 +73,7 @@ func TraceGitWorktreeCreate(
 		attrs[AttrGitRepoCommit] = repoCommit
 	}
 
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, nil, TelemetryOpGitWorktreeCreate, attrs, telemetry.WithoutLogger(fn))
+	return fn(ctx)
 }
 
 // TraceGitWorktreeRemove wraps a git worktree remove operation with telemetry.
@@ -85,11 +82,7 @@ func TraceGitWorktreeRemove(
 	ref, worktreeDir string,
 	fn func(ctx context.Context) error,
 ) error {
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, nil, TelemetryOpGitWorktreeRemove, map[string]any{
-			AttrGitRef:         ref,
-			AttrGitWorktreeDir: worktreeDir,
-		}, telemetry.WithoutLogger(fn))
+	return fn(ctx)
 }
 
 // TraceGitWorktreesCreate wraps multiple git worktree create operations with telemetry.
@@ -116,8 +109,7 @@ func TraceGitWorktreesCreate(
 		attrs[AttrGitRepoCommit] = repoCommit
 	}
 
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, nil, TelemetryOpGitWorktreesCreate, attrs, telemetry.WithoutLogger(fn))
+	return fn(ctx)
 }
 
 // TraceGitWorktreesCleanup wraps git worktrees cleanup with telemetry.
@@ -134,8 +126,7 @@ func TraceGitWorktreesCleanup(
 		attrs[AttrGitRepoRemote] = repoRemote
 	}
 
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, nil, TelemetryOpGitWorktreesCleanup, attrs, telemetry.WithoutLogger(fn))
+	return fn(ctx)
 }
 
 // TraceGitDiff wraps a git diff operation with telemetry.
@@ -152,8 +143,7 @@ func TraceGitDiff(
 		attrs[AttrGitRepoRemote] = repoRemote
 	}
 
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, nil, TelemetryOpGitDiff, attrs, telemetry.WithoutLogger(fn))
+	return fn(ctx)
 }
 
 // TraceGitWorktreeDiscovery wraps git worktree discovery operations with telemetry.
@@ -162,10 +152,7 @@ func TraceGitWorktreeDiscovery(
 	pairCount int,
 	fn func(ctx context.Context) error,
 ) error {
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, nil, TelemetryOpGitWorktreeDiscovery, map[string]any{
-			AttrWorktreePairCount: pairCount,
-		}, telemetry.WithoutLogger(fn))
+	return fn(ctx)
 }
 
 // TraceGitWorktreeStackWalk wraps git worktree stack walking operations with telemetry.
@@ -174,11 +161,7 @@ func TraceGitWorktreeStackWalk(
 	fromRef, toRef string,
 	fn func(ctx context.Context) error,
 ) error {
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, nil, TelemetryOpGitWorktreeStackWalk, map[string]any{
-			AttrGitFromRef: fromRef,
-			AttrGitToRef:   toRef,
-		}, telemetry.WithoutLogger(fn))
+	return fn(ctx)
 }
 
 // TraceGitWorktreeFilterApply wraps filter application to git worktrees with telemetry.
@@ -187,11 +170,7 @@ func TraceGitWorktreeFilterApply(
 	filterCount, resultCount int,
 	fn func(ctx context.Context) error,
 ) error {
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, nil, TelemetryOpGitWorktreeFilterApply, map[string]any{
-			AttrFilterCount: filterCount,
-			AttrResultCount: resultCount,
-		}, telemetry.WithoutLogger(fn))
+	return fn(ctx)
 }
 
 // TraceFilterEvaluate wraps filter evaluation with telemetry.
@@ -200,19 +179,12 @@ func TraceFilterEvaluate(
 	filterCount, componentCount int,
 	fn func(ctx context.Context) error,
 ) error {
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, nil, TelemetryOpFilterEvaluate, map[string]any{
-			AttrFilterCount:    filterCount,
-			AttrComponentCount: componentCount,
-		}, telemetry.WithoutLogger(fn))
+	return fn(ctx)
 }
 
 // TraceFilterParse wraps filter parsing with telemetry.
 func TraceFilterParse(ctx context.Context, query string, fn func(ctx context.Context) error) error {
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, nil, TelemetryOpFilterParse, map[string]any{
-			AttrFilterQuery: query,
-		}, telemetry.WithoutLogger(fn))
+	return fn(ctx)
 }
 
 // TraceGitFilterExpand wraps git filter expansion with telemetry.
@@ -222,14 +194,7 @@ func TraceGitFilterExpand(
 	addedCount, removedCount, changedCount int,
 	fn func(ctx context.Context) error,
 ) error {
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, nil, TelemetryOpGitFilterExpand, map[string]any{
-			AttrGitFromRef:     fromRef,
-			AttrGitToRef:       toRef,
-			AttrGitDiffAdded:   addedCount,
-			AttrGitDiffRemoved: removedCount,
-			AttrGitDiffChanged: changedCount,
-		}, telemetry.WithoutLogger(fn))
+	return fn(ctx)
 }
 
 // TraceGitFilterEvaluate wraps git filter evaluation with telemetry.
@@ -239,12 +204,7 @@ func TraceGitFilterEvaluate(
 	componentCount int,
 	fn func(ctx context.Context) error,
 ) error {
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, nil, TelemetryOpGitFilterEvaluate, map[string]any{
-			AttrGitFromRef:     fromRef,
-			AttrGitToRef:       toRef,
-			AttrComponentCount: componentCount,
-		}, telemetry.WithoutLogger(fn))
+	return fn(ctx)
 }
 
 // TraceGraphFilterTraverse wraps graph filter traversal with telemetry.
@@ -254,9 +214,5 @@ func TraceGraphFilterTraverse(
 	componentCount int,
 	fn func(ctx context.Context) error,
 ) error {
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, nil, TelemetryOpGraphFilterTraverse, map[string]any{
-			AttrFilterType:     filterType,
-			AttrComponentCount: componentCount,
-		}, telemetry.WithoutLogger(fn))
+	return fn(ctx)
 }
