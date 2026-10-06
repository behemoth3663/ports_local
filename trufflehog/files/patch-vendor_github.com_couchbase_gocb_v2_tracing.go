--- vendor/github.com/couchbase/gocb/v2/tracing.go.orig	2026-10-06 09:46:09 UTC
+++ vendor/github.com/couchbase/gocb/v2/tracing.go
@@ -1,9 +1,9 @@ import (
 package gocb
 
 import (
-	"github.com/couchbase/gocbcore/v10"
-	"go.opentelemetry.io/otel/trace"
 	"time"
+
+	"github.com/couchbase/gocbcore/v10"
 )
 
 func tracerAddRef(tracer RequestTracer) {
@@ -34,8 +34,8 @@ type OtelAwareRequestTracer interface {
 }
 
 type OtelAwareRequestTracer interface {
-	Wrapped() trace.Tracer
-	Provider() trace.TracerProvider
+	Wrapped() interface{}
+	Provider() interface{}
 }
 
 // RequestSpan is the interface for spans that are created by a RequestTracer.
@@ -47,11 +47,10 @@ type RequestSpan interface {
 }
 
 // RequestSpanContext is the interface for external span contexts that can be passed in into the SDK option blocks.
-type RequestSpanContext interface {
-}
+type RequestSpanContext interface{}
 
 type OtelAwareRequestSpan interface {
-	Wrapped() trace.Span
+	Wrapped() interface{}
 }
 
 type coreRequestTracerWrapper struct {
@@ -84,8 +83,10 @@ func (span *coreRequestSpanWrapper) AddEvent(key strin
 	span.span.SetAttribute(key, timestamp)
 }
 
-type noopSpan struct{}
-type noopSpanContext struct{}
+type (
+	noopSpan        struct{}
+	noopSpanContext struct{}
+)
 
 var (
 	defaultNoopSpanContext = noopSpanContext{}
