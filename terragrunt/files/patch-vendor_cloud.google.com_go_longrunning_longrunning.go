diff --git a/vendor/cloud.google.com/go/longrunning/longrunning.go b/vendor/cloud.google.com/go/longrunning/longrunning.go
index 666ea6b76..e64eefa76 100644
--- vendor/cloud.google.com/go/longrunning/longrunning.go.orig
+++ vendor/cloud.google.com/go/longrunning/longrunning.go
@@ -31,11 +31,6 @@ import (
 	pb "cloud.google.com/go/longrunning/autogen/longrunningpb"
 	gax "github.com/googleapis/gax-go/v2"
 	"github.com/googleapis/gax-go/v2/apierror"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
-	grpccodes "google.golang.org/grpc/codes"
 	"google.golang.org/grpc/status"
 	"google.golang.org/protobuf/proto"
 	"google.golang.org/protobuf/protoadapt"
@@ -50,7 +45,7 @@ type Operation struct {
 	c               operationsClient
 	proto           *pb.Operation
 	opName          string
-	initSpanContext trace.SpanContext
+	initSpanContext any
 }
 
 type operationsClient interface {
@@ -84,7 +79,7 @@ func InternalNewOperationWithMetadata(inner *autogen.OperationsClient, proto *pb
 // which is used to create a Span Link from the LRO Wait span.
 //
 // SetParentSpanContext is an EXPERIMENTAL API and may be changed or removed in the future.
-func (op *Operation) SetParentSpanContext(sc trace.SpanContext) {
+func (op *Operation) SetParentSpanContext(sc any) {
 	op.initSpanContext = sc
 }
 
@@ -172,35 +167,9 @@ func (op *Operation) waitWithInterval(ctx context.Context, resp protoadapt.Messa
 		bo.Max = bo.Initial
 	}
 
-	if !gax.IsFeatureEnabled("TRACING") {
-		return op.wait(ctx, resp, &bo, sl, opts...)
-	}
-
-	spanName := op.opName
-	if spanName == "" {
-		spanName = "*longrunning.Operation.Wait"
-	} else {
-		spanName = spanName + ".Wait"
-	}
-
-	var startOpts []trace.SpanStartOption
-	if op.initSpanContext.IsValid() {
-		startOpts = append(startOpts, trace.WithLinks(trace.Link{SpanContext: op.initSpanContext}))
-	}
-
-	tracer := otel.GetTracerProvider().Tracer("cloud.google.com/go")
-	ctx, span := tracer.Start(ctx, spanName, startOpts...)
-	defer span.End()
-	span.SetAttributes(
-		attribute.String("gcp.resource.destination.id", op.Name()),
-	)
-
-	err := op.waitTraced(ctx, resp, &bo, sl, opts...)
-	if err != nil {
-		span.SetStatus(codes.Error, err.Error())
-		span.RecordError(err)
-	}
-	return err
+	_ = op.opName
+	_ = op.initSpanContext
+	return op.wait(ctx, resp, &bo, sl, opts...)
 }
 
 type sleeper func(context.Context, time.Duration) error
@@ -221,57 +190,7 @@ func (op *Operation) wait(ctx context.Context, resp protoadapt.MessageV1, bo *ga
 }
 
 func (op *Operation) waitTraced(ctx context.Context, resp protoadapt.MessageV1, bo *gax.Backoff, sl sleeper, opts ...gax.CallOption) error {
-	tracer := otel.GetTracerProvider().Tracer("cloud.google.com/go")
-	pollAttempt := 0
-
-	for {
-		pollAttempt++
-		pollSpanName := "*longrunning.OperationsClient.GetOperation"
-
-		pollCtx, pollSpan := tracer.Start(ctx, pollSpanName)
-		pollSpan.SetAttributes(
-			attribute.String("gcp.resource.destination.id", op.Name()),
-			attribute.Int("gcp.longrunning.poll_attempt_count", pollAttempt),
-		)
-
-		err := op.Poll(pollCtx, resp, opts...)
-
-		pollSpan.SetAttributes(attribute.Bool("gcp.longrunning.done", op.Done()))
-		if op.Done() {
-			var statusCode int
-			if err != nil {
-				if apiErr, ok := apierror.FromError(err); ok {
-					statusCode = int(apiErr.GRPCStatus().Code())
-				} else {
-					statusCode = int(grpccodes.Unknown)
-				}
-			}
-			pollSpan.SetAttributes(attribute.Int("gcp.longrunning.status_code", statusCode))
-		}
-
-		if err != nil {
-			pollSpan.SetStatus(codes.Error, err.Error())
-			pollSpan.RecordError(err)
-			pollSpan.End()
-			return err
-		}
-
-		pollSpan.End()
-
-		if op.Done() {
-			return nil
-		}
-
-		_, sleepSpan := tracer.Start(ctx, "LRO Sleep")
-		sleepErr := sl(ctx, bo.Pause())
-		if sleepErr != nil {
-			sleepSpan.SetStatus(codes.Error, sleepErr.Error())
-			sleepSpan.RecordError(sleepErr)
-			sleepSpan.End()
-			return sleepErr
-		}
-		sleepSpan.End()
-	}
+	return op.wait(ctx, resp, bo, sl, opts...)
 }
 
 // Cancel starts asynchronous cancellation on a long-running operation. The server
