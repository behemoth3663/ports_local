--- go.mod.orig	2023-04-21 17:55:47 UTC
+++ go.mod
@@ -5,13 +5,28 @@ require (
 	github.com/aws/aws-sdk-go-v2 v1.17.8
 	github.com/aws/aws-sdk-go-v2/config v1.18.21
 	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v1.20.9
-	github.com/davecgh/go-spew v1.1.1 // indirect
 	github.com/fatih/color v1.15.0
 	github.com/jmespath/go-jmespath v0.4.0
-	github.com/mattn/go-colorable v0.1.13 // indirect
 	github.com/stretchr/testify v1.8.2
+)
+
+require (
+	github.com/aws/aws-sdk-go-v2/credentials v1.13.20 // indirect
+	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.2 // indirect
+	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.32 // indirect
+	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.26 // indirect
+	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.33 // indirect
+	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.26 // indirect
+	github.com/aws/aws-sdk-go-v2/service/sso v1.12.8 // indirect
+	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.8 // indirect
+	github.com/aws/aws-sdk-go-v2/service/sts v1.18.9 // indirect
+	github.com/aws/smithy-go v1.13.5 // indirect
+	github.com/davecgh/go-spew v1.1.1 // indirect
+	github.com/mattn/go-colorable v0.1.13 // indirect
+	github.com/mattn/go-isatty v0.0.17 // indirect
+	github.com/pmezard/go-difflib v1.0.0 // indirect
 	golang.org/x/sys v0.7.0 // indirect
 	gopkg.in/yaml.v3 v3.0.1 // indirect
 )
 
-go 1.13
+go 1.17
