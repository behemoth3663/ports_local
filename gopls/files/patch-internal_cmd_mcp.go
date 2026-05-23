--- internal/cmd/mcp.go.orig	2026-05-15 14:53:57 UTC
+++ internal/cmd/mcp.go
@@ -194,10 +194,8 @@ func (m *headlessMCP) Run(ctx context.Context, args ..
 	}()
 
 	if m.Address != "" {
-		countHeadlessMCPSSE.Inc()
 		return internalmcp.Serve(ctx, m.Address, &staticSessions{sess, cli.server}, false, watchRoots)
 	} else {
-		countHeadlessMCPStdIO.Inc()
 		var rpcLog io.Writer
 		if m.RPCTrace {
 			rpcLog = log.Writer() // possibly redirected by -logfile above
