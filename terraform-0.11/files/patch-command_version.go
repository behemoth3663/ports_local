--- command/version.go.orig	2019-05-16 20:40:50 UTC
+++ command/version.go
@@ -95,26 +95,6 @@ func (c *VersionCommand) Run(args []string) int {
 		}
 	}
 
-	// If we have a version check function, then let's check for
-	// the latest version as well.
-	if c.CheckFunc != nil {
-		// Separate the prior output with a newline
-		c.Ui.Output("")
-
-		// Check the latest version
-		info, err := c.CheckFunc()
-		if err != nil {
-			c.Ui.Error(fmt.Sprintf(
-				"Error checking latest version: %s", err))
-		}
-		if info.Outdated {
-			c.Ui.Output(fmt.Sprintf(
-				"Your version of Terraform is out of date! The latest version\n"+
-					"is %s. You can update by downloading from www.terraform.io/downloads.html",
-				info.Latest))
-		}
-	}
-
 	return 0
 }
 
