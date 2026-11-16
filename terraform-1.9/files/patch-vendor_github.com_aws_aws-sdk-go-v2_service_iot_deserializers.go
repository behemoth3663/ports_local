--- vendor/github.com/aws/aws-sdk-go-v2/service/iot/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/iot/deserializers.go
@@ -15,7 +15,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -47,10 +46,8 @@ func (m *awsRestjson1_deserializeOpAcceptCertificateTr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -68,7 +65,6 @@ func (m *awsRestjson1_deserializeOpAcceptCertificateTr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -159,10 +155,8 @@ func (m *awsRestjson1_deserializeOpAddThingToBillingGr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -174,7 +168,6 @@ func (m *awsRestjson1_deserializeOpAddThingToBillingGr
 	output := &AddThingToBillingGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -256,10 +249,8 @@ func (m *awsRestjson1_deserializeOpAddThingToThingGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -271,7 +262,6 @@ func (m *awsRestjson1_deserializeOpAddThingToThingGrou
 	output := &AddThingToThingGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -353,10 +343,8 @@ func (m *awsRestjson1_deserializeOpAssociateSbomWithPa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -396,7 +384,6 @@ func (m *awsRestjson1_deserializeOpAssociateSbomWithPa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -547,10 +534,8 @@ func (m *awsRestjson1_deserializeOpAssociateTargetsWit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -590,7 +575,6 @@ func (m *awsRestjson1_deserializeOpAssociateTargetsWit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -733,10 +717,8 @@ func (m *awsRestjson1_deserializeOpAttachPolicy) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -754,7 +736,6 @@ func (m *awsRestjson1_deserializeOpAttachPolicy) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -845,10 +826,8 @@ func (m *awsRestjson1_deserializeOpAttachPrincipalPoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -866,7 +845,6 @@ func (m *awsRestjson1_deserializeOpAttachPrincipalPoli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -957,10 +935,8 @@ func (m *awsRestjson1_deserializeOpAttachSecurityProfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -972,7 +948,6 @@ func (m *awsRestjson1_deserializeOpAttachSecurityProfi
 	output := &AttachSecurityProfileOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1060,10 +1035,8 @@ func (m *awsRestjson1_deserializeOpAttachThingPrincipa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1075,7 +1048,6 @@ func (m *awsRestjson1_deserializeOpAttachThingPrincipa
 	output := &AttachThingPrincipalOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1163,10 +1135,8 @@ func (m *awsRestjson1_deserializeOpCancelAuditMitigati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1178,7 +1148,6 @@ func (m *awsRestjson1_deserializeOpCancelAuditMitigati
 	output := &CancelAuditMitigationActionsTaskOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1260,10 +1229,8 @@ func (m *awsRestjson1_deserializeOpCancelAuditTask) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1275,7 +1242,6 @@ func (m *awsRestjson1_deserializeOpCancelAuditTask) Ha
 	output := &CancelAuditTaskOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1357,10 +1323,8 @@ func (m *awsRestjson1_deserializeOpCancelCertificateTr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1378,7 +1342,6 @@ func (m *awsRestjson1_deserializeOpCancelCertificateTr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1469,10 +1432,8 @@ func (m *awsRestjson1_deserializeOpCancelDetectMitigat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1484,7 +1445,6 @@ func (m *awsRestjson1_deserializeOpCancelDetectMitigat
 	output := &CancelDetectMitigationActionsTaskOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1566,10 +1526,8 @@ func (m *awsRestjson1_deserializeOpCancelJob) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1609,7 +1567,6 @@ func (m *awsRestjson1_deserializeOpCancelJob) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1752,10 +1709,8 @@ func (m *awsRestjson1_deserializeOpCancelJobExecution)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1773,7 +1728,6 @@ func (m *awsRestjson1_deserializeOpCancelJobExecution)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1861,10 +1815,8 @@ func (m *awsRestjson1_deserializeOpClearDefaultAuthori
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1876,7 +1828,6 @@ func (m *awsRestjson1_deserializeOpClearDefaultAuthori
 	output := &ClearDefaultAuthorizerOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1964,10 +1915,8 @@ func (m *awsRestjson1_deserializeOpConfirmTopicRuleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1979,7 +1928,6 @@ func (m *awsRestjson1_deserializeOpConfirmTopicRuleDes
 	output := &ConfirmTopicRuleDestinationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2064,10 +2012,8 @@ func (m *awsRestjson1_deserializeOpCreateAuditSuppress
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2079,7 +2025,6 @@ func (m *awsRestjson1_deserializeOpCreateAuditSuppress
 	output := &CreateAuditSuppressionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2164,10 +2109,8 @@ func (m *awsRestjson1_deserializeOpCreateAuthorizer) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2207,7 +2150,6 @@ func (m *awsRestjson1_deserializeOpCreateAuthorizer) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2347,10 +2289,8 @@ func (m *awsRestjson1_deserializeOpCreateBillingGroup)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2390,7 +2330,6 @@ func (m *awsRestjson1_deserializeOpCreateBillingGroup)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2530,10 +2469,8 @@ func (m *awsRestjson1_deserializeOpCreateCertificateFr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2573,7 +2510,6 @@ func (m *awsRestjson1_deserializeOpCreateCertificateFr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2716,10 +2652,8 @@ func (m *awsRestjson1_deserializeOpCreateCertificatePr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2759,7 +2693,6 @@ func (m *awsRestjson1_deserializeOpCreateCertificatePr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2899,10 +2832,8 @@ func (m *awsRestjson1_deserializeOpCreateCommand) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2942,7 +2873,6 @@ func (m *awsRestjson1_deserializeOpCreateCommand) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3076,10 +3006,8 @@ func (m *awsRestjson1_deserializeOpCreateCustomMetric)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3119,7 +3047,6 @@ func (m *awsRestjson1_deserializeOpCreateCustomMetric)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3253,10 +3180,8 @@ func (m *awsRestjson1_deserializeOpCreateDimension) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3296,7 +3221,6 @@ func (m *awsRestjson1_deserializeOpCreateDimension) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3430,10 +3354,8 @@ func (m *awsRestjson1_deserializeOpCreateDomainConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3473,7 +3395,6 @@ func (m *awsRestjson1_deserializeOpCreateDomainConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3616,10 +3537,8 @@ func (m *awsRestjson1_deserializeOpCreateDynamicThingG
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3659,7 +3578,6 @@ func (m *awsRestjson1_deserializeOpCreateDynamicThingG
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3835,10 +3753,8 @@ func (m *awsRestjson1_deserializeOpCreateFleetMetric) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3878,7 +3794,6 @@ func (m *awsRestjson1_deserializeOpCreateFleetMetric) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4030,10 +3945,8 @@ func (m *awsRestjson1_deserializeOpCreateJob) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4073,7 +3986,6 @@ func (m *awsRestjson1_deserializeOpCreateJob) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4219,10 +4131,8 @@ func (m *awsRestjson1_deserializeOpCreateJobTemplate) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4262,7 +4172,6 @@ func (m *awsRestjson1_deserializeOpCreateJobTemplate) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4399,10 +4308,8 @@ func (m *awsRestjson1_deserializeOpCreateKeysAndCertif
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4442,7 +4349,6 @@ func (m *awsRestjson1_deserializeOpCreateKeysAndCertif
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4590,10 +4496,8 @@ func (m *awsRestjson1_deserializeOpCreateMitigationAct
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4633,7 +4537,6 @@ func (m *awsRestjson1_deserializeOpCreateMitigationAct
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4767,10 +4670,8 @@ func (m *awsRestjson1_deserializeOpCreateOTAUpdate) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4810,7 +4711,6 @@ func (m *awsRestjson1_deserializeOpCreateOTAUpdate) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4980,10 +4880,8 @@ func (m *awsRestjson1_deserializeOpCreatePackage) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5023,7 +4921,6 @@ func (m *awsRestjson1_deserializeOpCreatePackage) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5166,10 +5063,8 @@ func (m *awsRestjson1_deserializeOpCreatePackageVersio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5209,7 +5104,6 @@ func (m *awsRestjson1_deserializeOpCreatePackageVersio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5384,10 +5278,8 @@ func (m *awsRestjson1_deserializeOpCreatePolicy) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5427,7 +5319,6 @@ func (m *awsRestjson1_deserializeOpCreatePolicy) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5585,10 +5476,8 @@ func (m *awsRestjson1_deserializeOpCreatePolicyVersion
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5628,7 +5517,6 @@ func (m *awsRestjson1_deserializeOpCreatePolicyVersion
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5789,10 +5677,8 @@ func (m *awsRestjson1_deserializeOpCreateProvisioningC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5832,7 +5718,6 @@ func (m *awsRestjson1_deserializeOpCreateProvisioningC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5990,10 +5875,8 @@ func (m *awsRestjson1_deserializeOpCreateProvisioningT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6033,7 +5916,6 @@ func (m *awsRestjson1_deserializeOpCreateProvisioningT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6183,10 +6065,8 @@ func (m *awsRestjson1_deserializeOpCreateProvisioningT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6226,7 +6106,6 @@ func (m *awsRestjson1_deserializeOpCreateProvisioningT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6388,10 +6267,8 @@ func (m *awsRestjson1_deserializeOpCreateRoleAlias) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6431,7 +6308,6 @@ func (m *awsRestjson1_deserializeOpCreateRoleAlias) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6571,10 +6447,8 @@ func (m *awsRestjson1_deserializeOpCreateScheduledAudi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6614,7 +6488,6 @@ func (m *awsRestjson1_deserializeOpCreateScheduledAudi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6739,10 +6612,8 @@ func (m *awsRestjson1_deserializeOpCreateSecurityProfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6782,7 +6653,6 @@ func (m *awsRestjson1_deserializeOpCreateSecurityProfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6913,10 +6783,8 @@ func (m *awsRestjson1_deserializeOpCreateStream) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6956,7 +6824,6 @@ func (m *awsRestjson1_deserializeOpCreateStream) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7121,10 +6988,8 @@ func (m *awsRestjson1_deserializeOpCreateThing) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7164,7 +7029,6 @@ func (m *awsRestjson1_deserializeOpCreateThing) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7313,10 +7177,8 @@ func (m *awsRestjson1_deserializeOpCreateThingGroup) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7356,7 +7218,6 @@ func (m *awsRestjson1_deserializeOpCreateThingGroup) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7496,10 +7357,8 @@ func (m *awsRestjson1_deserializeOpCreateThingType) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7539,7 +7398,6 @@ func (m *awsRestjson1_deserializeOpCreateThingType) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7685,10 +7543,8 @@ func (m *awsRestjson1_deserializeOpCreateTopicRule) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7706,7 +7562,6 @@ func (m *awsRestjson1_deserializeOpCreateTopicRule) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7794,10 +7649,8 @@ func (m *awsRestjson1_deserializeOpCreateTopicRuleDest
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7837,7 +7690,6 @@ func (m *awsRestjson1_deserializeOpCreateTopicRuleDest
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7958,10 +7810,8 @@ func (m *awsRestjson1_deserializeOpDeleteAccountAuditC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7973,7 +7823,6 @@ func (m *awsRestjson1_deserializeOpDeleteAccountAuditC
 	output := &DeleteAccountAuditConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8055,10 +7904,8 @@ func (m *awsRestjson1_deserializeOpDeleteAuditSuppress
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8070,7 +7917,6 @@ func (m *awsRestjson1_deserializeOpDeleteAuditSuppress
 	output := &DeleteAuditSuppressionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8149,10 +7995,8 @@ func (m *awsRestjson1_deserializeOpDeleteAuthorizer) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8164,7 +8008,6 @@ func (m *awsRestjson1_deserializeOpDeleteAuthorizer) H
 	output := &DeleteAuthorizerOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8255,10 +8098,8 @@ func (m *awsRestjson1_deserializeOpDeleteBillingGroup)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8270,7 +8111,6 @@ func (m *awsRestjson1_deserializeOpDeleteBillingGroup)
 	output := &DeleteBillingGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8352,10 +8192,8 @@ func (m *awsRestjson1_deserializeOpDeleteCACertificate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8367,7 +8205,6 @@ func (m *awsRestjson1_deserializeOpDeleteCACertificate
 	output := &DeleteCACertificateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8458,10 +8295,8 @@ func (m *awsRestjson1_deserializeOpDeleteCertificate) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8479,7 +8314,6 @@ func (m *awsRestjson1_deserializeOpDeleteCertificate) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8573,10 +8407,8 @@ func (m *awsRestjson1_deserializeOpDeleteCertificatePr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8588,7 +8420,6 @@ func (m *awsRestjson1_deserializeOpDeleteCertificatePr
 	output := &DeleteCertificateProviderOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8679,10 +8510,8 @@ func (m *awsRestjson1_deserializeOpDeleteCommand) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8699,7 +8528,6 @@ func (m *awsRestjson1_deserializeOpDeleteCommand) Hand
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to decode response with invalid Http bindings, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8791,10 +8619,8 @@ func (m *awsRestjson1_deserializeOpDeleteCommandExecut
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8806,7 +8632,6 @@ func (m *awsRestjson1_deserializeOpDeleteCommandExecut
 	output := &DeleteCommandExecutionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8888,10 +8713,8 @@ func (m *awsRestjson1_deserializeOpDeleteCustomMetric)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8903,7 +8726,6 @@ func (m *awsRestjson1_deserializeOpDeleteCustomMetric)
 	output := &DeleteCustomMetricOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8982,10 +8804,8 @@ func (m *awsRestjson1_deserializeOpDeleteDimension) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8997,7 +8817,6 @@ func (m *awsRestjson1_deserializeOpDeleteDimension) Ha
 	output := &DeleteDimensionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9076,10 +8895,8 @@ func (m *awsRestjson1_deserializeOpDeleteDomainConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9091,7 +8908,6 @@ func (m *awsRestjson1_deserializeOpDeleteDomainConfigu
 	output := &DeleteDomainConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9179,10 +8995,8 @@ func (m *awsRestjson1_deserializeOpDeleteDynamicThingG
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9194,7 +9008,6 @@ func (m *awsRestjson1_deserializeOpDeleteDynamicThingG
 	output := &DeleteDynamicThingGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9276,10 +9089,8 @@ func (m *awsRestjson1_deserializeOpDeleteFleetMetric) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9297,7 +9108,6 @@ func (m *awsRestjson1_deserializeOpDeleteFleetMetric) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9385,10 +9195,8 @@ func (m *awsRestjson1_deserializeOpDeleteJob) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9406,7 +9214,6 @@ func (m *awsRestjson1_deserializeOpDeleteJob) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9494,10 +9301,8 @@ func (m *awsRestjson1_deserializeOpDeleteJobExecution)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9515,7 +9320,6 @@ func (m *awsRestjson1_deserializeOpDeleteJobExecution)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9600,10 +9404,8 @@ func (m *awsRestjson1_deserializeOpDeleteJobTemplate) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9621,7 +9423,6 @@ func (m *awsRestjson1_deserializeOpDeleteJobTemplate) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9703,10 +9504,8 @@ func (m *awsRestjson1_deserializeOpDeleteMitigationAct
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9718,7 +9517,6 @@ func (m *awsRestjson1_deserializeOpDeleteMitigationAct
 	output := &DeleteMitigationActionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9797,10 +9595,8 @@ func (m *awsRestjson1_deserializeOpDeleteOTAUpdate) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9812,7 +9608,6 @@ func (m *awsRestjson1_deserializeOpDeleteOTAUpdate) Ha
 	output := &DeleteOTAUpdateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9903,10 +9698,8 @@ func (m *awsRestjson1_deserializeOpDeletePackage) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9918,7 +9711,6 @@ func (m *awsRestjson1_deserializeOpDeletePackage) Hand
 	output := &DeletePackageOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9997,10 +9789,8 @@ func (m *awsRestjson1_deserializeOpDeletePackageVersio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10012,7 +9802,6 @@ func (m *awsRestjson1_deserializeOpDeletePackageVersio
 	output := &DeletePackageVersionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10091,10 +9880,8 @@ func (m *awsRestjson1_deserializeOpDeletePolicy) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10112,7 +9899,6 @@ func (m *awsRestjson1_deserializeOpDeletePolicy) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10203,10 +9989,8 @@ func (m *awsRestjson1_deserializeOpDeletePolicyVersion
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10224,7 +10008,6 @@ func (m *awsRestjson1_deserializeOpDeletePolicyVersion
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10315,10 +10098,8 @@ func (m *awsRestjson1_deserializeOpDeleteProvisioningT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10330,7 +10111,6 @@ func (m *awsRestjson1_deserializeOpDeleteProvisioningT
 	output := &DeleteProvisioningTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10421,10 +10201,8 @@ func (m *awsRestjson1_deserializeOpDeleteProvisioningT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10436,7 +10214,6 @@ func (m *awsRestjson1_deserializeOpDeleteProvisioningT
 	output := &DeleteProvisioningTemplateVersionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10527,10 +10304,8 @@ func (m *awsRestjson1_deserializeOpDeleteRegistrationC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10542,7 +10317,6 @@ func (m *awsRestjson1_deserializeOpDeleteRegistrationC
 	output := &DeleteRegistrationCodeOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10627,10 +10401,8 @@ func (m *awsRestjson1_deserializeOpDeleteRoleAlias) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10642,7 +10414,6 @@ func (m *awsRestjson1_deserializeOpDeleteRoleAlias) Ha
 	output := &DeleteRoleAliasOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10733,10 +10504,8 @@ func (m *awsRestjson1_deserializeOpDeleteScheduledAudi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10748,7 +10517,6 @@ func (m *awsRestjson1_deserializeOpDeleteScheduledAudi
 	output := &DeleteScheduledAuditOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10830,10 +10598,8 @@ func (m *awsRestjson1_deserializeOpDeleteSecurityProfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10845,7 +10611,6 @@ func (m *awsRestjson1_deserializeOpDeleteSecurityProfi
 	output := &DeleteSecurityProfileOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10927,10 +10692,8 @@ func (m *awsRestjson1_deserializeOpDeleteStream) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10942,7 +10705,6 @@ func (m *awsRestjson1_deserializeOpDeleteStream) Handl
 	output := &DeleteStreamOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11033,10 +10795,8 @@ func (m *awsRestjson1_deserializeOpDeleteThing) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11048,7 +10808,6 @@ func (m *awsRestjson1_deserializeOpDeleteThing) Handle
 	output := &DeleteThingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11139,10 +10898,8 @@ func (m *awsRestjson1_deserializeOpDeleteThingGroup) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11154,7 +10911,6 @@ func (m *awsRestjson1_deserializeOpDeleteThingGroup) H
 	output := &DeleteThingGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11236,10 +10992,8 @@ func (m *awsRestjson1_deserializeOpDeleteThingType) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11251,7 +11005,6 @@ func (m *awsRestjson1_deserializeOpDeleteThingType) Ha
 	output := &DeleteThingTypeOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11339,10 +11092,8 @@ func (m *awsRestjson1_deserializeOpDeleteTopicRule) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11360,7 +11111,6 @@ func (m *awsRestjson1_deserializeOpDeleteTopicRule) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11445,10 +11195,8 @@ func (m *awsRestjson1_deserializeOpDeleteTopicRuleDest
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11460,7 +11208,6 @@ func (m *awsRestjson1_deserializeOpDeleteTopicRuleDest
 	output := &DeleteTopicRuleDestinationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11545,10 +11292,8 @@ func (m *awsRestjson1_deserializeOpDeleteV2LoggingLeve
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11566,7 +11311,6 @@ func (m *awsRestjson1_deserializeOpDeleteV2LoggingLeve
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11645,10 +11389,8 @@ func (m *awsRestjson1_deserializeOpDeprecateThingType)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11660,7 +11402,6 @@ func (m *awsRestjson1_deserializeOpDeprecateThingType)
 	output := &DeprecateThingTypeOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11748,10 +11489,8 @@ func (m *awsRestjson1_deserializeOpDescribeAccountAudi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11791,7 +11530,6 @@ func (m *awsRestjson1_deserializeOpDescribeAccountAudi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11917,10 +11655,8 @@ func (m *awsRestjson1_deserializeOpDescribeAuditFindin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11960,7 +11696,6 @@ func (m *awsRestjson1_deserializeOpDescribeAuditFindin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12078,10 +11813,8 @@ func (m *awsRestjson1_deserializeOpDescribeAuditMitiga
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12121,7 +11854,6 @@ func (m *awsRestjson1_deserializeOpDescribeAuditMitiga
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12295,10 +12027,8 @@ func (m *awsRestjson1_deserializeOpDescribeAuditSuppre
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12338,7 +12068,6 @@ func (m *awsRestjson1_deserializeOpDescribeAuditSuppre
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12499,10 +12228,8 @@ func (m *awsRestjson1_deserializeOpDescribeAuditTask) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12542,7 +12269,6 @@ func (m *awsRestjson1_deserializeOpDescribeAuditTask) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12708,10 +12434,8 @@ func (m *awsRestjson1_deserializeOpDescribeAuthorizer)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12751,7 +12475,6 @@ func (m *awsRestjson1_deserializeOpDescribeAuthorizer)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12875,10 +12598,8 @@ func (m *awsRestjson1_deserializeOpDescribeBillingGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12918,7 +12639,6 @@ func (m *awsRestjson1_deserializeOpDescribeBillingGrou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13081,10 +12801,8 @@ func (m *awsRestjson1_deserializeOpDescribeCACertifica
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13124,7 +12842,6 @@ func (m *awsRestjson1_deserializeOpDescribeCACertifica
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13253,10 +12970,8 @@ func (m *awsRestjson1_deserializeOpDescribeCertificate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13296,7 +13011,6 @@ func (m *awsRestjson1_deserializeOpDescribeCertificate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13420,10 +13134,8 @@ func (m *awsRestjson1_deserializeOpDescribeCertificate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13463,7 +13175,6 @@ func (m *awsRestjson1_deserializeOpDescribeCertificate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13646,10 +13357,8 @@ func (m *awsRestjson1_deserializeOpDescribeCustomMetri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13689,7 +13398,6 @@ func (m *awsRestjson1_deserializeOpDescribeCustomMetri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13870,10 +13578,8 @@ func (m *awsRestjson1_deserializeOpDescribeDefaultAuth
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13913,7 +13619,6 @@ func (m *awsRestjson1_deserializeOpDescribeDefaultAuth
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14037,10 +13742,8 @@ func (m *awsRestjson1_deserializeOpDescribeDetectMitig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14080,7 +13783,6 @@ func (m *awsRestjson1_deserializeOpDescribeDetectMitig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14198,10 +13900,8 @@ func (m *awsRestjson1_deserializeOpDescribeDimension) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14241,7 +13941,6 @@ func (m *awsRestjson1_deserializeOpDescribeDimension) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14418,10 +14117,8 @@ func (m *awsRestjson1_deserializeOpDescribeDomainConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14461,7 +14158,6 @@ func (m *awsRestjson1_deserializeOpDescribeDomainConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14693,10 +14389,8 @@ func (m *awsRestjson1_deserializeOpDescribeEndpoint) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14736,7 +14430,6 @@ func (m *awsRestjson1_deserializeOpDescribeEndpoint) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14858,10 +14551,8 @@ func (m *awsRestjson1_deserializeOpDescribeEventConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14901,7 +14592,6 @@ func (m *awsRestjson1_deserializeOpDescribeEventConfig
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15045,10 +14735,8 @@ func (m *awsRestjson1_deserializeOpDescribeFleetMetric
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15088,7 +14776,6 @@ func (m *awsRestjson1_deserializeOpDescribeFleetMetric
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15342,10 +15029,8 @@ func (m *awsRestjson1_deserializeOpDescribeIndex) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15385,7 +15070,6 @@ func (m *awsRestjson1_deserializeOpDescribeIndex) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15531,10 +15215,8 @@ func (m *awsRestjson1_deserializeOpDescribeJob) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15574,7 +15256,6 @@ func (m *awsRestjson1_deserializeOpDescribeJob) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15701,10 +15382,8 @@ func (m *awsRestjson1_deserializeOpDescribeJobExecutio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15744,7 +15423,6 @@ func (m *awsRestjson1_deserializeOpDescribeJobExecutio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15862,10 +15540,8 @@ func (m *awsRestjson1_deserializeOpDescribeJobTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15905,7 +15581,6 @@ func (m *awsRestjson1_deserializeOpDescribeJobTemplate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16114,10 +15789,8 @@ func (m *awsRestjson1_deserializeOpDescribeManagedJobT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16157,7 +15830,6 @@ func (m *awsRestjson1_deserializeOpDescribeManagedJobT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16325,10 +15997,8 @@ func (m *awsRestjson1_deserializeOpDescribeMitigationA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16368,7 +16038,6 @@ func (m *awsRestjson1_deserializeOpDescribeMitigationA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16563,10 +16232,8 @@ func (m *awsRestjson1_deserializeOpDescribeProvisionin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16606,7 +16273,6 @@ func (m *awsRestjson1_deserializeOpDescribeProvisionin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16835,10 +16501,8 @@ func (m *awsRestjson1_deserializeOpDescribeProvisionin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16878,7 +16542,6 @@ func (m *awsRestjson1_deserializeOpDescribeProvisionin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17041,10 +16704,8 @@ func (m *awsRestjson1_deserializeOpDescribeRoleAlias) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17084,7 +16745,6 @@ func (m *awsRestjson1_deserializeOpDescribeRoleAlias) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17208,10 +16868,8 @@ func (m *awsRestjson1_deserializeOpDescribeScheduledAu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17251,7 +16909,6 @@ func (m *awsRestjson1_deserializeOpDescribeScheduledAu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17414,10 +17071,8 @@ func (m *awsRestjson1_deserializeOpDescribeSecurityPro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17457,7 +17112,6 @@ func (m *awsRestjson1_deserializeOpDescribeSecurityPro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17667,10 +17321,8 @@ func (m *awsRestjson1_deserializeOpDescribeStream) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17710,7 +17362,6 @@ func (m *awsRestjson1_deserializeOpDescribeStream) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17834,10 +17485,8 @@ func (m *awsRestjson1_deserializeOpDescribeThing) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17877,7 +17526,6 @@ func (m *awsRestjson1_deserializeOpDescribeThing) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18068,10 +17716,8 @@ func (m *awsRestjson1_deserializeOpDescribeThingGroup)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18111,7 +17757,6 @@ func (m *awsRestjson1_deserializeOpDescribeThingGroup)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18310,10 +17955,8 @@ func (m *awsRestjson1_deserializeOpDescribeThingRegist
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18353,7 +17996,6 @@ func (m *awsRestjson1_deserializeOpDescribeThingRegist
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18603,10 +18245,8 @@ func (m *awsRestjson1_deserializeOpDescribeThingType) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18646,7 +18286,6 @@ func (m *awsRestjson1_deserializeOpDescribeThingType) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18802,10 +18441,8 @@ func (m *awsRestjson1_deserializeOpDetachPolicy) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18823,7 +18460,6 @@ func (m *awsRestjson1_deserializeOpDetachPolicy) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18911,10 +18547,8 @@ func (m *awsRestjson1_deserializeOpDetachPrincipalPoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18932,7 +18566,6 @@ func (m *awsRestjson1_deserializeOpDetachPrincipalPoli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19020,10 +18653,8 @@ func (m *awsRestjson1_deserializeOpDetachSecurityProfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19035,7 +18666,6 @@ func (m *awsRestjson1_deserializeOpDetachSecurityProfi
 	output := &DetachSecurityProfileOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19117,10 +18747,8 @@ func (m *awsRestjson1_deserializeOpDetachThingPrincipa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19132,7 +18760,6 @@ func (m *awsRestjson1_deserializeOpDetachThingPrincipa
 	output := &DetachThingPrincipalOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19220,10 +18847,8 @@ func (m *awsRestjson1_deserializeOpDisableTopicRule) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19241,7 +18866,6 @@ func (m *awsRestjson1_deserializeOpDisableTopicRule) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19326,10 +18950,8 @@ func (m *awsRestjson1_deserializeOpDisassociateSbomFro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19341,7 +18963,6 @@ func (m *awsRestjson1_deserializeOpDisassociateSbomFro
 	output := &DisassociateSbomFromPackageVersionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19426,10 +19047,8 @@ func (m *awsRestjson1_deserializeOpEnableTopicRule) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19447,7 +19066,6 @@ func (m *awsRestjson1_deserializeOpEnableTopicRule) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19532,10 +19150,8 @@ func (m *awsRestjson1_deserializeOpGetBehaviorModelTra
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19575,7 +19191,6 @@ func (m *awsRestjson1_deserializeOpGetBehaviorModelTra
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19702,10 +19317,8 @@ func (m *awsRestjson1_deserializeOpGetBucketsAggregati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19745,7 +19358,6 @@ func (m *awsRestjson1_deserializeOpGetBucketsAggregati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19891,10 +19503,8 @@ func (m *awsRestjson1_deserializeOpGetCardinality) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19934,7 +19544,6 @@ func (m *awsRestjson1_deserializeOpGetCardinality) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20075,10 +19684,8 @@ func (m *awsRestjson1_deserializeOpGetCommand) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20118,7 +19725,6 @@ func (m *awsRestjson1_deserializeOpGetCommand) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20345,10 +19951,8 @@ func (m *awsRestjson1_deserializeOpGetCommandExecution
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20388,7 +19992,6 @@ func (m *awsRestjson1_deserializeOpGetCommandExecution
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20645,10 +20248,8 @@ func (m *awsRestjson1_deserializeOpGetEffectivePolicie
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20688,7 +20289,6 @@ func (m *awsRestjson1_deserializeOpGetEffectivePolicie
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20815,10 +20415,8 @@ func (m *awsRestjson1_deserializeOpGetIndexingConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20858,7 +20456,6 @@ func (m *awsRestjson1_deserializeOpGetIndexingConfigur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20984,10 +20581,8 @@ func (m *awsRestjson1_deserializeOpGetJobDocument) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21027,7 +20622,6 @@ func (m *awsRestjson1_deserializeOpGetJobDocument) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21149,10 +20743,8 @@ func (m *awsRestjson1_deserializeOpGetLoggingOptions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21192,7 +20784,6 @@ func (m *awsRestjson1_deserializeOpGetLoggingOptions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21320,10 +20911,8 @@ func (m *awsRestjson1_deserializeOpGetOTAUpdate) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21363,7 +20952,6 @@ func (m *awsRestjson1_deserializeOpGetOTAUpdate) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21487,10 +21075,8 @@ func (m *awsRestjson1_deserializeOpGetPackage) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21530,7 +21116,6 @@ func (m *awsRestjson1_deserializeOpGetPackage) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21711,10 +21296,8 @@ func (m *awsRestjson1_deserializeOpGetPackageConfigura
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21754,7 +21337,6 @@ func (m *awsRestjson1_deserializeOpGetPackageConfigura
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21866,10 +21448,8 @@ func (m *awsRestjson1_deserializeOpGetPackageVersion) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21909,7 +21489,6 @@ func (m *awsRestjson1_deserializeOpGetPackageVersion) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22141,10 +21720,8 @@ func (m *awsRestjson1_deserializeOpGetPercentiles) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22184,7 +21761,6 @@ func (m *awsRestjson1_deserializeOpGetPercentiles) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22317,10 +21893,8 @@ func (m *awsRestjson1_deserializeOpGetPolicy) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22360,7 +21934,6 @@ func (m *awsRestjson1_deserializeOpGetPolicy) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22556,10 +22129,8 @@ func (m *awsRestjson1_deserializeOpGetPolicyVersion) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22599,7 +22170,6 @@ func (m *awsRestjson1_deserializeOpGetPolicyVersion) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22804,10 +22374,8 @@ func (m *awsRestjson1_deserializeOpGetRegistrationCode
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22847,7 +22415,6 @@ func (m *awsRestjson1_deserializeOpGetRegistrationCode
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22972,10 +22539,8 @@ func (m *awsRestjson1_deserializeOpGetStatistics) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23015,7 +22580,6 @@ func (m *awsRestjson1_deserializeOpGetStatistics) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23148,10 +22712,8 @@ func (m *awsRestjson1_deserializeOpGetThingConnectivit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23191,7 +22753,6 @@ func (m *awsRestjson1_deserializeOpGetThingConnectivit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23356,10 +22917,8 @@ func (m *awsRestjson1_deserializeOpGetTopicRule) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23399,7 +22958,6 @@ func (m *awsRestjson1_deserializeOpGetTopicRule) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23526,10 +23084,8 @@ func (m *awsRestjson1_deserializeOpGetTopicRuleDestina
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23569,7 +23125,6 @@ func (m *awsRestjson1_deserializeOpGetTopicRuleDestina
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23687,10 +23242,8 @@ func (m *awsRestjson1_deserializeOpGetV2LoggingOptions
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23730,7 +23283,6 @@ func (m *awsRestjson1_deserializeOpGetV2LoggingOptions
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23867,10 +23419,8 @@ func (m *awsRestjson1_deserializeOpListActiveViolation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23910,7 +23460,6 @@ func (m *awsRestjson1_deserializeOpListActiveViolation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24037,10 +23586,8 @@ func (m *awsRestjson1_deserializeOpListAttachedPolicie
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24080,7 +23627,6 @@ func (m *awsRestjson1_deserializeOpListAttachedPolicie
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24216,10 +23762,8 @@ func (m *awsRestjson1_deserializeOpListAuditFindings) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24259,7 +23803,6 @@ func (m *awsRestjson1_deserializeOpListAuditFindings) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24383,10 +23926,8 @@ func (m *awsRestjson1_deserializeOpListAuditMitigation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24426,7 +23967,6 @@ func (m *awsRestjson1_deserializeOpListAuditMitigation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24550,10 +24090,8 @@ func (m *awsRestjson1_deserializeOpListAuditMitigation
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24593,7 +24131,6 @@ func (m *awsRestjson1_deserializeOpListAuditMitigation
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24717,10 +24254,8 @@ func (m *awsRestjson1_deserializeOpListAuditSuppressio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24760,7 +24295,6 @@ func (m *awsRestjson1_deserializeOpListAuditSuppressio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24884,10 +24418,8 @@ func (m *awsRestjson1_deserializeOpListAuditTasks) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24927,7 +24459,6 @@ func (m *awsRestjson1_deserializeOpListAuditTasks) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25051,10 +24582,8 @@ func (m *awsRestjson1_deserializeOpListAuthorizers) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25094,7 +24623,6 @@ func (m *awsRestjson1_deserializeOpListAuthorizers) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25224,10 +24752,8 @@ func (m *awsRestjson1_deserializeOpListBillingGroups) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25267,7 +24793,6 @@ func (m *awsRestjson1_deserializeOpListBillingGroups) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25394,10 +24919,8 @@ func (m *awsRestjson1_deserializeOpListCACertificates)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25437,7 +24960,6 @@ func (m *awsRestjson1_deserializeOpListCACertificates)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25567,10 +25089,8 @@ func (m *awsRestjson1_deserializeOpListCertificateProv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25610,7 +25130,6 @@ func (m *awsRestjson1_deserializeOpListCertificateProv
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25740,10 +25259,8 @@ func (m *awsRestjson1_deserializeOpListCertificates) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25783,7 +25300,6 @@ func (m *awsRestjson1_deserializeOpListCertificates) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25913,10 +25429,8 @@ func (m *awsRestjson1_deserializeOpListCertificatesByC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25956,7 +25470,6 @@ func (m *awsRestjson1_deserializeOpListCertificatesByC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26086,10 +25599,8 @@ func (m *awsRestjson1_deserializeOpListCommandExecutio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26129,7 +25640,6 @@ func (m *awsRestjson1_deserializeOpListCommandExecutio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26256,10 +25766,8 @@ func (m *awsRestjson1_deserializeOpListCommands) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26299,7 +25807,6 @@ func (m *awsRestjson1_deserializeOpListCommands) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26423,10 +25930,8 @@ func (m *awsRestjson1_deserializeOpListCustomMetrics) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26466,7 +25971,6 @@ func (m *awsRestjson1_deserializeOpListCustomMetrics) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26590,10 +26094,8 @@ func (m *awsRestjson1_deserializeOpListDetectMitigatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26633,7 +26135,6 @@ func (m *awsRestjson1_deserializeOpListDetectMitigatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26757,10 +26258,8 @@ func (m *awsRestjson1_deserializeOpListDetectMitigatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26800,7 +26299,6 @@ func (m *awsRestjson1_deserializeOpListDetectMitigatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26924,10 +26422,8 @@ func (m *awsRestjson1_deserializeOpListDimensions) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26967,7 +26463,6 @@ func (m *awsRestjson1_deserializeOpListDimensions) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27091,10 +26586,8 @@ func (m *awsRestjson1_deserializeOpListDomainConfigura
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27134,7 +26627,6 @@ func (m *awsRestjson1_deserializeOpListDomainConfigura
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27264,10 +26756,8 @@ func (m *awsRestjson1_deserializeOpListFleetMetrics) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27307,7 +26797,6 @@ func (m *awsRestjson1_deserializeOpListFleetMetrics) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27437,10 +26926,8 @@ func (m *awsRestjson1_deserializeOpListIndices) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27480,7 +26967,6 @@ func (m *awsRestjson1_deserializeOpListIndices) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27610,10 +27096,8 @@ func (m *awsRestjson1_deserializeOpListJobExecutionsFo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27653,7 +27137,6 @@ func (m *awsRestjson1_deserializeOpListJobExecutionsFo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27780,10 +27263,8 @@ func (m *awsRestjson1_deserializeOpListJobExecutionsFo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27823,7 +27304,6 @@ func (m *awsRestjson1_deserializeOpListJobExecutionsFo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27950,10 +27430,8 @@ func (m *awsRestjson1_deserializeOpListJobs) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27993,7 +27471,6 @@ func (m *awsRestjson1_deserializeOpListJobs) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28120,10 +27597,8 @@ func (m *awsRestjson1_deserializeOpListJobTemplates) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28163,7 +27638,6 @@ func (m *awsRestjson1_deserializeOpListJobTemplates) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28287,10 +27761,8 @@ func (m *awsRestjson1_deserializeOpListManagedJobTempl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28330,7 +27802,6 @@ func (m *awsRestjson1_deserializeOpListManagedJobTempl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28457,10 +27928,8 @@ func (m *awsRestjson1_deserializeOpListMetricValues) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28500,7 +27969,6 @@ func (m *awsRestjson1_deserializeOpListMetricValues) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28627,10 +28095,8 @@ func (m *awsRestjson1_deserializeOpListMitigationActio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28670,7 +28136,6 @@ func (m *awsRestjson1_deserializeOpListMitigationActio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28794,10 +28259,8 @@ func (m *awsRestjson1_deserializeOpListOTAUpdates) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28837,7 +28300,6 @@ func (m *awsRestjson1_deserializeOpListOTAUpdates) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28967,10 +28429,8 @@ func (m *awsRestjson1_deserializeOpListOutgoingCertifi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29010,7 +28470,6 @@ func (m *awsRestjson1_deserializeOpListOutgoingCertifi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29140,10 +28599,8 @@ func (m *awsRestjson1_deserializeOpListPackages) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29183,7 +28640,6 @@ func (m *awsRestjson1_deserializeOpListPackages) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29307,10 +28763,8 @@ func (m *awsRestjson1_deserializeOpListPackageVersions
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29350,7 +28804,6 @@ func (m *awsRestjson1_deserializeOpListPackageVersions
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29474,10 +28927,8 @@ func (m *awsRestjson1_deserializeOpListPolicies) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29517,7 +28968,6 @@ func (m *awsRestjson1_deserializeOpListPolicies) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29647,10 +29097,8 @@ func (m *awsRestjson1_deserializeOpListPolicyPrincipal
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29690,7 +29138,6 @@ func (m *awsRestjson1_deserializeOpListPolicyPrincipal
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29823,10 +29270,8 @@ func (m *awsRestjson1_deserializeOpListPolicyVersions)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29866,7 +29311,6 @@ func (m *awsRestjson1_deserializeOpListPolicyVersions)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29990,10 +29434,8 @@ func (m *awsRestjson1_deserializeOpListPrincipalPolici
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30033,7 +29475,6 @@ func (m *awsRestjson1_deserializeOpListPrincipalPolici
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30166,10 +29607,8 @@ func (m *awsRestjson1_deserializeOpListPrincipalThings
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30209,7 +29648,6 @@ func (m *awsRestjson1_deserializeOpListPrincipalThings
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30342,10 +29780,8 @@ func (m *awsRestjson1_deserializeOpListPrincipalThings
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30385,7 +29821,6 @@ func (m *awsRestjson1_deserializeOpListPrincipalThings
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30518,10 +29953,8 @@ func (m *awsRestjson1_deserializeOpListProvisioningTem
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30561,7 +29994,6 @@ func (m *awsRestjson1_deserializeOpListProvisioningTem
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30688,10 +30120,8 @@ func (m *awsRestjson1_deserializeOpListProvisioningTem
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30731,7 +30161,6 @@ func (m *awsRestjson1_deserializeOpListProvisioningTem
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30861,10 +30290,8 @@ func (m *awsRestjson1_deserializeOpListRelatedResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30904,7 +30331,6 @@ func (m *awsRestjson1_deserializeOpListRelatedResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31031,10 +30457,8 @@ func (m *awsRestjson1_deserializeOpListRoleAliases) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31074,7 +30498,6 @@ func (m *awsRestjson1_deserializeOpListRoleAliases) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31204,10 +30627,8 @@ func (m *awsRestjson1_deserializeOpListSbomValidationR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31247,7 +30668,6 @@ func (m *awsRestjson1_deserializeOpListSbomValidationR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31374,10 +30794,8 @@ func (m *awsRestjson1_deserializeOpListScheduledAudits
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31417,7 +30835,6 @@ func (m *awsRestjson1_deserializeOpListScheduledAudits
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31541,10 +30958,8 @@ func (m *awsRestjson1_deserializeOpListSecurityProfile
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31584,7 +30999,6 @@ func (m *awsRestjson1_deserializeOpListSecurityProfile
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31711,10 +31125,8 @@ func (m *awsRestjson1_deserializeOpListSecurityProfile
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31754,7 +31166,6 @@ func (m *awsRestjson1_deserializeOpListSecurityProfile
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31881,10 +31292,8 @@ func (m *awsRestjson1_deserializeOpListStreams) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31924,7 +31333,6 @@ func (m *awsRestjson1_deserializeOpListStreams) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32054,10 +31462,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32097,7 +31503,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32224,10 +31629,8 @@ func (m *awsRestjson1_deserializeOpListTargetsForPolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32267,7 +31670,6 @@ func (m *awsRestjson1_deserializeOpListTargetsForPolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32403,10 +31805,8 @@ func (m *awsRestjson1_deserializeOpListTargetsForSecur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32446,7 +31846,6 @@ func (m *awsRestjson1_deserializeOpListTargetsForSecur
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32573,10 +31972,8 @@ func (m *awsRestjson1_deserializeOpListThingGroups) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32616,7 +32013,6 @@ func (m *awsRestjson1_deserializeOpListThingGroups) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32743,10 +32139,8 @@ func (m *awsRestjson1_deserializeOpListThingGroupsForT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32786,7 +32180,6 @@ func (m *awsRestjson1_deserializeOpListThingGroupsForT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32913,10 +32306,8 @@ func (m *awsRestjson1_deserializeOpListThingPrincipals
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32956,7 +32347,6 @@ func (m *awsRestjson1_deserializeOpListThingPrincipals
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33089,10 +32479,8 @@ func (m *awsRestjson1_deserializeOpListThingPrincipals
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -33132,7 +32520,6 @@ func (m *awsRestjson1_deserializeOpListThingPrincipals
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33265,10 +32652,8 @@ func (m *awsRestjson1_deserializeOpListThingRegistrati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -33308,7 +32693,6 @@ func (m *awsRestjson1_deserializeOpListThingRegistrati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33444,10 +32828,8 @@ func (m *awsRestjson1_deserializeOpListThingRegistrati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -33487,7 +32869,6 @@ func (m *awsRestjson1_deserializeOpListThingRegistrati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33614,10 +32995,8 @@ func (m *awsRestjson1_deserializeOpListThings) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -33657,7 +33036,6 @@ func (m *awsRestjson1_deserializeOpListThings) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33787,10 +33165,8 @@ func (m *awsRestjson1_deserializeOpListThingsInBilling
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -33830,7 +33206,6 @@ func (m *awsRestjson1_deserializeOpListThingsInBilling
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33957,10 +33332,8 @@ func (m *awsRestjson1_deserializeOpListThingsInThingGr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -34000,7 +33373,6 @@ func (m *awsRestjson1_deserializeOpListThingsInThingGr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -34127,10 +33499,8 @@ func (m *awsRestjson1_deserializeOpListThingTypes) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -34170,7 +33540,6 @@ func (m *awsRestjson1_deserializeOpListThingTypes) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -34300,10 +33669,8 @@ func (m *awsRestjson1_deserializeOpListTopicRuleDestin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -34343,7 +33710,6 @@ func (m *awsRestjson1_deserializeOpListTopicRuleDestin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -34470,10 +33836,8 @@ func (m *awsRestjson1_deserializeOpListTopicRules) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -34513,7 +33877,6 @@ func (m *awsRestjson1_deserializeOpListTopicRules) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -34637,10 +34000,8 @@ func (m *awsRestjson1_deserializeOpListV2LoggingLevels
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -34680,7 +34041,6 @@ func (m *awsRestjson1_deserializeOpListV2LoggingLevels
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -34807,10 +34167,8 @@ func (m *awsRestjson1_deserializeOpListViolationEvents
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -34850,7 +34208,6 @@ func (m *awsRestjson1_deserializeOpListViolationEvents
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -34974,10 +34331,8 @@ func (m *awsRestjson1_deserializeOpPutVerificationStat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -34989,7 +34344,6 @@ func (m *awsRestjson1_deserializeOpPutVerificationStat
 	output := &PutVerificationStateOnViolationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35068,10 +34422,8 @@ func (m *awsRestjson1_deserializeOpRegisterCACertifica
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35111,7 +34463,6 @@ func (m *awsRestjson1_deserializeOpRegisterCACertifica
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35260,10 +34611,8 @@ func (m *awsRestjson1_deserializeOpRegisterCertificate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35303,7 +34652,6 @@ func (m *awsRestjson1_deserializeOpRegisterCertificate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35449,10 +34797,8 @@ func (m *awsRestjson1_deserializeOpRegisterCertificate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35492,7 +34838,6 @@ func (m *awsRestjson1_deserializeOpRegisterCertificate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35635,10 +34980,8 @@ func (m *awsRestjson1_deserializeOpRegisterThing) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35678,7 +35021,6 @@ func (m *awsRestjson1_deserializeOpRegisterThing) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35814,10 +35156,8 @@ func (m *awsRestjson1_deserializeOpRejectCertificateTr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35835,7 +35175,6 @@ func (m *awsRestjson1_deserializeOpRejectCertificateTr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -35926,10 +35265,8 @@ func (m *awsRestjson1_deserializeOpRemoveThingFromBill
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -35941,7 +35278,6 @@ func (m *awsRestjson1_deserializeOpRemoveThingFromBill
 	output := &RemoveThingFromBillingGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36023,10 +35359,8 @@ func (m *awsRestjson1_deserializeOpRemoveThingFromThin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36038,7 +35372,6 @@ func (m *awsRestjson1_deserializeOpRemoveThingFromThin
 	output := &RemoveThingFromThingGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36120,10 +35453,8 @@ func (m *awsRestjson1_deserializeOpReplaceTopicRule) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36141,7 +35472,6 @@ func (m *awsRestjson1_deserializeOpReplaceTopicRule) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36229,10 +35559,8 @@ func (m *awsRestjson1_deserializeOpSearchIndex) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36272,7 +35600,6 @@ func (m *awsRestjson1_deserializeOpSearchIndex) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36416,10 +35743,8 @@ func (m *awsRestjson1_deserializeOpSetDefaultAuthorize
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36459,7 +35784,6 @@ func (m *awsRestjson1_deserializeOpSetDefaultAuthorize
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36599,10 +35923,8 @@ func (m *awsRestjson1_deserializeOpSetDefaultPolicyVer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36620,7 +35942,6 @@ func (m *awsRestjson1_deserializeOpSetDefaultPolicyVer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36708,10 +36029,8 @@ func (m *awsRestjson1_deserializeOpSetLoggingOptions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36729,7 +36048,6 @@ func (m *awsRestjson1_deserializeOpSetLoggingOptions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36808,10 +36126,8 @@ func (m *awsRestjson1_deserializeOpSetV2LoggingLevel) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36829,7 +36145,6 @@ func (m *awsRestjson1_deserializeOpSetV2LoggingLevel) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -36914,10 +36229,8 @@ func (m *awsRestjson1_deserializeOpSetV2LoggingOptions
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -36935,7 +36248,6 @@ func (m *awsRestjson1_deserializeOpSetV2LoggingOptions
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37014,10 +36326,8 @@ func (m *awsRestjson1_deserializeOpStartAuditMitigatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37057,7 +36367,6 @@ func (m *awsRestjson1_deserializeOpStartAuditMitigatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37182,10 +36491,8 @@ func (m *awsRestjson1_deserializeOpStartDetectMitigati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37225,7 +36532,6 @@ func (m *awsRestjson1_deserializeOpStartDetectMitigati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37350,10 +36656,8 @@ func (m *awsRestjson1_deserializeOpStartOnDemandAuditT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37393,7 +36697,6 @@ func (m *awsRestjson1_deserializeOpStartOnDemandAuditT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37515,10 +36818,8 @@ func (m *awsRestjson1_deserializeOpStartThingRegistrat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37558,7 +36859,6 @@ func (m *awsRestjson1_deserializeOpStartThingRegistrat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37680,10 +36980,8 @@ func (m *awsRestjson1_deserializeOpStopThingRegistrati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37695,7 +36993,6 @@ func (m *awsRestjson1_deserializeOpStopThingRegistrati
 	output := &StopThingRegistrationTaskOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37780,10 +37077,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37795,7 +37090,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -37880,10 +37174,8 @@ func (m *awsRestjson1_deserializeOpTestAuthorization) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -37923,7 +37215,6 @@ func (m *awsRestjson1_deserializeOpTestAuthorization) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38050,10 +37341,8 @@ func (m *awsRestjson1_deserializeOpTestInvokeAuthorize
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38093,7 +37382,6 @@ func (m *awsRestjson1_deserializeOpTestInvokeAuthorize
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38264,10 +37552,8 @@ func (m *awsRestjson1_deserializeOpTransferCertificate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38307,7 +37593,6 @@ func (m *awsRestjson1_deserializeOpTransferCertificate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38441,10 +37726,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38456,7 +37739,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38538,10 +37820,8 @@ func (m *awsRestjson1_deserializeOpUpdateAccountAuditC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38553,7 +37833,6 @@ func (m *awsRestjson1_deserializeOpUpdateAccountAuditC
 	output := &UpdateAccountAuditConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38632,10 +37911,8 @@ func (m *awsRestjson1_deserializeOpUpdateAuditSuppress
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38647,7 +37924,6 @@ func (m *awsRestjson1_deserializeOpUpdateAuditSuppress
 	output := &UpdateAuditSuppressionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38729,10 +38005,8 @@ func (m *awsRestjson1_deserializeOpUpdateAuthorizer) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38772,7 +38046,6 @@ func (m *awsRestjson1_deserializeOpUpdateAuthorizer) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -38912,10 +38185,8 @@ func (m *awsRestjson1_deserializeOpUpdateBillingGroup)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -38955,7 +38226,6 @@ func (m *awsRestjson1_deserializeOpUpdateBillingGroup)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39084,10 +38354,8 @@ func (m *awsRestjson1_deserializeOpUpdateCACertificate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39105,7 +38373,6 @@ func (m *awsRestjson1_deserializeOpUpdateCACertificate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39193,10 +38460,8 @@ func (m *awsRestjson1_deserializeOpUpdateCertificate) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39214,7 +38479,6 @@ func (m *awsRestjson1_deserializeOpUpdateCertificate) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39305,10 +38569,8 @@ func (m *awsRestjson1_deserializeOpUpdateCertificatePr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39348,7 +38610,6 @@ func (m *awsRestjson1_deserializeOpUpdateCertificatePr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39485,10 +38746,8 @@ func (m *awsRestjson1_deserializeOpUpdateCommand) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39528,7 +38787,6 @@ func (m *awsRestjson1_deserializeOpUpdateCommand) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39696,10 +38954,8 @@ func (m *awsRestjson1_deserializeOpUpdateCustomMetric)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39739,7 +38995,6 @@ func (m *awsRestjson1_deserializeOpUpdateCustomMetric)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -39920,10 +39175,8 @@ func (m *awsRestjson1_deserializeOpUpdateDimension) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -39963,7 +39216,6 @@ func (m *awsRestjson1_deserializeOpUpdateDimension) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40140,10 +39392,8 @@ func (m *awsRestjson1_deserializeOpUpdateDomainConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40183,7 +39433,6 @@ func (m *awsRestjson1_deserializeOpUpdateDomainConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40323,10 +39572,8 @@ func (m *awsRestjson1_deserializeOpUpdateDynamicThingG
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40366,7 +39613,6 @@ func (m *awsRestjson1_deserializeOpUpdateDynamicThingG
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40498,10 +39744,8 @@ func (m *awsRestjson1_deserializeOpUpdateEventConfigur
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40513,7 +39757,6 @@ func (m *awsRestjson1_deserializeOpUpdateEventConfigur
 	output := &UpdateEventConfigurationsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40592,10 +39835,8 @@ func (m *awsRestjson1_deserializeOpUpdateFleetMetric) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40613,7 +39854,6 @@ func (m *awsRestjson1_deserializeOpUpdateFleetMetric) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40713,10 +39953,8 @@ func (m *awsRestjson1_deserializeOpUpdateIndexingConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40728,7 +39966,6 @@ func (m *awsRestjson1_deserializeOpUpdateIndexingConfi
 	output := &UpdateIndexingConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40813,10 +40050,8 @@ func (m *awsRestjson1_deserializeOpUpdateJob) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40834,7 +40069,6 @@ func (m *awsRestjson1_deserializeOpUpdateJob) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -40916,10 +40150,8 @@ func (m *awsRestjson1_deserializeOpUpdateMitigationAct
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -40959,7 +40191,6 @@ func (m *awsRestjson1_deserializeOpUpdateMitigationAct
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41090,10 +40321,8 @@ func (m *awsRestjson1_deserializeOpUpdatePackage) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41105,7 +40334,6 @@ func (m *awsRestjson1_deserializeOpUpdatePackage) Hand
 	output := &UpdatePackageOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41190,10 +40418,8 @@ func (m *awsRestjson1_deserializeOpUpdatePackageConfig
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41205,7 +40431,6 @@ func (m *awsRestjson1_deserializeOpUpdatePackageConfig
 	output := &UpdatePackageConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41287,10 +40512,8 @@ func (m *awsRestjson1_deserializeOpUpdatePackageVersio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41302,7 +40525,6 @@ func (m *awsRestjson1_deserializeOpUpdatePackageVersio
 	output := &UpdatePackageVersionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41387,10 +40609,8 @@ func (m *awsRestjson1_deserializeOpUpdateProvisioningT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41402,7 +40622,6 @@ func (m *awsRestjson1_deserializeOpUpdateProvisioningT
 	output := &UpdateProvisioningTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41487,10 +40706,8 @@ func (m *awsRestjson1_deserializeOpUpdateRoleAlias) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41530,7 +40747,6 @@ func (m *awsRestjson1_deserializeOpUpdateRoleAlias) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41667,10 +40883,8 @@ func (m *awsRestjson1_deserializeOpUpdateScheduledAudi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41710,7 +40924,6 @@ func (m *awsRestjson1_deserializeOpUpdateScheduledAudi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -41832,10 +41045,8 @@ func (m *awsRestjson1_deserializeOpUpdateSecurityProfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -41875,7 +41086,6 @@ func (m *awsRestjson1_deserializeOpUpdateSecurityProfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -42088,10 +41298,8 @@ func (m *awsRestjson1_deserializeOpUpdateStream) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -42131,7 +41339,6 @@ func (m *awsRestjson1_deserializeOpUpdateStream) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -42293,10 +41500,8 @@ func (m *awsRestjson1_deserializeOpUpdateThing) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -42308,7 +41513,6 @@ func (m *awsRestjson1_deserializeOpUpdateThing) Handle
 	output := &UpdateThingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -42399,10 +41603,8 @@ func (m *awsRestjson1_deserializeOpUpdateThingGroup) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -42442,7 +41644,6 @@ func (m *awsRestjson1_deserializeOpUpdateThingGroup) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -42571,10 +41772,8 @@ func (m *awsRestjson1_deserializeOpUpdateThingGroupsFo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -42586,7 +41785,6 @@ func (m *awsRestjson1_deserializeOpUpdateThingGroupsFo
 	output := &UpdateThingGroupsForThingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -42668,10 +41866,8 @@ func (m *awsRestjson1_deserializeOpUpdateThingType) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -42683,7 +41879,6 @@ func (m *awsRestjson1_deserializeOpUpdateThingType) Ha
 	output := &UpdateThingTypeOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -42771,10 +41966,8 @@ func (m *awsRestjson1_deserializeOpUpdateTopicRuleDest
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -42786,7 +41979,6 @@ func (m *awsRestjson1_deserializeOpUpdateTopicRuleDest
 	output := &UpdateTopicRuleDestinationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -42871,10 +42063,8 @@ func (m *awsRestjson1_deserializeOpValidateSecurityPro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -42914,7 +42104,6 @@ func (m *awsRestjson1_deserializeOpValidateSecurityPro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
