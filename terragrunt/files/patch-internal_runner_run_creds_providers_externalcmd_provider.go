--- internal/runner/run/creds/providers/externalcmd/provider.go.orig	1979-11-29 21:00:00 UTC
+++ internal/runner/run/creds/providers/externalcmd/provider.go
@@ -13,7 +13,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds/providers"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds/providers/amazonsts"
 	"github.com/gruntwork-io/terragrunt/internal/shell"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"github.com/mattn/go-shellwords"
@@ -54,20 +53,7 @@ func (provider *Provider) GetCredentials(
 		return nil, nil
 	}
 
-	var creds *providers.Credentials
-
-	err := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "obtain_creds", map[string]any{
-		"auth_provider_cmd": provider.authProviderCmd,
-		"provider":          "external_cmd",
-	}, func(credsCtx context.Context, l log.Logger) error {
-		var fetchErr error
-
-		creds, fetchErr = provider.fetchCredentials(credsCtx, l, v)
-
-		return fetchErr
-	})
-
-	return creds, err
+	return provider.fetchCredentials(ctx, l, v)
 }
 
 // fetchCredentials runs the configured auth-provider command and decodes its JSON
