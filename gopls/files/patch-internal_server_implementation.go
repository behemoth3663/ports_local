--- internal/server/implementation.go.orig	2025-06-18 21:21:32 UTC
+++ internal/server/implementation.go
@@ -11,16 +11,10 @@ import (
 	"golang.org/x/tools/gopls/internal/golang"
 	"golang.org/x/tools/gopls/internal/label"
 	"golang.org/x/tools/gopls/internal/protocol"
-	"golang.org/x/tools/gopls/internal/telemetry"
 	"golang.org/x/tools/internal/event"
 )
 
 func (s *server) Implementation(ctx context.Context, params *protocol.ImplementationParams) (_ []protocol.Location, rerr error) {
-	recordLatency := telemetry.StartLatencyTimer("implementation")
-	defer func() {
-		recordLatency(ctx, rerr)
-	}()
-
 	ctx, done := event.Start(ctx, "server.Implementation", label.URI.Of(params.TextDocument.URI))
 	defer done()
 
