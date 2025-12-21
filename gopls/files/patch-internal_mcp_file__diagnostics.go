--- internal/mcp/file_diagnostics.go.orig	2025-12-05 20:12:17 UTC
+++ internal/mcp/file_diagnostics.go
@@ -29,7 +29,6 @@ func (h *handler) fileDiagnosticsHandler(ctx context.C
 }
 
 func (h *handler) fileDiagnosticsHandler(ctx context.Context, req *mcp.CallToolRequest, params diagnosticsParams) (*mcp.CallToolResult, any, error) {
-	countGoFileDiagnosticsMCP.Inc()
 	fh, snapshot, release, err := h.fileOf(ctx, params.File)
 	if err != nil {
 		return nil, nil, err
