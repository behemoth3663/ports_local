--- command/version.go.orig	2021-01-06 19:16:17 UTC
+++ command/version.go
@@ -105,21 +105,6 @@ func (c *VersionCommand) Run(args []string) int {
 		}
 	}
 
-	// If we have a version check function, then let's check for
-	// the latest version as well.
-	if c.CheckFunc != nil {
-		// Check the latest version
-		info, err := c.CheckFunc()
-		if err != nil && !jsonOutput {
-			c.Ui.Error(fmt.Sprintf(
-				"\nError checking latest version: %s", err))
-		}
-		if info.Outdated {
-			outdated = true
-			latest = info.Latest
-		}
-	}
-
 	if jsonOutput {
 		selectionsOutput := make(map[string]string)
 		for providerAddr, lock := range providerLocks {
