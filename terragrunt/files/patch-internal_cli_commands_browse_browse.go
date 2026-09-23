--- internal/cli/commands/browse/browse.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/commands/browse/browse.go
@@ -6,16 +6,11 @@ import (
 
 	"github.com/gruntwork-io/terragrunt/internal/cli/commands/browse/tui"
 	"github.com/gruntwork-io/terragrunt/internal/cli/commands/discoverysetup"
-	"github.com/gruntwork-io/terragrunt/internal/component"
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
@@ -102,22 +97,7 @@ func runDiscovery(
 	opts *Options,
 	d *discovery.Discovery,
 ) tui.DiscoveryResult {
-	var (
-		components  component.Components
-		discoverErr error
-	)
-
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
+	components, err := d.Discover(ctx, l, v, opts.TerragruntOptions)
 
 	components = components.Sort()
 
