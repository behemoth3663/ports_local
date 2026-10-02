--- vendor/github.com/dylibso/observe-sdk/go/adapter.go.orig	2026-10-02 12:18:10 UTC
+++ vendor/github.com/dylibso/observe-sdk/go/adapter.go
@@ -8,7 +8,6 @@ import (
 	"time"
 
 	"github.com/tetratelabs/wazero"
-	trace "go.opentelemetry.io/proto/otlp/trace/v1"
 )
 
 // The primary interface that every Adapter needs to follow
@@ -89,12 +88,12 @@ func (b *AdapterBase) Stop(wait bool) {
 }
 
 // MakeOtelCallSpans recursively constructs call spans in open telemetry format
-func (b *AdapterBase) MakeOtelCallSpans(event CallEvent, parentId []byte, traceId string) []*trace.Span {
+func (b *AdapterBase) MakeOtelCallSpans(event CallEvent, parentId []byte, traceId string) []*otelSpan {
 	name := event.FunctionName()
 	span := NewOtelSpan(traceId, parentId, name, event.Time, event.Time.Add(event.Duration))
 	span.Attributes = append(span.Attributes, NewOtelKeyValueString("function-name", fmt.Sprintf("function-call-%s", name)))
 
-	spans := []*trace.Span{span}
+	spans := []*otelSpan{span}
 	for _, ev := range event.Within() {
 		if call, ok := ev.(CallEvent); ok {
 			spans = append(spans, b.MakeOtelCallSpans(call, span.SpanId, traceId)...)
@@ -102,8 +101,9 @@ func (b *AdapterBase) MakeOtelCallSpans(event CallEven
 		if alloc, ok := ev.(MemoryGrowEvent); ok {
 			kv := NewOtelKeyValueInt64("allocation", int64(alloc.MemoryGrowAmount()))
 			i, existing := GetOtelAttrFromSpan("allocation", span)
-			if existing != nil {
-				span.Attributes[i] = AddOtelKeyValueInt64(kv, existing)
+			if existing != nil && i >= 0 {
+				sum := AddOtelKeyValueInt64(&kv, existing)
+				span.Attributes[i] = sum
 			} else {
 				span.Attributes = append(span.Attributes, kv)
 			}
