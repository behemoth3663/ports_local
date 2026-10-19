--- cargo-crates/libc-0.2.190/src/unix/bsd/freebsdlike/freebsd/mod.rs.orig	2006-07-24 01:21:28 UTC
+++ cargo-crates/libc-0.2.190/src/unix/bsd/freebsdlike/freebsd/mod.rs
@@ -2280,6 +2280,7 @@ pub const NOTE_PDATAMASK: u32 = 0x000fffff;
 pub const NOTE_FORK: u32 = 0x40000000;
 pub const NOTE_EXEC: u32 = 0x20000000;
 pub const NOTE_PDATAMASK: u32 = 0x000fffff;
+pub const NOTE_PCTRLMASK: u32 = 0xf0000000;
 pub const NOTE_TRACK: u32 = 0x00000001;
 pub const NOTE_TRACKERR: u32 = 0x00000002;
 pub const NOTE_CHILD: u32 = 0x00000004;
