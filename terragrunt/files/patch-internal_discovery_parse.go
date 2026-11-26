--- internal/discovery/parse.go.orig	1979-11-29 21:00:00 UTC
+++ internal/discovery/parse.go
@@ -15,7 +15,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/filter"
 	inthclparse "github.com/gruntwork-io/terragrunt/internal/hclparse"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
@@ -112,7 +111,6 @@ func ensureParsed(
 		// Cache hits are emitted as a counter rather than a span. A graph or
 		// relationship traversal can re-visit the same component many times, and
 		// a span per visit drowns out the parse misses that are the actual cost.
-		telemetry.TelemeterFromContext(ctx).Count(ctx, "discovery_parse_cache_hit", 1)
 		l.Debugf(
 			"Discovery: parse cache hit for %s (phase=%s, depth=%d)",
 			c.Path(), parsePhaseFromContext(ctx), parseDepthFromContext(ctx),
@@ -200,13 +198,6 @@ func parseComponent(
 	phase := parsePhaseFromContext(ctx)
 	depth := parseDepthFromContext(ctx)
 
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "discovery_parse_component", map[string]any{
-			"path":      c.Path(),
-			"phase":     phase,
-			"depth":     depth,
-			"cache_hit": false,
-		}, func(ctx context.Context, l log.Logger) error {
 			l.Debugf("Discovery: parsing %s (phase=%s, depth=%d)", c.Path(), phase, depth)
 
 			parseOpts := opts.Clone()
@@ -334,7 +325,6 @@ func parseComponent(
 			c.SetReading(sanitizeReadFiles(parsingCtx.FilesRead.Paths())...)
 
 			return nil
-		})
 }
 
 // trackStackComponentReads records in pctx.FilesRead the files read by the unit and
