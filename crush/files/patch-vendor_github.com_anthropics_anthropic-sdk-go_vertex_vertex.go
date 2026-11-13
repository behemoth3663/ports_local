--- vendor/github.com/anthropics/anthropic-sdk-go/vertex/vertex.go.orig	2026-11-13 20:41:32 UTC
+++ vendor/github.com/anthropics/anthropic-sdk-go/vertex/vertex.go
@@ -15,7 +15,6 @@ import (
 	"google.golang.org/api/transport"
 
 	"github.com/anthropics/anthropic-sdk-go/internal/requestconfig"
-	"github.com/anthropics/anthropic-sdk-go/internal/telemetry"
 	sdkoption "github.com/anthropics/anthropic-sdk-go/option"
 	"github.com/tidwall/gjson"
 	"github.com/tidwall/sjson"
@@ -113,7 +112,6 @@ func WithCredentials(ctx context.Context, region strin
 	}
 
 	return sdkoption.Join(sdkoption.WithoutEnvironmentDefaults(), requestconfig.RequestOptionFunc(func(rc *requestconfig.RequestConfig) error {
-		rc.OpenTelemetry.ProviderName, rc.OpenTelemetry.Platform = telemetry.ProviderVertex, "vertex"
 		client := defaultClient
 		if rc.HTTPClient != nil && rc.HTTPClient != http.DefaultClient {
 			client = authorizeClient(rc.HTTPClient, creds)
