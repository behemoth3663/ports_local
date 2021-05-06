--- command/version.go.orig	2020-09-30 18:07:32 UTC
+++ command/version.go
@@ -103,20 +103,6 @@ func (c *VersionCommand) Run(args []string) int {
 			pluginVersions = append(pluginVersions, fmt.Sprintf("+ provider %s v%s", providerAddr, version))
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
 
 	if jsonOutput {
 		selectionsOutput := make(map[string]string)
