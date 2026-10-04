--- cli-plugins/plugin/plugin.go.orig	2026-09-23 09:45:10 UTC
+++ cli-plugins/plugin/plugin.go
@@ -16,10 +16,8 @@ import (
 	"github.com/docker/cli/cli-plugins/socket"
 	"github.com/docker/cli/cli/command"
 	"github.com/docker/cli/cli/connhelper"
-	"github.com/docker/cli/cli/debug"
 	"github.com/moby/moby/client"
 	"github.com/spf13/cobra"
-	"go.opentelemetry.io/otel"
 )
 
 // PersistentPreRunE must be called by any plugin command (or
@@ -89,8 +87,6 @@ func Run(makeCmd func(command.Cli) *cobra.Command, met
 // makeCmd to construct the plugin command, then invokes the plugin command
 // using [RunPlugin].
 func Run(makeCmd func(command.Cli) *cobra.Command, meta metadata.Metadata, ops ...command.CLIOption) {
-	otel.SetErrorHandler(debug.OTELErrorHandler)
-
 	dockerCLI, err := command.NewDockerCli(ops...)
 	if err != nil {
 		_, _ = fmt.Fprintln(os.Stderr, err)
