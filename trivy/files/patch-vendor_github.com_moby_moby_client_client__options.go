--- vendor/github.com/moby/moby/client/client_options.go.orig	2026-10-02 12:18:10 UTC
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
+	// traceOpts is retained for compatibility and currently unused.
+	traceOpts []any
 }
 
 // ResponseHook is called for each HTTP response returned by the daemon.
@@ -396,15 +394,15 @@ func WithAPIVersionNegotiation() Opt {
 
 // WithTraceProvider sets the trace provider for the client.
 // If this is not set then the global trace provider is used.
-func WithTraceProvider(provider trace.TracerProvider) Opt {
+func WithTraceProvider(_ any) Opt {
 	return func(c *clientConfig) error {
-		c.traceOpts = append(c.traceOpts, otelhttp.WithTracerProvider(provider))
+		_ = c
 		return nil
 	}
 }
 
 // WithTraceOptions sets tracing span options for the client.
-func WithTraceOptions(opts ...otelhttp.Option) Opt {
+func WithTraceOptions(opts ...any) Opt {
 	return func(c *clientConfig) error {
 		c.traceOpts = append(c.traceOpts, opts...)
 		return nil
