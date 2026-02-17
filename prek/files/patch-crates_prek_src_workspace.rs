--- crates/prek/src/workspace.rs.orig	2026-02-15 13:20:37 UTC
+++ crates/prek/src/workspace.rs
@@ -121,7 +121,13 @@ impl Project {
             "Loading project configuration"
         );
 
-        let mut config = read_config(&config_path)?;
+        let mut config = match read_config(&config_path) {
+            Err(config::Error::Io(e)) if e.kind() == std::io::ErrorKind::NotFound => {
+                return Err(Error::MissingConfigFile);
+            }
+            other => other?,
+        };
+
         let size = config.repos.len();
 
         let config_dir = config_path
