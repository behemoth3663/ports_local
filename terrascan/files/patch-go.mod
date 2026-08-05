--- go.mod.orig	2024-09-18 07:31:55 UTC
+++ go.mod
@@ -235,14 +235,20 @@
 	k8s.io/kube-openapi v0.0.0-20231010175941-2dd684a91f00 // indirect
 	k8s.io/utils v0.0.0-20230726121419-3b25d923346b // indirect
 	lukechampine.com/uint128 v1.1.1 // indirect
-	modernc.org/cc/v3 v3.33.6 // indirect
-	modernc.org/ccgo/v3 v3.9.5 // indirect
-	modernc.org/libc v1.9.11 // indirect
-	modernc.org/mathutil v1.4.0 // indirect
-	modernc.org/memory v1.0.4 // indirect
+	modernc.org/cc/v3 v3.35.18 // indirect
+	modernc.org/ccgo/v3 v3.12.95 // indirect
+	modernc.org/libc v1.11.104 // indirect
+	modernc.org/mathutil v1.4.1 // indirect
+	modernc.org/memory v1.0.5 // indirect
 	modernc.org/opt v0.1.1 // indirect
 	modernc.org/strutil v1.1.1 // indirect
 	modernc.org/token v1.0.0 // indirect
 	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
 	sigs.k8s.io/yaml v1.3.0 // indirect
 )
+
+replace github.com/hashicorp/terraform => github.com/hashicorp/terraform v0.15.3
+
+replace modernc.org/libc => modernc.org/libc v1.11.105
+
+replace modernc.org/sqlite => modernc.org/sqlite v1.13.3
