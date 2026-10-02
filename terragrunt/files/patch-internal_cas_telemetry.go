--- internal/cas/telemetry.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cas/telemetry.go
@@ -4,7 +4,6 @@ import (
 	"context"
 	"maps"
 
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 )
 
@@ -63,11 +62,5 @@ func RecordFallback(
 
 	maps.Copy(all, attrs)
 
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "cas_fallback", all,
-		func(context.Context, log.Logger) error {
-			return nil
-		})
-	if err != nil {
-		l.Debugf("cas: failed to record fallback telemetry: %v", err)
-	}
+	_ = all
 }
