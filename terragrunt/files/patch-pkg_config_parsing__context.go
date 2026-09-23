diff --git a/pkg/config/parsing_context.go b/pkg/config/parsing_context.go
index e5a3d24cb..51103f705 100644
--- pkg/config/parsing_context.go.orig
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
@@ -245,7 +245,7 @@ func (ctx *ParsingContext) WithParseOption(parserOptions []hclparse.Option) *Par
 // This avoids false positive "There is no variable named dependency" errors during parsing
 // when dependency outputs haven't been resolved yet.
 func (ctx *ParsingContext) WithDiagnosticsSuppressed(l log.Logger) *ParsingContext {
-	var diagWriter = io.Discard
+	diagWriter := io.Discard
 	if l.Level() >= log.DebugLevel {
 		diagWriter = ctx.Venv.Writers.ErrWriter
 	}
