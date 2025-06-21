--- internal/server/definition.go.orig	2025-06-18 21:21:32 UTC
+++ internal/server/definition.go
@@ -13,17 +13,11 @@ import (
 	"golang.org/x/tools/gopls/internal/golang"
 	"golang.org/x/tools/gopls/internal/label"
 	"golang.org/x/tools/gopls/internal/protocol"
-	"golang.org/x/tools/gopls/internal/telemetry"
 	"golang.org/x/tools/gopls/internal/template"
 	"golang.org/x/tools/internal/event"
 )
 
 func (s *server) Definition(ctx context.Context, params *protocol.DefinitionParams) (_ []protocol.Location, rerr error) {
-	recordLatency := telemetry.StartLatencyTimer("definition")
-	defer func() {
-		recordLatency(ctx, rerr)
-	}()
-
 	ctx, done := event.Start(ctx, "server.Definition", label.URI.Of(params.TextDocument.URI))
 	defer done()
 
