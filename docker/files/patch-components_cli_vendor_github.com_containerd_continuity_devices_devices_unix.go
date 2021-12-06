--- components/cli/vendor/github.com/containerd/continuity/devices/devices_unix.go.orig	2021-12-06 17:19:09 UTC
+++ components/cli/vendor/github.com/containerd/continuity/devices/devices_unix.go
@@ -55,7 +55,7 @@ func Mknod(p string, mode os.FileMode, maj, min int) e
 		m |= unix.S_IFIFO
 	}
 
-	return unix.Mknod(p, m, int(dev))
+	return unix.Mknod(p, m, uint64(dev))
 }
 
 // syscallMode returns the syscall-specific mode bits from Go's portable mode bits.
