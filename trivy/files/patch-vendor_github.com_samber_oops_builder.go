--- vendor/github.com/samber/oops/builder.go.orig	2026-10-02 12:18:10 UTC
+++ vendor/github.com/samber/oops/builder.go
@@ -9,7 +9,6 @@ import (
 
 	"github.com/oklog/ulid/v2"
 	"github.com/samber/lo"
-	"go.opentelemetry.io/otel/trace"
 )
 
 /**
@@ -276,12 +275,15 @@ func (o OopsErrorBuilder) WithContext(ctx context.Cont
 		}
 	}
 
-	spanCtx := trace.SpanContextFromContext(ctx)
-	if spanCtx.HasTraceID() {
-		o2.trace = spanCtx.TraceID().String()
+	if v := ctx.Value("trace_id"); v != nil {
+		if s, ok := v.(string); ok {
+			o2.trace = s
+		}
 	}
-	if spanCtx.HasSpanID() {
-		o2.span = spanCtx.SpanID().String()
+	if v := ctx.Value("span_id"); v != nil {
+		if s, ok := v.(string); ok {
+			o2.span = s
+		}
 	}
 
 	return o2
