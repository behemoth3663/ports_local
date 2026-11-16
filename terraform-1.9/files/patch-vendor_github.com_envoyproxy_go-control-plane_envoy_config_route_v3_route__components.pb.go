--- vendor/github.com/envoyproxy/go-control-plane/envoy/config/route/v3/route_components.pb.go.orig	2026-01-13 06:26:49 UTC
+++ vendor/github.com/envoyproxy/go-control-plane/envoy/config/route/v3/route_components.pb.go
@@ -14,7 +14,6 @@ import (
 	v31 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
 	v32 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
 	v36 "github.com/envoyproxy/go-control-plane/envoy/type/metadata/v3"
-	v34 "github.com/envoyproxy/go-control-plane/envoy/type/tracing/v3"
 	v33 "github.com/envoyproxy/go-control-plane/envoy/type/v3"
 	_ "github.com/envoyproxy/protoc-gen-validate/validate"
 	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
@@ -3349,7 +3348,7 @@ type Tracing struct {
 	// configured in the HTTP connection manager. If two tags with the same name are configured
 	// each in the HTTP connection manager and the route level, the one configured here takes
 	// priority.
-	CustomTags []*v34.CustomTag `protobuf:"bytes,4,rep,name=custom_tags,json=customTags,proto3" json:"custom_tags,omitempty"`
+	CustomTags []*anypb.Any `protobuf:"bytes,4,rep,name=custom_tags,json=customTags,proto3" json:"custom_tags,omitempty"`
 	// The operation name of the span which will be used for tracing.
 	//
 	// The same :ref:`format specifier <config_access_log_format>` as used for
@@ -3431,7 +3430,7 @@ func (x *Tracing) GetOverallSampling() *v33.Fractional
 	return nil
 }
 
-func (x *Tracing) GetCustomTags() []*v34.CustomTag {
+func (x *Tracing) GetCustomTags() []*anypb.Any {
 	if x != nil {
 		return x.CustomTags
 	}
@@ -7888,7 +7887,7 @@ var file_envoy_config_route_v3_route_components_proto_
 	(*v33.FractionalPercent)(nil),                       // 83: envoy.type.v3.FractionalPercent
 	(*v31.DataSource)(nil),                              // 84: envoy.config.core.v3.DataSource
 	(*v31.SubstitutionFormatString)(nil),                // 85: envoy.config.core.v3.SubstitutionFormatString
-	(*v34.CustomTag)(nil),                               // 86: envoy.type.tracing.v3.CustomTag
+	(*anypb.Any)(nil),                                   // 86: envoy.type.tracing.v3.CustomTag
 	(*v33.Int64Range)(nil),                              // 87: envoy.type.v3.Int64Range
 	(*v35.HeaderMutation)(nil),                          // 88: envoy.config.common.mutation_rules.v3.HeaderMutation
 	(*v31.ProxyProtocolConfig)(nil),                     // 89: envoy.config.core.v3.ProxyProtocolConfig
