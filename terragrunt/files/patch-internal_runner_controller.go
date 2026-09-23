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
@@ -84,169 +80,153 @@ func (dr *Controller) Run(ctx context.Context, l log.L
 
 // Run drains the queue and returns an error summarizing every entry that failed.
 func (dr *Controller) Run(ctx context.Context, l log.Logger) error {
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "runner_pool_controller", map[string]any{
-			"total_tasks":             len(dr.q.Entries),
-			"concurrency":             dr.concurrency,
-			"fail_fast":               dr.q.FailFast,
-			"ignore_dependency_order": dr.q.IgnoreDependencyOrder,
-		}, func(childCtx context.Context, l log.Logger) error {
-			var (
-				wg        sync.WaitGroup
-				resultsMu sync.Mutex
-				sem       = make(chan struct{}, dr.concurrency)
-			)
+		var (
+			wg        sync.WaitGroup
+			resultsMu sync.Mutex
+			sem       = make(chan struct{}, dr.concurrency)
+		)
 
-			// Fresh map per execution, kept on the controller after Run returns:
-			// the report sweep must read only this execution's errors, never
-			// stale ones from a prior Run on the same controller.
-			results := make(map[string]error)
-			dr.unitErrs = results
+		// Fresh map per execution, kept on the controller after Run returns:
+		// the report sweep must read only this execution's errors, never
+		// stale ones from a prior Run on the same controller.
+		results := make(map[string]error)
+		dr.unitErrs = results
 
-			recordResult := func(path string, err error) {
-				resultsMu.Lock()
-				defer resultsMu.Unlock()
+		recordResult := func(path string, err error) {
+			resultsMu.Lock()
+			defer resultsMu.Unlock()
 
-				results[path] = err
-			}
+			results[path] = err
+		}
 
-			if dr.runner == nil {
-				return ErrRunnerNotSet
-			}
+		if dr.runner == nil {
+			return ErrRunnerNotSet
+		}
 
-			l.Debugf("Runner Pool Controller: starting with %d tasks, concurrency %d",
-				len(dr.q.Entries), dr.concurrency)
+		l.Debugf("Runner Pool Controller: starting with %d tasks, concurrency %d",
+			len(dr.q.Entries), dr.concurrency)
 
-			// Initial signal to start scheduling
-			select {
-			case dr.readyCh <- struct{}{}:
-			default:
-			}
+		// Initial signal to start scheduling
+		select {
+		case dr.readyCh <- struct{}{}:
+		default:
+		}
 
-			for {
-				readyEntries := dr.q.GetReadyWithDependencies(l)
-				l.Debugf("Runner Pool Controller: found %d readyEntries tasks", len(readyEntries))
+		for {
+			readyEntries := dr.q.GetReadyWithDependencies(l)
+			l.Debugf("Runner Pool Controller: found %d readyEntries tasks", len(readyEntries))
 
-				for _, e := range readyEntries {
-					if !dr.q.ClaimForRunning(e) {
-						l.Debugf(
-							"Runner Pool Controller: skipping %s; fail-fast cancelled before dispatch",
-							e.Component.Path(),
-						)
+			for _, e := range readyEntries {
+				if !dr.q.ClaimForRunning(e) {
+					l.Debugf(
+						"Runner Pool Controller: skipping %s; fail-fast cancelled before dispatch",
+						e.Component.Path(),
+					)
 
-						continue
-					}
+					continue
+				}
 
-					l.Debugf("Runner Pool Controller: running %s", e.Component.Path())
+				l.Debugf("Runner Pool Controller: running %s", e.Component.Path())
 
-					sem <- struct{}{}
+				sem <- struct{}{}
 
-					wg.Add(1)
+				wg.Add(1)
 
-					go func(ent *queue.Entry) {
-						defer func() {
-							<-sem
-							wg.Done()
+				go func(ent *queue.Entry) {
+					defer func() {
+						<-sem
+						wg.Done()
 
-							select {
-							case dr.readyCh <- struct{}{}:
-							default:
-							}
-						}()
-
-						unit := dr.unitsMap[ent.Component.Path()]
-						if unit == nil {
-							err := NewUnitNotDiscoveredError(ent.Component.Path())
-							l.Errorf(
-								"Runner Pool Controller: unit for path %s not found in discovered units, skipping execution",
-								ent.Component.Path(),
-							)
-							dr.q.FailEntry(ent)
-							recordResult(ent.Component.Path(), err)
-
-							return
+						select {
+						case dr.readyCh <- struct{}{}:
+						default:
 						}
+					}()
 
-						err := dr.runner(childCtx, unit)
+					unit := dr.unitsMap[ent.Component.Path()]
+					if unit == nil {
+						err := NewUnitNotDiscoveredError(ent.Component.Path())
+						l.Errorf(
+							"Runner Pool Controller: unit for path %s not found in discovered units, skipping execution",
+							ent.Component.Path(),
+						)
+						dr.q.FailEntry(ent)
 						recordResult(ent.Component.Path(), err)
 
-						if err != nil {
-							l.Debugf("Runner Pool Controller: %s failed", ent.Component.Path())
-							dr.q.FailEntry(ent)
+						return
+					}
 
-							return
-						}
+					err := dr.runner(ctx, unit)
+					recordResult(ent.Component.Path(), err)
 
-						l.Debugf("Runner Pool Controller: %s succeeded", ent.Component.Path())
-						dr.q.SetEntryStatus(ent, queue.StatusSucceeded)
-					}(e)
-				}
+					if err != nil {
+						l.Debugf("Runner Pool Controller: %s failed", ent.Component.Path())
+						dr.q.FailEntry(ent)
 
-				if dr.q.Finished() {
-					break
-				}
+						return
+					}
 
-				select {
-				case <-dr.readyCh:
-				case <-childCtx.Done():
-					wg.Wait()
-					return nil
-				}
+					l.Debugf("Runner Pool Controller: %s succeeded", ent.Component.Path())
+					dr.q.SetEntryStatus(ent, queue.StatusSucceeded)
+				}(e)
 			}
 
-			wg.Wait()
+			if dr.q.Finished() {
+				break
+			}
 
-			var errCollector []error
+			select {
+			case <-dr.readyCh:
+			case <-ctx.Done():
+				wg.Wait()
+				return nil
+			}
+		}
 
-			var succeeded, failed, earlyExit int
+		wg.Wait()
 
-			for _, entry := range dr.q.Entries {
-				switch entry.Status {
-				case queue.StatusSucceeded:
-					succeeded++
-				case queue.StatusFailed:
-					failed++
-				case queue.StatusEarlyExit:
-					earlyExit++
-				case queue.StatusPending,
-					queue.StatusBlocked,
-					queue.StatusUnsorted,
-					queue.StatusReady,
-					queue.StatusRunning:
-					// Non-terminal states are not counted in the summary.
-				}
+		var errCollector []error
 
-				if err, ok := results[entry.Component.Path()]; ok {
-					if err == nil {
-						continue
-					}
+		var succeeded, failed, earlyExit int
 
-					errCollector = append(errCollector, err)
+		for _, entry := range dr.q.Entries {
+			switch entry.Status {
+			case queue.StatusSucceeded:
+				succeeded++
+			case queue.StatusFailed:
+				failed++
+			case queue.StatusEarlyExit:
+				earlyExit++
+			case queue.StatusPending,
+				queue.StatusBlocked,
+				queue.StatusUnsorted,
+				queue.StatusReady,
+				queue.StatusRunning:
+				// Non-terminal states are not counted in the summary.
+			}
 
+			if err, ok := results[entry.Component.Path()]; ok {
+				if err == nil {
 					continue
 				}
 
-				if entry.Status == queue.StatusEarlyExit {
-					failedDep := findFailedDependency(entry, dr.q)
-					errCollector = append(
-						errCollector,
-						NewUnitEarlyExitError(entry.Component.Path(), failedDep),
-					)
-				}
+				errCollector = append(errCollector, err)
 
-				if entry.Status == queue.StatusFailed {
-					errCollector = append(errCollector, NewUnitFailedError(entry.Component.Path()))
-				}
+				continue
 			}
 
-			if span := trace.SpanFromContext(childCtx); span.IsRecording() {
-				span.SetAttributes(
-					attribute.Int("tasks_succeeded", succeeded),
-					attribute.Int("tasks_failed", failed),
-					attribute.Int("tasks_early_exit", earlyExit),
+			if entry.Status == queue.StatusEarlyExit {
+				failedDep := findFailedDependency(entry, dr.q)
+				errCollector = append(
+					errCollector,
+					NewUnitEarlyExitError(entry.Component.Path(), failedDep),
 				)
 			}
 
-			return multierror.Join(errCollector...)
-		})
+			if entry.Status == queue.StatusFailed {
+				errCollector = append(errCollector, NewUnitFailedError(entry.Component.Path()))
+			}
+		}
+
+		return multierror.Join(errCollector...)
 }
