--- vendor/github.com/moby/moby/client/client_options.go.orig	2026-09-23 09:45:10 UTC
+++ vendor/github.com/moby/moby/client/client_options.go
@@ -15,8 +15,6 @@ import (
 	cerrdefs "github.com/containerd/errdefs"
 	"github.com/docker/go-connections/sockets"
 	"github.com/docker/go-connections/tlsconfig"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
-	"go.opentelemetry.io/otel/trace"
 )
 
 type clientConfig struct {
@@ -61,8 +59,8 @@ type clientConfig struct {
 	// responseHooks is a list of custom response hooks to call on responses.
 	responseHooks []ResponseHook
 
-	// traceOpts is a list of options to configure the tracing span.
-	traceOpts []otelhttp.Option
+	// traceOpts are kept for backwards compatibility with WithTrace* options.
+	traceOpts []any
 }
 
 // ResponseHook is called for each HTTP response returned by the daemon.
@@ -394,19 +392,18 @@ func WithAPIVersionNegotiation() Opt {
 	}
 }
 
-// WithTraceProvider sets the trace provider for the client.
-// If this is not set then the global trace provider is used.
-func WithTraceProvider(provider trace.TracerProvider) Opt {
+// WithTraceProvider is a no-op kept for backwards compatibility.
+func WithTraceProvider(provider any) Opt {
 	return func(c *clientConfig) error {
-		c.traceOpts = append(c.traceOpts, otelhttp.WithTracerProvider(provider))
+		_ = provider
 		return nil
 	}
 }
 
-// WithTraceOptions sets tracing span options for the client.
-func WithTraceOptions(opts ...otelhttp.Option) Opt {
+// WithTraceOptions is a no-op kept for backwards compatibility.
+func WithTraceOptions(opts ...any) Opt {
 	return func(c *clientConfig) error {
-		c.traceOpts = append(c.traceOpts, opts...)
+		_ = opts
 		return nil
 	}
 }
