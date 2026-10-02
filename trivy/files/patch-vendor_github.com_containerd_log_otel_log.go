--- vendor/github.com/containerd/log/otel/log.go.orig	2026-10-02 12:18:10 UTC
+++ vendor/github.com/containerd/log/otel/log.go
@@ -1,37 +1,11 @@
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
-// Package otel provides integration between containerd/log and OpenTelemetry.
-//
-// In particular, it provides a hook that records log entries as events on
-// active OpenTelemetry spans.
 package otel
 
 import (
 	"slices"
 
 	"github.com/containerd/log"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 )
 
-// allLevels is the equivalent to [logrus.AllLevels].
-//
-// [logrus.AllLevels]: https://github.com/sirupsen/logrus/blob/v1.9.3/logrus.go#L80-L89
 var allLevels = []log.Level{
 	log.PanicLevel,
 	log.FatalLevel,
@@ -44,7 +18,6 @@ type HookOpt func(*LogrusHook)
 
 type HookOpt func(*LogrusHook)
 
-// NewLogrusHook creates a new logrus hook
 func NewLogrusHook(opts ...HookOpt) *LogrusHook {
 	hook := &LogrusHook{}
 	for _, opt := range opts {
@@ -62,8 +35,6 @@ func WithTraceIDField(enabled bool) HookOpt {
 	}
 }
 
-// WithLevel configures the minimum log level handled by the hook.
-// Entries below this level are ignored.
 func WithLevel(level log.Level) HookOpt {
 	return func(h *LogrusHook) {
 		for i, l := range allLevels {
@@ -75,26 +46,18 @@ func WithLevel(level log.Level) HookOpt {
 	}
 }
 
-// WithErrorStatusLevel configures the minimum log level that marks the
-// active span with an error status.
 func WithErrorStatusLevel(level log.Level) HookOpt {
 	return func(h *LogrusHook) {
 		h.errorStatusLevel = &level
 	}
 }
 
-// LogrusHook is a [logrus.Hook] which adds logrus events to active spans.
-// If the span is not recording or the span context is invalid, the hook
-// is a no-op.
-//
-// [logrus.Hook]: https://github.com/sirupsen/logrus/blob/v1.9.3/hooks.go#L3-L11
 type LogrusHook struct {
 	enableTraceIDField bool
 	errorStatusLevel   *log.Level
 	levels             []log.Level
 }
 
-// Levels returns the logrus levels that this hook is interested in.
 func (h *LogrusHook) Levels() []log.Level {
 	if h.levels == nil {
 		return allLevels
@@ -102,47 +65,6 @@ func (h *LogrusHook) Levels() []log.Level {
 	return h.levels
 }
 
-// Fire is called when a log event occurs.
-func (h *LogrusHook) Fire(entry *log.Entry) error {
-	if entry.Context == nil {
-		return nil
-	}
-
-	span := trace.SpanFromContext(entry.Context)
-	spanCtx := span.SpanContext()
-	if !spanCtx.IsValid() {
-		return nil
-	}
-
-	if h.enableTraceIDField {
-		entry.Data["trace_id"] = spanCtx.TraceID().String()
-	}
-
-	if !span.IsRecording() {
-		return nil
-	}
-
-	span.AddEvent(
-		entry.Message,
-		trace.WithAttributes(logrusDataToAttrs(entry.Data)...),
-		trace.WithAttributes(attribute.String("level", entry.Level.String())),
-		trace.WithTimestamp(entry.Time),
-	)
-
-	// Set the span status based on the log level, rather than the presence of
-	// an error field. Error values may be attached to lower-severity log entries
-	// without indicating that the operation represented by the span failed.
-	if h.errorStatusLevel != nil && entry.Level <= *h.errorStatusLevel {
-		span.SetStatus(codes.Error, entry.Message)
-	}
-
+func (h *LogrusHook) Fire(_ *log.Entry) error {
 	return nil
-}
-
-func logrusDataToAttrs(data map[string]any) []attribute.KeyValue {
-	attrs := make([]attribute.KeyValue, 0, len(data))
-	for k, v := range data {
-		attrs = append(attrs, keyValue(k, v))
-	}
-	return attrs
 }
