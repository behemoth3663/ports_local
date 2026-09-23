--- pkg/config/stack.go.orig	1979-11-29 21:00:00 UTC
+++ pkg/config/stack.go
@@ -2,6 +2,7 @@ import (
 
 import (
 	"context"
+	"errors"
 	"fmt"
 	"os"
 	"path/filepath"
@@ -15,7 +16,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/git"
 	inthclparse "github.com/gruntwork-io/terragrunt/internal/hclparse"
 	"github.com/gruntwork-io/terragrunt/internal/strict"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tf"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
@@ -33,8 +33,6 @@ import (
 
 	"github.com/zclconf/go-cty/cty"
 
-	"errors"
-
 	"github.com/gruntwork-io/terragrunt/pkg/config/hclparse"
 )
 
@@ -43,8 +41,8 @@ const (
 	StackDir      = inthclparse.StackDir
 	valuesFile    = "terragrunt.values.hcl"
 	manifestName  = ".terragrunt-stack-manifest"
-	unitDirPerm   = 0755
-	valueFilePerm = 0644
+	unitDirPerm   = 0o755
+	valueFilePerm = 0o644
 )
 
 // StackConfigFile represents the structure of terragrunt.stack.hcl stack file.
@@ -582,15 +580,7 @@ func generateUnits(
 				util.RelPathForLog(opts.rootWorkingDir, opts.sourceFile, opts.logShowAbsPaths),
 			)
 
-			return telemetry.TelemeterFromContext(ctx).
-				Collect(ctx, l, "stack_generate_unit", map[string]any{
-					"stack_file":  opts.sourceFile,
-					"unit_name":   unit.Name,
-					"unit_source": unit.Source,
-					"unit_path":   unit.Path,
-				}, func(ctx context.Context, l log.Logger) error {
-					return generateComponent(ctx, l, v, opts, &item)
-				})
+				return generateComponent(ctx, l, v, opts, &item)
 		})
 	}
 
@@ -635,15 +625,7 @@ func generateStacks(
 				util.RelPathForLog(opts.rootWorkingDir, opts.sourceFile, opts.logShowAbsPaths),
 			)
 
-			return telemetry.TelemeterFromContext(ctx).
-				Collect(ctx, l, "stack_generate_stack", map[string]any{
-					"stack_file":   opts.sourceFile,
-					"stack_name":   stack.Name,
-					"stack_source": stack.Source,
-					"stack_path":   stack.Path,
-				}, func(ctx context.Context, l log.Logger) error {
-					return generateComponent(ctx, l, v, opts, &item)
-				})
+				return generateComponent(ctx, l, v, opts, &item)
 		})
 	}
 
