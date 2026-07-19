--- crates/prek/src/workspace.rs.orig	2026-07-16 09:52:05 UTC
+++ crates/prek/src/workspace.rs
@@ -263,7 +263,12 @@ impl Project {
             "Loading project configuration"
         );
 
-        let config = read_config(&config_path)?;
+        let config = match read_config(&config_path) {
+            Err(config::Error::Io(e)) if e.kind() == std::io::ErrorKind::NotFound => {
+                return Err(Error::MissingConfigFile);
+            }
+            other => other?,
+        };
 
         let config_dir = config_path
             .parent()
