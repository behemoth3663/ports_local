--- components/cli/vendor/github.com/docker/docker/pkg/system/mknod.go.orig	2021-12-06 17:16:07 UTC
+++ components/cli/vendor/github.com/docker/docker/pkg/system/mknod.go
@@ -9,7 +9,7 @@ import (
 // Mknod creates a filesystem node (file, device special file or named pipe) named path
 // with attributes specified by mode and dev.
 func Mknod(path string, mode uint32, dev int) error {
-	return unix.Mknod(path, mode, dev)
+	return unix.Mknod(path, mode, uint64(dev))
 }
 
 // Mkdev is used to build the value of linux devices (in /dev/) which specifies major
