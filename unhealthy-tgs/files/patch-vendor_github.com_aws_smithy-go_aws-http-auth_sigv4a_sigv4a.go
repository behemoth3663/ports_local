--- vendor/github.com/aws/smithy-go/aws-http-auth/sigv4a/sigv4a.go.orig	2026-11-21 14:46:51 UTC
+++ vendor/github.com/aws/smithy-go/aws-http-auth/sigv4a/sigv4a.go
@@ -29,7 +29,6 @@ const algorithm = "AWS4-ECDSA-P256-SHA256"
 const algorithm = "AWS4-ECDSA-P256-SHA256"
 
 
-
 // Signer signs requests with AWS Signature Version 4a.
 //
 // Unlike Sigv4, AWS SigV4a signs requests with an ECDSA private key. This is
