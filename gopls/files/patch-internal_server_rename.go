--- internal/server/rename.go.orig	2025-07-28 18:28:48 UTC
+++ internal/server/rename.go
@@ -17,7 +17,6 @@ func (s *server) Rename(ctx context.Context, params *p
 )
 
 func (s *server) Rename(ctx context.Context, params *protocol.RenameParams) (*protocol.WorkspaceEdit, error) {
-	countRename.Inc()
 	ctx, done := event.Start(ctx, "server.Rename", label.URI.Of(params.TextDocument.URI))
 	defer done()
 
