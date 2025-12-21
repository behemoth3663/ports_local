--- internal/mcp/symbol_references.go.orig	2025-12-05 20:12:17 UTC
+++ internal/mcp/symbol_references.go
@@ -30,7 +30,6 @@ func (h *handler) symbolReferencesHandler(ctx context.
 // It finds all references to the requested symbol and describes their
 // locations.
 func (h *handler) symbolReferencesHandler(ctx context.Context, req *mcp.CallToolRequest, params symbolReferencesParams) (*mcp.CallToolResult, any, error) {
-	countGoSymbolReferencesMCP.Inc()
 	fh, snapshot, release, err := h.fileOf(ctx, params.File)
 	if err != nil {
 		return nil, nil, err
