--- vendor/cloud.google.com/go/storage/copy.go.orig	2026-08-26 17:25:54 UTC
+++ vendor/cloud.google.com/go/storage/copy.go
@@ -80,9 +80,6 @@ func (c *Copier) Run(ctx context.Context) (attrs *Obje
 
 // Run performs the copy.
 func (c *Copier) Run(ctx context.Context) (attrs *ObjectAttrs, err error) {
-	ctx, _ = startSpanWithBucket(ctx, c.dst.c, c.dst.bucket, "Copier.Run")
-	defer func() { endSpan(ctx, err) }()
-
 	if err := c.src.validate(); err != nil {
 		return nil, err
 	}
@@ -182,9 +179,6 @@ func (c *Composer) Run(ctx context.Context) (attrs *Ob
 
 // Run performs the compose operation.
 func (c *Composer) Run(ctx context.Context) (attrs *ObjectAttrs, err error) {
-	ctx, _ = startSpanWithBucket(ctx, c.dst.c, c.dst.bucket, "Composer.Run")
-	defer func() { endSpan(ctx, err) }()
-
 	if err := c.dst.validate(); err != nil {
 		return nil, err
 	}
