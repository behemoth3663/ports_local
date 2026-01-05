--- options/options.go.orig	2026-01-05 15:05:23 UTC
+++ options/options.go
@@ -25,7 +25,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"github.com/gruntwork-io/terragrunt/pkg/log/format"
 	"github.com/gruntwork-io/terragrunt/pkg/log/format/placeholders"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/util"
 	"github.com/hashicorp/go-version"
 	"github.com/puzpuzpuz/xsync/v3"
@@ -109,8 +108,6 @@ type TerragruntOptions struct {
 	FeatureFlags *xsync.MapOf[string, string] `clone:"shadowcopy"`
 	// Options to use engine for running IaC operations.
 	Engine *EngineOptions
-	// Telemetry are telemetry options.
-	Telemetry *telemetry.Options
 	// Attributes to override in AWS provider nested within modules as part of the aws-provider-patch command.
 	AwsProviderPatchOverrides map[string]string
 	// A command that can be used to run Terragrunt with the given options.
@@ -397,7 +394,6 @@ func NewTerragruntOptionsWithWriters(stdout, stderr io
 		Errors:                     defaultErrorsConfig(),
 		StrictControls:             controls.New(),
 		Experiments:                experiment.NewExperiments(),
-		Telemetry:                  new(telemetry.Options),
 		VersionManagerFileName:     defaultVersionManagerFileName,
 	}
 }
