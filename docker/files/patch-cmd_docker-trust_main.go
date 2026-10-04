--- cmd/docker-trust/main.go.orig	2026-09-23 09:45:10 UTC
+++ cmd/docker-trust/main.go
@@ -15,7 +15,6 @@ import (
 	"github.com/docker/cli/cli/command"
 	"github.com/docker/cli/cmd/docker-trust/internal/version"
 	"github.com/docker/cli/cmd/docker-trust/trust"
-	"go.opentelemetry.io/otel"
 )
 
 func runStandalone(cmd *command.DockerCli) error {
@@ -35,7 +34,7 @@ func flushMetrics(cmd *command.DockerCli) {
 func flushMetrics(cmd *command.DockerCli) {
 	if mp, ok := cmd.MeterProvider().(command.MeterProvider); ok {
 		if err := mp.ForceFlush(context.Background()); err != nil {
-			otel.Handle(err)
+			_, _ = fmt.Fprintln(os.Stderr, err)
 		}
 	}
 }
