--- vendor/github.com/go-openapi/runtime/client/runtime.go.orig	2024-03-09 19:19:52 UTC
+++ vendor/github.com/go-openapi/runtime/client/runtime.go
@@ -306,17 +306,11 @@ func NewWithClient(host, basePath string, schemes []st
 // A new client span is created for each request.
 // If the context of the client operation does not contain an active span, no span is created.
 // The provided opts are applied to each spans - for example to add global tags.
-func (r *Runtime) WithOpenTracing(opts ...opentracing.StartSpanOption) runtime.ClientTransport {
-	return newOpenTracingTransport(r, r.Host, opts)
-}
 
 // WithOpenTelemetry adds opentelemetry support to the provided runtime.
 // A new client span is created for each request.
 // If the context of the client operation does not contain an active span, no span is created.
 // The provided opts are applied to each spans - for example to add global tags.
-func (r *Runtime) WithOpenTelemetry(opts ...OpenTelemetryOpt) runtime.ClientTransport {
-	return newOpenTelemetryTransport(r, r.Host, opts)
-}
 
 func (r *Runtime) pickScheme(schemes []string) string {
 	if v := r.selectScheme(r.schemes); v != "" {
