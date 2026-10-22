--- pkg/options/options.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/options/options.go
@@ -4,13 +4,12 @@ import (
 import (
 	"context"
 	"encoding/json"
+	"errors"
 	"fmt"
 	"math"
 	"path/filepath"
 	"time"
 
-	"errors"
-
 	"github.com/gruntwork-io/terragrunt/internal/cloner"
 	"github.com/gruntwork-io/terragrunt/internal/engine"
 	"github.com/gruntwork-io/terragrunt/internal/errorconfig"
@@ -23,7 +22,6 @@ import (
 	semver "github.com/gruntwork-io/terragrunt/internal/semver"
 	"github.com/gruntwork-io/terragrunt/internal/strict"
 	"github.com/gruntwork-io/terragrunt/internal/strict/controls"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tfimpl"
 	"github.com/gruntwork-io/terragrunt/internal/tips"
 	"github.com/gruntwork-io/terragrunt/internal/util"
@@ -62,17 +60,26 @@ const (
 	DefaultLogLevel = log.InfoLevel
 )
 
-var (
-	defaultVersionManagerFileName = []string{
-		".terraform-version",
-		".tool-versions",
-		"mise.toml",
-		".mise.toml",
-	}
-)
+var defaultVersionManagerFileName = []string{
+	".terraform-version",
+	".tool-versions",
+	"mise.toml",
+	".mise.toml",
+}
 
 type ctxKey byte
 
+type TelemetryOptions struct {
+	TraceExporter                  string
+	TraceExporterInsecureEndpoint  bool
+	TraceExporterHTTPEndpoint      string
+	TraceParent                    string
+	MetricExporter                 string
+	MetricExporterInsecureEndpoint bool
+	LogsExporter                   string
+	LogsExporterInsecureEndpoint   bool
+}
+
 // TerragruntOptions represents options that configure the behavior of the Terragrunt program
 type TerragruntOptions struct {
 	// Version of terragrunt
@@ -84,7 +91,7 @@ type TerragruntOptions struct {
 	// EngineOptions groups CLI-supplied engine options.
 	EngineOptions *engine.EngineOptions
 	// Telemetry are telemetry options.
-	Telemetry *telemetry.Options
+	Telemetry *TelemetryOptions
 	// Attributes to override in AWS provider nested within modules as part of the aws-provider-patch command.
 	AwsProviderPatchOverrides map[string]string
 	// Version of terraform (obtained by running 'terraform version')
@@ -360,7 +367,7 @@ func NewTerragruntOptions(e vexec.Exec) *TerragruntOpt
 		StrictControls:         controls.New(),
 		Experiments:            experiment.NewExperiments(),
 		Tips:                   tips.NewTips(),
-		Telemetry:              new(telemetry.Options),
+		Telemetry:              new(TelemetryOptions),
 		EngineOptions:          new(engine.EngineOptions),
 		VersionManagerFileName: defaultVersionManagerFileName,
 		CASCloneDepth:          1,
@@ -721,7 +728,7 @@ func (opts *TerragruntOptions) handleIgnoreSignals(l l
 		return err
 	}
 
-	const ownerPerms = 0644
+	const ownerPerms = 0o644
 
 	l.Warnf("Writing error signals to %s", signalsFile)
 
