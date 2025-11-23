--- checkpoint.go.orig	2024-10-16 12:28:59 UTC
+++ checkpoint.go
@@ -12,7 +12,6 @@ import (
 	"github.com/hashicorp/go-checkpoint"
 	"github.com/hashicorp/terraform/internal/command"
 	"github.com/hashicorp/terraform/internal/command/cliconfig"
-	"go.opentelemetry.io/otel/codes"
 )
 
 func init() {
@@ -31,10 +30,6 @@ func runCheckpoint(ctx context.Context, c *cliconfig.C
 		return
 	}
 
-	ctx, span := tracer.Start(ctx, "HashiCorp Checkpoint")
-	_ = ctx // prevent staticcheck from complaining to avoid a maintenence hazard of having the wrong ctx in scope here
-	defer span.End()
-
 	configDir, err := cliconfig.ConfigDir()
 	if err != nil {
 		log.Printf("[ERR] Checkpoint setup error: %s", err)
@@ -61,10 +56,7 @@ func runCheckpoint(ctx context.Context, c *cliconfig.C
 	})
 	if err != nil {
 		log.Printf("[ERR] Checkpoint error: %s", err)
-		span.SetStatus(codes.Error, err.Error())
 		resp = nil
-	} else {
-		span.SetStatus(codes.Ok, "checkpoint request succeeded")
 	}
 
 	checkpointResult <- resp
