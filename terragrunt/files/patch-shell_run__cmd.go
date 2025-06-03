--- shell/run_cmd.go.orig	2025-06-02 18:34:14 UTC
+++ shell/run_cmd.go
@@ -3,7 +3,6 @@ import (
 
 import (
 	"context"
-	"fmt"
 	"io"
 	"strings"
 	"time"
@@ -12,8 +11,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/os/exec"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 
-	"github.com/gruntwork-io/terragrunt/telemetry"
-
 	"github.com/gruntwork-io/terragrunt/internal/errors"
 	"github.com/gruntwork-io/terragrunt/options"
 	"github.com/gruntwork-io/terragrunt/util"
@@ -59,11 +56,7 @@ func RunCommandWithOutput(
 		commandDir = opts.WorkingDir
 	}
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, "run_"+command, map[string]any{
-		"command": command,
-		"args":    fmt.Sprintf("%v", args),
-		"dir":     commandDir,
-	}, func(ctx context.Context) error {
+	err := func(ctx context.Context) error {
 		l.Debugf("Running command: %s %s", command, strings.Join(args, " "))
 
 		var (
@@ -71,14 +64,6 @@ func RunCommandWithOutput(
 			cmdStdout = io.MultiWriter(opts.Writer, &output.Stdout)
 		)
 
-		// Pass the traceparent to the child process if it is available in the context.
-		traceParent := telemetry.TraceParentFromContext(ctx, opts.Telemetry)
-
-		if traceParent != "" {
-			l.Debugf("Setting trace parent=%q for command %s", traceParent, fmt.Sprintf("%s %v", command, args))
-			opts.Env[telemetry.TraceParentEnv] = traceParent
-		}
-
 		if suppressStdout {
 			l.Debugf("Command output will be suppressed.")
 
@@ -152,7 +137,7 @@ func RunCommandWithOutput(
 		}
 
 		return nil
-	})
+	}
 
-	return &output, err
+	return &output, err(ctx)
 }
