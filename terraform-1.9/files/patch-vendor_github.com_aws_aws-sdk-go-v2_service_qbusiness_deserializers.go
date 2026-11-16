--- vendor/github.com/aws/aws-sdk-go-v2/service/qbusiness/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/qbusiness/deserializers.go
@@ -19,7 +19,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -50,10 +49,8 @@ func (m *awsRestjson1_deserializeOpAssociatePermission
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -93,7 +90,6 @@ func (m *awsRestjson1_deserializeOpAssociatePermission
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -224,10 +220,8 @@ func (m *awsRestjson1_deserializeOpBatchDeleteDocument
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -267,7 +261,6 @@ func (m *awsRestjson1_deserializeOpBatchDeleteDocument
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -391,10 +384,8 @@ func (m *awsRestjson1_deserializeOpBatchPutDocument) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -434,7 +425,6 @@ func (m *awsRestjson1_deserializeOpBatchPutDocument) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -561,10 +551,8 @@ func (m *awsRestjson1_deserializeOpCancelSubscription)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -604,7 +592,6 @@ func (m *awsRestjson1_deserializeOpCancelSubscription)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -739,10 +726,8 @@ func (m *awsRestjson1_deserializeOpChat) HandleDeseria
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -754,7 +739,6 @@ func (m *awsRestjson1_deserializeOpChat) HandleDeseria
 	output := &ChatOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -848,10 +832,8 @@ func (m *awsRestjson1_deserializeOpChatSync) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -891,7 +873,6 @@ func (m *awsRestjson1_deserializeOpChatSync) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1072,10 +1053,8 @@ func (m *awsRestjson1_deserializeOpCheckDocumentAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1115,7 +1094,6 @@ func (m *awsRestjson1_deserializeOpCheckDocumentAccess
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1255,10 +1233,8 @@ func (m *awsRestjson1_deserializeOpCreateAnonymousWebE
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1298,7 +1274,6 @@ func (m *awsRestjson1_deserializeOpCreateAnonymousWebE
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1426,10 +1401,8 @@ func (m *awsRestjson1_deserializeOpCreateApplication) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1469,7 +1442,6 @@ func (m *awsRestjson1_deserializeOpCreateApplication) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1609,10 +1581,8 @@ func (m *awsRestjson1_deserializeOpCreateChatResponseC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1652,7 +1622,6 @@ func (m *awsRestjson1_deserializeOpCreateChatResponseC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1792,10 +1761,8 @@ func (m *awsRestjson1_deserializeOpCreateDataAccessor)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1835,7 +1802,6 @@ func (m *awsRestjson1_deserializeOpCreateDataAccessor)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1984,10 +1950,8 @@ func (m *awsRestjson1_deserializeOpCreateDataSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2027,7 +1991,6 @@ func (m *awsRestjson1_deserializeOpCreateDataSource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2167,10 +2130,8 @@ func (m *awsRestjson1_deserializeOpCreateIndex) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2210,7 +2171,6 @@ func (m *awsRestjson1_deserializeOpCreateIndex) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2350,10 +2310,8 @@ func (m *awsRestjson1_deserializeOpCreatePlugin) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2393,7 +2351,6 @@ func (m *awsRestjson1_deserializeOpCreatePlugin) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2542,10 +2499,8 @@ func (m *awsRestjson1_deserializeOpCreateRetriever) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2585,7 +2540,6 @@ func (m *awsRestjson1_deserializeOpCreateRetriever) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2725,10 +2679,8 @@ func (m *awsRestjson1_deserializeOpCreateSubscription)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2768,7 +2720,6 @@ func (m *awsRestjson1_deserializeOpCreateSubscription)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2915,10 +2866,8 @@ func (m *awsRestjson1_deserializeOpCreateUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2930,7 +2879,6 @@ func (m *awsRestjson1_deserializeOpCreateUser) HandleD
 	output := &CreateUserOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3021,10 +2969,8 @@ func (m *awsRestjson1_deserializeOpCreateWebExperience
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3064,7 +3010,6 @@ func (m *awsRestjson1_deserializeOpCreateWebExperience
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3204,10 +3149,8 @@ func (m *awsRestjson1_deserializeOpDeleteApplication) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3219,7 +3162,6 @@ func (m *awsRestjson1_deserializeOpDeleteApplication) 
 	output := &DeleteApplicationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3307,10 +3249,8 @@ func (m *awsRestjson1_deserializeOpDeleteAttachment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3322,7 +3262,6 @@ func (m *awsRestjson1_deserializeOpDeleteAttachment) H
 	output := &DeleteAttachmentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3410,10 +3349,8 @@ func (m *awsRestjson1_deserializeOpDeleteChatControlsC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3425,7 +3362,6 @@ func (m *awsRestjson1_deserializeOpDeleteChatControlsC
 	output := &DeleteChatControlsConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3510,10 +3446,8 @@ func (m *awsRestjson1_deserializeOpDeleteChatResponseC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3525,7 +3459,6 @@ func (m *awsRestjson1_deserializeOpDeleteChatResponseC
 	output := &DeleteChatResponseConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3613,10 +3546,8 @@ func (m *awsRestjson1_deserializeOpDeleteConversation)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3628,7 +3559,6 @@ func (m *awsRestjson1_deserializeOpDeleteConversation)
 	output := &DeleteConversationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3716,10 +3646,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataAccessor)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3731,7 +3659,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataAccessor)
 	output := &DeleteDataAccessorOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3819,10 +3746,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3834,7 +3759,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataSource) H
 	output := &DeleteDataSourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3922,10 +3846,8 @@ func (m *awsRestjson1_deserializeOpDeleteGroup) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3937,7 +3859,6 @@ func (m *awsRestjson1_deserializeOpDeleteGroup) Handle
 	output := &DeleteGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4025,10 +3946,8 @@ func (m *awsRestjson1_deserializeOpDeleteIndex) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4040,7 +3959,6 @@ func (m *awsRestjson1_deserializeOpDeleteIndex) Handle
 	output := &DeleteIndexOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4128,10 +4046,8 @@ func (m *awsRestjson1_deserializeOpDeletePlugin) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4143,7 +4059,6 @@ func (m *awsRestjson1_deserializeOpDeletePlugin) Handl
 	output := &DeletePluginOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4231,10 +4146,8 @@ func (m *awsRestjson1_deserializeOpDeleteRetriever) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4246,7 +4159,6 @@ func (m *awsRestjson1_deserializeOpDeleteRetriever) Ha
 	output := &DeleteRetrieverOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4334,10 +4246,8 @@ func (m *awsRestjson1_deserializeOpDeleteUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4349,7 +4259,6 @@ func (m *awsRestjson1_deserializeOpDeleteUser) HandleD
 	output := &DeleteUserOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4437,10 +4346,8 @@ func (m *awsRestjson1_deserializeOpDeleteWebExperience
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4452,7 +4359,6 @@ func (m *awsRestjson1_deserializeOpDeleteWebExperience
 	output := &DeleteWebExperienceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4540,10 +4446,8 @@ func (m *awsRestjson1_deserializeOpDisassociatePermiss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4555,7 +4459,6 @@ func (m *awsRestjson1_deserializeOpDisassociatePermiss
 	output := &DisassociatePermissionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4643,10 +4546,8 @@ func (m *awsRestjson1_deserializeOpGetApplication) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4686,7 +4587,6 @@ func (m *awsRestjson1_deserializeOpGetApplication) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4955,10 +4855,8 @@ func (m *awsRestjson1_deserializeOpGetChatControlsConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4998,7 +4896,6 @@ func (m *awsRestjson1_deserializeOpGetChatControlsConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5157,10 +5054,8 @@ func (m *awsRestjson1_deserializeOpGetChatResponseConf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5200,7 +5095,6 @@ func (m *awsRestjson1_deserializeOpGetChatResponseConf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5369,10 +5263,8 @@ func (m *awsRestjson1_deserializeOpGetDataAccessor) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5412,7 +5304,6 @@ func (m *awsRestjson1_deserializeOpGetDataAccessor) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5624,10 +5515,8 @@ func (m *awsRestjson1_deserializeOpGetDataSource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5667,7 +5556,6 @@ func (m *awsRestjson1_deserializeOpGetDataSource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5930,10 +5818,8 @@ func (m *awsRestjson1_deserializeOpGetGroup) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5973,7 +5859,6 @@ func (m *awsRestjson1_deserializeOpGetGroup) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6102,10 +5987,8 @@ func (m *awsRestjson1_deserializeOpGetIndex) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6145,7 +6028,6 @@ func (m *awsRestjson1_deserializeOpGetIndex) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6376,10 +6258,8 @@ func (m *awsRestjson1_deserializeOpGetMedia) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6419,7 +6299,6 @@ func (m *awsRestjson1_deserializeOpGetMedia) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6563,10 +6442,8 @@ func (m *awsRestjson1_deserializeOpGetPlugin) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6606,7 +6483,6 @@ func (m *awsRestjson1_deserializeOpGetPlugin) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6836,10 +6712,8 @@ func (m *awsRestjson1_deserializeOpGetPolicy) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6879,7 +6753,6 @@ func (m *awsRestjson1_deserializeOpGetPolicy) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7004,10 +6877,8 @@ func (m *awsRestjson1_deserializeOpGetRetriever) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7047,7 +6918,6 @@ func (m *awsRestjson1_deserializeOpGetRetriever) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7263,10 +7133,8 @@ func (m *awsRestjson1_deserializeOpGetUser) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7306,7 +7174,6 @@ func (m *awsRestjson1_deserializeOpGetUser) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7430,10 +7297,8 @@ func (m *awsRestjson1_deserializeOpGetWebExperience) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7473,7 +7338,6 @@ func (m *awsRestjson1_deserializeOpGetWebExperience) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7741,10 +7605,8 @@ func (m *awsRestjson1_deserializeOpListApplications) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7784,7 +7646,6 @@ func (m *awsRestjson1_deserializeOpListApplications) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7911,10 +7772,8 @@ func (m *awsRestjson1_deserializeOpListAttachments) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7954,7 +7813,6 @@ func (m *awsRestjson1_deserializeOpListAttachments) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8087,10 +7945,8 @@ func (m *awsRestjson1_deserializeOpListChatResponseCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8130,7 +7986,6 @@ func (m *awsRestjson1_deserializeOpListChatResponseCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8260,10 +8115,8 @@ func (m *awsRestjson1_deserializeOpListConversations) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8303,7 +8156,6 @@ func (m *awsRestjson1_deserializeOpListConversations) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8436,10 +8288,8 @@ func (m *awsRestjson1_deserializeOpListDataAccessors) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8479,7 +8329,6 @@ func (m *awsRestjson1_deserializeOpListDataAccessors) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8609,10 +8458,8 @@ func (m *awsRestjson1_deserializeOpListDataSources) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8652,7 +8499,6 @@ func (m *awsRestjson1_deserializeOpListDataSources) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8782,10 +8628,8 @@ func (m *awsRestjson1_deserializeOpListDataSourceSyncJ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8825,7 +8669,6 @@ func (m *awsRestjson1_deserializeOpListDataSourceSyncJ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8958,10 +8801,8 @@ func (m *awsRestjson1_deserializeOpListDocuments) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9001,7 +8842,6 @@ func (m *awsRestjson1_deserializeOpListDocuments) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9131,10 +8971,8 @@ func (m *awsRestjson1_deserializeOpListGroups) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9174,7 +9012,6 @@ func (m *awsRestjson1_deserializeOpListGroups) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9307,10 +9144,8 @@ func (m *awsRestjson1_deserializeOpListIndices) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9350,7 +9185,6 @@ func (m *awsRestjson1_deserializeOpListIndices) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9480,10 +9314,8 @@ func (m *awsRestjson1_deserializeOpListMessages) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9523,7 +9355,6 @@ func (m *awsRestjson1_deserializeOpListMessages) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9656,10 +9487,8 @@ func (m *awsRestjson1_deserializeOpListPluginActions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9699,7 +9528,6 @@ func (m *awsRestjson1_deserializeOpListPluginActions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9829,10 +9657,8 @@ func (m *awsRestjson1_deserializeOpListPlugins) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9872,7 +9698,6 @@ func (m *awsRestjson1_deserializeOpListPlugins) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10002,10 +9827,8 @@ func (m *awsRestjson1_deserializeOpListPluginTypeActio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10045,7 +9868,6 @@ func (m *awsRestjson1_deserializeOpListPluginTypeActio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10172,10 +9994,8 @@ func (m *awsRestjson1_deserializeOpListPluginTypeMetad
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10215,7 +10035,6 @@ func (m *awsRestjson1_deserializeOpListPluginTypeMetad
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10342,10 +10161,8 @@ func (m *awsRestjson1_deserializeOpListRetrievers) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10385,7 +10202,6 @@ func (m *awsRestjson1_deserializeOpListRetrievers) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10515,10 +10331,8 @@ func (m *awsRestjson1_deserializeOpListSubscriptions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10558,7 +10372,6 @@ func (m *awsRestjson1_deserializeOpListSubscriptions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10691,10 +10504,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10734,7 +10545,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10855,10 +10665,8 @@ func (m *awsRestjson1_deserializeOpListWebExperiences)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10898,7 +10706,6 @@ func (m *awsRestjson1_deserializeOpListWebExperiences)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11028,10 +10835,8 @@ func (m *awsRestjson1_deserializeOpPutFeedback) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11049,7 +10854,6 @@ func (m *awsRestjson1_deserializeOpPutFeedback) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11134,10 +10938,8 @@ func (m *awsRestjson1_deserializeOpPutGroup) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11149,7 +10951,6 @@ func (m *awsRestjson1_deserializeOpPutGroup) HandleDes
 	output := &PutGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11240,10 +11041,8 @@ func (m *awsRestjson1_deserializeOpSearchRelevantConte
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11283,7 +11082,6 @@ func (m *awsRestjson1_deserializeOpSearchRelevantConte
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11416,10 +11214,8 @@ func (m *awsRestjson1_deserializeOpStartDataSourceSync
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11459,7 +11255,6 @@ func (m *awsRestjson1_deserializeOpStartDataSourceSync
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11590,10 +11385,8 @@ func (m *awsRestjson1_deserializeOpStopDataSourceSyncJ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11605,7 +11398,6 @@ func (m *awsRestjson1_deserializeOpStopDataSourceSyncJ
 	output := &StopDataSourceSyncJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11693,10 +11485,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11708,7 +11498,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11796,10 +11585,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11811,7 +11598,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11896,10 +11682,8 @@ func (m *awsRestjson1_deserializeOpUpdateApplication) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11911,7 +11695,6 @@ func (m *awsRestjson1_deserializeOpUpdateApplication) 
 	output := &UpdateApplicationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11999,10 +11782,8 @@ func (m *awsRestjson1_deserializeOpUpdateChatControlsC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12014,7 +11795,6 @@ func (m *awsRestjson1_deserializeOpUpdateChatControlsC
 	output := &UpdateChatControlsConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12105,10 +11885,8 @@ func (m *awsRestjson1_deserializeOpUpdateChatResponseC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12120,7 +11898,6 @@ func (m *awsRestjson1_deserializeOpUpdateChatResponseC
 	output := &UpdateChatResponseConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12208,10 +11985,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataAccessor)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12223,7 +11998,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataAccessor)
 	output := &UpdateDataAccessorOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12311,10 +12085,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12326,7 +12098,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataSource) H
 	output := &UpdateDataSourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12414,10 +12185,8 @@ func (m *awsRestjson1_deserializeOpUpdateIndex) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12429,7 +12198,6 @@ func (m *awsRestjson1_deserializeOpUpdateIndex) Handle
 	output := &UpdateIndexOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12520,10 +12288,8 @@ func (m *awsRestjson1_deserializeOpUpdatePlugin) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12535,7 +12301,6 @@ func (m *awsRestjson1_deserializeOpUpdatePlugin) Handl
 	output := &UpdatePluginOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12626,10 +12391,8 @@ func (m *awsRestjson1_deserializeOpUpdateRetriever) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12641,7 +12404,6 @@ func (m *awsRestjson1_deserializeOpUpdateRetriever) Ha
 	output := &UpdateRetrieverOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12732,10 +12494,8 @@ func (m *awsRestjson1_deserializeOpUpdateSubscription)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12775,7 +12535,6 @@ func (m *awsRestjson1_deserializeOpUpdateSubscription)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12913,10 +12672,8 @@ func (m *awsRestjson1_deserializeOpUpdateUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12956,7 +12713,6 @@ func (m *awsRestjson1_deserializeOpUpdateUser) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13093,10 +12849,8 @@ func (m *awsRestjson1_deserializeOpUpdateWebExperience
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13108,7 +12862,6 @@ func (m *awsRestjson1_deserializeOpUpdateWebExperience
 	output := &UpdateWebExperienceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
