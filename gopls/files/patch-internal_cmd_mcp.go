--- internal/cmd/mcp.go.orig	2025-07-28 18:28:48 UTC
+++ internal/cmd/mcp.go
@@ -143,10 +143,8 @@ func (m *headlessMCP) Run(ctx context.Context, args ..
 	}
 
 	if m.Address != "" {
-		countHeadlessMCPSSE.Inc()
 		return mcp.Serve(ctx, m.Address, &staticSessions{sess, cli.server}, false)
 	} else {
-		countHeadlessMCPStdIO.Inc()
 		var rpcLog io.Writer
 		if m.RPCTrace {
 			rpcLog = log.Writer() // possibly redirected by -logfile above
