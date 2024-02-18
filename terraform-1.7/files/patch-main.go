--- main.go.orig	2024-02-08 01:00:10 UTC
+++ main.go
@@ -14,7 +14,6 @@ import (
 	"runtime"
 	"strings"
 
-	"github.com/apparentlymart/go-shquot/shquot"
 	"github.com/hashicorp/cli"
 	"github.com/hashicorp/go-plugin"
 	"github.com/hashicorp/terraform-svchost/disco"
@@ -28,7 +27,6 @@ import (
 	"github.com/hashicorp/terraform/version"
 	"github.com/mattn/go-shellwords"
 	"github.com/mitchellh/colorstring"
-	"go.opentelemetry.io/otel/trace"
 
 	backendInit "github.com/hashicorp/terraform/internal/backend/init"
 )
@@ -69,23 +67,7 @@ func realMain() int {
 
 	var err error
 
-	err = openTelemetryInit()
-	if err != nil {
-		// openTelemetryInit can only fail if Terraform was run with an
-		// explicit environment variable to enable telemetry collection,
-		// so in typical use we cannot get here.
-		Ui.Error(fmt.Sprintf("Could not initialize telemetry: %s", err))
-		Ui.Error(fmt.Sprintf("Unset environment variable %s if you don't intend to collect telemetry from Terraform.", openTelemetryExporterEnvVar))
-		return 1
-	}
 	var ctx context.Context
-	var otelSpan trace.Span
-	{
-		// At minimum we emit a span covering the entire command execution.
-		_, displayArgs := shquot.POSIXShellSplit(os.Args)
-		ctx, otelSpan = tracer.Start(context.Background(), fmt.Sprintf("terraform %s", displayArgs))
-		defer otelSpan.End()
-	}
 
 	tmpLogPath := os.Getenv(envTmpLogPath)
 	if tmpLogPath != "" {
