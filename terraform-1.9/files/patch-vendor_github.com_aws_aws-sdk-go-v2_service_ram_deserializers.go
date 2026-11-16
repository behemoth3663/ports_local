--- vendor/github.com/aws/aws-sdk-go-v2/service/ram/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/ram/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpAcceptResourceShare
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpAcceptResourceShare
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -232,10 +228,8 @@ func (m *awsRestjson1_deserializeOpAssociateResourceSh
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -275,7 +269,6 @@ func (m *awsRestjson1_deserializeOpAssociateResourceSh
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -423,10 +416,8 @@ func (m *awsRestjson1_deserializeOpAssociateResourceSh
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -466,7 +457,6 @@ func (m *awsRestjson1_deserializeOpAssociateResourceSh
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -606,10 +596,8 @@ func (m *awsRestjson1_deserializeOpCreatePermission) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -649,7 +637,6 @@ func (m *awsRestjson1_deserializeOpCreatePermission) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -794,10 +781,8 @@ func (m *awsRestjson1_deserializeOpCreatePermissionVer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -837,7 +822,6 @@ func (m *awsRestjson1_deserializeOpCreatePermissionVer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -982,10 +966,8 @@ func (m *awsRestjson1_deserializeOpCreateResourceShare
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1025,7 +1007,6 @@ func (m *awsRestjson1_deserializeOpCreateResourceShare
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1176,10 +1157,8 @@ func (m *awsRestjson1_deserializeOpDeletePermission) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1219,7 +1198,6 @@ func (m *awsRestjson1_deserializeOpDeletePermission) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1368,10 +1346,8 @@ func (m *awsRestjson1_deserializeOpDeletePermissionVer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1411,7 +1387,6 @@ func (m *awsRestjson1_deserializeOpDeletePermissionVer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1563,10 +1538,8 @@ func (m *awsRestjson1_deserializeOpDeleteResourceShare
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1606,7 +1579,6 @@ func (m *awsRestjson1_deserializeOpDeleteResourceShare
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1752,10 +1724,8 @@ func (m *awsRestjson1_deserializeOpDisassociateResourc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1795,7 +1765,6 @@ func (m *awsRestjson1_deserializeOpDisassociateResourc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1940,10 +1909,8 @@ func (m *awsRestjson1_deserializeOpDisassociateResourc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1983,7 +1950,6 @@ func (m *awsRestjson1_deserializeOpDisassociateResourc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2126,10 +2092,8 @@ func (m *awsRestjson1_deserializeOpEnableSharingWithAw
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2169,7 +2133,6 @@ func (m *awsRestjson1_deserializeOpEnableSharingWithAw
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2288,10 +2251,8 @@ func (m *awsRestjson1_deserializeOpGetPermission) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2331,7 +2292,6 @@ func (m *awsRestjson1_deserializeOpGetPermission) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2455,10 +2415,8 @@ func (m *awsRestjson1_deserializeOpGetResourcePolicies
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2498,7 +2456,6 @@ func (m *awsRestjson1_deserializeOpGetResourcePolicies
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2631,10 +2588,8 @@ func (m *awsRestjson1_deserializeOpGetResourceShareAss
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2674,7 +2629,6 @@ func (m *awsRestjson1_deserializeOpGetResourceShareAss
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2810,10 +2764,8 @@ func (m *awsRestjson1_deserializeOpGetResourceShareInv
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2853,7 +2805,6 @@ func (m *awsRestjson1_deserializeOpGetResourceShareInv
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2992,10 +2943,8 @@ func (m *awsRestjson1_deserializeOpGetResourceShares) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3035,7 +2984,6 @@ func (m *awsRestjson1_deserializeOpGetResourceShares) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3168,10 +3116,8 @@ func (m *awsRestjson1_deserializeOpListPendingInvitati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3211,7 +3157,6 @@ func (m *awsRestjson1_deserializeOpListPendingInvitati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3353,10 +3298,8 @@ func (m *awsRestjson1_deserializeOpListPermissionAssoc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3396,7 +3339,6 @@ func (m *awsRestjson1_deserializeOpListPermissionAssoc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3526,10 +3468,8 @@ func (m *awsRestjson1_deserializeOpListPermissions) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3569,7 +3509,6 @@ func (m *awsRestjson1_deserializeOpListPermissions) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3699,10 +3638,8 @@ func (m *awsRestjson1_deserializeOpListPermissionVersi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3742,7 +3679,6 @@ func (m *awsRestjson1_deserializeOpListPermissionVersi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3878,10 +3814,8 @@ func (m *awsRestjson1_deserializeOpListPrincipals) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3921,7 +3855,6 @@ func (m *awsRestjson1_deserializeOpListPrincipals) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4054,10 +3987,8 @@ func (m *awsRestjson1_deserializeOpListReplacePermissi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4097,7 +4028,6 @@ func (m *awsRestjson1_deserializeOpListReplacePermissi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4224,10 +4154,8 @@ func (m *awsRestjson1_deserializeOpListResources) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4267,7 +4195,6 @@ func (m *awsRestjson1_deserializeOpListResources) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4403,10 +4330,8 @@ func (m *awsRestjson1_deserializeOpListResourceSharePe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4446,7 +4371,6 @@ func (m *awsRestjson1_deserializeOpListResourceSharePe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4582,10 +4506,8 @@ func (m *awsRestjson1_deserializeOpListResourceTypes) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4625,7 +4547,6 @@ func (m *awsRestjson1_deserializeOpListResourceTypes) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4752,10 +4673,8 @@ func (m *awsRestjson1_deserializeOpPromotePermissionCr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4795,7 +4714,6 @@ func (m *awsRestjson1_deserializeOpPromotePermissionCr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4931,10 +4849,8 @@ func (m *awsRestjson1_deserializeOpPromoteResourceShar
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4974,7 +4890,6 @@ func (m *awsRestjson1_deserializeOpPromoteResourceShar
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5114,10 +5029,8 @@ func (m *awsRestjson1_deserializeOpRejectResourceShare
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5157,7 +5070,6 @@ func (m *awsRestjson1_deserializeOpRejectResourceShare
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5302,10 +5214,8 @@ func (m *awsRestjson1_deserializeOpReplacePermissionAs
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5345,7 +5255,6 @@ func (m *awsRestjson1_deserializeOpReplacePermissionAs
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5484,10 +5393,8 @@ func (m *awsRestjson1_deserializeOpSetDefaultPermissio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5527,7 +5434,6 @@ func (m *awsRestjson1_deserializeOpSetDefaultPermissio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5667,10 +5573,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5682,7 +5586,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5776,10 +5679,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5791,7 +5692,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5876,10 +5776,8 @@ func (m *awsRestjson1_deserializeOpUpdateResourceShare
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5919,7 +5817,6 @@ func (m *awsRestjson1_deserializeOpUpdateResourceShare
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
