--- pkg/config/dependency_state.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/config/dependency_state.go
@@ -6,7 +6,6 @@ import (
 	"errors"
 	"io"
 
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 )
 
@@ -32,38 +31,26 @@ func readDependencyStateOutputs(
 		encrypted   error
 	)
 
-	err := telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, metric, attrs, func(ctx context.Context, l log.Logger) error {
-			reader, err := open(ctx, l)
-			if err != nil {
-				return err
-			}
+		reader, err := open(ctx, l)
+		if err != nil {
+			return nil, err
+		}
 
-			defer func() {
-				if err := reader.Close(); err != nil {
-					l.Warnf("Failed to close dependency state reader for %s: %v", location, err)
-				}
-			}()
-
-			jsonOutputs, err = stateOutputsJSON(reader, location)
-
-			// An encrypted state is an expected fallback rather than a failed read, so it is
-			// carried out of the callback to keep it off this metric's error counter.
-			if errors.Is(err, ErrDependencyStateEncrypted) {
-				encrypted = err
-
-				return nil
+		defer func() {
+			if err := reader.Close(); err != nil {
+				l.Warnf("Failed to close dependency state reader for %s: %v", location, err)
 			}
+		}()
 
-			return err
-		})
-	if err != nil {
-		return nil, err
-	}
+		jsonOutputs, err = stateOutputsJSON(reader, location)
 
-	if encrypted != nil {
-		return nil, encrypted
-	}
+		// An encrypted state is an expected fallback rather than a failed read, so it is
+		// carried out of the callback to keep it off this metric's error counter.
+		if errors.Is(err, ErrDependencyStateEncrypted) {
+			encrypted = err
+
+			return nil, encrypted
+		}
 
 	return jsonOutputs, nil
 }
