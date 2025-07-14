--- vendor/github.com/hashicorp/aws-sdk-go-base/v2/aws_config.go.orig	2023-12-12 19:24:56 UTC
+++ vendor/github.com/hashicorp/aws-sdk-go-base/v2/aws_config.go
@@ -354,31 +354,12 @@ func commonLoadOptions(ctx context.Context, c *Config)
 		apiOptions = append(apiOptions, awsmiddleware.AddUserAgentKey(v))
 	}
 
-	if !c.SuppressDebugLog {
-		apiOptions = append(apiOptions,
-			func(stack *middleware.Stack) error {
-				return stack.Initialize.Add(&logAttributeExtractor{}, middleware.After)
-			},
-			func(stack *middleware.Stack) error {
-				return stack.Deserialize.Add(&requestResponseLogger{}, middleware.After)
-			},
-		)
-	}
-
 	loadOptions := []func(*config.LoadOptions) error{
 		config.WithRegion(c.Region),
 		config.WithHTTPClient(httpClient),
 		config.WithAPIOptions(apiOptions),
 		config.WithEC2IMDSClientEnableState(c.EC2MetadataServiceEnableState),
 		config.WithLogConfigurationWarnings(true),
-	}
-
-	if !c.SuppressDebugLog {
-		loadOptions = append(
-			loadOptions,
-			config.WithClientLogMode(aws.LogDeprecatedUsage|aws.LogRetries),
-			config.WithLogger(debugLogger{}),
-		)
 	}
 
 	sharedCredentialsFiles, err := c.ResolveSharedCredentialsFiles()
