--- vendor/github.com/aws/aws-sdk-go-v2/service/codeartifact/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/codeartifact/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAssociateExternalCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpAssociateExternalCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -215,10 +211,8 @@ func (m *awsRestjson1_deserializeOpCopyPackageVersions
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -258,7 +252,6 @@ func (m *awsRestjson1_deserializeOpCopyPackageVersions
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -390,10 +383,8 @@ func (m *awsRestjson1_deserializeOpCreateDomain) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -433,7 +424,6 @@ func (m *awsRestjson1_deserializeOpCreateDomain) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -560,10 +550,8 @@ func (m *awsRestjson1_deserializeOpCreatePackageGroup)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -603,7 +591,6 @@ func (m *awsRestjson1_deserializeOpCreatePackageGroup)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -730,10 +717,8 @@ func (m *awsRestjson1_deserializeOpCreateRepository) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -773,7 +758,6 @@ func (m *awsRestjson1_deserializeOpCreateRepository) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -900,10 +884,8 @@ func (m *awsRestjson1_deserializeOpDeleteDomain) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -943,7 +925,6 @@ func (m *awsRestjson1_deserializeOpDeleteDomain) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1064,10 +1045,8 @@ func (m *awsRestjson1_deserializeOpDeleteDomainPermiss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1107,7 +1086,6 @@ func (m *awsRestjson1_deserializeOpDeleteDomainPermiss
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1231,10 +1209,8 @@ func (m *awsRestjson1_deserializeOpDeletePackage) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1274,7 +1250,6 @@ func (m *awsRestjson1_deserializeOpDeletePackage) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1398,10 +1373,8 @@ func (m *awsRestjson1_deserializeOpDeletePackageGroup)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1441,7 +1414,6 @@ func (m *awsRestjson1_deserializeOpDeletePackageGroup)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1568,10 +1540,8 @@ func (m *awsRestjson1_deserializeOpDeletePackageVersio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1611,7 +1581,6 @@ func (m *awsRestjson1_deserializeOpDeletePackageVersio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1740,10 +1709,8 @@ func (m *awsRestjson1_deserializeOpDeleteRepository) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1783,7 +1750,6 @@ func (m *awsRestjson1_deserializeOpDeleteRepository) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1907,10 +1873,8 @@ func (m *awsRestjson1_deserializeOpDeleteRepositoryPer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1950,7 +1914,6 @@ func (m *awsRestjson1_deserializeOpDeleteRepositoryPer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2074,10 +2037,8 @@ func (m *awsRestjson1_deserializeOpDescribeDomain) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2117,7 +2078,6 @@ func (m *awsRestjson1_deserializeOpDescribeDomain) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2238,10 +2198,8 @@ func (m *awsRestjson1_deserializeOpDescribePackage) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2281,7 +2239,6 @@ func (m *awsRestjson1_deserializeOpDescribePackage) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2402,10 +2359,8 @@ func (m *awsRestjson1_deserializeOpDescribePackageGrou
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2445,7 +2400,6 @@ func (m *awsRestjson1_deserializeOpDescribePackageGrou
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2566,10 +2520,8 @@ func (m *awsRestjson1_deserializeOpDescribePackageVers
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2609,7 +2561,6 @@ func (m *awsRestjson1_deserializeOpDescribePackageVers
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2733,10 +2684,8 @@ func (m *awsRestjson1_deserializeOpDescribeRepository)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2776,7 +2725,6 @@ func (m *awsRestjson1_deserializeOpDescribeRepository)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2897,10 +2845,8 @@ func (m *awsRestjson1_deserializeOpDisassociateExterna
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2940,7 +2886,6 @@ func (m *awsRestjson1_deserializeOpDisassociateExterna
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3067,10 +3012,8 @@ func (m *awsRestjson1_deserializeOpDisposePackageVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3110,7 +3053,6 @@ func (m *awsRestjson1_deserializeOpDisposePackageVersi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3239,10 +3181,8 @@ func (m *awsRestjson1_deserializeOpGetAssociatedPackag
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3282,7 +3222,6 @@ func (m *awsRestjson1_deserializeOpGetAssociatedPackag
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3409,10 +3348,8 @@ func (m *awsRestjson1_deserializeOpGetAuthorizationTok
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3452,7 +3389,6 @@ func (m *awsRestjson1_deserializeOpGetAuthorizationTok
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3593,10 +3529,8 @@ func (m *awsRestjson1_deserializeOpGetDomainPermission
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3636,7 +3570,6 @@ func (m *awsRestjson1_deserializeOpGetDomainPermission
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3757,10 +3690,8 @@ func (m *awsRestjson1_deserializeOpGetPackageVersionAs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3782,7 +3713,6 @@ func (m *awsRestjson1_deserializeOpGetPackageVersionAs
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3901,10 +3831,8 @@ func (m *awsRestjson1_deserializeOpGetPackageVersionRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3944,7 +3872,6 @@ func (m *awsRestjson1_deserializeOpGetPackageVersionRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4114,10 +4041,8 @@ func (m *awsRestjson1_deserializeOpGetRepositoryEndpoi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4157,7 +4082,6 @@ func (m *awsRestjson1_deserializeOpGetRepositoryEndpoi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4282,10 +4206,8 @@ func (m *awsRestjson1_deserializeOpGetRepositoryPermis
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4325,7 +4247,6 @@ func (m *awsRestjson1_deserializeOpGetRepositoryPermis
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4446,10 +4367,8 @@ func (m *awsRestjson1_deserializeOpListAllowedReposito
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4489,7 +4408,6 @@ func (m *awsRestjson1_deserializeOpListAllowedReposito
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4622,10 +4540,8 @@ func (m *awsRestjson1_deserializeOpListAssociatedPacka
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4665,7 +4581,6 @@ func (m *awsRestjson1_deserializeOpListAssociatedPacka
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4792,10 +4707,8 @@ func (m *awsRestjson1_deserializeOpListDomains) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4835,7 +4748,6 @@ func (m *awsRestjson1_deserializeOpListDomains) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4962,10 +4874,8 @@ func (m *awsRestjson1_deserializeOpListPackageGroups) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5005,7 +4915,6 @@ func (m *awsRestjson1_deserializeOpListPackageGroups) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5135,10 +5044,8 @@ func (m *awsRestjson1_deserializeOpListPackages) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5178,7 +5085,6 @@ func (m *awsRestjson1_deserializeOpListPackages) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5308,10 +5214,8 @@ func (m *awsRestjson1_deserializeOpListPackageVersionA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5351,7 +5255,6 @@ func (m *awsRestjson1_deserializeOpListPackageVersionA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5526,10 +5429,8 @@ func (m *awsRestjson1_deserializeOpListPackageVersionD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5569,7 +5470,6 @@ func (m *awsRestjson1_deserializeOpListPackageVersionD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5744,10 +5644,8 @@ func (m *awsRestjson1_deserializeOpListPackageVersions
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5787,7 +5685,6 @@ func (m *awsRestjson1_deserializeOpListPackageVersions
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5953,10 +5850,8 @@ func (m *awsRestjson1_deserializeOpListRepositories) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5996,7 +5891,6 @@ func (m *awsRestjson1_deserializeOpListRepositories) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6123,10 +6017,8 @@ func (m *awsRestjson1_deserializeOpListRepositoriesInD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6166,7 +6058,6 @@ func (m *awsRestjson1_deserializeOpListRepositoriesInD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6296,10 +6187,8 @@ func (m *awsRestjson1_deserializeOpListSubPackageGroup
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6339,7 +6228,6 @@ func (m *awsRestjson1_deserializeOpListSubPackageGroup
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6469,10 +6357,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6512,7 +6398,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6630,10 +6515,8 @@ func (m *awsRestjson1_deserializeOpPublishPackageVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6673,7 +6556,6 @@ func (m *awsRestjson1_deserializeOpPublishPackageVersi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6854,10 +6736,8 @@ func (m *awsRestjson1_deserializeOpPutDomainPermission
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6897,7 +6777,6 @@ func (m *awsRestjson1_deserializeOpPutDomainPermission
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7024,10 +6903,8 @@ func (m *awsRestjson1_deserializeOpPutPackageOriginCon
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7067,7 +6944,6 @@ func (m *awsRestjson1_deserializeOpPutPackageOriginCon
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7188,10 +7064,8 @@ func (m *awsRestjson1_deserializeOpPutRepositoryPermis
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7231,7 +7105,6 @@ func (m *awsRestjson1_deserializeOpPutRepositoryPermis
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7358,10 +7231,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7373,7 +7244,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7458,10 +7328,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7473,7 +7341,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7555,10 +7422,8 @@ func (m *awsRestjson1_deserializeOpUpdatePackageGroup)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7598,7 +7463,6 @@ func (m *awsRestjson1_deserializeOpUpdatePackageGroup)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7722,10 +7586,8 @@ func (m *awsRestjson1_deserializeOpUpdatePackageGroupO
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7765,7 +7627,6 @@ func (m *awsRestjson1_deserializeOpUpdatePackageGroupO
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7894,10 +7755,8 @@ func (m *awsRestjson1_deserializeOpUpdatePackageVersio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7937,7 +7796,6 @@ func (m *awsRestjson1_deserializeOpUpdatePackageVersio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8066,10 +7924,8 @@ func (m *awsRestjson1_deserializeOpUpdateRepository) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8109,7 +7965,6 @@ func (m *awsRestjson1_deserializeOpUpdateRepository) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
