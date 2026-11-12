--- internal/runner/builder_helpers.go.orig	1979-11-29 21:00:00 UTC
+++ internal/runner/builder_helpers.go
@@ -15,7 +15,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/discovery"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run"
 	"github.com/gruntwork-io/terragrunt/internal/spinner"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/worktrees"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
@@ -35,7 +34,9 @@ func doWithTelemetry(
 	fields map[string]any,
 	fn func(context.Context, log.Logger) error,
 ) error {
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, l, name, fields, fn)
+	_ = name
+	_ = fields
+	return fn(ctx, l)
 }
 
 // resolveWorkingDir determines the canonical working directory for discovery.
