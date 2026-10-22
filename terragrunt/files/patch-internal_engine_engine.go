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
@@ -236,17 +234,8 @@ func Run(
 
 	var output *util.CmdOutput
 
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
+	_ = found
+	output, runErr := invoke(ctx, l, v, execOptions, instance.engineClient)
 	if runErr != nil {
 		return nil, runErr
 	}
@@ -314,108 +303,103 @@ func downloadEngine(
 		return nil
 	}
 
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "engine_download", map[string]any{
-		"source":  e.Source,
-		"version": e.Version,
-	}, func(ctx context.Context, l log.Logger) error {
-		// If source is empty, we cannot download the engine
-		// This indicates an engine block was configured but source was not provided
-		if e.Source == "" {
-			return errors.New(
-				"engine block is configured but source is empty, please provide an engine source or remove the engine block",
-			)
-		}
+	// If source is empty, we cannot download the engine
+	// This indicates an engine block was configured but source was not provided
+	if e.Source == "" {
+		return errors.New(
+			"engine block is configured but source is empty, please provide an engine source or remove the engine block",
+		)
+	}
 
-		// identify engine version if not specified
-		if len(e.Version) == 0 {
-			if !isDirectURL(e.Source) {
-				v.RequireHTTP()
+	// identify engine version if not specified
+	if len(e.Version) == 0 {
+		if !isDirectURL(e.Source) {
+			v.RequireHTTP()
 
-				tag, err := lastReleaseVersion(ctx, v, execOptions)
-				if err != nil {
-					return err
-				}
-
-				e.Version = tag
+			tag, err := lastReleaseVersion(ctx, v, execOptions)
+			if err != nil {
+				return err
 			}
-		}
 
-		path, err := engineDir(v, execOptions)
-		if err != nil {
-			return err
+			e.Version = tag
 		}
+	}
 
-		if ensureErr := vfs.EnsureDirectory(v.FS, path); ensureErr != nil {
-			return ensureErr
-		}
+	path, err := engineDir(v, execOptions)
+	if err != nil {
+		return err
+	}
 
-		localEngineFile := filepath.Join(path, engineFileName(v, e))
+	if ensureErr := vfs.EnsureDirectory(v.FS, path); ensureErr != nil {
+		return ensureErr
+	}
 
-		// lock downloading process for only one instance
-		locks, err := downloadLocksFromContext(ctx)
-		if err != nil {
-			return err
-		}
-		// locking by file where engine is downloaded
-		// however, it will not help in case of multiple parallel Terragrunt runs
-		locks.Lock(localEngineFile)
-		defer locks.Unlock(localEngineFile)
+	localEngineFile := filepath.Join(path, engineFileName(v, e))
 
-		if vfs.Exists(v.FS, localEngineFile) {
-			return nil
-		}
+	// lock downloading process for only one instance
+	locks, err := downloadLocksFromContext(ctx)
+	if err != nil {
+		return err
+	}
+	// locking by file where engine is downloaded
+	// however, it will not help in case of multiple parallel Terragrunt runs
+	locks.Lock(localEngineFile)
+	defer locks.Unlock(localEngineFile)
 
-		downloadFile := filepath.Join(path, enginePackageName(v, e))
+	if vfs.Exists(v.FS, localEngineFile) {
+		return nil
+	}
 
-		// Prepare download assets
-		assets := &github.ReleaseAssets{
-			Repository:  e.Source,
-			Version:     e.Version,
-			PackageFile: downloadFile,
-		}
+	downloadFile := filepath.Join(path, enginePackageName(v, e))
 
-		var checksumFile, checksumSigFile string
+	// Prepare download assets
+	assets := &github.ReleaseAssets{
+		Repository:  e.Source,
+		Version:     e.Version,
+		PackageFile: downloadFile,
+	}
 
-		// Only add checksum files for GitHub releases (not direct URLs)
-		if !isDirectURL(e.Source) {
-			checksumFile = filepath.Join(path, engineChecksumName(e))
-			checksumSigFile = filepath.Join(path, engineChecksumSigName(e))
-			assets.ChecksumFile = checksumFile
-			assets.ChecksumSigFile = checksumSigFile
-		}
+	var checksumFile, checksumSigFile string
 
-		// Create download client and download assets
-		downloadClient := github.NewGitHubReleasesDownloadClient(l)
+	// Only add checksum files for GitHub releases (not direct URLs)
+	if !isDirectURL(e.Source) {
+		checksumFile = filepath.Join(path, engineChecksumName(e))
+		checksumSigFile = filepath.Join(path, engineChecksumSigName(e))
+		assets.ChecksumFile = checksumFile
+		assets.ChecksumSigFile = checksumSigFile
+	}
 
-		result, err := downloadClient.DownloadReleaseAssets(ctx, v, assets)
-		if err != nil {
-			return fmt.Errorf("failed to download engine assets: %w", err)
-		}
+	// Create download client and download assets
+	downloadClient := github.NewGitHubReleasesDownloadClient(l)
 
-		// Update file paths from result
-		downloadFile = result.PackageFile
-		checksumFile = result.ChecksumFile
-		checksumSigFile = result.ChecksumSigFile
+	result, err := downloadClient.DownloadReleaseAssets(ctx, v, assets)
+	if err != nil {
+		return fmt.Errorf("failed to download engine assets: %w", err)
+	}
 
-		if !execOptions.EngineOptions.SkipChecksumCheck && checksumFile != "" &&
-			checksumSigFile != "" {
-			l.Infof("Verifying checksum for %s", downloadFile)
+	// Update file paths from result
+	downloadFile = result.PackageFile
+	checksumFile = result.ChecksumFile
+	checksumSigFile = result.ChecksumSigFile
 
-			if err := verifyFile(v.FS, downloadFile, checksumFile, checksumSigFile); err != nil {
-				return err
-			}
-		} else {
-			l.Warnf("Skipping verification for %s", downloadFile)
-		}
+	if !execOptions.EngineOptions.SkipChecksumCheck && checksumFile != "" &&
+		checksumSigFile != "" {
+		l.Infof("Verifying checksum for %s", downloadFile)
 
-		if err := extractArchive(l, v.FS, downloadFile, localEngineFile); err != nil {
+		if err := verifyFile(v.FS, downloadFile, checksumFile, checksumSigFile); err != nil {
 			return err
 		}
+	} else {
+		l.Warnf("Skipping verification for %s", downloadFile)
+	}
 
-		l.Infof("Engine available as %s", path)
+	if err := extractArchive(l, v.FS, downloadFile, localEngineFile); err != nil {
+		return err
+	}
 
-		return nil
-	})
+	l.Infof("Engine available as %s", path)
+
+	return nil
 }
 
 func lastReleaseVersion(
@@ -821,120 +805,109 @@ func createEngine(
 		pluginClient *plugin.Client
 	)
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "engine_create", map[string]any{
-		"source":    execOptions.EngineConfig.Source,
-		"version":   execOptions.EngineConfig.Version,
-		"cache_dir": execOptions.CacheDir,
-	}, func(ctx context.Context, l log.Logger) error {
-		path, err := engineDir(v, execOptions)
-		if err != nil {
-			return err
-		}
+	path, err := engineDir(v, execOptions)
+	if err != nil {
+		return nil, nil, err
+	}
 
-		localEnginePath := filepath.Join(path, engineFileName(v, execOptions.EngineConfig))
-		localChecksumFile := filepath.Join(path, engineChecksumName(execOptions.EngineConfig))
-		localChecksumSigFile := filepath.Join(path, engineChecksumSigName(execOptions.EngineConfig))
+	localEnginePath := filepath.Join(path, engineFileName(v, execOptions.EngineConfig))
+	localChecksumFile := filepath.Join(path, engineChecksumName(execOptions.EngineConfig))
+	localChecksumSigFile := filepath.Join(path, engineChecksumSigName(execOptions.EngineConfig))
 
-		// validate engine before loading if verification is not disabled
-		skipCheck := execOptions.EngineOptions.SkipChecksumCheck
-		if !skipCheck && vfs.Exists(v.FS, localEnginePath) && vfs.Exists(v.FS, localChecksumFile) &&
-			vfs.Exists(v.FS, localChecksumSigFile) {
-			if err = verifyFile(
-				v.FS,
-				localEnginePath,
-				localChecksumFile,
-				localChecksumSigFile,
-			); err != nil {
-				return err
-			}
-		} else {
-			l.Warnf("Skipping verification for %s", localEnginePath)
+	// validate engine before loading if verification is not disabled
+	skipCheck := execOptions.EngineOptions.SkipChecksumCheck
+	if !skipCheck && vfs.Exists(v.FS, localEnginePath) && vfs.Exists(v.FS, localChecksumFile) &&
+		vfs.Exists(v.FS, localChecksumSigFile) {
+		if err = verifyFile(
+			v.FS,
+			localEnginePath,
+			localChecksumFile,
+			localChecksumSigFile,
+		); err != nil {
+			return nil, nil, err
 		}
+	} else {
+		l.Warnf("Skipping verification for %s", localEnginePath)
+	}
 
-		l.Debugf("Creating engine %s", localEnginePath)
+	l.Debugf("Creating engine %s", localEnginePath)
 
-		engineLogLevel := execOptions.EngineOptions.LogLevel
+	engineLogLevel := execOptions.EngineOptions.LogLevel
 
-		if len(engineLogLevel) == 0 {
-			engineLogLevel = hclog.Warn.String()
-			// update log level if it is different from info
-			if l.Level() != log.InfoLevel {
-				engineLogLevel = l.Level().String()
-			}
-			// turn off log formatting if disabled for Terragrunt
-			if l.Formatter().DisabledOutput() {
-				engineLogLevel = hclog.Off.String()
-			}
+	if len(engineLogLevel) == 0 {
+		engineLogLevel = hclog.Warn.String()
+		// update log level if it is different from info
+		if l.Level() != log.InfoLevel {
+			engineLogLevel = l.Level().String()
 		}
-
-		logger := hclog.NewInterceptLogger(&hclog.LoggerOptions{
-			Level:  hclog.LevelFromString(engineLogLevel),
-			Output: l.Writer(),
-		})
-
-		// We use without cancel here to ensure that the plugin isn't killed when the main context is cancelled,
-		// like it is in the RunCommandWithOutput function. This ensures that we don't cancel the shutdown
-		// when the command is cancelled.
-		cmd := v.Exec.Command(context.WithoutCancel(ctx), localEnginePath)
-		cmd.SetEnv([]string{fmt.Sprintf("%s=%s", engineLogLevelEnv, engineLogLevel)})
-		cmd.SetCancel(func() error {
-			sig := signal.SignalFromContext(ctx)
-			if sig == nil {
-				sig = os.Kill
-			}
-
-			if err := cmd.Signal(sig); err != nil && !errors.Is(err, vexec.ErrProcessNotStarted) {
-				return err
-			}
-
-			return nil
-		})
-
-		// hashicorp/go-plugin's ClientConfig requires a concrete *exec.Cmd.
-		osCmder, ok := cmd.(vexec.OSCmder)
-		if !ok {
-			return fmt.Errorf("engine plugin spawn: %w", vexec.ErrNotOSBacked)
+		// turn off log formatting if disabled for Terragrunt
+		if l.Formatter().DisabledOutput() {
+			engineLogLevel = hclog.Off.String()
 		}
+	}
 
-		client := plugin.NewClient(&plugin.ClientConfig{
-			Logger:           logger,
-			ProtocolVersion:  engineVersion,
-			MagicCookieKey:   engineCookieKey,
-			MagicCookieValue: engineCookieValue,
-			Plugins: map[string]plugin.Plugin{
-				"plugin": &engine.TerragruntGRPCEngine{},
-			},
-			Cmd: osCmder.OSCmd(),
-			GRPCDialOptions: []grpc.DialOption{
-				grpc.WithTransportCredentials(insecure.NewCredentials()),
-			},
-			AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
-		})
+	logger := hclog.NewInterceptLogger(&hclog.LoggerOptions{
+		Level:  hclog.LevelFromString(engineLogLevel),
+		Output: l.Writer(),
+	})
 
-		rpcClient, err := client.Client()
-		if err != nil {
-			return err
+	// We use without cancel here to ensure that the plugin isn't killed when the main context is cancelled,
+	// like it is in the RunCommandWithOutput function. This ensures that we don't cancel the shutdown
+	// when the command is cancelled.
+	cmd := v.Exec.Command(context.WithoutCancel(ctx), localEnginePath)
+	cmd.SetEnv([]string{fmt.Sprintf("%s=%s", engineLogLevelEnv, engineLogLevel)})
+	cmd.SetCancel(func() error {
+		sig := signal.SignalFromContext(ctx)
+		if sig == nil {
+			sig = os.Kill
 		}
 
-		rawClient, err := rpcClient.Dispense("plugin")
-		if err != nil {
+		if err := cmd.Signal(sig); err != nil && !errors.Is(err, vexec.ErrProcessNotStarted) {
 			return err
 		}
 
-		terragruntEngine, ok := rawClient.(proto.EngineClient)
-		if !ok {
-			return fmt.Errorf("engine plugin returned unexpected client type %T", rawClient)
-		}
+		return nil
+	})
 
-		engineClient = &terragruntEngine
-		pluginClient = client
+	// hashicorp/go-plugin's ClientConfig requires a concrete *exec.Cmd.
+	osCmder, ok := cmd.(vexec.OSCmder)
+	if !ok {
+		return nil, nil, fmt.Errorf("engine plugin spawn: %w", vexec.ErrNotOSBacked)
+	}
 
-		return nil
+	client := plugin.NewClient(&plugin.ClientConfig{
+		Logger:           logger,
+		ProtocolVersion:  engineVersion,
+		MagicCookieKey:   engineCookieKey,
+		MagicCookieValue: engineCookieValue,
+		Plugins: map[string]plugin.Plugin{
+			"plugin": &engine.TerragruntGRPCEngine{},
+		},
+		Cmd: osCmder.OSCmd(),
+		GRPCDialOptions: []grpc.DialOption{
+			grpc.WithTransportCredentials(insecure.NewCredentials()),
+		},
+		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
 	})
+
+	rpcClient, err := client.Client()
 	if err != nil {
 		return nil, nil, err
 	}
 
+	rawClient, err := rpcClient.Dispense("plugin")
+	if err != nil {
+		return nil, nil, err
+	}
+
+	terragruntEngine, ok := rawClient.(proto.EngineClient)
+	if !ok {
+		return nil, nil, fmt.Errorf("engine plugin returned unexpected client type %T", rawClient)
+	}
+
+	engineClient = &terragruntEngine
+	pluginClient = client
+
 	return engineClient, pluginClient, nil
 }
 
@@ -948,145 +921,133 @@ func invoke(
 ) (*util.CmdOutput, error) {
 	var result *util.CmdOutput
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "engine_invoke", map[string]any{
-		"command":   runOptions.Command,
-		"cache_dir": runOptions.CacheDir,
-	}, func(ctx context.Context, l log.Logger) error {
-		l = l.WithField(placeholders.TFPathKeyName, "engine")
+	l = l.WithField(placeholders.TFPathKeyName, "engine")
 
-		meta, err := ConvertMetaToProtobuf(runOptions.EngineConfig.Meta)
-		if err != nil {
-			return err
-		}
+	meta, err := ConvertMetaToProtobuf(runOptions.EngineConfig.Meta)
+	if err != nil {
+		return nil, err
+	}
 
-		response, err := (*client).Run(ctx, &proto.RunRequest{
-			Command:           runOptions.Command,
-			Args:              runOptions.Args,
-			AllocatePseudoTty: runOptions.AllocatePseudoTty,
-			WorkingDir:        runOptions.CacheDir,
-			Meta:              meta,
-			EnvVars:           v.Env,
-		})
-		if err != nil {
-			return err
-		}
+	response, err := (*client).Run(ctx, &proto.RunRequest{
+		Command:           runOptions.Command,
+		Args:              runOptions.Args,
+		AllocatePseudoTty: runOptions.AllocatePseudoTty,
+		WorkingDir:        runOptions.CacheDir,
+		Meta:              meta,
+		EnvVars:           v.Env,
+	})
+	if err != nil {
+		return nil, err
+	}
 
-		// Determine log levels based on headless mode (similar to buildOutWriter/buildErrWriter)
-		stdoutLogLevel := log.StdoutLevel
-		stderrLogLevel := log.StderrLevel
+	// Determine log levels based on headless mode (similar to buildOutWriter/buildErrWriter)
+	stdoutLogLevel := log.StdoutLevel
+	stderrLogLevel := log.StderrLevel
 
-		stdoutWriter := writer.ExtractOriginalWriter(v.Writers.Writer)
-		stderrWriter := writer.ExtractOriginalWriter(v.Writers.ErrWriter)
+	stdoutWriter := writer.ExtractOriginalWriter(v.Writers.Writer)
+	stderrWriter := writer.ExtractOriginalWriter(v.Writers.ErrWriter)
 
-		if runOptions.Headless && !runOptions.ForwardTFStdout {
-			stdoutLogLevel = log.InfoLevel
-			stderrLogLevel = log.ErrorLevel
-			stdoutWriter = writer.ExtractOriginalWriter(v.Writers.ErrWriter)
-		}
+	if runOptions.Headless && !runOptions.ForwardTFStdout {
+		stdoutLogLevel = log.InfoLevel
+		stderrLogLevel = log.ErrorLevel
+		stdoutWriter = writer.ExtractOriginalWriter(v.Writers.ErrWriter)
+	}
 
-		var (
-			output = util.CmdOutput{}
+	var (
+		output = util.CmdOutput{}
 
-			// Use the original output writers (before they were wrapped by logTFOutput)
-			// and create new writers with the engine logger
-			engineStdout = logwriter.New(
-				logwriter.WithLogger(l.WithOptions(log.WithOutput(stdoutWriter))),
-				logwriter.WithDefaultLevel(stdoutLogLevel),
-				logwriter.WithMsgSeparator("\n"),
-			)
-			engineStderr = logwriter.New(
-				logwriter.WithLogger(l.WithOptions(log.WithOutput(stderrWriter))),
-				logwriter.WithDefaultLevel(stderrLogLevel),
-				logwriter.WithMsgSeparator("\n"),
-			)
-
-			stdout = io.MultiWriter(engineStdout, &output.Stdout)
-			stderr = io.MultiWriter(engineStderr, &output.Stderr)
+		// Use the original output writers (before they were wrapped by logTFOutput)
+		// and create new writers with the engine logger
+		engineStdout = logwriter.New(
+			logwriter.WithLogger(l.WithOptions(log.WithOutput(stdoutWriter))),
+			logwriter.WithDefaultLevel(stdoutLogLevel),
+			logwriter.WithMsgSeparator("\n"),
 		)
-
-		var (
-			stdoutLineBuf, stderrLineBuf bytes.Buffer
-			resultCode                   int
+		engineStderr = logwriter.New(
+			logwriter.WithLogger(l.WithOptions(log.WithOutput(stderrWriter))),
+			logwriter.WithDefaultLevel(stderrLogLevel),
+			logwriter.WithMsgSeparator("\n"),
 		)
 
-		for {
-			runResp, recvErr := response.Recv()
-			if recvErr != nil || runResp == nil {
-				break
-			}
+		stdout = io.MultiWriter(engineStdout, &output.Stdout)
+		stderr = io.MultiWriter(engineStderr, &output.Stderr)
+	)
 
-			responseType := runResp.GetResponse()
-			if responseType == nil {
-				continue
-			}
+	var (
+		stdoutLineBuf, stderrLineBuf bytes.Buffer
+		resultCode                   int
+	)
 
-			switch resp := responseType.(type) {
-			case *proto.RunResponse_Stdout:
-				if resp.Stdout != nil {
-					if err = processStream(
-						resp.Stdout.GetContent(),
-						&stdoutLineBuf,
-						stdout,
-					); err != nil {
-						return err
-					}
+	for {
+		runResp, recvErr := response.Recv()
+		if recvErr != nil || runResp == nil {
+			break
+		}
+
+		responseType := runResp.GetResponse()
+		if responseType == nil {
+			continue
+		}
+
+		switch resp := responseType.(type) {
+		case *proto.RunResponse_Stdout:
+			if resp.Stdout != nil {
+				if err = processStream(
+					resp.Stdout.GetContent(),
+					&stdoutLineBuf,
+					stdout,
+				); err != nil {
+					return nil, err
 				}
-			case *proto.RunResponse_Stderr:
-				if resp.Stderr != nil {
-					if err = processStream(
-						resp.Stderr.GetContent(),
-						&stderrLineBuf,
-						stderr,
-					); err != nil {
-						return err
-					}
+			}
+		case *proto.RunResponse_Stderr:
+			if resp.Stderr != nil {
+				if err = processStream(
+					resp.Stderr.GetContent(),
+					&stderrLineBuf,
+					stderr,
+				); err != nil {
+					return nil, err
 				}
-			case *proto.RunResponse_ExitResult:
-				if resp.ExitResult != nil {
-					resultCode = int(resp.ExitResult.GetCode())
+			}
+		case *proto.RunResponse_ExitResult:
+			if resp.ExitResult != nil {
+				resultCode = int(resp.ExitResult.GetCode())
+			}
+		case *proto.RunResponse_Log:
+			if resp.Log != nil {
+				if logContent := resp.Log.GetContent(); logContent != "" {
+					logEngineMessage(l, resp.Log.GetLevel(), logContent)
 				}
-			case *proto.RunResponse_Log:
-				if resp.Log != nil {
-					if logContent := resp.Log.GetContent(); logContent != "" {
-						logEngineMessage(l, resp.Log.GetLevel(), logContent)
-					}
-				}
 			}
 		}
+	}
 
-		if err = flushBuffer(&stdoutLineBuf, stdout); err != nil {
-			return err
-		}
+	if err = flushBuffer(&stdoutLineBuf, stdout); err != nil {
+		return nil, err
+	}
 
-		if err = flushBuffer(&stderrLineBuf, stderr); err != nil {
-			return err
-		}
+	if err = flushBuffer(&stderrLineBuf, stderr); err != nil {
+		return nil, err
+	}
 
-		l.Debugf("Engine execution done in %v", runOptions.CacheDir)
+	l.Debugf("Engine execution done in %v", runOptions.CacheDir)
 
-		if resultCode != 0 {
-			err = &util.ProcessExecutionError{
-				Err:             fmt.Errorf("command failed with exit code %d", resultCode),
-				Output:          output,
-				WorkingDir:      runOptions.CacheDir,
-				RootWorkingDir:  runOptions.RootWorkingDir,
-				LogShowAbsPaths: runOptions.LogShowAbsPaths,
-				Command:         runOptions.Command,
-				Args:            runOptions.Args,
-				DisableSummary:  runOptions.LogDisableErrorSummary,
-			}
-
-			return err
+	if resultCode != 0 {
+		return nil, &util.ProcessExecutionError{
+			Err:             fmt.Errorf("command failed with exit code %d", resultCode),
+			Output:          output,
+			WorkingDir:      runOptions.CacheDir,
+			RootWorkingDir:  runOptions.RootWorkingDir,
+			LogShowAbsPaths: runOptions.LogShowAbsPaths,
+			Command:         runOptions.Command,
+			Args:            runOptions.Args,
+			DisableSummary:  runOptions.LogDisableErrorSummary,
 		}
-
-		result = &output
-
-		return nil
-	})
-	if err != nil {
-		return nil, err
 	}
 
+	result = &output
+
 	return result, nil
 }
 
@@ -1128,72 +1089,68 @@ func initialize(
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
-			}
+		outputLine := &OutputLine{}
 
-			if output == nil {
-				return nil, nil
+		//nolint:dupl // Similar structure to shutdown response handling, but different protobuf types
+		switch resp := output.GetResponse().(type) {
+		case *proto.InitResponse_Stdout:
+			if resp.Stdout != nil {
+				outputLine.Stdout = resp.Stdout.GetContent()
 			}
+		case *proto.InitResponse_Stderr:
+			if resp.Stderr != nil {
+				outputLine.Stderr = resp.Stderr.GetContent()
+			}
+		case *proto.InitResponse_ExitResult:
+			if resp.ExitResult != nil {
+				exitCode := int(resp.ExitResult.GetCode())
+				if exitCode != 0 {
+					l.Errorf("Engine init failed with exit code %d", exitCode)
 
-			outputLine := &OutputLine{}
-
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
+			}
+		case *proto.InitResponse_Log:
+			if resp.Log != nil {
+				if logContent := resp.Log.GetContent(); logContent != "" {
+					logEngineMessage(l, resp.Log.GetLevel(), logContent)
 				}
-			case *proto.InitResponse_ExitResult:
-				if resp.ExitResult != nil {
-					exitCode := int(resp.ExitResult.GetCode())
-					if exitCode != 0 {
-						l.Errorf("Engine init failed with exit code %d", exitCode)
-
-						return nil, fmt.Errorf(
-							"%w with exit code %d",
-							ErrEngineInitFailed,
-							exitCode,
-						)
-					}
-				}
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
 
@@ -1207,75 +1164,71 @@ func shutdown(
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
-			}
+		responseType := output.GetResponse()
+		if responseType == nil {
+			return outputLine, nil
+		}
 
-			if output == nil {
-				return nil, nil
+		//nolint:dupl // Similar structure to init response handling, but different protobuf types
+		switch resp := responseType.(type) {
+		case *proto.ShutdownResponse_Stdout:
+			if resp.Stdout != nil {
+				outputLine.Stdout = resp.Stdout.GetContent()
 			}
-
-			outputLine := &OutputLine{}
-
-			responseType := output.GetResponse()
-			if responseType == nil {
-				return outputLine, nil
+		case *proto.ShutdownResponse_Stderr:
+			if resp.Stderr != nil {
+				outputLine.Stderr = resp.Stderr.GetContent()
 			}
+		case *proto.ShutdownResponse_ExitResult:
+			if resp.ExitResult != nil {
+				exitCode := int(resp.ExitResult.GetCode())
+				if exitCode != 0 {
+					l.Errorf("Engine shutdown failed with exit code %d", exitCode)
 
-			//nolint:dupl // Similar structure to init response handling, but different protobuf types
-			switch resp := responseType.(type) {
-			case *proto.ShutdownResponse_Stdout:
-				if resp.Stdout != nil {
-					outputLine.Stdout = resp.Stdout.GetContent()
+					return nil, fmt.Errorf(
+						"%w with exit code %d",
+						ErrEngineShutdownFailed,
+						exitCode,
+					)
 				}
-			case *proto.ShutdownResponse_Stderr:
-				if resp.Stderr != nil {
-					outputLine.Stderr = resp.Stderr.GetContent()
+			}
+		case *proto.ShutdownResponse_Log:
+			if resp.Log != nil {
+				if logContent := resp.Log.GetContent(); logContent != "" {
+					logEngineMessage(l, resp.Log.GetLevel(), logContent)
 				}
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
 
