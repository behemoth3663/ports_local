--- pkg/cmd/root/root.go.orig	2026-10-26 18:02:47 UTC
+++ pkg/cmd/root/root.go
@@ -40,7 +40,6 @@ import (
 	runCmd "github.com/cli/cli/v2/pkg/cmd/run"
 	searchCmd "github.com/cli/cli/v2/pkg/cmd/search"
 	secretCmd "github.com/cli/cli/v2/pkg/cmd/secret"
-	sendTelemetryCmd "github.com/cli/cli/v2/pkg/cmd/send-telemetry"
 	skillsCmd "github.com/cli/cli/v2/pkg/cmd/skills"
 	sshKeyCmd "github.com/cli/cli/v2/pkg/cmd/ssh-key"
 	statusCmd "github.com/cli/cli/v2/pkg/cmd/status"
@@ -157,7 +156,6 @@ func NewCmdRoot(f *cmdutil.Factory, telemetry ghteleme
 	cmd.AddCommand(statusCmd.NewCmdStatus(f, nil))
 	cmd.AddCommand(creditsCmd.NewCmdCredits(f, nil))
 	cmd.AddCommand(licensesCmd.NewCmdLicenses(f))
-	cmd.AddCommand(sendTelemetryCmd.NewCmdSendTelemetry(f))
 
 	// below here at the commands that require the "intelligent" BaseRepo resolver
 	repoResolvingCmdFactory := *f
@@ -250,7 +248,6 @@ func NewCmdRoot(f *cmdutil.Factory, telemetry ghteleme
 	}
 
 	cmdutil.DisableAuthCheck(cmd)
-	cmdutil.RecordTelemetryForSubcommands(cmd, telemetry)
 
 	// The reference command produces paged output that displays information on every other command.
 	// Therefore, we explicitly set the Long text and HelpFunc here after all other commands are registered.
