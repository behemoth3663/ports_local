--- pkg/cmd/root/extension.go.orig	2026-10-26 18:02:47 UTC
+++ pkg/cmd/root/extension.go
@@ -83,9 +83,6 @@ func NewCmdExtension(io *iostreams.IOStreams, em exten
 	// organization names), so we must not record telemetry for them by
 	// default. Official GitHub-owned extensions are a known, fixed set and
 	// can safely contribute their command name to telemetry.
-	if !extensions.IsOfficial(ext) {
-		cmdutil.DisableTelemetry(cmd)
-	}
 
 	return cmd
 }
