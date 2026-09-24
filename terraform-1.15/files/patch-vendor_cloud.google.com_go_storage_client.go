--- vendor/cloud.google.com/go/storage/client.go.orig	2025-03-12 21:31:11 UTC
+++ vendor/cloud.google.com/go/storage/client.go
@@ -136,7 +136,6 @@ type settings struct {
 	// userProject is the user project that should be billed for the request.
 	userProject string
 
-	metricsContext *metricsContext
 }
 
 func initSettings(opts ...storageOption) *settings {
