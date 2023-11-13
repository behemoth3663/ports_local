--- internal/langserver/langserver.go.orig	2023-10-12 15:35:06 UTC
+++ internal/langserver/langserver.go
@@ -17,7 +17,6 @@ import (
 	"github.com/creachadair/jrpc2/channel"
 	"github.com/creachadair/jrpc2/server"
 	"github.com/hashicorp/terraform-ls/internal/langserver/session"
-	"go.opentelemetry.io/otel/trace"
 )
 
 type langServer struct {
@@ -39,8 +38,7 @@ func NewLangServer(srvCtx context.Context, sf session.
 		AllowPush:   true,
 		Concurrency: concurrency,
 		NewContext: func() context.Context {
-			spanCtx := trace.SpanContextFromContext(srvCtx)
-			return trace.ContextWithSpanContext(context.Background(), spanCtx)
+			return srvCtx
 		},
 	}
