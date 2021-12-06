--- components/cli/vendor/github.com/tonistiigi/fsutil/stat_unix.go.orig	2021-01-29 21:28:49 UTC
+++ components/cli/vendor/github.com/tonistiigi/fsutil/stat_unix.go
@@ -45,7 +45,7 @@ func setUnixOpt(fi os.FileInfo, stat *types.Stat, path
 			stat.Devminor = int64(minor(uint64(s.Rdev)))
 		}
 
-		ino := s.Ino
+		ino := uint64(s.Ino)
 		linked := false
 		if seenFiles != nil {
 			if s.Nlink > 1 {
