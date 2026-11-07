--- vendor/bazil.org/fuse/mount_freebsd.go.orig	2026-11-07 20:03:41 UTC
+++ vendor/bazil.org/fuse/mount_freebsd.go
@@ -56,10 +56,11 @@ func mount(dir string, conf *mountConfig) (*os.File, e
 		}
 	}
 
-	f, err := os.OpenFile("/dev/fuse", os.O_RDWR, 0o000)
+	fd, err := syscall.Open("/dev/fuse", os.O_RDWR|syscall.O_CLOEXEC, 0o000)
 	if err != nil {
 		return nil, err
 	}
+	f := os.NewFile(uintptr(fd), "/dev/fuse")
 
 	cmd := exec.Command(
 		"/sbin/mount_fusefs",
