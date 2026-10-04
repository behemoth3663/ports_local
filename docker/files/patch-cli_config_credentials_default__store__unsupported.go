--- cli/config/credentials/default_store_unsupported.go.orig	2026-09-23 09:45:10 UTC
+++ cli/config/credentials/default_store_unsupported.go
@@ -1,4 +1,4 @@
-//go:build !windows && !darwin && !linux
+//go:build !windows && !darwin && !linux && !freebsd
 
 package credentials
 
