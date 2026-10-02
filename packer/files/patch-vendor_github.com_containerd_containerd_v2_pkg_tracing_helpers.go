--- vendor/github.com/containerd/containerd/v2/pkg/tracing/helpers.go.orig	2026-10-02 20:06:39 UTC
+++ vendor/github.com/containerd/containerd/v2/pkg/tracing/helpers.go
@@ -19,69 +19,67 @@ import (
 import (
 	"encoding/json"
 	"fmt"
-
-	"go.opentelemetry.io/otel/attribute"
 )
 
-func keyValue(k string, v any) attribute.KeyValue {
+func keyValue(k string, v any) KeyValue {
 	if v == nil {
-		return attribute.String(k, "<nil>")
+		return KeyValue{Key: k, Value: "<nil>"}
 	}
 
 	switch typed := v.(type) {
 	case bool:
-		return attribute.Bool(k, typed)
+		return KeyValue{Key: k, Value: typed}
 	case []bool:
-		return attribute.BoolSlice(k, typed)
+		return KeyValue{Key: k, Value: typed}
 	case int:
-		return attribute.Int(k, typed)
+		return KeyValue{Key: k, Value: typed}
 	case []int:
-		return attribute.IntSlice(k, typed)
+		return KeyValue{Key: k, Value: typed}
 	case int8:
-		return attribute.Int(k, int(typed))
+		return KeyValue{Key: k, Value: int(typed)}
 	case []int8:
 		ls := make([]int, 0, len(typed))
 		for _, i := range typed {
 			ls = append(ls, int(i))
 		}
-		return attribute.IntSlice(k, ls)
+		return KeyValue{Key: k, Value: ls}
 	case int16:
-		return attribute.Int(k, int(typed))
+		return KeyValue{Key: k, Value: int(typed)}
 	case []int16:
 		ls := make([]int, 0, len(typed))
 		for _, i := range typed {
 			ls = append(ls, int(i))
 		}
-		return attribute.IntSlice(k, ls)
+		return KeyValue{Key: k, Value: ls}
 	case int32:
-		return attribute.Int64(k, int64(typed))
+		return KeyValue{Key: k, Value: int64(typed)}
 	case []int32:
 		ls := make([]int64, 0, len(typed))
 		for _, i := range typed {
 			ls = append(ls, int64(i))
 		}
-		return attribute.Int64Slice(k, ls)
+		return KeyValue{Key: k, Value: ls}
 	case int64:
-		return attribute.Int64(k, typed)
+		return KeyValue{Key: k, Value: typed}
 	case []int64:
-		return attribute.Int64Slice(k, typed)
+		return KeyValue{Key: k, Value: typed}
 	case float64:
-		return attribute.Float64(k, typed)
+		return KeyValue{Key: k, Value: typed}
 	case []float64:
-		return attribute.Float64Slice(k, typed)
+		return KeyValue{Key: k, Value: typed}
 	case string:
-		return attribute.String(k, typed)
+		return KeyValue{Key: k, Value: typed}
 	case []string:
-		return attribute.StringSlice(k, typed)
+		return KeyValue{Key: k, Value: typed}
 	case error:
-		return attribute.String(k, fmt.Sprint(typed))
+		return KeyValue{Key: k, Value: fmt.Sprint(typed)}
 	}
 
 	if stringer, ok := v.(fmt.Stringer); ok {
-		return attribute.String(k, fmt.Sprint(stringer))
+		return KeyValue{Key: k, Value: fmt.Sprint(stringer)}
 	}
 	if b, err := json.Marshal(v); b != nil && err == nil {
-		return attribute.String(k, string(b))
+		return KeyValue{Key: k, Value: string(b)}
 	}
-	return attribute.String(k, fmt.Sprint(v))
+	return KeyValue{Key: k, Value: fmt.Sprint(v)}
 }
