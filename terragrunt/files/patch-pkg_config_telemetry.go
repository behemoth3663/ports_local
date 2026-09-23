diff --git a/pkg/config/telemetry.go b/pkg/config/telemetry.go
index 092202bba..6604c17f8 100644
--- pkg/config/telemetry.go.orig
+++ pkg/config/telemetry.go
@@ -5,7 +5,6 @@ import (
 	"context"
 	"strings"
 
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 )
 
@@ -57,8 +56,7 @@ func TraceParseConfigFile(
 		attrs[AttrIncludeChildPath] = includeFromChild.Path
 	}
 
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, TelemetryOpParseConfigFile, attrs, fn)
+	return fn(ctx, l)
 }
 
 // TraceParseDependencies wraps dependency parsing with telemetry.
@@ -81,8 +79,7 @@ func TraceParseDependencies(
 		attrs[AttrDependencyNames] = strings.Join(dependencyNames, ",")
 	}
 
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, TelemetryOpParseDependencies, attrs, fn)
+	return fn(ctx, l)
 }
 
 // formatDecodeList converts a slice of PartialDecodeSectionType to a comma-separated string.
