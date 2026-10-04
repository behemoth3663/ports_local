--- cli/command/cli.go.orig	2026-09-23 09:45:10 UTC
+++ cli/command/cli.go
@@ -73,7 +73,6 @@ type DockerCli struct {
 	dockerEndpoint     docker.Endpoint
 	contextStoreConfig *store.Config
 	initTimeout        time.Duration
-	res                telemetryResource
 
 	// baseCtx is the base context used for internal operations. In the future
 	// this may be replaced by explicitly passing a context to functions that
