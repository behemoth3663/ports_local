--- go.mod.orig	2025-04-11 15:49:28 UTC
+++ go.mod
@@ -1,6 +1,8 @@ module github.com/jreisinger/checkip
 module github.com/jreisinger/checkip
 
-go 1.18
+go 1.23.0
+
+toolchain go1.24.2
 
 require (
 	github.com/go-ping/ping v1.1.0
