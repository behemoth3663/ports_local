--- vendor/github.com/go-openapi/runtime/client/runtime.go.orig	2024-03-09 19:19:52 UTC
+++ vendor/github.com/go-openapi/runtime/client/runtime.go
@@ -33,7 +33,6 @@ import (
 	"time"
 
 	"github.com/go-openapi/strfmt"
-	"github.com/opentracing/opentracing-go"
 
 	"github.com/go-openapi/runtime"
 	"github.com/go-openapi/runtime/logger"
@@ -306,9 +305,6 @@ func NewWithClient(host, basePath string, schemes []st
 // A new client span is created for each request.
 // If the context of the client operation does not contain an active span, no span is created.
 // The provided opts are applied to each spans - for example to add global tags.
-func (r *Runtime) WithOpenTracing(opts ...opentracing.StartSpanOption) runtime.ClientTransport {
-	return newOpenTracingTransport(r, r.Host, opts)
-}
 
 // WithOpenTelemetry adds opentelemetry support to the provided runtime.
 // A new client span is created for each request.
