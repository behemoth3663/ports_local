--- cli/config/credentials/default_store_freebsd.go.orig	2026-06-18 20:17:25 UTC
+++ cli/config/credentials/default_store_freebsd.go
@@ -0,0 +1,13 @@
+package credentials
+
+import (
+       "os/exec"
+)
+
+func defaultCredentialsStore() string {
+       if _, err := exec.LookPath("pass"); err == nil {
+               return "pass"
+       }
+
+       return "secretservice"
+}
