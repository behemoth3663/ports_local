--- internal/cli/commands/find/find.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/commands/find/find.go
@@ -8,7 +8,6 @@ import (
 	"path/filepath"
 	"strings"
 
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 
 	"errors"
@@ -55,17 +54,8 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 		discoverErr error
 	)
 
-	telemetryErr := telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "find_discover", map[string]any{
-			"working_dir":  opts.WorkingDir,
-			"no_hidden":    opts.NoHidden,
-			"dependencies": opts.Dependencies,
-			"mode":         opts.Mode,
-			"exclude":      opts.Exclude,
-		}, func(ctx context.Context, l log.Logger) error {
-			components, discoverErr = d.Discover(ctx, l, v, opts.TerragruntOptions)
-			return discoverErr
-		})
+	components, discoverErr = d.Discover(ctx, l, v, opts.TerragruntOptions)
+	telemetryErr := discoverErr
 	if telemetryErr != nil {
 		l.Debugf("Errors encountered while discovering components:\n%s", telemetryErr)
 	}
@@ -74,19 +64,13 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
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
@@ -98,15 +82,8 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 
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
