--- config/stack.go.orig	2025-11-10 20:41:38 UTC
+++ config/stack.go
@@ -9,7 +9,6 @@ import (
 	"strings"
 
 	"github.com/gruntwork-io/terragrunt/pkg/log"
-	"github.com/gruntwork-io/terragrunt/telemetry"
 	"github.com/gruntwork-io/terragrunt/tf"
 
 	"github.com/gruntwork-io/terragrunt/internal/ctyhelper"
@@ -121,14 +120,7 @@ func generateUnits(ctx context.Context, l log.Logger, 
 
 			l.Infof("Generating unit %s from %s", unit.Name, sourceFile)
 
-			return telemetry.TelemeterFromContext(ctx).Collect(ctx, "stack_generate_unit", map[string]any{
-				"stack_file":  sourceFile,
-				"unit_name":   unit.Name,
-				"unit_source": unit.Source,
-				"unit_path":   unit.Path,
-			}, func(ctx context.Context) error {
 				return generateComponent(ctx, l, opts, &item)
-			})
 		})
 	}
 
@@ -154,14 +146,7 @@ func generateStacks(ctx context.Context, l log.Logger,
 
 			l.Infof("Generating stack %s from %s", stack.Name, sourceFile)
 
-			return telemetry.TelemeterFromContext(ctx).Collect(ctx, "stack_generate_stack", map[string]any{
-				"stack_file":   sourceFile,
-				"stack_name":   stack.Name,
-				"stack_source": stack.Source,
-				"stack_path":   stack.Path,
-			}, func(ctx context.Context) error {
 				return generateComponent(ctx, l, opts, &item)
-			})
 		})
 	}
 
