--- vendor/github.com/dylibso/observe-sdk/go/otel_formatter.go.orig	2026-10-02 12:18:10 UTC
+++ vendor/github.com/dylibso/observe-sdk/go/otel_formatter.go
@@ -1,132 +1,75 @@ package observe
 package observe
 
-import (
-	"encoding/binary"
-	"encoding/hex"
-	"time"
+import "time"
 
-	common "go.opentelemetry.io/proto/otlp/common/v1"
-	resource "go.opentelemetry.io/proto/otlp/resource/v1"
-	trace "go.opentelemetry.io/proto/otlp/trace/v1"
-)
-
 type OtelTrace struct {
 	TraceId    string
-	TracesData *trace.TracesData
+	TracesData any
 }
 
-func NewOtelTrace(traceId string, serviceName string, spans []*trace.Span) *OtelTrace {
-	return &OtelTrace{
-		TraceId: traceId,
-		TracesData: &trace.TracesData{
-			ResourceSpans: []*trace.ResourceSpans{
-				{
-					Resource: &resource.Resource{
-						Attributes: []*common.KeyValue{
-							NewOtelKeyValueString("service.name", serviceName),
-						},
-					},
-					ScopeSpans: []*trace.ScopeSpans{
-						{
-							Spans: spans,
-						},
-					},
-				},
-			},
-		},
-	}
+type otelKeyValue struct {
+	Key   string
+	Value any
 }
 
-func (t *OtelTrace) SetMetadata(te *TraceEvent, meta map[string]string) {
-	for _, rs := range t.TracesData.ResourceSpans {
-		for _, ss := range rs.ScopeSpans {
-			for _, span := range ss.Spans {
-				for key, value := range meta {
-					span.Attributes = append(span.Attributes, NewOtelKeyValueString(key, value))
-				}
-			}
-		}
-	}
+type otelSpan struct {
+	SpanId       []byte
+	ParentSpanId []byte
+	Name         string
+	Start        time.Time
+	End          time.Time
+	Attributes   []otelKeyValue
 }
 
-func NewOtelSpan(traceId string, parentId []byte, name string, start, end time.Time) *trace.Span {
+func NewOtelTrace(traceId string, _ string, _ []*otelSpan) *OtelTrace {
+	return &OtelTrace{TraceId: traceId}
+}
+
+func (t *OtelTrace) SetMetadata(_ *TraceEvent, _ map[string]string) {}
+
+func NewOtelSpan(_ string, parentId []byte, name string, start, end time.Time) *otelSpan {
 	if parentId == nil {
 		parentId = []byte{}
 	}
-
-	traceIdB, err := hex.DecodeString(traceId)
-	if err != nil {
-		panic(err)
+	return &otelSpan{
+		ParentSpanId: parentId,
+		Name:         name,
+		Start:        start,
+		End:          end,
 	}
-
-	spanId := NewSpanId().Msb()
-	spanIdB := make([]byte, 8)
-	binary.LittleEndian.PutUint64(spanIdB, spanId)
-
-	return &trace.Span{
-		TraceId:           traceIdB,
-		SpanId:            spanIdB,
-		ParentSpanId:      parentId,
-		Name:              name,
-		Kind:              1,
-		StartTimeUnixNano: uint64(start.UnixNano()),
-		EndTimeUnixNano:   uint64(end.UnixNano()),
-		// uses empty defaults for remaining fields...
-	}
 }
 
-func NewOtelKeyValueString(key string, value string) *common.KeyValue {
-	strVal := &common.AnyValue_StringValue{
-		StringValue: value,
-	}
-	return &common.KeyValue{
-		Key: key,
-		Value: &common.AnyValue{
-			Value: strVal,
-		},
-	}
+func NewOtelKeyValueString(key string, value string) otelKeyValue {
+	return otelKeyValue{Key: key, Value: value}
 }
 
-func NewOtelKeyValueInt64(key string, value int64) *common.KeyValue {
-	intVal := &common.AnyValue_IntValue{
-		IntValue: value,
-	}
-	return &common.KeyValue{
-		Key: key,
-		Value: &common.AnyValue{
-			Value: intVal,
-		},
-	}
+func NewOtelKeyValueInt64(key string, value int64) otelKeyValue {
+	return otelKeyValue{Key: key, Value: value}
 }
 
-func GetOtelAttrFromSpan(attr string, span *trace.Span) (int, *common.KeyValue) {
-	for i, attr := range span.Attributes {
-		if attr.Key == "allocation" {
-			return i, attr
+func GetOtelAttrFromSpan(attr string, span *otelSpan) (int, *otelKeyValue) {
+	for i := range span.Attributes {
+		if span.Attributes[i].Key == attr {
+			return i, &span.Attributes[i]
 		}
 	}
 	return -1, nil
 }
 
-func AddOtelKeyValueInt64(kvs ...*common.KeyValue) *common.KeyValue {
-	if len(kvs) > 0 {
-		retKv := &common.KeyValue{
-			Key:   kvs[0].Key,
-			Value: kvs[0].Value,
+func AddOtelKeyValueInt64(kvs ...*otelKeyValue) otelKeyValue {
+	if len(kvs) == 0 || kvs[0] == nil {
+		return otelKeyValue{}
+	}
+	ret := *kvs[0]
+	for i := 1; i < len(kvs); i++ {
+		if kvs[i] == nil {
+			continue
 		}
-		for i := 1; i < len(kvs); i++ {
-			v, ok := retKv.Value.Value.(*common.AnyValue_IntValue)
-			if ok {
-				curr, ok := kvs[i].Value.Value.(*common.AnyValue_IntValue)
-				if ok {
-					intVal := &common.AnyValue_IntValue{
-						IntValue: v.IntValue + curr.IntValue,
-					}
-					retKv.Value.Value = intVal
-				}
-			}
+		lhs, ok1 := ret.Value.(int64)
+		rhs, ok2 := kvs[i].Value.(int64)
+		if ok1 && ok2 {
+			ret.Value = lhs + rhs
 		}
-		return retKv
 	}
-	return nil
+	return ret
 }
