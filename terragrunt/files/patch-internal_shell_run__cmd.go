diff --git a/internal/shell/run_cmd.go b/internal/shell/run_cmd.go
index a66a283dd..b60c82b41 100644
--- internal/shell/run_cmd.go.orig
+++ internal/shell/run_cmd.go
@@ -3,24 +3,18 @@ package shell
 
 import (
 	"context"
-	"fmt"
 	"io"
-	"path/filepath"
 	"strings"
 	"time"
 
 	"github.com/gruntwork-io/terragrunt/internal/engine"
 	"github.com/gruntwork-io/terragrunt/internal/experiment"
 	"github.com/gruntwork-io/terragrunt/internal/os/exec"
+	telemetryopts "github.com/gruntwork-io/terragrunt/internal/traceopts"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/writer"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
-
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
-
 	"github.com/gruntwork-io/terragrunt/internal/util"
 )
 
@@ -40,7 +34,7 @@ const SignalForwardingDelay = time.Second * 15
 type ShellOptions struct {
 	EngineOptions *engine.EngineOptions
 	EngineConfig  *engine.EngineConfig
-	Telemetry     *telemetry.Options
+	Telemetry     *telemetryopts.Options
 
 	RootWorkingDir         string
 	WorkingDir             string
@@ -59,16 +53,7 @@ type ShellOptions struct {
 // With* methods to override any field.
 func NewShellOptions(env map[string]string) *ShellOptions {
 	venv.RequireEnvMap(env)
-
-	opts := &ShellOptions{
-		Telemetry: &telemetry.Options{},
-	}
-
-	if tp := env[telemetry.TraceParentEnv]; tp != "" {
-		opts.Telemetry.TraceParent = tp
-	}
-
-	return opts
+	return &ShellOptions{Telemetry: new(telemetryopts.Options)}
 }
 
 // WithWorkingDir sets the working directory for command execution.
@@ -85,19 +70,8 @@ func (o *ShellOptions) WithUnitDir(dir string) *ShellOptions {
 	return o
 }
 
-// SetTraceParent explicitly overrides the TRACEPARENT value used for trace context propagation.
-func (o *ShellOptions) SetTraceParent(tp string) *ShellOptions {
-	if o.Telemetry == nil {
-		o.Telemetry = &telemetry.Options{}
-	}
-
-	o.Telemetry.TraceParent = tp
-
-	return o
-}
-
-// WithTelemetry sets the full telemetry options, replacing the defaults from the constructor.
-func (o *ShellOptions) WithTelemetry(t *telemetry.Options) *ShellOptions {
+// WithEngine sets the engine configuration and options.
+func (o *ShellOptions) WithTelemetry(t *telemetryopts.Options) *ShellOptions {
 	if t != nil {
 		o.Telemetry = t
 	}
@@ -205,41 +179,14 @@ func RunCommandWithOutput(
 		commandDir = runOpts.WorkingDir
 	}
 
-	err := telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "run_"+filepath.Base(command), map[string]any{
-			"binary":      filepath.Base(command),
-			"binary_path": command,
-			"args":        fmt.Sprintf("%v", args),
-			"dir":         commandDir,
-		}, func(ctx context.Context, l log.Logger) error {
-			runErr := runCommand(ctx, l, v, runOpts, RunCommandOptions{
-				CommandDir:     commandDir,
-				SuppressStdout: suppressStdout,
-				NeedsPTY:       needsPTY,
-				Command:        command,
-				Args:           args,
-				Output:         &output,
-			})
-
-			if span := trace.SpanFromContext(ctx); span.IsRecording() {
-				exitCode := 0
-
-				if runErr != nil {
-					exitCode = -1
-					if code, codeErr := util.GetExitCode(runErr); codeErr == nil {
-						exitCode = code
-					}
-				}
-
-				span.SetAttributes(
-					attribute.Int("exit_code", exitCode),
-					attribute.Int("stdout_bytes", output.Stdout.Len()),
-					attribute.Int("stderr_bytes", output.Stderr.Len()),
-				)
-			}
-
-			return runErr
-		})
+	err := runCommand(ctx, l, v, runOpts, RunCommandOptions{
+		CommandDir:     commandDir,
+		SuppressStdout: suppressStdout,
+		NeedsPTY:       needsPTY,
+		Command:        command,
+		Args:           args,
+		Output:         &output,
+	})
 
 	return &output, err
 }
@@ -255,10 +202,7 @@ type RunCommandOptions struct {
 	NeedsPTY       bool
 }
 
-// runCommand contains the actual subprocess execution logic, separated to keep
-// RunCommandWithOutput focused on telemetry framing.
-//
-// Requires v.Env: the traceparent is written into it before the child forks.
+// runCommand contains the actual subprocess execution logic.
 func runCommand(
 	ctx context.Context,
 	l log.Logger,
@@ -275,16 +219,6 @@ func runCommand(
 		cmdStdout = io.MultiWriter(v.Writers.Writer, &cmdOpts.Output.Stdout)
 	)
 
-	// Pass the traceparent to the child process if it is available in the context.
-	if traceParent := telemetry.TraceParentFromContext(ctx, runOpts.Telemetry); traceParent != "" {
-		l.Debugf(
-			"Setting trace parent=%q for command %s",
-			traceParent,
-			fmt.Sprintf("%s %v", cmdOpts.Command, cmdOpts.Args),
-		)
-		v.Env[telemetry.TraceParentEnv] = traceParent
-	}
-
 	if cmdOpts.SuppressStdout {
 		l.Debugf("Command output will be suppressed.")
 
