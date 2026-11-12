--- internal/engine/engine.go.orig	1979-11-29 21:00:00 UTC
+++ internal/engine/engine.go
@@ -21,7 +21,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/experiment"
 	"github.com/gruntwork-io/terragrunt/internal/github"
 	"github.com/gruntwork-io/terragrunt/internal/os/signal"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/vexec"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
@@ -235,17 +234,14 @@ func Run(
 
 	var output *util.CmdOutput
 
-	runErr := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "engine_run", map[string]any{
-		"command":            execOptions.Command,
-		"cache_dir":          cacheDir,
-		"engine_initialized": found,
-	}, func(runCtx context.Context, l log.Logger) error {
+	_ = found
+	runErr := func(runCtx context.Context, l log.Logger) error {
 		var invokeErr error
 
 		output, invokeErr = invoke(runCtx, l, v, execOptions, instance.engineClient)
 
 		return invokeErr
-	})
+	}(ctx, l)
 	if runErr != nil {
 		return nil, runErr
 	}
@@ -313,10 +309,7 @@ func downloadEngine(
 		return nil
 	}
 
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "engine_download", map[string]any{
-		"source":  e.Source,
-		"version": e.Version,
-	}, func(ctx context.Context, l log.Logger) error {
+	return func(ctx context.Context, l log.Logger) error {
 		// If source is empty, we cannot download the engine
 		// This indicates an engine block was configured but source was not provided
 		if e.Source == "" {
@@ -414,7 +407,7 @@ func downloadEngine(
 		l.Infof("Engine available as %s", path)
 
 		return nil
-	})
+	}(ctx, l)
 }
 
 func lastReleaseVersion(
@@ -820,11 +813,7 @@ func createEngine(
 		pluginClient *plugin.Client
 	)
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "engine_create", map[string]any{
-		"source":    execOptions.EngineConfig.Source,
-		"version":   execOptions.EngineConfig.Version,
-		"cache_dir": execOptions.CacheDir,
-	}, func(ctx context.Context, l log.Logger) error {
+	err := func(ctx context.Context, l log.Logger) error {
 		path, err := engineDir(v, execOptions)
 		if err != nil {
 			return err
@@ -927,7 +916,7 @@ func createEngine(
 		pluginClient = client
 
 		return nil
-	})
+	}(ctx, l)
 	if err != nil {
 		return nil, nil, err
 	}
@@ -945,10 +934,7 @@ func invoke(
 ) (*util.CmdOutput, error) {
 	var result *util.CmdOutput
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "engine_invoke", map[string]any{
-		"command":   runOptions.Command,
-		"cache_dir": runOptions.CacheDir,
-	}, func(ctx context.Context, l log.Logger) error {
+	err := func(ctx context.Context, l log.Logger) error {
 		l = l.WithField(placeholders.TFPathKeyName, "engine")
 
 		meta, err := ConvertMetaToProtobuf(runOptions.EngineConfig.Meta)
@@ -1079,7 +1065,7 @@ func invoke(
 		result = &output
 
 		return nil
-	})
+	}(ctx, l)
 	if err != nil {
 		return nil, err
 	}
@@ -1126,9 +1112,7 @@ func initialize(
 	runOptions *ExecutionOptions,
 	client *proto.EngineClient,
 ) error {
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "engine_initialize", map[string]any{
-		"cache_dir": runOptions.CacheDir,
-	}, func(ctx context.Context, l log.Logger) error {
+	return func(ctx context.Context, l log.Logger) error {
 		meta, err := ConvertMetaToProtobuf(runOptions.EngineConfig.Meta)
 		if err != nil {
 			return err
@@ -1192,7 +1176,7 @@ func initialize(
 
 			return outputLine, nil
 		})
-	})
+	}(ctx, l)
 }
 
 // ErrEngineShutdownFailed is wrapped by the error returned when the engine's Shutdown call exits with a non-zero code.
@@ -1206,9 +1190,7 @@ func shutdown(
 	runOptions *ExecutionOptions,
 	terragruntEngine *proto.EngineClient,
 ) error {
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "engine_shutdown", map[string]any{
-		"cache_dir": runOptions.CacheDir,
-	}, func(ctx context.Context, l log.Logger) error {
+	return func(ctx context.Context, l log.Logger) error {
 		meta, err := ConvertMetaToProtobuf(runOptions.EngineConfig.Meta)
 		if err != nil {
 			return err
@@ -1275,7 +1257,7 @@ func shutdown(
 
 			return outputLine, nil
 		})
-	})
+	}(ctx, l)
 }
 
 // OutputLine represents the output from the engine
