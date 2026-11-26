--- pkg/config/dependency.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/config/dependency.go
@@ -51,15 +51,10 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds/providers/amazonsts"
 	"github.com/gruntwork-io/terragrunt/internal/runner/runcfg"
 	"github.com/gruntwork-io/terragrunt/internal/shell"
-	"github.com/gruntwork-io/terragrunt/internal/spinner"
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
@@ -1697,70 +1692,28 @@ func getOutputJSONWithCaching(
 		fetchAttrs["dependency_name"] = name
 	}
 
-	err := telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "dependency_output_fetch", fetchAttrs, func(fetchCtx context.Context, l log.Logger) error {
-			jsonCache := cache.ContextCache[[]byte](fetchCtx, JSONOutputCacheContextKey)
-			if cached, found := jsonCache.Get(fetchCtx, targetConfig); found {
+			jsonCache := cache.ContextCache[[]byte](ctx, JSONOutputCacheContextKey)
+			if cached, found := jsonCache.Get(ctx, targetConfig); found {
 				l.Debugf("%s was run before. Using cached output.", targetConfig)
 
-				if span := trace.SpanFromContext(fetchCtx); span.IsRecording() {
-					span.SetAttributes(attribute.Bool("cache_hit", true))
-				}
-
-				newJSONBytes = cached
-
-				return nil
+				return cached, nil
 			}
 
-			var (
-				fetched  []byte
-				strategy string
-			)
-
-			target := util.RelPathForLog(pctx.RootWorkingDir, targetConfig, pctx.LogShowAbsPaths)
-
-			fetchErr := spinner.ShowAfter(fetchCtx, l, spinner.Messages{
-				Working: "Fetching outputs of dependency " + target + "...",
-				Done:    "Fetched outputs of dependency " + target,
-			}, func() error {
-				var err error
-
-				fetched, strategy, err = resolveOutputJSON(fetchCtx, l, v, s, pctx, targetConfig)
+				fetched, _, err := resolveOutputJSON(ctx, l, v, s, pctx, targetConfig)
 				if err != nil {
-					return err
+					return nil, err
 				}
 
 				trimmed, trimErr := tf.ExtractFirstJSONObject(fetched)
 				if trimErr != nil {
-					return TerragruntOutputParsingError{Path: targetConfig, Err: trimErr}
+					return nil, trimErr
 				}
 
 				fetched = trimmed
 
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
 			newJSONBytes = fetched
-			jsonCache.Put(fetchCtx, targetConfig, newJSONBytes)
+			jsonCache.Put(ctx, targetConfig, newJSONBytes)
 
-			return nil
-		})
-	if err != nil {
-		return nil, err
-	}
-
 	return newJSONBytes, nil
 }
 
@@ -2712,7 +2665,6 @@ func RunOptionsFromParsingContext(pctx *ParsingContext
 	runOpts.Debug = pctx.Debug
 	runOpts.AutoInit = pctx.AutoInit
 	runOpts.BackendBootstrap = pctx.BackendBootstrap
-	runOpts.Telemetry = pctx.Telemetry
 	runOpts.AuthProviderCmd = pctx.AuthProviderCmd
 	runOpts.NoCAS = pctx.NoCAS
 	runOpts.CASCloneDepth = pctx.CASCloneDepth
@@ -2727,7 +2679,6 @@ func shellRunOptsFromPctx(v *venv.Venv, pctx *ParsingC
 func shellRunOptsFromPctx(v *venv.Venv, pctx *ParsingContext) *shell.ShellOptions {
 	s := shell.NewShellOptions(v.Env).
 		WithWorkingDir(pctx.WorkingDir).
-		WithTelemetry(pctx.Telemetry).
 		WithEngine(pctx.EngineConfig, pctx.EngineOptions).
 		WithTFPath(tfimpl.BinaryOrDefault(v.Exec, pctx.TFPath)).
 		WithRootWorkingDir(pctx.RootWorkingDir).
