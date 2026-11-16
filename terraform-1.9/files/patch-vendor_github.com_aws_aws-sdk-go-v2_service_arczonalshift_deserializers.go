--- vendor/github.com/aws/aws-sdk-go-v2/service/arczonalshift/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/arczonalshift/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"math"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpCancelPracticeRun) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpCancelPracticeRun) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -284,10 +280,8 @@ func (m *awsRestjson1_deserializeOpCancelZonalShift) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -327,7 +321,6 @@ func (m *awsRestjson1_deserializeOpCancelZonalShift) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -523,10 +516,8 @@ func (m *awsRestjson1_deserializeOpCreatePracticeRunCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -566,7 +557,6 @@ func (m *awsRestjson1_deserializeOpCreatePracticeRunCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -717,10 +707,8 @@ func (m *awsRestjson1_deserializeOpDeletePracticeRunCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -760,7 +748,6 @@ func (m *awsRestjson1_deserializeOpDeletePracticeRunCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -906,10 +893,8 @@ func (m *awsRestjson1_deserializeOpGetAutoshiftObserve
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -949,7 +934,6 @@ func (m *awsRestjson1_deserializeOpGetAutoshiftObserve
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1068,10 +1052,8 @@ func (m *awsRestjson1_deserializeOpGetManagedResource)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1111,7 +1093,6 @@ func (m *awsRestjson1_deserializeOpGetManagedResource)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1274,10 +1255,8 @@ func (m *awsRestjson1_deserializeOpListAutoshifts) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1317,7 +1296,6 @@ func (m *awsRestjson1_deserializeOpListAutoshifts) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1444,10 +1422,8 @@ func (m *awsRestjson1_deserializeOpListManagedResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1487,7 +1463,6 @@ func (m *awsRestjson1_deserializeOpListManagedResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1614,10 +1589,8 @@ func (m *awsRestjson1_deserializeOpListZonalShifts) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1657,7 +1630,6 @@ func (m *awsRestjson1_deserializeOpListZonalShifts) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1784,10 +1756,8 @@ func (m *awsRestjson1_deserializeOpStartPracticeRun) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1827,7 +1797,6 @@ func (m *awsRestjson1_deserializeOpStartPracticeRun) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2023,10 +1992,8 @@ func (m *awsRestjson1_deserializeOpStartZonalShift) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2066,7 +2033,6 @@ func (m *awsRestjson1_deserializeOpStartZonalShift) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2262,10 +2228,8 @@ func (m *awsRestjson1_deserializeOpUpdateAutoshiftObse
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2305,7 +2269,6 @@ func (m *awsRestjson1_deserializeOpUpdateAutoshiftObse
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2427,10 +2390,8 @@ func (m *awsRestjson1_deserializeOpUpdatePracticeRunCo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2470,7 +2431,6 @@ func (m *awsRestjson1_deserializeOpUpdatePracticeRunCo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2621,10 +2581,8 @@ func (m *awsRestjson1_deserializeOpUpdateZonalAutoshif
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2664,7 +2622,6 @@ func (m *awsRestjson1_deserializeOpUpdateZonalAutoshif
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2801,10 +2758,8 @@ func (m *awsRestjson1_deserializeOpUpdateZonalShift) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2844,7 +2799,6 @@ func (m *awsRestjson1_deserializeOpUpdateZonalShift) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
