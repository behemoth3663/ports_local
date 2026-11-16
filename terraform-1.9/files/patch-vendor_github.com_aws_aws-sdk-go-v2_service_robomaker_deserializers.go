--- vendor/github.com/aws/aws-sdk-go-v2/service/robomaker/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/robomaker/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpBatchDeleteWorlds) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpBatchDeleteWorlds) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -203,10 +199,8 @@ func (m *awsRestjson1_deserializeOpBatchDescribeSimula
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -246,7 +240,6 @@ func (m *awsRestjson1_deserializeOpBatchDescribeSimula
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -369,10 +362,8 @@ func (m *awsRestjson1_deserializeOpCancelDeploymentJob
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -384,7 +375,6 @@ func (m *awsRestjson1_deserializeOpCancelDeploymentJob
 	output := &CancelDeploymentJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -466,10 +456,8 @@ func (m *awsRestjson1_deserializeOpCancelSimulationJob
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -481,7 +469,6 @@ func (m *awsRestjson1_deserializeOpCancelSimulationJob
 	output := &CancelSimulationJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -563,10 +550,8 @@ func (m *awsRestjson1_deserializeOpCancelSimulationJob
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -578,7 +563,6 @@ func (m *awsRestjson1_deserializeOpCancelSimulationJob
 	output := &CancelSimulationJobBatchOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -660,10 +644,8 @@ func (m *awsRestjson1_deserializeOpCancelWorldExportJo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -675,7 +657,6 @@ func (m *awsRestjson1_deserializeOpCancelWorldExportJo
 	output := &CancelWorldExportJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -757,10 +738,8 @@ func (m *awsRestjson1_deserializeOpCancelWorldGenerati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -772,7 +751,6 @@ func (m *awsRestjson1_deserializeOpCancelWorldGenerati
 	output := &CancelWorldGenerationJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -854,10 +832,8 @@ func (m *awsRestjson1_deserializeOpCreateDeploymentJob
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -897,7 +873,6 @@ func (m *awsRestjson1_deserializeOpCreateDeploymentJob
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1095,10 +1070,8 @@ func (m *awsRestjson1_deserializeOpCreateFleet) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1138,7 +1111,6 @@ func (m *awsRestjson1_deserializeOpCreateFleet) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1290,10 +1262,8 @@ func (m *awsRestjson1_deserializeOpCreateRobot) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1333,7 +1303,6 @@ func (m *awsRestjson1_deserializeOpCreateRobot) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1506,10 +1475,8 @@ func (m *awsRestjson1_deserializeOpCreateRobotApplicat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1549,7 +1516,6 @@ func (m *awsRestjson1_deserializeOpCreateRobotApplicat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1740,10 +1706,8 @@ func (m *awsRestjson1_deserializeOpCreateRobotApplicat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1783,7 +1747,6 @@ func (m *awsRestjson1_deserializeOpCreateRobotApplicat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1966,10 +1929,8 @@ func (m *awsRestjson1_deserializeOpCreateSimulationApp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2009,7 +1970,6 @@ func (m *awsRestjson1_deserializeOpCreateSimulationApp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2210,10 +2170,8 @@ func (m *awsRestjson1_deserializeOpCreateSimulationApp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2253,7 +2211,6 @@ func (m *awsRestjson1_deserializeOpCreateSimulationApp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2446,10 +2403,8 @@ func (m *awsRestjson1_deserializeOpCreateSimulationJob
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2489,7 +2444,6 @@ func (m *awsRestjson1_deserializeOpCreateSimulationJob
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2763,10 +2717,8 @@ func (m *awsRestjson1_deserializeOpCreateWorldExportJo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2806,7 +2758,6 @@ func (m *awsRestjson1_deserializeOpCreateWorldExportJo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2996,10 +2947,8 @@ func (m *awsRestjson1_deserializeOpCreateWorldGenerati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3039,7 +2988,6 @@ func (m *awsRestjson1_deserializeOpCreateWorldGenerati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3237,10 +3185,8 @@ func (m *awsRestjson1_deserializeOpCreateWorldTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3280,7 +3226,6 @@ func (m *awsRestjson1_deserializeOpCreateWorldTemplate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3447,10 +3392,8 @@ func (m *awsRestjson1_deserializeOpDeleteFleet) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3462,7 +3405,6 @@ func (m *awsRestjson1_deserializeOpDeleteFleet) Handle
 	output := &DeleteFleetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3541,10 +3483,8 @@ func (m *awsRestjson1_deserializeOpDeleteRobot) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3556,7 +3496,6 @@ func (m *awsRestjson1_deserializeOpDeleteRobot) Handle
 	output := &DeleteRobotOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3635,10 +3574,8 @@ func (m *awsRestjson1_deserializeOpDeleteRobotApplicat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3650,7 +3587,6 @@ func (m *awsRestjson1_deserializeOpDeleteRobotApplicat
 	output := &DeleteRobotApplicationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3729,10 +3665,8 @@ func (m *awsRestjson1_deserializeOpDeleteSimulationApp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3744,7 +3678,6 @@ func (m *awsRestjson1_deserializeOpDeleteSimulationApp
 	output := &DeleteSimulationApplicationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3823,10 +3756,8 @@ func (m *awsRestjson1_deserializeOpDeleteWorldTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3838,7 +3769,6 @@ func (m *awsRestjson1_deserializeOpDeleteWorldTemplate
 	output := &DeleteWorldTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3920,10 +3850,8 @@ func (m *awsRestjson1_deserializeOpDeregisterRobot) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3963,7 +3891,6 @@ func (m *awsRestjson1_deserializeOpDeregisterRobot) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4094,10 +4021,8 @@ func (m *awsRestjson1_deserializeOpDescribeDeploymentJ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4137,7 +4062,6 @@ func (m *awsRestjson1_deserializeOpDescribeDeploymentJ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4331,10 +4255,8 @@ func (m *awsRestjson1_deserializeOpDescribeFleet) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4374,7 +4296,6 @@ func (m *awsRestjson1_deserializeOpDescribeFleet) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4565,10 +4486,8 @@ func (m *awsRestjson1_deserializeOpDescribeRobot) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4608,7 +4527,6 @@ func (m *awsRestjson1_deserializeOpDescribeRobot) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4821,10 +4739,8 @@ func (m *awsRestjson1_deserializeOpDescribeRobotApplic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4864,7 +4780,6 @@ func (m *awsRestjson1_deserializeOpDescribeRobotApplic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5058,10 +4973,8 @@ func (m *awsRestjson1_deserializeOpDescribeSimulationA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5101,7 +5014,6 @@ func (m *awsRestjson1_deserializeOpDescribeSimulationA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5305,10 +5217,8 @@ func (m *awsRestjson1_deserializeOpDescribeSimulationJ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5348,7 +5258,6 @@ func (m *awsRestjson1_deserializeOpDescribeSimulationJ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5636,10 +5545,8 @@ func (m *awsRestjson1_deserializeOpDescribeSimulationJ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5679,7 +5586,6 @@ func (m *awsRestjson1_deserializeOpDescribeSimulationJ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5891,10 +5797,8 @@ func (m *awsRestjson1_deserializeOpDescribeWorld) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5934,7 +5838,6 @@ func (m *awsRestjson1_deserializeOpDescribeWorld) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6104,10 +6007,8 @@ func (m *awsRestjson1_deserializeOpDescribeWorldExport
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6147,7 +6048,6 @@ func (m *awsRestjson1_deserializeOpDescribeWorldExport
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6345,10 +6245,8 @@ func (m *awsRestjson1_deserializeOpDescribeWorldGenera
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6388,7 +6286,6 @@ func (m *awsRestjson1_deserializeOpDescribeWorldGenera
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6591,10 +6488,8 @@ func (m *awsRestjson1_deserializeOpDescribeWorldTempla
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6634,7 +6529,6 @@ func (m *awsRestjson1_deserializeOpDescribeWorldTempla
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6820,10 +6714,8 @@ func (m *awsRestjson1_deserializeOpGetWorldTemplateBod
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6863,7 +6755,6 @@ func (m *awsRestjson1_deserializeOpGetWorldTemplateBod
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6985,10 +6876,8 @@ func (m *awsRestjson1_deserializeOpListDeploymentJobs)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7028,7 +6917,6 @@ func (m *awsRestjson1_deserializeOpListDeploymentJobs)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7155,10 +7043,8 @@ func (m *awsRestjson1_deserializeOpListFleets) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7198,7 +7084,6 @@ func (m *awsRestjson1_deserializeOpListFleets) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7325,10 +7210,8 @@ func (m *awsRestjson1_deserializeOpListRobotApplicatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7368,7 +7251,6 @@ func (m *awsRestjson1_deserializeOpListRobotApplicatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7492,10 +7374,8 @@ func (m *awsRestjson1_deserializeOpListRobots) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7535,7 +7415,6 @@ func (m *awsRestjson1_deserializeOpListRobots) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7662,10 +7541,8 @@ func (m *awsRestjson1_deserializeOpListSimulationAppli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7705,7 +7582,6 @@ func (m *awsRestjson1_deserializeOpListSimulationAppli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7829,10 +7705,8 @@ func (m *awsRestjson1_deserializeOpListSimulationJobBa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7872,7 +7746,6 @@ func (m *awsRestjson1_deserializeOpListSimulationJobBa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7993,10 +7866,8 @@ func (m *awsRestjson1_deserializeOpListSimulationJobs)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8036,7 +7907,6 @@ func (m *awsRestjson1_deserializeOpListSimulationJobs)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8160,10 +8030,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8203,7 +8071,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8321,10 +8188,8 @@ func (m *awsRestjson1_deserializeOpListWorldExportJobs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8364,7 +8229,6 @@ func (m *awsRestjson1_deserializeOpListWorldExportJobs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8488,10 +8352,8 @@ func (m *awsRestjson1_deserializeOpListWorldGeneration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8531,7 +8393,6 @@ func (m *awsRestjson1_deserializeOpListWorldGeneration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8655,10 +8516,8 @@ func (m *awsRestjson1_deserializeOpListWorlds) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8698,7 +8557,6 @@ func (m *awsRestjson1_deserializeOpListWorlds) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8822,10 +8680,8 @@ func (m *awsRestjson1_deserializeOpListWorldTemplates)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8865,7 +8721,6 @@ func (m *awsRestjson1_deserializeOpListWorldTemplates)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8989,10 +8844,8 @@ func (m *awsRestjson1_deserializeOpRegisterRobot) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9032,7 +8885,6 @@ func (m *awsRestjson1_deserializeOpRegisterRobot) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9166,10 +9018,8 @@ func (m *awsRestjson1_deserializeOpRestartSimulationJo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9181,7 +9031,6 @@ func (m *awsRestjson1_deserializeOpRestartSimulationJo
 	output := &RestartSimulationJobOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9266,10 +9115,8 @@ func (m *awsRestjson1_deserializeOpStartSimulationJobB
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9309,7 +9156,6 @@ func (m *awsRestjson1_deserializeOpStartSimulationJobB
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9511,10 +9357,8 @@ func (m *awsRestjson1_deserializeOpSyncDeploymentJob) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9554,7 +9398,6 @@ func (m *awsRestjson1_deserializeOpSyncDeploymentJob) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9747,10 +9590,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9762,7 +9603,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9844,10 +9684,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9859,7 +9697,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9941,10 +9778,8 @@ func (m *awsRestjson1_deserializeOpUpdateRobotApplicat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9984,7 +9819,6 @@ func (m *awsRestjson1_deserializeOpUpdateRobotApplicat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10167,10 +10001,8 @@ func (m *awsRestjson1_deserializeOpUpdateSimulationApp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10210,7 +10042,6 @@ func (m *awsRestjson1_deserializeOpUpdateSimulationApp
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10403,10 +10234,8 @@ func (m *awsRestjson1_deserializeOpUpdateWorldTemplate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10446,7 +10275,6 @@ func (m *awsRestjson1_deserializeOpUpdateWorldTemplate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
