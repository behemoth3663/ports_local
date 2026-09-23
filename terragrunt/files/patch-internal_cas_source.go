--- internal/cas/source.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cas/source.go
@@ -12,11 +12,7 @@ import (
 	"path/filepath"
 	"time"
 
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
-
 	"github.com/gruntwork-io/terragrunt/internal/git"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
@@ -176,20 +172,12 @@ func (c *CAS) FetchSource(
 
 	maps.Copy(attrs, src.Attrs)
 
-	tlm := telemetry.TelemeterFromContext(ctx)
-
-	return tlm.Collect(
-		ctx,
-		l,
-		"cas_fetch_source",
-		attrs,
-		func(childCtx context.Context, l log.Logger) error {
-			suggestedKey, err := c.probeSource(childCtx, l, v, src)
+			suggestedKey, err := c.probeSource(ctx, l, v, src)
 			if err != nil {
 				return err
 			}
 
-			err = c.fetchAndLink(childCtx, l, v, opts, src, suggestedKey, IngestCached)
+			err = c.fetchAndLink(ctx, l, v, opts, src, suggestedKey, IngestCached)
 
 			var missing *MissingObjectError
 			if !errors.As(err, &missing) {
@@ -205,7 +193,7 @@ func (c *CAS) FetchSource(
 				missing.Hash,
 				RedactURL(src.URL),
 			)
-			RecordFallback(childCtx, l, FallbackReasonStoreRepair, map[string]any{
+			RecordFallback(ctx, l, FallbackReasonStoreRepair, map[string]any{
 				"url":    RedactURL(src.URL),
 				"scheme": src.Scheme,
 				"hash":   missing.Hash,
@@ -215,14 +203,12 @@ func (c *CAS) FetchSource(
 			// producing the object cannot produce it on a third pass
 			// either, so the second failure is the one the caller sees.
 			if err := c.fetchAndLink(
-				childCtx, l, v, opts, src, suggestedKey, IngestRepair,
+				ctx, l, v, opts, src, suggestedKey, IngestRepair,
 			); err != nil {
 				return fmt.Errorf("re-ingest %s: %w", RedactURL(src.URL), err)
 			}
 
 			return nil
-		},
-	)
 }
 
 // fetchAndLink ingests src unless the store already holds the tree
@@ -500,12 +486,8 @@ func recordFetchOutcome(ctx context.Context, cacheHit 
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
@@ -563,13 +545,7 @@ func (s *probeOriginSink) stamp(ctx context.Context) {
 	if s.origin == "" {
 		return
 	}
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
