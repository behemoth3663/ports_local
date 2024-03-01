--- internal/command/init.go.orig	2024-02-22 11:32:48 UTC
+++ internal/command/init.go
@@ -15,9 +15,6 @@ import (
 	svchost "github.com/hashicorp/terraform-svchost"
 	"github.com/posener/complete"
 	"github.com/zclconf/go-cty/cty"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 
 	"github.com/opentofu/opentofu/internal/addrs"
 	"github.com/opentofu/opentofu/internal/backend"
@@ -120,46 +117,6 @@ func (c *InitCommand) Run(args []string) int {
 	// to output a newline before the success message
 	var header bool
 
-	if flagFromModule != "" {
-		src := flagFromModule
-
-		empty, err := configs.IsEmptyDir(path)
-		if err != nil {
-			c.Ui.Error(fmt.Sprintf("Error validating destination directory: %s", err))
-			return 1
-		}
-		if !empty {
-			c.Ui.Error(strings.TrimSpace(errInitCopyNotEmpty))
-			return 1
-		}
-
-		c.Ui.Output(c.Colorize().Color(fmt.Sprintf(
-			"[reset][bold]Copying configuration[reset] from %q...", src,
-		)))
-		header = true
-
-		hooks := uiModuleInstallHooks{
-			Ui:             c.Ui,
-			ShowLocalPaths: false, // since they are in a weird location for init
-		}
-
-		ctx, span := tracer.Start(ctx, "-from-module=...", trace.WithAttributes(
-			attribute.String("module_source", src),
-		))
-
-		initDirFromModuleAbort, initDirFromModuleDiags := c.initDirFromModule(ctx, path, src, hooks)
-		diags = diags.Append(initDirFromModuleDiags)
-		if initDirFromModuleAbort || initDirFromModuleDiags.HasErrors() {
-			c.showDiagnostics(diags)
-			span.SetStatus(codes.Error, "module installation failed")
-			span.End()
-			return 1
-		}
-		span.End()
-
-		c.Ui.Output("")
-	}
-
 	// If our directory is empty, then we're done. We can't get or set up
 	// the backend with an empty directory.
 	empty, err := configs.IsEmptyDir(path)
@@ -372,11 +329,6 @@ func (c *InitCommand) getModules(ctx context.Context, 
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
@@ -412,10 +364,6 @@ func (c *InitCommand) initCloud(ctx context.Context, r
 }
 
 func (c *InitCommand) initCloud(ctx context.Context, root *configs.Module, extraConfig rawFlags) (be backend.Backend, output bool, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "initialize cloud backend")
-	_ = ctx // prevent staticcheck from complaining to avoid a maintenence hazard of having the wrong ctx in scope here
-	defer span.End()
-
 	c.Ui.Output(c.Colorize().Color("\n[reset][bold]Initializing cloud backend..."))
 
 	if len(extraConfig.AllItems()) != 0 {
@@ -440,10 +388,6 @@ func (c *InitCommand) initBackend(ctx context.Context,
 }
 
 func (c *InitCommand) initBackend(ctx context.Context, root *configs.Module, extraConfig rawFlags) (be backend.Backend, output bool, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "initialize backend")
-	_ = ctx // prevent staticcheck from complaining to avoid a maintenence hazard of having the wrong ctx in scope here
-	defer span.End()
-
 	c.Ui.Output(c.Colorize().Color("\n[reset][bold]Initializing the backend..."))
 
 	var backendConfig *configs.Backend
@@ -525,9 +469,6 @@ func (c *InitCommand) getProviders(ctx context.Context
 // Load the complete module tree, and fetch any missing providers.
 // This method outputs its own Ui.
 func (c *InitCommand) getProviders(ctx context.Context, config *configs.Config, state *states.State, upgrade bool, pluginDirs []string, flagLockfile string) (output, abort bool, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "install providers")
-	defer span.End()
-
 	// Dev overrides cause the result of "tofu init" to be irrelevant for
 	// any overridden providers, so we'll warn about it to avoid later
 	// confusion when OpenTofu ends up using a different provider than the
