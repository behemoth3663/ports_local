--- agent/session/utility/utility_unix.go.orig	2022-04-15 17:33:04 UTC
+++ agent/session/utility/utility_unix.go
@@ -33,11 +33,11 @@ import (
 	"github.com/aws/amazon-ssm-agent/agent/session/utility/model"
 )
 
-var ShellPluginCommandName = "sh"
+var ShellPluginCommandName = "/bin/sh"
 var ShellPluginCommandArgs = []string{"-c"}
 
 const (
-	sudoersFile                = "/etc/sudoers.d/ssm-agent-users"
+	sudoersFile                = "/usr/local/etc/sudoers.d/ssm-agent-users"
 	sudoersFileCreateWriteMode = 0640
 	sudoersFileReadOnlyMode    = 0440
 	fs_ioc_getflags            = uintptr(0x80086601)
@@ -79,6 +79,7 @@ func (u *SessionUtil) CreateLocalAdminUser(log log.T) 
 		}
 		// only create sudoers file when user does not exist
 		err = u.createSudoersFileIfNotPresent(log)
+		err = nil
 	}
 
 	return
