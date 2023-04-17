--- go.mod.orig	2023-04-03 20:28:47 UTC
+++ go.mod
@@ -1,13 +1,17 @@
 module github.com/Snawoot/dumbproxy
 
-go 1.13
+go 1.20
 
 require (
-	github.com/GehirnInc/crypt v0.0.0-20230320061759-8cc1b52080c5 // indirect
-	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
-	github.com/kr/pretty v0.3.1 // indirect
+	github.com/coreos/go-systemd/v22 v22.5.0
 	github.com/tg123/go-htpasswd v1.2.1
 	golang.org/x/crypto v0.7.0
 	golang.org/x/net v0.8.0
-	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
+)
+
+require (
+	github.com/GehirnInc/crypt v0.0.0-20230320061759-8cc1b52080c5 // indirect
+	golang.org/x/sys v0.6.0 // indirect
+	golang.org/x/term v0.6.0 // indirect
+	golang.org/x/text v0.8.0 // indirect
 )
