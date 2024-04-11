--- internal/walker/walker.go.orig	2023-10-12 15:35:06 UTC
+++ internal/walker/walker.go
@@ -15,10 +15,6 @@ import (
 	"github.com/hashicorp/terraform-ls/internal/document"
 	"github.com/hashicorp/terraform-ls/internal/job"
 	"github.com/hashicorp/terraform-ls/internal/terraform/ast"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 )
 
 var (
@@ -111,7 +107,7 @@ func (w *Walker) StartWalking(ctx context.Context) err
 
 	go func() {
 		for {
-			pathCtx, nextDir, err := w.pathStore.AwaitNextDir(ctx)
+			_, nextDir, err := w.pathStore.AwaitNextDir(ctx)
 			if err != nil {
 				if errors.Is(err, context.Canceled) {
 					return
@@ -121,35 +117,17 @@ func (w *Walker) StartWalking(ctx context.Context) err
 				return
 			}
 
-			spanCtx := trace.SpanContextFromContext(pathCtx)
-
-			ctx = trace.ContextWithSpanContext(ctx, spanCtx)
-
-			tracer := otel.Tracer(tracerName)
-			ctx, span := tracer.Start(ctx, "walk-path", trace.WithAttributes(attribute.KeyValue{
-				Key:   attribute.Key("URI"),
-				Value: attribute.StringValue(nextDir.URI),
-			}))
-
 			err = w.walk(ctx, nextDir)
 			if err != nil {
 				w.logger.Printf("walker: walking through %q failed: %s", nextDir, err)
 				w.collectError(err)
-				span.RecordError(err)
-				span.SetStatus(codes.Error, "walking failed")
-				span.End()
 				continue
 			}
-			span.SetStatus(codes.Ok, "walking finished")
-			span.End()
 
 			err = w.pathStore.RemoveDir(nextDir)
 			if err != nil {
 				w.logger.Printf("walker: removing dir %q from queue failed: %s", nextDir, err)
 				w.collectError(err)
-				span.RecordError(err)
-				span.SetStatus(codes.Error, "walking failed")
-				span.End()
 				continue
 			}
