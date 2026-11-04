--- internal/runner/controller.go.orig	1979-11-29 21:00:00 UTC
+++ internal/runner/controller.go
@@ -12,10 +12,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 
 	"github.com/gruntwork-io/terragrunt/internal/queue"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
-
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // UnitRunnerFunc executes a single [component.Unit].
@@ -84,13 +80,7 @@ func (dr *Controller) Run(ctx context.Context, l log.L
 
 // Run drains the queue and returns an error summarizing every entry that failed.
 func (dr *Controller) Run(ctx context.Context, l log.Logger) error {
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "runner_pool_controller", map[string]any{
-			"total_tasks":             len(dr.q.Entries),
-			"concurrency":             dr.concurrency,
-			"fail_fast":               dr.q.FailFast,
-			"ignore_dependency_order": dr.q.IgnoreDependencyOrder,
-		}, func(childCtx context.Context, l log.Logger) error {
+	childCtx := ctx
 			var (
 				wg        sync.WaitGroup
 				resultsMu sync.Mutex
@@ -238,14 +228,6 @@ func (dr *Controller) Run(ctx context.Context, l log.L
 				}
 			}
 
-			if span := trace.SpanFromContext(childCtx); span.IsRecording() {
-				span.SetAttributes(
-					attribute.Int("tasks_succeeded", succeeded),
-					attribute.Int("tasks_failed", failed),
-					attribute.Int("tasks_early_exit", earlyExit),
-				)
-			}
-
 			return errors.Join(errCollector...)
-		})
+
 }
