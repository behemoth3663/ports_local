--- crates/tirith/src/cli/daemon.rs.orig	2026-09-11 16:12:27 UTC
+++ crates/tirith/src/cli/daemon.rs
@@ -254,7 +254,7 @@ fn peer_euid(fd: std::os::unix::io::RawFd) -> Option<u
     let rc = unsafe {
         libc::getsockopt(
             fd,
-            libc::SOL_LOCAL,
+            0,
             libc::LOCAL_PEERCRED,
             cred.as_mut_ptr() as *mut libc::c_void,
             &mut len,
