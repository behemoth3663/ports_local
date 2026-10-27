--- pkg/cmd/auth/auth.go.orig	2026-10-26 18:02:47 UTC
+++ pkg/cmd/auth/auth.go
@@ -35,7 +35,6 @@ func NewCmdAuth(f *cmdutil.Factory, telemetry ghteleme
 	cmd.AddCommand(authTokenCmd.NewCmdToken(f, nil))
 	cmd.AddCommand(authSwitchCmd.NewCmdSwitch(f, nil))
 
-	cmdutil.DisableTelemetryForSubcommands(cmd)
 
 	return cmd
 }
