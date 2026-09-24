--- internal/command/init.go.orig	2026-09-23 11:57:56 UTC
+++ internal/command/init.go
@@ -16,8 +16,6 @@ import (
 	svchost "github.com/hashicorp/terraform-svchost"
 	"github.com/posener/complete"
 	"github.com/zclconf/go-cty/cty"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 
 	"github.com/hashicorp/terraform/internal/addrs"
 	"github.com/hashicorp/terraform/internal/backend"
@@ -92,11 +90,6 @@ func (c *InitCommand) getModules(ctx context.Context, 
 		return false, false, nil
 	}
 
-	ctx, span := tracer.Start(ctx, "install modules", trace.WithAttributes(
-		attribute.Bool("upgrade", upgrade),
-	))
-	defer span.End()
-
 	if upgrade {
 		view.Output(views.UpgradingModulesMessage)
 	} else {
@@ -142,10 +135,6 @@ func (c *InitCommand) initCloud(ctx context.Context, r
 }
 
 func (c *InitCommand) initCloud(ctx context.Context, root *configs.Module, extraConfig arguments.FlagNameValueSlice, viewType arguments.ViewType, view views.Init) (be backend.Backend, output bool, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "initialize HCP Terraform")
-	_ = ctx // prevent staticcheck from complaining to avoid a maintenance hazard of having the wrong ctx in scope here
-	defer span.End()
-
 	view.Output(views.InitializingTerraformCloudMessage)
 
 	if len(extraConfig.AllItems()) != 0 {
@@ -171,10 +160,6 @@ func (c *InitCommand) initBackend(ctx context.Context,
 }
 
 func (c *InitCommand) initBackend(ctx context.Context, root *configs.Module, initArgs *arguments.Init, configLocks *depsfile.Locks, view views.Init) (be backend.Backend, output bool, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "initialize backend")
-	_ = ctx // prevent staticcheck from complaining to avoid a maintenance hazard of having the wrong ctx in scope here
-	defer span.End()
-
 	if root.StateStore != nil {
 		view.Output(views.InitializingStateStoreMessage, root.StateStore.Type)
 	} else {
@@ -383,9 +368,6 @@ func (c *InitCommand) getProvidersFromPSSConfig(ctx co
 //
 // Calling code is responsible for validating inputs to this method, e.g. mutually exclusive flags.
 func (c *InitCommand) getProvidersFromPSSConfig(ctx context.Context, rootModEarly *configs.Module, previousLocks *depsfile.Locks, upgrade bool, pluginDirs []string, flagLockfile string, view views.Init) (output bool, resultingLocks *depsfile.Locks, safeInstallAction SafeStateStoreProviderInstallAction, authResult *getproviders.PackageAuthenticationResult, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "install providers for state store")
-	defer span.End()
-
 	// Dev overrides and unmanaged providers change the installation process in "terraform init";
 	// overridden and unmanaged providers are skipped during installation.
 	// This means that impacted providers won't be downloaded from external sources nor added
@@ -557,9 +539,6 @@ func (c *InitCommand) getProviders(ctx context.Context
 //
 // See getProvidersFromPSSConfig which is equivalent for state store providers.
 func (c *InitCommand) getProviders(ctx context.Context, config *configs.Config, state *states.State, upgrade bool, locks *depsfile.Locks, pluginDirs []string, view views.Init, installerHook providercache.InstallerHook) (output bool, resultingLocks *depsfile.Locks, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "install providers")
-	defer span.End()
-
 	// Dev overrides cause the result of "terraform init" to be irrelevant for
 	// any overridden providers, so we'll warn about it to avoid later
 	// confusion when Terraform ends up using a different provider than the
