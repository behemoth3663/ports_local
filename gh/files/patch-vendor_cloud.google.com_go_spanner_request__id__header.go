--- vendor/cloud.google.com/go/spanner/request_id_header.go.orig	2025-05-13 20:48:25 UTC
+++ vendor/cloud.google.com/go/spanner/request_id_header.go
@@ -26,8 +26,6 @@ import (
 	"time"
 
 	"github.com/googleapis/gax-go/v2"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 	"google.golang.org/api/iterator"
 	"google.golang.org/grpc"
 	"google.golang.org/grpc/codes"
@@ -144,13 +142,6 @@ func (wr *requestIDHeaderInjector) interceptUnary(ctx 
 	_, reqID, foundRequestID := gRPCCallOptionsToRequestID(opts)
 	if foundRequestID {
 		ctx = metadata.AppendToOutgoingContext(ctx, xSpannerRequestIDHeader, string(reqID))
-
-		// Associate the requestId as an attribute on the span in the current context.
-		span := trace.SpanFromContext(ctx)
-		span.SetAttributes(attribute.KeyValue{
-			Key:   xSpannerRequestIDSpanAttr,
-			Value: attribute.StringValue(string(reqID)),
-		})
 	}
 
 	err := invoker(ctx, method, req, reply, cc, opts...)
@@ -192,13 +183,6 @@ func (wr *requestIDHeaderInjector) interceptStream(ctx
 	_, reqID, foundRequestID := gRPCCallOptionsToRequestID(opts)
 	if foundRequestID {
 		ctx = metadata.AppendToOutgoingContext(ctx, xSpannerRequestIDHeader, string(reqID))
-
-		// Associate the requestId as an attribute on the span in the current context.
-		span := trace.SpanFromContext(ctx)
-		span.SetAttributes(attribute.KeyValue{
-			Key:   xSpannerRequestIDSpanAttr,
-			Value: attribute.StringValue(string(reqID)),
-		})
 	}
 
 	cs, err := streamer(ctx, desc, cc, method, opts...)
