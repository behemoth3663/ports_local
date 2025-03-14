--- internal/server/references.go.orig	2025-02-24 17:39:34 UTC
+++ internal/server/references.go
@@ -11,17 +11,11 @@ import (
 	"golang.org/x/tools/gopls/internal/golang"
 	"golang.org/x/tools/gopls/internal/label"
 	"golang.org/x/tools/gopls/internal/protocol"
-	"golang.org/x/tools/gopls/internal/telemetry"
 	"golang.org/x/tools/gopls/internal/template"
 	"golang.org/x/tools/internal/event"
 )
 
 func (s *server) References(ctx context.Context, params *protocol.ReferenceParams) (_ []protocol.Location, rerr error) {
-	recordLatency := telemetry.StartLatencyTimer("references")
-	defer func() {
-		recordLatency(ctx, rerr)
-	}()
-
 	ctx, done := event.Start(ctx, "lsp.Server.references", label.URI.Of(params.TextDocument.URI))
 	defer done()
 
