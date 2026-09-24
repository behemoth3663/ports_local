--- internal/rpcapi/setup.go.orig	2024-10-16 12:28:59 UTC
+++ internal/rpcapi/setup.go
@@ -52,9 +52,7 @@ func (s *setupServer) Handshake(ctx context.Context, r
 	var serverCaps *terraform1.ServerCapabilities
 	var err error
 	{
-		ctx, span := tracer.Start(ctx, "initialize RPC services")
 		serverCaps, err = s.initOthers(ctx, req, s.stopper)
-		span.End()
 	}
 	s.initOthers = nil // cannot handshake again
 	if err != nil {
