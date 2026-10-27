--- pkg/cmd/pr/shared/commentable.go.orig	2026-10-26 18:02:47 UTC
+++ pkg/cmd/pr/shared/commentable.go
@@ -63,7 +63,6 @@ type CommentableOptions struct {
 	BodyProvided     bool
 	KeepExistingBody bool
 	AttachFlag       *attachments.Flag
-	AttachEvent      *attachments.TelemetryEvent
 	Assets           []attachments.UserAsset
 }
 
@@ -106,7 +105,6 @@ func CommentablePreRun(cmd *cobra.Command, opts *Comme
 	}
 
 	var err error
-	opts.AttachEvent = attachments.BeginTelemetry(telemetry, cmd.CommandPath(), opts.AttachFlag.Count())
 	opts.Assets, err = opts.AttachFlag.UserAssets()
 	if err != nil {
 		return err
@@ -350,7 +348,6 @@ func bodyForWrite(opts *CommentableOptions, uploader *
 		return opts.Body, true, nil
 	}
 	body, uploadResult, err := uploader.UploadAndAttach(context.Background(), opts.Body, opts.BodyDir, opts.Assets)
-	opts.AttachEvent.RecordOperations(uploadResult)
 	if err != nil && uploadResult.Uploaded == 0 {
 		return "", false, err
 	}
