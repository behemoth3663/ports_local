--- cli/app.go.orig	2025-03-28 15:11:40 UTC
+++ cli/app.go
@@ -97,22 +97,6 @@ func (app *App) RunContext(ctx context.Context, args [
 
 	ctx = app.registerGracefullyShutdown(ctx)
 
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
-
 	ctx = config.WithConfigValues(ctx)
 	// configure engine context
 	ctx = engine.WithEngineValues(ctx)
