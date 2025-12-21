--- internal/mcp/search.go.orig	2025-12-05 20:12:17 UTC
+++ internal/mcp/search.go
@@ -18,7 +18,6 @@ func (h *handler) searchHandler(ctx context.Context, r
 }
 
 func (h *handler) searchHandler(ctx context.Context, req *mcp.CallToolRequest, params searchParams) (*mcp.CallToolResult, any, error) {
-	countGoSearchMCP.Inc()
 	query := params.Query
 	if len(query) == 0 {
 		return nil, nil, fmt.Errorf("empty query")
