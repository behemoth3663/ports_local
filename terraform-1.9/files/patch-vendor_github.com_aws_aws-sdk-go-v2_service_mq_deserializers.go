--- vendor/github.com/aws/aws-sdk-go-v2/service/mq/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/mq/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCreateBroker) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCreateBroker) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -222,10 +218,8 @@ func (m *awsRestjson1_deserializeOpCreateConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -265,7 +259,6 @@ func (m *awsRestjson1_deserializeOpCreateConfiguration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -432,10 +425,8 @@ func (m *awsRestjson1_deserializeOpCreateTags) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -453,7 +444,6 @@ func (m *awsRestjson1_deserializeOpCreateTags) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -535,10 +525,8 @@ func (m *awsRestjson1_deserializeOpCreateUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -550,7 +538,6 @@ func (m *awsRestjson1_deserializeOpCreateUser) HandleD
 	output := &CreateUserOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -635,10 +622,8 @@ func (m *awsRestjson1_deserializeOpDeleteBroker) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -678,7 +663,6 @@ func (m *awsRestjson1_deserializeOpDeleteBroker) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -800,10 +784,8 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -843,7 +825,6 @@ func (m *awsRestjson1_deserializeOpDeleteConfiguration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -968,10 +949,8 @@ func (m *awsRestjson1_deserializeOpDeleteTags) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -989,7 +968,6 @@ func (m *awsRestjson1_deserializeOpDeleteTags) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1071,10 +1049,8 @@ func (m *awsRestjson1_deserializeOpDeleteUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1086,7 +1062,6 @@ func (m *awsRestjson1_deserializeOpDeleteUser) HandleD
 	output := &DeleteUserOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1168,10 +1143,8 @@ func (m *awsRestjson1_deserializeOpDescribeBroker) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1211,7 +1184,6 @@ func (m *awsRestjson1_deserializeOpDescribeBroker) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1565,10 +1537,8 @@ func (m *awsRestjson1_deserializeOpDescribeBrokerEngin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1608,7 +1578,6 @@ func (m *awsRestjson1_deserializeOpDescribeBrokerEngin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1745,10 +1714,8 @@ func (m *awsRestjson1_deserializeOpDescribeBrokerInsta
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1788,7 +1755,6 @@ func (m *awsRestjson1_deserializeOpDescribeBrokerInsta
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1925,10 +1891,8 @@ func (m *awsRestjson1_deserializeOpDescribeConfigurati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1968,7 +1932,6 @@ func (m *awsRestjson1_deserializeOpDescribeConfigurati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2167,10 +2130,8 @@ func (m *awsRestjson1_deserializeOpDescribeConfigurati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2210,7 +2171,6 @@ func (m *awsRestjson1_deserializeOpDescribeConfigurati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2363,10 +2323,8 @@ func (m *awsRestjson1_deserializeOpDescribeUser) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2406,7 +2364,6 @@ func (m *awsRestjson1_deserializeOpDescribeUser) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2565,10 +2522,8 @@ func (m *awsRestjson1_deserializeOpListBrokers) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2608,7 +2563,6 @@ func (m *awsRestjson1_deserializeOpListBrokers) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2732,10 +2686,8 @@ func (m *awsRestjson1_deserializeOpListConfigurationRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2775,7 +2727,6 @@ func (m *awsRestjson1_deserializeOpListConfigurationRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2924,10 +2875,8 @@ func (m *awsRestjson1_deserializeOpListConfigurations)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2967,7 +2916,6 @@ func (m *awsRestjson1_deserializeOpListConfigurations)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3104,10 +3052,8 @@ func (m *awsRestjson1_deserializeOpListTags) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3147,7 +3093,6 @@ func (m *awsRestjson1_deserializeOpListTags) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3265,10 +3210,8 @@ func (m *awsRestjson1_deserializeOpListUsers) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3308,7 +3251,6 @@ func (m *awsRestjson1_deserializeOpListUsers) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3457,10 +3399,8 @@ func (m *awsRestjson1_deserializeOpPromote) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3500,7 +3440,6 @@ func (m *awsRestjson1_deserializeOpPromote) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3622,10 +3561,8 @@ func (m *awsRestjson1_deserializeOpRebootBroker) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3637,7 +3574,6 @@ func (m *awsRestjson1_deserializeOpRebootBroker) Handl
 	output := &RebootBrokerOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3719,10 +3655,8 @@ func (m *awsRestjson1_deserializeOpUpdateBroker) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3762,7 +3696,6 @@ func (m *awsRestjson1_deserializeOpUpdateBroker) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3976,10 +3909,8 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguration
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4019,7 +3950,6 @@ func (m *awsRestjson1_deserializeOpUpdateConfiguration
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4185,10 +4115,8 @@ func (m *awsRestjson1_deserializeOpUpdateUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4200,7 +4128,6 @@ func (m *awsRestjson1_deserializeOpUpdateUser) HandleD
 	output := &UpdateUserOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
