--- FileDescriptor.cpp.orig	2017-08-08 16:08:29 UTC
+++ FileDescriptor.cpp
@@ -21,6 +21,16 @@
 #define CAN_OPEN_SYMLINKS 0
 #endif
 
+#if defined(__FreeBSD__)
+  // FreeBSD libprocstat API and required system headers
+  #include <sys/param.h>
+  #include <sys/types.h>
+  #include <sys/queue.h>
+  #include <sys/user.h>
+  #include <sys/sysctl.h>   // defines KERN_PROC_PID
+  #include <libprocstat.h>  // procstat_* API and filestat_list
+#endif
+
 #ifdef _WIN32
 // We declare our own copy here because Ntifs.h is not included in the
 // standard install of the Visual Studio Community compiler.
@@ -425,6 +435,70 @@ w_string FileDescriptor::getOpenedPath() const {
 
   throw std::system_error(errno, std::generic_category(),
                           "readlink for getOpenedPath");
+#elif defined(__FreeBSD__)
+  // FreeBSD: Obtain the opened path using libprocstat(3) only.
+  // We intentionally do NOT rely on /dev/fd with fdescfs -o rdlnk,
+  // because it changes system-wide semantics and may break other programs.
+  // procfs is deprecated on FreeBSD; libprocstat is the supported interface. [man libprocstat(3), procfs(5)]
+  //
+  // Steps:
+  //   1. procstat_open_sysctl() to create a handle.
+  //   2. procstat_getprocs(... KERN_PROC_PID, getpid() ...) to get our kinfo_proc.
+  //   3. procstat_getfiles() to enumerate filestat entries.
+  //   4. Find entry whose fs_fd equals our fd_, and return fs_path if present (vnode only).
+  //
+  // Error policy:
+  //   - If libprocstat cannot provide a path (e.g., not a vnode or missing path),
+  //     throw ENOSYS to indicate that a filesystem path is not available.
+
+  struct procstat* pr = procstat_open_sysctl();
+  if (!pr) {
+    throw std::system_error(errno, std::generic_category(),
+                            "procstat_open_sysctl");
+  }
+
+  unsigned int nprocs = 0;
+  struct kinfo_proc* kp = procstat_getprocs(pr, KERN_PROC_PID, getpid(), &nprocs);
+  if (!kp || nprocs < 1) {
+    int err = errno;
+    if (kp) procstat_freeprocs(pr, kp);
+    procstat_close(pr);
+    throw std::system_error(err ? err : ENOSYS, std::generic_category(),
+                            "procstat_getprocs");
+  }
+
+  struct filestat_list* files = procstat_getfiles(pr, kp, 0 /* mmapped = 0 */);
+  if (!files) {
+    int err = errno;
+    procstat_freeprocs(pr, kp);
+    procstat_close(pr);
+    throw std::system_error(err ? err : ENOSYS, std::generic_category(),
+                            "procstat_getfiles");
+  }
+
+  w_string result;
+  for (auto fst = STAILQ_FIRST(files); fst; fst = STAILQ_NEXT(fst, next)) {
+    if ((int)fst->fs_fd != fd_) {
+      continue;
+    }
+    // For vnode-backed descriptors libprocstat provides fs_path.
+    if (fst->fs_path && fst->fs_path[0] != '\0') {
+      result = w_string(fst->fs_path);
+    }
+    break;
+  }
+
+  procstat_freefiles(pr, files);
+  procstat_freeprocs(pr, kp);
+  procstat_close(pr);
+
+  if (!result.empty()) {
+    return result;
+  }
+
+  // No filesystem path is available (e.g., pipe/socket/anon); report clearly.
+  throw std::system_error(ENOSYS, std::generic_category(),
+                          "getOpenedPath (FreeBSD): no filesystem path available for this descriptor");
 #elif defined(_WIN32)
   std::wstring wchar;
   wchar.resize(WATCHMAN_NAME_MAX);
