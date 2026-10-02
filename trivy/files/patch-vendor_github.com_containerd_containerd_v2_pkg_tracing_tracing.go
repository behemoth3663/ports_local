--- vendor/github.com/containerd/containerd/v2/pkg/tracing/tracing.go.orig	2026-10-02 12:18:10 UTC
+++ vendor/github.com/containerd/containerd/v2/pkg/tracing/tracing.go
@@ -1,134 +1,54 @@
-/*
-   Copyright The containerd Authors.
-
-   Licensed under the Apache License, Version 2.0 (the "License");
-   you may not use this file except in compliance with the License.
-   You may obtain a copy of the License at
-
-       http://www.apache.org/licenses/LICENSE-2.0
-
-   Unless required by applicable law or agreed to in writing, software
-   distributed under the License is distributed on an "AS IS" BASIS,
-   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
-   See the License for the specific language governing permissions and
-   limitations under the License.
-*/
-
 package tracing
 
 import (
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
 
-// StartConfig defines configuration for a new span object.
 type StartConfig struct {
-	spanOpts []trace.SpanStartOption
+	spanOpts []any
 }
 
 type SpanOpt func(config *StartConfig)
 
-// WithAttribute appends attributes to a new created span.
 func WithAttribute(k string, v any) SpanOpt {
 	return func(config *StartConfig) {
-		config.spanOpts = append(config.spanOpts,
-			trace.WithAttributes(Attribute(k, v)))
+		config.spanOpts = append(config.spanOpts, Attribute(k, v))
 	}
 }
 
-// UpdateHTTPClient updates the http client with the necessary otel transport
-func UpdateHTTPClient(client *http.Client, name string) {
-	client.Transport = otelhttp.NewTransport(
-		client.Transport,
-		otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
-			return name
-		}),
-	)
-}
+func UpdateHTTPClient(_ *http.Client, _ string) {}
 
-// StartSpan starts child span in a context.
-func StartSpan(ctx context.Context, opName string, opts ...SpanOpt) (context.Context, *Span) {
-	config := StartConfig{}
-	for _, fn := range opts {
-		fn(&config)
-	}
-	tracer := otel.Tracer("")
-	if parent := trace.SpanFromContext(ctx); parent != nil && parent.SpanContext().IsValid() {
-		tracer = parent.TracerProvider().Tracer("")
-	}
-	ctx, span := tracer.Start(ctx, opName, config.spanOpts...)
-	return ctx, &Span{otelSpan: span}
+func StartSpan(ctx context.Context, _ string, _ ...SpanOpt) (context.Context, *Span) {
+	return ctx, &Span{}
 }
 
-// SpanFromContext returns the current Span from the context.
-func SpanFromContext(ctx context.Context) *Span {
-	return &Span{
-		otelSpan: trace.SpanFromContext(ctx),
-	}
+func SpanFromContext(_ context.Context) *Span {
+	return &Span{}
 }
 
-// Span is wrapper around otel trace.Span.
-// Span is the individual component of a trace. It represents a
-// single named and timed operation of a workflow that is traced.
-type Span struct {
-	otelSpan trace.Span
-}
+type Span struct{}
 
-// End completes the span.
-func (s *Span) End() {
-	s.otelSpan.End()
-}
+func (s *Span) End()                                {}
+func (s *Span) AddEvent(_ string, _ ...AttributeKV) {}
+func (s *Span) RecordError(_ error, _ ...any)       {}
+func (s *Span) SetStatus(_ error)                   {}
+func (s *Span) SetAttributes(_ ...AttributeKV)      {}
 
-// AddEvent adds an event with provided name and options.
-func (s *Span) AddEvent(name string, attributes ...attribute.KeyValue) {
-	s.otelSpan.AddEvent(name, trace.WithAttributes(attributes...))
-}
-
-// RecordError will record err as an exception span event for this span
-func (s *Span) RecordError(err error, options ...trace.EventOption) {
-	s.otelSpan.RecordError(err, options...)
-}
-
-// SetStatus sets the status of the current span.
-// If an error is encountered, it records the error and sets span status to Error.
-func (s *Span) SetStatus(err error) {
-	if err != nil {
-		s.otelSpan.RecordError(err)
-		s.otelSpan.SetStatus(codes.Error, err.Error())
-	} else {
-		s.otelSpan.SetStatus(codes.Ok, "")
-	}
-}
-
-// SetAttributes sets kv as attributes of the span.
-func (s *Span) SetAttributes(kv ...attribute.KeyValue) {
-	s.otelSpan.SetAttributes(kv...)
-}
-
 const spanDelimiter = "."
 
-// Name sets the span name by joining a list of strings in dot separated format.
 func Name(names ...string) string {
 	return strings.Join(names, spanDelimiter)
 }
 
-// Attribute takes a key value pair and returns attribute.KeyValue type.
-func Attribute(k string, v any) attribute.KeyValue {
+func Attribute(k string, v any) AttributeKV {
 	return keyValue(k, v)
 }
 
-// HTTPStatusCodeAttributes generates HTTP response status code attributes
-// as specified by the current OpenTelemetry semantic conventions.
-func HTTPStatusCodeAttributes(code int) []attribute.KeyValue {
-	return []attribute.KeyValue{
-		attribute.Int("http.response.status_code", code),
-		attribute.Int("http.status_code", code), // Deprecated: SemConv <= v1.21
+func HTTPStatusCodeAttributes(code int) []AttributeKV {
+	return []AttributeKV{
+		{Key: "http.response.status_code", Value: code},
+		{Key: "http.status_code", Value: code},
 	}
 }
