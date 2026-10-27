--- internal/gh/ghtelemetry/telemetry.go.orig	2026-10-26 18:02:47 UTC
+++ internal/gh/ghtelemetry/telemetry.go
@@ -43,3 +43,28 @@ const SAMPLE_ALL = 100
 }
 
 const SAMPLE_ALL = 100
+
+type noOpPendingEvent struct{}
+
+func (noOpPendingEvent) UpsertDimensions(Dimensions) {}
+func (noOpPendingEvent) UpsertMeasures(Measures)     {}
+
+// NoOpService discards telemetry when collection is disabled.
+type NoOpService struct{}
+
+// Record discards the event.
+func (s *NoOpService) Record(Event) {}
+
+// Begin returns an inert handle without retaining the event.
+func (s *NoOpService) Begin(Event) PendingEvent {
+	return noOpPendingEvent{}
+}
+
+// Disable leaves telemetry disabled.
+func (s *NoOpService) Disable() {}
+
+// SetSampleRate leaves telemetry disabled.
+func (s *NoOpService) SetSampleRate(int) {}
+
+// Finish has no payload to complete.
+func (s *NoOpService) Finish() {}
