--- pkg/cmd/issue/artifact/create/create.go.orig	2026-10-26 18:02:47 UTC
+++ pkg/cmd/issue/artifact/create/create.go
@@ -41,7 +41,6 @@ type CreateOptions struct {
 	Type string
 
 	AttachFlag  *attachments.Flag
-	AttachEvent *attachments.TelemetryEvent
 	Assets      []attachments.UserAsset
 }
 
@@ -219,7 +218,6 @@ func NewCmdCreate(f *cmdutil.Factory, telemetry ghtele
 				return cmdutil.FlagErrorf("`--attach` can't be used with %s artifacts", kind)
 			}
 
-			opts.AttachEvent = attachments.BeginTelemetry(telemetry, cmd.CommandPath(), opts.AttachFlag.Count())
 			opts.Assets, err = opts.AttachFlag.UserAssets()
 			if err != nil {
 				return err
@@ -458,7 +456,6 @@ func attach(opts *CreateOptions, uploader *attachments
 		operations.AppendOperations += r.AppendOperations
 		operations.ReplaceOperations += r.ReplaceOperations
 	}
-	opts.AttachEvent.RecordOperations(operations)
 
 	// Without results, the files were refused before anything uploaded.
 	if len(results) == 0 {
