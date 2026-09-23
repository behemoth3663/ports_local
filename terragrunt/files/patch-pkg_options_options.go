diff --git a/pkg/options/options.go b/pkg/options/options.go
index 6874da231..9ba9a3309 100644
--- pkg/options/options.go.orig
+++ pkg/options/options.go
@@ -4,13 +4,12 @@ package options
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
@@ -23,9 +22,9 @@ import (
 	semver "github.com/gruntwork-io/terragrunt/internal/semver"
 	"github.com/gruntwork-io/terragrunt/internal/strict"
 	"github.com/gruntwork-io/terragrunt/internal/strict/controls"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tfimpl"
 	"github.com/gruntwork-io/terragrunt/internal/tips"
+	telemetryopts "github.com/gruntwork-io/terragrunt/internal/traceopts"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/vexec"
@@ -62,14 +61,12 @@ const (
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
 
@@ -83,8 +80,8 @@ type TerragruntOptions struct {
 	EngineConfig *engine.EngineConfig
 	// EngineOptions groups CLI-supplied engine options.
 	EngineOptions *engine.EngineOptions
-	// Telemetry are telemetry options.
-	Telemetry *telemetry.Options
+	// Telemetry holds inert telemetry-related CLI options.
+	Telemetry *telemetryopts.Options
 	// Attributes to override in AWS provider nested within modules as part of the aws-provider-patch command.
 	AwsProviderPatchOverrides map[string]string
 	// Version of terraform (obtained by running 'terraform version')
@@ -361,7 +358,7 @@ func NewTerragruntOptions(e vexec.Exec) *TerragruntOptions {
 		StrictControls:         controls.New(),
 		Experiments:            experiment.NewExperiments(),
 		Tips:                   tips.NewTips(),
-		Telemetry:              new(telemetry.Options),
+		Telemetry:              new(telemetryopts.Options),
 		EngineOptions:          new(engine.EngineOptions),
 		VersionManagerFileName: defaultVersionManagerFileName,
 		CASCloneDepth:          1,
@@ -724,7 +721,7 @@ func (opts *TerragruntOptions) handleIgnoreSignals(l log.Logger, fsys vfs.FS, si
 		return err
 	}
 
-	const ownerPerms = 0644
+	const ownerPerms = 0o644
 
 	l.Warnf("Writing error signals to %s", signalsFile)
 
