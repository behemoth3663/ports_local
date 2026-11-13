--- vendor/github.com/anthropics/anthropic-sdk-go/bedrock/bedrock.go.orig	2026-11-13 20:40:39 UTC
+++ vendor/github.com/anthropics/anthropic-sdk-go/bedrock/bedrock.go
@@ -28,7 +28,6 @@ import (
 
 	"github.com/anthropics/anthropic-sdk-go/internal/awsauth"
 	"github.com/anthropics/anthropic-sdk-go/internal/requestconfig"
-	"github.com/anthropics/anthropic-sdk-go/internal/telemetry"
 	"github.com/anthropics/anthropic-sdk-go/option"
 )
 
@@ -210,7 +209,6 @@ var withPlatform = requestconfig.RequestOptionFunc(fun
 // withPlatform names Amazon Bedrock as the gen_ai.provider.name of the
 // OpenTelemetry spans, and this package in their instrumentation scope.
 var withPlatform = requestconfig.RequestOptionFunc(func(rc *requestconfig.RequestConfig) error {
-	rc.OpenTelemetry.ProviderName, rc.OpenTelemetry.Platform = telemetry.ProviderBedrock, "bedrock"
 	return nil
 })
 
