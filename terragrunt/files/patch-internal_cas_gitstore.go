--- internal/cas/gitstore.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cas/gitstore.go
@@ -2,16 +2,12 @@ import (
 
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
