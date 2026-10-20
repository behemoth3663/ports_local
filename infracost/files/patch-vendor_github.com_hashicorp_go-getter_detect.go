--- vendor/github.com/hashicorp/go-getter/detect.go.orig	2026-10-20 09:33:13 UTC
+++ vendor/github.com/hashicorp/go-getter/detect.go
@@ -30,7 +30,6 @@ func init() {
 		new(GitDetector),
 		new(BitBucketDetector),
 		new(S3Detector),
-		new(GCSDetector),
 		new(FileDetector),
 	}
 }
