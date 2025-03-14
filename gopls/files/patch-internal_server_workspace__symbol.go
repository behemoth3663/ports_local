--- internal/server/workspace_symbol.go.orig	2025-02-24 17:39:34 UTC
+++ internal/server/workspace_symbol.go
@@ -10,16 +10,10 @@ import (
 	"golang.org/x/tools/gopls/internal/cache"
 	"golang.org/x/tools/gopls/internal/golang"
 	"golang.org/x/tools/gopls/internal/protocol"
-	"golang.org/x/tools/gopls/internal/telemetry"
 	"golang.org/x/tools/internal/event"
 )
 
 func (s *server) Symbol(ctx context.Context, params *protocol.WorkspaceSymbolParams) (_ []protocol.SymbolInformation, rerr error) {
-	recordLatency := telemetry.StartLatencyTimer("symbol")
-	defer func() {
-		recordLatency(ctx, rerr)
-	}()
-
 	ctx, done := event.Start(ctx, "lsp.Server.symbol")
 	defer done()
 
