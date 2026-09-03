--- crates/branding/build.rs.orig	2026-07-18 10:25:51 UTC
+++ crates/branding/build.rs
@@ -62,6 +62,10 @@ fn workspace_root(manifest_dir: &Path) -> Option<PathB
 }
 
 fn workspace_root(manifest_dir: &Path) -> Option<PathBuf> {
+    if let Ok(root) = env::var("CARGO_WORKSPACE_DIR") {
+        return Some(PathBuf::from(root));
+    }
+
     run_git(manifest_dir, &["rev-parse", "--show-toplevel"]).map(PathBuf::from)
 }
 
