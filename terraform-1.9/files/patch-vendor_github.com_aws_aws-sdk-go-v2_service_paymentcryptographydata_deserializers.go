--- vendor/github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/paymentcryptographydata/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpDecryptData) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpDecryptData) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -230,10 +226,8 @@ func (m *awsRestjson1_deserializeOpEncryptData) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -273,7 +267,6 @@ func (m *awsRestjson1_deserializeOpEncryptData) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -416,10 +409,8 @@ func (m *awsRestjson1_deserializeOpGenerateCardValidat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -459,7 +450,6 @@ func (m *awsRestjson1_deserializeOpGenerateCardValidat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -602,10 +592,8 @@ func (m *awsRestjson1_deserializeOpGenerateMac) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -645,7 +633,6 @@ func (m *awsRestjson1_deserializeOpGenerateMac) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -788,10 +775,8 @@ func (m *awsRestjson1_deserializeOpGenerateMacEmvPinCh
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -831,7 +816,6 @@ func (m *awsRestjson1_deserializeOpGenerateMacEmvPinCh
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1024,10 +1008,8 @@ func (m *awsRestjson1_deserializeOpGeneratePinData) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1067,7 +1049,6 @@ func (m *awsRestjson1_deserializeOpGeneratePinData) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1233,10 +1214,8 @@ func (m *awsRestjson1_deserializeOpReEncryptData) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1276,7 +1255,6 @@ func (m *awsRestjson1_deserializeOpReEncryptData) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1419,10 +1397,8 @@ func (m *awsRestjson1_deserializeOpTranslatePinData) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1462,7 +1438,6 @@ func (m *awsRestjson1_deserializeOpTranslatePinData) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1605,10 +1580,8 @@ func (m *awsRestjson1_deserializeOpVerifyAuthRequestCr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1648,7 +1621,6 @@ func (m *awsRestjson1_deserializeOpVerifyAuthRequestCr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1794,10 +1766,8 @@ func (m *awsRestjson1_deserializeOpVerifyCardValidatio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1837,7 +1807,6 @@ func (m *awsRestjson1_deserializeOpVerifyCardValidatio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1974,10 +1943,8 @@ func (m *awsRestjson1_deserializeOpVerifyMac) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2017,7 +1984,6 @@ func (m *awsRestjson1_deserializeOpVerifyMac) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2154,10 +2120,8 @@ func (m *awsRestjson1_deserializeOpVerifyPinData) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2197,7 +2161,6 @@ func (m *awsRestjson1_deserializeOpVerifyPinData) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
