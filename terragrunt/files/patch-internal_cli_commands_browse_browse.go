--- internal/cli/commands/browse/browse.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/commands/browse/browse.go
@@ -9,13 +9,9 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/component"
 	"github.com/gruntwork-io/terragrunt/internal/discovery"
 	"github.com/gruntwork-io/terragrunt/internal/os/stdout"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	viewtui "github.com/gruntwork-io/terragrunt/internal/view/tui"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // Run runs the browse command. It opens the browser immediately over a tree the
@@ -108,22 +104,10 @@ func runDiscovery(
 	opts *Options,
 	d *discovery.Discovery,
 ) tui.DiscoveryResult {
-	var (
-		components  component.Components
-		discoverErr error
-	)
+	var components component.Components
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "browse_discover", map[string]any{
-		"working_dir": opts.WorkingDir,
-	}, func(ctx context.Context, l log.Logger) error {
-		components, discoverErr = d.Discover(ctx, l, v, opts.TerragruntOptions)
-
-		if span := trace.SpanFromContext(ctx); span.IsRecording() {
-			span.SetAttributes(attribute.Int("component_count", len(components)))
-		}
-
-		return discoverErr
-	})
+	_ = opts.WorkingDir
+	components, err := d.Discover(ctx, l, v, opts.TerragruntOptions)
 
 	components = components.Sort()
 
