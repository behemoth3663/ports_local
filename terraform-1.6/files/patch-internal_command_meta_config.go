--- internal/command/meta_config.go.orig	2023-10-04 16:30:07 UTC
+++ internal/command/meta_config.go
@@ -14,8 +14,6 @@ import (
 	"github.com/hashicorp/hcl/v2/hclsyntax"
 	"github.com/zclconf/go-cty/cty"
 	"github.com/zclconf/go-cty/cty/convert"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 
 	"github.com/hashicorp/terraform/internal/configs"
 	"github.com/hashicorp/terraform/internal/configs/configload"
@@ -225,11 +223,6 @@ func (m *Meta) installModules(ctx context.Context, roo
 // this package has a reasonable implementation for displaying notifications
 // via a provided cli.Ui.
 func (m *Meta) initDirFromModule(ctx context.Context, targetDir string, addr string, hooks initwd.ModuleInstallHooks) (abort bool, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "initialize directory from module", trace.WithAttributes(
-		attribute.String("source_addr", addr),
-	))
-	defer span.End()
-
 	loader, err := m.initConfigLoader()
 	if err != nil {
 		diags = diags.Append(err)
