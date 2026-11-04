--- internal/cli/flags/global/flags.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/flags/global/flags.go
@@ -335,109 +335,11 @@ func NewFlags(l log.Logger, opts *options.TerragruntOp
 		),
 	}
 
-	flags = flags.Add(NewTelemetryFlags(opts, nil)...)
 	flags = flags.Add(NewProfileFlags(opts, nil)...)
 	flags = flags.Sort()
 	flags = flags.Add(NewHelpVersionFlags(l, opts)...)
 
 	return flags
-}
-
-// NewTelemetryFlags creates telemetry related flags.
-func NewTelemetryFlags(opts *options.TerragruntOptions, prefix flags.Prefix) clihelper.Flags {
-	tgPrefix := prefix.Prepend(flags.TgPrefix)
-	terragruntPrefix := prefix.Prepend(flags.TerragruntPrefix)
-
-	return clihelper.Flags{
-		flags.NewFlag(
-			&clihelper.GenericFlag[string]{
-				EnvVars:     tgPrefix.EnvVars(TelemetryTraceExporterFlagName),
-				Destination: &opts.Telemetry.TraceExporter,
-			},
-			flags.WithDeprecatedEnvVars(
-				terragruntPrefix.EnvVars("telemetry-trace-exporter"),
-				opts.StrictControls,
-			),
-			flags.WithDeprecatedEnvVars(
-				terragruntPrefix.EnvVars("telemerty-trace-exporter"),
-				opts.StrictControls,
-			),
-		),
-
-		flags.NewFlag(
-			&clihelper.BoolFlag{
-				EnvVars:     tgPrefix.EnvVars(TelemetryTraceExporterInsecureEndpointFlagName),
-				Destination: &opts.Telemetry.TraceExporterInsecureEndpoint,
-			},
-			flags.WithDeprecatedEnvVars(
-				terragruntPrefix.EnvVars("telemetry-trace-exporter-insecure-endpoint"),
-				opts.StrictControls,
-			),
-			flags.WithDeprecatedEnvVars(
-				terragruntPrefix.EnvVars("telemerty-trace-exporter-insecure-endpoint"),
-				opts.StrictControls,
-			),
-		),
-
-		flags.NewFlag(
-			&clihelper.GenericFlag[string]{
-				EnvVars:     tgPrefix.EnvVars(TelemetryTraceExporterHTTPEndpointFlagName),
-				Destination: &opts.Telemetry.TraceExporterHTTPEndpoint,
-			},
-			flags.WithDeprecatedEnvVars(
-				terragruntPrefix.EnvVars("telemetry-trace-exporter-http-endpoint"),
-				opts.StrictControls,
-			),
-			flags.WithDeprecatedEnvVars(
-				terragruntPrefix.EnvVars("telemerty-trace-exporter-http-endpoint"),
-				opts.StrictControls,
-			),
-		),
-
-		flags.NewFlag(&clihelper.GenericFlag[string]{
-			EnvVars:     flags.Prefix{}.EnvVars(TraceparentFlagName),
-			Destination: &opts.Telemetry.TraceParent,
-		}),
-		flags.NewFlag(
-			&clihelper.GenericFlag[string]{
-				EnvVars:     tgPrefix.EnvVars(TelemetryMetricExporterFlagName),
-				Destination: &opts.Telemetry.MetricExporter,
-			},
-			flags.WithDeprecatedEnvVars(
-				terragruntPrefix.EnvVars("telemetry-metric-exporter"),
-				opts.StrictControls,
-			),
-			flags.WithDeprecatedEnvVars(
-				terragruntPrefix.EnvVars("telemerty-metric-exporter"),
-				opts.StrictControls,
-			),
-		),
-
-		flags.NewFlag(
-			&clihelper.BoolFlag{
-				EnvVars:     tgPrefix.EnvVars(TelemetryMetricExporterInsecureEndpointFlagName),
-				Destination: &opts.Telemetry.MetricExporterInsecureEndpoint,
-			},
-			flags.WithDeprecatedEnvVars(
-				terragruntPrefix.EnvVars("telemetry-metric-exporter-insecure-endpoint"),
-				opts.StrictControls,
-			),
-			flags.WithDeprecatedEnvVars(
-				terragruntPrefix.EnvVars("telemerty-metric-exporter-insecure-endpoint"),
-				opts.StrictControls,
-			),
-		),
-
-		flags.NewFlag(&clihelper.GenericFlag[string]{
-			EnvVars:     tgPrefix.EnvVars(TelemetryLogsExporterFlagName),
-			Destination: &opts.Telemetry.LogsExporter,
-		}),
-
-		flags.NewFlag(&clihelper.BoolFlag{
-			EnvVars:     tgPrefix.EnvVars(TelemetryLogsExporterInsecureEndpointFlagName),
-			Destination: &opts.Telemetry.LogsExporterInsecureEndpoint,
-		}),
-	}
 }
 
 // NewProfileFlags creates the profiling flags backing the opts.Profile* fields.
