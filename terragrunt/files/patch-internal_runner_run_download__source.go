--- internal/runner/run/download_source.go.orig	1979-11-29 21:00:00 UTC
+++ internal/runner/run/download_source.go
@@ -20,7 +20,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/runner/runcfg"
 	"github.com/gruntwork-io/terragrunt/internal/spinner"
 	"github.com/gruntwork-io/terragrunt/internal/strict/controls"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tf"
 	"github.com/gruntwork-io/terragrunt/internal/util"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
@@ -146,19 +145,14 @@ func DownloadTerraformSource(
 
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
+			l,
+			v.FS,
+			opts.UnitDir,
+			terraformSource.WorkingDir,
+			ModuleManifestName,
+			copyOpts...,
+		)
 		if err != nil {
 			return nil, err
 		}
@@ -640,12 +634,6 @@ func tryCASDownload(
 		}
 
 		l.Warnf("Failed to initialize CAS: %v. Falling back to standard getter.", err)
-		cas.RecordFallback(
-			ctx,
-			l,
-			cas.FallbackReasonInitError,
-			map[string]any{"url": reportedSourceURL},
-		)
 
 		return false, nil
 	}
@@ -656,12 +644,6 @@ func tryCASDownload(
 		}
 
 		l.Warnf("Failed to initialize CAS environment: %v. Falling back to standard getter.", err)
-		cas.RecordFallback(
-			ctx,
-			l,
-			cas.FallbackReasonInitError,
-			map[string]any{"url": reportedSourceURL},
-		)
 
 		return false, nil
 	}
@@ -710,12 +692,6 @@ func tryCASDownload(
 		}
 
 		l.Warnf("CAS download failed: %v. Falling back to standard getter.", err)
-		cas.RecordFallback(
-			ctx,
-			l,
-			cas.FallbackReasonGetterError,
-			map[string]any{"url": reportedSourceURL},
-		)
 
 		// Clear any partial CAS output before the fallback runs; mixing
 		// leftover CAS files with the standard getter's output leaves the
