--- pkg/cmd/root/alias.go.orig	2026-10-26 18:02:47 UTC
+++ pkg/cmd/root/alias.go
@@ -46,7 +46,6 @@ func NewCmdShellAlias(io *iostreams.IOStreams, aliasNa
 	// Aliases are user-defined names and must not be reported as telemetry
 	// dimensions, since the name itself may be sensitive (e.g. project or
 	// organization names).
-	cmdutil.DisableTelemetry(cmd)
 	return cmd
 }
 
@@ -73,7 +72,6 @@ func NewCmdAlias(io *iostreams.IOStreams, aliasName, a
 	// Aliases are user-defined names and must not be reported as telemetry
 	// dimensions, since the name itself may be sensitive (e.g. project or
 	// organization names).
-	cmdutil.DisableTelemetry(cmd)
 	return cmd
 }
 
