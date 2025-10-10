--- vendor/github.com/gruntwork-io/terragrunt/shell/run_shell_cmd.go.orig	2025-01-27 15:23:56 UTC
+++ vendor/github.com/gruntwork-io/terragrunt/shell/run_shell_cmd.go
@@ -3,7 +3,6 @@ import (
 
 import (
 	"context"
-	"fmt"
 	"io"
 	"os"
 	"path/filepath"
@@ -20,8 +19,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/pkg/log/writer"
 	"github.com/gruntwork-io/terragrunt/terraform"
 
-	"github.com/gruntwork-io/terragrunt/telemetry"
-
 	"github.com/gruntwork-io/go-commons/collections"
 	"github.com/gruntwork-io/terragrunt/internal/errors"
 	"github.com/gruntwork-io/terragrunt/options"
@@ -121,11 +118,7 @@ func RunShellCommandWithOutput(
 		commandDir = opts.WorkingDir
 	}
 
-	err := telemetry.Telemetry(ctx, opts, "run_"+command, map[string]interface{}{
-		"command": command,
-		"args":    fmt.Sprintf("%v", args),
-		"dir":     commandDir,
-	}, func(childCtx context.Context) error {
+	err := func(ctx context.Context) error {
 		opts.Logger.Debugf("Running command: %s %s", command, strings.Join(args, " "))
 
 		var (
@@ -249,9 +242,9 @@ func RunShellCommandWithOutput(
 		}
 
 		return nil
-	})
+	}
 
-	return &output, err
+	return &output, err(ctx)
 }
 
 // buildOutWriter returns the writer for the command's stdout.
