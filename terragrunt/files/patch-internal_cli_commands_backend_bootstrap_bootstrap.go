--- internal/cli/commands/backend/bootstrap/bootstrap.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/commands/backend/bootstrap/bootstrap.go
@@ -11,7 +11,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/cli/commands/discoverysetup"
 	rewrite "github.com/gruntwork-io/terragrunt/internal/config/rewrite"
 	"github.com/gruntwork-io/terragrunt/internal/configbridge"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
@@ -35,19 +34,14 @@ func runBootstrap(
 	v *venv.Venv,
 	opts *options.TerragruntOptions,
 ) error {
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "backend_bootstrap", map[string]any{
-		"working_dir":            opts.WorkingDir,
-		"terragrunt_config_path": opts.TerragruntConfigPath,
-	}, func(ctx context.Context, l log.Logger) error {
-		pctx := configbridge.NewParsingContext(opts)
+	pctx := configbridge.NewParsingContext(opts)
 
-		remoteState, err := config.ParseRemoteState(ctx, l, v, pctx)
-		if err != nil || remoteState == nil {
-			return err
-		}
+	remoteState, err := config.ParseRemoteState(ctx, l, v, pctx)
+	if err != nil || remoteState == nil {
+		return err
+	}
 
-		return remoteState.Bootstrap(ctx, l, v, configbridge.RemoteStateOptsFromOpts(v.Env, opts))
-	})
+	return remoteState.Bootstrap(ctx, l, v, configbridge.RemoteStateOptsFromOpts(v.Env, opts))
 }
 
 func runAll(
@@ -65,49 +59,42 @@ func runAll(
 		return err
 	}
 
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "backend_bootstrap_all", map[string]any{
-			"working_dir": opts.WorkingDir,
-			"unit_count":  len(units),
-			"fail_fast":   opts.FailFast,
-		}, func(ctx context.Context, l log.Logger) error {
-			var errs []error
+	var errs []error
 
-			for _, unit := range units {
-				unitOpts := opts.Clone()
-				unitOpts.WorkingDir = unit.Path()
+	for _, unit := range units {
+		unitOpts := opts.Clone()
+		unitOpts.WorkingDir = unit.Path()
 
-				configFilename := config.DefaultTerragruntConfigPath
-				if len(opts.TerragruntConfigPath) > 0 {
-					configFilename = filepath.Base(opts.TerragruntConfigPath)
-				}
+		configFilename := config.DefaultTerragruntConfigPath
+		if len(opts.TerragruntConfigPath) > 0 {
+			configFilename = filepath.Base(opts.TerragruntConfigPath)
+		}
 
-				unitOpts.TerragruntConfigPath = filepath.Join(unit.Path(), configFilename)
-				unitOpts.OriginalTerragruntConfigPath = unitOpts.TerragruntConfigPath
+		unitOpts.TerragruntConfigPath = filepath.Join(unit.Path(), configFilename)
+		unitOpts.OriginalTerragruntConfigPath = unitOpts.TerragruntConfigPath
 
-				// Parsing can write obtained credentials into the env, so each
-				// unit gets its own clone to keep them from leaking to siblings.
-				unitV := v.WithEnvCloned()
-				if err := runBootstrap(ctx, l, unitV, unitOpts); err != nil {
-					if opts.FailFast {
-						return err
-					}
-
-					errs = append(
-						errs,
-						fmt.Errorf(
-							"backend bootstrap for unit %s failed: %w",
-							unit.Path(),
-							err,
-						),
-					)
-				}
+		// Parsing can write obtained credentials into the env, so each
+		// unit gets its own clone to keep them from leaking to siblings.
+		unitV := v.WithEnvCloned()
+		if err := runBootstrap(ctx, l, unitV, unitOpts); err != nil {
+			if opts.FailFast {
+				return err
 			}
 
-			if len(errs) > 0 {
-				return errors.Join(errs...)
-			}
+			errs = append(
+				errs,
+				fmt.Errorf(
+					"backend bootstrap for unit %s failed: %w",
+					unit.Path(),
+					err,
+				),
+			)
+		}
+	}
 
-			return nil
-		})
+	if len(errs) > 0 {
+		return errors.Join(errs...)
+	}
+
+	return nil
 }
