--- vendor/github.com/aws/aws-sdk-go-v2/service/directoryservicedata/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/directoryservicedata/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAddGroupMember) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -60,7 +57,6 @@ func (m *awsRestjson1_deserializeOpAddGroupMember) Han
 	output := &AddGroupMemberOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -151,10 +147,8 @@ func (m *awsRestjson1_deserializeOpCreateGroup) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -194,7 +188,6 @@ func (m *awsRestjson1_deserializeOpCreateGroup) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -340,10 +333,8 @@ func (m *awsRestjson1_deserializeOpCreateUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -383,7 +374,6 @@ func (m *awsRestjson1_deserializeOpCreateUser) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -529,10 +519,8 @@ func (m *awsRestjson1_deserializeOpDeleteGroup) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -544,7 +532,6 @@ func (m *awsRestjson1_deserializeOpDeleteGroup) Handle
 	output := &DeleteGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -635,10 +622,8 @@ func (m *awsRestjson1_deserializeOpDeleteUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -650,7 +635,6 @@ func (m *awsRestjson1_deserializeOpDeleteUser) HandleD
 	output := &DeleteUserOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -741,10 +725,8 @@ func (m *awsRestjson1_deserializeOpDescribeGroup) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -784,7 +766,6 @@ func (m *awsRestjson1_deserializeOpDescribeGroup) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -971,10 +952,8 @@ func (m *awsRestjson1_deserializeOpDescribeUser) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1014,7 +993,6 @@ func (m *awsRestjson1_deserializeOpDescribeUser) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1228,10 +1206,8 @@ func (m *awsRestjson1_deserializeOpDisableUser) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1243,7 +1219,6 @@ func (m *awsRestjson1_deserializeOpDisableUser) Handle
 	output := &DisableUserOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1334,10 +1309,8 @@ func (m *awsRestjson1_deserializeOpListGroupMembers) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1377,7 +1350,6 @@ func (m *awsRestjson1_deserializeOpListGroupMembers) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1537,10 +1509,8 @@ func (m *awsRestjson1_deserializeOpListGroups) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1580,7 +1550,6 @@ func (m *awsRestjson1_deserializeOpListGroups) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1728,10 +1697,8 @@ func (m *awsRestjson1_deserializeOpListGroupsForMember
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1771,7 +1738,6 @@ func (m *awsRestjson1_deserializeOpListGroupsForMember
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1931,10 +1897,8 @@ func (m *awsRestjson1_deserializeOpListUsers) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1974,7 +1938,6 @@ func (m *awsRestjson1_deserializeOpListUsers) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2122,10 +2085,8 @@ func (m *awsRestjson1_deserializeOpRemoveGroupMember) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2137,7 +2098,6 @@ func (m *awsRestjson1_deserializeOpRemoveGroupMember) 
 	output := &RemoveGroupMemberOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2228,10 +2188,8 @@ func (m *awsRestjson1_deserializeOpSearchGroups) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2271,7 +2229,6 @@ func (m *awsRestjson1_deserializeOpSearchGroups) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2419,10 +2376,8 @@ func (m *awsRestjson1_deserializeOpSearchUsers) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2462,7 +2417,6 @@ func (m *awsRestjson1_deserializeOpSearchUsers) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2610,10 +2564,8 @@ func (m *awsRestjson1_deserializeOpUpdateGroup) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2625,7 +2577,6 @@ func (m *awsRestjson1_deserializeOpUpdateGroup) Handle
 	output := &UpdateGroupOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2716,10 +2667,8 @@ func (m *awsRestjson1_deserializeOpUpdateUser) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2731,7 +2680,6 @@ func (m *awsRestjson1_deserializeOpUpdateUser) HandleD
 	output := &UpdateUserOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
