--- cmd/tfsec/main.go.orig	2023-09-07 14:36:57 UTC
+++ cmd/tfsec/main.go
@@ -8,21 +8,7 @@ import (
 	"github.com/aquasecurity/tfsec/internal/app/tfsec/cmd"
 )
 
-const transitionMsg = `
-======================================================
-tfsec is joining the Trivy family
-
-tfsec will continue to remain available 
-for the time being, although our engineering 
-attention will be directed at Trivy going forward.
-
-You can read more here: 
-https://github.com/aquasecurity/tfsec/discussions/1994
-======================================================
-`
-
 func main() {
-	fmt.Print(transitionMsg)
 	if err := cmd.Root().Execute(); err != nil {
 		if err.Error() != "" {
 			fmt.Printf("Error: %s\n", err)
