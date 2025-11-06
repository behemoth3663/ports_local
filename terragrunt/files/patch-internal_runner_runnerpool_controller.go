--- internal/runner/runnerpool/controller.go.orig	2025-11-05 17:03:04 UTC
+++ internal/runner/runnerpool/controller.go
@@ -12,7 +12,6 @@ import (
 
 	"github.com/gruntwork-io/terragrunt/internal/queue"
 	"github.com/gruntwork-io/terragrunt/internal/runner/common"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 
 	"github.com/puzpuzpuz/xsync/v3"
 )
@@ -81,12 +80,7 @@ func (dr *Controller) Run(ctx context.Context, l log.L
 
 // Run executes the Queue return error summarizing all entries that failed to run.
 func (dr *Controller) Run(ctx context.Context, l log.Logger) error {
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, "runner_pool_controller", map[string]any{
-		"total_tasks":             len(dr.q.Entries),
-		"concurrency":             dr.concurrency,
-		"fail_fast":               dr.q.FailFast,
-		"ignore_dependency_order": dr.q.IgnoreDependencyOrder,
-	}, func(childCtx context.Context) error {
+	return func(childCtx context.Context) error {
 		var (
 			wg      sync.WaitGroup
 			sem     = make(chan struct{}, dr.concurrency)
@@ -193,5 +187,5 @@ func (dr *Controller) Run(ctx context.Context, l log.L
 		}
 
 		return errCollector.ErrorOrNil()
-	})
+	}(ctx)
 }
