--- vendor/github.com/docker/docker/client/options.go.orig	2026-10-06 09:46:09 UTC
+++ vendor/github.com/docker/docker/client/options.go
@@ -12,8 +12,6 @@ import (
 	"github.com/docker/go-connections/sockets"
 	"github.com/docker/go-connections/tlsconfig"
 	"github.com/pkg/errors"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // Opt is a configuration option to initialize a [Client].
@@ -227,14 +225,15 @@ func WithAPIVersionNegotiation() Opt {
 
 // WithTraceProvider sets the trace provider for the client.
 // If this is not set then the global trace provider will be used.
-func WithTraceProvider(provider trace.TracerProvider) Opt {
-	return WithTraceOptions(otelhttp.WithTracerProvider(provider))
+func WithTraceProvider(provider interface{}) Opt {
+	return func(c *Client) error {
+		return nil
+	}
 }
 
 // WithTraceOptions sets tracing span options for the client.
-func WithTraceOptions(opts ...otelhttp.Option) Opt {
+func WithTraceOptions(opts ...interface{}) Opt {
 	return func(c *Client) error {
-		c.traceOpts = append(c.traceOpts, opts...)
 		return nil
 	}
 }
