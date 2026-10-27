--- pkg/cmd/pr/create/create.go.orig	2026-10-26 18:02:47 UTC
+++ pkg/cmd/pr/create/create.go
@@ -80,7 +80,6 @@ type CreateOptions struct {
 
 	BodyDir     string
 	AttachFlag  *attachments.Flag
-	AttachEvent *attachments.TelemetryEvent
 	Assets      []attachments.UserAsset
 }
 
@@ -371,7 +370,6 @@ func NewCmdCreate(f *cmdutil.Factory, telemetry ghtele
 				return err
 			}
 
-			opts.AttachEvent = attachments.BeginTelemetry(telemetry, cmd.CommandPath(), opts.AttachFlag.Count())
 			opts.Assets, err = opts.AttachFlag.UserAssets()
 			if err != nil {
 				return err
@@ -1121,7 +1119,6 @@ func submitPR(opts CreateOptions, ctx CreateContext, s
 	var uploadErr error
 	if uploader != nil {
 		body, uploadResult, err := uploader.UploadAndAttach(context.Background(), state.Body, opts.BodyDir, opts.Assets)
-		opts.AttachEvent.RecordOperations(uploadResult)
 		// With nothing uploaded, a body that lost the files it was written
 		// around is not what the caller asked to create. The branch is already
 		// pushed by now, so the message says what was not created rather than
