--- internal/mcp/file_metadata.go.orig	2025-12-05 20:12:17 UTC
+++ internal/mcp/file_metadata.go
@@ -17,7 +17,6 @@ func (h *handler) fileMetadataHandler(ctx context.Cont
 }
 
 func (h *handler) fileMetadataHandler(ctx context.Context, req *mcp.CallToolRequest, params fileMetadataParams) (*mcp.CallToolResult, any, error) {
-	countGoFileMetadataMCP.Inc()
 	fh, snapshot, release, err := h.fileOf(ctx, params.File)
 	if err != nil {
 		return nil, nil, err
