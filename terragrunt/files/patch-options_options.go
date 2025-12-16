--- options/options.go.orig	2025-12-15 17:07:12 UTC
+++ options/options.go
@@ -24,7 +24,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"github.com/gruntwork-io/terragrunt/pkg/log/format"
 	"github.com/gruntwork-io/terragrunt/pkg/log/format/placeholders"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/util"
 	"github.com/hashicorp/go-version"
 	"github.com/puzpuzpuz/xsync/v3"
@@ -108,8 +107,6 @@ type TerragruntOptions struct {
 	FeatureFlags *xsync.MapOf[string, string] `clone:"shadowcopy"`
 	// Options to use engine for running IaC operations.
 	Engine *EngineOptions
-	// Telemetry are telemetry options.
-	Telemetry *telemetry.Options
 	// Attributes to override in AWS provider nested within modules as part of the aws-provider-patch command.
 	AwsProviderPatchOverrides map[string]string
 	// A command that can be used to run Terragrunt with the given options.
@@ -433,7 +430,6 @@ func NewTerragruntOptionsWithWriters(stdout, stderr io
 		Errors:                           defaultErrorsConfig(),
 		StrictControls:                   controls.New(),
 		Experiments:                      experiment.NewExperiments(),
-		Telemetry:                        new(telemetry.Options),
 		NoStackValidate:                  false,
 		NoStackGenerate:                  false,
 		VersionManagerFileName:           defaultVersionManagerFileName,
