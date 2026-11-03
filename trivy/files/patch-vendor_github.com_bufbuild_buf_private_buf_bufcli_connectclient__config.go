--- vendor/github.com/bufbuild/buf/private/buf/bufcli/connectclient_config.go.orig	2026-11-02 21:24:36 UTC
+++ vendor/github.com/bufbuild/buf/private/buf/bufcli/connectclient_config.go
@@ -17,7 +17,6 @@ import (
 import (
 	"buf.build/go/app/appext"
 	"connectrpc.com/connect"
-	otelconnect "connectrpc.com/otelconnect"
 	"github.com/bufbuild/buf/private/buf/bufapp"
 	"github.com/bufbuild/buf/private/bufpkg/bufconnect"
 	"github.com/bufbuild/buf/private/bufpkg/buftransport"
@@ -64,10 +63,6 @@ func newConnectClientConfigWithOptions(container appex
 	if err != nil {
 		return nil, err
 	}
-	otelconnectInterceptor, err := otelconnect.NewInterceptor()
-	if err != nil {
-		return nil, err
-	}
 	client := httpclient.NewClient(config.TLS)
 	options := []connectclient.ConfigOption{
 		connectclient.WithAddressMapper(func(address string) string {
@@ -82,7 +77,6 @@ func newConnectClientConfigWithOptions(container appex
 				bufconnect.NewSetCLIVersionInterceptor(Version),
 				bufconnect.NewCLIWarningInterceptor(container),
 				bufconnect.NewDebugLoggingInterceptor(container),
-				otelconnectInterceptor,
 			},
 		),
 	}
