--- vendor/github.com/containerd/containerd/v2/core/unpack/unpacker.go.orig	2026-05-20 23:39:32 UTC
+++ vendor/github.com/containerd/containerd/v2/core/unpack/unpacker.go
@@ -24,7 +24,6 @@ import (
 	"errors"
 	"fmt"
 	"slices"
-	"strconv"
 	"sync"
 	"sync/atomic"
 	"time"
@@ -45,7 +44,6 @@ import (
 	"github.com/containerd/containerd/v2/internal/cleanup"
 	"github.com/containerd/containerd/v2/internal/kmutex"
 	"github.com/containerd/containerd/v2/pkg/labels"
-	"github.com/containerd/containerd/v2/pkg/tracing"
 )
 
 const (
@@ -218,11 +216,6 @@ func (u *Unpacker) Unpack(h images.Handler) images.Han
 	}
 
 	return images.HandlerFunc(func(ctx context.Context, desc ocispec.Descriptor) ([]ocispec.Descriptor, error) {
-		ctx, span := tracing.StartSpan(ctx, tracing.Name(unpackSpanPrefix, "UnpackHandler"))
-		defer span.End()
-		span.SetAttributes(
-			tracing.Attribute("descriptor.media.type", desc.MediaType),
-			tracing.Attribute("descriptor.digest", desc.Digest.String()))
 		unlock, err := u.lockBlobDescriptor(ctx, desc)
 		if err != nil {
 			return nil, err
@@ -238,10 +231,7 @@ func (u *Unpacker) Unpack(h images.Handler) images.Han
 			var manifestLayers []ocispec.Descriptor
 			// Split layers from non-layers, layers will be handled after
 			// the config
-			for i, child := range children {
-				span.SetAttributes(
-					tracing.Attribute("descriptor.child."+strconv.Itoa(i), []string{child.MediaType, child.Digest.String()}),
-				)
+			for _, child := range children {
 				if images.IsLayerType(child.MediaType) || layerTypes[child.MediaType] {
 					manifestLayers = append(manifestLayers, child)
 				} else {
@@ -295,7 +285,6 @@ type unpackStatus struct {
 	err     error
 	desc    ocispec.Descriptor
 	bottomF func(bool) error
-	span    *tracing.Span
 	startAt time.Time
 }
 
@@ -305,8 +294,6 @@ func (u *Unpacker) unpack(
 	layers []ocispec.Descriptor,
 ) error {
 	ctx := u.ctx
-	ctx, layerSpan := tracing.StartSpan(ctx, tracing.Name(unpackSpanPrefix, "unpack"))
-	defer layerSpan.End()
 	unpackStart := time.Now()
 	p, err := content.ReadBlob(ctx, u.content, config)
 	if err != nil {
@@ -370,7 +357,7 @@ func (u *Unpacker) unpack(
 	copy(chainIDs, diffIDs)
 	chainIDs = identity.ChainIDs(chainIDs)
 
-	topHalf := func(i int, desc ocispec.Descriptor, span *tracing.Span, startAt time.Time) (<-chan *unpackStatus, error) {
+	topHalf := func(i int, desc ocispec.Descriptor, startAt time.Time) (<-chan *unpackStatus, error) {
 		var (
 			err     error
 			parent  string
@@ -472,7 +459,6 @@ func (u *Unpacker) unpack(
 
 			status := &unpackStatus{
 				desc:    desc,
-				span:    span,
 				startAt: startAt,
 				bottomF: func(shouldAbort bool) error {
 					defer unlock()
@@ -566,8 +552,6 @@ func (u *Unpacker) unpack(
 			err = s.bottomF(false)
 		}
 
-		s.span.SetStatus(err)
-		s.span.End()
 		if err == nil {
 			log.G(ctx).WithFields(log.Fields{
 				"layer":    s.desc.Digest,
@@ -580,26 +564,17 @@ func (u *Unpacker) unpack(
 	var statusChans []<-chan *unpackStatus
 
 	for i, desc := range layers {
-		_, layerSpan := tracing.StartSpan(ctx, tracing.Name(unpackSpanPrefix, "unpackLayer"))
 		unpackLayerStart := time.Now()
-		layerSpan.SetAttributes(
-			tracing.Attribute("layer.media.type", desc.MediaType),
-			tracing.Attribute("layer.media.size", desc.Size),
-			tracing.Attribute("layer.media.digest", desc.Digest.String()),
-		)
-		statusCh, err := topHalf(i, desc, layerSpan, unpackLayerStart)
+		statusCh, err := topHalf(i, desc, unpackLayerStart)
 		if err != nil {
 			if parallel {
 				break
 			} else {
-				layerSpan.SetStatus(err)
-				layerSpan.End()
 				return err
 			}
 		}
 		if statusCh == nil {
 			// nothing to do, already exists
-			layerSpan.End()
 			continue
 		}
 		if parallel {
@@ -651,12 +626,6 @@ func (u *Unpacker) fetch(ctx context.Context, h images
 func (u *Unpacker) fetch(ctx context.Context, h images.Handler, layers []ocispec.Descriptor, done []chan struct{}) error {
 	eg, ctx2 := errgroup.WithContext(ctx)
 	for i, desc := range layers {
-		ctx2, layerSpan := tracing.StartSpan(ctx2, tracing.Name(unpackSpanPrefix, "fetchLayer"))
-		layerSpan.SetAttributes(
-			tracing.Attribute("layer.media.type", desc.MediaType),
-			tracing.Attribute("layer.media.size", desc.Size),
-			tracing.Attribute("layer.media.digest", desc.Digest.String()),
-		)
 		var ch chan struct{}
 		if done != nil {
 			ch = done[i]
@@ -667,8 +636,6 @@ func (u *Unpacker) fetch(ctx context.Context, h images
 		}
 
 		eg.Go(func() error {
-			defer layerSpan.End()
-
 			unlock, err := u.lockBlobDescriptor(ctx2, desc)
 			if err != nil {
 				u.release(u.limiter)
