--- pkg/config/parsing_context.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/config/parsing_context.go
@@ -14,7 +14,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/iam"
 	pcoptions "github.com/gruntwork-io/terragrunt/internal/providercache/options"
 	"github.com/gruntwork-io/terragrunt/internal/strict"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tfimpl"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
@@ -41,8 +40,8 @@ type ParsingContext struct {
 	// FeatureFlags contains explicit feature flag overrides supplied by the user.
 	FeatureFlags map[string]string
 
-	FilesRead *FilesRead
-	Telemetry *telemetry.Options
+	FilesRead   *FilesRead
+	TraceParent string
 
 	DecodedDependencies *cty.Value
 	Values              *cty.Value
