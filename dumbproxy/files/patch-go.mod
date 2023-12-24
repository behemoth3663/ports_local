--- go.mod.orig	2023-12-19 11:42:37 UTC
+++ go.mod
@@ -1,15 +1,19 @@ module github.com/Snawoot/dumbproxy
 module github.com/Snawoot/dumbproxy
 
-go 1.13
+go 1.21
 
 require (
-	github.com/GehirnInc/crypt v0.0.0-20230320061759-8cc1b52080c5 // indirect
 	github.com/coreos/go-systemd/v22 v22.5.0
-	github.com/hashicorp/errwrap v1.1.0 // indirect
 	github.com/hashicorp/go-multierror v1.1.1
-	github.com/kr/pretty v0.3.1 // indirect
 	github.com/tg123/go-htpasswd v1.2.2
 	golang.org/x/crypto v0.17.0
 	golang.org/x/net v0.19.0
-	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
+)
+
+require (
+	github.com/GehirnInc/crypt v0.0.0-20230320061759-8cc1b52080c5 // indirect
+	github.com/hashicorp/errwrap v1.1.0 // indirect
+	golang.org/x/sys v0.15.0 // indirect
+	golang.org/x/term v0.15.0 // indirect
+	golang.org/x/text v0.14.0 // indirect
 )
