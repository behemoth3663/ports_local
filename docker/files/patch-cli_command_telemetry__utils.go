--- cli/command/telemetry_utils.go.orig	2026-09-23 09:45:10 UTC
+++ cli/command/telemetry_utils.go
@@ -5,33 +5,29 @@ import (
 	"errors"
 	"fmt"
 	"strings"
-	"time"
 
-	"github.com/docker/cli/cli/version"
 	"github.com/spf13/cobra"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/metric"
 )
 
-// BaseCommandAttributes returns an attribute.Set containing attributes to attach to metrics/traces
-func BaseCommandAttributes(cmd *cobra.Command, streams Streams) []attribute.KeyValue {
-	return append([]attribute.KeyValue{
-		attribute.String("command.name", getCommandName(cmd)),
-	}, stdioAttributes(streams)...)
+// BaseCommandAttributes returns command metadata used by telemetry integrations.
+func BaseCommandAttributes(cmd *cobra.Command, streams Streams) map[string]string {
+	attrs := map[string]string{
+		"command.name":         getCommandName(cmd),
+		"command.stdin.isatty": fmt.Sprintf("%t", streams.In().IsTerminal()),
+		"command.stdout.isatty": fmt.Sprintf("%t",
+			streams.Out().IsTerminal()),
+		"command.stderr.isatty": fmt.Sprintf("%t",
+			streams.Err().IsTerminal()),
+	}
+	return attrs
 }
 
-// InstrumentCobraCommands wraps all cobra commands' RunE funcs to set a command duration metric using otel.
-//
-// Note: this should be the last func to wrap/modify the PersistentRunE/RunE funcs before command execution.
-//
-// can also be used for spans!
+// InstrumentCobraCommands wraps all cobra commands' RunE funcs.
 func (cli *DockerCli) InstrumentCobraCommands(ctx context.Context, cmd *cobra.Command) {
-	// If PersistentPreRunE is nil, make it execute PersistentPreRun and return nil by default
+	_ = ctx
 	ogPersistentPreRunE := cmd.PersistentPreRunE
 	if ogPersistentPreRunE == nil {
 		ogPersistentPreRun := cmd.PersistentPreRun
-		//nolint:unparam // necessary because error will always be nil here
 		ogPersistentPreRunE = func(cmd *cobra.Command, args []string) error {
 			ogPersistentPreRun(cmd, args)
 			return nil
@@ -39,13 +35,10 @@ func (cli *DockerCli) InstrumentCobraCommands(ctx cont
 		cmd.PersistentPreRun = nil
 	}
 
-	// wrap RunE in PersistentPreRunE so that this operation gets executed on all children commands
 	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
-		// If RunE is nil, make it execute Run and return nil by default
 		ogRunE := cmd.RunE
 		if ogRunE == nil {
 			ogRun := cmd.Run
-			//nolint:unparam // necessary because error will always be nil here
 			ogRunE = func(cmd *cobra.Command, args []string) error {
 				ogRun(cmd, args)
 				return nil
@@ -53,83 +46,35 @@ func (cli *DockerCli) InstrumentCobraCommands(ctx cont
 			cmd.Run = nil
 		}
 		cmd.RunE = func(cmd *cobra.Command, args []string) error {
-			// start the timer as the first step of every cobra command
 			stopInstrumentation := cli.StartInstrumentation(cmd)
 			cmdErr := ogRunE(cmd, args)
 			stopInstrumentation(cmdErr)
 			return cmdErr
 		}
-
 		return ogPersistentPreRunE(cmd, args)
 	}
 }
 
-// StartInstrumentation instruments CLI commands with the individual metrics and spans configured.
-// It's the main command OTel utility, and new command-related metrics should be added to it.
-// It should be called immediately before command execution, and returns a stopInstrumentation function
-// that must be called with the error resulting from the command execution.
+// StartInstrumentation returns a no-op stopper while preserving the call-flow.
 func (cli *DockerCli) StartInstrumentation(cmd *cobra.Command) (stopInstrumentation func(error)) {
-	baseAttrs := BaseCommandAttributes(cmd, cli)
-	return startCobraCommandTimer(cli.MeterProvider(), baseAttrs)
+	_ = BaseCommandAttributes(cmd, cli)
+	return func(error) {}
 }
 
-func startCobraCommandTimer(mp metric.MeterProvider, attrs []attribute.KeyValue) func(err error) {
-	meter := getDefaultMeter(mp)
-	durationCounter, _ := meter.Float64Counter(
-		"command.time",
-		metric.WithDescription("Measures the duration of the cobra command"),
-		metric.WithUnit("ms"),
-	)
-	start := time.Now()
-
-	return func(err error) {
-		// Use a new context for the export so that the command being cancelled
-		// doesn't affect the metrics, and we get metrics for cancelled commands.
-		ctx, cancel := context.WithTimeout(context.Background(), exportTimeout)
-		defer cancel()
-
-		duration := float64(time.Since(start)) / float64(time.Millisecond)
-		cmdStatusAttrs := attributesFromError(err)
-		durationCounter.Add(ctx, duration,
-			metric.WithAttributes(attrs...),
-			metric.WithAttributes(cmdStatusAttrs...),
-		)
-		if mp, ok := mp.(MeterProvider); ok {
-			if err := mp.ForceFlush(ctx); err != nil {
-				otel.Handle(err)
-			}
-		}
-	}
-}
-
-func stdioAttributes(streams Streams) []attribute.KeyValue {
-	return []attribute.KeyValue{
-		attribute.Bool("command.stdin.isatty", streams.In().IsTerminal()),
-		attribute.Bool("command.stdout.isatty", streams.Out().IsTerminal()),
-		attribute.Bool("command.stderr.isatty", streams.Err().IsTerminal()),
-	}
-}
-
-func attributesFromError(err error) []attribute.KeyValue {
-	attrs := []attribute.KeyValue{}
+func attributesFromError(err error) map[string]string {
+	attrs := map[string]string{}
 	exitCode := 0
 	if err != nil {
 		exitCode = 1
-		if stderr, ok := err.(statusError); ok {
-			// StatusError should only be used for errors, and all errors should
-			// have a non-zero exit status, so only set this here if this value isn't 0
-			if stderr.StatusCode != 0 {
-				exitCode = stderr.StatusCode
-			}
+		if stderr, ok := err.(statusError); ok && stderr.StatusCode != 0 {
+			exitCode = stderr.StatusCode
 		}
-		attrs = append(attrs, attribute.String("command.error.type", otelErrorType(err)))
+		attrs["command.error.type"] = otelErrorType(err)
 	}
-	attrs = append(attrs, attribute.Int("command.status.code", exitCode))
-
+	attrs["command.status.code"] = fmt.Sprintf("%d", exitCode)
 	return attrs
 }
 
-// otelErrorType returns an attribute for the error type based on the error category.
 func otelErrorType(err error) string {
 	name := "generic"
 	if errors.Is(err, context.Canceled) {
@@ -138,7 +83,6 @@ func otelErrorType(err error) string {
 	return name
 }
 
-// statusError reports an unsuccessful exit by a command.
 type statusError struct {
 	Status     string
 	StatusCode int
@@ -148,11 +92,6 @@ func (e statusError) Error() string {
 	return fmt.Sprintf("Status: %s, Code: %d", e.Status, e.StatusCode)
 }
 
-// getCommandName gets the cobra command name in the format
-// `... parentCommandName commandName` by traversing it's parent commands recursively.
-// until the root command is reached.
-//
-// Note: The root command's name is excluded. If cmd is the root cmd, return ""
 func getCommandName(cmd *cobra.Command) string {
 	fullCmdName := getFullCommandName(cmd)
 	_, after, ok := strings.Cut(fullCmdName, " ")
@@ -162,21 +101,9 @@ func getCommandName(cmd *cobra.Command) string {
 	return after
 }
 
-// getFullCommandName gets the full cobra command name in the format
-// `... parentCommandName commandName` by traversing it's parent commands recursively
-// until the root command is reached.
 func getFullCommandName(cmd *cobra.Command) string {
 	if cmd.HasParent() {
 		return fmt.Sprintf("%s %s", getFullCommandName(cmd.Parent()), cmd.Name())
 	}
 	return cmd.Name()
-}
-
-// getDefaultMeter gets the default metric.Meter for the application
-// using the given metric.MeterProvider
-func getDefaultMeter(mp metric.MeterProvider) metric.Meter {
-	return mp.Meter(
-		"github.com/docker/cli",
-		metric.WithInstrumentationVersion(version.Version),
-	)
 }
