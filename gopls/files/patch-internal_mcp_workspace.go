--- internal/mcp/workspace.go.orig	2025-12-05 20:12:17 UTC
+++ internal/mcp/workspace.go
@@ -20,7 +20,6 @@ func (h *handler) workspaceHandler(ctx context.Context
 )
 
 func (h *handler) workspaceHandler(ctx context.Context, req *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, any, error) {
-	countGoWorkspaceMCP.Inc()
 	var summary bytes.Buffer
 	views := h.session.Views()
 	for _, v := range views {
