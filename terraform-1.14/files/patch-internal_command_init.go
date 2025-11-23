--- internal/command/init.go.orig	2025-11-19 13:33:06 UTC
+++ internal/command/init.go
@@ -18,8 +18,6 @@ import (
 	svchost "github.com/hashicorp/terraform-svchost"
 	"github.com/posener/complete"
 	"github.com/zclconf/go-cty/cty"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 
 	"github.com/hashicorp/terraform/internal/addrs"
 	"github.com/hashicorp/terraform/internal/backend"
@@ -90,11 +88,6 @@ func (c *InitCommand) getModules(ctx context.Context, 
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
@@ -131,10 +124,6 @@ func (c *InitCommand) initCloud(ctx context.Context, r
 }
 
 func (c *InitCommand) initCloud(ctx context.Context, root *configs.Module, extraConfig arguments.FlagNameValueSlice, viewType arguments.ViewType, view views.Init) (be backend.Backend, output bool, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "initialize HCP Terraform")
-	_ = ctx // prevent staticcheck from complaining to avoid a maintenance hazard of having the wrong ctx in scope here
-	defer span.End()
-
 	view.Output(views.InitializingTerraformCloudMessage)
 
 	if len(extraConfig.AllItems()) != 0 {
@@ -160,10 +149,6 @@ func (c *InitCommand) initBackend(ctx context.Context,
 }
 
 func (c *InitCommand) initBackend(ctx context.Context, root *configs.Module, extraConfig arguments.FlagNameValueSlice, viewType arguments.ViewType, view views.Init) (be backend.Backend, output bool, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "initialize backend")
-	_ = ctx // prevent staticcheck from complaining to avoid a maintenance hazard of having the wrong ctx in scope here
-	defer span.End()
-
 	if root.StateStore != nil {
 		view.Output(views.InitializingStateStoreMessage)
 	} else {
@@ -368,9 +353,6 @@ func (c *InitCommand) getProviders(ctx context.Context
 //
 // This method outputs to the provided view. The returned `output` boolean lets calling code know if anything has been rendered via the view.
 func (c *InitCommand) getProviders(ctx context.Context, config *configs.Config, state *states.State, upgrade bool, pluginDirs []string, flagLockfile string, view views.Init) (output, abort bool, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "install providers")
-	defer span.End()
-
 	// Dev overrides cause the result of "terraform init" to be irrelevant for
 	// any overridden providers, so we'll warn about it to avoid later
 	// confusion when Terraform ends up using a different provider than the
@@ -820,9 +802,6 @@ func (c *InitCommand) getProvidersFromConfig(ctx conte
 // dependency lock data based on the configuration.
 // The dependency lock file itself isn't updated here.
 func (c *InitCommand) getProvidersFromConfig(ctx context.Context, config *configs.Config, upgrade bool, pluginDirs []string, flagLockfile string, view views.Init) (output bool, resultingLocks *depsfile.Locks, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "install providers from config")
-	defer span.End()
-
 	// Dev overrides cause the result of "terraform init" to be irrelevant for
 	// any overridden providers, so we'll warn about it to avoid later
 	// confusion when Terraform ends up using a different provider than the
@@ -921,9 +900,6 @@ func (c *InitCommand) getProvidersFromState(ctx contex
 // supply the configLocks argument.
 // The dependency lock file itself isn't updated here.
 func (c *InitCommand) getProvidersFromState(ctx context.Context, state *states.State, configLocks *depsfile.Locks, upgrade bool, pluginDirs []string, flagLockfile string, view views.Init) (output bool, resultingLocks *depsfile.Locks, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "install providers from state")
-	defer span.End()
-
 	// Dev overrides cause the result of "terraform init" to be irrelevant for
 	// any overridden providers, so we'll warn about it to avoid later
 	// confusion when Terraform ends up using a different provider than the
