--- vendor/github.com/samber/oops/builder.go.orig	2026-11-02 21:24:37 UTC
+++ vendor/github.com/samber/oops/builder.go
@@ -9,7 +9,6 @@ import (
 
 	"github.com/oklog/ulid/v2"
 	"github.com/samber/lo"
-	"go.opentelemetry.io/otel/trace"
 )
 
 /**
@@ -274,14 +273,6 @@ func (o OopsErrorBuilder) WithContext(ctx context.Cont
 		default:
 			o2.context[fmt.Sprint(k)] = contextValueOrNil(ctx, k)
 		}
-	}
-
-	spanCtx := trace.SpanContextFromContext(ctx)
-	if spanCtx.HasTraceID() {
-		o2.trace = spanCtx.TraceID().String()
-	}
-	if spanCtx.HasSpanID() {
-		o2.span = spanCtx.SpanID().String()
 	}
 
 	return o2
