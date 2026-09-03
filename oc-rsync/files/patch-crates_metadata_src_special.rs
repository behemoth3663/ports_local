--- crates/metadata/src/special.rs.orig	2026-07-18 10:25:51 UTC
+++ crates/metadata/src/special.rs
@@ -220,7 +220,7 @@ fn create_fifo_parts_inner(
     mode_bits: u32,
     is_socket: bool,
 ) -> Result<(), MetadataError> {
-    use nix::sys::stat::{Mode, SFlag, makedev, mknod};
+    use nix::sys::stat::{Mode, SFlag, mknod};
 
     let (kind, context) = if is_socket {
         // Linux defines MKNOD_CREATES_SOCKETS, so upstream materialises socket
@@ -235,7 +235,7 @@ fn create_fifo_parts_inner(
 
     // `nix::sys::stat::mknod` wraps the libc `mknod` symbol, so an LD_PRELOAD
     // interposer such as fakeroot can fake CAP_MKNOD; see mknod note below.
-    mknod(destination, kind, perm, makedev(0, 0))
+    mknod(destination, kind, perm, libc::makedev(0, 0))
         .map_err(|error| MetadataError::new(context, destination, io::Error::from(error)))
 }
 
