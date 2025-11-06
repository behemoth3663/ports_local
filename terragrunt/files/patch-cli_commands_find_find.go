--- cli/commands/find/find.go.orig	2025-11-05 17:03:04 UTC
+++ cli/commands/find/find.go
@@ -7,7 +7,6 @@ import (
 	"strings"
 
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/util"
 
 	"github.com/gruntwork-io/terragrunt/config"
@@ -42,17 +41,10 @@ func Run(ctx context.Context, l log.Logger, opts *Opti
 		discoverErr error
 	)
 
-	telemetryErr := telemetry.TelemeterFromContext(ctx).Collect(ctx, "find_discover", map[string]any{
-		"working_dir":  opts.WorkingDir,
-		"hidden":       opts.Hidden,
-		"dependencies": opts.Dependencies,
-		"external":     opts.External,
-		"mode":         opts.Mode,
-		"exclude":      opts.Exclude,
-	}, func(ctx context.Context) error {
+	telemetryErr := func(ctx context.Context) error {
 		components, discoverErr = d.Discover(ctx, l, opts.TerragruntOptions)
 		return discoverErr
-	})
+	}(ctx)
 	if telemetryErr != nil {
 		l.Debugf("Errors encountered while discovering components:\n%s", telemetryErr)
 	}
@@ -61,10 +53,7 @@ func Run(ctx context.Context, l log.Logger, opts *Opti
 	case ModeNormal:
 		components = components.Sort()
 	case ModeDAG:
-		err = telemetry.TelemeterFromContext(ctx).Collect(ctx, "find_mode_dag", map[string]any{
-			"working_dir":  opts.WorkingDir,
-			"config_count": len(components),
-		}, func(ctx context.Context) error {
+		err = func(ctx context.Context) error {
 			q, queueErr := queue.NewQueue(components)
 			if queueErr != nil {
 				return queueErr
@@ -73,7 +62,7 @@ func Run(ctx context.Context, l log.Logger, opts *Opti
 			components = q.Components()
 
 			return nil
-		})
+		}(ctx)
 		if err != nil {
 			return errors.New(err)
 		}
@@ -85,16 +74,13 @@ func Run(ctx context.Context, l log.Logger, opts *Opti
 
 	var foundComponents FoundComponents
 
-	err = telemetry.TelemeterFromContext(ctx).Collect(ctx, "find_discovered_to_found", map[string]any{
-		"working_dir":  opts.WorkingDir,
-		"config_count": len(components),
-	}, func(ctx context.Context) error {
+	err = func(ctx context.Context) error {
 		var convErr error
 
 		foundComponents, convErr = discoveredToFound(components, opts)
 
 		return convErr
-	})
+	}(ctx)
 	if err != nil {
 		return errors.New(err)
 	}
