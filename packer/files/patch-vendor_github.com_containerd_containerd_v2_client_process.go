--- vendor/github.com/containerd/containerd/v2/client/process.go.orig	2026-05-20 23:39:32 UTC
+++ vendor/github.com/containerd/containerd/v2/client/process.go
@@ -29,7 +29,6 @@ import (
 
 	"github.com/containerd/containerd/v2/pkg/cio"
 	"github.com/containerd/containerd/v2/pkg/protobuf"
-	"github.com/containerd/containerd/v2/pkg/tracing"
 )
 
 // Process represents a system process
@@ -121,11 +120,6 @@ func (p *process) Start(ctx context.Context) error {
 
 // Start starts the exec process
 func (p *process) Start(ctx context.Context) error {
-	ctx, span := tracing.StartSpan(ctx, "process.Start",
-		tracing.WithAttribute("process.id", p.ID()),
-		tracing.WithAttribute("process.task.id", p.task.ID()),
-	)
-	defer span.End()
 	r, err := p.task.client.TaskService().Start(ctx, &tasks.StartRequest{
 		ContainerID: p.task.id,
 		ExecID:      p.id,
@@ -138,18 +132,11 @@ func (p *process) Start(ctx context.Context) error {
 		}
 		return errgrpc.ToNative(err)
 	}
-	span.SetAttributes(tracing.Attribute("process.pid", int(r.Pid)))
 	p.pid = r.Pid
 	return nil
 }
 
 func (p *process) Kill(ctx context.Context, s syscall.Signal, opts ...KillOpts) error {
-	ctx, span := tracing.StartSpan(ctx, "process.Kill",
-		tracing.WithAttribute("process.id", p.ID()),
-		tracing.WithAttribute("process.pid", int(p.Pid())),
-		tracing.WithAttribute("process.task.id", p.task.ID()),
-	)
-	defer span.End()
 	var i KillInfo
 	for _, o := range opts {
 		if err := o(ctx, &i); err != nil {
@@ -169,11 +156,6 @@ func (p *process) Wait(ctx context.Context) (<-chan Ex
 	c := make(chan ExitStatus, 1)
 	go func() {
 		defer close(c)
-		ctx, span := tracing.StartSpan(ctx, "process.Wait",
-			tracing.WithAttribute("process.id", p.ID()),
-			tracing.WithAttribute("process.task.id", p.task.ID()),
-		)
-		defer span.End()
 		r, err := p.task.client.TaskService().Wait(ctx, &tasks.WaitRequest{
 			ContainerID: p.task.id,
 			ExecID:      p.id,
@@ -194,10 +176,6 @@ func (p *process) CloseIO(ctx context.Context, opts ..
 }
 
 func (p *process) CloseIO(ctx context.Context, opts ...IOCloserOpts) error {
-	ctx, span := tracing.StartSpan(ctx, "process.CloseIO",
-		tracing.WithAttribute("process.id", p.ID()),
-	)
-	defer span.End()
 	r := &tasks.CloseIORequest{
 		ContainerID: p.task.id,
 		ExecID:      p.id,
@@ -216,10 +194,6 @@ func (p *process) Resize(ctx context.Context, w, h uin
 }
 
 func (p *process) Resize(ctx context.Context, w, h uint32) error {
-	ctx, span := tracing.StartSpan(ctx, "process.Resize",
-		tracing.WithAttribute("process.id", p.ID()),
-	)
-	defer span.End()
 	_, err := p.task.client.TaskService().ResizePty(ctx, &tasks.ResizePtyRequest{
 		ContainerID: p.task.id,
 		Width:       w,
@@ -230,10 +204,6 @@ func (p *process) Delete(ctx context.Context, opts ...
 }
 
 func (p *process) Delete(ctx context.Context, opts ...ProcessDeleteOpts) (*ExitStatus, error) {
-	ctx, span := tracing.StartSpan(ctx, "process.Delete",
-		tracing.WithAttribute("process.id", p.ID()),
-	)
-	defer span.End()
 	for _, o := range opts {
 		if err := o(ctx, p); err != nil {
 			return nil, err
