--- internal/mcp/vulncheck.go.orig	2025-12-05 20:12:17 UTC
+++ internal/mcp/vulncheck.go
@@ -33,7 +33,6 @@ func (h *handler) vulncheckHandler(ctx context.Context
 }
 
 func (h *handler) vulncheckHandler(ctx context.Context, req *mcp.CallToolRequest, params *vulncheckParams) (*mcp.CallToolResult, *VulncheckResultOutput, error) {
-	countGoVulncheckMCP.Inc()
 	snapshot, release, err := h.snapshot()
 	if err != nil {
 		return nil, nil, err
