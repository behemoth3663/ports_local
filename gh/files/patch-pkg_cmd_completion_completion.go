--- pkg/cmd/completion/completion.go.orig	2026-10-26 18:02:47 UTC
+++ pkg/cmd/completion/completion.go
@@ -93,7 +93,6 @@ func NewCmdCompletion(io *iostreams.IOStreams) *cobra.
 
 	cmdutil.DisableAuthCheck(cmd)
 	cmdutil.StringEnumFlag(cmd, &shellType, "shell", "s", "", []string{"bash", "zsh", "fish", "powershell"}, "Shell type")
-	cmdutil.DisableTelemetry(cmd)
 
 	return cmd
 }
