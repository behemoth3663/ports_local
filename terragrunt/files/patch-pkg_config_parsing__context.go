--- pkg/config/parsing_context.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/config/parsing_context.go
@@ -18,8 +18,8 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/iam"
 	pcoptions "github.com/gruntwork-io/terragrunt/internal/providercache/options"
 	"github.com/gruntwork-io/terragrunt/internal/strict"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tfimpl"
+	telemetryopts "github.com/gruntwork-io/terragrunt/internal/traceopts"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/pkg/config/hclparse"
@@ -56,7 +56,7 @@ type ParsingContext struct {
 	FeatureFlags map[string]string
 
 	FilesRead *FilesRead
-	Telemetry *telemetry.Options
+	Telemetry *telemetryopts.Options
 
 	DecodedDependencies *cty.Value
 	Values              *cty.Value
@@ -245,7 +245,7 @@ func (ctx *ParsingContext) WithDiagnosticsSuppressed(l
 // This avoids false positive "There is no variable named dependency" errors during parsing
 // when dependency outputs haven't been resolved yet.
 func (ctx *ParsingContext) WithDiagnosticsSuppressed(l log.Logger) *ParsingContext {
-	var diagWriter = io.Discard
+	diagWriter := io.Discard
 	if l.Level() >= log.DebugLevel {
 		diagWriter = ctx.Venv.Writers.ErrWriter
 	}
