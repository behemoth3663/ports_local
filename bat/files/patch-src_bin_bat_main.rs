--- src/bin/bat/main.rs.orig	2023-10-11 17:14:12 UTC
+++ src/bin/bat/main.rs
@@ -246,7 +246,6 @@ fn invoke_bugreport(app: &App, cache_dir: &Path) {
 
     let mut report = bugreport!()
         .info(SoftwareVersion::default())
-        .info(OperatingSystem::default())
         .info(CommandLine::default())
         .info(EnvironmentVariables::list(&[
             "SHELL",
