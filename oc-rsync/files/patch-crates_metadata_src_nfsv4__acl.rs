--- crates/metadata/src/nfsv4_acl.rs.orig	2026-07-18 10:25:51 UTC
+++ crates/metadata/src/nfsv4_acl.rs
@@ -488,7 +488,7 @@ pub fn get_nfsv4_acl(path: &Path, follow_symlinks: boo
             let kind = e.kind();
             if kind == io::ErrorKind::NotFound
                 || kind == io::ErrorKind::Unsupported
-                || e.raw_os_error() == Some(libc::ENODATA)
+                || e.raw_os_error() == Some(libc::ENOATTR)
                 || e.raw_os_error() == Some(libc::EOPNOTSUPP)
             {
                 Ok(None)
@@ -533,7 +533,7 @@ pub fn set_nfsv4_acl(
             };
             match result {
                 Ok(()) => Ok(()),
-                Err(e) if e.raw_os_error() == Some(libc::ENODATA) => Ok(()),
+                Err(e) if e.raw_os_error() == Some(libc::ENOATTR) => Ok(()),
                 Err(e) if e.kind() == io::ErrorKind::NotFound => Ok(()),
                 Err(e) => Err(MetadataError::new("remove NFSv4 ACL", path, e)),
             }
