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
@@ -177,53 +173,44 @@ func (c *CAS) FetchSource(
 
 	maps.Copy(attrs, src.Attrs)
 
-	tlm := telemetry.TelemeterFromContext(ctx)
+	_ = attrs
+	suggestedKey, err := c.probeSource(ctx, l, v, src)
+	if err != nil {
+		return err
+	}
 
-	return tlm.Collect(
-		ctx,
-		l,
-		"cas_fetch_source",
-		attrs,
-		func(childCtx context.Context, l log.Logger) error {
-			suggestedKey, err := c.probeSource(childCtx, l, v, src)
-			if err != nil {
-				return err
-			}
+	err = c.fetchAndLink(ctx, l, v, opts, src, suggestedKey, IngestCached)
 
-			err = c.fetchAndLink(childCtx, l, v, opts, src, suggestedKey, IngestCached)
+	var missing *MissingObjectError
+	if !errors.As(err, &missing) {
+		return err
+	}
 
-			var missing *MissingObjectError
-			if !errors.As(err, &missing) {
-				return err
-			}
+	if c.probeMode == ProbeModeOffline {
+		return &OfflineRepairError{Missing: missing, Source: src.URL}
+	}
 
-			if c.probeMode == ProbeModeOffline {
-				return &OfflineRepairError{Missing: missing, Source: src.URL}
-			}
+	l.Warnf(
+		"cas: store is missing object %s, re-ingesting %s to restore it",
+		missing.Hash,
+		src.URL,
+	)
+	RecordFallback(ctx, l, FallbackReasonStoreRepair, map[string]any{
+		"url":    src.URL,
+		"scheme": src.Scheme,
+		"hash":   missing.Hash,
+	})
 
-			l.Warnf(
-				"cas: store is missing object %s, re-ingesting %s to restore it",
-				missing.Hash,
-				src.URL,
-			)
-			RecordFallback(childCtx, l, FallbackReasonStoreRepair, map[string]any{
-				"url":    src.URL,
-				"scheme": src.Scheme,
-				"hash":   missing.Hash,
-			})
+	// One attempt. A source that answered the re-ingest without
+	// producing the object cannot produce it on a third pass
+	// either, so the second failure is the one the caller sees.
+	if err := c.fetchAndLink(
+		ctx, l, v, opts, src, suggestedKey, IngestRepair,
+	); err != nil {
+		return fmt.Errorf("re-ingest %s: %w", src.URL, err)
+	}
 
-			// One attempt. A source that answered the re-ingest without
-			// producing the object cannot produce it on a third pass
-			// either, so the second failure is the one the caller sees.
-			if err := c.fetchAndLink(
-				childCtx, l, v, opts, src, suggestedKey, IngestRepair,
-			); err != nil {
-				return fmt.Errorf("re-ingest %s: %w", src.URL, err)
-			}
-
-			return nil
-		},
-	)
+	return nil
 }
 
 // fetchAndLink ingests src unless the store already holds the tree
@@ -501,12 +488,8 @@ func recordFetchOutcome(ctx context.Context, cacheHit 
 // recordFetchOutcome stamps cache_hit on the active cas_fetch_source span
 // so dashboards can distinguish probe short-circuits from network fetches.
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
@@ -561,16 +544,7 @@ func (s *probeOriginSink) stamp(ctx context.Context) {
 // stamp puts the reported origin on the active cas_fetch_source span so
 // dashboards can tell cached probes from network ones.
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
 }
 
 // linkStoredTree materializes the tree at key into opts.Dir, then the
