--- internal/getter/oci_auth.go.orig	1979-11-29 21:00:00 UTC
+++ internal/getter/oci_auth.go
@@ -15,7 +15,6 @@ import (
 	"sync"
 	"time"
 
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/version"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
@@ -442,10 +441,10 @@ func ociEnvSlice(ctx context.Context, env map[string]s
 // ociEnvSlice renders env as a non-nil KEY=VALUE slice and injects TRACEPARENT from ctx when present.
 func ociEnvSlice(ctx context.Context, env map[string]string) []string {
 	n := len(env)
-	traceParent := telemetry.TraceParentFromContext(ctx, nil)
+	traceParent := ""
 
 	if traceParent != "" {
-		if _, ok := env[telemetry.TraceParentEnv]; !ok {
+		if _, ok := env["TRACEPARENT"]; !ok {
 			n++
 		}
 	}
@@ -453,7 +452,7 @@ func ociEnvSlice(ctx context.Context, env map[string]s
 	out := make([]string, 0, n)
 
 	for _, key := range slices.Sorted(maps.Keys(env)) {
-		if key == telemetry.TraceParentEnv && traceParent != "" {
+		if key == "TRACEPARENT" && traceParent != "" {
 			continue
 		}
 
@@ -461,7 +460,7 @@ func ociEnvSlice(ctx context.Context, env map[string]s
 	}
 
 	if traceParent != "" {
-		out = append(out, telemetry.TraceParentEnv+"="+traceParent)
+		out = append(out, "TRACEPARENT"+"="+traceParent)
 	}
 
 	return out
