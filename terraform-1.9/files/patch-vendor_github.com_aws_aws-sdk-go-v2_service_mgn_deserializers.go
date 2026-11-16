--- vendor/github.com/aws/aws-sdk-go-v2/service/mgn/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/mgn/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -47,10 +46,8 @@ func (m *awsRestjson1_deserializeOpArchiveApplication)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -90,7 +87,6 @@ func (m *awsRestjson1_deserializeOpArchiveApplication)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -285,10 +281,8 @@ func (m *awsRestjson1_deserializeOpArchiveWave) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -328,7 +322,6 @@ func (m *awsRestjson1_deserializeOpArchiveWave) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -514,10 +507,8 @@ func (m *awsRestjson1_deserializeOpAssociateApplicatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -529,7 +520,6 @@ func (m *awsRestjson1_deserializeOpAssociateApplicatio
 	output := &AssociateApplicationsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -611,10 +601,8 @@ func (m *awsRestjson1_deserializeOpAssociateSourceServ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -626,7 +614,6 @@ func (m *awsRestjson1_deserializeOpAssociateSourceServ
 	output := &AssociateSourceServersOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -708,10 +695,8 @@ func (m *awsRestjson1_deserializeOpChangeServerLifeCyc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -751,7 +736,6 @@ func (m *awsRestjson1_deserializeOpChangeServerLifeCyc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -966,10 +950,8 @@ func (m *awsRestjson1_deserializeOpCreateApplication) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1009,7 +991,6 @@ func (m *awsRestjson1_deserializeOpCreateApplication) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1201,10 +1182,8 @@ func (m *awsRestjson1_deserializeOpCreateConnector) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1244,7 +1223,6 @@ func (m *awsRestjson1_deserializeOpCreateConnector) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1397,10 +1375,8 @@ func (m *awsRestjson1_deserializeOpCreateLaunchConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1440,7 +1416,6 @@ func (m *awsRestjson1_deserializeOpCreateLaunchConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1687,10 +1662,8 @@ func (m *awsRestjson1_deserializeOpCreateReplicationCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1730,7 +1703,6 @@ func (m *awsRestjson1_deserializeOpCreateReplicationCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1976,10 +1948,8 @@ func (m *awsRestjson1_deserializeOpCreateWave) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2019,7 +1989,6 @@ func (m *awsRestjson1_deserializeOpCreateWave) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2202,10 +2171,8 @@ func (m *awsRestjson1_deserializeOpDeleteApplication) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2217,7 +2184,6 @@ func (m *awsRestjson1_deserializeOpDeleteApplication) 
 	output := &DeleteApplicationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2296,10 +2262,8 @@ func (m *awsRestjson1_deserializeOpDeleteConnector) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2317,7 +2281,6 @@ func (m *awsRestjson1_deserializeOpDeleteConnector) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2396,10 +2359,8 @@ func (m *awsRestjson1_deserializeOpDeleteJob) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2411,7 +2372,6 @@ func (m *awsRestjson1_deserializeOpDeleteJob) HandleDe
 	output := &DeleteJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2490,10 +2450,8 @@ func (m *awsRestjson1_deserializeOpDeleteLaunchConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2505,7 +2463,6 @@ func (m *awsRestjson1_deserializeOpDeleteLaunchConfigu
 	output := &DeleteLaunchConfigurationTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2584,10 +2541,8 @@ func (m *awsRestjson1_deserializeOpDeleteReplicationCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2599,7 +2554,6 @@ func (m *awsRestjson1_deserializeOpDeleteReplicationCo
 	output := &DeleteReplicationConfigurationTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2678,10 +2632,8 @@ func (m *awsRestjson1_deserializeOpDeleteSourceServer)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2693,7 +2645,6 @@ func (m *awsRestjson1_deserializeOpDeleteSourceServer)
 	output := &DeleteSourceServerOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2772,10 +2723,8 @@ func (m *awsRestjson1_deserializeOpDeleteVcenterClient
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2793,7 +2742,6 @@ func (m *awsRestjson1_deserializeOpDeleteVcenterClient
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2872,10 +2820,8 @@ func (m *awsRestjson1_deserializeOpDeleteWave) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2887,7 +2833,6 @@ func (m *awsRestjson1_deserializeOpDeleteWave) HandleD
 	output := &DeleteWaveOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2966,10 +2911,8 @@ func (m *awsRestjson1_deserializeOpDescribeJobLogItems
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3009,7 +2952,6 @@ func (m *awsRestjson1_deserializeOpDescribeJobLogItems
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3130,10 +3072,8 @@ func (m *awsRestjson1_deserializeOpDescribeJobs) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3173,7 +3113,6 @@ func (m *awsRestjson1_deserializeOpDescribeJobs) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3294,10 +3233,8 @@ func (m *awsRestjson1_deserializeOpDescribeLaunchConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3337,7 +3274,6 @@ func (m *awsRestjson1_deserializeOpDescribeLaunchConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3461,10 +3397,8 @@ func (m *awsRestjson1_deserializeOpDescribeReplication
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3504,7 +3438,6 @@ func (m *awsRestjson1_deserializeOpDescribeReplication
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3628,10 +3561,8 @@ func (m *awsRestjson1_deserializeOpDescribeSourceServe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3671,7 +3602,6 @@ func (m *awsRestjson1_deserializeOpDescribeSourceServe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3792,10 +3722,8 @@ func (m *awsRestjson1_deserializeOpDescribeVcenterClie
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3835,7 +3763,6 @@ func (m *awsRestjson1_deserializeOpDescribeVcenterClie
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3959,10 +3886,8 @@ func (m *awsRestjson1_deserializeOpDisassociateApplica
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3974,7 +3899,6 @@ func (m *awsRestjson1_deserializeOpDisassociateApplica
 	output := &DisassociateApplicationsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4053,10 +3977,8 @@ func (m *awsRestjson1_deserializeOpDisassociateSourceS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4068,7 +3990,6 @@ func (m *awsRestjson1_deserializeOpDisassociateSourceS
 	output := &DisassociateSourceServersOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4147,10 +4068,8 @@ func (m *awsRestjson1_deserializeOpDisconnectFromServi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4190,7 +4109,6 @@ func (m *awsRestjson1_deserializeOpDisconnectFromServi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4402,10 +4320,8 @@ func (m *awsRestjson1_deserializeOpFinalizeCutover) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4445,7 +4361,6 @@ func (m *awsRestjson1_deserializeOpFinalizeCutover) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4660,10 +4575,8 @@ func (m *awsRestjson1_deserializeOpGetLaunchConfigurat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4703,7 +4616,6 @@ func (m *awsRestjson1_deserializeOpGetLaunchConfigurat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4910,10 +4822,8 @@ func (m *awsRestjson1_deserializeOpGetReplicationConfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4953,7 +4863,6 @@ func (m *awsRestjson1_deserializeOpGetReplicationConfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5196,10 +5105,8 @@ func (m *awsRestjson1_deserializeOpInitializeService) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5211,7 +5118,6 @@ func (m *awsRestjson1_deserializeOpInitializeService) 
 	output := &InitializeServiceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5287,10 +5193,8 @@ func (m *awsRestjson1_deserializeOpListApplications) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5330,7 +5234,6 @@ func (m *awsRestjson1_deserializeOpListApplications) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5448,10 +5351,8 @@ func (m *awsRestjson1_deserializeOpListConnectors) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5491,7 +5392,6 @@ func (m *awsRestjson1_deserializeOpListConnectors) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5612,10 +5512,8 @@ func (m *awsRestjson1_deserializeOpListExportErrors) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5655,7 +5553,6 @@ func (m *awsRestjson1_deserializeOpListExportErrors) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5776,10 +5673,8 @@ func (m *awsRestjson1_deserializeOpListExports) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5819,7 +5714,6 @@ func (m *awsRestjson1_deserializeOpListExports) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5937,10 +5831,8 @@ func (m *awsRestjson1_deserializeOpListImportErrors) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5980,7 +5872,6 @@ func (m *awsRestjson1_deserializeOpListImportErrors) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6101,10 +5992,8 @@ func (m *awsRestjson1_deserializeOpListImports) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6144,7 +6033,6 @@ func (m *awsRestjson1_deserializeOpListImports) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6265,10 +6153,8 @@ func (m *awsRestjson1_deserializeOpListManagedAccounts
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6308,7 +6194,6 @@ func (m *awsRestjson1_deserializeOpListManagedAccounts
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6429,10 +6314,8 @@ func (m *awsRestjson1_deserializeOpListSourceServerAct
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6472,7 +6355,6 @@ func (m *awsRestjson1_deserializeOpListSourceServerAct
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6593,10 +6475,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6636,7 +6516,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6757,10 +6636,8 @@ func (m *awsRestjson1_deserializeOpListTemplateActions
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6800,7 +6677,6 @@ func (m *awsRestjson1_deserializeOpListTemplateActions
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6921,10 +6797,8 @@ func (m *awsRestjson1_deserializeOpListWaves) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6964,7 +6838,6 @@ func (m *awsRestjson1_deserializeOpListWaves) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7082,10 +6955,8 @@ func (m *awsRestjson1_deserializeOpMarkAsArchived) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7125,7 +6996,6 @@ func (m *awsRestjson1_deserializeOpMarkAsArchived) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7337,10 +7207,8 @@ func (m *awsRestjson1_deserializeOpPauseReplication) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7380,7 +7248,6 @@ func (m *awsRestjson1_deserializeOpPauseReplication) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7598,10 +7465,8 @@ func (m *awsRestjson1_deserializeOpPutSourceServerActi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7641,7 +7506,6 @@ func (m *awsRestjson1_deserializeOpPutSourceServerActi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7862,10 +7726,8 @@ func (m *awsRestjson1_deserializeOpPutTemplateAction) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7905,7 +7767,6 @@ func (m *awsRestjson1_deserializeOpPutTemplateAction) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8135,10 +7996,8 @@ func (m *awsRestjson1_deserializeOpRemoveSourceServerA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8150,7 +8009,6 @@ func (m *awsRestjson1_deserializeOpRemoveSourceServerA
 	output := &RemoveSourceServerActionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8229,10 +8087,8 @@ func (m *awsRestjson1_deserializeOpRemoveTemplateActio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8244,7 +8100,6 @@ func (m *awsRestjson1_deserializeOpRemoveTemplateActio
 	output := &RemoveTemplateActionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8323,10 +8178,8 @@ func (m *awsRestjson1_deserializeOpResumeReplication) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8366,7 +8219,6 @@ func (m *awsRestjson1_deserializeOpResumeReplication) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8584,10 +8436,8 @@ func (m *awsRestjson1_deserializeOpRetryDataReplicatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8627,7 +8477,6 @@ func (m *awsRestjson1_deserializeOpRetryDataReplicatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8839,10 +8688,8 @@ func (m *awsRestjson1_deserializeOpStartCutover) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8882,7 +8729,6 @@ func (m *awsRestjson1_deserializeOpStartCutover) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8997,10 +8843,8 @@ func (m *awsRestjson1_deserializeOpStartExport) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9040,7 +8884,6 @@ func (m *awsRestjson1_deserializeOpStartExport) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9155,10 +8998,8 @@ func (m *awsRestjson1_deserializeOpStartImport) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9198,7 +9039,6 @@ func (m *awsRestjson1_deserializeOpStartImport) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9319,10 +9159,8 @@ func (m *awsRestjson1_deserializeOpStartReplication) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9362,7 +9200,6 @@ func (m *awsRestjson1_deserializeOpStartReplication) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9580,10 +9417,8 @@ func (m *awsRestjson1_deserializeOpStartTest) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9623,7 +9458,6 @@ func (m *awsRestjson1_deserializeOpStartTest) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9738,10 +9572,8 @@ func (m *awsRestjson1_deserializeOpStopReplication) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9781,7 +9613,6 @@ func (m *awsRestjson1_deserializeOpStopReplication) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9999,10 +9830,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10020,7 +9849,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10105,10 +9933,8 @@ func (m *awsRestjson1_deserializeOpTerminateTargetInst
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10148,7 +9974,6 @@ func (m *awsRestjson1_deserializeOpTerminateTargetInst
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10263,10 +10088,8 @@ func (m *awsRestjson1_deserializeOpUnarchiveApplicatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10306,7 +10129,6 @@ func (m *awsRestjson1_deserializeOpUnarchiveApplicatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10498,10 +10320,8 @@ func (m *awsRestjson1_deserializeOpUnarchiveWave) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10541,7 +10361,6 @@ func (m *awsRestjson1_deserializeOpUnarchiveWave) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10724,10 +10543,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10745,7 +10562,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10830,10 +10646,8 @@ func (m *awsRestjson1_deserializeOpUpdateApplication) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10873,7 +10687,6 @@ func (m *awsRestjson1_deserializeOpUpdateApplication) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11065,10 +10878,8 @@ func (m *awsRestjson1_deserializeOpUpdateConnector) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11108,7 +10919,6 @@ func (m *awsRestjson1_deserializeOpUpdateConnector) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11264,10 +11074,8 @@ func (m *awsRestjson1_deserializeOpUpdateLaunchConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11307,7 +11115,6 @@ func (m *awsRestjson1_deserializeOpUpdateLaunchConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11520,10 +11327,8 @@ func (m *awsRestjson1_deserializeOpUpdateLaunchConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11563,7 +11368,6 @@ func (m *awsRestjson1_deserializeOpUpdateLaunchConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11813,10 +11617,8 @@ func (m *awsRestjson1_deserializeOpUpdateReplicationCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11856,7 +11658,6 @@ func (m *awsRestjson1_deserializeOpUpdateReplicationCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12108,10 +11909,8 @@ func (m *awsRestjson1_deserializeOpUpdateReplicationCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12151,7 +11950,6 @@ func (m *awsRestjson1_deserializeOpUpdateReplicationCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12400,10 +12198,8 @@ func (m *awsRestjson1_deserializeOpUpdateSourceServer)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12443,7 +12239,6 @@ func (m *awsRestjson1_deserializeOpUpdateSourceServer)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12655,10 +12450,8 @@ func (m *awsRestjson1_deserializeOpUpdateSourceServerR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12698,7 +12491,6 @@ func (m *awsRestjson1_deserializeOpUpdateSourceServerR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12913,10 +12705,8 @@ func (m *awsRestjson1_deserializeOpUpdateWave) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12956,7 +12746,6 @@ func (m *awsRestjson1_deserializeOpUpdateWave) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
