--- agent/appconfig/constants_unix.go.orig	2020-11-17 21:26:34 UTC
+++ agent/appconfig/constants_unix.go
@@ -28,7 +28,7 @@ import (
 const (
 
 	// AgentExtensions specified the root folder for various kinds of downloaded content
-	AgentData = "/var/lib/amazon/ssm/"
+	AgentData = "/var/run/amazon/ssm/"
 
 	// PackageRoot specifies the directory under which packages will be downloaded and installed
 	PackageRoot = AgentData + "packages"
@@ -61,10 +61,10 @@ const (
 	DefaultDataStorePath = AgentData
 
 	// EC2ConfigDataStorePath represents the directory for storing ec2 config data
-	EC2ConfigDataStorePath = "/var/lib/amazon/ec2config/"
+	EC2ConfigDataStorePath = "/var/run/amazon/ec2config/"
 
 	// EC2ConfigSettingPath represents the directory for storing ec2 config settings
-	EC2ConfigSettingPath = "/var/lib/amazon/ec2configservice/"
+	EC2ConfigSettingPath = "/var/run/amazon/ec2configservice/"
 
 	// UpdaterArtifactsRoot represents the directory for storing update related information
 	UpdaterArtifactsRoot = AgentData + "update/"
@@ -112,9 +112,9 @@ const (
 var PowerShellPluginCommandName string
 
 // DefaultProgramFolder is the default folder for SSM
-var DefaultProgramFolder = "/etc/amazon/ssm/"
+var DefaultProgramFolder = "/usr/local/etc/amazon/ssm/"
 
-var defaultWorkerPath = "/usr/bin/"
+var defaultWorkerPath = "/usr/local/sbin/"
 var DefaultSSMAgentWorker = defaultWorkerPath + "ssm-agent-worker"
 var DefaultDocumentWorker = defaultWorkerPath + "ssm-document-worker"
 var DefaultSessionWorker = defaultWorkerPath + "ssm-session-worker"
@@ -133,9 +133,9 @@ func init() {
 	/*
 	   Powershell command used to be poweshell in alpha versions, now it's pwsh in prod versions
 	*/
-	PowerShellPluginCommandName = "/usr/bin/powershell"
+	PowerShellPluginCommandName = "/usr/local/bin/powershell"
 	if _, err := os.Stat(PowerShellPluginCommandName); err != nil {
-		PowerShellPluginCommandName = "/usr/bin/pwsh"
+		PowerShellPluginCommandName = "/usr/local/bin/pwsh"
 	}
 
 	// curdir is amazon-ssm-agent current directory path
