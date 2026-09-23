diff --git a/internal/cas/gitstore.go b/internal/cas/gitstore.go
index 27a7d12f8..01113be3a 100644
--- internal/cas/gitstore.go.orig
+++ internal/cas/gitstore.go
@@ -2,16 +2,12 @@ package cas
 
 import (
 	"context"
+	"errors"
 	"fmt"
 	"path/filepath"
 	"strings"
 	"time"
 
-	"errors"
-
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
-
 	"github.com/gruntwork-io/terragrunt/internal/git"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
@@ -424,14 +420,8 @@ func recordCommitFetchPath(
 	url, ref string,
 	path commitFetchPath,
 ) {
+	_ = ctx
 	l.Debugf("git store: fetched %s from %s via %s", ref, RedactURL(url), path)
-
-	span := trace.SpanFromContext(ctx)
-	if !span.IsRecording() {
-		return
-	}
-
-	span.SetAttributes(attribute.String("git_store_commit_fetch", string(path)))
 }
 
 // repoSession bundles the locked repo handle, the runner pointed at it,
