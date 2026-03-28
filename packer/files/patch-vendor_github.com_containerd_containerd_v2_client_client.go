--- vendor/github.com/containerd/containerd/v2/client/client.go.orig	2026-05-20 23:39:32 UTC
+++ vendor/github.com/containerd/containerd/v2/client/client.go
@@ -76,7 +76,6 @@ import (
 	"github.com/containerd/containerd/v2/pkg/dialer"
 	"github.com/containerd/containerd/v2/pkg/namespaces"
 	ptypes "github.com/containerd/containerd/v2/pkg/protobuf/types"
-	"github.com/containerd/containerd/v2/pkg/tracing"
 	"github.com/containerd/containerd/v2/plugins"
 )
 
@@ -338,8 +337,6 @@ func (c *Client) NewContainer(ctx context.Context, id 
 // NewContainer will create a new container with the provided id.
 // The id must be unique within the namespace.
 func (c *Client) NewContainer(ctx context.Context, id string, opts ...NewContainerOpts) (Container, error) {
-	ctx, span := tracing.StartSpan(ctx, "client.NewContainer")
-	defer span.End()
 	ctx, done, err := c.WithLease(ctx)
 	if err != nil {
 		return nil, err
@@ -363,12 +360,6 @@ func (c *Client) NewContainer(ctx context.Context, id 
 		}
 	}
 
-	span.SetAttributes(
-		tracing.Attribute("container.id", container.ID),
-		tracing.Attribute("container.image.ref", container.Image),
-		tracing.Attribute("container.runtime.name", container.Runtime.Name),
-		tracing.Attribute("container.snapshotter.name", container.Snapshotter),
-	)
 	r, err := c.ContainerService().Create(ctx, container)
 	if err != nil {
 		return nil, err
@@ -378,21 +369,11 @@ func (c *Client) LoadContainer(ctx context.Context, id
 
 // LoadContainer loads an existing container from metadata
 func (c *Client) LoadContainer(ctx context.Context, id string) (Container, error) {
-	ctx, span := tracing.StartSpan(ctx, "client.LoadContainer")
-	defer span.End()
 	r, err := c.ContainerService().Get(ctx, id)
 	if err != nil {
 		return nil, err
 	}
 
-	span.SetAttributes(
-		tracing.Attribute("container.id", r.ID),
-		tracing.Attribute("container.image.ref", r.Image),
-		tracing.Attribute("container.runtime.name", r.Runtime.Name),
-		tracing.Attribute("container.snapshotter.name", r.Snapshotter),
-		tracing.Attribute("container.createdAt", r.CreatedAt.Format(time.RFC3339)),
-		tracing.Attribute("container.updatedAt", r.UpdatedAt.Format(time.RFC3339)),
-	)
 	return containerFromRecord(c, r), nil
 }
 
