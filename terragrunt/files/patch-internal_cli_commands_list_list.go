--- internal/cli/commands/list/list.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/commands/list/list.go
@@ -9,8 +9,6 @@ import (
 	"sort"
 	"strings"
 
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
-
 	"errors"
 
 	"charm.land/lipgloss/v2/tree"
@@ -22,9 +20,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/view/dag"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // Run runs the list command.
@@ -55,20 +50,8 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 		discoverErr error
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
+	components, discoverErr = d.Discover(ctx, l, v, opts.TerragruntOptions)
+	err = discoverErr
 	if err != nil {
 		l.Debugf("Errors encountered while discovering components:\n%s", err)
 	}
@@ -77,19 +60,13 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
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
+		q, queueErr := queue.NewQueue(components)
+		if queueErr != nil {
+			return queueErr
+		}
 
-			components = q.Components()
-
-			return nil
-		})
+		components = q.Components()
+		err = nil
 		if err != nil {
 			return err
 		}
@@ -101,19 +78,8 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 
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
+	listedComponents = discoveredToListed(l, components, opts)
+	err = nil
 	if err != nil {
 		return err
 	}
