--- vendor/github.com/moby/moby/client/client_options.go.orig	2026-03-05 13:48:11 UTC
+++ vendor/github.com/moby/moby/client/client_options.go
@@ -13,8 +13,6 @@ import (
 
 	"github.com/docker/go-connections/sockets"
 	"github.com/docker/go-connections/tlsconfig"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
-	"go.opentelemetry.io/otel/trace"
 )
 
 type clientConfig struct {
@@ -58,9 +56,6 @@ type clientConfig struct {
 
 	// responseHooks is a list of custom response hooks to call on responses.
 	responseHooks []ResponseHook
-
-	// traceOpts is a list of options to configure the tracing span.
-	traceOpts []otelhttp.Option
 }
 
 // ResponseHook is called for each HTTP response returned by the daemon.
@@ -341,20 +336,6 @@ func WithAPIVersionNegotiation() Opt {
 // or [WithAPIVersionFromEnv] to disable API version negotiation.
 func WithAPIVersionNegotiation() Opt {
 	return func(c *clientConfig) error {
-		return nil
-	}
-}
-
-// WithTraceProvider sets the trace provider for the client.
-// If this is not set then the global trace provider is used.
-func WithTraceProvider(provider trace.TracerProvider) Opt {
-	return WithTraceOptions(otelhttp.WithTracerProvider(provider))
-}
-
-// WithTraceOptions sets tracing span options for the client.
-func WithTraceOptions(opts ...otelhttp.Option) Opt {
-	return func(c *clientConfig) error {
-		c.traceOpts = append(c.traceOpts, opts...)
 		return nil
 	}
 }
