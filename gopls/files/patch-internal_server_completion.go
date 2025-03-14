--- internal/server/completion.go.orig	2025-02-24 17:39:34 UTC
+++ internal/server/completion.go
@@ -15,18 +15,12 @@ import (
 	"golang.org/x/tools/gopls/internal/label"
 	"golang.org/x/tools/gopls/internal/protocol"
 	"golang.org/x/tools/gopls/internal/settings"
-	"golang.org/x/tools/gopls/internal/telemetry"
 	"golang.org/x/tools/gopls/internal/template"
 	"golang.org/x/tools/gopls/internal/work"
 	"golang.org/x/tools/internal/event"
 )
 
 func (s *server) Completion(ctx context.Context, params *protocol.CompletionParams) (_ *protocol.CompletionList, rerr error) {
-	recordLatency := telemetry.StartLatencyTimer("completion")
-	defer func() {
-		recordLatency(ctx, rerr)
-	}()
-
 	ctx, done := event.Start(ctx, "lsp.Server.completion", label.URI.Of(params.TextDocument.URI))
 	defer done()
 
@@ -61,7 +55,6 @@ func (s *server) Completion(ctx context.Context, param
 		event.Error(ctx, "no completions found", err, label.Position.Of(params.Position))
 	}
 	if candidates == nil || surrounding == nil {
-		complEmpty.Inc()
 		return &protocol.CompletionList{
 			IsIncomplete: true,
 			Items:        []protocol.CompletionItem{},
@@ -81,12 +74,6 @@ func (s *server) Completion(ctx context.Context, param
 		s.saveLastCompletion(fh.URI(), fh.Version(), items, params.Position)
 	}
 
-	if len(items) > 10 {
-		// TODO(pjw): long completions are ok for field lists
-		complLong.Inc()
-	} else {
-		complShort.Inc()
-	}
 	return &protocol.CompletionList{
 		IsIncomplete: incompleteResults,
 		Items:        items,
