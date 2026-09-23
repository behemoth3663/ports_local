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
@@ -52,44 +50,23 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 
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
-			components, discoverErr = d.Discover(ctx, l, v, opts.TerragruntOptions)
-			return discoverErr
-		})
-	if telemetryErr != nil {
-		l.Debugf("Errors encountered while discovering components:\n%s", telemetryErr)
+	components, err = d.Discover(ctx, l, v, opts.TerragruntOptions)
+	if err != nil {
+		return err
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
@@ -98,18 +75,7 @@ func Run(ctx context.Context, l log.Logger, v *venv.Ve
 
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
-	if err != nil {
-		return err
-	}
+	foundComponents = discoveredToFound(l, components, opts)
 
 	switch opts.Format {
 	case FormatText:
