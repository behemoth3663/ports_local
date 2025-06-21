--- internal/server/text_synchronization.go.orig	2025-06-18 21:21:32 UTC
+++ internal/server/text_synchronization.go
@@ -315,7 +315,6 @@ func (s *server) changedText(ctx context.Context, uri 
 	// Check if the client sent the full content of the file.
 	// We accept a full content change even if the server expected incremental changes.
 	if len(changes) == 1 && changes[0].Range == nil && changes[0].RangeLength == nil {
-		changeFull.Inc()
 		return []byte(changes[0].Text), nil
 	}
 	return s.applyIncrementalChanges(ctx, uri, changes)
@@ -396,17 +395,14 @@ func (s *server) checkEfficacy(uri protocol.DocumentUR
 			ix := strings.Index(edit.NewText, "$")
 			if ix < 0 && strings.HasPrefix(change.Text, edit.NewText) {
 				// not a snippet, suggested completion is a prefix of the change
-				complUsed.Inc()
 				return
 			}
 			if ix > 1 && strings.HasPrefix(change.Text, edit.NewText[:ix]) {
 				// a snippet, suggested completion up to $ marker is a prefix of the change
-				complUsed.Inc()
 				return
 			}
 		}
 	}
-	complUnused.Inc()
 }
 
 func changeTypeToFileAction(ct protocol.FileChangeType) file.Action {
