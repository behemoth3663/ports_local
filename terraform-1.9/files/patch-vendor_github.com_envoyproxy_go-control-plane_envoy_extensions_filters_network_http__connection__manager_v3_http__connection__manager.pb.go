--- vendor/github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3/http_connection_manager.pb.go.orig	2026-01-13 06:26:49 UTC
+++ vendor/github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3/http_connection_manager.pb.go
@@ -13,9 +13,7 @@ import (
 	v31 "github.com/envoyproxy/go-control-plane/envoy/config/accesslog/v3"
 	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
 	v33 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
-	v36 "github.com/envoyproxy/go-control-plane/envoy/config/trace/v3"
 	v37 "github.com/envoyproxy/go-control-plane/envoy/type/http/v3"
-	v35 "github.com/envoyproxy/go-control-plane/envoy/type/tracing/v3"
 	v34 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
 	_ "github.com/envoyproxy/protoc-gen-validate/validate"
 	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
@@ -2108,10 +2106,10 @@ type HttpConnectionManager_Tracing struct {
 	// Default: 256
 	MaxPathTagLength *wrapperspb.UInt32Value `protobuf:"bytes,7,opt,name=max_path_tag_length,json=maxPathTagLength,proto3" json:"max_path_tag_length,omitempty"`
 	// A list of custom tags with unique tag name to create tags for the active span.
-	CustomTags []*v35.CustomTag `protobuf:"bytes,8,rep,name=custom_tags,json=customTags,proto3" json:"custom_tags,omitempty"`
+	CustomTags []*anypb.Any `protobuf:"bytes,8,rep,name=custom_tags,json=customTags,proto3" json:"custom_tags,omitempty"`
 	// Configuration for an external tracing provider.
 	// If not specified, no tracing will be performed.
-	Provider *v36.Tracing_Http `protobuf:"bytes,9,opt,name=provider,proto3" json:"provider,omitempty"`
+	Provider *anypb.Any `protobuf:"bytes,9,opt,name=provider,proto3" json:"provider,omitempty"`
 	// Create separate tracing span for each upstream request if true. And if this flag is set to true,
 	// the tracing provider will assume that Envoy will be independent hop in the trace chain and may
 	// set span type to client or server based on this flag.
@@ -2221,14 +2219,14 @@ func (x *HttpConnectionManager_Tracing) GetMaxPathTagL
 	return nil
 }
 
-func (x *HttpConnectionManager_Tracing) GetCustomTags() []*v35.CustomTag {
+func (x *HttpConnectionManager_Tracing) GetCustomTags() []*anypb.Any {
 	if x != nil {
 		return x.CustomTags
 	}
 	return nil
 }
 
-func (x *HttpConnectionManager_Tracing) GetProvider() *v36.Tracing_Http {
+func (x *HttpConnectionManager_Tracing) GetProvider() *anypb.Any {
 	if x != nil {
 		return x.Provider
 	}
@@ -3433,8 +3431,8 @@ var file_envoy_extensions_filters_network_http_connect
 	(*anypb.Any)(nil),                                                                   // 45: google.protobuf.Any
 	(*v3.ExtensionConfigSource)(nil),                                                    // 46: envoy.config.core.v3.ExtensionConfigSource
 	(*v34.Percent)(nil),                                                                 // 47: envoy.type.v3.Percent
-	(*v35.CustomTag)(nil),                                                               // 48: envoy.type.tracing.v3.CustomTag
-	(*v36.Tracing_Http)(nil),                                                            // 49: envoy.config.trace.v3.Tracing.Http
+	(*anypb.Any)(nil),                                                                   // 48: envoy.type.tracing.v3.CustomTag
+	(*anypb.Any)(nil),                                                                   // 49: envoy.config.trace.v3.Tracing.Http
 	(*v3.CidrRange)(nil),                                                                // 50: envoy.config.core.v3.CidrRange
 	(*v37.PathTransformation)(nil),                                                      // 51: envoy.type.http.v3.PathTransformation
 }
