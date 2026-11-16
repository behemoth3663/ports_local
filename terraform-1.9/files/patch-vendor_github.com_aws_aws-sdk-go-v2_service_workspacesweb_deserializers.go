--- vendor/github.com/aws/aws-sdk-go-v2/service/workspacesweb/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/workspacesweb/deserializers.go
@@ -15,7 +15,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strconv"
@@ -46,10 +45,8 @@ func (m *awsRestjson1_deserializeOpAssociateBrowserSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -89,7 +86,6 @@ func (m *awsRestjson1_deserializeOpAssociateBrowserSet
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -226,10 +222,8 @@ func (m *awsRestjson1_deserializeOpAssociateDataProtec
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -269,7 +263,6 @@ func (m *awsRestjson1_deserializeOpAssociateDataProtec
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -406,10 +399,8 @@ func (m *awsRestjson1_deserializeOpAssociateIpAccessSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -449,7 +440,6 @@ func (m *awsRestjson1_deserializeOpAssociateIpAccessSe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -586,10 +576,8 @@ func (m *awsRestjson1_deserializeOpAssociateNetworkSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -629,7 +617,6 @@ func (m *awsRestjson1_deserializeOpAssociateNetworkSet
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -766,10 +753,8 @@ func (m *awsRestjson1_deserializeOpAssociateTrustStore
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -809,7 +794,6 @@ func (m *awsRestjson1_deserializeOpAssociateTrustStore
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -946,10 +930,8 @@ func (m *awsRestjson1_deserializeOpAssociateUserAccess
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -989,7 +971,6 @@ func (m *awsRestjson1_deserializeOpAssociateUserAccess
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1126,10 +1107,8 @@ func (m *awsRestjson1_deserializeOpAssociateUserSettin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1169,7 +1148,6 @@ func (m *awsRestjson1_deserializeOpAssociateUserSettin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1306,10 +1284,8 @@ func (m *awsRestjson1_deserializeOpCreateBrowserSettin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1349,7 +1325,6 @@ func (m *awsRestjson1_deserializeOpCreateBrowserSettin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1480,10 +1455,8 @@ func (m *awsRestjson1_deserializeOpCreateDataProtectio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1523,7 +1496,6 @@ func (m *awsRestjson1_deserializeOpCreateDataProtectio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1654,10 +1626,8 @@ func (m *awsRestjson1_deserializeOpCreateIdentityProvi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1697,7 +1667,6 @@ func (m *awsRestjson1_deserializeOpCreateIdentityProvi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1828,10 +1797,8 @@ func (m *awsRestjson1_deserializeOpCreateIpAccessSetti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1871,7 +1838,6 @@ func (m *awsRestjson1_deserializeOpCreateIpAccessSetti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1999,10 +1965,8 @@ func (m *awsRestjson1_deserializeOpCreateNetworkSettin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2042,7 +2006,6 @@ func (m *awsRestjson1_deserializeOpCreateNetworkSettin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2170,10 +2133,8 @@ func (m *awsRestjson1_deserializeOpCreatePortal) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2213,7 +2174,6 @@ func (m *awsRestjson1_deserializeOpCreatePortal) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2353,10 +2313,8 @@ func (m *awsRestjson1_deserializeOpCreateTrustStore) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2396,7 +2354,6 @@ func (m *awsRestjson1_deserializeOpCreateTrustStore) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2524,10 +2481,8 @@ func (m *awsRestjson1_deserializeOpCreateUserAccessLog
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2567,7 +2522,6 @@ func (m *awsRestjson1_deserializeOpCreateUserAccessLog
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2695,10 +2649,8 @@ func (m *awsRestjson1_deserializeOpCreateUserSettings)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2738,7 +2690,6 @@ func (m *awsRestjson1_deserializeOpCreateUserSettings)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2866,10 +2817,8 @@ func (m *awsRestjson1_deserializeOpDeleteBrowserSettin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2881,7 +2830,6 @@ func (m *awsRestjson1_deserializeOpDeleteBrowserSettin
 	output := &DeleteBrowserSettingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2966,10 +2914,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataProtectio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2981,7 +2927,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataProtectio
 	output := &DeleteDataProtectionSettingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3066,10 +3011,8 @@ func (m *awsRestjson1_deserializeOpDeleteIdentityProvi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3081,7 +3024,6 @@ func (m *awsRestjson1_deserializeOpDeleteIdentityProvi
 	output := &DeleteIdentityProviderOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3166,10 +3108,8 @@ func (m *awsRestjson1_deserializeOpDeleteIpAccessSetti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3181,7 +3121,6 @@ func (m *awsRestjson1_deserializeOpDeleteIpAccessSetti
 	output := &DeleteIpAccessSettingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3266,10 +3205,8 @@ func (m *awsRestjson1_deserializeOpDeleteNetworkSettin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3281,7 +3218,6 @@ func (m *awsRestjson1_deserializeOpDeleteNetworkSettin
 	output := &DeleteNetworkSettingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3366,10 +3302,8 @@ func (m *awsRestjson1_deserializeOpDeletePortal) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3381,7 +3315,6 @@ func (m *awsRestjson1_deserializeOpDeletePortal) Handl
 	output := &DeletePortalOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3466,10 +3399,8 @@ func (m *awsRestjson1_deserializeOpDeleteTrustStore) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3481,7 +3412,6 @@ func (m *awsRestjson1_deserializeOpDeleteTrustStore) H
 	output := &DeleteTrustStoreOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3566,10 +3496,8 @@ func (m *awsRestjson1_deserializeOpDeleteUserAccessLog
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3581,7 +3509,6 @@ func (m *awsRestjson1_deserializeOpDeleteUserAccessLog
 	output := &DeleteUserAccessLoggingSettingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3666,10 +3593,8 @@ func (m *awsRestjson1_deserializeOpDeleteUserSettings)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3681,7 +3606,6 @@ func (m *awsRestjson1_deserializeOpDeleteUserSettings)
 	output := &DeleteUserSettingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3766,10 +3690,8 @@ func (m *awsRestjson1_deserializeOpDisassociateBrowser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3781,7 +3703,6 @@ func (m *awsRestjson1_deserializeOpDisassociateBrowser
 	output := &DisassociateBrowserSettingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3869,10 +3790,8 @@ func (m *awsRestjson1_deserializeOpDisassociateDataPro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3884,7 +3803,6 @@ func (m *awsRestjson1_deserializeOpDisassociateDataPro
 	output := &DisassociateDataProtectionSettingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3972,10 +3890,8 @@ func (m *awsRestjson1_deserializeOpDisassociateIpAcces
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3987,7 +3903,6 @@ func (m *awsRestjson1_deserializeOpDisassociateIpAcces
 	output := &DisassociateIpAccessSettingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4075,10 +3990,8 @@ func (m *awsRestjson1_deserializeOpDisassociateNetwork
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4090,7 +4003,6 @@ func (m *awsRestjson1_deserializeOpDisassociateNetwork
 	output := &DisassociateNetworkSettingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4178,10 +4090,8 @@ func (m *awsRestjson1_deserializeOpDisassociateTrustSt
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4193,7 +4103,6 @@ func (m *awsRestjson1_deserializeOpDisassociateTrustSt
 	output := &DisassociateTrustStoreOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4281,10 +4190,8 @@ func (m *awsRestjson1_deserializeOpDisassociateUserAcc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4296,7 +4203,6 @@ func (m *awsRestjson1_deserializeOpDisassociateUserAcc
 	output := &DisassociateUserAccessLoggingSettingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4384,10 +4290,8 @@ func (m *awsRestjson1_deserializeOpDisassociateUserSet
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4399,7 +4303,6 @@ func (m *awsRestjson1_deserializeOpDisassociateUserSet
 	output := &DisassociateUserSettingsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4487,10 +4390,8 @@ func (m *awsRestjson1_deserializeOpExpireSession) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4502,7 +4403,6 @@ func (m *awsRestjson1_deserializeOpExpireSession) Hand
 	output := &ExpireSessionOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4587,10 +4487,8 @@ func (m *awsRestjson1_deserializeOpGetBrowserSettings)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4630,7 +4528,6 @@ func (m *awsRestjson1_deserializeOpGetBrowserSettings)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4751,10 +4648,8 @@ func (m *awsRestjson1_deserializeOpGetDataProtectionSe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4794,7 +4689,6 @@ func (m *awsRestjson1_deserializeOpGetDataProtectionSe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4915,10 +4809,8 @@ func (m *awsRestjson1_deserializeOpGetIdentityProvider
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4958,7 +4850,6 @@ func (m *awsRestjson1_deserializeOpGetIdentityProvider
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5079,10 +4970,8 @@ func (m *awsRestjson1_deserializeOpGetIpAccessSettings
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5122,7 +5011,6 @@ func (m *awsRestjson1_deserializeOpGetIpAccessSettings
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5243,10 +5131,8 @@ func (m *awsRestjson1_deserializeOpGetNetworkSettings)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5286,7 +5172,6 @@ func (m *awsRestjson1_deserializeOpGetNetworkSettings)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5407,10 +5292,8 @@ func (m *awsRestjson1_deserializeOpGetPortal) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5450,7 +5333,6 @@ func (m *awsRestjson1_deserializeOpGetPortal) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5571,10 +5453,8 @@ func (m *awsRestjson1_deserializeOpGetPortalServicePro
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5614,7 +5494,6 @@ func (m *awsRestjson1_deserializeOpGetPortalServicePro
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5748,10 +5627,8 @@ func (m *awsRestjson1_deserializeOpGetSession) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5791,7 +5668,6 @@ func (m *awsRestjson1_deserializeOpGetSession) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5912,10 +5788,8 @@ func (m *awsRestjson1_deserializeOpGetTrustStore) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5955,7 +5829,6 @@ func (m *awsRestjson1_deserializeOpGetTrustStore) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6076,10 +5949,8 @@ func (m *awsRestjson1_deserializeOpGetTrustStoreCertif
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6119,7 +5990,6 @@ func (m *awsRestjson1_deserializeOpGetTrustStoreCertif
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6249,10 +6119,8 @@ func (m *awsRestjson1_deserializeOpGetUserAccessLoggin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6292,7 +6160,6 @@ func (m *awsRestjson1_deserializeOpGetUserAccessLoggin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6413,10 +6280,8 @@ func (m *awsRestjson1_deserializeOpGetUserSettings) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6456,7 +6321,6 @@ func (m *awsRestjson1_deserializeOpGetUserSettings) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6577,10 +6441,8 @@ func (m *awsRestjson1_deserializeOpListBrowserSettings
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6620,7 +6482,6 @@ func (m *awsRestjson1_deserializeOpListBrowserSettings
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6747,10 +6608,8 @@ func (m *awsRestjson1_deserializeOpListDataProtectionS
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6790,7 +6649,6 @@ func (m *awsRestjson1_deserializeOpListDataProtectionS
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6917,10 +6775,8 @@ func (m *awsRestjson1_deserializeOpListIdentityProvide
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6960,7 +6816,6 @@ func (m *awsRestjson1_deserializeOpListIdentityProvide
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7087,10 +6942,8 @@ func (m *awsRestjson1_deserializeOpListIpAccessSetting
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7130,7 +6983,6 @@ func (m *awsRestjson1_deserializeOpListIpAccessSetting
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7257,10 +7109,8 @@ func (m *awsRestjson1_deserializeOpListNetworkSettings
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7300,7 +7150,6 @@ func (m *awsRestjson1_deserializeOpListNetworkSettings
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7427,10 +7276,8 @@ func (m *awsRestjson1_deserializeOpListPortals) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7470,7 +7317,6 @@ func (m *awsRestjson1_deserializeOpListPortals) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7597,10 +7443,8 @@ func (m *awsRestjson1_deserializeOpListSessions) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7640,7 +7484,6 @@ func (m *awsRestjson1_deserializeOpListSessions) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7770,10 +7613,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7813,7 +7654,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7934,10 +7774,8 @@ func (m *awsRestjson1_deserializeOpListTrustStoreCerti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7977,7 +7815,6 @@ func (m *awsRestjson1_deserializeOpListTrustStoreCerti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8116,10 +7953,8 @@ func (m *awsRestjson1_deserializeOpListTrustStores) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8159,7 +7994,6 @@ func (m *awsRestjson1_deserializeOpListTrustStores) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8286,10 +8120,8 @@ func (m *awsRestjson1_deserializeOpListUserAccessLoggi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8329,7 +8161,6 @@ func (m *awsRestjson1_deserializeOpListUserAccessLoggi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8456,10 +8287,8 @@ func (m *awsRestjson1_deserializeOpListUserSettings) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8499,7 +8328,6 @@ func (m *awsRestjson1_deserializeOpListUserSettings) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8626,10 +8454,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8641,7 +8467,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8729,10 +8554,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8744,7 +8567,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8829,10 +8651,8 @@ func (m *awsRestjson1_deserializeOpUpdateBrowserSettin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8872,7 +8692,6 @@ func (m *awsRestjson1_deserializeOpUpdateBrowserSettin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8993,10 +8812,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataProtectio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9036,7 +8853,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataProtectio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9157,10 +8973,8 @@ func (m *awsRestjson1_deserializeOpUpdateIdentityProvi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9200,7 +9014,6 @@ func (m *awsRestjson1_deserializeOpUpdateIdentityProvi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9321,10 +9134,8 @@ func (m *awsRestjson1_deserializeOpUpdateIpAccessSetti
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9364,7 +9175,6 @@ func (m *awsRestjson1_deserializeOpUpdateIpAccessSetti
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9485,10 +9295,8 @@ func (m *awsRestjson1_deserializeOpUpdateNetworkSettin
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9528,7 +9336,6 @@ func (m *awsRestjson1_deserializeOpUpdateNetworkSettin
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9649,10 +9456,8 @@ func (m *awsRestjson1_deserializeOpUpdatePortal) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9692,7 +9497,6 @@ func (m *awsRestjson1_deserializeOpUpdatePortal) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9819,10 +9623,8 @@ func (m *awsRestjson1_deserializeOpUpdateTrustStore) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9862,7 +9664,6 @@ func (m *awsRestjson1_deserializeOpUpdateTrustStore) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9990,10 +9791,8 @@ func (m *awsRestjson1_deserializeOpUpdateUserAccessLog
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10033,7 +9832,6 @@ func (m *awsRestjson1_deserializeOpUpdateUserAccessLog
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10154,10 +9952,8 @@ func (m *awsRestjson1_deserializeOpUpdateUserSettings)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10197,7 +9993,6 @@ func (m *awsRestjson1_deserializeOpUpdateUserSettings)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
