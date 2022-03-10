--- agent/platform/platform_unix.go.orig	2022-02-28 21:39:37 UTC
+++ agent/platform/platform_unix.go
@@ -30,7 +30,7 @@ import (
 )
 
 const (
-	osReleaseFile           = "/etc/os-release"
+	osReleaseFile           = "/var/run/os-release"
 	systemReleaseFile       = "/etc/system-release"
 	centosReleaseFile       = "/etc/centos-release"
 	redhatReleaseFile       = "/etc/redhat-release"
