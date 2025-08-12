--- internal/cmd/serve.go.orig	2025-07-28 18:28:48 UTC
+++ internal/cmd/serve.go
@@ -123,7 +123,6 @@ func (s *Serve) Run(ctx context.Context, args ...strin
 
 	// Start MCP server.
 	if sessions != nil {
-		countAttachedMCP.Inc()
 		group.Go(func() (err error) {
 			defer func() {
 				if err == nil {
