--- command/version.go.orig	2020-07-22 15:42:34 UTC
+++ command/version.go
@@ -95,24 +95,6 @@ func (c *VersionCommand) Run(args []string) int {
 		}
 	}
 
-	// If we have a version check function, then let's check for
-	// the latest version as well.
-	if c.CheckFunc != nil {
-
-		// Check the latest version
-		info, err := c.CheckFunc()
-		if err != nil {
-			c.Ui.Error(fmt.Sprintf(
-				"\nError checking latest version: %s", err))
-		}
-		if info.Outdated {
-			c.Ui.Output(fmt.Sprintf(
-				"\nYour version of Terraform is out of date! The latest version\n"+
-					"is %s. You can update by downloading from https://www.terraform.io/downloads.html",
-				info.Latest))
-		}
-	}
-
 	return 0
 }
