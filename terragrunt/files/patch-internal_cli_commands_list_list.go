--- internal/cli/commands/list/list.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/commands/list/list.go
@@ -2,6 +2,7 @@ import (
 
 import (
 	"context"
+	"errors"
 	"fmt"
 	"io"
 	"os"
@@ -9,10 +10,6 @@ import (
 	"sort"
 	"strings"
 
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
-
-	"errors"
-
 	"charm.land/lipgloss/v2/tree"
 	"github.com/gruntwork-io/terragrunt/internal/cli/commands/discoverysetup"
 	"github.com/gruntwork-io/terragrunt/internal/component"
@@ -22,9 +19,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/view/dag"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // Run runs the list command.
@@ -52,23 +46,9 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 
 	var (
 		components  component.Components
-		discoverErr error
 	)
 
-	// Wrap discovery with telemetry
-	err = telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "list_discover", map[string]any{
-		"working_dir":  opts.WorkingDir,
-		"no_hidden":    opts.NoHidden,
-		"dependencies": opts.Dependencies || opts.Mode == ModeDAG,
-	}, func(ctx context.Context, l log.Logger) error {
-		components, discoverErr = d.Discover(ctx, l, v, opts.TerragruntOptions)
-
-		if span := trace.SpanFromContext(ctx); span.IsRecording() {
-			span.SetAttributes(attribute.Int("component_count", len(components)))
-		}
-
-		return discoverErr
-	})
+	components, err = d.Discover(ctx, l, v, opts.TerragruntOptions)
 	if err != nil {
 		l.Debugf("Errors encountered while discovering components:\n%s", err)
 	}
@@ -77,22 +57,12 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
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
+		q, err := queue.NewQueue(components)
 		if err != nil {
 			return err
 		}
+
+		components = q.Components()
 	default:
 		// This should never happen, because of validation in the command.
 		// If it happens, we want to throw so we can fix the validation.
@@ -101,22 +71,7 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 
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
