--- vendor/cloud.google.com/go/auth/httptransport/httptransport.go.orig	2026-10-21 09:39:40 UTC
+++ vendor/cloud.google.com/go/auth/httptransport/httptransport.go
@@ -200,10 +200,6 @@ func AddAuthorizationMiddleware(client *http.Client, c
 			base = http.DefaultTransport
 		}
 	}
-	client.Transport = &authTransport{
-		creds: creds,
-		base:  base,
-	}
 	return nil
 }
 
