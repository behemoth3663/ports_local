--- vendor/github.com/aws/aws-sdk-go-v2/service/amplifyuibuilder/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/amplifyuibuilder/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpCreateComponent) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpCreateComponent) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -207,10 +203,8 @@ func (m *awsRestjson1_deserializeOpCreateForm) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -250,7 +244,6 @@ func (m *awsRestjson1_deserializeOpCreateForm) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -368,10 +361,8 @@ func (m *awsRestjson1_deserializeOpCreateTheme) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -411,7 +402,6 @@ func (m *awsRestjson1_deserializeOpCreateTheme) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -529,10 +519,8 @@ func (m *awsRestjson1_deserializeOpDeleteComponent) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -550,7 +538,6 @@ func (m *awsRestjson1_deserializeOpDeleteComponent) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -629,10 +616,8 @@ func (m *awsRestjson1_deserializeOpDeleteForm) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -650,7 +635,6 @@ func (m *awsRestjson1_deserializeOpDeleteForm) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -729,10 +713,8 @@ func (m *awsRestjson1_deserializeOpDeleteTheme) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -750,7 +732,6 @@ func (m *awsRestjson1_deserializeOpDeleteTheme) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -829,10 +810,8 @@ func (m *awsRestjson1_deserializeOpExchangeCodeForToke
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -872,7 +851,6 @@ func (m *awsRestjson1_deserializeOpExchangeCodeForToke
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1007,10 +985,8 @@ func (m *awsRestjson1_deserializeOpExportComponents) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1050,7 +1026,6 @@ func (m *awsRestjson1_deserializeOpExportComponents) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1171,10 +1146,8 @@ func (m *awsRestjson1_deserializeOpExportForms) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1214,7 +1187,6 @@ func (m *awsRestjson1_deserializeOpExportForms) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1335,10 +1307,8 @@ func (m *awsRestjson1_deserializeOpExportThemes) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1378,7 +1348,6 @@ func (m *awsRestjson1_deserializeOpExportThemes) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1499,10 +1468,8 @@ func (m *awsRestjson1_deserializeOpGetCodegenJob) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1542,7 +1509,6 @@ func (m *awsRestjson1_deserializeOpGetCodegenJob) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1660,10 +1626,8 @@ func (m *awsRestjson1_deserializeOpGetComponent) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1703,7 +1667,6 @@ func (m *awsRestjson1_deserializeOpGetComponent) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1818,10 +1781,8 @@ func (m *awsRestjson1_deserializeOpGetForm) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1861,7 +1822,6 @@ func (m *awsRestjson1_deserializeOpGetForm) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1976,10 +1936,8 @@ func (m *awsRestjson1_deserializeOpGetMetadata) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2019,7 +1977,6 @@ func (m *awsRestjson1_deserializeOpGetMetadata) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2131,10 +2088,8 @@ func (m *awsRestjson1_deserializeOpGetTheme) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2174,7 +2129,6 @@ func (m *awsRestjson1_deserializeOpGetTheme) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2289,10 +2243,8 @@ func (m *awsRestjson1_deserializeOpListCodegenJobs) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2332,7 +2284,6 @@ func (m *awsRestjson1_deserializeOpListCodegenJobs) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2456,10 +2407,8 @@ func (m *awsRestjson1_deserializeOpListComponents) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2499,7 +2448,6 @@ func (m *awsRestjson1_deserializeOpListComponents) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2620,10 +2568,8 @@ func (m *awsRestjson1_deserializeOpListForms) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2663,7 +2609,6 @@ func (m *awsRestjson1_deserializeOpListForms) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2784,10 +2729,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2827,7 +2770,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2948,10 +2890,8 @@ func (m *awsRestjson1_deserializeOpListThemes) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2991,7 +2931,6 @@ func (m *awsRestjson1_deserializeOpListThemes) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3112,10 +3051,8 @@ func (m *awsRestjson1_deserializeOpPutMetadataFlag) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3133,7 +3070,6 @@ func (m *awsRestjson1_deserializeOpPutMetadataFlag) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3209,10 +3145,8 @@ func (m *awsRestjson1_deserializeOpRefreshToken) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3252,7 +3186,6 @@ func (m *awsRestjson1_deserializeOpRefreshToken) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3378,10 +3311,8 @@ func (m *awsRestjson1_deserializeOpStartCodegenJob) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3421,7 +3352,6 @@ func (m *awsRestjson1_deserializeOpStartCodegenJob) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3536,10 +3466,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3551,7 +3479,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3636,10 +3563,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3651,7 +3576,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3736,10 +3660,8 @@ func (m *awsRestjson1_deserializeOpUpdateComponent) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3779,7 +3701,6 @@ func (m *awsRestjson1_deserializeOpUpdateComponent) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3894,10 +3815,8 @@ func (m *awsRestjson1_deserializeOpUpdateForm) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3937,7 +3856,6 @@ func (m *awsRestjson1_deserializeOpUpdateForm) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4052,10 +3970,8 @@ func (m *awsRestjson1_deserializeOpUpdateTheme) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4095,7 +4011,6 @@ func (m *awsRestjson1_deserializeOpUpdateTheme) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
