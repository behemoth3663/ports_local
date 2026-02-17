--- crates/prek/src/config.rs.orig	2026-02-15 13:20:37 UTC
+++ crates/prek/src/config.rs
@@ -1014,6 +1014,37 @@ pub(crate) fn read_config(path: &Path) -> Result<Confi
 
 /// Read the configuration file from the given path, and warn about certain issues.
 pub(crate) fn read_config(path: &Path) -> Result<Config, Error> {
+    use std::{fs, io, io::Read};
+
+    // If this is .pre-commit-config.yaml, we ignore an empty/logically empty file by returning Io(NotFound)
+    // - at the top level this is interpreted as "config is missing"
+    if path
+        .file_name()
+        .and_then(|n| n.to_str())
+        .is_some_and(|n| n == ".pre-commit-config.yaml")
+    {
+        // Empty size?
+        if fs::metadata(path).map(|m| m.len() == 0).unwrap_or(false) {
+            tracing::debug!("Skipping empty YAML config: {}", path.display());
+            return Err(Error::Io(io::Error::from(io::ErrorKind::NotFound)));
+        }
+        // Logically empty: only spaces and/or comments
+        if let Ok(mut f) = fs::File::open(path) {
+            let mut s = String::new();
+            // read_to_string may return an error - it will be handled in the normal parsing below,
+            // so here we ignore it and check what was read
+            let _ = f.read_to_string(&mut s);
+            let logical_empty = s
+                .lines()
+                .filter(|l| !l.trim_start().starts_with('#'))
+                .all(|l| l.trim().is_empty());
+            if logical_empty {
+                tracing::debug!("Skipping logically-empty YAML config: {}", path.display());
+                return Err(Error::Io(io::Error::from(io::ErrorKind::NotFound)));
+            }
+        }
+    }
+
     let config = load_config(path)?;
 
     let unused_paths = collect_unused_paths(&config);
