--- internal/mcp/references.go.orig	2025-12-05 20:12:17 UTC
+++ internal/mcp/references.go
@@ -21,7 +21,6 @@ func (h *handler) referencesHandler(ctx context.Contex
 }
 
 func (h *handler) referencesHandler(ctx context.Context, req *mcp.CallToolRequest, params findReferencesParams) (*mcp.CallToolResult, any, error) {
-	countGoReferencesMCP.Inc()
 	fh, snapshot, release, err := h.session.FileOf(ctx, params.Location.URI)
 	if err != nil {
 		return nil, nil, err
