--- FileDescriptor.cpp.orig	2017-08-08 16:08:29 UTC
+++ FileDescriptor.cpp
@@ -425,6 +425,17 @@ w_string FileDescriptor::getOpenedPath() const {
 
   throw std::system_error(errno, std::generic_category(),
                           "readlink for getOpenedPath");
+#elif defined(__FreeBSD__)
+  char procpath[64];
+  snprintf(procpath, sizeof(procpath), "/dev/fd/%d", fd_);
+
+  char buf[WATCHMAN_NAME_MAX];
+  ssize_t len = readlink(procpath, buf, sizeof(buf));
+  if (len >= 0 && len < (ssize_t)sizeof(buf)) {
+    return w_string(buf, len);
+  }
+  throw std::system_error(errno ? errno : ENOSYS, std::generic_category(),
+                          "getOpenedPath (/dev/fd): ensure fdescfs mounted with -o rdlnk");
 #elif defined(_WIN32)
   std::wstring wchar;
   wchar.resize(WATCHMAN_NAME_MAX);
