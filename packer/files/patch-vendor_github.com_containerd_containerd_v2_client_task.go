--- vendor/github.com/containerd/containerd/v2/client/task.go.orig	2026-05-20 23:39:32 UTC
+++ vendor/github.com/containerd/containerd/v2/client/task.go
@@ -46,7 +46,6 @@ import (
 	"github.com/containerd/containerd/v2/pkg/protobuf"
 	google_protobuf "github.com/containerd/containerd/v2/pkg/protobuf/types"
 	"github.com/containerd/containerd/v2/pkg/rootfs"
-	"github.com/containerd/containerd/v2/pkg/tracing"
 	"github.com/containerd/containerd/v2/plugins"
 )
 
@@ -241,10 +240,6 @@ func (t *task) Start(ctx context.Context) error {
 }
 
 func (t *task) Start(ctx context.Context) error {
-	ctx, span := tracing.StartSpan(ctx, "task.Start",
-		tracing.WithAttribute("task.id", t.ID()),
-	)
-	defer span.End()
 	r, err := t.client.TaskService().Start(ctx, &tasks.StartRequest{
 		ContainerID: t.id,
 	})
@@ -255,17 +250,11 @@ func (t *task) Start(ctx context.Context) error {
 		}
 		return errgrpc.ToNative(err)
 	}
-	span.SetAttributes(tracing.Attribute("task.pid", r.Pid))
 	t.pid = r.Pid
 	return nil
 }
 
 func (t *task) Kill(ctx context.Context, s syscall.Signal, opts ...KillOpts) error {
-	ctx, span := tracing.StartSpan(ctx, "task.Kill",
-		tracing.WithAttribute("task.id", t.ID()),
-		tracing.WithAttribute("task.pid", int(t.Pid())),
-	)
-	defer span.End()
 	var i KillInfo
 	for _, o := range opts {
 		if err := o(ctx, &i); err != nil {
@@ -273,10 +262,6 @@ func (t *task) Kill(ctx context.Context, s syscall.Sig
 		}
 	}
 
-	span.SetAttributes(
-		tracing.Attribute("task.exec.id", i.ExecID),
-		tracing.Attribute("task.exec.killall", i.All),
-	)
 	_, err := t.client.TaskService().Kill(ctx, &tasks.KillRequest{
 		Signal:      uint32(s),
 		ContainerID: t.id,
@@ -290,10 +275,6 @@ func (t *task) Pause(ctx context.Context) error {
 }
 
 func (t *task) Pause(ctx context.Context) error {
-	ctx, span := tracing.StartSpan(ctx, "task.Pause",
-		tracing.WithAttribute("task.id", t.ID()),
-	)
-	defer span.End()
 	_, err := t.client.TaskService().Pause(ctx, &tasks.PauseTaskRequest{
 		ContainerID: t.id,
 	})
@@ -301,10 +282,6 @@ func (t *task) Resume(ctx context.Context) error {
 }
 
 func (t *task) Resume(ctx context.Context) error {
-	ctx, span := tracing.StartSpan(ctx, "task.Resume",
-		tracing.WithAttribute("task.id", t.ID()),
-	)
-	defer span.End()
 	_, err := t.client.TaskService().Resume(ctx, &tasks.ResumeTaskRequest{
 		ContainerID: t.id,
 	})
@@ -333,10 +310,6 @@ func (t *task) Wait(ctx context.Context) (<-chan ExitS
 	c := make(chan ExitStatus, 1)
 	go func() {
 		defer close(c)
-		ctx, span := tracing.StartSpan(ctx, "task.Wait",
-			tracing.WithAttribute("task.id", t.ID()),
-		)
-		defer span.End()
 		r, err := t.client.TaskService().Wait(ctx, &tasks.WaitRequest{
 			ContainerID: t.id,
 		})
@@ -359,10 +332,6 @@ func (t *task) Delete(ctx context.Context, opts ...Pro
 // it returns the exit status of the task and any errors that were encountered
 // during cleanup
 func (t *task) Delete(ctx context.Context, opts ...ProcessDeleteOpts) (*ExitStatus, error) {
-	ctx, span := tracing.StartSpan(ctx, "task.Delete",
-		tracing.WithAttribute("task.id", t.ID()),
-	)
-	defer span.End()
 	for _, o := range opts {
 		if err := o(ctx, t); err != nil {
 			return nil, err
@@ -422,14 +391,9 @@ func (t *task) Exec(ctx context.Context, id string, sp
 }
 
 func (t *task) Exec(ctx context.Context, id string, spec *specs.Process, ioCreate cio.Creator) (_ Process, retErr error) {
-	ctx, span := tracing.StartSpan(ctx, "task.Exec",
-		tracing.WithAttribute("task.id", t.ID()),
-	)
-	defer span.End()
 	if id == "" {
 		return nil, fmt.Errorf("exec id must not be empty: %w", errdefs.ErrInvalidArgument)
 	}
-	span.SetAttributes(tracing.Attribute("task.exec.id", id))
 	i, err := ioCreate(id)
 	if err != nil {
 		return nil, err
@@ -485,10 +449,6 @@ func (t *task) CloseIO(ctx context.Context, opts ...IO
 }
 
 func (t *task) CloseIO(ctx context.Context, opts ...IOCloserOpts) error {
-	ctx, span := tracing.StartSpan(ctx, "task.CloseIO",
-		tracing.WithAttribute("task.id", t.ID()),
-	)
-	defer span.End()
 	r := &tasks.CloseIORequest{
 		ContainerID: t.id,
 	}
@@ -507,10 +467,6 @@ func (t *task) Resize(ctx context.Context, w, h uint32
 }
 
 func (t *task) Resize(ctx context.Context, w, h uint32) error {
-	ctx, span := tracing.StartSpan(ctx, "task.Resize",
-		tracing.WithAttribute("task.id", t.ID()),
-	)
-	defer span.End()
 	_, err := t.client.TaskService().ResizePty(ctx, &tasks.ResizePtyRequest{
 		ContainerID: t.id,
 		Width:       w,
@@ -624,10 +580,6 @@ func (t *task) Update(ctx context.Context, opts ...Upd
 type UpdateTaskOpts func(context.Context, *Client, *UpdateTaskInfo) error
 
 func (t *task) Update(ctx context.Context, opts ...UpdateTaskOpts) error {
-	ctx, span := tracing.StartSpan(ctx, "task.Update",
-		tracing.WithAttribute("task.id", t.ID()),
-	)
-	defer span.End()
 	request := &tasks.UpdateTaskRequest{
 		ContainerID: t.id,
 	}
