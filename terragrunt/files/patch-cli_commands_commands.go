--- cli/commands/commands.go.orig	2025-11-05 17:03:04 UTC
+++ cli/commands/commands.go
@@ -3,7 +3,6 @@ import (
 
 import (
 	"context"
-	"fmt"
 	"os"
 	"path/filepath"
 	"slices"
@@ -37,7 +36,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/errors"
 	"github.com/gruntwork-io/terragrunt/internal/os/exec"
 	"github.com/gruntwork-io/terragrunt/pkg/log/format/placeholders"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/hashicorp/go-version"
 )
 
@@ -124,11 +122,7 @@ func WrapWithTelemetry(l log.Logger, opts *options.Ter
 // WrapWithTelemetry wraps CLI command execution with setting of telemetry context and labels, if telemetry is disabled, just runAction the command.
 func WrapWithTelemetry(l log.Logger, opts *options.TerragruntOptions) func(ctx *cli.Context, action cli.ActionFunc) error {
 	return func(ctx *cli.Context, action cli.ActionFunc) error {
-		return telemetry.TelemeterFromContext(ctx).Collect(ctx.Context, fmt.Sprintf("%s %s", ctx.Command.Name, opts.TerraformCommand), map[string]any{
-			"terraformCommand": opts.TerraformCommand,
-			"args":             opts.TerraformCliArgs,
-			"dir":              opts.WorkingDir,
-		}, func(childCtx context.Context) error {
+		return func(childCtx context.Context) error {
 			ctx.Context = childCtx //nolint:fatcontext
 			if err := initialSetup(ctx, l, opts); err != nil {
 				return err
@@ -136,7 +130,7 @@ func WrapWithTelemetry(l log.Logger, opts *options.Ter
 
 			// TODO: See if this lint should be ignored
 			return runAction(ctx, l, opts, action) //nolint:contextcheck
-		})
+		}(ctx)
 	}
 }
 
