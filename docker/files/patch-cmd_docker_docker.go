--- cmd/docker/docker.go.orig	2026-09-23 09:45:10 UTC
+++ cmd/docker/docker.go
@@ -29,7 +29,6 @@ import (
 	"github.com/sirupsen/logrus"
 	"github.com/spf13/cobra"
 	"github.com/spf13/pflag"
-	"go.opentelemetry.io/otel"
 )
 
 type errCtxSignalTerminated struct {
@@ -89,7 +88,6 @@ func dockerMain(ctx context.Context) error {
 		return err
 	}
 	logrus.SetOutput(dockerCli.Err())
-	otel.SetErrorHandler(debug.OTELErrorHandler)
 
 	return runDocker(ctx, dockerCli)
 }
@@ -516,11 +514,11 @@ func runDocker(ctx context.Context, dockerCli *command
 	if mp, ok := mp.(command.MeterProvider); ok {
 		defer func() {
 			if err := mp.Shutdown(ctx); err != nil {
-				otel.Handle(err)
+				debug.LogTelemetryError(err)
 			}
 		}()
 	} else {
-		_, _ = fmt.Fprint(dockerCli.Err(), "Warning: Unexpected OTEL error, metrics may not be flushed")
+		_, _ = fmt.Fprint(dockerCli.Err(), "Warning: telemetry provider unavailable, metrics may not be flushed")
 	}
 
 	dockerCli.InstrumentCobraCommands(ctx, cmd)
