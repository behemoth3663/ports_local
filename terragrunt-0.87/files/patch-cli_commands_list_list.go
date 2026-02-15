--- cli/commands/list/list.go.orig	2025-09-25 14:56:12 UTC
+++ cli/commands/list/list.go
@@ -7,8 +7,6 @@ import (
 	"sort"
 	"strings"
 
-	"github.com/gruntwork-io/terragrunt/telemetry"
-
 	"github.com/charmbracelet/lipgloss"
 	"github.com/charmbracelet/lipgloss/tree"
 	"github.com/charmbracelet/x/term"
@@ -46,20 +44,7 @@ func Run(ctx context.Context, l log.Logger, opts *Opti
 		})
 	}
 
-	var cfgs discovery.DiscoveredConfigs
-
-	var discoverErr error
-
-	// Wrap discovery with telemetry
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "list_discover", map[string]any{
-		"working_dir":  opts.WorkingDir,
-		"hidden":       opts.Hidden,
-		"dependencies": shouldDiscoverDependencies(opts),
-		"external":     opts.External,
-	}, func(ctx context.Context) error {
-		cfgs, discoverErr = d.Discover(ctx, l, opts.TerragruntOptions)
-		return discoverErr
-	})
+	cfgs, err := d.Discover(ctx, l, opts.TerragruntOptions)
 	if err != nil {
 		l.Debugf("Errors encountered while discovering configurations:\n%s", err)
 	}
@@ -68,40 +53,18 @@ func Run(ctx context.Context, l log.Logger, opts *Opti
 	case ModeNormal:
 		cfgs = cfgs.Sort()
 	case ModeDAG:
-		err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "list_mode_dag", map[string]any{
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
 
-	var listedCfgs ListedConfigs
-
-	err = telemetry.TelemeterFromContext(ctx).Collect(ctx, "list_discovered_to_listed", map[string]any{
-		"working_dir":  opts.WorkingDir,
-		"config_count": len(cfgs),
-	}, func(ctx context.Context) error {
-		var convErr error
-
-		listedCfgs, convErr = discoveredToListed(cfgs, opts)
-
-		return convErr
-	})
+	listedCfgs, err := discoveredToListed(cfgs, opts)
 	if err != nil {
 		return errors.New(err)
 	}
