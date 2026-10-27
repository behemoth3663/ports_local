--- pkg/cmd/pr/edit/edit.go.orig	2026-10-26 18:02:47 UTC
+++ pkg/cmd/pr/edit/edit.go
@@ -42,7 +42,6 @@ type EditOptions struct {
 
 	BodyDir     string
 	AttachFlag  *attachments.Flag
-	AttachEvent *attachments.TelemetryEvent
 	Assets      []attachments.UserAsset
 
 	shared.Editable
@@ -221,7 +220,6 @@ func NewCmdEdit(f *cmdutil.Factory, telemetry ghteleme
 			}
 
 			var err error
-			opts.AttachEvent = attachments.BeginTelemetry(telemetry, cmd.CommandPath(), opts.AttachFlag.Count())
 			opts.Assets, err = opts.AttachFlag.UserAssets()
 			if err != nil {
 				return err
@@ -419,7 +417,6 @@ func editRun(opts *EditOptions) error {
 		// Nothing that can prompt or cancel may follow this.
 		var uploadResult attachments.UploadResult
 		body, uploadResult, uploadErr = uploader.UploadAndAttach(context.Background(), body, opts.BodyDir, opts.Assets)
-		opts.AttachEvent.RecordOperations(uploadResult)
 
 		// With nothing uploaded, even a body the caller typed goes unwritten:
 		// its references are still local paths, which render broken. The other
