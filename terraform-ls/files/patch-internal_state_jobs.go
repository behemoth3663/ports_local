--- internal/state/jobs.go.orig	2023-10-12 15:35:06 UTC
+++ internal/state/jobs.go
@@ -16,10 +16,6 @@ import (
 	lsctx "github.com/hashicorp/terraform-ls/internal/context"
 	"github.com/hashicorp/terraform-ls/internal/document"
 	"github.com/hashicorp/terraform-ls/internal/job"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 )
 
 type JobStore struct {
@@ -47,19 +43,12 @@ type ScheduledJob struct {
 
 	// EnqueueTime tracks time when the job was originally put into the queue
 	EnqueueTime time.Time
-	// TraceSpan represents a tracing span for the entire job lifecycle
-	// (from queuing to finishing execution).
-	TraceSpan trace.Span
+
 	// DocumentContext contains information from when & where the job was scheduled from
 	DocumentContext lsctx.Document
 }
 
 func (sj *ScheduledJob) Copy() *ScheduledJob {
-	// This may be an awkward way to copy the Span
-	// but the upstream doesn't seem to offer any better way.
-	newCtx := trace.ContextWithSpan(context.Background(), sj.TraceSpan)
-	traceSpan := trace.SpanFromContext(newCtx)
-
 	return &ScheduledJob{
 		ID:              sj.ID,
 		Job:             sj.Job.Copy(),
@@ -68,7 +57,6 @@ func (sj *ScheduledJob) Copy() *ScheduledJob {
 		JobErr:          sj.JobErr,
 		DeferredJobIDs:  sj.DeferredJobIDs.Copy(),
 		EnqueueTime:     sj.EnqueueTime,
-		TraceSpan:       traceSpan,
 		DocumentContext: sj.DocumentContext.Copy(),
 	}
 }
@@ -101,34 +89,12 @@ func (js *JobStore) EnqueueJob(ctx context.Context, ne
 	newJob.DependsOn = dependsOn
 	dirOpen := isDirOpen(txn, newJob.Dir)
 
-	_, jobSpan := otel.Tracer(tracerName).Start(ctx, "job",
-		trace.WithAttributes(attribute.KeyValue{
-			Key:   attribute.Key("JobID"),
-			Value: attribute.StringValue(newJobID.String()),
-		}, attribute.KeyValue{
-			Key:   attribute.Key("JobType"),
-			Value: attribute.StringValue(newJob.Type),
-		}, attribute.KeyValue{
-			Key:   attribute.Key("IsDirOpen"),
-			Value: attribute.BoolValue(dirOpen),
-		}, attribute.KeyValue{
-			Key:   attribute.Key("Priority"),
-			Value: attribute.IntValue(int(newJob.Priority)),
-		}, attribute.KeyValue{
-			Key:   attribute.Key("URI"),
-			Value: attribute.StringValue(newJob.Dir.URI),
-		}, attribute.KeyValue{
-			Key:   attribute.Key("DependsOn"),
-			Value: attribute.StringSliceValue(dependsOn.StringSlice()),
-		}))
-
 	sJob := &ScheduledJob{
 		ID:              newJobID,
 		Job:             newJob,
 		IsDirOpen:       dirOpen,
 		State:           StateQueued,
 		EnqueueTime:     time.Now(),
-		TraceSpan:       jobSpan,
 		DocumentContext: lsctx.DocumentContext(ctx),
 	}
 
@@ -186,8 +152,6 @@ func (js *JobStore) DequeueJobsForDir(dir document.Dir
 		if err != nil {
 			return err
 		}
-		sJob.TraceSpan.SetStatus(codes.Ok, "job dequeued")
-		sJob.TraceSpan.End()
 	}
 
 	txn.Commit()
@@ -322,28 +286,6 @@ func (js *JobStore) awaitNextJob(ctx context.Context, 
 		sJob.ID, priority, sJob.Priority, sJob.IsDirOpen, sJob.Type, sJob.Dir)
 
 	ctx = lsctx.WithDocumentContext(ctx, sJob.DocumentContext)
-	ctx = trace.ContextWithSpan(ctx, sJob.TraceSpan)
-
-	_, span := otel.Tracer(tracerName).Start(ctx, "job-wait",
-		trace.WithTimestamp(sJob.EnqueueTime),
-		trace.WithAttributes(attribute.KeyValue{
-			Key:   attribute.Key("JobID"),
-			Value: attribute.StringValue(sJob.ID.String()),
-		}, attribute.KeyValue{
-			Key:   attribute.Key("JobType"),
-			Value: attribute.StringValue(sJob.Type),
-		}, attribute.KeyValue{
-			Key:   attribute.Key("IsDirOpen"),
-			Value: attribute.BoolValue(sJob.IsDirOpen),
-		}, attribute.KeyValue{
-			Key:   attribute.Key("Priority"),
-			Value: attribute.IntValue(int(sJob.Priority)),
-		}, attribute.KeyValue{
-			Key:   attribute.Key("URI"),
-			Value: attribute.StringValue(sJob.Dir.URI),
-		}),
-	)
-	span.End()
 
 	return ctx, sJob.ID, sJob.Job, nil
 }
