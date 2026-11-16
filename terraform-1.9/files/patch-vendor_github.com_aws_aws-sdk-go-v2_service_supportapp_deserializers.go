--- vendor/github.com/aws/aws-sdk-go-v2/service/supportapp/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/supportapp/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpCreateSlackChannelC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -59,7 +56,6 @@ func (m *awsRestjson1_deserializeOpCreateSlackChannelC
 	output := &CreateSlackChannelConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -144,10 +140,8 @@ func (m *awsRestjson1_deserializeOpDeleteAccountAlias)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -159,7 +153,6 @@ func (m *awsRestjson1_deserializeOpDeleteAccountAlias)
 	output := &DeleteAccountAliasOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -238,10 +231,8 @@ func (m *awsRestjson1_deserializeOpDeleteSlackChannelC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -253,7 +244,6 @@ func (m *awsRestjson1_deserializeOpDeleteSlackChannelC
 	output := &DeleteSlackChannelConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -338,10 +328,8 @@ func (m *awsRestjson1_deserializeOpDeleteSlackWorkspac
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -353,7 +341,6 @@ func (m *awsRestjson1_deserializeOpDeleteSlackWorkspac
 	output := &DeleteSlackWorkspaceConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -438,10 +425,8 @@ func (m *awsRestjson1_deserializeOpGetAccountAlias) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -481,7 +466,6 @@ func (m *awsRestjson1_deserializeOpGetAccountAlias) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -594,10 +578,8 @@ func (m *awsRestjson1_deserializeOpListSlackChannelCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -637,7 +619,6 @@ func (m *awsRestjson1_deserializeOpListSlackChannelCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -758,10 +739,8 @@ func (m *awsRestjson1_deserializeOpListSlackWorkspaceC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -801,7 +780,6 @@ func (m *awsRestjson1_deserializeOpListSlackWorkspaceC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -922,10 +900,8 @@ func (m *awsRestjson1_deserializeOpPutAccountAlias) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -937,7 +913,6 @@ func (m *awsRestjson1_deserializeOpPutAccountAlias) Ha
 	output := &PutAccountAliasOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1016,10 +991,8 @@ func (m *awsRestjson1_deserializeOpRegisterSlackWorksp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1059,7 +1032,6 @@ func (m *awsRestjson1_deserializeOpRegisterSlackWorksp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1202,10 +1174,8 @@ func (m *awsRestjson1_deserializeOpUpdateSlackChannelC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1245,7 +1215,6 @@ func (m *awsRestjson1_deserializeOpUpdateSlackChannelC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
