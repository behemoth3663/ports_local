--- pkg/config/dependency.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/config/dependency.go
@@ -50,14 +50,10 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/runner/runcfg"
 	"github.com/gruntwork-io/terragrunt/internal/shell"
 	"github.com/gruntwork-io/terragrunt/internal/spinner"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tf"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
 	"github.com/gruntwork-io/terragrunt/pkg/config/hclparse"
-
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 )
 
 const (
@@ -1769,70 +1765,49 @@ func getOutputJSONWithCaching(
 		fetchAttrs["dependency_name"] = name
 	}
 
-	err := telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "dependency_output_fetch", fetchAttrs, func(fetchCtx context.Context, l log.Logger) error {
-			jsonCache := cache.ContextCache[[]byte](fetchCtx, JSONOutputCacheContextKey)
-			if cached, found := jsonCache.Get(fetchCtx, targetConfig); found {
-				l.Debugf("%s was run before. Using cached output.", targetConfig)
+	_ = fetchAttrs
+	jsonCache := cache.ContextCache[[]byte](ctx, JSONOutputCacheContextKey)
+	if cached, found := jsonCache.Get(ctx, targetConfig); found {
+		l.Debugf("%s was run before. Using cached output.", targetConfig)
+		newJSONBytes = cached
+		return newJSONBytes, nil
+	}
 
-				if span := trace.SpanFromContext(fetchCtx); span.IsRecording() {
-					span.SetAttributes(attribute.Bool("cache_hit", true))
-				}
+	var (
+		fetched  []byte
+		strategy string
+	)
 
-				newJSONBytes = cached
+	target := util.RelPathForLog(pctx.RootWorkingDir, targetConfig, pctx.LogShowAbsPaths)
 
-				return nil
-			}
+	err := spinner.ShowAfter(ctx, l, spinner.Messages{
+		Working: "Fetching outputs of dependency " + target + "...",
+		Done:    "Fetched outputs of dependency " + target,
+	}, func() error {
+		var err error
 
-			var (
-				fetched  []byte
-				strategy string
-			)
+		fetched, strategy, err = resolveOutputJSON(ctx, l, v, pctx, targetConfig)
+		if err != nil {
+			return err
+		}
 
-			target := util.RelPathForLog(pctx.RootWorkingDir, targetConfig, pctx.LogShowAbsPaths)
+		trimmed, trimErr := tf.ExtractFirstJSONObject(fetched)
+		if trimErr != nil {
+			return TerragruntOutputParsingError{Path: targetConfig, Err: trimErr}
+		}
 
-			fetchErr := spinner.ShowAfter(fetchCtx, l, spinner.Messages{
-				Working: "Fetching outputs of dependency " + target + "...",
-				Done:    "Fetched outputs of dependency " + target,
-			}, func() error {
-				var err error
+		fetched = trimmed
 
-				fetched, strategy, err = resolveOutputJSON(fetchCtx, l, v, pctx, targetConfig)
-				if err != nil {
-					return err
-				}
-
-				trimmed, trimErr := tf.ExtractFirstJSONObject(fetched)
-				if trimErr != nil {
-					return TerragruntOutputParsingError{Path: targetConfig, Err: trimErr}
-				}
-
-				fetched = trimmed
-
-				return nil
-			})
-
-			if span := trace.SpanFromContext(fetchCtx); span.IsRecording() {
-				span.SetAttributes(attribute.Bool("cache_hit", false))
-
-				if strategy != "" {
-					span.SetAttributes(attribute.String("strategy", strategy))
-				}
-			}
-
-			if fetchErr != nil {
-				return fetchErr
-			}
-
-			newJSONBytes = fetched
-			jsonCache.Put(fetchCtx, targetConfig, newJSONBytes)
-
-			return nil
-		})
+		return nil
+	})
+	_ = strategy
 	if err != nil {
 		return nil, err
 	}
 
+	newJSONBytes = fetched
+	jsonCache.Put(ctx, targetConfig, newJSONBytes)
+
 	return newJSONBytes, nil
 }
 
@@ -2765,7 +2740,6 @@ func RunOptionsFromParsingContext(pctx *ParsingContext
 	runOpts.Debug = pctx.Debug
 	runOpts.AutoInit = pctx.AutoInit
 	runOpts.BackendBootstrap = pctx.BackendBootstrap
-	runOpts.Telemetry = pctx.Telemetry
 	runOpts.AuthProviderCmd = pctx.AuthProviderCmd
 	runOpts.NoCAS = pctx.NoCAS
 	runOpts.CASCloneDepth = pctx.CASCloneDepth
@@ -2780,7 +2754,6 @@ func shellRunOptsFromPctx(v *venv.Venv, pctx *ParsingC
 func shellRunOptsFromPctx(v *venv.Venv, pctx *ParsingContext) *shell.ShellOptions {
 	s := shell.NewShellOptions(v.Env).
 		WithWorkingDir(pctx.WorkingDir).
-		WithTelemetry(pctx.Telemetry).
 		WithEngine(pctx.EngineConfig, pctx.EngineOptions).
 		WithTFPath(pctx.TFPath).
 		WithRootWorkingDir(pctx.RootWorkingDir).
