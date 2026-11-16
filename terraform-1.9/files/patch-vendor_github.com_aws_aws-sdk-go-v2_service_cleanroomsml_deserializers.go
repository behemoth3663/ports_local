--- vendor/github.com/aws/aws-sdk-go-v2/service/cleanroomsml/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/cleanroomsml/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpCancelTrainedModel)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -67,7 +64,6 @@ func (m *awsRestjson1_deserializeOpCancelTrainedModel)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -152,10 +148,8 @@ func (m *awsRestjson1_deserializeOpCancelTrainedModelI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -173,7 +167,6 @@ func (m *awsRestjson1_deserializeOpCancelTrainedModelI
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -258,10 +251,8 @@ func (m *awsRestjson1_deserializeOpCreateAudienceModel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -301,7 +292,6 @@ func (m *awsRestjson1_deserializeOpCreateAudienceModel
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -426,10 +416,8 @@ func (m *awsRestjson1_deserializeOpCreateConfiguredAud
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -469,7 +457,6 @@ func (m *awsRestjson1_deserializeOpCreateConfiguredAud
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -594,10 +581,8 @@ func (m *awsRestjson1_deserializeOpCreateConfiguredMod
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -637,7 +622,6 @@ func (m *awsRestjson1_deserializeOpCreateConfiguredMod
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -759,10 +743,8 @@ func (m *awsRestjson1_deserializeOpCreateConfiguredMod
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -802,7 +784,6 @@ func (m *awsRestjson1_deserializeOpCreateConfiguredMod
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -930,10 +911,8 @@ func (m *awsRestjson1_deserializeOpCreateMLInputChanne
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -973,7 +952,6 @@ func (m *awsRestjson1_deserializeOpCreateMLInputChanne
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1101,10 +1079,8 @@ func (m *awsRestjson1_deserializeOpCreateTrainedModel)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1144,7 +1120,6 @@ func (m *awsRestjson1_deserializeOpCreateTrainedModel)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1284,10 +1259,8 @@ func (m *awsRestjson1_deserializeOpCreateTrainingDatas
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1327,7 +1300,6 @@ func (m *awsRestjson1_deserializeOpCreateTrainingDatas
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1446,10 +1418,8 @@ func (m *awsRestjson1_deserializeOpDeleteAudienceGener
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1467,7 +1437,6 @@ func (m *awsRestjson1_deserializeOpDeleteAudienceGener
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1549,10 +1518,8 @@ func (m *awsRestjson1_deserializeOpDeleteAudienceModel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1570,7 +1537,6 @@ func (m *awsRestjson1_deserializeOpDeleteAudienceModel
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1652,10 +1618,8 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguredAud
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1673,7 +1637,6 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguredAud
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1755,10 +1718,8 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguredAud
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1776,7 +1737,6 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguredAud
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1855,10 +1815,8 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguredMod
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1876,7 +1834,6 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguredMod
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1958,10 +1915,8 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguredMod
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1979,7 +1934,6 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguredMod
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2064,10 +2018,8 @@ func (m *awsRestjson1_deserializeOpDeleteMLConfigurati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2085,7 +2037,6 @@ func (m *awsRestjson1_deserializeOpDeleteMLConfigurati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2167,10 +2118,8 @@ func (m *awsRestjson1_deserializeOpDeleteMLInputChanne
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2188,7 +2137,6 @@ func (m *awsRestjson1_deserializeOpDeleteMLInputChanne
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2273,10 +2221,8 @@ func (m *awsRestjson1_deserializeOpDeleteTrainedModelO
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2294,7 +2240,6 @@ func (m *awsRestjson1_deserializeOpDeleteTrainedModelO
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2379,10 +2324,8 @@ func (m *awsRestjson1_deserializeOpDeleteTrainingDatas
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2400,7 +2343,6 @@ func (m *awsRestjson1_deserializeOpDeleteTrainingDatas
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2482,10 +2424,8 @@ func (m *awsRestjson1_deserializeOpGetAudienceGenerati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2525,7 +2465,6 @@ func (m *awsRestjson1_deserializeOpGetAudienceGenerati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2762,10 +2701,8 @@ func (m *awsRestjson1_deserializeOpGetAudienceModel) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2805,7 +2742,6 @@ func (m *awsRestjson1_deserializeOpGetAudienceModel) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3031,10 +2967,8 @@ func (m *awsRestjson1_deserializeOpGetCollaborationCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3074,7 +3008,6 @@ func (m *awsRestjson1_deserializeOpGetCollaborationCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3281,10 +3214,8 @@ func (m *awsRestjson1_deserializeOpGetCollaborationMLI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3324,7 +3255,6 @@ func (m *awsRestjson1_deserializeOpGetCollaborationMLI
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3562,10 +3492,8 @@ func (m *awsRestjson1_deserializeOpGetCollaborationTra
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3605,7 +3533,6 @@ func (m *awsRestjson1_deserializeOpGetCollaborationTra
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3899,10 +3826,8 @@ func (m *awsRestjson1_deserializeOpGetConfiguredAudien
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3942,7 +3867,6 @@ func (m *awsRestjson1_deserializeOpGetConfiguredAudien
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4165,10 +4089,8 @@ func (m *awsRestjson1_deserializeOpGetConfiguredAudien
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4208,7 +4130,6 @@ func (m *awsRestjson1_deserializeOpGetConfiguredAudien
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4345,10 +4266,8 @@ func (m *awsRestjson1_deserializeOpGetConfiguredModelA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4388,7 +4307,6 @@ func (m *awsRestjson1_deserializeOpGetConfiguredModelA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4584,10 +4502,8 @@ func (m *awsRestjson1_deserializeOpGetConfiguredModelA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4627,7 +4543,6 @@ func (m *awsRestjson1_deserializeOpGetConfiguredModelA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4830,10 +4745,8 @@ func (m *awsRestjson1_deserializeOpGetMLConfiguration)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4873,7 +4786,6 @@ func (m *awsRestjson1_deserializeOpGetMLConfiguration)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5026,10 +4938,8 @@ func (m *awsRestjson1_deserializeOpGetMLInputChannel) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5069,7 +4979,6 @@ func (m *awsRestjson1_deserializeOpGetMLInputChannel) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5394,10 +5303,8 @@ func (m *awsRestjson1_deserializeOpGetTrainedModel) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5437,7 +5344,6 @@ func (m *awsRestjson1_deserializeOpGetTrainedModel) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5751,10 +5657,8 @@ func (m *awsRestjson1_deserializeOpGetTrainedModelInfe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5794,7 +5698,6 @@ func (m *awsRestjson1_deserializeOpGetTrainedModelInfe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6094,10 +5997,8 @@ func (m *awsRestjson1_deserializeOpGetTrainingDataset)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6137,7 +6038,6 @@ func (m *awsRestjson1_deserializeOpGetTrainingDataset)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6328,10 +6228,8 @@ func (m *awsRestjson1_deserializeOpListAudienceExportJ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6371,7 +6269,6 @@ func (m *awsRestjson1_deserializeOpListAudienceExportJ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6492,10 +6389,8 @@ func (m *awsRestjson1_deserializeOpListAudienceGenerat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6535,7 +6430,6 @@ func (m *awsRestjson1_deserializeOpListAudienceGenerat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6656,10 +6550,8 @@ func (m *awsRestjson1_deserializeOpListAudienceModels)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6699,7 +6591,6 @@ func (m *awsRestjson1_deserializeOpListAudienceModels)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6820,10 +6711,8 @@ func (m *awsRestjson1_deserializeOpListCollaborationCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6863,7 +6752,6 @@ func (m *awsRestjson1_deserializeOpListCollaborationCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6987,10 +6875,8 @@ func (m *awsRestjson1_deserializeOpListCollaborationML
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7030,7 +6916,6 @@ func (m *awsRestjson1_deserializeOpListCollaborationML
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7154,10 +7039,8 @@ func (m *awsRestjson1_deserializeOpListCollaborationTr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7197,7 +7080,6 @@ func (m *awsRestjson1_deserializeOpListCollaborationTr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7321,10 +7203,8 @@ func (m *awsRestjson1_deserializeOpListCollaborationTr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7364,7 +7244,6 @@ func (m *awsRestjson1_deserializeOpListCollaborationTr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7488,10 +7367,8 @@ func (m *awsRestjson1_deserializeOpListCollaborationTr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7531,7 +7408,6 @@ func (m *awsRestjson1_deserializeOpListCollaborationTr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7655,10 +7531,8 @@ func (m *awsRestjson1_deserializeOpListConfiguredAudie
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7698,7 +7572,6 @@ func (m *awsRestjson1_deserializeOpListConfiguredAudie
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7819,10 +7692,8 @@ func (m *awsRestjson1_deserializeOpListConfiguredModel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7862,7 +7733,6 @@ func (m *awsRestjson1_deserializeOpListConfiguredModel
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7986,10 +7856,8 @@ func (m *awsRestjson1_deserializeOpListConfiguredModel
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8029,7 +7897,6 @@ func (m *awsRestjson1_deserializeOpListConfiguredModel
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8150,10 +8017,8 @@ func (m *awsRestjson1_deserializeOpListMLInputChannels
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8193,7 +8058,6 @@ func (m *awsRestjson1_deserializeOpListMLInputChannels
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8317,10 +8181,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8360,7 +8222,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8475,10 +8336,8 @@ func (m *awsRestjson1_deserializeOpListTrainedModelInf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8518,7 +8377,6 @@ func (m *awsRestjson1_deserializeOpListTrainedModelInf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8642,10 +8500,8 @@ func (m *awsRestjson1_deserializeOpListTrainedModels) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8685,7 +8541,6 @@ func (m *awsRestjson1_deserializeOpListTrainedModels) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8809,10 +8664,8 @@ func (m *awsRestjson1_deserializeOpListTrainedModelVer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8852,7 +8705,6 @@ func (m *awsRestjson1_deserializeOpListTrainedModelVer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8979,10 +8831,8 @@ func (m *awsRestjson1_deserializeOpListTrainingDataset
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9022,7 +8872,6 @@ func (m *awsRestjson1_deserializeOpListTrainingDataset
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9143,10 +8992,8 @@ func (m *awsRestjson1_deserializeOpPutConfiguredAudien
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9186,7 +9033,6 @@ func (m *awsRestjson1_deserializeOpPutConfiguredAudien
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9314,10 +9160,8 @@ func (m *awsRestjson1_deserializeOpPutMLConfiguration)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9335,7 +9179,6 @@ func (m *awsRestjson1_deserializeOpPutMLConfiguration)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9414,10 +9257,8 @@ func (m *awsRestjson1_deserializeOpStartAudienceExport
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9435,7 +9276,6 @@ func (m *awsRestjson1_deserializeOpStartAudienceExport
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9520,10 +9360,8 @@ func (m *awsRestjson1_deserializeOpStartAudienceGenera
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9563,7 +9401,6 @@ func (m *awsRestjson1_deserializeOpStartAudienceGenera
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9691,10 +9528,8 @@ func (m *awsRestjson1_deserializeOpStartTrainedModelEx
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9712,7 +9547,6 @@ func (m *awsRestjson1_deserializeOpStartTrainedModelEx
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9797,10 +9631,8 @@ func (m *awsRestjson1_deserializeOpStartTrainedModelIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9840,7 +9672,6 @@ func (m *awsRestjson1_deserializeOpStartTrainedModelIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9968,10 +9799,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9983,7 +9812,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10062,10 +9890,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10077,7 +9903,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10156,10 +9981,8 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguredAud
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10199,7 +10022,6 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguredAud
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
