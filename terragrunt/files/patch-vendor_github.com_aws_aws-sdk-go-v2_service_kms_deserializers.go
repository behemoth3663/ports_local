--- vendor/github.com/aws/aws-sdk-go-v2/service/kms/deserializers.go.orig	2026-11-15 15:04:53 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/kms/deserializers.go
@@ -15,7 +15,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -36,10 +35,9 @@ func (m *awsAwsjson11_deserializeOpCancelKeyDeletion) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -159,10 +157,9 @@ func (m *awsAwsjson11_deserializeOpConnectCustomKeySto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -282,10 +279,9 @@ func (m *awsAwsjson11_deserializeOpCreateAlias) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -389,10 +385,9 @@ func (m *awsAwsjson11_deserializeOpCreateCustomKeyStor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -548,10 +543,9 @@ func (m *awsAwsjson11_deserializeOpCreateGrant) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -683,10 +677,9 @@ func (m *awsAwsjson11_deserializeOpCreateKey) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -830,10 +823,9 @@ func (m *awsAwsjson11_deserializeOpDecrypt) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -971,10 +963,9 @@ func (m *awsAwsjson11_deserializeOpDeleteAlias) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1069,10 +1060,9 @@ func (m *awsAwsjson11_deserializeOpDeleteCustomKeyStor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1189,10 +1179,9 @@ func (m *awsAwsjson11_deserializeOpDeleteImportedKeyMa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1315,10 +1304,9 @@ func (m *awsAwsjson11_deserializeOpDeriveSharedSecret)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1450,10 +1438,9 @@ func (m *awsAwsjson11_deserializeOpDescribeCustomKeySt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1567,10 +1554,9 @@ func (m *awsAwsjson11_deserializeOpDescribeKey) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1687,10 +1673,9 @@ func (m *awsAwsjson11_deserializeOpDisableKey) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1788,10 +1773,9 @@ func (m *awsAwsjson11_deserializeOpDisableKeyRotation)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1895,10 +1879,9 @@ func (m *awsAwsjson11_deserializeOpDisconnectCustomKey
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2012,10 +1995,9 @@ func (m *awsAwsjson11_deserializeOpEnableKey) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2116,10 +2098,9 @@ func (m *awsAwsjson11_deserializeOpEnableKeyRotation) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2223,10 +2204,9 @@ func (m *awsAwsjson11_deserializeOpEncrypt) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2358,10 +2338,9 @@ func (m *awsAwsjson11_deserializeOpGenerateDataKey) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2493,10 +2472,9 @@ func (m *awsAwsjson11_deserializeOpGenerateDataKeyPair
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2631,10 +2609,9 @@ func (m *awsAwsjson11_deserializeOpGenerateDataKeyPair
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2769,10 +2746,9 @@ func (m *awsAwsjson11_deserializeOpGenerateDataKeyWith
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2904,10 +2880,9 @@ func (m *awsAwsjson11_deserializeOpGenerateMac) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3036,10 +3011,9 @@ func (m *awsAwsjson11_deserializeOpGenerateRandom) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3159,10 +3133,9 @@ func (m *awsAwsjson11_deserializeOpGetKeyLastUsage) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3279,10 +3252,9 @@ func (m *awsAwsjson11_deserializeOpGetKeyPolicy) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3402,10 +3374,9 @@ func (m *awsAwsjson11_deserializeOpGetKeyRotationStatu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3528,10 +3499,9 @@ func (m *awsAwsjson11_deserializeOpGetParametersForImp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3654,10 +3624,9 @@ func (m *awsAwsjson11_deserializeOpGetPublicKey) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3792,10 +3761,9 @@ func (m *awsAwsjson11_deserializeOpImportKeyMaterial) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3930,10 +3898,9 @@ func (m *awsAwsjson11_deserializeOpListAliases) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4053,10 +4020,9 @@ func (m *awsAwsjson11_deserializeOpListGrants) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4182,10 +4148,9 @@ func (m *awsAwsjson11_deserializeOpListKeyPolicies) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4305,10 +4270,9 @@ func (m *awsAwsjson11_deserializeOpListKeyRotations) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4431,10 +4395,9 @@ func (m *awsAwsjson11_deserializeOpListKeys) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4548,10 +4511,9 @@ func (m *awsAwsjson11_deserializeOpListResourceTags) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4668,10 +4630,9 @@ func (m *awsAwsjson11_deserializeOpListRetirableGrants
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4791,10 +4752,9 @@ func (m *awsAwsjson11_deserializeOpPutKeyPolicy) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4901,10 +4861,9 @@ func (m *awsAwsjson11_deserializeOpReEncrypt) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5042,10 +5001,9 @@ func (m *awsAwsjson11_deserializeOpReplicateKey) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5180,10 +5138,9 @@ func (m *awsAwsjson11_deserializeOpRetireGrant) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5290,10 +5247,9 @@ func (m *awsAwsjson11_deserializeOpRevokeGrant) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5397,10 +5353,9 @@ func (m *awsAwsjson11_deserializeOpRotateKeyOnDemand) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5532,10 +5487,9 @@ func (m *awsAwsjson11_deserializeOpScheduleKeyDeletion
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5655,10 +5609,9 @@ func (m *awsAwsjson11_deserializeOpSign) HandleDeseria
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5790,10 +5743,9 @@ func (m *awsAwsjson11_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5894,10 +5846,9 @@ func (m *awsAwsjson11_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5995,10 +5946,9 @@ func (m *awsAwsjson11_deserializeOpUpdateAlias) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6096,10 +6046,9 @@ func (m *awsAwsjson11_deserializeOpUpdateCustomKeyStor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6255,10 +6204,9 @@ func (m *awsAwsjson11_deserializeOpUpdateKeyDescriptio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6356,10 +6304,9 @@ func (m *awsAwsjson11_deserializeOpUpdatePrimaryRegion
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6460,10 +6407,9 @@ func (m *awsAwsjson11_deserializeOpVerify) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6598,10 +6544,9 @@ func (m *awsAwsjson11_deserializeOpVerifyMac) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
+
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
