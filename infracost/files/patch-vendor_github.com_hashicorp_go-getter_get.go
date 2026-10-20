--- vendor/github.com/hashicorp/go-getter/get.go.orig	2026-10-20 09:33:13 UTC
+++ vendor/github.com/hashicorp/go-getter/get.go
@@ -70,7 +70,6 @@ func init() {
 	Getters = map[string]Getter{
 		"file":  new(FileGetter),
 		"git":   new(GitGetter),
-		"gcs":   new(GCSGetter),
 		"hg":    new(HgGetter),
 		"s3":    new(S3Getter),
 		"http":  httpGetter,
