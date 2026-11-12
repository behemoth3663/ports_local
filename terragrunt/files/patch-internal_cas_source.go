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
@@ -183,28 +179,20 @@ func (c *CAS) FetchSource(
 
 	maps.Copy(attrs, src.Attrs)
 
-	tlm := telemetry.TelemeterFromContext(ctx)
-
-	err = tlm.Collect(
-		ctx,
-		l,
-		"cas_fetch_source",
-		attrs,
-		func(childCtx context.Context, l log.Logger) error {
-			suggestedKey, err := c.probeSource(childCtx, l, v, src)
+			suggestedKey, err := c.probeSource(ctx, l, v, src)
 			if err != nil {
-				return err
+				return "", err
 			}
 
-			linkedTreeKey, err = c.fetchAndLink(childCtx, l, v, opts, src, suggestedKey, IngestCached)
+			linkedTreeKey, err = c.fetchAndLink(ctx, l, v, opts, src, suggestedKey, IngestCached)
 
 			var missing *MissingObjectError
 			if !errors.As(err, &missing) {
-				return err
+				return "", err
 			}
 
 			if c.probeMode == ProbeModeOffline {
-				return &OfflineRepairError{Missing: missing, Source: src.URL}
+				return "", &OfflineRepairError{Missing: missing, Source: src.URL}
 			}
 
 			l.Warnf(
@@ -212,7 +200,7 @@ func (c *CAS) FetchSource(
 				missing.Hash,
 				src.URL,
 			)
-			RecordFallback(childCtx, l, FallbackReasonStoreRepair, map[string]any{
+			RecordFallback(ctx, l, FallbackReasonStoreRepair, map[string]any{
 				"url":    src.URL,
 				"scheme": src.Scheme,
 				"hash":   missing.Hash,
@@ -221,18 +209,11 @@ func (c *CAS) FetchSource(
 			// One attempt. A source that answered the re-ingest without
 			// producing the object cannot produce it on a third pass
 			// either, so the second failure is the one the caller sees.
-			linkedTreeKey, err = c.fetchAndLink(childCtx, l, v, opts, src, suggestedKey, IngestRepair)
+			linkedTreeKey, err = c.fetchAndLink(ctx, l, v, opts, src, suggestedKey, IngestRepair)
 			if err != nil {
-				return fmt.Errorf("re-ingest %s: %w", src.URL, err)
+				return "", fmt.Errorf("re-ingest %s: %w", src.URL, err)
 			}
 
-			return nil
-		},
-	)
-	if err != nil {
-		return "", err
-	}
-
 	return linkedTreeKey, nil
 }
 
@@ -517,15 +498,10 @@ func (c *CAS) recordSourceProbe(l log.Logger, v *venv.
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
@@ -577,19 +553,10 @@ func recordProbeOrigin(ctx context.Context, origin pro
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
