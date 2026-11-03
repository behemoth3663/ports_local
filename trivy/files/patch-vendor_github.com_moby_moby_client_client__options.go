--- vendor/github.com/moby/moby/client/client_options.go.orig	2026-11-02 21:24:37 UTC
+++ vendor/github.com/moby/moby/client/client_options.go
@@ -12,11 +12,9 @@ import (
 	"strings"
 	"time"
 
-	cerrdefs "github.com/containerd/errdefs"
+	cerrdefs 	"github.com/containerd/errdefs"
 	"github.com/docker/go-connections/sockets"
 	"github.com/docker/go-connections/tlsconfig"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
-	"go.opentelemetry.io/otel/trace"
 )
 
 type clientConfig struct {
@@ -63,8 +61,8 @@ type clientConfig struct {
 	// responseHooks contains hooks to call on responses.
 	responseHooks []ResponseHook
 
-	// traceOpts is a list of options to configure the tracing span.
-	traceOpts []otelhttp.Option
+	// traceOpts are kept for backward compatibility and ignored.
+	traceOpts []interface{}
 }
 
 // RequestHook is called for each HTTP request made by the client.
@@ -406,15 +404,15 @@ func WithAPIVersionNegotiation() Opt {
 
 // WithTraceProvider sets the trace provider for the client.
 // If this is not set then the global trace provider is used.
-func WithTraceProvider(provider trace.TracerProvider) Opt {
+func WithTraceProvider(provider interface{}) Opt {
 	return func(c *clientConfig) error {
-		c.traceOpts = append(c.traceOpts, otelhttp.WithTracerProvider(provider))
+		c.traceOpts = append(c.traceOpts, provider)
 		return nil
 	}
 }
 
 // WithTraceOptions sets tracing span options for the client.
-func WithTraceOptions(opts ...otelhttp.Option) Opt {
+func WithTraceOptions(opts ...interface{}) Opt {
 	return func(c *clientConfig) error {
 		c.traceOpts = append(c.traceOpts, opts...)
 		return nil
