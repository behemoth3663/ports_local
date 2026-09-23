diff --git a/vendor/github.com/moby/moby/client/client_options.go b/vendor/github.com/moby/moby/client/client_options.go
index 399255723..4c94d8898 100644
--- vendor/github.com/moby/moby/client/client_options.go.orig
+++ vendor/github.com/moby/moby/client/client_options.go
@@ -15,8 +15,6 @@ import (
 	cerrdefs "github.com/containerd/errdefs"
 	"github.com/docker/go-connections/sockets"
 	"github.com/docker/go-connections/tlsconfig"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
-	"go.opentelemetry.io/otel/trace"
 )
 
 type clientConfig struct {
@@ -62,7 +60,7 @@ type clientConfig struct {
 	responseHooks []ResponseHook
 
 	// traceOpts is a list of options to configure the tracing span.
-	traceOpts []otelhttp.Option
+	traceOpts []any
 }
 
 // ResponseHook is called for each HTTP response returned by the daemon.
@@ -396,17 +394,18 @@ func WithAPIVersionNegotiation() Opt {
 
 // WithTraceProvider sets the trace provider for the client.
 // If this is not set then the global trace provider is used.
-func WithTraceProvider(provider trace.TracerProvider) Opt {
+func WithTraceProvider(provider any) Opt {
 	return func(c *clientConfig) error {
-		c.traceOpts = append(c.traceOpts, otelhttp.WithTracerProvider(provider))
+		_ = provider
 		return nil
 	}
 }
 
 // WithTraceOptions sets tracing span options for the client.
-func WithTraceOptions(opts ...otelhttp.Option) Opt {
+func WithTraceOptions(opts ...any) Opt {
 	return func(c *clientConfig) error {
-		c.traceOpts = append(c.traceOpts, opts...)
+		_ = c
+		_ = opts
 		return nil
 	}
 }
