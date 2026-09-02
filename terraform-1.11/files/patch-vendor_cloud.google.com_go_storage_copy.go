--- vendor/cloud.google.com/go/storage/copy.go.orig	2026-06-04 04:52:32 UTC
+++ vendor/cloud.google.com/go/storage/copy.go
@@ -80,8 +80,6 @@ func (c *Copier) Run(ctx context.Context) (attrs *Obje
 
 // Run performs the copy.
 func (c *Copier) Run(ctx context.Context) (attrs *ObjectAttrs, err error) {
-	ctx, _ = startSpan(ctx, "Copier.Run")
-	defer func() { endSpan(ctx, err) }()
 
 	if err := c.src.validate(); err != nil {
 		return nil, err
@@ -178,8 +176,6 @@ func (c *Composer) Run(ctx context.Context) (attrs *Ob
 
 // Run performs the compose operation.
 func (c *Composer) Run(ctx context.Context) (attrs *ObjectAttrs, err error) {
-	ctx, _ = startSpan(ctx, "Composer.Run")
-	defer func() { endSpan(ctx, err) }()
 
 	if err := c.dst.validate(); err != nil {
 		return nil, err
