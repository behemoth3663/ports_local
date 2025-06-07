--- vendor/cloud.google.com/go/internal/trace/trace.go.orig	2023-09-11 19:22:13 UTC
+++ vendor/cloud.google.com/go/internal/trace/trace.go
@@ -16,43 +16,19 @@ import (
 
 import (
 	"context"
-	"fmt"
 
-	"go.opencensus.io/trace"
-	"golang.org/x/xerrors"
-	"google.golang.org/api/googleapi"
 	"google.golang.org/genproto/googleapis/rpc/code"
-	"google.golang.org/grpc/status"
 )
 
 // StartSpan adds a span to the trace with the given name.
 func StartSpan(ctx context.Context, name string) context.Context {
-	ctx, _ = trace.StartSpan(ctx, name)
 	return ctx
 }
 
 // EndSpan ends a span with the given error.
 func EndSpan(ctx context.Context, err error) {
-	span := trace.FromContext(ctx)
-	if err != nil {
-		span.SetStatus(toStatus(err))
-	}
-	span.End()
 }
 
-// toStatus interrogates an error and converts it to an appropriate
-// OpenCensus status.
-func toStatus(err error) trace.Status {
-	var err2 *googleapi.Error
-	if ok := xerrors.As(err, &err2); ok {
-		return trace.Status{Code: httpStatusCodeToOCCode(err2.Code), Message: err2.Message}
-	} else if s, ok := status.FromError(err); ok {
-		return trace.Status{Code: int32(s.Code()), Message: s.Message()}
-	} else {
-		return trace.Status{Code: int32(code.Code_UNKNOWN), Message: err.Error()}
-	}
-}
-
 // TODO(deklerk): switch to using OpenCensus function when it becomes available.
 // Reference: https://github.com/googleapis/googleapis/blob/26b634d2724ac5dd30ae0b0cbfb01f07f2e4050e/google/rpc/code.proto
 func httpStatusCodeToOCCode(httpStatusCode int) int32 {
@@ -90,22 +66,4 @@ func TracePrintf(ctx context.Context, attrMap map[stri
 // incurred from using trace.FromContext(ctx) yet we could avoid
 // throwing away the work done by ctx, span := trace.StartSpan.
 func TracePrintf(ctx context.Context, attrMap map[string]interface{}, format string, args ...interface{}) {
-	var attrs []trace.Attribute
-	for k, v := range attrMap {
-		var a trace.Attribute
-		switch v := v.(type) {
-		case string:
-			a = trace.StringAttribute(k, v)
-		case bool:
-			a = trace.BoolAttribute(k, v)
-		case int:
-			a = trace.Int64Attribute(k, int64(v))
-		case int64:
-			a = trace.Int64Attribute(k, v)
-		default:
-			a = trace.StringAttribute(k, fmt.Sprintf("%#v", v))
-		}
-		attrs = append(attrs, a)
-	}
-	trace.FromContext(ctx).Annotatef(attrs, format, args...)
 }
