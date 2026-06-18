--- cli/config/credentials/default_store_unsupported.go.orig	2026-06-03 17:16:33 UTC
+++ cli/config/credentials/default_store_unsupported.go
@@ -1,4 +1,4 @@
-//go:build !windows && !darwin && !linux
+//go:build !windows && !darwin && !linux && !freebsd
 
 package credentials
 
