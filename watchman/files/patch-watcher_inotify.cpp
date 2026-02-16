--- watcher/inotify.cpp.orig	2025-08-23 19:27:40 UTC
+++ watcher/inotify.cpp
@@ -16,6 +16,7 @@ using watchman::inotify_category;
 #ifndef IN_EXCL_UNLINK
 /* defined in <linux/inotify.h> but we can't include that without
  * breaking userspace */
+#include <sys/inotify.h> // from the libinotify package
 # define WATCHMAN_IN_EXCL_UNLINK 0x04000000
 #else
 # define WATCHMAN_IN_EXCL_UNLINK IN_EXCL_UNLINK
