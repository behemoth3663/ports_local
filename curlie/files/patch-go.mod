--- go.mod.orig	2023-05-08 14:01:18 UTC
+++ go.mod
@@ -5,4 +5,6 @@ require (
 	golang.org/x/sys v0.1.0
 )
 
-go 1.13
+require golang.org/x/term v0.1.0 // indirect
+
+go 1.17
