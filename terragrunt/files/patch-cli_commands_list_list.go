--- cli/commands/list/list.go.orig	2025-11-05 17:03:04 UTC
+++ cli/commands/list/list.go
@@ -8,8 +8,6 @@ import (
 	"sort"
 	"strings"
 
-	"github.com/gruntwork-io/terragrunt/telemetry"
-
 	"github.com/charmbracelet/lipgloss"
 	"github.com/charmbracelet/lipgloss/tree"
 	"github.com/charmbracelet/x/term"
@@ -43,15 +41,10 @@ func Run(ctx context.Context, l log.Logger, opts *Opti
 	)
 
 	// Wrap discovery with telemetry
-	err = telemetry.TelemeterFromContext(ctx).Collect(ctx, "list_discover", map[string]any{
-		"working_dir":  opts.WorkingDir,
-		"hidden":       opts.Hidden,
-		"dependencies": shouldDiscoverDependencies(opts),
-		"external":     opts.External,
-	}, func(ctx context.Context) error {
+	err = func(ctx context.Context) error {
 		components, discoverErr = d.Discover(ctx, l, opts.TerragruntOptions)
 		return discoverErr
-	})
+	}(ctx)
 	if err != nil {
 		l.Debugf("Errors encountered while discovering components:\n%s", err)
 	}
@@ -60,10 +53,7 @@ func Run(ctx context.Context, l log.Logger, opts *Opti
 	case ModeNormal:
 		components = components.Sort()
 	case ModeDAG:
-		err = telemetry.TelemeterFromContext(ctx).Collect(ctx, "list_mode_dag", map[string]any{
-			"working_dir":  opts.WorkingDir,
-			"config_count": len(components),
-		}, func(ctx context.Context) error {
+		err = func(ctx context.Context) error {
 			q, queueErr := queue.NewQueue(components)
 			if queueErr != nil {
 				return queueErr
@@ -72,7 +62,7 @@ func Run(ctx context.Context, l log.Logger, opts *Opti
 			components = q.Components()
 
 			return nil
-		})
+		}(ctx)
 		if err != nil {
 			return errors.New(err)
 		}
@@ -84,16 +74,13 @@ func Run(ctx context.Context, l log.Logger, opts *Opti
 
 	var listedComponents ListedComponents
 
-	err = telemetry.TelemeterFromContext(ctx).Collect(ctx, "list_discovered_to_listed", map[string]any{
-		"working_dir":  opts.WorkingDir,
-		"config_count": len(components),
-	}, func(ctx context.Context) error {
+	err = func(ctx context.Context) error {
 		var convErr error
 
 		listedComponents, convErr = discoveredToListed(components, opts)
 
 		return convErr
-	})
+	}(ctx)
 	if err != nil {
 		return errors.New(err)
 	}
