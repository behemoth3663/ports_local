--- vendor/github.com/aws/aws-sdk-go-v2/service/macie2/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/macie2/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAcceptInvitation) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -60,7 +57,6 @@ func (m *awsRestjson1_deserializeOpAcceptInvitation) H
 	output := &AcceptInvitationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -151,10 +147,8 @@ func (m *awsRestjson1_deserializeOpBatchGetCustomDataI
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -194,7 +188,6 @@ func (m *awsRestjson1_deserializeOpBatchGetCustomDataI
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -326,10 +319,8 @@ func (m *awsRestjson1_deserializeOpBatchUpdateAutomate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -369,7 +360,6 @@ func (m *awsRestjson1_deserializeOpBatchUpdateAutomate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -490,10 +480,8 @@ func (m *awsRestjson1_deserializeOpCreateAllowList) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -533,7 +521,6 @@ func (m *awsRestjson1_deserializeOpCreateAllowList) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -673,10 +660,8 @@ func (m *awsRestjson1_deserializeOpCreateClassificatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -716,7 +701,6 @@ func (m *awsRestjson1_deserializeOpCreateClassificatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -856,10 +840,8 @@ func (m *awsRestjson1_deserializeOpCreateCustomDataIde
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -899,7 +881,6 @@ func (m *awsRestjson1_deserializeOpCreateCustomDataIde
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1030,10 +1011,8 @@ func (m *awsRestjson1_deserializeOpCreateFindingsFilte
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1073,7 +1052,6 @@ func (m *awsRestjson1_deserializeOpCreateFindingsFilte
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1213,10 +1191,8 @@ func (m *awsRestjson1_deserializeOpCreateInvitations) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1256,7 +1232,6 @@ func (m *awsRestjson1_deserializeOpCreateInvitations) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1383,10 +1358,8 @@ func (m *awsRestjson1_deserializeOpCreateMember) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1426,7 +1399,6 @@ func (m *awsRestjson1_deserializeOpCreateMember) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1557,10 +1529,8 @@ func (m *awsRestjson1_deserializeOpCreateSampleFinding
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1572,7 +1542,6 @@ func (m *awsRestjson1_deserializeOpCreateSampleFinding
 	output := &CreateSampleFindingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1663,10 +1632,8 @@ func (m *awsRestjson1_deserializeOpDeclineInvitations)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1706,7 +1673,6 @@ func (m *awsRestjson1_deserializeOpDeclineInvitations)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1833,10 +1799,8 @@ func (m *awsRestjson1_deserializeOpDeleteAllowList) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1848,7 +1812,6 @@ func (m *awsRestjson1_deserializeOpDeleteAllowList) Ha
 	output := &DeleteAllowListOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1933,10 +1896,8 @@ func (m *awsRestjson1_deserializeOpDeleteCustomDataIde
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1948,7 +1909,6 @@ func (m *awsRestjson1_deserializeOpDeleteCustomDataIde
 	output := &DeleteCustomDataIdentifierOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2039,10 +1999,8 @@ func (m *awsRestjson1_deserializeOpDeleteFindingsFilte
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2054,7 +2012,6 @@ func (m *awsRestjson1_deserializeOpDeleteFindingsFilte
 	output := &DeleteFindingsFilterOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2145,10 +2102,8 @@ func (m *awsRestjson1_deserializeOpDeleteInvitations) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2188,7 +2143,6 @@ func (m *awsRestjson1_deserializeOpDeleteInvitations) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2315,10 +2269,8 @@ func (m *awsRestjson1_deserializeOpDeleteMember) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2330,7 +2282,6 @@ func (m *awsRestjson1_deserializeOpDeleteMember) Handl
 	output := &DeleteMemberOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2421,10 +2372,8 @@ func (m *awsRestjson1_deserializeOpDescribeBuckets) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2464,7 +2413,6 @@ func (m *awsRestjson1_deserializeOpDescribeBuckets) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2600,10 +2548,8 @@ func (m *awsRestjson1_deserializeOpDescribeClassificat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2643,7 +2589,6 @@ func (m *awsRestjson1_deserializeOpDescribeClassificat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2930,10 +2875,8 @@ func (m *awsRestjson1_deserializeOpDescribeOrganizatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2973,7 +2916,6 @@ func (m *awsRestjson1_deserializeOpDescribeOrganizatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3113,10 +3055,8 @@ func (m *awsRestjson1_deserializeOpDisableMacie) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3128,7 +3068,6 @@ func (m *awsRestjson1_deserializeOpDisableMacie) Handl
 	output := &DisableMacieOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3219,10 +3158,8 @@ func (m *awsRestjson1_deserializeOpDisableOrganization
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3234,7 +3171,6 @@ func (m *awsRestjson1_deserializeOpDisableOrganization
 	output := &DisableOrganizationAdminAccountOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3325,10 +3261,8 @@ func (m *awsRestjson1_deserializeOpDisassociateFromAdm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3340,7 +3274,6 @@ func (m *awsRestjson1_deserializeOpDisassociateFromAdm
 	output := &DisassociateFromAdministratorAccountOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3431,10 +3364,8 @@ func (m *awsRestjson1_deserializeOpDisassociateFromMas
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3446,7 +3377,6 @@ func (m *awsRestjson1_deserializeOpDisassociateFromMas
 	output := &DisassociateFromMasterAccountOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3537,10 +3467,8 @@ func (m *awsRestjson1_deserializeOpDisassociateMember)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3552,7 +3480,6 @@ func (m *awsRestjson1_deserializeOpDisassociateMember)
 	output := &DisassociateMemberOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3643,10 +3570,8 @@ func (m *awsRestjson1_deserializeOpEnableMacie) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3658,7 +3583,6 @@ func (m *awsRestjson1_deserializeOpEnableMacie) Handle
 	output := &EnableMacieOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3749,10 +3673,8 @@ func (m *awsRestjson1_deserializeOpEnableOrganizationA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3764,7 +3686,6 @@ func (m *awsRestjson1_deserializeOpEnableOrganizationA
 	output := &EnableOrganizationAdminAccountOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3855,10 +3776,8 @@ func (m *awsRestjson1_deserializeOpGetAdministratorAcc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3898,7 +3817,6 @@ func (m *awsRestjson1_deserializeOpGetAdministratorAcc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4025,10 +3943,8 @@ func (m *awsRestjson1_deserializeOpGetAllowList) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4068,7 +3984,6 @@ func (m *awsRestjson1_deserializeOpGetAllowList) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4261,10 +4176,8 @@ func (m *awsRestjson1_deserializeOpGetAutomatedDiscove
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4304,7 +4217,6 @@ func (m *awsRestjson1_deserializeOpGetAutomatedDiscove
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4492,10 +4404,8 @@ func (m *awsRestjson1_deserializeOpGetBucketStatistics
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4535,7 +4445,6 @@ func (m *awsRestjson1_deserializeOpGetBucketStatistics
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4783,10 +4692,8 @@ func (m *awsRestjson1_deserializeOpGetClassificationEx
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4826,7 +4733,6 @@ func (m *awsRestjson1_deserializeOpGetClassificationEx
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4953,10 +4859,8 @@ func (m *awsRestjson1_deserializeOpGetClassificationSc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4996,7 +4900,6 @@ func (m *awsRestjson1_deserializeOpGetClassificationSc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5135,10 +5038,8 @@ func (m *awsRestjson1_deserializeOpGetCustomDataIdenti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5178,7 +5079,6 @@ func (m *awsRestjson1_deserializeOpGetCustomDataIdenti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5400,10 +5300,8 @@ func (m *awsRestjson1_deserializeOpGetFindings) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5443,7 +5341,6 @@ func (m *awsRestjson1_deserializeOpGetFindings) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5570,10 +5467,8 @@ func (m *awsRestjson1_deserializeOpGetFindingsFilter) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5613,7 +5508,6 @@ func (m *awsRestjson1_deserializeOpGetFindingsFilter) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5803,10 +5697,8 @@ func (m *awsRestjson1_deserializeOpGetFindingsPublicat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5846,7 +5738,6 @@ func (m *awsRestjson1_deserializeOpGetFindingsPublicat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5973,10 +5864,8 @@ func (m *awsRestjson1_deserializeOpGetFindingStatistic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6016,7 +5905,6 @@ func (m *awsRestjson1_deserializeOpGetFindingStatistic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6143,10 +6031,8 @@ func (m *awsRestjson1_deserializeOpGetInvitationsCount
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6186,7 +6072,6 @@ func (m *awsRestjson1_deserializeOpGetInvitationsCount
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6321,10 +6206,8 @@ func (m *awsRestjson1_deserializeOpGetMacieSession) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6364,7 +6247,6 @@ func (m *awsRestjson1_deserializeOpGetMacieSession) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6539,10 +6421,8 @@ func (m *awsRestjson1_deserializeOpGetMasterAccount) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6582,7 +6462,6 @@ func (m *awsRestjson1_deserializeOpGetMasterAccount) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6709,10 +6588,8 @@ func (m *awsRestjson1_deserializeOpGetMember) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6752,7 +6629,6 @@ func (m *awsRestjson1_deserializeOpGetMember) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6959,10 +6835,8 @@ func (m *awsRestjson1_deserializeOpGetResourceProfile)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7002,7 +6876,6 @@ func (m *awsRestjson1_deserializeOpGetResourceProfile)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7161,10 +7034,8 @@ func (m *awsRestjson1_deserializeOpGetRevealConfigurat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7204,7 +7075,6 @@ func (m *awsRestjson1_deserializeOpGetRevealConfigurat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7327,10 +7197,8 @@ func (m *awsRestjson1_deserializeOpGetSensitiveDataOcc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7370,7 +7238,6 @@ func (m *awsRestjson1_deserializeOpGetSensitiveDataOcc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7512,10 +7379,8 @@ func (m *awsRestjson1_deserializeOpGetSensitiveDataOcc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7555,7 +7420,6 @@ func (m *awsRestjson1_deserializeOpGetSensitiveDataOcc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7682,10 +7546,8 @@ func (m *awsRestjson1_deserializeOpGetSensitivityInspe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7725,7 +7587,6 @@ func (m *awsRestjson1_deserializeOpGetSensitivityInspe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7878,10 +7739,8 @@ func (m *awsRestjson1_deserializeOpGetUsageStatistics)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7921,7 +7780,6 @@ func (m *awsRestjson1_deserializeOpGetUsageStatistics)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8066,10 +7924,8 @@ func (m *awsRestjson1_deserializeOpGetUsageTotals) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8109,7 +7965,6 @@ func (m *awsRestjson1_deserializeOpGetUsageTotals) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8245,10 +8100,8 @@ func (m *awsRestjson1_deserializeOpListAllowLists) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8288,7 +8141,6 @@ func (m *awsRestjson1_deserializeOpListAllowLists) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8415,10 +8267,8 @@ func (m *awsRestjson1_deserializeOpListAutomatedDiscov
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8458,7 +8308,6 @@ func (m *awsRestjson1_deserializeOpListAutomatedDiscov
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8588,10 +8437,8 @@ func (m *awsRestjson1_deserializeOpListClassificationJ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8631,7 +8478,6 @@ func (m *awsRestjson1_deserializeOpListClassificationJ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8767,10 +8613,8 @@ func (m *awsRestjson1_deserializeOpListClassificationS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8810,7 +8654,6 @@ func (m *awsRestjson1_deserializeOpListClassificationS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8937,10 +8780,8 @@ func (m *awsRestjson1_deserializeOpListCustomDataIdent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8980,7 +8821,6 @@ func (m *awsRestjson1_deserializeOpListCustomDataIdent
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9116,10 +8956,8 @@ func (m *awsRestjson1_deserializeOpListFindings) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9159,7 +8997,6 @@ func (m *awsRestjson1_deserializeOpListFindings) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9295,10 +9132,8 @@ func (m *awsRestjson1_deserializeOpListFindingsFilters
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9338,7 +9173,6 @@ func (m *awsRestjson1_deserializeOpListFindingsFilters
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9474,10 +9308,8 @@ func (m *awsRestjson1_deserializeOpListInvitations) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9517,7 +9349,6 @@ func (m *awsRestjson1_deserializeOpListInvitations) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9653,10 +9484,8 @@ func (m *awsRestjson1_deserializeOpListManagedDataIden
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9696,7 +9525,6 @@ func (m *awsRestjson1_deserializeOpListManagedDataIden
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9811,10 +9639,8 @@ func (m *awsRestjson1_deserializeOpListMembers) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9854,7 +9680,6 @@ func (m *awsRestjson1_deserializeOpListMembers) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9990,10 +9815,8 @@ func (m *awsRestjson1_deserializeOpListOrganizationAdm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10033,7 +9856,6 @@ func (m *awsRestjson1_deserializeOpListOrganizationAdm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10169,10 +9991,8 @@ func (m *awsRestjson1_deserializeOpListResourceProfile
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10212,7 +10032,6 @@ func (m *awsRestjson1_deserializeOpListResourceProfile
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10342,10 +10161,8 @@ func (m *awsRestjson1_deserializeOpListResourceProfile
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10385,7 +10202,6 @@ func (m *awsRestjson1_deserializeOpListResourceProfile
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10518,10 +10334,8 @@ func (m *awsRestjson1_deserializeOpListSensitivityInsp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10561,7 +10375,6 @@ func (m *awsRestjson1_deserializeOpListSensitivityInsp
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
 
@@ -10840,10 +10650,8 @@ func (m *awsRestjson1_deserializeOpPutClassificationEx
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10883,7 +10691,6 @@ func (m *awsRestjson1_deserializeOpPutClassificationEx
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11010,10 +10817,8 @@ func (m *awsRestjson1_deserializeOpPutFindingsPublicat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11025,7 +10830,6 @@ func (m *awsRestjson1_deserializeOpPutFindingsPublicat
 	output := &PutFindingsPublicationConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11116,10 +10920,8 @@ func (m *awsRestjson1_deserializeOpSearchResources) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11159,7 +10961,6 @@ func (m *awsRestjson1_deserializeOpSearchResources) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11295,10 +11096,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11310,7 +11109,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11380,10 +11178,8 @@ func (m *awsRestjson1_deserializeOpTestCustomDataIdent
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11423,7 +11219,6 @@ func (m *awsRestjson1_deserializeOpTestCustomDataIdent
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11558,10 +11353,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11573,7 +11366,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11643,10 +11435,8 @@ func (m *awsRestjson1_deserializeOpUpdateAllowList) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11686,7 +11476,6 @@ func (m *awsRestjson1_deserializeOpUpdateAllowList) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11820,10 +11609,8 @@ func (m *awsRestjson1_deserializeOpUpdateAutomatedDisc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11835,7 +11622,6 @@ func (m *awsRestjson1_deserializeOpUpdateAutomatedDisc
 	output := &UpdateAutomatedDiscoveryConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11917,10 +11703,8 @@ func (m *awsRestjson1_deserializeOpUpdateClassificatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11932,7 +11716,6 @@ func (m *awsRestjson1_deserializeOpUpdateClassificatio
 	output := &UpdateClassificationJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12023,10 +11806,8 @@ func (m *awsRestjson1_deserializeOpUpdateClassificatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12038,7 +11819,6 @@ func (m *awsRestjson1_deserializeOpUpdateClassificatio
 	output := &UpdateClassificationScopeOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12123,10 +11903,8 @@ func (m *awsRestjson1_deserializeOpUpdateFindingsFilte
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12166,7 +11944,6 @@ func (m *awsRestjson1_deserializeOpUpdateFindingsFilte
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12306,10 +12083,8 @@ func (m *awsRestjson1_deserializeOpUpdateMacieSession)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12321,7 +12096,6 @@ func (m *awsRestjson1_deserializeOpUpdateMacieSession)
 	output := &UpdateMacieSessionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12412,10 +12186,8 @@ func (m *awsRestjson1_deserializeOpUpdateMemberSession
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12427,7 +12199,6 @@ func (m *awsRestjson1_deserializeOpUpdateMemberSession
 	output := &UpdateMemberSessionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12518,10 +12289,8 @@ func (m *awsRestjson1_deserializeOpUpdateOrganizationC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12533,7 +12302,6 @@ func (m *awsRestjson1_deserializeOpUpdateOrganizationC
 	output := &UpdateOrganizationConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12624,10 +12392,8 @@ func (m *awsRestjson1_deserializeOpUpdateResourceProfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12639,7 +12405,6 @@ func (m *awsRestjson1_deserializeOpUpdateResourceProfi
 	output := &UpdateResourceProfileOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12727,10 +12492,8 @@ func (m *awsRestjson1_deserializeOpUpdateResourceProfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12742,7 +12505,6 @@ func (m *awsRestjson1_deserializeOpUpdateResourceProfi
 	output := &UpdateResourceProfileDetectionsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12830,10 +12592,8 @@ func (m *awsRestjson1_deserializeOpUpdateRevealConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12873,7 +12633,6 @@ func (m *awsRestjson1_deserializeOpUpdateRevealConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12996,10 +12755,8 @@ func (m *awsRestjson1_deserializeOpUpdateSensitivityIn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13011,7 +12768,6 @@ func (m *awsRestjson1_deserializeOpUpdateSensitivityIn
 	output := &UpdateSensitivityInspectionTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
