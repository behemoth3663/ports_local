--- crates/daemon/src/daemon/module_state/connection_limiter.rs.orig	2026-07-18 10:25:51 UTC
+++ crates/daemon/src/daemon/module_state/connection_limiter.rs
@@ -139,7 +139,19 @@ impl ConnectionLimiter {
 ///
 /// upstream: util1.c:634-640 sets `l_type = F_WRLCK`, `l_whence = SEEK_SET`,
 /// `l_start = offset`, `l_len = len`.
-#[cfg(unix)]
+#[cfg(target_os = "freebsd")]
+fn slot_lock(offset: i64) -> nix::libc::flock {
+    nix::libc::flock {
+        l_type: nix::libc::F_WRLCK as nix::libc::c_short,
+        l_whence: nix::libc::SEEK_SET as nix::libc::c_short,
+        l_start: offset as nix::libc::off_t,
+        l_len: SLOT_LEN as nix::libc::off_t,
+        l_pid: 0,
+        l_sysid: 0,
+    }
+}
+
+#[cfg(all(unix, not(target_os = "freebsd")))]
 fn slot_lock(offset: i64) -> nix::libc::flock {
     nix::libc::flock {
         l_type: nix::libc::F_WRLCK as nix::libc::c_short,
