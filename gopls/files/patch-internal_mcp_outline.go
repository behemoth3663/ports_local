--- internal/mcp/outline.go.orig	2025-12-05 20:12:17 UTC
+++ internal/mcp/outline.go
@@ -18,7 +18,6 @@ func (h *handler) outlineHandler(ctx context.Context, 
 }
 
 func (h *handler) outlineHandler(ctx context.Context, req *mcp.CallToolRequest, params outlineParams) (*mcp.CallToolResult, any, error) {
-	countGoPackageAPIMCP.Inc()
 	snapshot, release, err := h.snapshot()
 	if err != nil {
 		return nil, nil, err
