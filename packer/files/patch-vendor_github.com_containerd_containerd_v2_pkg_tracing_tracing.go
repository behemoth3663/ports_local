--- vendor/github.com/containerd/containerd/v2/pkg/tracing/tracing.go.orig	2026-10-02 20:06:39 UTC
+++ vendor/github.com/containerd/containerd/v2/pkg/tracing/tracing.go
@@ -20,17 +20,16 @@ import (
 	"context"
 	"net/http"
 	"strings"
-
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 )
 
+type KeyValue struct {
+	Key   string
+	Value any
+}
+
 // StartConfig defines configuration for a new span object.
 type StartConfig struct {
-	spanOpts []trace.SpanStartOption
+	spanOpts []any
 }
 
 type SpanOpt func(config *StartConfig)
@@ -38,19 +37,14 @@ func WithAttribute(k string, v any) SpanOpt {
 // WithAttribute appends attributes to a new created span.
 func WithAttribute(k string, v any) SpanOpt {
 	return func(config *StartConfig) {
-		config.spanOpts = append(config.spanOpts,
-			trace.WithAttributes(Attribute(k, v)))
+		config.spanOpts = append(config.spanOpts, Attribute(k, v))
 	}
 }
 
 // UpdateHTTPClient updates the http client with the necessary otel transport
 func UpdateHTTPClient(client *http.Client, name string) {
-	client.Transport = otelhttp.NewTransport(
-		client.Transport,
-		otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
-			return name
-		}),
-	)
+	_ = client
+	_ = name
 }
 
 // StartSpan starts child span in a context.
@@ -59,59 +53,64 @@ func StartSpan(ctx context.Context, opName string, opt
 	for _, fn := range opts {
 		fn(&config)
 	}
-	tracer := otel.Tracer("")
-	if parent := trace.SpanFromContext(ctx); parent != nil && parent.SpanContext().IsValid() {
-		tracer = parent.TracerProvider().Tracer("")
-	}
-	ctx, span := tracer.Start(ctx, opName, config.spanOpts...)
-	return ctx, &Span{otelSpan: span}
+	_ = config
+	_ = opName
+	return ctx, &Span{}
 }
 
 // SpanFromContext returns the current Span from the context.
 func SpanFromContext(ctx context.Context) *Span {
-	return &Span{
-		otelSpan: trace.SpanFromContext(ctx),
-	}
+	_ = ctx
+	return &Span{}
 }
 
 // Span is wrapper around otel trace.Span.
 // Span is the individual component of a trace. It represents a
 // single named and timed operation of a workflow that is traced.
-type Span struct {
-	otelSpan trace.Span
-}
+type Span struct{}
 
 // End completes the span.
 func (s *Span) End() {
-	s.otelSpan.End()
+	_ = s
 }
 
 // AddEvent adds an event with provided name and options.
-func (s *Span) AddEvent(name string, attributes ...attribute.KeyValue) {
-	s.otelSpan.AddEvent(name, trace.WithAttributes(attributes...))
+func (s *Span) AddEvent(name string, attributes ...KeyValue) {
+	_ = s
+	_ = name
+	_ = attributes
 }
 
 // RecordError will record err as an exception span event for this span
-func (s *Span) RecordError(err error, options ...trace.EventOption) {
-	s.otelSpan.RecordError(err, options...)
+func (s *Span) RecordError(err error, options ...any) {
+	_ = s
+	_ = err
+	_ = options
 }
 
 // SetStatus sets the status of the current span.
 // If an error is encountered, it records the error and sets span status to Error.
 func (s *Span) SetStatus(err error) {
-	if err != nil {
-		s.otelSpan.RecordError(err)
-		s.otelSpan.SetStatus(codes.Error, err.Error())
-	} else {
-		s.otelSpan.SetStatus(codes.Ok, "")
-	}
+	_ = s
+	_ = err
 }
 
 // SetAttributes sets kv as attributes of the span.
-func (s *Span) SetAttributes(kv ...attribute.KeyValue) {
-	s.otelSpan.SetAttributes(kv...)
+func (s *Span) SetAttributes(kv ...KeyValue) {
+	_ = s
+	_ = kv
 }
 
+func (s *Span) IsRecording() bool {
+	_ = s
+	return false
+}
+
+func (s *Span) TraceID() string {
+	_ = s
+	return ""
+}
+
 const spanDelimiter = "."
 
 // Name sets the span name by joining a list of strings in dot separated format.
@@ -120,15 +119,15 @@ func Name(names ...string) string {
 }
 
 // Attribute takes a key value pair and returns attribute.KeyValue type.
-func Attribute(k string, v any) attribute.KeyValue {
+func Attribute(k string, v any) KeyValue {
 	return keyValue(k, v)
 }
 
 // HTTPStatusCodeAttributes generates HTTP response status code attributes
 // as specified by the current OpenTelemetry semantic conventions.
-func HTTPStatusCodeAttributes(code int) []attribute.KeyValue {
-	return []attribute.KeyValue{
-		attribute.Int("http.response.status_code", code),
-		attribute.Int("http.status_code", code), // Deprecated: SemConv <= v1.21
+func HTTPStatusCodeAttributes(code int) []KeyValue {
+	return []KeyValue{
+		{Key: "http.response.status_code", Value: code},
+		{Key: "http.status_code", Value: code}, // Deprecated: SemConv <= v1.21
 	}
 }
