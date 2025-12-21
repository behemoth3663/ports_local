--- internal/mcp/workspace_diagnostics.go.orig	2025-12-05 20:12:17 UTC
+++ internal/mcp/workspace_diagnostics.go
@@ -23,7 +23,6 @@ func (h *handler) workspaceDiagnosticsHandler(ctx cont
 }
 
 func (h *handler) workspaceDiagnosticsHandler(ctx context.Context, req *mcp.CallToolRequest, params workspaceDiagnosticsParams) (*mcp.CallToolResult, any, error) {
-	countGoDiagnosticsMCP.Inc()
 	var (
 		fh       file.Handle
 		snapshot *cache.Snapshot
