--- config/telemetry.go.orig	2026-01-09 15:01:11 UTC
+++ config/telemetry.go
@@ -7,8 +7,6 @@ import (
 
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"github.com/gruntwork-io/terragrunt/telemetry"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // Telemetry operation names for config parsing operations.
@@ -106,54 +104,6 @@ func TraceParseBaseBlocksResult(
 	configPath string,
 	baseBlocks *DecodedBaseBlocks,
 ) {
-	span := trace.SpanFromContext(ctx)
-	if span == nil || !span.IsRecording() {
-		return
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
