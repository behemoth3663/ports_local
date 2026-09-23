--- pkg/config/dependency.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/config/dependency.go
@@ -47,14 +47,10 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds/providers/amazonsts"
 	"github.com/gruntwork-io/terragrunt/internal/runner/runcfg"
 	"github.com/gruntwork-io/terragrunt/internal/shell"
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
@@ -1594,59 +1590,39 @@ func getOutputJSONWithCaching(
 		fetchAttrs["dependency_name"] = name
 	}
 
-	err := telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "dependency_output_fetch", fetchAttrs, func(fetchCtx context.Context, l log.Logger) error {
-			jsonCache := cache.ContextCache[[]byte](fetchCtx, JSONOutputCacheContextKey)
-			if cached, found := jsonCache.Get(fetchCtx, targetConfig); found {
-				l.Debugf("%s was run before. Using cached output.", targetConfig)
+		jsonCache := cache.ContextCache[[]byte](ctx, JSONOutputCacheContextKey)
+		if cached, found := jsonCache.Get(ctx, targetConfig); found {
+			l.Debugf("%s was run before. Using cached output.", targetConfig)
 
-				if span := trace.SpanFromContext(fetchCtx); span.IsRecording() {
-					span.SetAttributes(attribute.Bool("cache_hit", true))
-				}
+			newJSONBytes = cached
 
-				newJSONBytes = cached
+			return newJSONBytes, nil
+		}
 
-				return nil
-			}
+		fetched, _, fetchErr := resolveOutputJSON(ctx, pctx, l, targetConfig)
 
-			fetched, strategy, fetchErr := resolveOutputJSON(fetchCtx, pctx, l, targetConfig)
+		if fetchErr != nil {
+			return nil, fetchErr
+		}
 
-			if span := trace.SpanFromContext(fetchCtx); span.IsRecording() {
-				span.SetAttributes(attribute.Bool("cache_hit", false))
+		// `tofu/terraform output -json` stdout can be polluted with non-JSON text on either side of the JSON object:
+		//   - Leading: AWS Client Side Monitoring (CSM) logs (e.g., "2023/05/04 20:22:43 Enabling CSM"),
+		//     ANSI color escape sequences from warning blocks.
+		//     Refs: https://github.com/aws/aws-sdk-go/blob/81d1cbbc6a2028023aff7bcab0fe1be320cd39f7/aws/session/session.go#L444
+		//           https://github.com/gruntwork-io/terragrunt/issues/2233
+		//   - Trailing: Terraform 1.15+ emits backend deprecation warnings (e.g., for the S3
+		//     `dynamodb_table` parameter) on stdout after the JSON has already been printed.
+		//     Refs: https://github.com/gruntwork-io/terragrunt/issues/6001
+		//
+		// To make parsing robust to either, isolate the first JSON object in the buffer.
+		trimmed, trimErr := extractFirstJSONObject(fetched)
+		if trimErr != nil {
+			return nil, TerragruntOutputParsingError{Path: targetConfig, Err: trimErr}
+		}
 
-				if strategy != "" {
-					span.SetAttributes(attribute.String("strategy", strategy))
-				}
-			}
+		newJSONBytes = trimmed
+		jsonCache.Put(ctx, targetConfig, newJSONBytes)
 
-			if fetchErr != nil {
-				return fetchErr
-			}
-
-			// `tofu/terraform output -json` stdout can be polluted with non-JSON text on either side of the JSON object:
-			//   - Leading: AWS Client Side Monitoring (CSM) logs (e.g., "2023/05/04 20:22:43 Enabling CSM"),
-			//     ANSI color escape sequences from warning blocks.
-			//     Refs: https://github.com/aws/aws-sdk-go/blob/81d1cbbc6a2028023aff7bcab0fe1be320cd39f7/aws/session/session.go#L444
-			//           https://github.com/gruntwork-io/terragrunt/issues/2233
-			//   - Trailing: Terraform 1.15+ emits backend deprecation warnings (e.g., for the S3
-			//     `dynamodb_table` parameter) on stdout after the JSON has already been printed.
-			//     Refs: https://github.com/gruntwork-io/terragrunt/issues/6001
-			//
-			// To make parsing robust to either, isolate the first JSON object in the buffer.
-			trimmed, trimErr := extractFirstJSONObject(fetched)
-			if trimErr != nil {
-				return TerragruntOutputParsingError{Path: targetConfig, Err: trimErr}
-			}
-
-			newJSONBytes = trimmed
-			jsonCache.Put(fetchCtx, targetConfig, newJSONBytes)
-
-			return nil
-		})
-	if err != nil {
-		return nil, err
-	}
-
 	return newJSONBytes, nil
 }
 
@@ -2547,7 +2523,6 @@ func RunOptionsFromParsingContext(pctx *ParsingContext
 	runOpts.Debug = pctx.Debug
 	runOpts.AutoInit = pctx.AutoInit
 	runOpts.BackendBootstrap = pctx.BackendBootstrap
-	runOpts.Telemetry = pctx.Telemetry
 	runOpts.AuthProviderCmd = pctx.AuthProviderCmd
 	runOpts.NoCAS = pctx.NoCAS
 	runOpts.CASCloneDepth = pctx.CASCloneDepth
@@ -2562,7 +2537,6 @@ func shellRunOptsFromPctx(pctx *ParsingContext) *shell
 func shellRunOptsFromPctx(pctx *ParsingContext) *shell.ShellOptions {
 	s := shell.NewShellOptions(pctx.Venv.Env).
 		WithWorkingDir(pctx.WorkingDir).
-		WithTelemetry(pctx.Telemetry).
 		WithEngine(pctx.EngineConfig, pctx.EngineOptions).
 		WithTFPath(pctx.TFPath).
 		WithRootWorkingDir(pctx.RootWorkingDir).
