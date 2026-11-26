--- internal/cli/commands/list/list.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/commands/list/list.go
@@ -9,8 +9,6 @@ import (
 	"sort"
 	"strings"
 
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
-
 	"errors"
 
 	"charm.land/lipgloss/v2/tree"
@@ -24,9 +22,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/view/dag"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // Run runs the list command.
@@ -52,33 +47,15 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 		return err
 	}
 
-	var (
-		components  component.Components
-		discoverErr error
-	)
+	var components component.Components
 
-	// Wrap discovery with telemetry
-	err = telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "list_discover", map[string]any{
-		"working_dir":  opts.WorkingDir,
-		"no_hidden":    opts.NoHidden,
-		"dependencies": opts.Dependencies || opts.Mode == ModeDAG,
-	}, func(ctx context.Context, l log.Logger) error {
-		discoverErr = spinner.ShowAfter(ctx, l, spinner.Messages{
-			Working: "Discovering units...",
-			Done:    "Discovered units",
-		}, func() error {
-			var err error
-
-			components, err = d.Discover(ctx, l, v, s, opts.TerragruntOptions)
-
-			return err
-		})
-
-		if span := trace.SpanFromContext(ctx); span.IsRecording() {
-			span.SetAttributes(attribute.Int("component_count", len(components)))
-		}
-
-		return discoverErr
+	err = spinner.ShowAfter(ctx, l, spinner.Messages{
+		Working: "Discovering units...",
+		Done:    "Discovered units",
+	}, func() error {
+		var err error
+		components, err = d.Discover(ctx, l, v, s, opts.TerragruntOptions)
+		return err
 	})
 	if err != nil {
 		// Discover returns no components with this error, so logging it
@@ -94,22 +71,12 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 	case ModeNormal:
 		components = components.Sort()
 	case ModeDAG:
-		err = telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "list_mode_dag", map[string]any{
-			"working_dir":  opts.WorkingDir,
-			"config_count": len(components),
-		}, func(ctx context.Context, l log.Logger) error {
-			q, queueErr := queue.NewQueue(components)
-			if queueErr != nil {
-				return queueErr
-			}
-
-			components = q.Components()
-
-			return nil
-		})
-		if err != nil {
-			return err
+		q, queueErr := queue.NewQueue(components)
+		if queueErr != nil {
+			return queueErr
 		}
+
+		components = q.Components()
 	default:
 		// This should never happen, because of validation in the command.
 		// If it happens, we want to throw so we can fix the validation.
@@ -118,22 +85,7 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 
 	var listedComponents dag.ListedComponents
 
-	err = telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "list_discovered_to_listed", map[string]any{
-			"working_dir":  opts.WorkingDir,
-			"config_count": len(components),
-		}, func(ctx context.Context, l log.Logger) error {
-			listedComponents = discoveredToListed(l, components, opts)
-
-			if span := trace.SpanFromContext(ctx); span.IsRecording() {
-				span.SetAttributes(attribute.Int("listed_count", len(listedComponents)))
-			}
-
-			return nil
-		})
-	if err != nil {
-		return err
-	}
+	listedComponents = discoveredToListed(l, components, opts)
 
 	switch opts.Format {
 	case FormatText:
