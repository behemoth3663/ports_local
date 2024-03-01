--- cmd/tofu/main.go.orig	2024-02-22 11:32:48 UTC
+++ cmd/tofu/main.go
@@ -14,7 +14,6 @@ import (
 	"runtime"
 	"strings"
 
-	"github.com/apparentlymart/go-shquot/shquot"
 	"github.com/hashicorp/go-plugin"
 	"github.com/hashicorp/terraform-svchost/disco"
 	"github.com/mattn/go-shellwords"
@@ -28,7 +27,6 @@ import (
 	"github.com/opentofu/opentofu/internal/logging"
 	"github.com/opentofu/opentofu/internal/terminal"
 	"github.com/opentofu/opentofu/version"
-	"go.opentelemetry.io/otel/trace"
 
 	backendInit "github.com/opentofu/opentofu/internal/backend/init"
 )
@@ -69,23 +67,7 @@ func realMain() int {
 
 	var err error
 
-	err = openTelemetryInit()
-	if err != nil {
-		// openTelemetryInit can only fail if OpenTofu was run with an
-		// explicit environment variable to enable telemetry collection,
-		// so in typical use we cannot get here.
-		Ui.Error(fmt.Sprintf("Could not initialize telemetry: %s", err))
-		Ui.Error(fmt.Sprintf("Unset environment variable %s if you don't intend to collect telemetry from OpenTofu.", openTelemetryExporterEnvVar))
-		return 1
-	}
 	var ctx context.Context
-	var otelSpan trace.Span
-	{
-		// At minimum we emit a span covering the entire command execution.
-		_, displayArgs := shquot.POSIXShellSplit(os.Args)
-		ctx, otelSpan = tracer.Start(context.Background(), fmt.Sprintf("tofu %s", displayArgs))
-		defer otelSpan.End()
-	}
 
 	tmpLogPath := os.Getenv(envTmpLogPath)
 	if tmpLogPath != "" {
