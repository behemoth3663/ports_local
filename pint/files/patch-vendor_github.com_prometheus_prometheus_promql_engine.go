--- vendor/github.com/prometheus/prometheus/promql/engine.go.orig	2026-09-25 07:25:38 UTC
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
 
@@ -248,10 +242,6 @@ func (q *query) Exec(ctx context.Context) *Result {
 
 // Exec implements the Query interface.
 func (q *query) Exec(ctx context.Context) *Result {
-	if span := trace.SpanFromContext(ctx); span != nil {
-		span.SetAttributes(attribute.String(queryTag, q.stmt.String()))
-	}
-
 	// Exec query.
 	res, warnings, err := q.ng.exec(ctx, q)
 	return &Result{Err: err, Value: res, Warnings: warnings}
@@ -714,13 +704,6 @@ func (ng *Engine) exec(ctx context.Context, q *query) 
 				f = append(f, slog.Any("error", err))
 			}
 			f = append(f, slog.Any("stats", stats.NewQueryStats(q.Stats())))
-			if span := trace.SpanFromContext(ctx); span != nil {
-				spanCtx := span.SpanContext()
-				f = append(f,
-					slog.Any("spanID", spanCtx.SpanID()),
-					slog.Any("traceID", spanCtx.TraceID()),
-				)
-			}
 			if origin := ctx.Value(QueryOrigin{}); origin != nil {
 				for k, v := range origin.(map[string]any) {
 					f = append(f, slog.Any(k, v))
@@ -789,19 +772,11 @@ func (ng *Engine) execEvalStmt(ctx context.Context, qu
 	prepareSpanTimer, ctxPrepare := query.stats.GetSpanTimer(ctx, stats.QueryPreparationTime, ng.metrics.queryPrepareTime, ng.metrics.queryPrepareTimeHistogram)
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
 		prepareSpanTimer.Finish()
 		return nil, nil, err
 	}
-	querierSpan.End()
 	defer querier.Close()
 
 	ng.populateSeries(ctxPrepare, querier, s)
@@ -1092,13 +1067,7 @@ func (ng *Engine) populateSeries(ctx context.Context, 
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
@@ -1146,11 +1115,6 @@ func checkAndExpandSeriesSet(ctx context.Context, expr
 		if e.Series != nil {
 			return nil, nil
 		}
-		// This span is created only when a selector is read from storage. The
-		// result is cached in e.Series. At most one span is produced per storage
-		// selector per query. The span is not produced per step or per series.
-		ctx, span := otel.Tracer("").Start(ctx, "promqlExpandSeries", trace.WithAttributes(attribute.String("selector", e.String())))
-		defer span.End()
 		series, ws, err := expandSeriesSet(ctx, e.UnexpandedSeriesSet)
 		if e.SkipHistogramBuckets {
 			for i := range series {
@@ -1158,11 +1122,6 @@ func checkAndExpandSeriesSet(ctx context.Context, expr
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
@@ -2119,13 +2078,6 @@ func (ev *evaluator) eval(ctx context.Context, expr pa
 
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
@@ -2581,13 +2533,11 @@ func (ev *evaluator) eval(ctx context.Context, expr pa
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
