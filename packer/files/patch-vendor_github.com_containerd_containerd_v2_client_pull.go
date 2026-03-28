--- vendor/github.com/containerd/containerd/v2/client/pull.go.orig	2026-05-20 23:39:32 UTC
+++ vendor/github.com/containerd/containerd/v2/client/pull.go
@@ -29,7 +29,6 @@ import (
 	"github.com/containerd/containerd/v2/core/remotes/docker"
 	"github.com/containerd/containerd/v2/core/transfer"
 	"github.com/containerd/containerd/v2/core/unpack"
-	"github.com/containerd/containerd/v2/pkg/tracing"
 	"github.com/containerd/errdefs"
 	"github.com/containerd/platforms"
 )
@@ -41,9 +40,6 @@ func (c *Client) Pull(ctx context.Context, ref string,
 // Pull downloads the provided content into containerd's content store
 // and returns a platform specific image object
 func (c *Client) Pull(ctx context.Context, ref string, opts ...RemoteOpt) (_ Image, retErr error) {
-	ctx, span := tracing.StartSpan(ctx, tracing.Name(pullSpanPrefix, "Pull"))
-	defer span.End()
-
 	pullCtx := defaultRemoteContext()
 
 	for _, o := range opts {
@@ -75,14 +71,6 @@ func (c *Client) Pull(ctx context.Context, ref string,
 		}
 	}
 
-	span.SetAttributes(
-		tracing.Attribute("image.ref", ref),
-		tracing.Attribute("unpack", pullCtx.Unpack),
-		tracing.Attribute("max.concurrent.downloads", pullCtx.MaxConcurrentDownloads),
-		tracing.Attribute("concurrent.layer.fetch.buffer", pullCtx.ConcurrentLayerFetchBuffer),
-		tracing.Attribute("platforms.count", len(pullCtx.Platforms)),
-	)
-
 	ctx, done, err := c.WithLease(ctx)
 	if err != nil {
 		return nil, err
@@ -96,7 +84,6 @@ func (c *Client) Pull(ctx context.Context, ref string,
 		if err != nil {
 			return nil, fmt.Errorf("unable to resolve snapshotter: %w", err)
 		}
-		span.SetAttributes(tracing.Attribute("snapshotter.name", snapshotterName))
 
 		snCapabilities, err := c.GetSnapshotterCapabilities(ctx, snapshotterName)
 		if err != nil {
@@ -161,13 +148,9 @@ func (c *Client) Pull(ctx context.Context, ref string,
 	// download).
 	var ur unpack.Result
 	if unpacker != nil {
-		_, unpackSpan := tracing.StartSpan(ctx, tracing.Name(pullSpanPrefix, "UnpackWait"))
 		if ur, err = unpacker.Wait(); err != nil {
-			unpackSpan.SetStatus(err)
-			unpackSpan.End()
 			return nil, err
 		}
-		unpackSpan.End()
 	}
 
 	img, err = c.createNewImage(ctx, img)
@@ -176,7 +159,6 @@ func (c *Client) Pull(ctx context.Context, ref string,
 	}
 
 	i := NewImageWithPlatform(c, img, pullCtx.PlatformMatcher)
-	span.SetAttributes(tracing.Attribute("image.ref", i.Name()))
 
 	if unpacker != nil && ur.Unpacks == 0 {
 		// Unpack was tried previously but nothing was unpacked
@@ -190,8 +172,6 @@ func (c *Client) fetch(ctx context.Context, rCtx *Remo
 }
 
 func (c *Client) fetch(ctx context.Context, rCtx *RemoteContext, ref string, limit int) (images.Image, error) {
-	ctx, span := tracing.StartSpan(ctx, tracing.Name(pullSpanPrefix, "fetch"))
-	defer span.End()
 	store := c.ContentStore()
 	name, desc, err := rCtx.Resolver.Resolve(ctx, ref)
 	if err != nil {
@@ -286,8 +266,6 @@ func (c *Client) createNewImage(ctx context.Context, i
 }
 
 func (c *Client) createNewImage(ctx context.Context, img images.Image) (images.Image, error) {
-	ctx, span := tracing.StartSpan(ctx, tracing.Name(pullSpanPrefix, "pull.createNewImage"))
-	defer span.End()
 	is := c.ImageService()
 	for {
 		if created, err := is.Create(ctx, img); err != nil {
