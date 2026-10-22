--- internal/shell/run_cmd.go.orig	1979-11-29 21:00:00 UTC
+++ internal/shell/run_cmd.go
@@ -6,7 +6,6 @@ import (
 	"errors"
 	"fmt"
 	"io"
-	"path/filepath"
 	"strings"
 	"time"
 
@@ -17,11 +16,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/writer"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
-
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
-
 	"github.com/gruntwork-io/terragrunt/internal/util"
 )
 
@@ -45,7 +39,7 @@ type ShellOptions struct {
 type ShellOptions struct {
 	EngineOptions *engine.EngineOptions
 	EngineConfig  *engine.EngineConfig
-	Telemetry     *telemetry.Options
+	TraceParent   string
 
 	RootWorkingDir         string
 	WorkingDir             string
@@ -59,20 +53,16 @@ type ShellOptions struct {
 	LogDisableErrorSummary bool
 }
 
-// NewShellOptions creates ShellOptions with sensible defaults. Telemetry is
-// always non-nil; TRACEPARENT is read from the environment when set. Use the
-// With* methods to override any field.
+// NewShellOptions creates ShellOptions with sensible defaults.
+// TRACEPARENT is read from the environment when set.
 func NewShellOptions(env map[string]string) *ShellOptions {
 	venv.RequireEnvMap(env)
 
-	opts := &ShellOptions{
-		Telemetry: &telemetry.Options{},
+	opts := &ShellOptions{}
+	if tp := env["TRACEPARENT"]; tp != "" {
+		opts.TraceParent = tp
 	}
 
-	if tp := env[telemetry.TraceParentEnv]; tp != "" {
-		opts.Telemetry.TraceParent = tp
-	}
-
 	return opts
 }
 
@@ -92,24 +82,11 @@ func (o *ShellOptions) SetTraceParent(tp string) *Shel
 
 // SetTraceParent explicitly overrides the TRACEPARENT value used for trace context propagation.
 func (o *ShellOptions) SetTraceParent(tp string) *ShellOptions {
-	if o.Telemetry == nil {
-		o.Telemetry = &telemetry.Options{}
-	}
+	o.TraceParent = tp
 
-	o.Telemetry.TraceParent = tp
-
 	return o
 }
 
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
@@ -222,42 +199,16 @@ func RunCommandWithOutput(
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
+	_ = fmt.Sprintf("%v", args)
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
 
@@ -294,7 +245,7 @@ func runCommand(
 	)
 
 	// Pass the traceparent to the child process if it is available in the context.
-	if traceParent := telemetry.TraceParentFromContext(ctx, runOpts.Telemetry); traceParent != "" {
+	if traceParent := runOpts.TraceParent; traceParent != "" {
 		l.Debugf(
 			"Setting trace parent=%q for command %s",
 			traceParent,
@@ -302,7 +253,7 @@ func runCommand(
 		)
 
 		v = v.WithEnvCloned()
-		v.Env[telemetry.TraceParentEnv] = traceParent
+		v.Env["TRACEPARENT"] = traceParent
 	}
 
 	if cmdOpts.SuppressStdout {
