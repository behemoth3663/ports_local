--- internal/runner/run/download_source.go.orig	1979-11-29 21:00:00 UTC
+++ internal/runner/run/download_source.go
@@ -19,7 +19,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/report"
 	"github.com/gruntwork-io/terragrunt/internal/runner/runcfg"
 	"github.com/gruntwork-io/terragrunt/internal/strict/controls"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tf"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
@@ -131,19 +130,13 @@ func DownloadTerraformSource(
 
 		copyOpts := moduleCopyOptions(opts, cfg)
 
-		err = telemetry.TelemeterFromContext(ctx).
-			Collect(ctx, l, "copy_folder_contents", map[string]any{
-				"src":  opts.UnitDir,
-				"dest": terraformSource.WorkingDir,
-			}, func(_ context.Context, l log.Logger) error {
-				return util.CopyFolderContents(
-					l,
-					v.FS,
-					opts.UnitDir,
-					terraformSource.WorkingDir,
-					ModuleManifestName,
-					copyOpts...)
-			})
+		err = util.CopyFolderContents(
+				l,
+				v.FS,
+				opts.UnitDir,
+				terraformSource.WorkingDir,
+				ModuleManifestName,
+				copyOpts...)
 		if err != nil {
 			return nil, err
 		}
@@ -348,7 +341,7 @@ func DownloadTerraformSourceIfNecessary(
 		}
 	}
 
-	var previousVersion = ""
+	previousVersion := ""
 	// read previous source version
 	// https://github.com/gruntwork-io/terragrunt/issues/1921
 	versionFileExists, err := vfs.FileExists(v.FS, terraformSource.VersionFile)
