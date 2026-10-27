--- internal/ghcmd/cmd.go.orig	2026-10-26 18:02:47 UTC
+++ internal/ghcmd/cmd.go
@@ -10,7 +10,6 @@ import (
 	"os/exec"
 	"path/filepath"
 	"slices"
-	"strconv"
 	"strings"
 	"time"
 
@@ -19,11 +18,9 @@ import (
 	"github.com/cli/cli/v2/api"
 	"github.com/cli/cli/v2/internal/agents"
 	"github.com/cli/cli/v2/internal/build"
-	"github.com/cli/cli/v2/internal/ci"
 	"github.com/cli/cli/v2/internal/config"
 	"github.com/cli/cli/v2/internal/gh"
 	"github.com/cli/cli/v2/internal/gh/ghtelemetry"
-	"github.com/cli/cli/v2/internal/telemetry"
 	"github.com/cli/cli/v2/internal/update"
 	"github.com/cli/cli/v2/pkg/cmd/auth/shared"
 	"github.com/cli/cli/v2/pkg/cmd/factory"
@@ -70,63 +67,7 @@ func Main() exitCode {
 
 	ghExecutablePath := executablePath("gh")
 
-	additionalCommonDimensions := ghtelemetry.Dimensions{
-		"version":             strings.TrimPrefix(buildVersion, "v"),
-		"is_tty":              strconv.FormatBool(ioStreams.IsStdoutTTY()),
-		"agent":               string(invokingAgent),
-		"ci":                  strconv.FormatBool(ci.IsCI()),
-		"github_actions":      strconv.FormatBool(ci.IsGitHubActions()),
-		"accessible_colors":   strconv.FormatBool(ioStreams.AccessibleColorsEnabled()),
-		"accessible_prompter": strconv.FormatBool(ioStreams.AccessiblePrompterEnabled()),
-		"color_labels":        strconv.FormatBool(ioStreams.ColorLabels()),
-		"spinner_disabled":    strconv.FormatBool(ioStreams.GetSpinnerDisabled()),
-	}
-
-	var telemetryService ghtelemetry.Service
-	switch {
-	case cfgErr != nil:
-		// Without a valid on-disk config we can't honour user telemetry preferences, so disable it to be safe.
-		telemetryService = &telemetry.NoOpService{}
-	default:
-		telemetryState := telemetry.ParseTelemetryState(cfg.Telemetry().Value)
-		telemetryDisabled := mightBeGHESUser(cfg)
-
-		switch telemetryState {
-		case telemetry.Disabled:
-			telemetryService = &telemetry.NoOpService{}
-		case telemetry.Logged:
-			// Always construct the real service in log mode so that the log
-			// flusher runs and surfaces an explicit "Telemetry payload: none"
-			// marker when no events will be sent. This gives the user an
-			// observable signal that telemetry is wired up even when their
-			// context (e.g. GHES) causes events to be dropped.
-			telemetryService = telemetry.NewService(
-				telemetry.LogFlusher(ioStreams.ErrOut, ioStreams.ColorEnabled()),
-				telemetry.WithAdditionalCommonDimensions(additionalCommonDimensions),
-			)
-			if telemetryDisabled {
-				telemetryService.Disable()
-			}
-		case telemetry.Enabled:
-			if telemetryDisabled {
-				telemetryService = &telemetry.NoOpService{}
-				break
-			}
-			sampleRate := 1
-			if v, err := strconv.Atoi(os.Getenv("GH_TELEMETRY_SAMPLE_RATE")); err == nil && v >= 0 && v <= 100 {
-				sampleRate = v
-			}
-			additionalCommonDimensions["sample_rate"] = strconv.Itoa(sampleRate)
-			telemetryService = telemetry.NewService(
-				telemetry.GitHubFlusher(ghExecutablePath),
-				telemetry.WithAdditionalCommonDimensions(additionalCommonDimensions),
-				telemetry.WithSampleRate(sampleRate),
-			)
-		default:
-			fmt.Fprintf(stderr, "invalid telemetry configuration: %q\n", cfg.Telemetry().Value)
-			return exitError
-		}
-	}
+	telemetryService := ghtelemetry.Service(&ghtelemetry.NoOpService{})
 	// Complete and send events even when returning before Cobra reaches RunE.
 	defer telemetryService.Finish()
 
