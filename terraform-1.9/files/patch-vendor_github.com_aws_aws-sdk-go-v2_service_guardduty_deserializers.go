--- vendor/github.com/aws/aws-sdk-go-v2/service/guardduty/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/guardduty/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpAcceptAdministrator
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -61,7 +58,6 @@ func (m *awsRestjson1_deserializeOpAcceptAdministrator
 	output := &AcceptAdministratorInvitationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -137,10 +133,8 @@ func (m *awsRestjson1_deserializeOpAcceptInvitation) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -152,7 +146,6 @@ func (m *awsRestjson1_deserializeOpAcceptInvitation) H
 	output := &AcceptInvitationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -228,10 +221,8 @@ func (m *awsRestjson1_deserializeOpArchiveFindings) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -243,7 +234,6 @@ func (m *awsRestjson1_deserializeOpArchiveFindings) Ha
 	output := &ArchiveFindingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -319,10 +309,8 @@ func (m *awsRestjson1_deserializeOpCreateDetector) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -362,7 +350,6 @@ func (m *awsRestjson1_deserializeOpCreateDetector) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -483,10 +470,8 @@ func (m *awsRestjson1_deserializeOpCreateFilter) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -526,7 +511,6 @@ func (m *awsRestjson1_deserializeOpCreateFilter) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -642,10 +626,8 @@ func (m *awsRestjson1_deserializeOpCreateIPSet) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -685,7 +667,6 @@ func (m *awsRestjson1_deserializeOpCreateIPSet) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -804,10 +785,8 @@ func (m *awsRestjson1_deserializeOpCreateMalwareProtec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -847,7 +826,6 @@ func (m *awsRestjson1_deserializeOpCreateMalwareProtec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -969,10 +947,8 @@ func (m *awsRestjson1_deserializeOpCreateMembers) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1012,7 +988,6 @@ func (m *awsRestjson1_deserializeOpCreateMembers) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1124,10 +1099,8 @@ func (m *awsRestjson1_deserializeOpCreatePublishingDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1167,7 +1140,6 @@ func (m *awsRestjson1_deserializeOpCreatePublishingDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1283,10 +1255,8 @@ func (m *awsRestjson1_deserializeOpCreateSampleFinding
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1298,7 +1268,6 @@ func (m *awsRestjson1_deserializeOpCreateSampleFinding
 	output := &CreateSampleFindingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1374,10 +1343,8 @@ func (m *awsRestjson1_deserializeOpCreateThreatIntelSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1417,7 +1384,6 @@ func (m *awsRestjson1_deserializeOpCreateThreatIntelSe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1536,10 +1502,8 @@ func (m *awsRestjson1_deserializeOpDeclineInvitations)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1579,7 +1543,6 @@ func (m *awsRestjson1_deserializeOpDeclineInvitations)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1691,10 +1654,8 @@ func (m *awsRestjson1_deserializeOpDeleteDetector) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1706,7 +1667,6 @@ func (m *awsRestjson1_deserializeOpDeleteDetector) Han
 	output := &DeleteDetectorOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1782,10 +1742,8 @@ func (m *awsRestjson1_deserializeOpDeleteFilter) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1797,7 +1755,6 @@ func (m *awsRestjson1_deserializeOpDeleteFilter) Handl
 	output := &DeleteFilterOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1873,10 +1830,8 @@ func (m *awsRestjson1_deserializeOpDeleteInvitations) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1916,7 +1871,6 @@ func (m *awsRestjson1_deserializeOpDeleteInvitations) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2028,10 +1982,8 @@ func (m *awsRestjson1_deserializeOpDeleteIPSet) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2043,7 +1995,6 @@ func (m *awsRestjson1_deserializeOpDeleteIPSet) Handle
 	output := &DeleteIPSetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2119,10 +2070,8 @@ func (m *awsRestjson1_deserializeOpDeleteMalwareProtec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2140,7 +2089,6 @@ func (m *awsRestjson1_deserializeOpDeleteMalwareProtec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2222,10 +2170,8 @@ func (m *awsRestjson1_deserializeOpDeleteMembers) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2265,7 +2211,6 @@ func (m *awsRestjson1_deserializeOpDeleteMembers) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2377,10 +2322,8 @@ func (m *awsRestjson1_deserializeOpDeletePublishingDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2392,7 +2335,6 @@ func (m *awsRestjson1_deserializeOpDeletePublishingDes
 	output := &DeletePublishingDestinationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2468,10 +2410,8 @@ func (m *awsRestjson1_deserializeOpDeleteThreatIntelSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2483,7 +2423,6 @@ func (m *awsRestjson1_deserializeOpDeleteThreatIntelSe
 	output := &DeleteThreatIntelSetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2559,10 +2498,8 @@ func (m *awsRestjson1_deserializeOpDescribeMalwareScan
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2602,7 +2539,6 @@ func (m *awsRestjson1_deserializeOpDescribeMalwareScan
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2723,10 +2659,8 @@ func (m *awsRestjson1_deserializeOpDescribeOrganizatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2766,7 +2700,6 @@ func (m *awsRestjson1_deserializeOpDescribeOrganizatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2919,10 +2852,8 @@ func (m *awsRestjson1_deserializeOpDescribePublishingD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2962,7 +2893,6 @@ func (m *awsRestjson1_deserializeOpDescribePublishingD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3114,10 +3044,8 @@ func (m *awsRestjson1_deserializeOpDisableOrganization
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3129,7 +3057,6 @@ func (m *awsRestjson1_deserializeOpDisableOrganization
 	output := &DisableOrganizationAdminAccountOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3205,10 +3132,8 @@ func (m *awsRestjson1_deserializeOpDisassociateFromAdm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3220,7 +3145,6 @@ func (m *awsRestjson1_deserializeOpDisassociateFromAdm
 	output := &DisassociateFromAdministratorAccountOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3296,10 +3220,8 @@ func (m *awsRestjson1_deserializeOpDisassociateFromMas
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3311,7 +3233,6 @@ func (m *awsRestjson1_deserializeOpDisassociateFromMas
 	output := &DisassociateFromMasterAccountOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3387,10 +3308,8 @@ func (m *awsRestjson1_deserializeOpDisassociateMembers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3430,7 +3349,6 @@ func (m *awsRestjson1_deserializeOpDisassociateMembers
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3542,10 +3460,8 @@ func (m *awsRestjson1_deserializeOpEnableOrganizationA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3557,7 +3473,6 @@ func (m *awsRestjson1_deserializeOpEnableOrganizationA
 	output := &EnableOrganizationAdminAccountOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3633,10 +3548,8 @@ func (m *awsRestjson1_deserializeOpGetAdministratorAcc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3676,7 +3589,6 @@ func (m *awsRestjson1_deserializeOpGetAdministratorAcc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3788,10 +3700,8 @@ func (m *awsRestjson1_deserializeOpGetCoverageStatisti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3831,7 +3741,6 @@ func (m *awsRestjson1_deserializeOpGetCoverageStatisti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3943,10 +3852,8 @@ func (m *awsRestjson1_deserializeOpGetDetector) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3986,7 +3893,6 @@ func (m *awsRestjson1_deserializeOpGetDetector) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4153,10 +4059,8 @@ func (m *awsRestjson1_deserializeOpGetFilter) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4196,7 +4100,6 @@ func (m *awsRestjson1_deserializeOpGetFilter) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4353,10 +4256,8 @@ func (m *awsRestjson1_deserializeOpGetFindings) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4396,7 +4297,6 @@ func (m *awsRestjson1_deserializeOpGetFindings) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4508,10 +4408,8 @@ func (m *awsRestjson1_deserializeOpGetFindingsStatisti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4551,7 +4449,6 @@ func (m *awsRestjson1_deserializeOpGetFindingsStatisti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4672,10 +4569,8 @@ func (m *awsRestjson1_deserializeOpGetInvitationsCount
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4715,7 +4610,6 @@ func (m *awsRestjson1_deserializeOpGetInvitationsCount
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4835,10 +4729,8 @@ func (m *awsRestjson1_deserializeOpGetIPSet) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4878,7 +4770,6 @@ func (m *awsRestjson1_deserializeOpGetIPSet) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5035,10 +4926,8 @@ func (m *awsRestjson1_deserializeOpGetMalwareProtectio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5078,7 +4967,6 @@ func (m *awsRestjson1_deserializeOpGetMalwareProtectio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5254,10 +5142,8 @@ func (m *awsRestjson1_deserializeOpGetMalwareScanSetti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5297,7 +5183,6 @@ func (m *awsRestjson1_deserializeOpGetMalwareScanSetti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5418,10 +5303,8 @@ func (m *awsRestjson1_deserializeOpGetMasterAccount) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5461,7 +5344,6 @@ func (m *awsRestjson1_deserializeOpGetMasterAccount) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5573,10 +5455,8 @@ func (m *awsRestjson1_deserializeOpGetMemberDetectors)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5616,7 +5496,6 @@ func (m *awsRestjson1_deserializeOpGetMemberDetectors)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5733,10 +5612,8 @@ func (m *awsRestjson1_deserializeOpGetMembers) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5776,7 +5653,6 @@ func (m *awsRestjson1_deserializeOpGetMembers) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5893,10 +5769,8 @@ func (m *awsRestjson1_deserializeOpGetOrganizationStat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5936,7 +5810,6 @@ func (m *awsRestjson1_deserializeOpGetOrganizationStat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6048,10 +5921,8 @@ func (m *awsRestjson1_deserializeOpGetRemainingFreeTri
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6091,7 +5962,6 @@ func (m *awsRestjson1_deserializeOpGetRemainingFreeTri
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6208,10 +6078,8 @@ func (m *awsRestjson1_deserializeOpGetThreatIntelSet) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6251,7 +6119,6 @@ func (m *awsRestjson1_deserializeOpGetThreatIntelSet) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6408,10 +6275,8 @@ func (m *awsRestjson1_deserializeOpGetUsageStatistics)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6451,7 +6316,6 @@ func (m *awsRestjson1_deserializeOpGetUsageStatistics)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6572,10 +6436,8 @@ func (m *awsRestjson1_deserializeOpInviteMembers) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6615,7 +6477,6 @@ func (m *awsRestjson1_deserializeOpInviteMembers) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6727,10 +6588,8 @@ func (m *awsRestjson1_deserializeOpListCoverage) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6770,7 +6629,6 @@ func (m *awsRestjson1_deserializeOpListCoverage) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6891,10 +6749,8 @@ func (m *awsRestjson1_deserializeOpListDetectors) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6934,7 +6790,6 @@ func (m *awsRestjson1_deserializeOpListDetectors) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7055,10 +6910,8 @@ func (m *awsRestjson1_deserializeOpListFilters) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7098,7 +6951,6 @@ func (m *awsRestjson1_deserializeOpListFilters) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7219,10 +7071,8 @@ func (m *awsRestjson1_deserializeOpListFindings) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7262,7 +7112,6 @@ func (m *awsRestjson1_deserializeOpListFindings) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7383,10 +7232,8 @@ func (m *awsRestjson1_deserializeOpListInvitations) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7426,7 +7273,6 @@ func (m *awsRestjson1_deserializeOpListInvitations) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7547,10 +7393,8 @@ func (m *awsRestjson1_deserializeOpListIPSets) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7590,7 +7434,6 @@ func (m *awsRestjson1_deserializeOpListIPSets) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7711,10 +7554,8 @@ func (m *awsRestjson1_deserializeOpListMalwareProtecti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7754,7 +7595,6 @@ func (m *awsRestjson1_deserializeOpListMalwareProtecti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7878,10 +7718,8 @@ func (m *awsRestjson1_deserializeOpListMembers) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7921,7 +7759,6 @@ func (m *awsRestjson1_deserializeOpListMembers) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8042,10 +7879,8 @@ func (m *awsRestjson1_deserializeOpListOrganizationAdm
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8085,7 +7920,6 @@ func (m *awsRestjson1_deserializeOpListOrganizationAdm
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8206,10 +8040,8 @@ func (m *awsRestjson1_deserializeOpListPublishingDesti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8249,7 +8081,6 @@ func (m *awsRestjson1_deserializeOpListPublishingDesti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8370,10 +8201,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8413,7 +8242,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8528,10 +8356,8 @@ func (m *awsRestjson1_deserializeOpListThreatIntelSets
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8571,7 +8397,6 @@ func (m *awsRestjson1_deserializeOpListThreatIntelSets
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8692,10 +8517,8 @@ func (m *awsRestjson1_deserializeOpStartMalwareScan) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8735,7 +8558,6 @@ func (m *awsRestjson1_deserializeOpStartMalwareScan) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8854,10 +8676,8 @@ func (m *awsRestjson1_deserializeOpStartMonitoringMemb
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8897,7 +8717,6 @@ func (m *awsRestjson1_deserializeOpStartMonitoringMemb
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9009,10 +8828,8 @@ func (m *awsRestjson1_deserializeOpStopMonitoringMembe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9052,7 +8869,6 @@ func (m *awsRestjson1_deserializeOpStopMonitoringMembe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9164,10 +8980,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9179,7 +8993,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9258,10 +9071,8 @@ func (m *awsRestjson1_deserializeOpUnarchiveFindings) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9273,7 +9084,6 @@ func (m *awsRestjson1_deserializeOpUnarchiveFindings) 
 	output := &UnarchiveFindingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9349,10 +9159,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9364,7 +9172,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9443,10 +9250,8 @@ func (m *awsRestjson1_deserializeOpUpdateDetector) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9458,7 +9263,6 @@ func (m *awsRestjson1_deserializeOpUpdateDetector) Han
 	output := &UpdateDetectorOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9534,10 +9338,8 @@ func (m *awsRestjson1_deserializeOpUpdateFilter) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9577,7 +9379,6 @@ func (m *awsRestjson1_deserializeOpUpdateFilter) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9693,10 +9494,8 @@ func (m *awsRestjson1_deserializeOpUpdateFindingsFeedb
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9708,7 +9507,6 @@ func (m *awsRestjson1_deserializeOpUpdateFindingsFeedb
 	output := &UpdateFindingsFeedbackOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9784,10 +9582,8 @@ func (m *awsRestjson1_deserializeOpUpdateIPSet) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9799,7 +9595,6 @@ func (m *awsRestjson1_deserializeOpUpdateIPSet) Handle
 	output := &UpdateIPSetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9878,10 +9673,8 @@ func (m *awsRestjson1_deserializeOpUpdateMalwareProtec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9899,7 +9692,6 @@ func (m *awsRestjson1_deserializeOpUpdateMalwareProtec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9981,10 +9773,8 @@ func (m *awsRestjson1_deserializeOpUpdateMalwareScanSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9996,7 +9786,6 @@ func (m *awsRestjson1_deserializeOpUpdateMalwareScanSe
 	output := &UpdateMalwareScanSettingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10072,10 +9861,8 @@ func (m *awsRestjson1_deserializeOpUpdateMemberDetecto
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10115,7 +9902,6 @@ func (m *awsRestjson1_deserializeOpUpdateMemberDetecto
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10227,10 +10013,8 @@ func (m *awsRestjson1_deserializeOpUpdateOrganizationC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10242,7 +10026,6 @@ func (m *awsRestjson1_deserializeOpUpdateOrganizationC
 	output := &UpdateOrganizationConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10318,10 +10101,8 @@ func (m *awsRestjson1_deserializeOpUpdatePublishingDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10333,7 +10114,6 @@ func (m *awsRestjson1_deserializeOpUpdatePublishingDes
 	output := &UpdatePublishingDestinationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10409,10 +10189,8 @@ func (m *awsRestjson1_deserializeOpUpdateThreatIntelSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10424,7 +10202,6 @@ func (m *awsRestjson1_deserializeOpUpdateThreatIntelSe
 	output := &UpdateThreatIntelSetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
