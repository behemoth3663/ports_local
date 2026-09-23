--- pkg/config/config_helpers.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/config/config_helpers.go
@@ -3,6 +3,7 @@ import (
 import (
 	"context"
 	"encoding/json"
+	"errors"
 	"fmt"
 	"io"
 	"maps"
@@ -26,8 +27,6 @@ import (
 	"github.com/zclconf/go-cty/cty/function"
 	"github.com/zclconf/go-cty/cty/gocty"
 
-	"errors"
-
 	"github.com/gruntwork-io/terragrunt/internal/awshelper"
 	"github.com/gruntwork-io/terragrunt/internal/cache"
 	"github.com/gruntwork-io/terragrunt/internal/clihelper"
@@ -38,7 +37,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/retry"
 	"github.com/gruntwork-io/terragrunt/internal/shell"
 	"github.com/gruntwork-io/terragrunt/internal/strict/controls"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tf"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/vsops"
@@ -1040,21 +1038,7 @@ func getAWSField(
 		return "", err
 	}
 
-	var result string
-
-	err = telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "config_get_aws_field", map[string]any{
-			"config_path": pctx.TerragruntConfigPath,
-			"role_arn":    pctx.IAMRoleOptions.RoleARN,
-		}, func(ctx context.Context, l log.Logger) error {
-			var fetchErr error
-
-			result, fetchErr = fetchFn(ctx, &awsConfig)
-
-			return fetchErr
-		})
-
-	return result, err
+	return fetchFn(ctx, &awsConfig)
 }
 
 func getAWSAccountAlias(ctx context.Context, pctx *ParsingContext, l log.Logger) (string, error) {
@@ -1308,7 +1292,7 @@ func getModulePathFromSourceURL(sourceURL string) (str
 func getModulePathFromSourceURL(sourceURL string) (string, error) {
 	// Regexp for module name extraction. It assumes that the query string has already been stripped off.
 	// Then we simply capture anything after the last slash, and before `.` or end of string.
-	var moduleNameRegexp = regexp.MustCompile(`(?:.+/)(.+?)(?:\.|$)`)
+	moduleNameRegexp := regexp.MustCompile(`(?:.+/)(.+?)(?:\.|$)`)
 
 	// strip off the query string if present
 	sourceURL = strings.Split(sourceURL, "?")[0]
@@ -1389,19 +1373,7 @@ func SopsDecryptFileWithDecrypter(
 
 	l.Debugf("sops decrypt: decrypting %s (format=%s)", path, format)
 
-	var rawData []byte
-
-	err := telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "config_sops_decrypt", map[string]any{
-			"path":   path,
-			"format": format,
-		}, func(ctx context.Context, l log.Logger) error {
-			var decryptErr error
-
-			rawData, decryptErr = d.DecryptFile(pctx.Venv.Env, path, format)
-
-			return decryptErr
-		})
+	rawData, err := d.DecryptFile(pctx.Venv.Env, path, format)
 	if err != nil {
 		return "", err
 	}
