--- vendor/github.com/aws/aws-sdk-go-v2/service/auditmanager/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/auditmanager/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpAssociateAssessment
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -59,7 +56,6 @@ func (m *awsRestjson1_deserializeOpAssociateAssessment
 	output := &AssociateAssessmentReportEvidenceFolderOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -141,10 +137,8 @@ func (m *awsRestjson1_deserializeOpBatchAssociateAsses
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -184,7 +178,6 @@ func (m *awsRestjson1_deserializeOpBatchAssociateAsses
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -307,10 +300,8 @@ func (m *awsRestjson1_deserializeOpBatchCreateDelegati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -350,7 +341,6 @@ func (m *awsRestjson1_deserializeOpBatchCreateDelegati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -473,10 +463,8 @@ func (m *awsRestjson1_deserializeOpBatchDeleteDelegati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -516,7 +504,6 @@ func (m *awsRestjson1_deserializeOpBatchDeleteDelegati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -634,10 +621,8 @@ func (m *awsRestjson1_deserializeOpBatchDisassociateAs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -677,7 +662,6 @@ func (m *awsRestjson1_deserializeOpBatchDisassociateAs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -800,10 +784,8 @@ func (m *awsRestjson1_deserializeOpBatchImportEvidence
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -843,7 +825,6 @@ func (m *awsRestjson1_deserializeOpBatchImportEvidence
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -964,10 +945,8 @@ func (m *awsRestjson1_deserializeOpCreateAssessment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1007,7 +986,6 @@ func (m *awsRestjson1_deserializeOpCreateAssessment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1131,10 +1109,8 @@ func (m *awsRestjson1_deserializeOpCreateAssessmentFra
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1174,7 +1150,6 @@ func (m *awsRestjson1_deserializeOpCreateAssessmentFra
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1295,10 +1270,8 @@ func (m *awsRestjson1_deserializeOpCreateAssessmentRep
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1338,7 +1311,6 @@ func (m *awsRestjson1_deserializeOpCreateAssessmentRep
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1456,10 +1428,8 @@ func (m *awsRestjson1_deserializeOpCreateControl) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1499,7 +1469,6 @@ func (m *awsRestjson1_deserializeOpCreateControl) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1620,10 +1589,8 @@ func (m *awsRestjson1_deserializeOpDeleteAssessment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1635,7 +1602,6 @@ func (m *awsRestjson1_deserializeOpDeleteAssessment) H
 	output := &DeleteAssessmentOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1717,10 +1683,8 @@ func (m *awsRestjson1_deserializeOpDeleteAssessmentFra
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1732,7 +1696,6 @@ func (m *awsRestjson1_deserializeOpDeleteAssessmentFra
 	output := &DeleteAssessmentFrameworkOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1814,10 +1777,8 @@ func (m *awsRestjson1_deserializeOpDeleteAssessmentFra
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1829,7 +1790,6 @@ func (m *awsRestjson1_deserializeOpDeleteAssessmentFra
 	output := &DeleteAssessmentFrameworkShareOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1911,10 +1871,8 @@ func (m *awsRestjson1_deserializeOpDeleteAssessmentRep
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1926,7 +1884,6 @@ func (m *awsRestjson1_deserializeOpDeleteAssessmentRep
 	output := &DeleteAssessmentReportOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2008,10 +1965,8 @@ func (m *awsRestjson1_deserializeOpDeleteControl) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2023,7 +1978,6 @@ func (m *awsRestjson1_deserializeOpDeleteControl) Hand
 	output := &DeleteControlOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2105,10 +2059,8 @@ func (m *awsRestjson1_deserializeOpDeregisterAccount) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2148,7 +2100,6 @@ func (m *awsRestjson1_deserializeOpDeregisterAccount) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2270,10 +2221,8 @@ func (m *awsRestjson1_deserializeOpDeregisterOrganizat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2285,7 +2234,6 @@ func (m *awsRestjson1_deserializeOpDeregisterOrganizat
 	output := &DeregisterOrganizationAdminAccountOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2367,10 +2315,8 @@ func (m *awsRestjson1_deserializeOpDisassociateAssessm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2382,7 +2328,6 @@ func (m *awsRestjson1_deserializeOpDisassociateAssessm
 	output := &DisassociateAssessmentReportEvidenceFolderOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2464,10 +2409,8 @@ func (m *awsRestjson1_deserializeOpGetAccountStatus) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2507,7 +2450,6 @@ func (m *awsRestjson1_deserializeOpGetAccountStatus) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2620,10 +2562,8 @@ func (m *awsRestjson1_deserializeOpGetAssessment) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2663,7 +2603,6 @@ func (m *awsRestjson1_deserializeOpGetAssessment) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2786,10 +2725,8 @@ func (m *awsRestjson1_deserializeOpGetAssessmentFramew
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2829,7 +2766,6 @@ func (m *awsRestjson1_deserializeOpGetAssessmentFramew
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2947,10 +2883,8 @@ func (m *awsRestjson1_deserializeOpGetAssessmentReport
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2990,7 +2924,6 @@ func (m *awsRestjson1_deserializeOpGetAssessmentReport
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3108,10 +3041,8 @@ func (m *awsRestjson1_deserializeOpGetChangeLogs) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3151,7 +3082,6 @@ func (m *awsRestjson1_deserializeOpGetChangeLogs) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3278,10 +3208,8 @@ func (m *awsRestjson1_deserializeOpGetControl) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3321,7 +3249,6 @@ func (m *awsRestjson1_deserializeOpGetControl) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3439,10 +3366,8 @@ func (m *awsRestjson1_deserializeOpGetDelegations) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3482,7 +3407,6 @@ func (m *awsRestjson1_deserializeOpGetDelegations) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3606,10 +3530,8 @@ func (m *awsRestjson1_deserializeOpGetEvidence) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3649,7 +3571,6 @@ func (m *awsRestjson1_deserializeOpGetEvidence) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3767,10 +3688,8 @@ func (m *awsRestjson1_deserializeOpGetEvidenceByEviden
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3810,7 +3729,6 @@ func (m *awsRestjson1_deserializeOpGetEvidenceByEviden
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3937,10 +3855,8 @@ func (m *awsRestjson1_deserializeOpGetEvidenceFileUplo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3980,7 +3896,6 @@ func (m *awsRestjson1_deserializeOpGetEvidenceFileUplo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4111,10 +4026,8 @@ func (m *awsRestjson1_deserializeOpGetEvidenceFolder) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4154,7 +4067,6 @@ func (m *awsRestjson1_deserializeOpGetEvidenceFolder) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4272,10 +4184,8 @@ func (m *awsRestjson1_deserializeOpGetEvidenceFoldersB
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4315,7 +4225,6 @@ func (m *awsRestjson1_deserializeOpGetEvidenceFoldersB
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4442,10 +4351,8 @@ func (m *awsRestjson1_deserializeOpGetEvidenceFoldersB
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4485,7 +4392,6 @@ func (m *awsRestjson1_deserializeOpGetEvidenceFoldersB
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4612,10 +4518,8 @@ func (m *awsRestjson1_deserializeOpGetInsights) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4655,7 +4559,6 @@ func (m *awsRestjson1_deserializeOpGetInsights) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4767,10 +4670,8 @@ func (m *awsRestjson1_deserializeOpGetInsightsByAssess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4810,7 +4711,6 @@ func (m *awsRestjson1_deserializeOpGetInsightsByAssess
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4928,10 +4828,8 @@ func (m *awsRestjson1_deserializeOpGetOrganizationAdmi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4971,7 +4869,6 @@ func (m *awsRestjson1_deserializeOpGetOrganizationAdmi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5102,10 +4999,8 @@ func (m *awsRestjson1_deserializeOpGetServicesInScope)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5145,7 +5040,6 @@ func (m *awsRestjson1_deserializeOpGetServicesInScope)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5260,10 +5154,8 @@ func (m *awsRestjson1_deserializeOpGetSettings) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5303,7 +5195,6 @@ func (m *awsRestjson1_deserializeOpGetSettings) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5415,10 +5306,8 @@ func (m *awsRestjson1_deserializeOpListAssessmentContr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5458,7 +5347,6 @@ func (m *awsRestjson1_deserializeOpListAssessmentContr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5585,10 +5473,8 @@ func (m *awsRestjson1_deserializeOpListAssessmentFrame
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5628,7 +5514,6 @@ func (m *awsRestjson1_deserializeOpListAssessmentFrame
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5752,10 +5637,8 @@ func (m *awsRestjson1_deserializeOpListAssessmentFrame
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5795,7 +5678,6 @@ func (m *awsRestjson1_deserializeOpListAssessmentFrame
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5919,10 +5801,8 @@ func (m *awsRestjson1_deserializeOpListAssessmentRepor
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5962,7 +5842,6 @@ func (m *awsRestjson1_deserializeOpListAssessmentRepor
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6086,10 +5965,8 @@ func (m *awsRestjson1_deserializeOpListAssessments) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6129,7 +6006,6 @@ func (m *awsRestjson1_deserializeOpListAssessments) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6253,10 +6129,8 @@ func (m *awsRestjson1_deserializeOpListControlDomainIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6296,7 +6170,6 @@ func (m *awsRestjson1_deserializeOpListControlDomainIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6423,10 +6296,8 @@ func (m *awsRestjson1_deserializeOpListControlDomainIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6466,7 +6337,6 @@ func (m *awsRestjson1_deserializeOpListControlDomainIn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6593,10 +6463,8 @@ func (m *awsRestjson1_deserializeOpListControlInsights
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6636,7 +6504,6 @@ func (m *awsRestjson1_deserializeOpListControlInsights
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6763,10 +6630,8 @@ func (m *awsRestjson1_deserializeOpListControls) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6806,7 +6671,6 @@ func (m *awsRestjson1_deserializeOpListControls) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6930,10 +6794,8 @@ func (m *awsRestjson1_deserializeOpListKeywordsForData
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6973,7 +6835,6 @@ func (m *awsRestjson1_deserializeOpListKeywordsForData
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7097,10 +6958,8 @@ func (m *awsRestjson1_deserializeOpListNotifications) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7140,7 +6999,6 @@ func (m *awsRestjson1_deserializeOpListNotifications) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7264,10 +7122,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7307,7 +7163,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7422,10 +7277,8 @@ func (m *awsRestjson1_deserializeOpRegisterAccount) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7465,7 +7318,6 @@ func (m *awsRestjson1_deserializeOpRegisterAccount) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7590,10 +7442,8 @@ func (m *awsRestjson1_deserializeOpRegisterOrganizatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7633,7 +7483,6 @@ func (m *awsRestjson1_deserializeOpRegisterOrganizatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7767,10 +7616,8 @@ func (m *awsRestjson1_deserializeOpStartAssessmentFram
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7810,7 +7657,6 @@ func (m *awsRestjson1_deserializeOpStartAssessmentFram
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7928,10 +7774,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7943,7 +7787,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8022,10 +7865,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8037,7 +7878,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8116,10 +7956,8 @@ func (m *awsRestjson1_deserializeOpUpdateAssessment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8159,7 +7997,6 @@ func (m *awsRestjson1_deserializeOpUpdateAssessment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8283,10 +8120,8 @@ func (m *awsRestjson1_deserializeOpUpdateAssessmentCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8326,7 +8161,6 @@ func (m *awsRestjson1_deserializeOpUpdateAssessmentCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8444,10 +8278,8 @@ func (m *awsRestjson1_deserializeOpUpdateAssessmentCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8487,7 +8319,6 @@ func (m *awsRestjson1_deserializeOpUpdateAssessmentCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8605,10 +8436,8 @@ func (m *awsRestjson1_deserializeOpUpdateAssessmentFra
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8648,7 +8477,6 @@ func (m *awsRestjson1_deserializeOpUpdateAssessmentFra
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8766,10 +8594,8 @@ func (m *awsRestjson1_deserializeOpUpdateAssessmentFra
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8809,7 +8635,6 @@ func (m *awsRestjson1_deserializeOpUpdateAssessmentFra
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8930,10 +8755,8 @@ func (m *awsRestjson1_deserializeOpUpdateAssessmentSta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8973,7 +8796,6 @@ func (m *awsRestjson1_deserializeOpUpdateAssessmentSta
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9094,10 +8916,8 @@ func (m *awsRestjson1_deserializeOpUpdateControl) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9137,7 +8957,6 @@ func (m *awsRestjson1_deserializeOpUpdateControl) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9255,10 +9074,8 @@ func (m *awsRestjson1_deserializeOpUpdateSettings) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9298,7 +9115,6 @@ func (m *awsRestjson1_deserializeOpUpdateSettings) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9413,10 +9229,8 @@ func (m *awsRestjson1_deserializeOpValidateAssessmentR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9456,7 +9270,6 @@ func (m *awsRestjson1_deserializeOpValidateAssessmentR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
