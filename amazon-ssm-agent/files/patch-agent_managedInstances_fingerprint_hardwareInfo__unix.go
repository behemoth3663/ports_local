--- agent/managedInstances/fingerprint/hardwareInfo_unix.go.orig	2020-09-25 14:40:04 UTC
+++ agent/managedInstances/fingerprint/hardwareInfo_unix.go
@@ -24,9 +24,9 @@ import (
 )
 
 const (
-	systemDMachineIDPath = "/etc/machine-id"
-	upstartMachineIDPath = "/var/lib/dbus/machine-id"
-	dmidecodeCommand     = "/usr/sbin/dmidecode"
+	systemDMachineIDPath = "/usr/local/etc/machine-id"
+	upstartMachineIDPath = "/var/run/dbus/machine-id"
+	dmidecodeCommand     = "/usr/local/sbin/dmidecode"
 	hardwareID           = "machine-id"
 )
 
