--- vendor/github.com/go-openapi/runtime/client/runtime.go.orig	2023-12-09 18:15:16 UTC
+++ vendor/github.com/go-openapi/runtime/client/runtime.go
@@ -36,7 +36,6 @@ import (
 	"github.com/go-openapi/runtime/middleware"
 	"github.com/go-openapi/runtime/yamlpc"
 	"github.com/go-openapi/strfmt"
-	"github.com/opentracing/opentracing-go"
 )
 
 const (
@@ -304,9 +303,6 @@ func NewWithClient(host, basePath string, schemes []st
 // A new client span is created for each request.
 // If the context of the client operation does not contain an active span, no span is created.
 // The provided opts are applied to each spans - for example to add global tags.
-func (r *Runtime) WithOpenTracing(opts ...opentracing.StartSpanOption) runtime.ClientTransport {
-	return newOpenTracingTransport(r, r.Host, opts)
-}
 
 // WithOpenTelemetry adds opentelemetry support to the provided runtime.
 // A new client span is created for each request.
