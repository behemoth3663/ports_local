--- pkg/cmd/issue/create/create.go.orig	2026-10-26 18:02:47 UTC
+++ pkg/cmd/issue/create/create.go
@@ -59,7 +59,6 @@ type CreateOptions struct {
 
 	BodyDir     string
 	AttachFlag  *attachments.Flag
-	AttachEvent *attachments.TelemetryEvent
 	Assets      []attachments.UserAsset
 }
 
@@ -164,7 +163,6 @@ func NewCmdCreate(f *cmdutil.Factory, telemetry ghtele
 				return err
 			}
 
-			opts.AttachEvent = attachments.BeginTelemetry(telemetry, cmd.CommandPath(), opts.AttachFlag.Count())
 			opts.Assets, err = opts.AttachFlag.UserAssets()
 			if err != nil {
 				return err
@@ -458,7 +456,6 @@ func createRun(opts *CreateOptions) (err error) {
 		// issue is created and the failures are reported.
 		if uploader != nil {
 			body, uploadResult, uploadErr := uploader.UploadAndAttach(context.Background(), tb.Body, opts.BodyDir, opts.Assets)
-			opts.AttachEvent.RecordOperations(uploadResult)
 			if uploadErr != nil && uploadResult.Uploaded == 0 {
 				err = uploadErr
 				return
