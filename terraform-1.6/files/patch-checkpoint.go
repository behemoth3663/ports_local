--- checkpoint.go.orig	2023-10-04 16:30:07 UTC
+++ checkpoint.go
@@ -5,14 +5,10 @@ package main
 
 import (
 	"context"
-	"fmt"
-	"log"
-	"path/filepath"
 
 	"github.com/hashicorp/go-checkpoint"
 	"github.com/hashicorp/terraform/internal/command"
 	"github.com/hashicorp/terraform/internal/command/cliconfig"
-	"go.opentelemetry.io/otel/codes"
 )
 
 func init() {
@@ -24,50 +20,6 @@ var checkpointResult chan *checkpoint.CheckResponse
 // runCheckpoint runs a HashiCorp Checkpoint request. You can read about
 // Checkpoint here: https://github.com/hashicorp/go-checkpoint.
 func runCheckpoint(ctx context.Context, c *cliconfig.Config) {
-	// If the user doesn't want checkpoint at all, then return.
-	if c.DisableCheckpoint {
-		log.Printf("[INFO] Checkpoint disabled. Not running.")
-		checkpointResult <- nil
-		return
-	}
-
-	ctx, span := tracer.Start(ctx, "HashiCorp Checkpoint")
-	_ = ctx // prevent staticcheck from complaining to avoid a maintenence hazard of having the wrong ctx in scope here
-	defer span.End()
-
-	configDir, err := cliconfig.ConfigDir()
-	if err != nil {
-		log.Printf("[ERR] Checkpoint setup error: %s", err)
-		checkpointResult <- nil
-		return
-	}
-
-	version := Version
-	if VersionPrerelease != "" {
-		version += fmt.Sprintf("-%s", VersionPrerelease)
-	}
-
-	signaturePath := filepath.Join(configDir, "checkpoint_signature")
-	if c.DisableCheckpointSignature {
-		log.Printf("[INFO] Checkpoint signature disabled")
-		signaturePath = ""
-	}
-
-	resp, err := checkpoint.Check(&checkpoint.CheckParams{
-		Product:       "terraform",
-		Version:       version,
-		SignatureFile: signaturePath,
-		CacheFile:     filepath.Join(configDir, "checkpoint_cache"),
-	})
-	if err != nil {
-		log.Printf("[ERR] Checkpoint error: %s", err)
-		span.SetStatus(codes.Error, err.Error())
-		resp = nil
-	} else {
-		span.SetStatus(codes.Ok, "checkpoint request succeeded")
-	}
-
-	checkpointResult <- resp
 }
 
 // commandVersionCheck implements command.VersionCheckFunc and is used
