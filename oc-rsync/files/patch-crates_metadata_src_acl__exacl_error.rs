--- crates/metadata/src/acl_exacl/error.rs.orig	2026-07-18 10:25:51 UTC
+++ crates/metadata/src/acl_exacl/error.rs
@@ -22,7 +22,7 @@ pub(super) fn is_unsupported_error(e: &io::Error) -> b
     }
 
     match e.raw_os_error() {
-        Some(libc::ENOTSUP) | Some(libc::ENOSYS) | Some(libc::EINVAL) | Some(libc::ENODATA) => {
+        Some(libc::ENOTSUP) | Some(libc::ENOSYS) | Some(libc::EINVAL) | Some(libc::ENOATTR) => {
             return true;
         }
         // upstream: lib/sysacls.c:2780-2782 - macOS reports ENOENT for the
