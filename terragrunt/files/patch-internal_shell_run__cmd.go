--- internal/shell/run_cmd.go.orig	1979-11-29 21:00:00 UTC
+++ internal/shell/run_cmd.go
@@ -4,9 +4,7 @@ import (
 import (
 	"context"
 	"errors"
-	"fmt"
 	"io"
-	"path/filepath"
 	"strings"
 	"time"
 
@@ -17,11 +15,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/writer"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
-
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
-
 	"github.com/gruntwork-io/terragrunt/internal/util"
 )
 
@@ -45,7 +38,6 @@ type ShellOptions struct {
 type ShellOptions struct {
 	EngineOptions *engine.EngineOptions
 	EngineConfig  *engine.EngineConfig
-	Telemetry     *telemetry.Options
 
 	RootWorkingDir         string
 	WorkingDir             string
@@ -59,21 +51,11 @@ type ShellOptions struct {
 	LogDisableErrorSummary bool
 }
 
-// NewShellOptions creates ShellOptions with sensible defaults. Telemetry is
-// always non-nil; TRACEPARENT is read from the environment when set. Use the
-// With* methods to override any field.
+// NewShellOptions creates ShellOptions with sensible defaults.
 func NewShellOptions(env map[string]string) *ShellOptions {
 	venv.RequireEnvMap(env)
 
-	opts := &ShellOptions{
-		Telemetry: &telemetry.Options{},
-	}
-
-	if tp := env[telemetry.TraceParentEnv]; tp != "" {
-		opts.Telemetry.TraceParent = tp
-	}
-
-	return opts
+	return &ShellOptions{}
 }
 
 // WithWorkingDir sets the working directory for command execution.
@@ -90,26 +72,6 @@ func (o *ShellOptions) WithUnitDir(dir string) *ShellO
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
-	if t != nil {
-		o.Telemetry = t
-	}
-
-	return o
-}
-
 // WithEngine sets the engine configuration and options.
 func (o *ShellOptions) WithEngine(
 	cfg *engine.EngineConfig,
@@ -222,42 +184,15 @@ func RunCommandWithOutput(
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
+	err := runCommand(ctx, l, v, runOpts, RunCommandOptions{
+		CommandDir:     commandDir,
+		SuppressStdout: suppressStdout,
+		NeedsPTY:       needsPTY,
+		Command:        command,
+		Args:           args,
+		Output:         &output,
+	})
 
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
-
 	return &output, err
 }
 
@@ -275,8 +210,7 @@ type RunCommandOptions struct {
 // runCommand contains the actual subprocess execution logic, separated to keep
 // RunCommandWithOutput focused on telemetry framing.
 //
-// Requires v.Env. The traceparent goes into a copy of it, so the caller's map
-// is never written.
+// Requires v.Env.
 func runCommand(
 	ctx context.Context,
 	l log.Logger,
@@ -292,18 +226,6 @@ func runCommand(
 		cmdStderr = io.MultiWriter(v.Writers.ErrWriter, &cmdOpts.Output.Stderr)
 		cmdStdout = io.MultiWriter(v.Writers.Writer, &cmdOpts.Output.Stdout)
 	)
-
-	// Pass the traceparent to the child process if it is available in the context.
-	if traceParent := telemetry.TraceParentFromContext(ctx, runOpts.Telemetry); traceParent != "" {
-		l.Debugf(
-			"Setting trace parent=%q for command %s",
-			traceParent,
-			fmt.Sprintf("%s %v", cmdOpts.Command, cmdOpts.Args),
-		)
-
-		v = v.WithEnvCloned()
-		v.Env[telemetry.TraceParentEnv] = traceParent
-	}
 
 	if cmdOpts.SuppressStdout {
 		l.Debugf("Command output will be suppressed.")
