--- internal/state/walker_paths.go.orig	2023-10-12 15:35:06 UTC
+++ internal/state/walker_paths.go
@@ -12,7 +12,6 @@ import (
 
 	"github.com/hashicorp/go-memdb"
 	"github.com/hashicorp/terraform-ls/internal/document"
-	"go.opentelemetry.io/otel/trace"
 )
 
 type WalkerPathStore struct {
@@ -28,7 +27,6 @@ type WalkerPath struct {
 	Dir            document.DirHandle
 	IsDirOpen      bool
 	State          PathState
-	EnqueueContext trace.SpanContext
 }
 
 type PathContext struct {
@@ -43,15 +41,9 @@ const (
 )
 
 func (wp *WalkerPath) Copy() *WalkerPath {
-	// This may be an awkward way to copy the SpanContext
-	// but the upstream doesn't seem to offer any better way.
-	newCtx := trace.ContextWithSpanContext(context.Background(), wp.EnqueueContext)
-	spanContext := trace.SpanContextFromContext(newCtx)
-
 	return &WalkerPath{
 		Dir:            wp.Dir,
 		IsDirOpen:      wp.IsDirOpen,
-		EnqueueContext: spanContext,
 	}
 }
 
@@ -65,7 +57,7 @@ func (pa *PathAwaiter) AwaitNextDir(ctx context.Contex
 	if err != nil {
 		return ctx, document.DirHandle{}, err
 	}
-	return trace.ContextWithSpanContext(ctx, wp.EnqueueContext), wp.Dir, nil
+	return ctx, wp.Dir, nil
 }
 
 func (pa *PathAwaiter) RemoveDir(dir document.DirHandle) error {
@@ -96,7 +88,6 @@ func (wps *WalkerPathStore) EnqueueDir(ctx context.Con
 		Dir:            dir,
 		IsDirOpen:      false,
 		State:          PathStateQueued,
-		EnqueueContext: trace.SpanContextFromContext(ctx),
 	})
 	if err != nil {
 		return err
