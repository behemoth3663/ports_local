--- vendor/github.com/dylibso/observe-sdk/go/adapter.go.orig	2026-11-02 21:24:36 UTC
+++ vendor/github.com/dylibso/observe-sdk/go/adapter.go
@@ -2,13 +2,10 @@ import (
 
 import (
 	"context"
-	"fmt"
 	"log"
-	"strings"
 	"time"
 
 	"github.com/tetratelabs/wazero"
-	trace "go.opentelemetry.io/proto/otlp/trace/v1"
 )
 
 // The primary interface that every Adapter needs to follow
@@ -86,42 +83,6 @@ func (b *AdapterBase) Stop(wait bool) {
 	if wait {
 		b.eventBucket.Wait()
 	}
-}
-
-// MakeOtelCallSpans recursively constructs call spans in open telemetry format
-func (b *AdapterBase) MakeOtelCallSpans(event CallEvent, parentId []byte, traceId string) []*trace.Span {
-	name := event.FunctionName()
-	span := NewOtelSpan(traceId, parentId, name, event.Time, event.Time.Add(event.Duration))
-	span.Attributes = append(span.Attributes, NewOtelKeyValueString("function-name", fmt.Sprintf("function-call-%s", name)))
-
-	spans := []*trace.Span{span}
-	for _, ev := range event.Within() {
-		if call, ok := ev.(CallEvent); ok {
-			spans = append(spans, b.MakeOtelCallSpans(call, span.SpanId, traceId)...)
-		}
-		if alloc, ok := ev.(MemoryGrowEvent); ok {
-			kv := NewOtelKeyValueInt64("allocation", int64(alloc.MemoryGrowAmount()))
-			i, existing := GetOtelAttrFromSpan("allocation", span)
-			if existing != nil {
-				span.Attributes[i] = AddOtelKeyValueInt64(kv, existing)
-			} else {
-				span.Attributes = append(span.Attributes, kv)
-			}
-		}
-		if tags, ok := ev.(SpanTagsEvent); ok {
-			for _, tag := range tags.Tags {
-				parts := strings.Split(tag, ":")
-				if len(parts) != 2 {
-					log.Printf("Invalid tag: %s\n", tag)
-					continue
-				}
-
-				kv := NewOtelKeyValueString(parts[0], parts[1])
-				span.Attributes = append(span.Attributes, kv)
-			}
-		}
-	}
-	return spans
 }
 
 // Definition of how to filter our Spans to reduce noise
