--- pkg/options/options.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/options/options.go
@@ -23,7 +23,6 @@ import (
 	semver "github.com/gruntwork-io/terragrunt/internal/semver"
 	"github.com/gruntwork-io/terragrunt/internal/strict"
 	"github.com/gruntwork-io/terragrunt/internal/strict/controls"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tfimpl"
 	"github.com/gruntwork-io/terragrunt/internal/tips"
 	"github.com/gruntwork-io/terragrunt/internal/util"
@@ -83,8 +82,6 @@ type TerragruntOptions struct {
 	EngineConfig *engine.EngineConfig
 	// EngineOptions groups CLI-supplied engine options.
 	EngineOptions *engine.EngineOptions
-	// Telemetry are telemetry options.
-	Telemetry *telemetry.Options
 	// Attributes to override in AWS provider nested within modules as part of the aws-provider-patch command.
 	AwsProviderPatchOverrides map[string]string
 	// Version of terraform (obtained by running 'terraform version')
@@ -360,7 +357,6 @@ func NewTerragruntOptions(e vexec.Exec) *TerragruntOpt
 		StrictControls:         controls.New(),
 		Experiments:            experiment.NewExperiments(),
 		Tips:                   tips.NewTips(),
-		Telemetry:              new(telemetry.Options),
 		EngineOptions:          new(engine.EngineOptions),
 		VersionManagerFileName: defaultVersionManagerFileName,
 		CASCloneDepth:          1,
