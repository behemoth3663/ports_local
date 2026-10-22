--- internal/cli/commands/find/find.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/commands/find/find.go
@@ -3,16 +3,14 @@ import (
 import (
 	"context"
 	"encoding/json"
+	"errors"
 	"fmt"
 	"io"
 	"path/filepath"
 	"strings"
 
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 
-	"errors"
-
 	"github.com/gruntwork-io/terragrunt/internal/cli/commands/discoverysetup"
 	"github.com/gruntwork-io/terragrunt/internal/component"
 	"github.com/gruntwork-io/terragrunt/internal/discovery"
@@ -50,22 +48,10 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 		return err
 	}
 
-	var (
-		components  component.Components
-		discoverErr error
-	)
+	var components component.Components
 
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
+	_, _, _, _, _ = opts.WorkingDir, opts.NoHidden, opts.Dependencies, opts.Mode, opts.Exclude
+	components, telemetryErr := d.Discover(ctx, l, v, opts.TerragruntOptions)
 	if telemetryErr != nil {
 		l.Debugf("Errors encountered while discovering components:\n%s", telemetryErr)
 	}
@@ -74,19 +60,14 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
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
+		_, _ = opts.WorkingDir, len(components)
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
@@ -98,15 +79,9 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 
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
+	_, _ = opts.WorkingDir, len(components)
+	foundComponents = discoveredToFound(l, components, opts)
+	err = nil
 	if err != nil {
 		return err
 	}
