--- vendor/github.com/aws/aws-sdk-go-v2/service/socialmessaging/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/socialmessaging/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"strings"
@@ -44,10 +43,8 @@ func (m *awsRestjson1_deserializeOpAssociateWhatsAppBu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -87,7 +84,6 @@ func (m *awsRestjson1_deserializeOpAssociateWhatsAppBu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -221,10 +217,8 @@ func (m *awsRestjson1_deserializeOpDeleteWhatsAppMessa
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -264,7 +258,6 @@ func (m *awsRestjson1_deserializeOpDeleteWhatsAppMessa
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -398,10 +391,8 @@ func (m *awsRestjson1_deserializeOpDisassociateWhatsAp
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -413,7 +404,6 @@ func (m *awsRestjson1_deserializeOpDisassociateWhatsAp
 	output := &DisassociateWhatsAppBusinessAccountOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -501,10 +491,8 @@ func (m *awsRestjson1_deserializeOpGetLinkedWhatsAppBu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -544,7 +532,6 @@ func (m *awsRestjson1_deserializeOpGetLinkedWhatsAppBu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -671,10 +658,8 @@ func (m *awsRestjson1_deserializeOpGetLinkedWhatsAppBu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -714,7 +699,6 @@ func (m *awsRestjson1_deserializeOpGetLinkedWhatsAppBu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -850,10 +834,8 @@ func (m *awsRestjson1_deserializeOpGetWhatsAppMessageM
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -893,7 +875,6 @@ func (m *awsRestjson1_deserializeOpGetWhatsAppMessageM
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1040,10 +1021,8 @@ func (m *awsRestjson1_deserializeOpListLinkedWhatsAppB
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1083,7 +1062,6 @@ func (m *awsRestjson1_deserializeOpListLinkedWhatsAppB
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1216,10 +1194,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1259,7 +1235,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1393,10 +1368,8 @@ func (m *awsRestjson1_deserializeOpPostWhatsAppMessage
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1436,7 +1409,6 @@ func (m *awsRestjson1_deserializeOpPostWhatsAppMessage
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1570,10 +1542,8 @@ func (m *awsRestjson1_deserializeOpPutWhatsAppBusiness
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1585,7 +1555,6 @@ func (m *awsRestjson1_deserializeOpPutWhatsAppBusiness
 	output := &PutWhatsAppBusinessAccountEventDestinationsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1670,10 +1639,8 @@ func (m *awsRestjson1_deserializeOpSendWhatsAppMessage
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1713,7 +1680,6 @@ func (m *awsRestjson1_deserializeOpSendWhatsAppMessage
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1844,10 +1810,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1887,7 +1851,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2016,10 +1979,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2059,7 +2020,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
