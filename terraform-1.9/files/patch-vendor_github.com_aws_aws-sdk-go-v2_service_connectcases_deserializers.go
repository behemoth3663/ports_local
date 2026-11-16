--- vendor/github.com/aws/aws-sdk-go-v2/service/connectcases/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/connectcases/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -47,10 +46,8 @@ func (m *awsRestjson1_deserializeOpBatchGetCaseRule) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -90,7 +87,6 @@ func (m *awsRestjson1_deserializeOpBatchGetCaseRule) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -216,10 +212,8 @@ func (m *awsRestjson1_deserializeOpBatchGetField) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -259,7 +253,6 @@ func (m *awsRestjson1_deserializeOpBatchGetField) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -385,10 +378,8 @@ func (m *awsRestjson1_deserializeOpBatchPutFieldOption
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -428,7 +419,6 @@ func (m *awsRestjson1_deserializeOpBatchPutFieldOption
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -552,10 +542,8 @@ func (m *awsRestjson1_deserializeOpCreateCase) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -595,7 +583,6 @@ func (m *awsRestjson1_deserializeOpCreateCase) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -732,10 +719,8 @@ func (m *awsRestjson1_deserializeOpCreateCaseRule) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -775,7 +760,6 @@ func (m *awsRestjson1_deserializeOpCreateCaseRule) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -915,10 +899,8 @@ func (m *awsRestjson1_deserializeOpCreateDomain) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -958,7 +940,6 @@ func (m *awsRestjson1_deserializeOpCreateDomain) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1104,10 +1085,8 @@ func (m *awsRestjson1_deserializeOpCreateField) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1147,7 +1126,6 @@ func (m *awsRestjson1_deserializeOpCreateField) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1287,10 +1265,8 @@ func (m *awsRestjson1_deserializeOpCreateLayout) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1330,7 +1306,6 @@ func (m *awsRestjson1_deserializeOpCreateLayout) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1470,10 +1445,8 @@ func (m *awsRestjson1_deserializeOpCreateRelatedItem) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1513,7 +1486,6 @@ func (m *awsRestjson1_deserializeOpCreateRelatedItem) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1650,10 +1622,8 @@ func (m *awsRestjson1_deserializeOpCreateTemplate) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1693,7 +1663,6 @@ func (m *awsRestjson1_deserializeOpCreateTemplate) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1833,10 +1802,8 @@ func (m *awsRestjson1_deserializeOpDeleteCase) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1848,7 +1815,6 @@ func (m *awsRestjson1_deserializeOpDeleteCase) HandleD
 	output := &DeleteCaseOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1933,10 +1899,8 @@ func (m *awsRestjson1_deserializeOpDeleteCaseRule) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1948,7 +1912,6 @@ func (m *awsRestjson1_deserializeOpDeleteCaseRule) Han
 	output := &DeleteCaseRuleOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2033,10 +1996,8 @@ func (m *awsRestjson1_deserializeOpDeleteDomain) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2048,7 +2009,6 @@ func (m *awsRestjson1_deserializeOpDeleteDomain) Handl
 	output := &DeleteDomainOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2136,10 +2096,8 @@ func (m *awsRestjson1_deserializeOpDeleteField) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2151,7 +2109,6 @@ func (m *awsRestjson1_deserializeOpDeleteField) Handle
 	output := &DeleteFieldOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2242,10 +2199,8 @@ func (m *awsRestjson1_deserializeOpDeleteLayout) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2257,7 +2212,6 @@ func (m *awsRestjson1_deserializeOpDeleteLayout) Handl
 	output := &DeleteLayoutOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2345,10 +2299,8 @@ func (m *awsRestjson1_deserializeOpDeleteRelatedItem) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2360,7 +2312,6 @@ func (m *awsRestjson1_deserializeOpDeleteRelatedItem) 
 	output := &DeleteRelatedItemOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2445,10 +2396,8 @@ func (m *awsRestjson1_deserializeOpDeleteTemplate) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2460,7 +2409,6 @@ func (m *awsRestjson1_deserializeOpDeleteTemplate) Han
 	output := &DeleteTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2548,10 +2496,8 @@ func (m *awsRestjson1_deserializeOpGetCase) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2591,7 +2537,6 @@ func (m *awsRestjson1_deserializeOpGetCase) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2735,10 +2680,8 @@ func (m *awsRestjson1_deserializeOpGetCaseAuditEvents)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2778,7 +2721,6 @@ func (m *awsRestjson1_deserializeOpGetCaseAuditEvents)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2908,10 +2850,8 @@ func (m *awsRestjson1_deserializeOpGetCaseEventConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2951,7 +2891,6 @@ func (m *awsRestjson1_deserializeOpGetCaseEventConfigu
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3072,10 +3011,8 @@ func (m *awsRestjson1_deserializeOpGetDomain) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3115,7 +3052,6 @@ func (m *awsRestjson1_deserializeOpGetDomain) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3285,10 +3221,8 @@ func (m *awsRestjson1_deserializeOpGetLayout) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3328,7 +3262,6 @@ func (m *awsRestjson1_deserializeOpGetLayout) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3516,10 +3449,8 @@ func (m *awsRestjson1_deserializeOpGetTemplate) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3559,7 +3490,6 @@ func (m *awsRestjson1_deserializeOpGetTemplate) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3775,10 +3705,8 @@ func (m *awsRestjson1_deserializeOpListCaseRules) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3818,7 +3746,6 @@ func (m *awsRestjson1_deserializeOpListCaseRules) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3948,10 +3875,8 @@ func (m *awsRestjson1_deserializeOpListCasesForContact
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3991,7 +3916,6 @@ func (m *awsRestjson1_deserializeOpListCasesForContact
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4121,10 +4045,8 @@ func (m *awsRestjson1_deserializeOpListDomains) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4164,7 +4086,6 @@ func (m *awsRestjson1_deserializeOpListDomains) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4291,10 +4212,8 @@ func (m *awsRestjson1_deserializeOpListFieldOptions) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4334,7 +4253,6 @@ func (m *awsRestjson1_deserializeOpListFieldOptions) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4464,10 +4382,8 @@ func (m *awsRestjson1_deserializeOpListFields) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4507,7 +4423,6 @@ func (m *awsRestjson1_deserializeOpListFields) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4637,10 +4552,8 @@ func (m *awsRestjson1_deserializeOpListLayouts) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4680,7 +4593,6 @@ func (m *awsRestjson1_deserializeOpListLayouts) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4810,10 +4722,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4853,7 +4763,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4974,10 +4883,8 @@ func (m *awsRestjson1_deserializeOpListTemplates) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5017,7 +4924,6 @@ func (m *awsRestjson1_deserializeOpListTemplates) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5147,10 +5053,8 @@ func (m *awsRestjson1_deserializeOpPutCaseEventConfigu
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5162,7 +5066,6 @@ func (m *awsRestjson1_deserializeOpPutCaseEventConfigu
 	output := &PutCaseEventConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5250,10 +5153,8 @@ func (m *awsRestjson1_deserializeOpSearchCases) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5293,7 +5194,6 @@ func (m *awsRestjson1_deserializeOpSearchCases) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5423,10 +5323,8 @@ func (m *awsRestjson1_deserializeOpSearchRelatedItems)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5466,7 +5364,6 @@ func (m *awsRestjson1_deserializeOpSearchRelatedItems)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5596,10 +5493,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5617,7 +5512,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5702,10 +5596,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5723,7 +5615,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5808,10 +5699,8 @@ func (m *awsRestjson1_deserializeOpUpdateCase) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5823,7 +5712,6 @@ func (m *awsRestjson1_deserializeOpUpdateCase) HandleD
 	output := &UpdateCaseOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5908,10 +5796,8 @@ func (m *awsRestjson1_deserializeOpUpdateCaseRule) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5923,7 +5809,6 @@ func (m *awsRestjson1_deserializeOpUpdateCaseRule) Han
 	output := &UpdateCaseRuleOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6011,10 +5896,8 @@ func (m *awsRestjson1_deserializeOpUpdateField) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6026,7 +5909,6 @@ func (m *awsRestjson1_deserializeOpUpdateField) Handle
 	output := &UpdateFieldOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6114,10 +5996,8 @@ func (m *awsRestjson1_deserializeOpUpdateLayout) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6129,7 +6009,6 @@ func (m *awsRestjson1_deserializeOpUpdateLayout) Handl
 	output := &UpdateLayoutOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6220,10 +6099,8 @@ func (m *awsRestjson1_deserializeOpUpdateTemplate) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6235,7 +6112,6 @@ func (m *awsRestjson1_deserializeOpUpdateTemplate) Han
 	output := &UpdateTemplateOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
