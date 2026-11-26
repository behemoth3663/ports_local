--- pkg/config/parsing_context.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/config/parsing_context.go
@@ -14,7 +14,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/iam"
 	pcoptions "github.com/gruntwork-io/terragrunt/internal/providercache/options"
 	"github.com/gruntwork-io/terragrunt/internal/strict"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/pkg/config/hclparse"
@@ -142,9 +141,6 @@ type RunSettings struct {
 
 	// FeatureFlags contains explicit feature flag overrides supplied by the user.
 	FeatureFlags map[string]string
-
-	// Telemetry configures the run's traces and metrics.
-	Telemetry *telemetry.Options
 
 	// SourceMap replaces the prefix of a terraform source URL, from the --source-map flag.
 	SourceMap map[string]string
