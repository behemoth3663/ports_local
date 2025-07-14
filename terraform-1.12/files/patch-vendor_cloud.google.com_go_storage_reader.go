--- vendor/cloud.google.com/go/storage/reader.go.orig	2023-03-21 19:34:12 UTC
+++ vendor/cloud.google.com/go/storage/reader.go
@@ -23,8 +23,6 @@ import (
 	"net/http"
 	"strings"
 	"time"
-
-	"cloud.google.com/go/internal/trace"
 )
 
 var crc32cTable = crc32.MakeTable(crc32.Castagnoli)
@@ -87,9 +85,6 @@ func (o *ObjectHandle) NewRangeReader(ctx context.Cont
 // that file will be served back whole, regardless of the requested range as
 // Google Cloud Storage dictates.
 func (o *ObjectHandle) NewRangeReader(ctx context.Context, offset, length int64) (r *Reader, err error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.Object.NewRangeReader")
-	defer func() { trace.EndSpan(ctx, err) }()
-
 	if err := o.validate(); err != nil {
 		return nil, err
 	}
