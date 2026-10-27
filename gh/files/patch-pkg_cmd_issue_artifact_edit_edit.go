--- pkg/cmd/issue/artifact/edit/edit.go.orig	2026-10-26 18:02:47 UTC
+++ pkg/cmd/issue/artifact/edit/edit.go
@@ -42,7 +42,6 @@ type EditOptions struct {
 	RestoreVersion int
 
 	AttachFlag  *attachments.Flag
-	AttachEvent *attachments.TelemetryEvent
 	Assets      []attachments.UserAsset
 }
 
@@ -112,7 +111,6 @@ func NewCmdEdit(f *cmdutil.Factory, telemetry ghteleme
 				return cmdutil.FlagErrorf("`--body` cannot be blank")
 			}
 
-			opts.AttachEvent = attachments.BeginTelemetry(telemetry, cmd.CommandPath(), opts.AttachFlag.Count())
 			opts.Assets, err = opts.AttachFlag.UserAssets()
 			if err != nil {
 				return err
@@ -197,7 +195,6 @@ func editRun(opts *EditOptions) error {
 		// Other content has no directory of its own, so its references
 		// resolve against the working directory.
 		md, uploadResult, err := uploader.UploadAndAttach(context.Background(), *body, dir, opts.Assets)
-		opts.AttachEvent.RecordOperations(uploadResult)
 		// Uploads can't be undone, so the artifact is saved if any file
 		// uploaded, even when another failed. It isn't saved when none did.
 		if err != nil && uploadResult.Uploaded == 0 {
