--- agent/platform/platform_unix.go.orig	2020-11-17 21:26:34 UTC
+++ agent/platform/platform_unix.go
@@ -29,7 +29,7 @@ import (
 )
 
 const (
-	osReleaseFile          = "/etc/os-release"
+	osReleaseFile          = "/var/run/os-release"
 	systemReleaseFile      = "/etc/system-release"
 	centosReleaseFile      = "/etc/centos-release"
 	redhatReleaseFile      = "/etc/redhat-release"
