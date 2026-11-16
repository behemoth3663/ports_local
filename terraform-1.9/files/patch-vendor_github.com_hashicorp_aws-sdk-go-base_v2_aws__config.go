--- vendor/github.com/hashicorp/aws-sdk-go-base/v2/aws_config.go.orig	2024-08-27 14:10:44 UTC
+++ vendor/github.com/hashicorp/aws-sdk-go-base/v2/aws_config.go
@@ -384,31 +384,12 @@ func commonLoadOptions(ctx context.Context, c *Config)
 		apiOptions = append(apiOptions, withUserAgentAppender(v))
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
