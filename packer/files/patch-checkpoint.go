--- checkpoint.go.orig	2023-12-05 19:39:03 UTC
+++ checkpoint.go
@@ -4,14 +4,8 @@ import (
 package main
 
 import (
-	"fmt"
-	"log"
-	"path/filepath"
-
 	"github.com/hashicorp/go-checkpoint"
-	"github.com/hashicorp/packer-plugin-sdk/pathing"
 	"github.com/hashicorp/packer/command"
-	packerVersion "github.com/hashicorp/packer/version"
 )
 
 func init() {
@@ -23,43 +17,7 @@ func runCheckpoint(c *config) {
 // runCheckpoint runs a HashiCorp Checkpoint request. You can read about
 // Checkpoint here: https://github.com/hashicorp/go-checkpoint.
 func runCheckpoint(c *config) {
-	// If the user doesn't want checkpoint at all, then return.
-	if c.DisableCheckpoint {
-		log.Printf("[INFO] Checkpoint disabled. Not running.")
-		checkpointResult <- nil
-		return
-	}
-
-	configDir, err := pathing.ConfigDir()
-	if err != nil {
-		log.Printf("[ERR] Checkpoint setup error: %s", err)
-		checkpointResult <- nil
-		return
-	}
-
-	version := packerVersion.Version
-	if packerVersion.VersionPrerelease != "" {
-		version += fmt.Sprintf("-%s", packerVersion.VersionPrerelease)
-	}
-
-	signaturePath := filepath.Join(configDir, "checkpoint_signature")
-	if c.DisableCheckpointSignature {
-		log.Printf("[INFO] Checkpoint signature disabled")
-		signaturePath = ""
-	}
-
-	resp, err := checkpoint.Check(&checkpoint.CheckParams{
-		Product:       "packer",
-		Version:       version,
-		SignatureFile: signaturePath,
-		CacheFile:     filepath.Join(configDir, "checkpoint_cache"),
-	})
-	if err != nil {
-		log.Printf("[ERR] Checkpoint error: %s", err)
-		resp = nil
-	}
-
-	checkpointResult <- resp
+	checkpointResult <- nil
 }
 
 // commandVersionCheck implements command.VersionCheckFunc and is used
