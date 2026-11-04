--- internal/runner/run/options.go.orig	1979-11-29 21:00:00 UTC
+++ internal/runner/run/options.go
@@ -19,7 +19,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/report"
 	"github.com/gruntwork-io/terragrunt/internal/shell"
 	"github.com/gruntwork-io/terragrunt/internal/strict"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tf"
 	"github.com/gruntwork-io/terragrunt/internal/tfimpl"
 	"github.com/gruntwork-io/terragrunt/internal/tflint"
@@ -46,7 +45,6 @@ type Options struct {
 	EngineOptions                *engine.EngineOptions
 	Errors                       *errorconfig.Config
 	FeatureFlags                 map[string]string
-	Telemetry                    *telemetry.Options
 	SourceMap                    map[string]string
 	TFPath                       string
 	TerraformCommand             string
@@ -199,7 +197,6 @@ func (o *Options) shellRunOptions(env map[string]strin
 	s := shell.NewShellOptions(env).
 		WithWorkingDir(o.CacheDir).
 		WithUnitDir(o.UnitDir).
-		WithTelemetry(o.Telemetry).
 		WithEngine(o.EngineConfig, o.EngineOptions).
 		WithTFPath(o.TFPath).
 		WithRootWorkingDir(o.RootWorkingDir).
