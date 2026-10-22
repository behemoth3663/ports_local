--- internal/cas/telemetry.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cas/telemetry.go
@@ -2,9 +2,7 @@ import (
 
 import (
 	"context"
-	"maps"
 
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 )
 
@@ -59,15 +57,5 @@ func RecordFallback(
 	reason FallbackReason,
 	attrs map[string]any,
 ) {
-	all := map[string]any{"reason": string(reason)}
-
-	maps.Copy(all, attrs)
-
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "cas_fallback", all,
-		func(context.Context, log.Logger) error {
-			return nil
-		})
-	if err != nil {
-		l.Debugf("cas: failed to record fallback telemetry: %v", err)
-	}
+	_, _, _, _ = ctx, l, reason, attrs
 }
