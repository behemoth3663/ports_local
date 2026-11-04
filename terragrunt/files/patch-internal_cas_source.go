--- internal/cas/source.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cas/source.go
@@ -12,12 +12,8 @@ import (
 	"path/filepath"
 	"time"
 
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
-
 	"github.com/gruntwork-io/terragrunt/internal/git"
 	"github.com/gruntwork-io/terragrunt/internal/redact"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
@@ -177,14 +173,8 @@ func (c *CAS) FetchSource(
 
 	maps.Copy(attrs, src.Attrs)
 
-	tlm := telemetry.TelemeterFromContext(ctx)
-
-	return tlm.Collect(
-		ctx,
-		l,
-		"cas_fetch_source",
-		attrs,
-		func(childCtx context.Context, l log.Logger) error {
+	_ = attrs
+	return func(childCtx context.Context, l log.Logger) error {
 			suggestedKey, err := c.probeSource(childCtx, l, v, src)
 			if err != nil {
 				return err
@@ -222,8 +212,7 @@ func (c *CAS) FetchSource(
 			}
 
 			return nil
-		},
-	)
+		}(ctx, l)
 }
 
 // fetchAndLink ingests src unless the store already holds the tree
@@ -498,15 +487,10 @@ func (c *CAS) recordSourceProbe(l log.Logger, v *venv.
 	}
 }
 
-// recordFetchOutcome stamps cache_hit on the active cas_fetch_source span
-// so dashboards can distinguish probe short-circuits from network fetches.
+// recordFetchOutcome records whether the fetch was served from cache.
 func recordFetchOutcome(ctx context.Context, cacheHit bool) {
-	span := trace.SpanFromContext(ctx)
-	if !span.IsRecording() {
-		return
-	}
-
-	span.SetAttributes(attribute.Bool("cache_hit", cacheHit))
+	_ = ctx
+	_ = cacheHit
 }
 
 // probeOrigin names where a probe answer came from. It travels as the
@@ -558,19 +542,10 @@ func recordProbeOrigin(ctx context.Context, origin pro
 	sink.origin = origin
 }
 
-// stamp puts the reported origin on the active cas_fetch_source span so
-// dashboards can tell cached probes from network ones.
+// stamp records the reported origin.
 func (s *probeOriginSink) stamp(ctx context.Context) {
-	if s.origin == "" {
-		return
-	}
-
-	span := trace.SpanFromContext(ctx)
-	if !span.IsRecording() {
-		return
-	}
-
-	span.SetAttributes(attribute.String("probe_origin", string(s.origin)))
+	_ = ctx
+	_ = s.origin
 }
 
 // linkStoredTree materializes the tree at key into opts.Dir, then the
