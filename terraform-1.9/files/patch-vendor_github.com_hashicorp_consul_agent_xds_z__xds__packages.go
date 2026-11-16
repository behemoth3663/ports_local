--- vendor/github.com/hashicorp/consul/agent/xds/z_xds_packages.go.orig	2026-11-16 12:46:35 UTC
+++ vendor/github.com/hashicorp/consul/agent/xds/z_xds_packages.go
@@ -118,9 +118,6 @@ import (
 	_ "github.com/envoyproxy/go-control-plane/envoy/config/retry/previous_priorities"
 	_ "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
 	_ "github.com/envoyproxy/go-control-plane/envoy/config/tap/v3"
-	_ "github.com/envoyproxy/go-control-plane/envoy/config/trace/v2"
-	_ "github.com/envoyproxy/go-control-plane/envoy/config/trace/v2alpha"
-	_ "github.com/envoyproxy/go-control-plane/envoy/config/trace/v3"
 	_ "github.com/envoyproxy/go-control-plane/envoy/config/transport_socket/alts/v2alpha"
 	_ "github.com/envoyproxy/go-control-plane/envoy/config/transport_socket/raw_buffer/v2"
 	_ "github.com/envoyproxy/go-control-plane/envoy/config/transport_socket/tap/v2alpha"
