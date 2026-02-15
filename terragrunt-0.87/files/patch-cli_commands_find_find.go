--- cli/commands/find/find.go.orig	2025-09-25 14:56:12 UTC
+++ cli/commands/find/find.go
@@ -6,7 +6,6 @@ import (
 	"path/filepath"
 
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/util"
 
 	"github.com/gruntwork-io/terragrunt/config"
@@ -50,21 +49,7 @@ func Run(ctx context.Context, l log.Logger, opts *Opti
 		})
 	}
 
-	var cfgs discovery.DiscoveredConfigs
-
-	var discoverErr error
-
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "find_discover", map[string]any{
-		"working_dir":  opts.WorkingDir,
-		"hidden":       opts.Hidden,
-		"dependencies": opts.Dependencies,
-		"external":     opts.External,
-		"mode":         opts.Mode,
-		"exclude":      opts.Exclude,
-	}, func(ctx context.Context) error {
-		cfgs, discoverErr = d.Discover(ctx, l, opts.TerragruntOptions)
-		return discoverErr
-	})
+	cfgs, err := d.Discover(ctx, l, opts.TerragruntOptions)
 	if err != nil {
 		l.Debugf("Errors encountered while discovering configurations:\n%s", err)
 	}
@@ -73,40 +58,18 @@ func Run(ctx context.Context, l log.Logger, opts *Opti
 	case ModeNormal:
 		cfgs = cfgs.Sort()
 	case ModeDAG:
-		err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "find_mode_dag", map[string]any{
-			"working_dir":  opts.WorkingDir,
-			"config_count": len(cfgs),
-		}, func(ctx context.Context) error {
-			q, queueErr := queue.NewQueue(cfgs)
-			if queueErr != nil {
-				return queueErr
-			}
-
-			cfgs = q.Configs()
-
-			return nil
-		})
+		q, err := queue.NewQueue(cfgs)
 		if err != nil {
 			return errors.New(err)
 		}
+		cfgs = q.Configs()
 	default:
 		// This should never happen, because of validation in the command.
 		// If it happens, we want to throw so we can fix the validation.
 		return errors.New("invalid mode: " + opts.Mode)
 	}
 
-	var foundCfgs FoundConfigs
-
-	err = telemetry.TelemeterFromContext(ctx).Collect(ctx, "find_discovered_to_found", map[string]any{
-		"working_dir":  opts.WorkingDir,
-		"config_count": len(cfgs),
-	}, func(ctx context.Context) error {
-		var convErr error
-
-		foundCfgs, convErr = discoveredToFound(cfgs, opts)
-
-		return convErr
-	})
+	foundCfgs, err := discoveredToFound(cfgs, opts)
 	if err != nil {
 		return errors.New(err)
 	}
