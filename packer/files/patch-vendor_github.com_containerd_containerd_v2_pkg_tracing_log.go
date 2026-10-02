--- vendor/github.com/containerd/containerd/v2/pkg/tracing/log.go.orig	2026-10-02 20:06:39 UTC
+++ vendor/github.com/containerd/containerd/v2/pkg/tracing/log.go
@@ -18,8 +18,6 @@ import (
 
 import (
 	"github.com/containerd/log"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // allLevels is the equivalent to [logrus.AllLevels].
@@ -68,17 +66,15 @@ func (h *LogrusHook) Fire(entry *log.Entry) error {
 
 // Fire is called when a log event occurs.
 func (h *LogrusHook) Fire(entry *log.Entry) error {
-	span := trace.SpanFromContext(entry.Context)
+	span := SpanFromContext(entry.Context)
 	if span == nil {
 		return nil
 	}
 
-	if !span.SpanContext().IsValid() {
-		return nil
-	}
-
 	if h.enableTraceIDField {
-		entry.Data["trace_id"] = span.SpanContext().TraceID().String()
+		if traceID := span.TraceID(); traceID != "" {
+			entry.Data["trace_id"] = traceID
+		}
 	}
 
 	if !span.IsRecording() {
@@ -87,16 +83,15 @@ func (h *LogrusHook) Fire(entry *log.Entry) error {
 
 	span.AddEvent(
 		entry.Message,
-		trace.WithAttributes(logrusDataToAttrs(entry.Data)...),
-		trace.WithAttributes(attribute.String("level", entry.Level.String())),
-		trace.WithTimestamp(entry.Time),
+		logrusDataToAttrs(entry.Data)...,
 	)
+	span.SetAttributes(KeyValue{Key: "level", Value: entry.Level.String()})
 
 	return nil
 }
 
-func logrusDataToAttrs(data map[string]any) []attribute.KeyValue {
-	attrs := make([]attribute.KeyValue, 0, len(data))
+func logrusDataToAttrs(data map[string]any) []KeyValue {
+	attrs := make([]KeyValue, 0, len(data))
 	for k, v := range data {
 		attrs = append(attrs, keyValue(k, v))
 	}
