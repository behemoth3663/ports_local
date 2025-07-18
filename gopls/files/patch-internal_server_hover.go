--- internal/server/hover.go.orig	2025-06-18 21:21:32 UTC
+++ internal/server/hover.go
@@ -13,18 +13,12 @@ import (
 	"golang.org/x/tools/gopls/internal/mod"
 	"golang.org/x/tools/gopls/internal/protocol"
 	"golang.org/x/tools/gopls/internal/settings"
-	"golang.org/x/tools/gopls/internal/telemetry"
 	"golang.org/x/tools/gopls/internal/template"
 	"golang.org/x/tools/gopls/internal/work"
 	"golang.org/x/tools/internal/event"
 )
 
 func (s *server) Hover(ctx context.Context, params *protocol.HoverParams) (_ *protocol.Hover, rerr error) {
-	recordLatency := telemetry.StartLatencyTimer("hover")
-	defer func() {
-		recordLatency(ctx, rerr)
-	}()
-
 	ctx, done := event.Start(ctx, "server.Hover", label.URI.Of(params.TextDocument.URI))
 	defer done()
 
