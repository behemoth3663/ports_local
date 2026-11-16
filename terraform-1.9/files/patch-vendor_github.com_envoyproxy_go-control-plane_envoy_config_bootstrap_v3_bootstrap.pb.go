--- vendor/github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3/bootstrap.pb.go.orig	2026-11-16 12:46:35 UTC
+++ vendor/github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v3/bootstrap.pb.go
@@ -15,7 +15,6 @@ import (
 	v37 "github.com/envoyproxy/go-control-plane/envoy/config/listener/v3"
 	v31 "github.com/envoyproxy/go-control-plane/envoy/config/metrics/v3"
 	v33 "github.com/envoyproxy/go-control-plane/envoy/config/overload/v3"
-	v32 "github.com/envoyproxy/go-control-plane/envoy/config/trace/v3"
 	v39 "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
 	v35 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
 	v36 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
@@ -241,7 +240,7 @@ type Bootstrap struct {
 	//	<envoy_v3_api_field_extensions.filters.network.http_connection_manager.v3.HttpConnectionManager.Tracing.provider>`.
 	//
 	// Deprecated: Marked as deprecated in envoy/config/bootstrap/v3/bootstrap.proto.
-	Tracing *v32.Tracing `protobuf:"bytes,9,opt,name=tracing,proto3" json:"tracing,omitempty"`
+	Tracing *structpb.Struct `protobuf:"bytes,9,opt,name=tracing,proto3" json:"tracing,omitempty"`
 	// Configuration for the runtime configuration provider. If not
 	// specified, a “null” provider will be used which will result in all defaults
 	// being used.
@@ -551,7 +550,7 @@ func (x *Bootstrap) GetWatchdogs() *Watchdogs {
 }
 
 // Deprecated: Marked as deprecated in envoy/config/bootstrap/v3/bootstrap.proto.
-func (x *Bootstrap) GetTracing() *v32.Tracing {
+func (x *Bootstrap) GetTracing() *structpb.Struct {
 	if x != nil {
 		return x.Tracing
 	}
@@ -2499,7 +2498,7 @@ var file_envoy_config_bootstrap_v3_bootstrap_proto_goT
 	(*v31.StatsSink)(nil),                            // 27: envoy.config.metrics.v3.StatsSink
 	(*v31.StatsConfig)(nil),                          // 28: envoy.config.metrics.v3.StatsConfig
 	(*durationpb.Duration)(nil),                      // 29: google.protobuf.Duration
-	(*v32.Tracing)(nil),                              // 30: envoy.config.trace.v3.Tracing
+	(*structpb.Struct)(nil),                          // 30: envoy.config.trace.v3.Tracing
 	(*v33.OverloadManager)(nil),                      // 31: envoy.config.overload.v3.OverloadManager
 	(*wrapperspb.UInt64Value)(nil),                   // 32: google.protobuf.UInt64Value
 	(*v3.DnsResolutionConfig)(nil),                   // 33: envoy.config.core.v3.DnsResolutionConfig
