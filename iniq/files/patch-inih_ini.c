--- inih/ini.c.orig	2022-05-16 16:29:19 UTC
+++ inih/ini.c
@@ -25,8 +25,8 @@ https://github.com/benhoyt/inih
 #include <stdlib.h>
 #endif
 
-#define MAX_SECTION 50
-#define MAX_NAME 50
+#define MAX_SECTION 100
+#define MAX_NAME 100
 
 /* Used by ini_parse_string() to keep track of string parsing state. */
 typedef struct {
