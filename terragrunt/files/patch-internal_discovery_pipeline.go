--- internal/discovery/pipeline.go.orig	1979-11-29 21:00:00 UTC
+++ internal/discovery/pipeline.go
@@ -9,9 +9,6 @@ import (
 	"strings"
 	"sync"
 
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
-
 	"github.com/gruntwork-io/terragrunt/internal/component"
 	"github.com/gruntwork-io/terragrunt/internal/experiment"
 	"github.com/gruntwork-io/terragrunt/internal/filter"
@@ -164,14 +161,6 @@ func (d *Discovery) runPipeline(
 		return walked{}, err
 	}
 
-	if span := trace.SpanFromContext(ctx); span.IsRecording() {
-		span.SetAttributes(
-			attribute.Int("walked", len(inWalk.discovered)+len(inWalk.candidates)),
-			attribute.Int("in_worktrees", len(inWorktrees.discovered)+len(inWorktrees.candidates)),
-			attribute.Int("parsed", len(collected.parsed)+len(collected.parseErrs)),
-		)
-	}
-
 	results, parseErrs, err := collected.results(d, inWalk)
 	if err != nil {
 		return walked{}, err
@@ -272,13 +261,6 @@ func (p *pipeline) walk(ctx context.Context, l log.Log
 	roots, reasons := d.classifier.WalkRoots(d.discoveryContext.WorkingDir)
 
 	l.Debugf("Discovery: pruned walk from %d root(s): %s", len(roots), strings.Join(reasons, "; "))
-
-	if span := trace.SpanFromContext(ctx); span.IsRecording() {
-		span.SetAttributes(
-			attribute.Int("walk_roots", len(roots)),
-			attribute.StringSlice("walk_reasons", reasons),
-		)
-	}
 
 	walkRoot, err := walkRootUnder(v.FS, d.discoveryContext.WorkingDir, d.walkRoot)
 	if err != nil {
