--- node_modules/newman/lib/run/secure-fs.js.orig	1970-01-01 00:00:00 UTC
+++ node_modules/newman/lib/run/secure-fs.js
@@ -143,7 +143,7 @@ Object.getOwnPropertyNames(fs).map((prop) => {
 // Attach all functions in fs to postman-fs
 Object.getOwnPropertyNames(fs).map((prop) => {
     // Bail-out early to prevent fs module from logging warning for deprecated and experimental methods
-    if (prop === DEPRECATED_SYNC_WRITE_STREAM || prop === EXPERIMENTAL_PROMISE || typeof fs[prop] !== FUNCTION) {
+    if (prop === DEPRECATED_SYNC_WRITE_STREAM || prop === EXPERIMENTAL_PROMISE || typeof fs["constants"][prop] !== FUNCTION) {
         return;
     }
 
