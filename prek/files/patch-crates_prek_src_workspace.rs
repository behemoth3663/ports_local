--- crates/prek/src/workspace.rs.orig	2026-06-15 11:20:08 UTC
+++ crates/prek/src/workspace.rs
@@ -263,7 +263,12 @@ impl Project {
             "Loading project configuration"
         );
 
-        let mut config = read_config(&config_path)?;
+        let mut config = match read_config(&config_path) {
+            Err(config::Error::Io(e)) if e.kind() == std::io::ErrorKind::NotFound => {
+                return Err(Error::MissingConfigFile);
+            }
+            other => other?,
+        };
 
         let config_dir = config_path
             .parent()
