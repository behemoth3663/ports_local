--- internal/scheduler/scheduler.go.orig	2023-10-12 15:35:06 UTC
+++ internal/scheduler/scheduler.go
@@ -10,10 +10,6 @@ import (
 	"log"
 
 	"github.com/hashicorp/terraform-ls/internal/job"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 )
 
 const tracerName = "github.com/hashicorp/terraform-ls/internal/scheduler"
@@ -75,38 +71,7 @@ func (s *Scheduler) eval(ctx context.Context) {
 		}
 
 		ctx = job.WithIgnoreState(ctx, nextJob.IgnoreState)
-		jobSpan := trace.SpanFromContext(ctx)
-
-		ctx, span := otel.Tracer(tracerName).Start(ctx, "job-eval:"+nextJob.Type,
-			trace.WithAttributes(attribute.KeyValue{
-				Key:   attribute.Key("JobID"),
-				Value: attribute.StringValue(id.String()),
-			}, attribute.KeyValue{
-				Key:   attribute.Key("JobType"),
-				Value: attribute.StringValue(nextJob.Type),
-			}, attribute.KeyValue{
-				Key:   attribute.Key("Priority"),
-				Value: attribute.IntValue(int(nextJob.Priority)),
-			}, attribute.KeyValue{
-				Key:   attribute.Key("URI"),
-				Value: attribute.StringValue(nextJob.Dir.URI),
-			}))
-
 		jobErr := nextJob.Func(ctx)
-
-		if jobErr != nil {
-			if errors.Is(jobErr, job.StateNotChangedErr{Dir: nextJob.Dir}) {
-				span.SetStatus(codes.Ok, "state not changed")
-			} else {
-				span.RecordError(jobErr)
-				span.SetStatus(codes.Error, "job failed")
-			}
-		} else {
-			span.SetStatus(codes.Ok, "job finished")
-		}
-		span.End()
-		jobSpan.SetStatus(codes.Ok, "ok")
-		jobSpan.End()
 
 		deferredJobIds := make(job.IDs, 0)
 		if nextJob.Defer != nil {
