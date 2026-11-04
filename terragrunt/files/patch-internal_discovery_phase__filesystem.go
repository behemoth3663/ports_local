--- internal/discovery/phase_filesystem.go.orig	1979-11-29 21:00:00 UTC
+++ internal/discovery/phase_filesystem.go
@@ -6,9 +6,6 @@ import (
 	"path/filepath"
 	"strings"
 
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
-
 	"github.com/gruntwork-io/terragrunt/internal/experiment"
 	"github.com/gruntwork-io/terragrunt/internal/filter"
 	"github.com/gruntwork-io/terragrunt/internal/util"
@@ -101,13 +98,6 @@ func (p *FilesystemPhase) Run(
 	roots, reasons := input.Classifier.WalkRoots(discoveryContext.WorkingDir)
 
 	l.Debugf("Discovery: pruned walk from %d root(s): %s", len(roots), strings.Join(reasons, "; "))
-
-	if span := trace.SpanFromContext(ctx); span.IsRecording() {
-		span.SetAttributes(
-			attribute.Int("walk_roots", len(roots)),
-			attribute.StringSlice("walk_reasons", reasons),
-		)
-	}
 
 	walkRoot, err := walkRootUnder(v.FS, discoveryContext.WorkingDir, discovery.walkRoot)
 	if err != nil {
