--- internal/cli/commands/scaffold/scaffold.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cli/commands/scaffold/scaffold.go
@@ -2,6 +2,7 @@ import (
 
 import (
 	"context"
+	"errors"
 	"fmt"
 	"io/fs"
 	"net/url"
@@ -15,7 +16,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/configbridge"
 	"github.com/gruntwork-io/terragrunt/internal/services/catalog/component"
 	"github.com/gruntwork-io/terragrunt/internal/shell"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
 	"github.com/gruntwork-io/terragrunt/internal/view/tui/form"
@@ -26,8 +26,6 @@ import (
 
 	"github.com/gruntwork-io/terragrunt/internal/util"
 
-	"errors"
-
 	"github.com/gruntwork-io/boilerplate/manifest"
 	boilerplateoptions "github.com/gruntwork-io/boilerplate/options"
 	"github.com/gruntwork-io/boilerplate/templates"
@@ -262,17 +260,8 @@ func Prepare(
 
 	l.Debugf("Scaffolding a new Terragrunt module %s to %s", resolvedURL, outputDir)
 
-	if err := telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "scaffold_get_module", map[string]any{
-			"module_url": resolvedURL,
-		}, func(ctx context.Context, l log.Logger) error {
-			if _, getErr := getter.GetAny(ctx, l, v, tempDir, resolvedURL); getErr != nil {
-				return fmt.Errorf("downloading scaffold module from %s: %w", resolvedURL, getErr)
-			}
-
-			return nil
-		}); err != nil {
-		return nil, err
+	if _, err := getter.GetAny(ctx, l, v, tempDir, resolvedURL); err != nil {
+		return nil, fmt.Errorf("downloading scaffold module from %s: %w", resolvedURL, err)
 	}
 
 	markers, err := component.Inspect(v.FS, tempDir)
@@ -679,21 +668,12 @@ func downloadTemplate(
 	l.Debugf("Downloading template from %s into %s", baseURL.String(), templateDir)
 	// Downloading baseURL to support boilerplate dependencies and partials.
 	// Go-getter discards all but specified folder if one is provided.
-	if err := telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "scaffold_get_template", map[string]any{
-			"template_url": baseURL.String(),
-		}, func(ctx context.Context, l log.Logger) error {
-			if _, getErr := getter.GetAny(ctx, l, v, templateDir, baseURL.String()); getErr != nil {
-				return fmt.Errorf(
-					"downloading scaffold template from %s: %w",
-					baseURL.String(),
-					getErr,
-				)
-			}
-
-			return nil
-		}); err != nil {
-		return "", err
+	if _, err := getter.GetAny(ctx, l, v, templateDir, baseURL.String()); err != nil {
+		return "", fmt.Errorf(
+			"downloading scaffold template from %s: %w",
+			baseURL.String(),
+			err,
+		)
 	}
 
 	// Add subfolder to templateDir if provided, as scaffold needs path to boilerplate.yml file
