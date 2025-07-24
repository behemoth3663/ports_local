--- vendor/cloud.google.com/go/storage/copy.go.orig	2025-05-13 20:48:25 UTC
+++ vendor/cloud.google.com/go/storage/copy.go
@@ -18,8 +18,6 @@ import (
 	"context"
 	"errors"
 	"fmt"
-
-	"cloud.google.com/go/internal/trace"
 )
 
 // CopierFrom creates a Copier that can copy src to dst.
@@ -82,9 +80,6 @@ func (c *Copier) Run(ctx context.Context) (attrs *Obje
 
 // Run performs the copy.
 func (c *Copier) Run(ctx context.Context) (attrs *ObjectAttrs, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.Copier.Run")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	if err := c.src.validate(); err != nil {
 		return nil, err
 	}
@@ -180,9 +175,6 @@ func (c *Composer) Run(ctx context.Context) (attrs *Ob
 
 // Run performs the compose operation.
 func (c *Composer) Run(ctx context.Context) (attrs *ObjectAttrs, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.Composer.Run")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	if err := c.dst.validate(); err != nil {
 		return nil, err
 	}
