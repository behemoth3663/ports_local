--- internal/mcp/file_context.go.orig	2025-12-05 20:12:17 UTC
+++ internal/mcp/file_context.go
@@ -21,7 +21,6 @@ func (h *handler) fileContextHandler(ctx context.Conte
 }
 
 func (h *handler) fileContextHandler(ctx context.Context, req *mcp.CallToolRequest, params fileContextParams) (*mcp.CallToolResult, any, error) {
-	countGoFileContextMCP.Inc()
 	fh, snapshot, release, err := h.fileOf(ctx, params.File)
 	if err != nil {
 		return nil, nil, err
