--- vendor/github.com/prometheus/prometheus/promql/engine.go.orig	2026-11-05 15:06:51 UTC
+++ vendor/github.com/prometheus/prometheus/promql/engine.go
@@ -23,7 +23,6 @@ import (
 	"log/slog"
 	"maps"
 	"math"
-	"reflect"
 	"runtime"
 	"slices"
 	"sort"
@@ -35,10 +34,6 @@ import (
 	"github.com/prometheus/client_golang/prometheus"
 	"github.com/prometheus/common/model"
 	"github.com/prometheus/common/promslog"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 
 	"github.com/prometheus/prometheus/model/histogram"
 	"github.com/prometheus/prometheus/model/labels"
@@ -60,7 +55,6 @@ const (
 const (
 	namespace            = "prometheus"
 	subsystem            = "engine"
-	queryTag             = "query"
 	env                  = "query execution"
 	defaultLookbackDelta = 5 * time.Minute
 
@@ -383,10 +377,6 @@ func (q *query) Exec(ctx context.Context) *Result {
 
 // Exec implements the Query interface.
 func (q *query) Exec(ctx context.Context) *Result {
-	if span := trace.SpanFromContext(ctx); span != nil {
-		span.SetAttributes(attribute.String(queryTag, q.stmt.String()))
-	}
-
 	// Exec query.
 	res, warnings, err := q.ng.exec(ctx, q)
 	return &Result{Err: err, Value: res, Warnings: warnings}
@@ -1111,12 +1101,6 @@ func (ng *Engine) exec(ctx context.Context, q *query) 
 				f = append(f, slog.Any("error", err))
 			}
 			f = append(f, slog.Any("stats", stats.NewQueryStats(q.Stats())))
-			if spanCtx := trace.SpanContextFromContext(ctx); spanCtx.IsValid() {
-				f = append(f,
-					slog.Any("spanID", spanCtx.SpanID()),
-					slog.Any("traceID", spanCtx.TraceID()),
-				)
-			}
 			if origin := ctx.Value(QueryOrigin{}); origin != nil {
 				for k, v := range origin.(map[string]any) {
 					f = append(f, slog.Any(k, v))
@@ -1130,9 +1114,6 @@ func (ng *Engine) exec(ctx context.Context, q *query) 
 		ng.queryLoggerLock.RUnlock()
 	}()
 
-	execSpanTimer, ctx := q.stats.GetSpanTimer(ctx, stats.ExecTotalTime)
-	defer execSpanTimer.Finish()
-
 	finishQueue, err := ng.queueActive(ctx, q)
 	if err != nil {
 		if ng.enableQueryCost {
@@ -1142,9 +1123,6 @@ func (ng *Engine) exec(ctx context.Context, q *query) 
 	}
 	defer finishQueue()
 
-	evalSpanTimer, ctx := q.stats.GetSpanTimer(ctx, stats.EvalTotalTime)
-	defer evalSpanTimer.Finish()
-
 	// The base context might already be canceled on the first iteration (e.g. during shutdown).
 	if err := contextDone(ctx, env); err != nil {
 		return nil, nil, err
@@ -1194,9 +1172,7 @@ func (ng *Engine) queueActive(ctx context.Context, q *
 	if ng.activeQueryTracker == nil {
 		return func() {}, nil
 	}
-	queueSpanTimer, _ := q.stats.GetSpanTimer(ctx, stats.ExecQueueTime, ng.metrics.queryQueueTime, ng.metrics.queryQueueTimeHistogram)
 	queryIndex, err := ng.activeQueryTracker.Insert(ctx, q.q)
-	queueSpanTimer.Finish()
 	if err != nil {
 		return func() {}, err
 	}
@@ -1252,31 +1228,19 @@ func (ng *Engine) execEvalStmt(ctx context.Context, qu
 			query.sampleStats.SamplesRead = budget.samplesRead
 		}()
 	}
-	prepareSpanTimer, ctxPrepare := query.stats.GetSpanTimer(ctx, stats.QueryPreparationTime, ng.metrics.queryPrepareTime, ng.metrics.queryPrepareTimeHistogram)
 	mint, maxt := FindMinMaxTime(s)
 
-	_, querierSpan := otel.Tracer("").Start(ctxPrepare, "Querier", trace.WithAttributes(
-		attribute.Int64("mint", mint),
-		attribute.Int64("maxt", maxt),
-	))
 	querier, err := query.queryable.Querier(mint, maxt)
 	if err != nil {
-		querierSpan.RecordError(err)
-		querierSpan.SetStatus(codes.Error, err.Error())
-		querierSpan.End()
-		prepareSpanTimer.Finish()
 		return nil, nil, err
 	}
-	querierSpan.End()
 	defer querier.Close()
 
-	ng.populateSeries(ctxPrepare, querier, s)
-	prepareSpanTimer.Finish()
+	ng.populateSeries(ctx, querier, s)
 
 	// Modify the offset of vector and matrix selectors for the @ modifier
 	// w.r.t. the start time since only 1 evaluation will be done on them.
 	setOffsetForAtModifier(timeMilliseconds(s.Start), s.Expr)
-	evalSpanTimer, ctxInnerEval := query.stats.GetSpanTimer(ctx, stats.InnerEvalTime, ng.metrics.queryInnerEval, ng.metrics.queryInnerEvalHistogram)
 	// Instant evaluation. This is executed as a range evaluation with one step.
 	if s.Start.Equal(s.End) && s.Interval == 0 {
 		start := timeMilliseconds(s.Start)
@@ -1297,10 +1261,8 @@ func (ng *Engine) execEvalStmt(ctx context.Context, qu
 		}
 		query.sampleStats.InitStepTracking(start, start, 1)
 
-		val, warnings, err := evaluator.Eval(ctxInnerEval, s.Expr)
+		val, warnings, err := evaluator.Eval(ctx, s.Expr)
 
-		evalSpanTimer.Finish()
-
 		if err != nil {
 			return nil, warnings, err
 		}
@@ -1358,10 +1320,8 @@ func (ng *Engine) execEvalStmt(ctx context.Context, qu
 		querier:                  querier,
 	}
 	query.sampleStats.InitStepTracking(evaluator.startTimestamp, evaluator.endTimestamp, evaluator.interval)
-	val, warnings, err := evaluator.Eval(ctxInnerEval, s.Expr)
+	val, warnings, err := evaluator.Eval(ctx, s.Expr)
 
-	evalSpanTimer.Finish()
-
 	if err != nil {
 		return nil, warnings, err
 	}
@@ -1383,9 +1343,7 @@ func (ng *Engine) sortMatrixResult(ctx context.Context
 }
 
 func (ng *Engine) sortMatrixResult(ctx context.Context, query *query, mat Matrix) {
-	sortSpanTimer, _ := query.stats.GetSpanTimer(ctx, stats.ResultSortTime, ng.metrics.queryResultSort, ng.metrics.queryResultSortHistogram)
 	sort.Sort(mat)
-	sortSpanTimer.Finish()
 }
 
 // subqueryTimes returns the sum of offsets and ranges of all subqueries in the path.
@@ -1560,13 +1518,7 @@ func (ng *Engine) populateSeries(ctx context.Context, 
 			}
 			evalRange = 0
 			hints.By, hints.Grouping = extractGroupsFromPath(path)
-			selectCtx, selectSpan := otel.Tracer("").Start(ctx, "querierSelect", trace.WithAttributes(
-				attribute.String("selector", n.String()),
-				attribute.Int64("start", hints.Start),
-				attribute.Int64("end", hints.End),
-			))
-			n.UnexpandedSeriesSet = querier.Select(selectCtx, false, hints, n.LabelMatchers...)
-			selectSpan.End()
+			n.UnexpandedSeriesSet = querier.Select(ctx, false, hints, n.LabelMatchers...)
 		case *parser.MatrixSelector:
 			evalRange = n.Range
 		}
@@ -1623,8 +1575,6 @@ func (ev *evaluator) checkAndExpandSeriesSet(ctx conte
 		// This span is created only when a selector is read from storage. The
 		// result is cached in e.Series. At most one span is produced per storage
 		// selector per query. The span is not produced per step or per series.
-		ctx, span := otel.Tracer("").Start(ctx, "promqlExpandSeries", trace.WithAttributes(attribute.String("selector", e.String())))
-		defer span.End()
 		series, ws, err := ev.expandSeriesSet(ctx, e.UnexpandedSeriesSet)
 		if e.SkipHistogramBuckets {
 			for i := range series {
@@ -1632,11 +1582,6 @@ func (ev *evaluator) checkAndExpandSeriesSet(ctx conte
 			}
 		}
 		e.Series = series
-		span.SetAttributes(attribute.Int("num_series", len(series)))
-		if err != nil {
-			span.RecordError(err)
-			span.SetStatus(codes.Error, err.Error())
-		}
 		return ws, err
 	}
 	return nil, nil
@@ -2664,13 +2609,6 @@ func (ev *evaluator) eval(ctx context.Context, expr pa
 
 	numSteps := int((ev.endTimestamp-ev.startTimestamp)/ev.interval) + 1
 
-	// Create a new span to help investigate inner evaluation performances.
-	ctx, span := otel.Tracer("").Start(ctx, stats.InnerEvalTime.SpanOperation()+" eval "+reflect.TypeOf(expr).String())
-	defer span.End()
-	if ss, ok := expr.(interface{ ShortString() string }); ok {
-		span.SetAttributes(attribute.String("operation", ss.ShortString()))
-	}
-
 	switch e := expr.(type) {
 	case *parser.AggregateExpr:
 		// Grouping labels must be sorted (expected both by generateGroupingKey() and aggregation()).
@@ -3137,13 +3075,11 @@ func (ev *evaluator) eval(ctx context.Context, expr pa
 		}
 
 	case *parser.NumberLiteral:
-		span.SetAttributes(attribute.Float64("value", e.Val))
 		return ev.rangeEval(ctx, nil, func(_ []Vector, _ Matrix, _ [][]EvalSeriesHelper, enh *EvalNodeHelper) (Vector, annotations.Annotations) {
 			return append(enh.Out, Sample{F: e.Val, Metric: labels.EmptyLabels()}), nil
 		})
 
 	case *parser.StringLiteral:
-		span.SetAttributes(attribute.String("value", e.Val))
 		return String{V: e.Val, T: ev.startTimestamp}, nil
 
 	case *parser.VectorSelector:
