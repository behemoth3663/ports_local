--- vendor/github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3/http_connection_manager.pb.go.orig	2026-11-15 15:04:54 UTC
+++ vendor/github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3/http_connection_manager.pb.go
@@ -15,7 +15,6 @@ import (
 	v33 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
 	v36 "github.com/envoyproxy/go-control-plane/envoy/config/trace/v3"
 	v37 "github.com/envoyproxy/go-control-plane/envoy/type/http/v3"
-	v35 "github.com/envoyproxy/go-control-plane/envoy/type/tracing/v3"
 	v34 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
 	_ "github.com/envoyproxy/protoc-gen-validate/validate"
 	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
@@ -2108,7 +2107,7 @@ type HttpConnectionManager_Tracing struct {
 	// Default: 256
 	MaxPathTagLength *wrapperspb.UInt32Value `protobuf:"bytes,7,opt,name=max_path_tag_length,json=maxPathTagLength,proto3" json:"max_path_tag_length,omitempty"`
 	// A list of custom tags with unique tag name to create tags for the active span.
-	CustomTags []*v35.CustomTag `protobuf:"bytes,8,rep,name=custom_tags,json=customTags,proto3" json:"custom_tags,omitempty"`
+	CustomTags []*anypb.Any `protobuf:"bytes,8,rep,name=custom_tags,json=customTags,proto3" json:"custom_tags,omitempty"`
 	// Configuration for an external tracing provider.
 	// If not specified, no tracing will be performed.
 	Provider *v36.Tracing_Http `protobuf:"bytes,9,opt,name=provider,proto3" json:"provider,omitempty"`
@@ -2221,7 +2220,7 @@ func (x *HttpConnectionManager_Tracing) GetMaxPathTagL
 	return nil
 }
 
-func (x *HttpConnectionManager_Tracing) GetCustomTags() []*v35.CustomTag {
+func (x *HttpConnectionManager_Tracing) GetCustomTags() []*anypb.Any {
 	if x != nil {
 		return x.CustomTags
 	}
@@ -3433,7 +3432,7 @@ var file_envoy_extensions_filters_network_http_connect
 	(*anypb.Any)(nil),                                                                   // 45: google.protobuf.Any
 	(*v3.ExtensionConfigSource)(nil),                                                    // 46: envoy.config.core.v3.ExtensionConfigSource
 	(*v34.Percent)(nil),                                                                 // 47: envoy.type.v3.Percent
-	(*v35.CustomTag)(nil),                                                               // 48: envoy.type.tracing.v3.CustomTag
+	(*anypb.Any)(nil),                                                                   // 48: envoy.type.tracing.v3.CustomTag
 	(*v36.Tracing_Http)(nil),                                                            // 49: envoy.config.trace.v3.Tracing.Http
 	(*v3.CidrRange)(nil),                                                                // 50: envoy.config.core.v3.CidrRange
 	(*v37.PathTransformation)(nil),                                                      // 51: envoy.type.http.v3.PathTransformation
