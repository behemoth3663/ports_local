--- agent/session/utility/utility_unix.go.orig	2021-01-14 23:52:11 UTC
+++ agent/session/utility/utility_unix.go
@@ -32,11 +32,11 @@ import (
 	"github.com/aws/amazon-ssm-agent/agent/session/utility/model"
 )
 
-var ShellPluginCommandName = "sh"
+var ShellPluginCommandName = "/bin/sh"
 var ShellPluginCommandArgs = []string{"-c"}
 
 const (
-	sudoersFile     = "/etc/sudoers.d/ssm-agent-users"
+	sudoersFile     = "/usr/local/etc/sudoers.d/ssm-agent-users"
 	sudoersFileMode = 0440
 	fs_ioc_getflags = uintptr(0x80086601)
 	fs_ioc_setflags = uintptr(0x40086602)
@@ -77,6 +77,7 @@ func (u *SessionUtil) CreateLocalAdminUser(log log.T) 
 		}
 		// only create sudoers file when user does not exist
 		err = u.createSudoersFileIfNotPresent(log)
+		err = nil
 	}
 
 	return
