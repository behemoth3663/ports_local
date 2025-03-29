--- cli/flags/global/flags.go.orig	2025-03-28 15:11:40 UTC
+++ cli/flags/global/flags.go
@@ -48,15 +48,6 @@ const (
 	HelpFlagName    = "help"
 	VersionFlagName = "version"
 
-	// Telemetry flags.
-
-	TelemetryTraceExporterFlagName                  = "telemetry-trace-exporter"
-	TelemetryTraceExporterInsecureEndpointFlagName  = "telemetry-trace-exporter-insecure-endpoint"
-	TelemetryTraceExporterHTTPEndpointFlagName      = "telemetry-trace-exporter-http-endpoint"
-	TraceparentFlagName                             = "traceparent"
-	TelemetryMetricExporterFlagName                 = "telemetry-metric-exporter"
-	TelemetryMetricExporterInsecureEndpointFlagName = "telemetry-metric-exporter-insecure-endpoint"
-
 	// Renamed flags.
 
 	DeprecatedLogLevelFlagName        = "log-level"
@@ -279,59 +270,10 @@ func NewFlags(opts *options.TerragruntOptions, prefix 
 			flags.WithDeprecatedNames(terragruntPrefix.FlagNames(DeprecatedStrictControlFlagName), terragruntPrefixControl)),
 	}
 
-	flags = flags.Add(NewTelemetryFlags(opts, nil)...)
 	flags = flags.Sort()
 	flags = flags.Add(NewHelpVersionFlags(opts)...)
 
 	return flags
-}
-
-// NewTelemetryFlags creates telemetry related flags.
-func NewTelemetryFlags(opts *options.TerragruntOptions, prefix flags.Prefix) cli.Flags {
-	tgPrefix := prefix.Prepend(flags.TgPrefix)
-	terragruntPrefix := prefix.Prepend(flags.TerragruntPrefix)
-	terragruntPrefixControl := flags.StrictControlsByGlobalFlags(opts.StrictControls)
-
-	return cli.Flags{
-		flags.NewFlag(&cli.GenericFlag[string]{
-			EnvVars:     tgPrefix.EnvVars(TelemetryTraceExporterFlagName),
-			Destination: &opts.Telemetry.TraceExporter,
-		},
-			flags.WithDeprecatedNames(terragruntPrefix.FlagNames("telemetry-trace-exporter"), terragruntPrefixControl),
-			flags.WithDeprecatedNames(terragruntPrefix.FlagNames("telemerty-trace-exporter"), terragruntPrefixControl)),
-
-		flags.NewFlag(&cli.BoolFlag{
-			EnvVars:     tgPrefix.EnvVars(TelemetryTraceExporterInsecureEndpointFlagName),
-			Destination: &opts.Telemetry.TraceExporterInsecureEndpoint,
-		},
-			flags.WithDeprecatedNames(terragruntPrefix.FlagNames("telemetry-trace-exporter-insecure-endpoint"), terragruntPrefixControl),
-			flags.WithDeprecatedNames(terragruntPrefix.FlagNames("telemerty-trace-exporter-insecure-endpoint"), terragruntPrefixControl)),
-
-		flags.NewFlag(&cli.GenericFlag[string]{
-			EnvVars:     tgPrefix.EnvVars(TelemetryTraceExporterHTTPEndpointFlagName),
-			Destination: &opts.Telemetry.TraceExporterHTTPEndpoint,
-		},
-			flags.WithDeprecatedNames(terragruntPrefix.FlagNames("telemetry-trace-exporter-http-endpoint"), terragruntPrefixControl),
-			flags.WithDeprecatedNames(terragruntPrefix.FlagNames("telemerty-trace-exporter-http-endpoint"), terragruntPrefixControl)),
-
-		flags.NewFlag(&cli.GenericFlag[string]{
-			EnvVars:     flags.Prefix{}.EnvVars(TraceparentFlagName),
-			Destination: &opts.Telemetry.TraceParent,
-		}),
-		flags.NewFlag(&cli.GenericFlag[string]{
-			EnvVars:     tgPrefix.EnvVars(TelemetryMetricExporterFlagName),
-			Destination: &opts.Telemetry.MetricExporter,
-		},
-			flags.WithDeprecatedNames(terragruntPrefix.FlagNames("telemetry-metric-exporter"), terragruntPrefixControl),
-			flags.WithDeprecatedNames(terragruntPrefix.FlagNames("telemerty-metric-exporter"), terragruntPrefixControl)),
-
-		flags.NewFlag(&cli.BoolFlag{
-			EnvVars:     tgPrefix.EnvVars(TelemetryMetricExporterInsecureEndpointFlagName),
-			Destination: &opts.Telemetry.MetricExporterInsecureEndpoint,
-		},
-			flags.WithDeprecatedNames(terragruntPrefix.FlagNames("telemetry-metric-exporter-insecure-endpoint"), terragruntPrefixControl),
-			flags.WithDeprecatedNames(terragruntPrefix.FlagNames("telemerty-metric-exporter-insecure-endpoint"), terragruntPrefixControl)),
-	}
 }
 
 func NewLogLevelFlag(opts *options.TerragruntOptions, prefix flags.Prefix) *flags.Flag {
