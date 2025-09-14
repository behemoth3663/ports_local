--- vendor/github.com/go-openapi/runtime/client/runtime.go.orig	2025-09-27 06:04:17 UTC
+++ vendor/github.com/go-openapi/runtime/client/runtime.go
@@ -313,26 +313,11 @@ func NewWithClient(host, basePath string, schemes []st
 // usual opentracing options and opentracing-enabled transport.
 //
 // Passed options are ignored unless they are of type [OpenTelemetryOpt].
-func (r *Runtime) WithOpenTracing(opts ...any) runtime.ClientTransport {
-	otelOpts := make([]OpenTelemetryOpt, 0, len(opts))
-	for _, o := range opts {
-		otelOpt, ok := o.(OpenTelemetryOpt)
-		if !ok {
-			continue
-		}
-		otelOpts = append(otelOpts, otelOpt)
-	}
 
-	return r.WithOpenTelemetry(otelOpts...)
-}
-
 // WithOpenTelemetry adds opentelemetry support to the provided runtime.
 // A new client span is created for each request.
 // If the context of the client operation does not contain an active span, no span is created.
 // The provided opts are applied to each spans - for example to add global tags.
-func (r *Runtime) WithOpenTelemetry(opts ...OpenTelemetryOpt) runtime.ClientTransport {
-	return newOpenTelemetryTransport(r, r.Host, opts)
-}
 
 // EnableConnectionReuse drains the remaining body from a response
 // so that go will reuse the TCP connections.
