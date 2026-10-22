--- internal/configbridge/bridge.go.orig	1979-11-29 21:00:00 UTC
+++ internal/configbridge/bridge.go
@@ -79,7 +79,9 @@ func populateFromOpts(pctx *config.ParsingContext, opt
 	pctx.Headless = opts.Headless
 	pctx.BackendBootstrap = opts.BackendBootstrap
 	pctx.CheckDependentUnits = opts.CheckDependentUnits
-	pctx.Telemetry = opts.Telemetry
+	if opts.Telemetry != nil {
+		pctx.TraceParent = opts.Telemetry.TraceParent
+	}
 	pctx.NoStackValidate = opts.NoStackValidate
 	pctx.NoCAS = opts.NoCAS
 	pctx.CASCloneDepth = opts.CASCloneDepth
@@ -95,13 +97,15 @@ func ShellRunOptsFromOpts(env map[string]string, opts 
 func ShellRunOptsFromOpts(env map[string]string, opts *options.TerragruntOptions) *shell.ShellOptions {
 	s := shell.NewShellOptions(env).
 		WithWorkingDir(opts.WorkingDir).
-		WithTelemetry(opts.Telemetry).
 		WithEngine(opts.EngineConfig, opts.EngineOptions).
 		WithTFPath(opts.TFPath).
 		WithRootWorkingDir(opts.RootWorkingDir).
 		WithExperiments(opts.Experiments).
 		WithHeadless(opts.Headless).
 		WithForwardTFStdout(opts.ForwardTFStdout)
+	if opts.Telemetry != nil && opts.Telemetry.TraceParent != "" {
+		s.SetTraceParent(opts.Telemetry.TraceParent)
+	}
 	s.LogShowAbsPaths = opts.LogShowAbsPaths
 	s.LogDisableErrorSummary = opts.LogDisableErrorSummary
 
@@ -176,7 +180,9 @@ func NewRunOptions(opts *options.TerragruntOptions) *r
 	runOpts.AutoInit = opts.AutoInit
 	runOpts.AutoRetry = opts.AutoRetry
 	runOpts.BackendBootstrap = opts.BackendBootstrap
-	runOpts.Telemetry = opts.Telemetry
+	if opts.Telemetry != nil {
+		runOpts.TraceParent = opts.Telemetry.TraceParent
+	}
 	runOpts.AuthProviderCmd = opts.AuthProviderCmd
 	runOpts.MaxFoldersToCheck = opts.MaxFoldersToCheck
 	runOpts.FailIfBucketCreationRequired = opts.FailIfBucketCreationRequired
