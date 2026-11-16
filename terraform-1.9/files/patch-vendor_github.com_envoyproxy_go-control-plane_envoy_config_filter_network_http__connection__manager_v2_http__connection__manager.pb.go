--- vendor/github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.pb.go.orig	2026-11-16 12:50:00 UTC
+++ vendor/github.com/envoyproxy/go-control-plane/envoy/config/filter/network/http_connection_manager/v2/http_connection_manager.pb.go
@@ -12,7 +12,6 @@ import (
 	v21 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
 	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
 	v2 "github.com/envoyproxy/go-control-plane/envoy/config/filter/accesslog/v2"
-	v23 "github.com/envoyproxy/go-control-plane/envoy/config/trace/v2"
 	_type "github.com/envoyproxy/go-control-plane/envoy/type"
 	v22 "github.com/envoyproxy/go-control-plane/envoy/type/tracing/v2"
 	_ "github.com/envoyproxy/protoc-gen-validate/validate"
@@ -1298,7 +1297,7 @@ type HttpConnectionManager_Tracing struct {
 	CustomTags []*v22.CustomTag `protobuf:"bytes,8,rep,name=custom_tags,json=customTags,proto3" json:"custom_tags,omitempty"`
 	// Configuration for an external tracing provider.
 	// If not specified, no tracing will be performed.
-	Provider      *v23.Tracing_Http `protobuf:"bytes,9,opt,name=provider,proto3" json:"provider,omitempty"`
+	Provider      *anypb.Any `protobuf:"bytes,9,opt,name=provider,proto3" json:"provider,omitempty"`
 	unknownFields protoimpl.UnknownFields
 	sizeCache     protoimpl.SizeCache
 }
@@ -1391,7 +1390,7 @@ func (x *HttpConnectionManager_Tracing) GetCustomTags(
 	return nil
 }
 
-func (x *HttpConnectionManager_Tracing) GetProvider() *v23.Tracing_Http {
+func (x *HttpConnectionManager_Tracing) GetProvider() *anypb.Any {
 	if x != nil {
 		return x.Provider
 	}
@@ -2108,7 +2107,7 @@ var file_envoy_config_filter_network_http_connection_m
 	(*anypb.Any)(nil),                                                                   // 30: google.protobuf.Any
 	(*_type.Percent)(nil),                                                               // 31: envoy.type.Percent
 	(*v22.CustomTag)(nil),                                                               // 32: envoy.type.tracing.v2.CustomTag
-	(*v23.Tracing_Http)(nil),                                                            // 33: envoy.config.trace.v2.Tracing.Http
+	(*anypb.Any)(nil),                                                                   // 33: envoy.config.trace.v2.Tracing.Http
 }
 var file_envoy_config_filter_network_http_connection_manager_v2_http_connection_manager_proto_depIdxs = []int32{
 	0,  // 0: envoy.config.filter.network.http_connection_manager.v2.HttpConnectionManager.codec_type:type_name -> envoy.config.filter.network.http_connection_manager.v2.HttpConnectionManager.CodecType
