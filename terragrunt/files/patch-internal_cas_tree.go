--- internal/cas/tree.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cas/tree.go
@@ -13,13 +13,10 @@ import (
 	"sync/atomic"
 	"unicode"
 
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 	"golang.org/x/sync/errgroup"
 	"golang.org/x/text/unicode/norm"
 
 	"github.com/gruntwork-io/terragrunt/internal/git"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
@@ -182,16 +179,9 @@ func LinkTree(
 		mutable:     o.mutable,
 	}
 
-	return telemetry.TelemeterFromContext(ctx).Collect(ctx, nil, "cas_link_tree", map[string]any{
-		"path": targetDir,
-		"mode": mode.String(),
-	}, telemetry.WithoutLogger(func(childCtx context.Context) error {
-		err := linkTree(l, v, linker, t, targetDir, o.fsWorkers)
-
-		linker.report(childCtx)
-
-		return err
-	}))
+	err := linkTree(l, v, linker, t, targetDir, o.fsWorkers)
+	linker.report(ctx)
+	return err
 }
 
 // linkTree materializes t and everything nested below it, one level of the
@@ -706,11 +696,6 @@ func (tl *treeLinker) report(ctx context.Context) {
 // report sets the per-mode file counts, the bytes copied, and the fallback on
 // the span in ctx.
 func (tl *treeLinker) report(ctx context.Context) {
-	span := trace.SpanFromContext(ctx)
-	if !span.IsRecording() {
-		return
-	}
-
 	counts := map[LinkMode]int64{
 		LinkModeHardlink: tl.linked.Load(),
 		LinkModeClone:    tl.cloned.Load(),
@@ -726,14 +711,8 @@ func (tl *treeLinker) report(ctx context.Context) {
 	if counts[tl.mode] < total {
 		fallback = linkFallbackByMode[tl.mode]
 	}
-
-	span.SetAttributes(
-		attribute.Int64("files_linked", counts[LinkModeHardlink]),
-		attribute.Int64("files_cloned", counts[LinkModeClone]),
-		attribute.Int64("files_copied", counts[LinkModeCopy]),
-		attribute.Int64("bytes_copied", tl.bytesCopied.Load()),
-		attribute.String("fallback", string(fallback)),
-	)
+	_ = ctx
+	_ = fallback
 }
 
 // gitFilePerm extracts the unix permission bits from a git tree entry mode
