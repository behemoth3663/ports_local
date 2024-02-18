--- internal/command/init.go.orig	2024-02-08 01:00:10 UTC
+++ internal/command/init.go
@@ -15,9 +15,6 @@ import (
 	svchost "github.com/hashicorp/terraform-svchost"
 	"github.com/posener/complete"
 	"github.com/zclconf/go-cty/cty"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 
 	"github.com/hashicorp/terraform/internal/addrs"
 	"github.com/hashicorp/terraform/internal/backend"
@@ -142,19 +139,12 @@ func (c *InitCommand) Run(args []string) int {
 			ShowLocalPaths: false, // since they are in a weird location for init
 		}
 
-		ctx, span := tracer.Start(ctx, "-from-module=...", trace.WithAttributes(
-			attribute.String("module_source", src),
-		))
-
 		initDirFromModuleAbort, initDirFromModuleDiags := c.initDirFromModule(ctx, path, src, hooks)
 		diags = diags.Append(initDirFromModuleDiags)
 		if initDirFromModuleAbort || initDirFromModuleDiags.HasErrors() {
 			c.showDiagnostics(diags)
-			span.SetStatus(codes.Error, "module installation failed")
-			span.End()
 			return 1
 		}
-		span.End()
 
 		c.Ui.Output("")
 	}
@@ -359,11 +349,6 @@ func (c *InitCommand) getModules(ctx context.Context, 
 		return false, false, nil
 	}
 
-	ctx, span := tracer.Start(ctx, "install modules", trace.WithAttributes(
-		attribute.Bool("upgrade", upgrade),
-	))
-	defer span.End()
-
 	if upgrade {
 		c.Ui.Output(c.Colorize().Color("[reset][bold]Upgrading modules..."))
 	} else {
@@ -399,10 +384,6 @@ func (c *InitCommand) initCloud(ctx context.Context, r
 }
 
 func (c *InitCommand) initCloud(ctx context.Context, root *configs.Module, extraConfig rawFlags) (be backend.Backend, output bool, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "initialize Terraform Cloud")
-	_ = ctx // prevent staticcheck from complaining to avoid a maintenence hazard of having the wrong ctx in scope here
-	defer span.End()
-
 	c.Ui.Output(c.Colorize().Color("\n[reset][bold]Initializing Terraform Cloud..."))
 
 	if len(extraConfig.AllItems()) != 0 {
@@ -427,10 +408,6 @@ func (c *InitCommand) initBackend(ctx context.Context,
 }
 
 func (c *InitCommand) initBackend(ctx context.Context, root *configs.Module, extraConfig rawFlags) (be backend.Backend, output bool, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "initialize backend")
-	_ = ctx // prevent staticcheck from complaining to avoid a maintenence hazard of having the wrong ctx in scope here
-	defer span.End()
-
 	c.Ui.Output(c.Colorize().Color("\n[reset][bold]Initializing the backend..."))
 
 	var backendConfig *configs.Backend
@@ -512,9 +489,6 @@ func (c *InitCommand) getProviders(ctx context.Context
 // Load the complete module tree, and fetch any missing providers.
 // This method outputs its own Ui.
 func (c *InitCommand) getProviders(ctx context.Context, config *configs.Config, state *states.State, upgrade bool, pluginDirs []string, flagLockfile string) (output, abort bool, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "install providers")
-	defer span.End()
-
 	// Dev overrides cause the result of "terraform init" to be irrelevant for
 	// any overridden providers, so we'll warn about it to avoid later
 	// confusion when Terraform ends up using a different provider than the
