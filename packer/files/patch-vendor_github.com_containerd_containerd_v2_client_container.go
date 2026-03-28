--- vendor/github.com/containerd/containerd/v2/client/container.go.orig	2026-05-20 23:39:32 UTC
+++ vendor/github.com/containerd/containerd/v2/client/container.go
@@ -40,7 +40,6 @@ import (
 	"github.com/containerd/containerd/v2/core/images"
 	"github.com/containerd/containerd/v2/pkg/cio"
 	"github.com/containerd/containerd/v2/pkg/oci"
-	"github.com/containerd/containerd/v2/pkg/tracing"
 )
 
 const (
@@ -146,10 +145,6 @@ func (c *container) SetLabels(ctx context.Context, lab
 }
 
 func (c *container) SetLabels(ctx context.Context, labels map[string]string) (map[string]string, error) {
-	ctx, span := tracing.StartSpan(ctx, "container.SetLabels",
-		tracing.WithAttribute("container.id", c.id),
-	)
-	defer span.End()
 	container := containers.Container{
 		ID:     c.id,
 		Labels: labels,
@@ -185,10 +180,6 @@ func (c *container) Delete(ctx context.Context, opts .
 // Delete deletes an existing container
 // an error is returned if the container has running tasks
 func (c *container) Delete(ctx context.Context, opts ...DeleteOpts) error {
-	ctx, span := tracing.StartSpan(ctx, "container.Delete",
-		tracing.WithAttribute("container.id", c.id),
-	)
-	defer span.End()
 	if _, err := c.loadTask(ctx, nil); err == nil {
 		return fmt.Errorf("cannot delete running task %v: %w", c.id, errdefs.ErrFailedPrecondition)
 	}
@@ -225,8 +216,6 @@ func (c *container) NewTask(ctx context.Context, ioCre
 }
 
 func (c *container) NewTask(ctx context.Context, ioCreate cio.Creator, opts ...NewTaskOpts) (_ Task, retErr error) {
-	ctx, span := tracing.StartSpan(ctx, "container.NewTask")
-	defer span.End()
 	i, err := ioCreate(c.id)
 	if err != nil {
 		return nil, err
@@ -288,27 +277,17 @@ func (c *container) NewTask(ctx context.Context, ioCre
 		request.Checkpoint = info.Checkpoint
 	}
 
-	span.SetAttributes(
-		tracing.Attribute("task.container.id", request.ContainerID),
-		tracing.Attribute("task.request.options", request.Options.String()),
-		tracing.Attribute("task.runtime.name", info.runtime),
-	)
 	response, err := c.client.TaskService().Create(ctx, request)
 	if err != nil {
 		return nil, errgrpc.ToNative(err)
 	}
 
-	span.AddEvent("task created",
-		tracing.Attribute("task.process.id", int(response.Pid)),
-	)
 	t.pid = response.Pid
 	return t, nil
 }
 
 func (c *container) Update(ctx context.Context, opts ...UpdateContainerOpts) error {
 	// fetch the current container config before updating it
-	ctx, span := tracing.StartSpan(ctx, "container.Update")
-	defer span.End()
 	r, err := c.get(ctx)
 	if err != nil {
 		return err
