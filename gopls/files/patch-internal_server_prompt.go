--- internal/server/prompt.go.orig	2025-02-24 17:39:34 UTC
+++ internal/server/prompt.go
@@ -14,8 +14,6 @@ import (
 	"testing"
 	"time"
 
-	"golang.org/x/telemetry"
-	"golang.org/x/telemetry/counter"
 	"golang.org/x/tools/gopls/internal/protocol"
 	"golang.org/x/tools/internal/event"
 )
@@ -68,23 +66,14 @@ func (s *server) telemetryMode() string {
 // telemetryMode returns the current effective telemetry mode.
 // By default this is x/telemetry.Mode(), but it may be overridden for tests.
 func (s *server) telemetryMode() string {
-	if fake := s.getenv(FakeTelemetryModefileEnvvar); fake != "" {
-		if data, err := os.ReadFile(fake); err == nil {
-			return string(data)
-		}
 		return "local"
-	}
-	return telemetry.Mode()
 }
 
 // setTelemetryMode sets the current telemetry mode.
 // By default this calls x/telemetry.SetMode, but it may be overridden for
 // tests.
 func (s *server) setTelemetryMode(mode string) error {
-	if fake := s.getenv(FakeTelemetryModefileEnvvar); fake != "" {
-		return os.WriteFile(fake, []byte(mode), 0666)
-	}
-	return telemetry.SetMode(mode)
+	return nil
 }
 
 // maybePromptForTelemetry checks for the right conditions, and then prompts
@@ -172,12 +161,9 @@ func (s *server) maybePromptForTelemetry(ctx context.C
 		//
 		// But record this in telemetry, in case some users enable telemetry by
 		// other means.
-		counter.New("gopls/telemetryprompt/corrupted").Inc()
 		return
 	}
 
-	counter.New(fmt.Sprintf("gopls/telemetryprompt/attempts:%d", attempts)).Inc()
-
 	// Check terminal conditions.
 
 	if state == pYes {
@@ -188,7 +174,6 @@ func (s *server) maybePromptForTelemetry(ctx context.C
 		// counter file at the time telemetry is enabled, we'd never upload it,
 		// because we exclude any counter files that overlap with a time period
 		// that has telemetry uploading is disabled.
-		counter.New("gopls/telemetryprompt/accepted").Inc()
 		return
 	}
 	if state == pNo {
@@ -197,14 +182,12 @@ func (s *server) maybePromptForTelemetry(ctx context.C
 		// other means later on. If we see a significant number of users that have
 		// accepted telemetry but declined the prompt, it may be an indication that
 		// the prompt is not working well.
-		counter.New("gopls/telemetryprompt/declined").Inc()
 		return
 	}
 	if attempts >= 5 { // pPending or pFailed
 		// We've tried asking enough; give up. Record that the prompt expired, in
 		// case the user decides to enable telemetry by other means later on.
 		// (see also the pNo case).
-		counter.New("gopls/telemetryprompt/expired").Inc()
 		return
 	}
 
