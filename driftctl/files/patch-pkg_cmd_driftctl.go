--- pkg/cmd/driftctl.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/cmd/driftctl.go
@@ -63,8 +63,8 @@ func NewDriftctlCmd(build build.BuildInterface) *Drift
 
 	cmd.PersistentFlags().BoolP("help", "h", false, "Display help for command")
 	if cmd.build.IsUsageReportingEnabled() {
-		cmd.PersistentFlags().BoolP("no-version-check", "", false, "Disable the version check")
-		cmd.PersistentFlags().BoolP("disable-telemetry", "", false, "Disable telemetry")
+		cmd.PersistentFlags().BoolP("no-version-check", "", true, "Disable the version check")
+		cmd.PersistentFlags().BoolP("disable-telemetry", "", true, "Disable telemetry")
 	}
 	cmd.PersistentFlags().BoolP("send-crash-report", "", false, "Enable error reporting. Crash data will be sent to us via Sentry.\nWARNING: may leak sensitive data (please read the documentation for more details)\nThis flag should be used only if an error occurs during execution")
 
