--- vendor/cloud.google.com/go/storage/client.go.orig	2025-02-06 22:11:02 UTC
+++ vendor/cloud.google.com/go/storage/client.go
@@ -135,8 +135,6 @@ type settings struct {
 
 	// userProject is the user project that should be billed for the request.
 	userProject string
-
-	metricsContext *metricsContext
 }
 
 func initSettings(opts ...storageOption) *settings {
