--- vendor/github.com/docker/docker/client/options.go.orig	2025-11-28 10:28:17 UTC
+++ vendor/github.com/docker/docker/client/options.go
@@ -12,8 +12,6 @@ import (
 	"github.com/docker/go-connections/sockets"
 	"github.com/docker/go-connections/tlsconfig"
 	"github.com/pkg/errors"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // Opt is a configuration option to initialize a [Client].
@@ -221,20 +219,6 @@ func WithAPIVersionNegotiation() Opt {
 func WithAPIVersionNegotiation() Opt {
 	return func(c *Client) error {
 		c.negotiateVersion = true
-		return nil
-	}
-}
-
-// WithTraceProvider sets the trace provider for the client.
-// If this is not set then the global trace provider will be used.
-func WithTraceProvider(provider trace.TracerProvider) Opt {
-	return WithTraceOptions(otelhttp.WithTracerProvider(provider))
-}
-
-// WithTraceOptions sets tracing span options for the client.
-func WithTraceOptions(opts ...otelhttp.Option) Opt {
-	return func(c *Client) error {
-		c.traceOpts = append(c.traceOpts, opts...)
 		return nil
 	}
 }
