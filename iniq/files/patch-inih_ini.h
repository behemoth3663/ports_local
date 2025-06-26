--- inih/ini.h.orig	2022-05-16 16:29:19 UTC
+++ inih/ini.h
@@ -114,7 +114,7 @@ int ini_parse_string(const char* string, ini_handler h
 /* Maximum line length for any line in INI file (stack or heap). Note that
    this must be 3 more than the longest line (due to '\r', '\n', and '\0'). */
 #ifndef INI_MAX_LINE
-#define INI_MAX_LINE 200
+#define INI_MAX_LINE 2000
 #endif
 
 /* Nonzero to allow heap line buffer to grow via realloc(), zero for a
@@ -127,7 +127,7 @@ int ini_parse_string(const char* string, ini_handler h
 /* Initial size in bytes for heap line buffer. Only applies if INI_USE_STACK
    is zero. */
 #ifndef INI_INITIAL_ALLOC
-#define INI_INITIAL_ALLOC 200
+#define INI_INITIAL_ALLOC 2000
 #endif
 
 /* Stop parsing on first error (default is to keep parsing). */
