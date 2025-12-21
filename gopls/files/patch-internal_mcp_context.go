--- internal/mcp/context.go.orig	2025-12-05 20:12:17 UTC
+++ internal/mcp/context.go
@@ -32,7 +32,6 @@ func (h *handler) contextHandler(ctx context.Context, 
 }
 
 func (h *handler) contextHandler(ctx context.Context, req *mcp.CallToolRequest, params ContextParams) (*mcp.CallToolResult, any, error) {
-	countGoContextMCP.Inc()
 	fh, snapshot, release, err := h.fileOf(ctx, params.File)
 	if err != nil {
 		return nil, nil, err
