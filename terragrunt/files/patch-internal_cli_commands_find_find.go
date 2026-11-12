--- internal/cli/commands/find/find.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/commands/find/find.go
@@ -9,7 +9,6 @@ import (
 	"strings"
 
 	rewrite "github.com/gruntwork-io/terragrunt/internal/config/rewrite"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 
 	"errors"
@@ -54,51 +53,24 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 
 	var (
 		components  component.Components
-		discoverErr error
 	)
 
-	telemetryErr := telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "find_discover", map[string]any{
-			"working_dir":  opts.WorkingDir,
-			"no_hidden":    opts.NoHidden,
-			"dependencies": opts.Dependencies,
-			"mode":         opts.Mode,
-			"exclude":      opts.Exclude,
-		}, func(ctx context.Context, l log.Logger) error {
-			discoverErr = spinner.ShowAfter(ctx, l, spinner.Messages{
-				Working: "Discovering units...",
-				Done:    "Discovered units",
-			}, func() error {
-				var err error
-
 				components, err = d.Discover(ctx, l, v, s, opts.TerragruntOptions)
-
-				return err
-			})
-
-			return discoverErr
-		})
-	if telemetryErr != nil {
-		l.Debugf("Errors encountered while discovering components:\n%s", telemetryErr)
+	if err != nil {
+		l.Debugf("Errors encountered while discovering components:\n%s", err)
 	}
 
 	switch opts.Mode {
 	case ModeNormal:
 		components = components.Sort()
 	case ModeDAG:
-		err = telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "find_mode_dag", map[string]any{
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
@@ -110,15 +82,8 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 
 	var foundComponents FoundComponents
 
-	err = telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "find_discovered_to_found", map[string]any{
-			"working_dir":  opts.WorkingDir,
-			"config_count": len(components),
-		}, func(ctx context.Context, l log.Logger) error {
-			foundComponents = discoveredToFound(l, components, opts)
-
-			return nil
-		})
+	foundComponents = discoveredToFound(l, components, opts)
+	err = nil
 	if err != nil {
 		return err
 	}
