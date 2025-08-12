--- internal/server/command.go.orig	2025-07-28 18:28:48 UTC
+++ internal/server/command.go
@@ -25,7 +25,6 @@ import (
 
 	"github.com/fatih/gomodifytags/modifytags"
 	"golang.org/x/mod/modfile"
-	"golang.org/x/telemetry/counter"
 	"golang.org/x/tools/go/ast/astutil"
 	"golang.org/x/tools/gopls/internal/cache"
 	"golang.org/x/tools/gopls/internal/cache/metadata"
@@ -279,7 +278,6 @@ func (*commandHandler) AddTelemetryCounters(_ context.
 		if n == "" || v < 0 {
 			continue
 		}
-		counter.Add("fwd/"+n, v)
 	}
 	return nil
 }
@@ -1604,7 +1602,6 @@ func (c *commandHandler) ChangeSignature(ctx context.C
 }
 
 func (c *commandHandler) ChangeSignature(ctx context.Context, args command.ChangeSignatureArgs) (*protocol.WorkspaceEdit, error) {
-	countChangeSignature.Inc()
 	var result *protocol.WorkspaceEdit
 	err := c.run(ctx, commandConfig{
 		forURI: args.Location.URI,
