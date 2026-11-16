--- vendor/github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2/bootstrap.pb.go.orig	2026-11-16 12:50:00 UTC
+++ vendor/github.com/envoyproxy/go-control-plane/envoy/config/bootstrap/v2/bootstrap.pb.go
@@ -14,7 +14,6 @@ import (
 	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
 	v2 "github.com/envoyproxy/go-control-plane/envoy/config/metrics/v2"
 	v2alpha "github.com/envoyproxy/go-control-plane/envoy/config/overload/v2alpha"
-	v21 "github.com/envoyproxy/go-control-plane/envoy/config/trace/v2"
 	_ "github.com/envoyproxy/protoc-gen-validate/validate"
 	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
 	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
@@ -70,7 +69,7 @@ type Bootstrap struct {
 	//
 	//	This field has been deprecated in favor of :ref:`HttpConnectionManager.Tracing.provider
 	//	<envoy_api_field_config.filter.network.http_connection_manager.v2.HttpConnectionManager.Tracing.provider>`.
-	Tracing *v21.Tracing `protobuf:"bytes,9,opt,name=tracing,proto3" json:"tracing,omitempty"`
+	Tracing *structpb.Struct `protobuf:"bytes,9,opt,name=tracing,proto3" json:"tracing,omitempty"`
 	// Configuration for the runtime configuration provider (deprecated). If not
 	// specified, a “null” provider will be used which will result in all defaults
 	// being used.
@@ -219,7 +218,7 @@ func (x *Bootstrap) GetWatchdog() *Watchdog {
 	return nil
 }
 
-func (x *Bootstrap) GetTracing() *v21.Tracing {
+func (x *Bootstrap) GetTracing() *structpb.Struct {
 	if x != nil {
 		return x.Tracing
 	}
@@ -1275,7 +1274,7 @@ var file_envoy_config_bootstrap_v2_bootstrap_proto_goT
 	(*v2.StatsSink)(nil),                    // 15: envoy.config.metrics.v2.StatsSink
 	(*v2.StatsConfig)(nil),                  // 16: envoy.config.metrics.v2.StatsConfig
 	(*durationpb.Duration)(nil),             // 17: google.protobuf.Duration
-	(*v21.Tracing)(nil),                     // 18: envoy.config.trace.v2.Tracing
+	(*structpb.Struct)(nil),                 // 18: envoy.config.trace.v2.Tracing
 	(*v2alpha.OverloadManager)(nil),         // 19: envoy.config.overload.v2alpha.OverloadManager
 	(*wrapperspb.UInt64Value)(nil),          // 20: google.protobuf.UInt64Value
 	(*core.Address)(nil),                    // 21: envoy.api.v2.core.Address
