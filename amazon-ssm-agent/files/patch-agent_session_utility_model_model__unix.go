--- agent/session/utility/model/model_unix.go.orig	2020-12-17 08:46:51 UTC
+++ agent/session/utility/model/model_unix.go
@@ -17,5 +17,5 @@
 package model
 
 const (
-	AddUserCommand = "useradd -m %s"
+	AddUserCommand = "pw useradd %s -G wheel -m"
 )
