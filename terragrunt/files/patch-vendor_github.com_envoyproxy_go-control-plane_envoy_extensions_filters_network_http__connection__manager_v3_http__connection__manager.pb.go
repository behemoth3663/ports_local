--- vendor/github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3/http_connection_manager.pb.go.orig	2026-11-26 20:52:46 UTC
+++ vendor/github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3/http_connection_manager.pb.go
@@ -15,7 +15,6 @@ import (
 	v34 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
 	v36 "github.com/envoyproxy/go-control-plane/envoy/config/trace/v3"
 	v37 "github.com/envoyproxy/go-control-plane/envoy/type/http/v3"
-	v35 "github.com/envoyproxy/go-control-plane/envoy/type/tracing/v3"
 	v31 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
 	_ "github.com/envoyproxy/protoc-gen-validate/validate"
 	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
@@ -2275,7 +2274,7 @@ type HttpConnectionManager_Tracing struct {
 	// Default: 256
 	MaxPathTagLength *wrapperspb.UInt32Value `protobuf:"bytes,7,opt,name=max_path_tag_length,json=maxPathTagLength,proto3" json:"max_path_tag_length,omitempty"`
 	// A list of custom tags with unique tag name to create tags for the active span.
-	CustomTags []*v35.CustomTag `protobuf:"bytes,8,rep,name=custom_tags,json=customTags,proto3" json:"custom_tags,omitempty"`
+	CustomTags []*anypb.Any `protobuf:"bytes,8,rep,name=custom_tags,json=customTags,proto3" json:"custom_tags,omitempty"`
 	// Configuration for an external tracing provider.
 	// If not specified, no tracing will be performed.
 	Provider *v36.Tracing_Http `protobuf:"bytes,9,opt,name=provider,proto3" json:"provider,omitempty"`
@@ -2401,7 +2400,7 @@ func (x *HttpConnectionManager_Tracing) GetMaxPathTagL
 	return nil
 }
 
-func (x *HttpConnectionManager_Tracing) GetCustomTags() []*v35.CustomTag {
+func (x *HttpConnectionManager_Tracing) GetCustomTags() []*anypb.Any {
 	if x != nil {
 		return x.CustomTags
 	}
@@ -3660,7 +3659,7 @@ var file_envoy_extensions_filters_network_http_connect
 	(*v34.ScopedRouteConfiguration)(nil),                                                // 47: envoy.config.route.v3.ScopedRouteConfiguration
 	(*anypb.Any)(nil),                                                                   // 48: google.protobuf.Any
 	(*v3.ExtensionConfigSource)(nil),                                                    // 49: envoy.config.core.v3.ExtensionConfigSource
-	(*v35.CustomTag)(nil),                                                               // 50: envoy.type.tracing.v3.CustomTag
+	(*anypb.Any)(nil),                                                                   // 48: envoy.type.tracing.v3.CustomTag
 	(*v36.Tracing_Http)(nil),                                                            // 51: envoy.config.trace.v3.Tracing.Http
 	(*v3.CidrRange)(nil),                                                                // 52: envoy.config.core.v3.CidrRange
 	(*v37.PathTransformation)(nil),                                                      // 53: envoy.type.http.v3.PathTransformation
