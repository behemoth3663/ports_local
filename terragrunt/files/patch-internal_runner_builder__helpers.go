diff --git a/internal/runner/builder_helpers.go b/internal/runner/builder_helpers.go
index e6c1869de..6573b757d 100644
--- internal/runner/builder_helpers.go.orig
+++ internal/runner/builder_helpers.go
@@ -13,7 +13,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/configbridge"
 	"github.com/gruntwork-io/terragrunt/internal/discovery"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/worktrees"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
@@ -33,7 +32,9 @@ func doWithTelemetry(
 	fields map[string]any,
 	fn func(context.Context, log.Logger) error,
 ) error {
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, l, name, fields, fn)
+	_ = name
+	_ = fields
+	return fn(ctx, l)
 }
 
 // resolveWorkingDir determines the canonical working directory for discovery.
