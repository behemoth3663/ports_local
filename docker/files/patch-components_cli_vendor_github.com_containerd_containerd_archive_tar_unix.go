--- components/cli/vendor/github.com/containerd/containerd/archive/tar_unix.go.orig	2021-01-29 21:28:49 UTC
+++ components/cli/vendor/github.com/containerd/containerd/archive/tar_unix.go
@@ -122,7 +122,7 @@ func handleTarTypeBlockCharFifo(hdr *tar.Header, path 
 		mode |= unix.S_IFIFO
 	}
 
-	return unix.Mknod(path, mode, int(unix.Mkdev(uint32(hdr.Devmajor), uint32(hdr.Devminor))))
+	return unix.Mknod(path, mode, uint64(unix.Mkdev(uint32(hdr.Devmajor), uint32(hdr.Devminor))))
 }
 
 func handleLChmod(hdr *tar.Header, path string, hdrInfo os.FileInfo) error {
