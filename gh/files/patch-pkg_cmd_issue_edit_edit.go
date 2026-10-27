--- pkg/cmd/issue/edit/edit.go.orig	2026-10-26 18:02:47 UTC
+++ pkg/cmd/issue/edit/edit.go
@@ -53,7 +53,6 @@ type EditOptions struct {
 
 	BodyDir     string
 	AttachFlag  *attachments.Flag
-	AttachEvent *attachments.TelemetryEvent
 	Assets      []attachments.UserAsset
 
 	prShared.Editable
@@ -211,7 +210,6 @@ func NewCmdEdit(f *cmdutil.Factory, telemetry ghteleme
 				opts.Editable.IssueType.Edited = true
 			}
 
-			opts.AttachEvent = attachments.BeginTelemetry(telemetry, cmd.CommandPath(), opts.AttachFlag.Count())
 			opts.Assets, err = opts.AttachFlag.UserAssets()
 			if err != nil {
 				return err
@@ -412,7 +410,6 @@ func editRun(opts *EditOptions) error {
 		// prompt or cancel may follow an upload.
 		var uploadResult attachments.UploadResult
 		body, uploadResult, uploadErr = uploader.UploadAndAttach(context.Background(), body, opts.BodyDir, opts.Assets)
-		opts.AttachEvent.RecordOperations(uploadResult)
 
 		if uploadResult.Uploaded > 0 {
 			editable.Body.Value = body
