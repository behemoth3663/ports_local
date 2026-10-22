--- internal/runner/run/options.go.orig	1979-11-29 21:00:00 UTC
+++ internal/runner/run/options.go
@@ -3,12 +3,11 @@ import (
 import (
 	"context"
 	"encoding/json"
+	"errors"
 	"fmt"
 	"path/filepath"
 	"time"
 
-	"errors"
-
 	"github.com/gruntwork-io/terragrunt/internal/cloner"
 	"github.com/gruntwork-io/terragrunt/internal/engine"
 	"github.com/gruntwork-io/terragrunt/internal/errorconfig"
@@ -19,7 +18,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/report"
 	"github.com/gruntwork-io/terragrunt/internal/shell"
 	"github.com/gruntwork-io/terragrunt/internal/strict"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tf"
 	"github.com/gruntwork-io/terragrunt/internal/tfimpl"
 	"github.com/gruntwork-io/terragrunt/internal/tflint"
@@ -46,7 +44,7 @@ type Options struct {
 	EngineOptions                *engine.EngineOptions
 	Errors                       *errorconfig.Config
 	FeatureFlags                 map[string]string
-	Telemetry                    *telemetry.Options
+	TraceParent                  string
 	SourceMap                    map[string]string
 	TFPath                       string
 	TerraformCommand             string
@@ -199,13 +197,15 @@ func (o *Options) shellRunOptions(env map[string]strin
 	s := shell.NewShellOptions(env).
 		WithWorkingDir(o.CacheDir).
 		WithUnitDir(o.UnitDir).
-		WithTelemetry(o.Telemetry).
 		WithEngine(o.EngineConfig, o.EngineOptions).
 		WithTFPath(o.TFPath).
 		WithRootWorkingDir(o.RootWorkingDir).
 		WithExperiments(o.Experiments).
 		WithHeadless(o.Headless).
 		WithForwardTFStdout(o.ForwardTFStdout)
+	if o.TraceParent != "" {
+		s.SetTraceParent(o.TraceParent)
+	}
 	s.LogShowAbsPaths = o.LogShowAbsPaths
 	s.LogDisableErrorSummary = o.LogDisableErrorSummary
 
@@ -367,7 +367,7 @@ func (o *Options) handleIgnoreSignals(l log.Logger, fs
 		return err
 	}
 
-	const ownerPerms = 0644
+	const ownerPerms = 0o644
 
 	l.Warnf("Writing error signals to %s", signalsFile)
 
