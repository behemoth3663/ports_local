--- internal/mcp/workspace.go.orig	2026-09-30 18:55:09 UTC
+++ internal/mcp/workspace.go
@@ -34,7 +34,6 @@ func (h *handler) workspaceHandler(ctx context.Context
 }
 
 func (h *handler) workspaceHandler(ctx context.Context, req *mcp.CallToolRequest, params WorkspaceParams) (*mcp.CallToolResult, any, error) {
-	countGoWorkspaceMCP.Inc()
 	var summary bytes.Buffer
 	views := h.session.Views()
 	if params.Dir != "" {
