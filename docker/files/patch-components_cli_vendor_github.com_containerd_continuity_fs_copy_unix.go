--- components/cli/vendor/github.com/containerd/continuity/fs/copy_unix.go.orig	2021-12-06 17:14:45 UTC
+++ components/cli/vendor/github.com/containerd/continuity/fs/copy_unix.go
@@ -108,5 +108,5 @@ func copyDevice(dst string, fi os.FileInfo) error {
 	if !ok {
 		return errors.New("unsupported stat type")
 	}
-	return unix.Mknod(dst, uint32(fi.Mode()), int(st.Rdev))
+	return unix.Mknod(dst, uint32(fi.Mode()), uint64(st.Rdev))
 }
