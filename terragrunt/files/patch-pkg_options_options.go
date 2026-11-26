--- pkg/options/options.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/options/options.go
@@ -23,7 +23,6 @@ import (
 	semver "github.com/gruntwork-io/terragrunt/internal/semver"
 	"github.com/gruntwork-io/terragrunt/internal/strict"
 	"github.com/gruntwork-io/terragrunt/internal/strict/controls"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tips"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
@@ -81,8 +80,6 @@ type TerragruntOptions struct {
 	EngineConfig *engine.EngineConfig
 	// EngineOptions groups CLI-supplied engine options.
 	EngineOptions *engine.EngineOptions
-	// Telemetry are telemetry options.
-	Telemetry *telemetry.Options
 	// Attributes to override in AWS provider nested within modules as part of the aws-provider-patch command.
 	AwsProviderPatchOverrides map[string]string
 	// Errors is a configuration for error handling.
@@ -356,7 +353,6 @@ func NewTerragruntOptions() *TerragruntOptions {
 		StrictControls:         controls.New(),
 		Experiments:            experiment.NewExperiments(),
 		Tips:                   tips.NewTips(),
-		Telemetry:              new(telemetry.Options),
 		EngineOptions:          new(engine.EngineOptions),
 		VersionManagerFileName: defaultVersionManagerFileName,
 		CASCloneDepth:          1,
