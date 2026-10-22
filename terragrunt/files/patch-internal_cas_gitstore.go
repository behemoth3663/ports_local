--- internal/cas/gitstore.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cas/gitstore.go
@@ -9,9 +9,6 @@ import (
 
 	"errors"
 
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
-
 	"github.com/gruntwork-io/terragrunt/internal/git"
 	"github.com/gruntwork-io/terragrunt/internal/redact"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
@@ -442,13 +439,7 @@ func recordCommitFetchPath(
 	path commitFetchPath,
 ) {
 	l.Debugf("git store: fetched %s from %s via %s", ref, u, path)
-
-	span := trace.SpanFromContext(ctx)
-	if !span.IsRecording() {
-		return
-	}
-
-	span.SetAttributes(attribute.String("git_store_commit_fetch", string(path)))
+	_ = ctx
 }
 
 // repoSession bundles the locked repo handle, the runner pointed at it,
