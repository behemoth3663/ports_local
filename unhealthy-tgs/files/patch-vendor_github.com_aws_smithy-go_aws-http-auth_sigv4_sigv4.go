--- vendor/github.com/aws/smithy-go/aws-http-auth/sigv4/sigv4.go.orig	2026-11-21 14:46:51 UTC
+++ vendor/github.com/aws/smithy-go/aws-http-auth/sigv4/sigv4.go
@@ -18,7 +18,6 @@ const algorithm = "AWS4-HMAC-SHA256"
 const algorithm = "AWS4-HMAC-SHA256"
 
 
-
 // Signer signs requests with AWS Signature version 4.
 type Signer struct {
 	options v4.SignerOptions
