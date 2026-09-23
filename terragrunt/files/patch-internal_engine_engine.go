--- internal/engine/engine.go.orig	1979-11-29 21:00:00 UTC
+++ internal/engine/engine.go
@@ -5,6 +5,7 @@ import (
 	"bytes"
 	"context"
 	"encoding/json"
+	"errors"
 	"fmt"
 	"io"
 	"os"
@@ -21,7 +22,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/experiment"
 	"github.com/gruntwork-io/terragrunt/internal/github"
 	"github.com/gruntwork-io/terragrunt/internal/os/signal"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/vexec"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
@@ -31,8 +31,6 @@ import (
 
 	"google.golang.org/grpc/credentials/insecure"
 
-	"errors"
-
 	"github.com/gruntwork-io/terragrunt-engine-go/engine"
 	"github.com/gruntwork-io/terragrunt-engine-go/proto"
 	"github.com/gruntwork-io/terragrunt/internal/util"
@@ -223,7 +221,7 @@ func Run(
 
 	cacheDir := execOptions.CacheDir
 
-	instance, found, err := engineClients.loadOrCreate(
+	instance, _, err := engineClients.loadOrCreate(
 		cacheDir,
 		execOptions,
 		func() (*engineInstance, error) {
@@ -234,21 +232,9 @@ func Run(
 		return nil, err
 	}
 
-	var output *util.CmdOutput
-
-	runErr := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "engine_run", map[string]any{
-		"command":            execOptions.Command,
-		"cache_dir":          cacheDir,
-		"engine_initialized": found,
-	}, func(runCtx context.Context, l log.Logger) error {
-		var invokeErr error
-
-		output, invokeErr = invoke(runCtx, l, v, execOptions, instance.engineClient)
-
-		return invokeErr
-	})
-	if runErr != nil {
-		return nil, runErr
+	output, err := invoke(ctx, l, v, execOptions, instance.engineClient)
+	if err != nil {
+		return nil, err
 	}
 
 	return output, nil
@@ -314,10 +300,6 @@ func downloadEngine(
 		return nil
 	}
 
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "engine_download", map[string]any{
-		"source":  e.Source,
-		"version": e.Version,
-	}, func(ctx context.Context, l log.Logger) error {
 		// If source is empty, we cannot download the engine
 		// This indicates an engine block was configured but source was not provided
 		if e.Source == "" {
@@ -415,7 +397,6 @@ func downloadEngine(
 		l.Infof("Engine available as %s", path)
 
 		return nil
-	})
 }
 
 func lastReleaseVersion(
@@ -821,14 +802,9 @@ func createEngine(
 		pluginClient *plugin.Client
 	)
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "engine_create", map[string]any{
-		"source":    execOptions.EngineConfig.Source,
-		"version":   execOptions.EngineConfig.Version,
-		"cache_dir": execOptions.CacheDir,
-	}, func(ctx context.Context, l log.Logger) error {
 		path, err := engineDir(v, execOptions)
 		if err != nil {
-			return err
+			return nil, nil, err
 		}
 
 		localEnginePath := filepath.Join(path, engineFileName(v, execOptions.EngineConfig))
@@ -845,7 +821,7 @@ func createEngine(
 				localChecksumFile,
 				localChecksumSigFile,
 			); err != nil {
-				return err
+				return nil, nil, err
 			}
 		} else {
 			l.Warnf("Skipping verification for %s", localEnginePath)
@@ -893,7 +869,7 @@ func createEngine(
 		// hashicorp/go-plugin's ClientConfig requires a concrete *exec.Cmd.
 		osCmder, ok := cmd.(vexec.OSCmder)
 		if !ok {
-			return fmt.Errorf("engine plugin spawn: %w", vexec.ErrNotOSBacked)
+			return nil, nil, fmt.Errorf("engine plugin spawn: %w", vexec.ErrNotOSBacked)
 		}
 
 		client := plugin.NewClient(&plugin.ClientConfig{
@@ -913,28 +889,22 @@ func createEngine(
 
 		rpcClient, err := client.Client()
 		if err != nil {
-			return err
+			return nil, nil, err
 		}
 
 		rawClient, err := rpcClient.Dispense("plugin")
 		if err != nil {
-			return err
+			return nil, nil, err
 		}
 
 		terragruntEngine, ok := rawClient.(proto.EngineClient)
 		if !ok {
-			return fmt.Errorf("engine plugin returned unexpected client type %T", rawClient)
+			return nil, nil, fmt.Errorf("engine plugin returned unexpected client type %T", rawClient)
 		}
 
 		engineClient = &terragruntEngine
 		pluginClient = client
 
-		return nil
-	})
-	if err != nil {
-		return nil, nil, err
-	}
-
 	return engineClient, pluginClient, nil
 }
 
@@ -948,15 +918,11 @@ func invoke(
 ) (*util.CmdOutput, error) {
 	var result *util.CmdOutput
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "engine_invoke", map[string]any{
-		"command":   runOptions.Command,
-		"cache_dir": runOptions.CacheDir,
-	}, func(ctx context.Context, l log.Logger) error {
 		l = l.WithField(placeholders.TFPathKeyName, "engine")
 
 		meta, err := ConvertMetaToProtobuf(runOptions.EngineConfig.Meta)
 		if err != nil {
-			return err
+			return nil, err
 		}
 
 		response, err := (*client).Run(ctx, &proto.RunRequest{
@@ -968,7 +934,7 @@ func invoke(
 			EnvVars:           v.Env,
 		})
 		if err != nil {
-			return err
+			return nil, err
 		}
 
 		// Determine log levels based on headless mode (similar to buildOutWriter/buildErrWriter)
@@ -1028,7 +994,7 @@ func invoke(
 						&stdoutLineBuf,
 						stdout,
 					); err != nil {
-						return err
+						return nil, err
 					}
 				}
 			case *proto.RunResponse_Stderr:
@@ -1038,7 +1004,7 @@ func invoke(
 						&stderrLineBuf,
 						stderr,
 					); err != nil {
-						return err
+						return nil, err
 					}
 				}
 			case *proto.RunResponse_ExitResult:
@@ -1055,11 +1021,11 @@ func invoke(
 		}
 
 		if err = flushBuffer(&stdoutLineBuf, stdout); err != nil {
-			return err
+			return nil, err
 		}
 
 		if err = flushBuffer(&stderrLineBuf, stderr); err != nil {
-			return err
+			return nil, err
 		}
 
 		l.Debugf("Engine execution done in %v", runOptions.CacheDir)
@@ -1076,17 +1042,11 @@ func invoke(
 				DisableSummary:  runOptions.LogDisableErrorSummary,
 			}
 
-			return err
+			return nil, err
 		}
 
 		result = &output
 
-		return nil
-	})
-	if err != nil {
-		return nil, err
-	}
-
 	return result, nil
 }
 
@@ -1128,72 +1088,71 @@ func initialize(
 	runOptions *ExecutionOptions,
 	client *proto.EngineClient,
 ) error {
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "engine_initialize", map[string]any{
-		"cache_dir": runOptions.CacheDir,
-	}, func(ctx context.Context, l log.Logger) error {
-		meta, err := ConvertMetaToProtobuf(runOptions.EngineConfig.Meta)
-		if err != nil {
-			return err
-		}
+	meta, err := ConvertMetaToProtobuf(runOptions.EngineConfig.Meta)
+	if err != nil {
+		return err
+	}
 
-		l.Debugf("Running init for engine in %s", runOptions.CacheDir)
+	l.Debugf("Running init for engine in %s", runOptions.CacheDir)
 
-		request, err := (*client).Init(ctx, &proto.InitRequest{
-			EnvVars:    v.Env,
-			WorkingDir: runOptions.CacheDir,
-			Meta:       meta,
-		})
+	request, err := (*client).Init(ctx, &proto.InitRequest{
+		EnvVars:    v.Env,
+		WorkingDir: runOptions.CacheDir,
+		Meta:       meta,
+	})
+	if err != nil {
+		return err
+	}
+
+	l.Debugf("Reading init output for engine in %s", runOptions.CacheDir)
+
+	return ReadEngineOutput(v, true, func() (*OutputLine, error) {
+		output, err := request.Recv()
 		if err != nil {
-			return err
+			return nil, err
 		}
 
-		l.Debugf("Reading init output for engine in %s", runOptions.CacheDir)
+		if output == nil {
+			return nil, nil
+		}
 
-		return ReadEngineOutput(v, true, func() (*OutputLine, error) {
-			output, err := request.Recv()
-			if err != nil {
-				return nil, err
+		outputLine := &OutputLine{}
+
+		//nolint:dupl // Similar structure to shutdown response handling, but different protobuf types
+		switch resp := output.GetResponse().(type) {
+		case *proto.InitResponse_Stdout:
+			if resp.Stdout != nil {
+				outputLine.Stdout = resp.Stdout.GetContent()
 			}
 
-			if output == nil {
-				return nil, nil
+		case *proto.InitResponse_Stderr:
+			if resp.Stderr != nil {
+				outputLine.Stderr = resp.Stderr.GetContent()
 			}
 
-			outputLine := &OutputLine{}
+		case *proto.InitResponse_ExitResult:
+			if resp.ExitResult != nil {
+				exitCode := int(resp.ExitResult.GetCode())
+				if exitCode != 0 {
+					l.Errorf("Engine init failed with exit code %d", exitCode)
 
-			//nolint:dupl // Similar structure to shutdown response handling, but different protobuf types
-			switch resp := output.GetResponse().(type) {
-			case *proto.InitResponse_Stdout:
-				if resp.Stdout != nil {
-					outputLine.Stdout = resp.Stdout.GetContent()
+					return nil, fmt.Errorf(
+						"%w with exit code %d",
+						ErrEngineInitFailed,
+						exitCode,
+					)
 				}
-			case *proto.InitResponse_Stderr:
-				if resp.Stderr != nil {
-					outputLine.Stderr = resp.Stderr.GetContent()
-				}
-			case *proto.InitResponse_ExitResult:
-				if resp.ExitResult != nil {
-					exitCode := int(resp.ExitResult.GetCode())
-					if exitCode != 0 {
-						l.Errorf("Engine init failed with exit code %d", exitCode)
+			}
 
-						return nil, fmt.Errorf(
-							"%w with exit code %d",
-							ErrEngineInitFailed,
-							exitCode,
-						)
-					}
+		case *proto.InitResponse_Log:
+			if resp.Log != nil {
+				if logContent := resp.Log.GetContent(); logContent != "" {
+					logEngineMessage(l, resp.Log.GetLevel(), logContent)
 				}
-			case *proto.InitResponse_Log:
-				if resp.Log != nil {
-					if logContent := resp.Log.GetContent(); logContent != "" {
-						logEngineMessage(l, resp.Log.GetLevel(), logContent)
-					}
-				}
 			}
+		}
 
-			return outputLine, nil
-		})
+		return outputLine, nil
 	})
 }
 
@@ -1207,75 +1166,74 @@ func shutdown(
 	runOptions *ExecutionOptions,
 	terragruntEngine *proto.EngineClient,
 ) error {
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "engine_shutdown", map[string]any{
-		"cache_dir": runOptions.CacheDir,
-	}, func(ctx context.Context, l log.Logger) error {
-		meta, err := ConvertMetaToProtobuf(runOptions.EngineConfig.Meta)
+	meta, err := ConvertMetaToProtobuf(runOptions.EngineConfig.Meta)
+	if err != nil {
+		return err
+	}
+
+	request, err := (*terragruntEngine).Shutdown(ctx, &proto.ShutdownRequest{
+		WorkingDir: runOptions.CacheDir,
+		Meta:       meta,
+		EnvVars:    v.Env,
+	})
+	if err != nil {
+		return err
+	}
+
+	l.Debugf("Reading shutdown output for engine in %s", runOptions.CacheDir)
+
+	return ReadEngineOutput(v, true, func() (*OutputLine, error) {
+		output, err := request.Recv()
 		if err != nil {
-			return err
+			return nil, err
 		}
 
-		request, err := (*terragruntEngine).Shutdown(ctx, &proto.ShutdownRequest{
-			WorkingDir: runOptions.CacheDir,
-			Meta:       meta,
-			EnvVars:    v.Env,
-		})
-		if err != nil {
-			return err
+		if output == nil {
+			return nil, nil
 		}
 
-		l.Debugf("Reading shutdown output for engine in %s", runOptions.CacheDir)
+		outputLine := &OutputLine{}
 
-		return ReadEngineOutput(v, true, func() (*OutputLine, error) {
-			output, err := request.Recv()
-			if err != nil {
-				return nil, err
+		responseType := output.GetResponse()
+		if responseType == nil {
+			return outputLine, nil
+		}
+
+		//nolint:dupl // Similar structure to init response handling, but different protobuf types
+		switch resp := responseType.(type) {
+		case *proto.ShutdownResponse_Stdout:
+			if resp.Stdout != nil {
+				outputLine.Stdout = resp.Stdout.GetContent()
 			}
 
-			if output == nil {
-				return nil, nil
+		case *proto.ShutdownResponse_Stderr:
+			if resp.Stderr != nil {
+				outputLine.Stderr = resp.Stderr.GetContent()
 			}
 
-			outputLine := &OutputLine{}
+		case *proto.ShutdownResponse_ExitResult:
+			if resp.ExitResult != nil {
+				exitCode := int(resp.ExitResult.GetCode())
+				if exitCode != 0 {
+					l.Errorf("Engine shutdown failed with exit code %d", exitCode)
 
-			responseType := output.GetResponse()
-			if responseType == nil {
-				return outputLine, nil
+					return nil, fmt.Errorf(
+						"%w with exit code %d",
+						ErrEngineShutdownFailed,
+						exitCode,
+					)
+				}
 			}
 
-			//nolint:dupl // Similar structure to init response handling, but different protobuf types
-			switch resp := responseType.(type) {
-			case *proto.ShutdownResponse_Stdout:
-				if resp.Stdout != nil {
-					outputLine.Stdout = resp.Stdout.GetContent()
+		case *proto.ShutdownResponse_Log:
+			if resp.Log != nil {
+				if logContent := resp.Log.GetContent(); logContent != "" {
+					logEngineMessage(l, resp.Log.GetLevel(), logContent)
 				}
-			case *proto.ShutdownResponse_Stderr:
-				if resp.Stderr != nil {
-					outputLine.Stderr = resp.Stderr.GetContent()
-				}
-			case *proto.ShutdownResponse_ExitResult:
-				if resp.ExitResult != nil {
-					exitCode := int(resp.ExitResult.GetCode())
-					if exitCode != 0 {
-						l.Errorf("Engine shutdown failed with exit code %d", exitCode)
-
-						return nil, fmt.Errorf(
-							"%w with exit code %d",
-							ErrEngineShutdownFailed,
-							exitCode,
-						)
-					}
-				}
-			case *proto.ShutdownResponse_Log:
-				if resp.Log != nil {
-					if logContent := resp.Log.GetContent(); logContent != "" {
-						logEngineMessage(l, resp.Log.GetLevel(), logContent)
-					}
-				}
 			}
+		}
 
-			return outputLine, nil
-		})
+		return outputLine, nil
 	})
 }
 
