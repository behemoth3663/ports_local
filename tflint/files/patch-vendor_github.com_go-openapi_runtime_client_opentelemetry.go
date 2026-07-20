--- vendor/github.com/go-openapi/runtime/client/opentelemetry.go.orig	2026-06-19 07:30:21 UTC
+++ vendor/github.com/go-openapi/runtime/client/opentelemetry.go
@@ -6,16 +6,8 @@ import (
 import (
 	"context"
 	"fmt"
-	"net/http"
 	"strings"
 
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/propagation"
-	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
-	"go.opentelemetry.io/otel/trace"
-
 	"github.com/go-openapi/runtime"
 	"github.com/go-openapi/strfmt"
 )
@@ -72,11 +64,7 @@ type config struct {
 }
 
 type config struct {
-	Tracer            trace.Tracer
-	Propagator        propagation.TextMapPropagator
-	SpanStartOptions  []trace.SpanStartOption
 	SpanNameFormatter func(*runtime.ClientOperation) string
-	TracerProvider    trace.TracerProvider
 }
 
 type OpenTelemetryOpt interface {
@@ -89,34 +77,6 @@ func (o optionFunc) apply(c *config) {
 	o(c)
 }
 
-// WithTracerProvider specifies a tracer provider to use for creating a tracer.
-// If none is specified, the global provider is used.
-func WithTracerProvider(provider trace.TracerProvider) OpenTelemetryOpt {
-	return optionFunc(func(c *config) {
-		if provider != nil {
-			c.TracerProvider = provider
-		}
-	})
-}
-
-// WithPropagators configures specific propagators. If this
-// option isn't specified, then the global TextMapPropagator is used.
-func WithPropagators(ps propagation.TextMapPropagator) OpenTelemetryOpt {
-	return optionFunc(func(c *config) {
-		if ps != nil {
-			c.Propagator = ps
-		}
-	})
-}
-
-// WithSpanOptions configures an additional set of
-// trace.SpanOptions, which are applied to each new span.
-func WithSpanOptions(opts ...trace.SpanStartOption) OpenTelemetryOpt {
-	return optionFunc(func(c *config) {
-		c.SpanStartOptions = append(c.SpanStartOptions, opts...)
-	})
-}
-
 // WithSpanNameFormatter takes a function that will be called on every
 // request and the returned string will become the Span Name.
 func WithSpanNameFormatter(f func(op *runtime.ClientOperation) string) OpenTelemetryOpt {
@@ -136,7 +96,6 @@ type openTelemetryTransport struct {
 type openTelemetryTransport struct {
 	transport runtime.ClientTransport
 	host      string
-	tracer    trace.Tracer
 	config    *config
 }
 
@@ -149,10 +108,7 @@ func newOpenTelemetryTransport(transport runtime.Clien
 	const baseOptions = 4
 	defaultOpts := make([]OpenTelemetryOpt, 0, len(opts)+baseOptions)
 	defaultOpts = append(defaultOpts,
-		WithSpanOptions(trace.WithSpanKind(trace.SpanKindClient)),
 		WithSpanNameFormatter(defaultTransportFormatter),
-		WithPropagators(otel.GetTextMapPropagator()),
-		WithTracerProvider(otel.GetTracerProvider()),
 	)
 
 	c := newConfig(append(defaultOpts, opts...)...)
@@ -189,39 +145,15 @@ func (t *openTelemetryTransport) SubmitContext(ctx con
 	params := op.Params
 	reader := op.Reader
 
-	var span trace.Span
-	defer func() {
-		if span != nil {
-			span.End()
-		}
-	}()
-
 	op.Params = runtime.ClientRequestWriterFunc(func(req runtime.ClientRequest, reg strfmt.Registry) error {
-		span = t.newOpenTelemetrySpan(ctx, op, req.GetHeaderParams())
 		return params.WriteToRequest(req, reg)
 	})
 
 	op.Reader = runtime.ClientResponseReaderFunc(func(response runtime.ClientResponse, consumer runtime.Consumer) (any, error) {
-		if span != nil {
-			statusCode := response.Code()
-			// NOTE: this is replaced by semconv.HTTPResponseStatusCode in semconv v1.21
-			span.SetAttributes(semconv.HTTPResponseStatusCode(statusCode))
-			// NOTE: the conversion from HTTP status code to trace code is no longer available with
-			// semconv v1.21
-			const minHTTPStatusIsError = 400
-			if statusCode >= minHTTPStatusIsError {
-				span.SetStatus(codes.Error, http.StatusText(statusCode))
-			}
-		}
-
 		return reader.ReadResponse(response, consumer)
 	})
 
 	submit, err := t.submitWrapped(ctx, op)
-	if err != nil && span != nil {
-		span.RecordError(err)
-		span.SetStatus(codes.Error, err.Error())
-	}
 
 	return submit, err
 }
@@ -237,53 +169,12 @@ func (t *openTelemetryTransport) submitWrapped(ctx con
 	return t.transport.Submit(op)
 }
 
-func (t *openTelemetryTransport) newOpenTelemetrySpan(ctx context.Context, op *runtime.ClientOperation, header http.Header) trace.Span {
-	tracer := t.tracer
-	if tracer == nil {
-		if span := trace.SpanFromContext(ctx); span.SpanContext().IsValid() {
-			tracer = newTracer(span.TracerProvider())
-		} else {
-			tracer = newTracer(otel.GetTracerProvider())
-		}
-	}
-
-	ctx, span := tracer.Start(ctx, t.config.SpanNameFormatter(op), t.config.SpanStartOptions...)
-
-	var scheme string
-	if len(op.Schemes) > 0 {
-		scheme = op.Schemes[0]
-	}
-
-	span.SetAttributes(
-		attribute.String("net.peer.name", t.host),
-		attribute.String(string(semconv.HTTPRouteKey), op.PathPattern),
-		attribute.String(string(semconv.HTTPRequestMethodKey), op.Method),
-		attribute.String("span.kind", trace.SpanKindClient.String()),
-		attribute.String("http.scheme", scheme),
-	)
-
-	carrier := propagation.HeaderCarrier(header)
-	t.config.Propagator.Inject(ctx, carrier)
-
-	return span
-}
-
-func newTracer(tp trace.TracerProvider) trace.Tracer {
-	return tp.Tracer(tracerName, trace.WithInstrumentationVersion(version()))
-}
-
 func newConfig(opts ...OpenTelemetryOpt) *config {
 	c := &config{
-		Propagator: otel.GetTextMapPropagator(),
 	}
 
 	for _, opt := range opts {
 		opt.apply(c)
-	}
-
-	// Tracer is only initialized if manually specified. Otherwise, can be passed with the tracing context.
-	if c.TracerProvider != nil {
-		c.Tracer = newTracer(c.TracerProvider)
 	}
 
 	return c
