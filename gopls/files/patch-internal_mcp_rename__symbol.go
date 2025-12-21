--- internal/mcp/rename_symbol.go.orig	2025-12-05 20:12:17 UTC
+++ internal/mcp/rename_symbol.go
@@ -24,7 +24,6 @@ func (h *handler) renameSymbolHandler(ctx context.Cont
 }
 
 func (h *handler) renameSymbolHandler(ctx context.Context, req *mcp.CallToolRequest, params renameSymbolParams) (*mcp.CallToolResult, any, error) {
-	countGoRenameSymbolMCP.Inc()
 	fh, snapshot, release, err := h.fileOf(ctx, params.File)
 	if err != nil {
 		return nil, nil, err
