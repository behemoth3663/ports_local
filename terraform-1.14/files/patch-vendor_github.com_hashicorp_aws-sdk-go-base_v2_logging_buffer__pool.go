--- vendor/github.com/hashicorp/aws-sdk-go-base/v2/logging/buffer_pool.go.orig	2026-07-29 05:02:03 UTC
+++ vendor/github.com/hashicorp/aws-sdk-go-base/v2/logging/buffer_pool.go
@@ -8,6 +8,8 @@ import (
 	"sync"
 )
 
+const MaxResponseBodyLen = 4096
+
 var BufferPool = newBufferPool()
 
 type bufferPool struct {
