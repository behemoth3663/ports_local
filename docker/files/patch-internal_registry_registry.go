--- internal/registry/registry.go.orig	2026-09-23 09:45:10 UTC
+++ internal/registry/registry.go
@@ -17,7 +17,6 @@ import (
 	"github.com/containerd/log"
 	"github.com/docker/distribution/registry/client/transport"
 	"github.com/docker/go-connections/tlsconfig"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
 )
 
 // hostCertsDir returns the config directory for a specific host.
@@ -134,17 +133,15 @@ func newTransport(tlsConfig *tls.Config) http.RoundTri
 		tlsConfig = tlsconfig.ServerDefault()
 	}
 
-	return otelhttp.NewTransport(
-		&http.Transport{
-			Proxy: http.ProxyFromEnvironment,
-			DialContext: (&net.Dialer{
-				Timeout:   30 * time.Second,
-				KeepAlive: 30 * time.Second,
-			}).DialContext,
-			TLSHandshakeTimeout: 10 * time.Second,
-			TLSClientConfig:     tlsConfig,
-			// TODO(dmcgowan): Call close idle connections when complete and use keep alive
-			DisableKeepAlives: true,
-		},
-	)
+	return &http.Transport{
+		Proxy: http.ProxyFromEnvironment,
+		DialContext: (&net.Dialer{
+			Timeout:   30 * time.Second,
+			KeepAlive: 30 * time.Second,
+		}).DialContext,
+		TLSHandshakeTimeout: 10 * time.Second,
+		TLSClientConfig:     tlsConfig,
+		// TODO(dmcgowan): Call close idle connections when complete and use keep alive
+		DisableKeepAlives: true,
+	}
 }
