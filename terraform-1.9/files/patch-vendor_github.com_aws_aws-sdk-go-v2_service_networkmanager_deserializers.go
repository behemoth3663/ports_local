--- vendor/github.com/aws/aws-sdk-go-v2/service/networkmanager/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/networkmanager/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAcceptAttachment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpAcceptAttachment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -212,10 +208,8 @@ func (m *awsRestjson1_deserializeOpAssociateConnectPee
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -255,7 +249,6 @@ func (m *awsRestjson1_deserializeOpAssociateConnectPee
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -382,10 +375,8 @@ func (m *awsRestjson1_deserializeOpAssociateCustomerGa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -425,7 +416,6 @@ func (m *awsRestjson1_deserializeOpAssociateCustomerGa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -552,10 +542,8 @@ func (m *awsRestjson1_deserializeOpAssociateLink) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -595,7 +583,6 @@ func (m *awsRestjson1_deserializeOpAssociateLink) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -722,10 +709,8 @@ func (m *awsRestjson1_deserializeOpAssociateTransitGat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -765,7 +750,6 @@ func (m *awsRestjson1_deserializeOpAssociateTransitGat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -892,10 +876,8 @@ func (m *awsRestjson1_deserializeOpCreateConnectAttach
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -935,7 +917,6 @@ func (m *awsRestjson1_deserializeOpCreateConnectAttach
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1059,10 +1040,8 @@ func (m *awsRestjson1_deserializeOpCreateConnection) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1102,7 +1081,6 @@ func (m *awsRestjson1_deserializeOpCreateConnection) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1226,10 +1204,8 @@ func (m *awsRestjson1_deserializeOpCreateConnectPeer) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1269,7 +1245,6 @@ func (m *awsRestjson1_deserializeOpCreateConnectPeer) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1393,10 +1368,8 @@ func (m *awsRestjson1_deserializeOpCreateCoreNetwork) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1436,7 +1409,6 @@ func (m *awsRestjson1_deserializeOpCreateCoreNetwork) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1563,10 +1535,8 @@ func (m *awsRestjson1_deserializeOpCreateDevice) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1606,7 +1576,6 @@ func (m *awsRestjson1_deserializeOpCreateDevice) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1733,10 +1702,8 @@ func (m *awsRestjson1_deserializeOpCreateDirectConnect
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1776,7 +1743,6 @@ func (m *awsRestjson1_deserializeOpCreateDirectConnect
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1900,10 +1866,8 @@ func (m *awsRestjson1_deserializeOpCreateGlobalNetwork
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1943,7 +1907,6 @@ func (m *awsRestjson1_deserializeOpCreateGlobalNetwork
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2067,10 +2030,8 @@ func (m *awsRestjson1_deserializeOpCreateLink) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2110,7 +2071,6 @@ func (m *awsRestjson1_deserializeOpCreateLink) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2237,10 +2197,8 @@ func (m *awsRestjson1_deserializeOpCreateSite) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2280,7 +2238,6 @@ func (m *awsRestjson1_deserializeOpCreateSite) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2407,10 +2364,8 @@ func (m *awsRestjson1_deserializeOpCreateSiteToSiteVpn
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2450,7 +2405,6 @@ func (m *awsRestjson1_deserializeOpCreateSiteToSiteVpn
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2574,10 +2528,8 @@ func (m *awsRestjson1_deserializeOpCreateTransitGatewa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2617,7 +2569,6 @@ func (m *awsRestjson1_deserializeOpCreateTransitGatewa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2741,10 +2692,8 @@ func (m *awsRestjson1_deserializeOpCreateTransitGatewa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2784,7 +2733,6 @@ func (m *awsRestjson1_deserializeOpCreateTransitGatewa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2908,10 +2856,8 @@ func (m *awsRestjson1_deserializeOpCreateVpcAttachment
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2951,7 +2897,6 @@ func (m *awsRestjson1_deserializeOpCreateVpcAttachment
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3075,10 +3020,8 @@ func (m *awsRestjson1_deserializeOpDeleteAttachment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3118,7 +3061,6 @@ func (m *awsRestjson1_deserializeOpDeleteAttachment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3242,10 +3184,8 @@ func (m *awsRestjson1_deserializeOpDeleteConnection) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3285,7 +3225,6 @@ func (m *awsRestjson1_deserializeOpDeleteConnection) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3409,10 +3348,8 @@ func (m *awsRestjson1_deserializeOpDeleteConnectPeer) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3452,7 +3389,6 @@ func (m *awsRestjson1_deserializeOpDeleteConnectPeer) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3576,10 +3512,8 @@ func (m *awsRestjson1_deserializeOpDeleteCoreNetwork) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3619,7 +3553,6 @@ func (m *awsRestjson1_deserializeOpDeleteCoreNetwork) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3743,10 +3676,8 @@ func (m *awsRestjson1_deserializeOpDeleteCoreNetworkPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3786,7 +3717,6 @@ func (m *awsRestjson1_deserializeOpDeleteCoreNetworkPo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3910,10 +3840,8 @@ func (m *awsRestjson1_deserializeOpDeleteDevice) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3953,7 +3881,6 @@ func (m *awsRestjson1_deserializeOpDeleteDevice) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4077,10 +4004,8 @@ func (m *awsRestjson1_deserializeOpDeleteGlobalNetwork
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4120,7 +4045,6 @@ func (m *awsRestjson1_deserializeOpDeleteGlobalNetwork
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4244,10 +4168,8 @@ func (m *awsRestjson1_deserializeOpDeleteLink) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4287,7 +4209,6 @@ func (m *awsRestjson1_deserializeOpDeleteLink) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4411,10 +4332,8 @@ func (m *awsRestjson1_deserializeOpDeletePeering) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4454,7 +4373,6 @@ func (m *awsRestjson1_deserializeOpDeletePeering) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4578,10 +4496,8 @@ func (m *awsRestjson1_deserializeOpDeleteResourcePolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4593,7 +4509,6 @@ func (m *awsRestjson1_deserializeOpDeleteResourcePolic
 	output := &DeleteResourcePolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4678,10 +4593,8 @@ func (m *awsRestjson1_deserializeOpDeleteSite) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4721,7 +4634,6 @@ func (m *awsRestjson1_deserializeOpDeleteSite) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4845,10 +4757,8 @@ func (m *awsRestjson1_deserializeOpDeregisterTransitGa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4888,7 +4798,6 @@ func (m *awsRestjson1_deserializeOpDeregisterTransitGa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5012,10 +4921,8 @@ func (m *awsRestjson1_deserializeOpDescribeGlobalNetwo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5055,7 +4962,6 @@ func (m *awsRestjson1_deserializeOpDescribeGlobalNetwo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5185,10 +5091,8 @@ func (m *awsRestjson1_deserializeOpDisassociateConnect
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5228,7 +5132,6 @@ func (m *awsRestjson1_deserializeOpDisassociateConnect
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5352,10 +5255,8 @@ func (m *awsRestjson1_deserializeOpDisassociateCustome
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5395,7 +5296,6 @@ func (m *awsRestjson1_deserializeOpDisassociateCustome
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5519,10 +5419,8 @@ func (m *awsRestjson1_deserializeOpDisassociateLink) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5562,7 +5460,6 @@ func (m *awsRestjson1_deserializeOpDisassociateLink) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5686,10 +5583,8 @@ func (m *awsRestjson1_deserializeOpDisassociateTransit
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5729,7 +5624,6 @@ func (m *awsRestjson1_deserializeOpDisassociateTransit
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5853,10 +5747,8 @@ func (m *awsRestjson1_deserializeOpExecuteCoreNetworkC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5868,7 +5760,6 @@ func (m *awsRestjson1_deserializeOpExecuteCoreNetworkC
 	output := &ExecuteCoreNetworkChangeSetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5956,10 +5847,8 @@ func (m *awsRestjson1_deserializeOpGetConnectAttachmen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5999,7 +5888,6 @@ func (m *awsRestjson1_deserializeOpGetConnectAttachmen
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6120,10 +6008,8 @@ func (m *awsRestjson1_deserializeOpGetConnections) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6163,7 +6049,6 @@ func (m *awsRestjson1_deserializeOpGetConnections) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6293,10 +6178,8 @@ func (m *awsRestjson1_deserializeOpGetConnectPeer) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6336,7 +6219,6 @@ func (m *awsRestjson1_deserializeOpGetConnectPeer) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6457,10 +6339,8 @@ func (m *awsRestjson1_deserializeOpGetConnectPeerAssoc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6500,7 +6380,6 @@ func (m *awsRestjson1_deserializeOpGetConnectPeerAssoc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6633,10 +6512,8 @@ func (m *awsRestjson1_deserializeOpGetCoreNetwork) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6676,7 +6553,6 @@ func (m *awsRestjson1_deserializeOpGetCoreNetwork) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6797,10 +6673,8 @@ func (m *awsRestjson1_deserializeOpGetCoreNetworkChang
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6840,7 +6714,6 @@ func (m *awsRestjson1_deserializeOpGetCoreNetworkChang
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6970,10 +6843,8 @@ func (m *awsRestjson1_deserializeOpGetCoreNetworkChang
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7013,7 +6884,6 @@ func (m *awsRestjson1_deserializeOpGetCoreNetworkChang
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7143,10 +7013,8 @@ func (m *awsRestjson1_deserializeOpGetCoreNetworkPolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7186,7 +7054,6 @@ func (m *awsRestjson1_deserializeOpGetCoreNetworkPolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7307,10 +7174,8 @@ func (m *awsRestjson1_deserializeOpGetCustomerGatewayA
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7350,7 +7215,6 @@ func (m *awsRestjson1_deserializeOpGetCustomerGatewayA
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7483,10 +7347,8 @@ func (m *awsRestjson1_deserializeOpGetDevices) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7526,7 +7388,6 @@ func (m *awsRestjson1_deserializeOpGetDevices) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7656,10 +7517,8 @@ func (m *awsRestjson1_deserializeOpGetDirectConnectGat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7699,7 +7558,6 @@ func (m *awsRestjson1_deserializeOpGetDirectConnectGat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7820,10 +7678,8 @@ func (m *awsRestjson1_deserializeOpGetLinkAssociations
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7863,7 +7719,6 @@ func (m *awsRestjson1_deserializeOpGetLinkAssociations
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7993,10 +7848,8 @@ func (m *awsRestjson1_deserializeOpGetLinks) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8036,7 +7889,6 @@ func (m *awsRestjson1_deserializeOpGetLinks) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8166,10 +8018,8 @@ func (m *awsRestjson1_deserializeOpGetNetworkResourceC
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8209,7 +8059,6 @@ func (m *awsRestjson1_deserializeOpGetNetworkResourceC
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8336,10 +8185,8 @@ func (m *awsRestjson1_deserializeOpGetNetworkResourceR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8379,7 +8226,6 @@ func (m *awsRestjson1_deserializeOpGetNetworkResourceR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8509,10 +8355,8 @@ func (m *awsRestjson1_deserializeOpGetNetworkResources
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8552,7 +8396,6 @@ func (m *awsRestjson1_deserializeOpGetNetworkResources
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8682,10 +8525,8 @@ func (m *awsRestjson1_deserializeOpGetNetworkRoutes) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8725,7 +8566,6 @@ func (m *awsRestjson1_deserializeOpGetNetworkRoutes) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8885,10 +8725,8 @@ func (m *awsRestjson1_deserializeOpGetNetworkTelemetry
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8928,7 +8766,6 @@ func (m *awsRestjson1_deserializeOpGetNetworkTelemetry
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9058,10 +8895,8 @@ func (m *awsRestjson1_deserializeOpGetResourcePolicy) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9101,7 +8936,6 @@ func (m *awsRestjson1_deserializeOpGetResourcePolicy) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9223,10 +9057,8 @@ func (m *awsRestjson1_deserializeOpGetRouteAnalysis) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9266,7 +9098,6 @@ func (m *awsRestjson1_deserializeOpGetRouteAnalysis) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9387,10 +9218,8 @@ func (m *awsRestjson1_deserializeOpGetSites) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9430,7 +9259,6 @@ func (m *awsRestjson1_deserializeOpGetSites) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9560,10 +9388,8 @@ func (m *awsRestjson1_deserializeOpGetSiteToSiteVpnAtt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9603,7 +9429,6 @@ func (m *awsRestjson1_deserializeOpGetSiteToSiteVpnAtt
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9724,10 +9549,8 @@ func (m *awsRestjson1_deserializeOpGetTransitGatewayCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9767,7 +9590,6 @@ func (m *awsRestjson1_deserializeOpGetTransitGatewayCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9900,10 +9722,8 @@ func (m *awsRestjson1_deserializeOpGetTransitGatewayPe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9943,7 +9763,6 @@ func (m *awsRestjson1_deserializeOpGetTransitGatewayPe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10064,10 +9883,8 @@ func (m *awsRestjson1_deserializeOpGetTransitGatewayRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10107,7 +9924,6 @@ func (m *awsRestjson1_deserializeOpGetTransitGatewayRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10237,10 +10053,8 @@ func (m *awsRestjson1_deserializeOpGetTransitGatewayRo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10280,7 +10094,6 @@ func (m *awsRestjson1_deserializeOpGetTransitGatewayRo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10401,10 +10214,8 @@ func (m *awsRestjson1_deserializeOpGetVpcAttachment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10444,7 +10255,6 @@ func (m *awsRestjson1_deserializeOpGetVpcAttachment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10565,10 +10375,8 @@ func (m *awsRestjson1_deserializeOpListAttachments) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10608,7 +10416,6 @@ func (m *awsRestjson1_deserializeOpListAttachments) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10735,10 +10542,8 @@ func (m *awsRestjson1_deserializeOpListConnectPeers) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10778,7 +10583,6 @@ func (m *awsRestjson1_deserializeOpListConnectPeers) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10905,10 +10709,8 @@ func (m *awsRestjson1_deserializeOpListCoreNetworkPoli
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10948,7 +10750,6 @@ func (m *awsRestjson1_deserializeOpListCoreNetworkPoli
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11078,10 +10879,8 @@ func (m *awsRestjson1_deserializeOpListCoreNetworks) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11121,7 +10920,6 @@ func (m *awsRestjson1_deserializeOpListCoreNetworks) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11248,10 +11046,8 @@ func (m *awsRestjson1_deserializeOpListOrganizationSer
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11291,7 +11087,6 @@ func (m *awsRestjson1_deserializeOpListOrganizationSer
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11406,10 +11201,8 @@ func (m *awsRestjson1_deserializeOpListPeerings) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11449,7 +11242,6 @@ func (m *awsRestjson1_deserializeOpListPeerings) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11576,10 +11368,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11619,7 +11409,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11740,10 +11529,8 @@ func (m *awsRestjson1_deserializeOpPutCoreNetworkPolic
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11783,7 +11570,6 @@ func (m *awsRestjson1_deserializeOpPutCoreNetworkPolic
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11910,10 +11696,8 @@ func (m *awsRestjson1_deserializeOpPutResourcePolicy) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11925,7 +11709,6 @@ func (m *awsRestjson1_deserializeOpPutResourcePolicy) 
 	output := &PutResourcePolicyOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12013,10 +11796,8 @@ func (m *awsRestjson1_deserializeOpRegisterTransitGate
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12056,7 +11837,6 @@ func (m *awsRestjson1_deserializeOpRegisterTransitGate
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12180,10 +11960,8 @@ func (m *awsRestjson1_deserializeOpRejectAttachment) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12223,7 +12001,6 @@ func (m *awsRestjson1_deserializeOpRejectAttachment) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12347,10 +12124,8 @@ func (m *awsRestjson1_deserializeOpRestoreCoreNetworkP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12390,7 +12165,6 @@ func (m *awsRestjson1_deserializeOpRestoreCoreNetworkP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12514,10 +12288,8 @@ func (m *awsRestjson1_deserializeOpStartOrganizationSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12557,7 +12329,6 @@ func (m *awsRestjson1_deserializeOpStartOrganizationSe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12681,10 +12452,8 @@ func (m *awsRestjson1_deserializeOpStartRouteAnalysis)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12724,7 +12493,6 @@ func (m *awsRestjson1_deserializeOpStartRouteAnalysis)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12848,10 +12616,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12863,7 +12629,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12954,10 +12719,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12969,7 +12732,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13057,10 +12819,8 @@ func (m *awsRestjson1_deserializeOpUpdateConnection) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13100,7 +12860,6 @@ func (m *awsRestjson1_deserializeOpUpdateConnection) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13224,10 +12983,8 @@ func (m *awsRestjson1_deserializeOpUpdateCoreNetwork) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13267,7 +13024,6 @@ func (m *awsRestjson1_deserializeOpUpdateCoreNetwork) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13391,10 +13147,8 @@ func (m *awsRestjson1_deserializeOpUpdateDevice) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13434,7 +13188,6 @@ func (m *awsRestjson1_deserializeOpUpdateDevice) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13558,10 +13311,8 @@ func (m *awsRestjson1_deserializeOpUpdateDirectConnect
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13601,7 +13352,6 @@ func (m *awsRestjson1_deserializeOpUpdateDirectConnect
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13725,10 +13475,8 @@ func (m *awsRestjson1_deserializeOpUpdateGlobalNetwork
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13768,7 +13516,6 @@ func (m *awsRestjson1_deserializeOpUpdateGlobalNetwork
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13892,10 +13639,8 @@ func (m *awsRestjson1_deserializeOpUpdateLink) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13935,7 +13680,6 @@ func (m *awsRestjson1_deserializeOpUpdateLink) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14062,10 +13806,8 @@ func (m *awsRestjson1_deserializeOpUpdateNetworkResour
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14105,7 +13847,6 @@ func (m *awsRestjson1_deserializeOpUpdateNetworkResour
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14238,10 +13979,8 @@ func (m *awsRestjson1_deserializeOpUpdateSite) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14281,7 +14020,6 @@ func (m *awsRestjson1_deserializeOpUpdateSite) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14405,10 +14143,8 @@ func (m *awsRestjson1_deserializeOpUpdateVpcAttachment
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14448,7 +14184,6 @@ func (m *awsRestjson1_deserializeOpUpdateVpcAttachment
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
