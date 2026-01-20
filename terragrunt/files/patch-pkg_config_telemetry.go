--- pkg/config/telemetry.go.orig	2026-01-19 21:45:47 UTC
+++ pkg/config/telemetry.go
@@ -5,10 +5,7 @@ import (
 	"context"
 	"strings"
 
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // Telemetry operation names for config parsing operations.
@@ -56,23 +53,7 @@ func TraceParseConfigFile(
 	cacheHit bool,
 	fn func(ctx context.Context) error,
 ) error {
-	attrs := map[string]any{
-		AttrConfigPath:       configPath,
-		AttrWorkingDir:       workingDir,
-		AttrIsPartial:        isPartial,
-		AttrCacheHit:         cacheHit,
-		AttrIncludeFromChild: includeFromChild != nil,
-	}
-
-	if len(decodeList) > 0 {
-		attrs[AttrDecodeList] = formatDecodeList(decodeList)
-	}
-
-	if includeFromChild != nil {
-		attrs[AttrIncludeChildPath] = includeFromChild.Path
-	}
-
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, TelemetryOpParseConfigFile, attrs, fn)
+	return nil
 }
 
 // TraceParseBaseBlocks wraps base blocks parsing with telemetry.
@@ -87,15 +68,7 @@ func TraceParseBaseBlocks(
 		resultErr error
 	)
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, TelemetryOpParseBaseBlocks, map[string]any{
-		AttrConfigPath: configPath,
-	}, func(childCtx context.Context) error {
-		result, resultErr = fn(childCtx)
-		return resultErr
-	})
-	if err != nil {
-		l.Warnf("Telemetry error during base blocks parsing: %v", err)
-	}
+		result, resultErr = fn(ctx)
 
 	return result, resultErr
 }
@@ -106,54 +79,7 @@ func TraceParseBaseBlocksResult(
 	configPath string,
 	baseBlocks *DecodedBaseBlocks,
 ) {
-	span := trace.SpanFromContext(ctx)
-	if span == nil || !span.IsRecording() {
 		return
-	}
-
-	attrs := []attribute.KeyValue{
-		attribute.String(AttrConfigPath, configPath),
-	}
-
-	if baseBlocks != nil {
-		// Include information
-		if baseBlocks.TrackInclude != nil && len(baseBlocks.TrackInclude.CurrentList) > 0 {
-			attrs = append(attrs,
-				attribute.Bool(AttrHasIncludes, true),
-				attribute.Int(AttrIncludeCount, len(baseBlocks.TrackInclude.CurrentList)),
-				attribute.String(AttrIncludePaths, formatIncludePaths(baseBlocks.TrackInclude.CurrentList)),
-			)
-		} else {
-			attrs = append(attrs,
-				attribute.Bool(AttrHasIncludes, false),
-				attribute.Int(AttrIncludeCount, 0),
-			)
-		}
-
-		// Locals information
-		if baseBlocks.Locals != nil && !baseBlocks.Locals.IsNull() {
-			localsMap := baseBlocks.Locals.AsValueMap()
-			attrs = append(attrs,
-				attribute.Int(AttrLocalsCount, len(localsMap)),
-				attribute.String(AttrLocalsNames, formatMapKeys(localsMap)),
-			)
-		} else {
-			attrs = append(attrs, attribute.Int(AttrLocalsCount, 0))
-		}
-
-		// Feature flags information
-		if baseBlocks.FeatureFlags != nil && !baseBlocks.FeatureFlags.IsNull() {
-			flagsMap := baseBlocks.FeatureFlags.AsValueMap()
-			attrs = append(attrs,
-				attribute.Int(AttrFeatureFlagCount, len(flagsMap)),
-				attribute.String(AttrFeatureFlagNames, formatMapKeys(flagsMap)),
-			)
-		} else {
-			attrs = append(attrs, attribute.Int(AttrFeatureFlagCount, 0))
-		}
-	}
-
-	span.SetAttributes(attrs...)
 }
 
 // TraceParseDependencies wraps dependency parsing with telemetry.
@@ -165,17 +91,7 @@ func TraceParseDependencies(
 	dependencyNames []string,
 	fn func(ctx context.Context) error,
 ) error {
-	attrs := map[string]any{
-		AttrConfigPath:      configPath,
-		AttrSkipOutputs:     skipOutputsResolution,
-		AttrDependencyCount: dependencyCount,
-	}
-
-	if len(dependencyNames) > 0 {
-		attrs[AttrDependencyNames] = strings.Join(dependencyNames, ",")
-	}
-
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, TelemetryOpParseDependencies, attrs, fn)
+	return nil
 }
 
 // TraceParseDependency wraps individual dependency output resolution with telemetry.
@@ -185,10 +101,7 @@ func TraceParseDependency(
 	dependencyPath string,
 	fn func(ctx context.Context) error,
 ) error {
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, TelemetryOpParseDependency, map[string]any{
-		AttrDependencyName: dependencyName,
-		AttrDependencyPath: dependencyPath,
-	}, fn)
+	return nil
 }
 
 // TraceParseConfigDecode wraps config decoding with telemetry.
@@ -197,9 +110,7 @@ func TraceParseConfigDecode(
 	configPath string,
 	fn func(ctx context.Context) error,
 ) error {
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, TelemetryOpParseConfigDecode, map[string]any{
-		AttrConfigPath: configPath,
-	}, fn)
+	return nil
 }
 
 // TraceParseIncludeMerge wraps include merging with telemetry.
@@ -210,16 +121,7 @@ func TraceParseIncludeMerge(
 	includePaths []string,
 	fn func(ctx context.Context) error,
 ) error {
-	attrs := map[string]any{
-		AttrConfigPath:   configPath,
-		AttrIncludeCount: includeCount,
-	}
-
-	if len(includePaths) > 0 {
-		attrs[AttrIncludePaths] = strings.Join(includePaths, ",")
-	}
-
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, TelemetryOpParseIncludeMerge, attrs, fn)
+	return nil
 }
 
 // formatDecodeList converts a slice of PartialDecodeSectionType to a comma-separated string.
