--- cli/app.go.orig	2025-05-02 16:13:41 UTC
+++ cli/app.go
@@ -9,7 +9,6 @@ import (
 
 	"github.com/gruntwork-io/terragrunt/engine"
 	"github.com/gruntwork-io/terragrunt/internal/os/signal"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"golang.org/x/text/cases"
 	"golang.org/x/text/language"
 
@@ -85,22 +84,6 @@ func (app *App) RunContext(ctx context.Context, args [
 	defer cancel()
 
 	ctx = app.registerGracefullyShutdown(ctx)
-
-	if err := global.NewTelemetryFlags(app.opts, nil).Parse(os.Args); err != nil {
-		return err
-	}
-
-	telemeter, err := telemetry.NewTelemeter(ctx, app.Name, app.Version, app.Writer, app.opts.Telemetry)
-	if err != nil {
-		return err
-	}
-	defer func(ctx context.Context) {
-		if err := telemeter.Shutdown(ctx); err != nil {
-			_, _ = app.ErrWriter.Write([]byte(err.Error()))
-		}
-	}(ctx)
-
-	ctx = telemetry.ContextWithTelemeter(ctx, telemeter)
 
 	ctx = config.WithConfigValues(ctx)
 	// configure engine context
