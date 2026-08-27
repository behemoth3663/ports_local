--- crates/prek/src/workspace.rs.orig	2026-08-27 03:27:58 UTC
+++ crates/prek/src/workspace.rs
@@ -228,7 +228,12 @@ impl Project {
         );
 
         let config_path = std::path::absolute(config_path).map_err(config::Error::from)?;
-        let config = read_config(&config_path)?;
+        let config = match read_config(&config_path) {
+            Err(config::Error::Io(e)) if e.kind() == std::io::ErrorKind::NotFound => {
+                return Err(Error::MissingConfigFile);
+            }
+            other => other?,
+        };
 
         let config_dir = config_path
             .parent()
