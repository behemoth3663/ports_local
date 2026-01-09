--- cmd/checkmake/main.go.orig	2026-01-08 13:03:34 UTC
+++ cmd/checkmake/main.go
@@ -54,8 +54,8 @@ func newRootCmd() *cobra.Command {
 	cmd.PersistentFlags().StringVarP(&output, "output", "o", "text", "Output format: 'text' (default) or 'json' (mutually exclusive with --format)")
 	cmd.MarkFlagsMutuallyExclusive("format", "output")
 
-	cmd.Version = fmt.Sprintf("%s %s built at %s by %s with %s",
-		"checkmake", version, buildTime, builder, goversion)
+	cmd.Version = fmt.Sprintf("%s built at %s by %s with %s",
+		version, buildTime, builder, goversion)
 
 	cmd.AddCommand(&cobra.Command{
 		Use:   "list-rules",
