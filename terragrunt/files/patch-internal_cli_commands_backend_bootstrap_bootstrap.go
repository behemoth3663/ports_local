--- internal/cli/commands/backend/bootstrap/bootstrap.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/commands/backend/bootstrap/bootstrap.go
@@ -9,7 +9,6 @@ import (
 
 	"github.com/gruntwork-io/terragrunt/internal/cli/commands/discoverysetup"
 	"github.com/gruntwork-io/terragrunt/internal/configbridge"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
@@ -34,10 +33,6 @@ func runBootstrap(
 	s *config.Store,
 	opts *options.TerragruntOptions,
 ) error {
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "backend_bootstrap", map[string]any{
-		"working_dir":            opts.WorkingDir,
-		"terragrunt_config_path": opts.TerragruntConfigPath,
-	}, func(ctx context.Context, l log.Logger) error {
 		pctx := configbridge.NewParsingContext(opts)
 
 		remoteState, err := config.ParseRemoteState(ctx, l, v, s, pctx)
@@ -46,7 +41,6 @@ func runBootstrap(
 		}
 
 		return remoteState.Bootstrap(ctx, l, v, configbridge.RemoteStateOptsFromOpts(v.Exec, v.Env, opts))
-	})
 }
 
 func runAll(
@@ -64,12 +58,6 @@ func runAll(
 		return err
 	}
 
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "backend_bootstrap_all", map[string]any{
-			"working_dir": opts.WorkingDir,
-			"unit_count":  len(units),
-			"fail_fast":   opts.FailFast,
-		}, func(ctx context.Context, l log.Logger) error {
 			var errs []error
 
 			for _, unit := range units {
@@ -100,5 +88,4 @@ func runAll(
 			}
 
 			return nil
-		})
 }
