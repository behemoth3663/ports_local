--- vendor/cloud.google.com/go/storage/storage.go.orig	2023-03-21 19:34:12 UTC
+++ vendor/cloud.google.com/go/storage/storage.go
@@ -39,7 +39,6 @@ import (
 	"unicode/utf8"
 
 	"cloud.google.com/go/internal/optional"
-	"cloud.google.com/go/internal/trace"
 	"cloud.google.com/go/storage/internal"
 	storagepb "cloud.google.com/go/storage/internal/apiv2/stubs"
 	"github.com/googleapis/gax-go/v2"
@@ -917,9 +916,6 @@ func (o *ObjectHandle) Attrs(ctx context.Context) (att
 // Attrs returns meta information about the object.
 // ErrObjectNotExist will be returned if the object is not found.
 func (o *ObjectHandle) Attrs(ctx context.Context) (attrs *ObjectAttrs, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.Object.Attrs")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	if err := o.validate(); err != nil {
 		return nil, err
 	}
@@ -931,9 +927,6 @@ func (o *ObjectHandle) Update(ctx context.Context, uat
 // ObjectAttrsToUpdate docs for details on treatment of zero values.
 // ErrObjectNotExist will be returned if the object is not found.
 func (o *ObjectHandle) Update(ctx context.Context, uattrs ObjectAttrsToUpdate) (oa *ObjectAttrs, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.Object.Update")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	if err := o.validate(); err != nil {
 		return nil, err
 	}
