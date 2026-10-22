--- internal/getter/oci_auth.go.orig	1979-11-29 21:00:00 UTC
+++ internal/getter/oci_auth.go
@@ -15,7 +15,6 @@ import (
 	"sync"
 	"time"
 
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/version"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
@@ -441,27 +440,11 @@ func ociEnvSlice(ctx context.Context, env map[string]s
 
 // ociEnvSlice renders env as a non-nil KEY=VALUE slice and injects TRACEPARENT from ctx when present.
 func ociEnvSlice(ctx context.Context, env map[string]string) []string {
-	n := len(env)
-	traceParent := telemetry.TraceParentFromContext(ctx, nil)
+	_ = ctx
+	out := make([]string, 0, len(env))
 
-	if traceParent != "" {
-		if _, ok := env[telemetry.TraceParentEnv]; !ok {
-			n++
-		}
-	}
-
-	out := make([]string, 0, n)
-
 	for _, key := range slices.Sorted(maps.Keys(env)) {
-		if key == telemetry.TraceParentEnv && traceParent != "" {
-			continue
-		}
-
 		out = append(out, key+"="+env[key])
-	}
-
-	if traceParent != "" {
-		out = append(out, telemetry.TraceParentEnv+"="+traceParent)
 	}
 
 	return out
