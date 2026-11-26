--- internal/cli/commands/find/find.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/commands/find/find.go
@@ -8,7 +8,6 @@ import (
 	"path/filepath"
 	"strings"
 
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 
 	"errors"
@@ -18,7 +17,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/discovery"
 	"github.com/gruntwork-io/terragrunt/internal/os/stdout"
 	"github.com/gruntwork-io/terragrunt/internal/queue"
-	"github.com/gruntwork-io/terragrunt/internal/spinner"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
 
@@ -53,60 +51,22 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 
 	var (
 		components  component.Components
-		discoverErr error
 	)
-
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
-		// Discover returns no components with this error, so logging it
-		// and carrying on would print an empty result.
-		if coexistenceErr, ok := errors.AsType[discovery.CoexistenceError](telemetryErr); ok {
-			return coexistenceErr
-		}
-
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
+			q, err := queue.NewQueue(components)
+			if err != nil {
+				return err
 			}
 
 			components = q.Components()
-
-			return nil
-		})
-		if err != nil {
-			return err
-		}
 	default:
 		// This should never happen, because of validation in the command.
 		// If it happens, we want to throw so we can fix the validation.
@@ -115,15 +75,8 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 
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
