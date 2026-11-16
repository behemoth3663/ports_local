--- vendor/github.com/aws/aws-sdk-go-v2/service/sesv2/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/sesv2/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpBatchGetMetricData)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpBatchGetMetricData)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -211,10 +207,8 @@ func (m *awsRestjson1_deserializeOpCancelExportJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -226,7 +220,6 @@ func (m *awsRestjson1_deserializeOpCancelExportJob) Ha
 	output := &CancelExportJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -305,10 +298,8 @@ func (m *awsRestjson1_deserializeOpCreateConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -320,7 +311,6 @@ func (m *awsRestjson1_deserializeOpCreateConfiguration
 	output := &CreateConfigurationSetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -408,10 +398,8 @@ func (m *awsRestjson1_deserializeOpCreateConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -423,7 +411,6 @@ func (m *awsRestjson1_deserializeOpCreateConfiguration
 	output := &CreateConfigurationSetEventDestinationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -508,10 +495,8 @@ func (m *awsRestjson1_deserializeOpCreateContact) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -523,7 +508,6 @@ func (m *awsRestjson1_deserializeOpCreateContact) Hand
 	output := &CreateContactOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -605,10 +589,8 @@ func (m *awsRestjson1_deserializeOpCreateContactList) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -620,7 +602,6 @@ func (m *awsRestjson1_deserializeOpCreateContactList) 
 	output := &CreateContactListOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -702,10 +683,8 @@ func (m *awsRestjson1_deserializeOpCreateCustomVerific
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -717,7 +696,6 @@ func (m *awsRestjson1_deserializeOpCreateCustomVerific
 	output := &CreateCustomVerificationEmailTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -802,10 +780,8 @@ func (m *awsRestjson1_deserializeOpCreateDedicatedIpPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -817,7 +793,6 @@ func (m *awsRestjson1_deserializeOpCreateDedicatedIpPo
 	output := &CreateDedicatedIpPoolOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -902,10 +877,8 @@ func (m *awsRestjson1_deserializeOpCreateDeliverabilit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -945,7 +918,6 @@ func (m *awsRestjson1_deserializeOpCreateDeliverabilit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1091,10 +1063,8 @@ func (m *awsRestjson1_deserializeOpCreateEmailIdentity
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1134,7 +1104,6 @@ func (m *awsRestjson1_deserializeOpCreateEmailIdentity
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1276,10 +1245,8 @@ func (m *awsRestjson1_deserializeOpCreateEmailIdentity
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1291,7 +1258,6 @@ func (m *awsRestjson1_deserializeOpCreateEmailIdentity
 	output := &CreateEmailIdentityPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1376,10 +1342,8 @@ func (m *awsRestjson1_deserializeOpCreateEmailTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1391,7 +1355,6 @@ func (m *awsRestjson1_deserializeOpCreateEmailTemplate
 	output := &CreateEmailTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1473,10 +1436,8 @@ func (m *awsRestjson1_deserializeOpCreateExportJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1516,7 +1477,6 @@ func (m *awsRestjson1_deserializeOpCreateExportJob) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1638,10 +1598,8 @@ func (m *awsRestjson1_deserializeOpCreateImportJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1681,7 +1639,6 @@ func (m *awsRestjson1_deserializeOpCreateImportJob) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1800,10 +1757,8 @@ func (m *awsRestjson1_deserializeOpCreateMultiRegionEn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1843,7 +1798,6 @@ func (m *awsRestjson1_deserializeOpCreateMultiRegionEn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1974,10 +1928,8 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1989,7 +1941,6 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguration
 	output := &DeleteConfigurationSetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2071,10 +2022,8 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2086,7 +2035,6 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguration
 	output := &DeleteConfigurationSetEventDestinationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2165,10 +2113,8 @@ func (m *awsRestjson1_deserializeOpDeleteContact) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2180,7 +2126,6 @@ func (m *awsRestjson1_deserializeOpDeleteContact) Hand
 	output := &DeleteContactOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2259,10 +2204,8 @@ func (m *awsRestjson1_deserializeOpDeleteContactList) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2274,7 +2217,6 @@ func (m *awsRestjson1_deserializeOpDeleteContactList) 
 	output := &DeleteContactListOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2356,10 +2298,8 @@ func (m *awsRestjson1_deserializeOpDeleteCustomVerific
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2371,7 +2311,6 @@ func (m *awsRestjson1_deserializeOpDeleteCustomVerific
 	output := &DeleteCustomVerificationEmailTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2450,10 +2389,8 @@ func (m *awsRestjson1_deserializeOpDeleteDedicatedIpPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2465,7 +2402,6 @@ func (m *awsRestjson1_deserializeOpDeleteDedicatedIpPo
 	output := &DeleteDedicatedIpPoolOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2547,10 +2483,8 @@ func (m *awsRestjson1_deserializeOpDeleteEmailIdentity
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2562,7 +2496,6 @@ func (m *awsRestjson1_deserializeOpDeleteEmailIdentity
 	output := &DeleteEmailIdentityOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2644,10 +2577,8 @@ func (m *awsRestjson1_deserializeOpDeleteEmailIdentity
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2659,7 +2590,6 @@ func (m *awsRestjson1_deserializeOpDeleteEmailIdentity
 	output := &DeleteEmailIdentityPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2738,10 +2668,8 @@ func (m *awsRestjson1_deserializeOpDeleteEmailTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2753,7 +2681,6 @@ func (m *awsRestjson1_deserializeOpDeleteEmailTemplate
 	output := &DeleteEmailTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2832,10 +2759,8 @@ func (m *awsRestjson1_deserializeOpDeleteMultiRegionEn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2875,7 +2800,6 @@ func (m *awsRestjson1_deserializeOpDeleteMultiRegionEn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2997,10 +2921,8 @@ func (m *awsRestjson1_deserializeOpDeleteSuppressedDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3012,7 +2934,6 @@ func (m *awsRestjson1_deserializeOpDeleteSuppressedDes
 	output := &DeleteSuppressedDestinationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3091,10 +3012,8 @@ func (m *awsRestjson1_deserializeOpGetAccount) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3134,7 +3053,6 @@ func (m *awsRestjson1_deserializeOpGetAccount) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3297,10 +3215,8 @@ func (m *awsRestjson1_deserializeOpGetBlacklistReports
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3340,7 +3256,6 @@ func (m *awsRestjson1_deserializeOpGetBlacklistReports
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3455,10 +3370,8 @@ func (m *awsRestjson1_deserializeOpGetConfigurationSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3498,7 +3411,6 @@ func (m *awsRestjson1_deserializeOpGetConfigurationSet
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3657,10 +3569,8 @@ func (m *awsRestjson1_deserializeOpGetConfigurationSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3700,7 +3610,6 @@ func (m *awsRestjson1_deserializeOpGetConfigurationSet
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3815,10 +3724,8 @@ func (m *awsRestjson1_deserializeOpGetContact) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3858,7 +3765,6 @@ func (m *awsRestjson1_deserializeOpGetContact) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4046,10 +3952,8 @@ func (m *awsRestjson1_deserializeOpGetContactList) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4089,7 +3993,6 @@ func (m *awsRestjson1_deserializeOpGetContactList) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4259,10 +4162,8 @@ func (m *awsRestjson1_deserializeOpGetCustomVerificati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4302,7 +4203,6 @@ func (m *awsRestjson1_deserializeOpGetCustomVerificati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4466,10 +4366,8 @@ func (m *awsRestjson1_deserializeOpGetDedicatedIp) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4509,7 +4407,6 @@ func (m *awsRestjson1_deserializeOpGetDedicatedIp) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4624,10 +4521,8 @@ func (m *awsRestjson1_deserializeOpGetDedicatedIpPool)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4667,7 +4562,6 @@ func (m *awsRestjson1_deserializeOpGetDedicatedIpPool)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4782,10 +4676,8 @@ func (m *awsRestjson1_deserializeOpGetDedicatedIps) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4825,7 +4717,6 @@ func (m *awsRestjson1_deserializeOpGetDedicatedIps) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4949,10 +4840,8 @@ func (m *awsRestjson1_deserializeOpGetDeliverabilityDa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4992,7 +4881,6 @@ func (m *awsRestjson1_deserializeOpGetDeliverabilityDa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5146,10 +5034,8 @@ func (m *awsRestjson1_deserializeOpGetDeliverabilityTe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5189,7 +5075,6 @@ func (m *awsRestjson1_deserializeOpGetDeliverabilityTe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5328,10 +5213,8 @@ func (m *awsRestjson1_deserializeOpGetDomainDeliverabi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5371,7 +5254,6 @@ func (m *awsRestjson1_deserializeOpGetDomainDeliverabi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5486,10 +5368,8 @@ func (m *awsRestjson1_deserializeOpGetDomainStatistics
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5529,7 +5409,6 @@ func (m *awsRestjson1_deserializeOpGetDomainStatistics
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5649,10 +5528,8 @@ func (m *awsRestjson1_deserializeOpGetEmailIdentity) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5692,7 +5569,6 @@ func (m *awsRestjson1_deserializeOpGetEmailIdentity) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5872,10 +5748,8 @@ func (m *awsRestjson1_deserializeOpGetEmailIdentityPol
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5915,7 +5789,6 @@ func (m *awsRestjson1_deserializeOpGetEmailIdentityPol
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6030,10 +5903,8 @@ func (m *awsRestjson1_deserializeOpGetEmailTemplate) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6073,7 +5944,6 @@ func (m *awsRestjson1_deserializeOpGetEmailTemplate) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6197,10 +6067,8 @@ func (m *awsRestjson1_deserializeOpGetExportJob) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6240,7 +6108,6 @@ func (m *awsRestjson1_deserializeOpGetExportJob) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6429,10 +6296,8 @@ func (m *awsRestjson1_deserializeOpGetImportJob) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6472,7 +6337,6 @@ func (m *awsRestjson1_deserializeOpGetImportJob) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6673,10 +6537,8 @@ func (m *awsRestjson1_deserializeOpGetMessageInsights)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6716,7 +6578,6 @@ func (m *awsRestjson1_deserializeOpGetMessageInsights)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6863,10 +6724,8 @@ func (m *awsRestjson1_deserializeOpGetMultiRegionEndpo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6906,7 +6765,6 @@ func (m *awsRestjson1_deserializeOpGetMultiRegionEndpo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7080,10 +6938,8 @@ func (m *awsRestjson1_deserializeOpGetSuppressedDestin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7123,7 +6979,6 @@ func (m *awsRestjson1_deserializeOpGetSuppressedDestin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7238,10 +7093,8 @@ func (m *awsRestjson1_deserializeOpListConfigurationSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7281,7 +7134,6 @@ func (m *awsRestjson1_deserializeOpListConfigurationSe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7402,10 +7254,8 @@ func (m *awsRestjson1_deserializeOpListContactLists) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7445,7 +7295,6 @@ func (m *awsRestjson1_deserializeOpListContactLists) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7566,10 +7415,8 @@ func (m *awsRestjson1_deserializeOpListContacts) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7609,7 +7456,6 @@ func (m *awsRestjson1_deserializeOpListContacts) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7733,10 +7579,8 @@ func (m *awsRestjson1_deserializeOpListCustomVerificat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7776,7 +7620,6 @@ func (m *awsRestjson1_deserializeOpListCustomVerificat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7897,10 +7740,8 @@ func (m *awsRestjson1_deserializeOpListDedicatedIpPool
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7940,7 +7781,6 @@ func (m *awsRestjson1_deserializeOpListDedicatedIpPool
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8061,10 +7901,8 @@ func (m *awsRestjson1_deserializeOpListDeliverabilityT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8104,7 +7942,6 @@ func (m *awsRestjson1_deserializeOpListDeliverabilityT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8228,10 +8065,8 @@ func (m *awsRestjson1_deserializeOpListDomainDeliverab
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8271,7 +8106,6 @@ func (m *awsRestjson1_deserializeOpListDomainDeliverab
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8395,10 +8229,8 @@ func (m *awsRestjson1_deserializeOpListEmailIdentities
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8438,7 +8270,6 @@ func (m *awsRestjson1_deserializeOpListEmailIdentities
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8559,10 +8390,8 @@ func (m *awsRestjson1_deserializeOpListEmailTemplates)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8602,7 +8431,6 @@ func (m *awsRestjson1_deserializeOpListEmailTemplates)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8723,10 +8551,8 @@ func (m *awsRestjson1_deserializeOpListExportJobs) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8766,7 +8592,6 @@ func (m *awsRestjson1_deserializeOpListExportJobs) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8887,10 +8712,8 @@ func (m *awsRestjson1_deserializeOpListImportJobs) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8930,7 +8753,6 @@ func (m *awsRestjson1_deserializeOpListImportJobs) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9051,10 +8873,8 @@ func (m *awsRestjson1_deserializeOpListMultiRegionEndp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9094,7 +8914,6 @@ func (m *awsRestjson1_deserializeOpListMultiRegionEndp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9215,10 +9034,8 @@ func (m *awsRestjson1_deserializeOpListRecommendations
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9258,7 +9075,6 @@ func (m *awsRestjson1_deserializeOpListRecommendations
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9382,10 +9198,8 @@ func (m *awsRestjson1_deserializeOpListSuppressedDesti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9425,7 +9239,6 @@ func (m *awsRestjson1_deserializeOpListSuppressedDesti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9549,10 +9362,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9592,7 +9403,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9707,10 +9517,8 @@ func (m *awsRestjson1_deserializeOpPutAccountDedicated
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9722,7 +9530,6 @@ func (m *awsRestjson1_deserializeOpPutAccountDedicated
 	output := &PutAccountDedicatedIpWarmupAttributesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9798,10 +9605,8 @@ func (m *awsRestjson1_deserializeOpPutAccountDetails) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9813,7 +9618,6 @@ func (m *awsRestjson1_deserializeOpPutAccountDetails) 
 	output := &PutAccountDetailsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9892,10 +9696,8 @@ func (m *awsRestjson1_deserializeOpPutAccountSendingAt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9907,7 +9709,6 @@ func (m *awsRestjson1_deserializeOpPutAccountSendingAt
 	output := &PutAccountSendingAttributesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9983,10 +9784,8 @@ func (m *awsRestjson1_deserializeOpPutAccountSuppressi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9998,7 +9797,6 @@ func (m *awsRestjson1_deserializeOpPutAccountSuppressi
 	output := &PutAccountSuppressionAttributesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10074,10 +9872,8 @@ func (m *awsRestjson1_deserializeOpPutAccountVdmAttrib
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10089,7 +9885,6 @@ func (m *awsRestjson1_deserializeOpPutAccountVdmAttrib
 	output := &PutAccountVdmAttributesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10165,10 +9960,8 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10180,7 +9973,6 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 	output := &PutConfigurationSetArchivingOptionsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10259,10 +10051,8 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10274,7 +10064,6 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 	output := &PutConfigurationSetDeliveryOptionsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10353,10 +10142,8 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10368,7 +10155,6 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 	output := &PutConfigurationSetReputationOptionsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10447,10 +10233,8 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10462,7 +10246,6 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 	output := &PutConfigurationSetSendingOptionsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10541,10 +10324,8 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10556,7 +10337,6 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 	output := &PutConfigurationSetSuppressionOptionsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10635,10 +10415,8 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10650,7 +10428,6 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 	output := &PutConfigurationSetTrackingOptionsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10729,10 +10506,8 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10744,7 +10519,6 @@ func (m *awsRestjson1_deserializeOpPutConfigurationSet
 	output := &PutConfigurationSetVdmOptionsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10823,10 +10597,8 @@ func (m *awsRestjson1_deserializeOpPutDedicatedIpInPoo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10838,7 +10610,6 @@ func (m *awsRestjson1_deserializeOpPutDedicatedIpInPoo
 	output := &PutDedicatedIpInPoolOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10917,10 +10688,8 @@ func (m *awsRestjson1_deserializeOpPutDedicatedIpPoolS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10932,7 +10701,6 @@ func (m *awsRestjson1_deserializeOpPutDedicatedIpPoolS
 	output := &PutDedicatedIpPoolScalingAttributesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11014,10 +10782,8 @@ func (m *awsRestjson1_deserializeOpPutDedicatedIpWarmu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11029,7 +10795,6 @@ func (m *awsRestjson1_deserializeOpPutDedicatedIpWarmu
 	output := &PutDedicatedIpWarmupAttributesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11108,10 +10873,8 @@ func (m *awsRestjson1_deserializeOpPutDeliverabilityDa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11123,7 +10886,6 @@ func (m *awsRestjson1_deserializeOpPutDeliverabilityDa
 	output := &PutDeliverabilityDashboardOptionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11208,10 +10970,8 @@ func (m *awsRestjson1_deserializeOpPutEmailIdentityCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11223,7 +10983,6 @@ func (m *awsRestjson1_deserializeOpPutEmailIdentityCon
 	output := &PutEmailIdentityConfigurationSetAttributesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11302,10 +11061,8 @@ func (m *awsRestjson1_deserializeOpPutEmailIdentityDki
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11317,7 +11074,6 @@ func (m *awsRestjson1_deserializeOpPutEmailIdentityDki
 	output := &PutEmailIdentityDkimAttributesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11396,10 +11152,8 @@ func (m *awsRestjson1_deserializeOpPutEmailIdentityDki
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11439,7 +11193,6 @@ func (m *awsRestjson1_deserializeOpPutEmailIdentityDki
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11563,10 +11316,8 @@ func (m *awsRestjson1_deserializeOpPutEmailIdentityFee
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11578,7 +11329,6 @@ func (m *awsRestjson1_deserializeOpPutEmailIdentityFee
 	output := &PutEmailIdentityFeedbackAttributesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11657,10 +11407,8 @@ func (m *awsRestjson1_deserializeOpPutEmailIdentityMai
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11672,7 +11420,6 @@ func (m *awsRestjson1_deserializeOpPutEmailIdentityMai
 	output := &PutEmailIdentityMailFromAttributesOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11751,10 +11498,8 @@ func (m *awsRestjson1_deserializeOpPutSuppressedDestin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11766,7 +11511,6 @@ func (m *awsRestjson1_deserializeOpPutSuppressedDestin
 	output := &PutSuppressedDestinationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11842,10 +11586,8 @@ func (m *awsRestjson1_deserializeOpSendBulkEmail) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11885,7 +11627,6 @@ func (m *awsRestjson1_deserializeOpSendBulkEmail) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12015,10 +11756,8 @@ func (m *awsRestjson1_deserializeOpSendCustomVerificat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12058,7 +11797,6 @@ func (m *awsRestjson1_deserializeOpSendCustomVerificat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12189,10 +11927,8 @@ func (m *awsRestjson1_deserializeOpSendEmail) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12232,7 +11968,6 @@ func (m *awsRestjson1_deserializeOpSendEmail) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12366,10 +12101,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12381,7 +12114,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12463,10 +12195,8 @@ func (m *awsRestjson1_deserializeOpTestRenderEmailTemp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12506,7 +12236,6 @@ func (m *awsRestjson1_deserializeOpTestRenderEmailTemp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12625,10 +12354,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12640,7 +12367,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12722,10 +12448,8 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12737,7 +12461,6 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguration
 	output := &UpdateConfigurationSetEventDestinationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12816,10 +12539,8 @@ func (m *awsRestjson1_deserializeOpUpdateContact) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12831,7 +12552,6 @@ func (m *awsRestjson1_deserializeOpUpdateContact) Hand
 	output := &UpdateContactOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12913,10 +12633,8 @@ func (m *awsRestjson1_deserializeOpUpdateContactList) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12928,7 +12646,6 @@ func (m *awsRestjson1_deserializeOpUpdateContactList) 
 	output := &UpdateContactListOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13010,10 +12727,8 @@ func (m *awsRestjson1_deserializeOpUpdateCustomVerific
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13025,7 +12740,6 @@ func (m *awsRestjson1_deserializeOpUpdateCustomVerific
 	output := &UpdateCustomVerificationEmailTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13104,10 +12818,8 @@ func (m *awsRestjson1_deserializeOpUpdateEmailIdentity
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13119,7 +12831,6 @@ func (m *awsRestjson1_deserializeOpUpdateEmailIdentity
 	output := &UpdateEmailIdentityPolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13198,10 +12909,8 @@ func (m *awsRestjson1_deserializeOpUpdateEmailTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13213,7 +12922,6 @@ func (m *awsRestjson1_deserializeOpUpdateEmailTemplate
 	output := &UpdateEmailTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
