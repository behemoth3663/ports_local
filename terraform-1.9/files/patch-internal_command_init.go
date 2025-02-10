--- internal/command/init.go.orig	2024-10-16 12:28:59 UTC
+++ internal/command/init.go
@@ -16,9 +16,6 @@ import (
 	svchost "github.com/hashicorp/terraform-svchost"
 	"github.com/posener/complete"
 	"github.com/zclconf/go-cty/cty"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 
 	"github.com/hashicorp/terraform/internal/addrs"
 	"github.com/hashicorp/terraform/internal/backend"
@@ -129,19 +126,12 @@ func (c *InitCommand) Run(args []string) int {
 			View:           view,
 		}
 
-		ctx, span := tracer.Start(ctx, "-from-module=...", trace.WithAttributes(
-			attribute.String("module_source", src),
-		))
-
 		initDirFromModuleAbort, initDirFromModuleDiags := c.initDirFromModule(ctx, path, src, hooks)
 		diags = diags.Append(initDirFromModuleDiags)
 		if initDirFromModuleAbort || initDirFromModuleDiags.HasErrors() {
 			view.Diagnostics(diags)
-			span.SetStatus(codes.Error, "module installation failed")
-			span.End()
 			return 1
 		}
-		span.End()
 
 		view.Output(views.EmptyMessage)
 	}
@@ -348,11 +338,6 @@ func (c *InitCommand) getModules(ctx context.Context, 
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
@@ -389,10 +374,6 @@ func (c *InitCommand) initCloud(ctx context.Context, r
 }
 
 func (c *InitCommand) initCloud(ctx context.Context, root *configs.Module, extraConfig arguments.FlagNameValueSlice, viewType arguments.ViewType, view views.Init) (be backend.Backend, output bool, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "initialize HCP Terraform")
-	_ = ctx // prevent staticcheck from complaining to avoid a maintenence hazard of having the wrong ctx in scope here
-	defer span.End()
-
 	view.Output(views.InitializingTerraformCloudMessage)
 
 	if len(extraConfig.AllItems()) != 0 {
@@ -418,10 +399,6 @@ func (c *InitCommand) initBackend(ctx context.Context,
 }
 
 func (c *InitCommand) initBackend(ctx context.Context, root *configs.Module, extraConfig arguments.FlagNameValueSlice, viewType arguments.ViewType, view views.Init) (be backend.Backend, output bool, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "initialize backend")
-	_ = ctx // prevent staticcheck from complaining to avoid a maintenence hazard of having the wrong ctx in scope here
-	defer span.End()
-
 	view.Output(views.InitializingBackendMessage)
 
 	var backendConfig *configs.Backend
@@ -504,9 +481,6 @@ func (c *InitCommand) getProviders(ctx context.Context
 // Load the complete module tree, and fetch any missing providers.
 // This method outputs its own Ui.
 func (c *InitCommand) getProviders(ctx context.Context, config *configs.Config, state *states.State, upgrade bool, pluginDirs []string, flagLockfile string, view views.Init) (output, abort bool, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "install providers")
-	defer span.End()
-
 	// Dev overrides cause the result of "terraform init" to be irrelevant for
 	// any overridden providers, so we'll warn about it to avoid later
 	// confusion when Terraform ends up using a different provider than the
