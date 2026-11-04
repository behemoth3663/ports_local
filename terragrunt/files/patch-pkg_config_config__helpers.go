--- pkg/config/config_helpers.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/config/config_helpers.go
@@ -38,7 +38,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/retry"
 	"github.com/gruntwork-io/terragrunt/internal/shell"
 	"github.com/gruntwork-io/terragrunt/internal/strict/controls"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tf"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/vendored/opentofu/patch"
@@ -1144,18 +1143,8 @@ func getAWSField(
 
 	var result string
 
-	err = telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "config_get_aws_field", map[string]any{
-			"config_path": pctx.TerragruntConfigPath,
-			"role_arn":    pctx.IAMRoleOptions.RoleARN,
-		}, func(ctx context.Context, l log.Logger) error {
-			var fetchErr error
+	result, err = fetchFn(ctx, &awsConfig)
 
-			result, fetchErr = fetchFn(ctx, &awsConfig)
-
-			return fetchErr
-		})
-
 	return result, err
 }
 
@@ -1504,17 +1493,7 @@ func SopsDecryptFileWithDecrypter(
 
 	var rawData []byte
 
-	err := telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "config_sops_decrypt", map[string]any{
-			"path":   path,
-			"format": format,
-		}, func(ctx context.Context, l log.Logger) error {
-			var decryptErr error
-
-			rawData, decryptErr = d.DecryptFile(v.Env, path, format)
-
-			return decryptErr
-		})
+	rawData, err := d.DecryptFile(v.Env, path, format)
 	if err != nil {
 		return "", err
 	}
