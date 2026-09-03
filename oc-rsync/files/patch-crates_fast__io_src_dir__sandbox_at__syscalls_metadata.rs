--- crates/fast_io/src/dir_sandbox/at_syscalls/metadata.rs.orig	2026-07-18 10:25:51 UTC
+++ crates/fast_io/src/dir_sandbox/at_syscalls/metadata.rs
@@ -124,7 +124,7 @@ pub(super) fn widen_mode(value: libc::mode_t) -> u32 {
 /// `u32` on every supported glibc/musl target.
 #[cfg(not(target_os = "macos"))]
 pub(super) fn widen_mode(value: libc::mode_t) -> u32 {
-    value
+    value.into()
 }
 
 /// Issue `fstatat(dirfd, name, &mut stat, AT_SYMLINK_NOFOLLOW)`.
