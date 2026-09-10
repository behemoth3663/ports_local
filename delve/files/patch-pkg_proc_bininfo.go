--- pkg/proc/bininfo.go.orig	2026-09-09 15:48:17 UTC
+++ pkg/proc/bininfo.go
@@ -2377,7 +2377,6 @@ func loadBinaryInfoGoRuntimeElf(bi *BinaryInfo, image 
 	defer func() {
 		ierr := recover()
 		if ierr != nil {
-			logflags.Bug.Inc()
 			err = fmt.Errorf("error loading binary info from Go runtime: %v", ierr)
 		}
 	}()
@@ -2433,7 +2432,6 @@ func loadBinaryInfoGoRuntimeMacho(bi *BinaryInfo, imag
 	defer func() {
 		ierr := recover()
 		if ierr != nil {
-			logflags.Bug.Inc()
 			err = fmt.Errorf("error loading binary info from Go runtime: %v", ierr)
 		}
 	}()
