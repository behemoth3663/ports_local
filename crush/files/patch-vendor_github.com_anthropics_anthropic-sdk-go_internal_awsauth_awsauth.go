--- vendor/github.com/anthropics/anthropic-sdk-go/internal/awsauth/awsauth.go.orig	2026-11-13 20:38:10 UTC
+++ vendor/github.com/anthropics/anthropic-sdk-go/internal/awsauth/awsauth.go
@@ -16,7 +16,6 @@ import (
 	v4 "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
 	"github.com/aws/aws-sdk-go-v2/config"
 	"github.com/aws/aws-sdk-go-v2/credentials"
-	"go.opentelemetry.io/otel"
 
 	"github.com/anthropics/anthropic-sdk-go/option"
 )
@@ -207,9 +206,6 @@ func SignHTTP(ctx context.Context, signer *v4.Signer, 
 // while it signs.
 func SignHTTP(ctx context.Context, signer *v4.Signer, creds awssdk.Credentials, r *http.Request, payloadHash, service, region string, signingTime time.Time) error {
 	unsigned := map[string]bool{"traceparent": true, "tracestate": true, "baggage": true}
-	for _, field := range otel.GetTextMapPropagator().Fields() {
-		unsigned[strings.ToLower(field)] = true
-	}
 	held := http.Header{}
 	for name, values := range r.Header {
 		if unsigned[strings.ToLower(name)] {
