--- vendor/github.com/aws/aws-sdk-go-v2/service/datazone/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/datazone/deserializers.go
@@ -14,7 +14,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -45,10 +44,8 @@ func (m *awsRestjson1_deserializeOpAcceptPredictions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -88,7 +85,6 @@ func (m *awsRestjson1_deserializeOpAcceptPredictions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -237,10 +233,8 @@ func (m *awsRestjson1_deserializeOpAcceptSubscriptionR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -280,7 +274,6 @@ func (m *awsRestjson1_deserializeOpAcceptSubscriptionR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -530,10 +523,8 @@ func (m *awsRestjson1_deserializeOpAddEntityOwner) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -545,7 +536,6 @@ func (m *awsRestjson1_deserializeOpAddEntityOwner) Han
 	output := &AddEntityOwnerOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -639,10 +629,8 @@ func (m *awsRestjson1_deserializeOpAddPolicyGrant) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -654,7 +642,6 @@ func (m *awsRestjson1_deserializeOpAddPolicyGrant) Han
 	output := &AddPolicyGrantOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -745,10 +732,8 @@ func (m *awsRestjson1_deserializeOpAssociateEnvironmen
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -760,7 +745,6 @@ func (m *awsRestjson1_deserializeOpAssociateEnvironmen
 	output := &AssociateEnvironmentRoleOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -851,10 +835,8 @@ func (m *awsRestjson1_deserializeOpCancelMetadataGener
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -866,7 +848,6 @@ func (m *awsRestjson1_deserializeOpCancelMetadataGener
 	output := &CancelMetadataGenerationRunOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -957,10 +938,8 @@ func (m *awsRestjson1_deserializeOpCancelSubscription)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1000,7 +979,6 @@ func (m *awsRestjson1_deserializeOpCancelSubscription)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1227,10 +1205,8 @@ func (m *awsRestjson1_deserializeOpCreateAsset) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1270,7 +1246,6 @@ func (m *awsRestjson1_deserializeOpCreateAsset) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1556,10 +1531,8 @@ func (m *awsRestjson1_deserializeOpCreateAssetFilter) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1599,7 +1572,6 @@ func (m *awsRestjson1_deserializeOpCreateAssetFilter) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -1822,10 +1794,8 @@ func (m *awsRestjson1_deserializeOpCreateAssetRevision
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -1865,7 +1835,6 @@ func (m *awsRestjson1_deserializeOpCreateAssetRevision
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2148,10 +2117,8 @@ func (m *awsRestjson1_deserializeOpCreateAssetType) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2191,7 +2158,6 @@ func (m *awsRestjson1_deserializeOpCreateAssetType) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2431,10 +2397,8 @@ func (m *awsRestjson1_deserializeOpCreateConnection) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2474,7 +2438,6 @@ func (m *awsRestjson1_deserializeOpCreateConnection) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2681,10 +2644,8 @@ func (m *awsRestjson1_deserializeOpCreateDataProduct) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -2724,7 +2685,6 @@ func (m *awsRestjson1_deserializeOpCreateDataProduct) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -2977,10 +2937,8 @@ func (m *awsRestjson1_deserializeOpCreateDataProductRe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3020,7 +2978,6 @@ func (m *awsRestjson1_deserializeOpCreateDataProductRe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3270,10 +3227,8 @@ func (m *awsRestjson1_deserializeOpCreateDataSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3313,7 +3268,6 @@ func (m *awsRestjson1_deserializeOpCreateDataSource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3615,10 +3569,8 @@ func (m *awsRestjson1_deserializeOpCreateDomain) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3658,7 +3610,6 @@ func (m *awsRestjson1_deserializeOpCreateDomain) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -3892,10 +3843,8 @@ func (m *awsRestjson1_deserializeOpCreateDomainUnit) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -3935,7 +3884,6 @@ func (m *awsRestjson1_deserializeOpCreateDomainUnit) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4137,10 +4085,8 @@ func (m *awsRestjson1_deserializeOpCreateEnvironment) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4180,7 +4126,6 @@ func (m *awsRestjson1_deserializeOpCreateEnvironment) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4480,10 +4425,8 @@ func (m *awsRestjson1_deserializeOpCreateEnvironmentAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4523,7 +4466,6 @@ func (m *awsRestjson1_deserializeOpCreateEnvironmentAc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4695,10 +4637,8 @@ func (m *awsRestjson1_deserializeOpCreateEnvironmentPr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -4738,7 +4678,6 @@ func (m *awsRestjson1_deserializeOpCreateEnvironmentPr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -4975,10 +4914,8 @@ func (m *awsRestjson1_deserializeOpCreateFormType) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5018,7 +4955,6 @@ func (m *awsRestjson1_deserializeOpCreateFormType) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5203,10 +5139,8 @@ func (m *awsRestjson1_deserializeOpCreateGlossary) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5246,7 +5180,6 @@ func (m *awsRestjson1_deserializeOpCreateGlossary) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5422,10 +5355,8 @@ func (m *awsRestjson1_deserializeOpCreateGlossaryTerm)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5465,7 +5396,6 @@ func (m *awsRestjson1_deserializeOpCreateGlossaryTerm)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5658,10 +5588,8 @@ func (m *awsRestjson1_deserializeOpCreateGroupProfile)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5701,7 +5629,6 @@ func (m *awsRestjson1_deserializeOpCreateGroupProfile)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -5856,10 +5783,8 @@ func (m *awsRestjson1_deserializeOpCreateListingChange
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -5899,7 +5824,6 @@ func (m *awsRestjson1_deserializeOpCreateListingChange
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6051,10 +5975,8 @@ func (m *awsRestjson1_deserializeOpCreateProject) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6094,7 +6016,6 @@ func (m *awsRestjson1_deserializeOpCreateProject) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6337,10 +6258,8 @@ func (m *awsRestjson1_deserializeOpCreateProjectMember
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6352,7 +6271,6 @@ func (m *awsRestjson1_deserializeOpCreateProjectMember
 	output := &CreateProjectMembershipOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6440,10 +6358,8 @@ func (m *awsRestjson1_deserializeOpCreateProjectProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6483,7 +6399,6 @@ func (m *awsRestjson1_deserializeOpCreateProjectProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6702,10 +6617,8 @@ func (m *awsRestjson1_deserializeOpCreateRule) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -6745,7 +6658,6 @@ func (m *awsRestjson1_deserializeOpCreateRule) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -6964,10 +6876,8 @@ func (m *awsRestjson1_deserializeOpCreateSubscriptionG
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7007,7 +6917,6 @@ func (m *awsRestjson1_deserializeOpCreateSubscriptionG
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7234,10 +7143,8 @@ func (m *awsRestjson1_deserializeOpCreateSubscriptionR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7277,7 +7184,6 @@ func (m *awsRestjson1_deserializeOpCreateSubscriptionR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7527,10 +7433,8 @@ func (m *awsRestjson1_deserializeOpCreateSubscriptionT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7570,7 +7474,6 @@ func (m *awsRestjson1_deserializeOpCreateSubscriptionT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -7829,10 +7732,8 @@ func (m *awsRestjson1_deserializeOpCreateUserProfile) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -7872,7 +7773,6 @@ func (m *awsRestjson1_deserializeOpCreateUserProfile) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8032,10 +7932,8 @@ func (m *awsRestjson1_deserializeOpDeleteAsset) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8047,7 +7945,6 @@ func (m *awsRestjson1_deserializeOpDeleteAsset) Handle
 	output := &DeleteAssetOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8138,10 +8035,8 @@ func (m *awsRestjson1_deserializeOpDeleteAssetFilter) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8159,7 +8054,6 @@ func (m *awsRestjson1_deserializeOpDeleteAssetFilter) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8250,10 +8144,8 @@ func (m *awsRestjson1_deserializeOpDeleteAssetType) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8265,7 +8157,6 @@ func (m *awsRestjson1_deserializeOpDeleteAssetType) Ha
 	output := &DeleteAssetTypeOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8356,10 +8247,8 @@ func (m *awsRestjson1_deserializeOpDeleteConnection) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8399,7 +8288,6 @@ func (m *awsRestjson1_deserializeOpDeleteConnection) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8527,10 +8415,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataProduct) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8542,7 +8428,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataProduct) 
 	output := &DeleteDataProductOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8633,10 +8518,8 @@ func (m *awsRestjson1_deserializeOpDeleteDataSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -8676,7 +8559,6 @@ func (m *awsRestjson1_deserializeOpDeleteDataSource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -8987,10 +8869,8 @@ func (m *awsRestjson1_deserializeOpDeleteDomain) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9030,7 +8910,6 @@ func (m *awsRestjson1_deserializeOpDeleteDomain) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9161,10 +9040,8 @@ func (m *awsRestjson1_deserializeOpDeleteDomainUnit) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9176,7 +9053,6 @@ func (m *awsRestjson1_deserializeOpDeleteDomainUnit) H
 	output := &DeleteDomainUnitOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9267,10 +9143,8 @@ func (m *awsRestjson1_deserializeOpDeleteEnvironment) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9288,7 +9162,6 @@ func (m *awsRestjson1_deserializeOpDeleteEnvironment) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9376,10 +9249,8 @@ func (m *awsRestjson1_deserializeOpDeleteEnvironmentAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9397,7 +9268,6 @@ func (m *awsRestjson1_deserializeOpDeleteEnvironmentAc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9488,10 +9358,8 @@ func (m *awsRestjson1_deserializeOpDeleteEnvironmentBl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9503,7 +9371,6 @@ func (m *awsRestjson1_deserializeOpDeleteEnvironmentBl
 	output := &DeleteEnvironmentBlueprintConfigurationOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9588,10 +9455,8 @@ func (m *awsRestjson1_deserializeOpDeleteEnvironmentPr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9609,7 +9474,6 @@ func (m *awsRestjson1_deserializeOpDeleteEnvironmentPr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9697,10 +9561,8 @@ func (m *awsRestjson1_deserializeOpDeleteFormType) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9712,7 +9574,6 @@ func (m *awsRestjson1_deserializeOpDeleteFormType) Han
 	output := &DeleteFormTypeOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9803,10 +9664,8 @@ func (m *awsRestjson1_deserializeOpDeleteGlossary) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9818,7 +9677,6 @@ func (m *awsRestjson1_deserializeOpDeleteGlossary) Han
 	output := &DeleteGlossaryOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -9909,10 +9767,8 @@ func (m *awsRestjson1_deserializeOpDeleteGlossaryTerm)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -9924,7 +9780,6 @@ func (m *awsRestjson1_deserializeOpDeleteGlossaryTerm)
 	output := &DeleteGlossaryTermOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10015,10 +9870,8 @@ func (m *awsRestjson1_deserializeOpDeleteListing) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10030,7 +9883,6 @@ func (m *awsRestjson1_deserializeOpDeleteListing) Hand
 	output := &DeleteListingOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10121,10 +9973,8 @@ func (m *awsRestjson1_deserializeOpDeleteProject) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10136,7 +9986,6 @@ func (m *awsRestjson1_deserializeOpDeleteProject) Hand
 	output := &DeleteProjectOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10224,10 +10073,8 @@ func (m *awsRestjson1_deserializeOpDeleteProjectMember
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10239,7 +10086,6 @@ func (m *awsRestjson1_deserializeOpDeleteProjectMember
 	output := &DeleteProjectMembershipOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10330,10 +10176,8 @@ func (m *awsRestjson1_deserializeOpDeleteProjectProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10345,7 +10189,6 @@ func (m *awsRestjson1_deserializeOpDeleteProjectProfil
 	output := &DeleteProjectProfileOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10433,10 +10276,8 @@ func (m *awsRestjson1_deserializeOpDeleteRule) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10448,7 +10289,6 @@ func (m *awsRestjson1_deserializeOpDeleteRule) HandleD
 	output := &DeleteRuleOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10539,10 +10379,8 @@ func (m *awsRestjson1_deserializeOpDeleteSubscriptionG
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10582,7 +10420,6 @@ func (m *awsRestjson1_deserializeOpDeleteSubscriptionG
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10809,10 +10646,8 @@ func (m *awsRestjson1_deserializeOpDeleteSubscriptionR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10830,7 +10665,6 @@ func (m *awsRestjson1_deserializeOpDeleteSubscriptionR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -10921,10 +10755,8 @@ func (m *awsRestjson1_deserializeOpDeleteSubscriptionT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -10942,7 +10774,6 @@ func (m *awsRestjson1_deserializeOpDeleteSubscriptionT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11033,10 +10864,8 @@ func (m *awsRestjson1_deserializeOpDeleteTimeSeriesDat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11048,7 +10877,6 @@ func (m *awsRestjson1_deserializeOpDeleteTimeSeriesDat
 	output := &DeleteTimeSeriesDataPointsOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11136,10 +10964,8 @@ func (m *awsRestjson1_deserializeOpDisassociateEnviron
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11151,7 +10977,6 @@ func (m *awsRestjson1_deserializeOpDisassociateEnviron
 	output := &DisassociateEnvironmentRoleOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11242,10 +11067,8 @@ func (m *awsRestjson1_deserializeOpGetAsset) HandleDes
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11285,7 +11108,6 @@ func (m *awsRestjson1_deserializeOpGetAsset) HandleDes
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11560,10 +11382,8 @@ func (m *awsRestjson1_deserializeOpGetAssetFilter) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11603,7 +11423,6 @@ func (m *awsRestjson1_deserializeOpGetAssetFilter) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -11820,10 +11639,8 @@ func (m *awsRestjson1_deserializeOpGetAssetType) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -11863,7 +11680,6 @@ func (m *awsRestjson1_deserializeOpGetAssetType) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12100,10 +11916,8 @@ func (m *awsRestjson1_deserializeOpGetConnection) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12143,7 +11957,6 @@ func (m *awsRestjson1_deserializeOpGetConnection) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12358,10 +12171,8 @@ func (m *awsRestjson1_deserializeOpGetDataProduct) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12401,7 +12212,6 @@ func (m *awsRestjson1_deserializeOpGetDataProduct) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -12648,10 +12458,8 @@ func (m *awsRestjson1_deserializeOpGetDataSource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -12691,7 +12499,6 @@ func (m *awsRestjson1_deserializeOpGetDataSource) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13011,10 +12818,8 @@ func (m *awsRestjson1_deserializeOpGetDataSourceRun) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13054,7 +12859,6 @@ func (m *awsRestjson1_deserializeOpGetDataSourceRun) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13309,10 +13113,8 @@ func (m *awsRestjson1_deserializeOpGetDomain) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13352,7 +13154,6 @@ func (m *awsRestjson1_deserializeOpGetDomain) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13615,10 +13416,8 @@ func (m *awsRestjson1_deserializeOpGetDomainUnit) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13658,7 +13457,6 @@ func (m *awsRestjson1_deserializeOpGetDomainUnit) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -13877,10 +13675,8 @@ func (m *awsRestjson1_deserializeOpGetEnvironment) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -13920,7 +13716,6 @@ func (m *awsRestjson1_deserializeOpGetEnvironment) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14217,10 +14012,8 @@ func (m *awsRestjson1_deserializeOpGetEnvironmentActio
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14260,7 +14053,6 @@ func (m *awsRestjson1_deserializeOpGetEnvironmentActio
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14429,10 +14221,8 @@ func (m *awsRestjson1_deserializeOpGetEnvironmentBluep
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14472,7 +14262,6 @@ func (m *awsRestjson1_deserializeOpGetEnvironmentBluep
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14673,10 +14462,8 @@ func (m *awsRestjson1_deserializeOpGetEnvironmentBluep
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14716,7 +14503,6 @@ func (m *awsRestjson1_deserializeOpGetEnvironmentBluep
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -14921,10 +14707,8 @@ func (m *awsRestjson1_deserializeOpGetEnvironmentCrede
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -14964,7 +14748,6 @@ func (m *awsRestjson1_deserializeOpGetEnvironmentCrede
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15123,10 +14906,8 @@ func (m *awsRestjson1_deserializeOpGetEnvironmentProfi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15166,7 +14947,6 @@ func (m *awsRestjson1_deserializeOpGetEnvironmentProfi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15397,10 +15177,8 @@ func (m *awsRestjson1_deserializeOpGetFormType) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15440,7 +15218,6 @@ func (m *awsRestjson1_deserializeOpGetFormType) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15666,10 +15443,8 @@ func (m *awsRestjson1_deserializeOpGetGlossary) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15709,7 +15484,6 @@ func (m *awsRestjson1_deserializeOpGetGlossary) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -15932,10 +15706,8 @@ func (m *awsRestjson1_deserializeOpGetGlossaryTerm) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -15975,7 +15747,6 @@ func (m *awsRestjson1_deserializeOpGetGlossaryTerm) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16212,10 +15983,8 @@ func (m *awsRestjson1_deserializeOpGetGroupProfile) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16255,7 +16024,6 @@ func (m *awsRestjson1_deserializeOpGetGroupProfile) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16410,10 +16178,8 @@ func (m *awsRestjson1_deserializeOpGetIamPortalLoginUr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16453,7 +16219,6 @@ func (m *awsRestjson1_deserializeOpGetIamPortalLoginUr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16593,10 +16358,8 @@ func (m *awsRestjson1_deserializeOpGetJobRun) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16636,7 +16399,6 @@ func (m *awsRestjson1_deserializeOpGetJobRun) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -16876,10 +16638,8 @@ func (m *awsRestjson1_deserializeOpGetLineageEvent) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -16901,7 +16661,6 @@ func (m *awsRestjson1_deserializeOpGetLineageEvent) Ha
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("failed to deserialize response payload, %w", err)}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17056,10 +16815,8 @@ func (m *awsRestjson1_deserializeOpGetLineageNode) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17099,7 +16856,6 @@ func (m *awsRestjson1_deserializeOpGetLineageNode) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17362,10 +17118,8 @@ func (m *awsRestjson1_deserializeOpGetListing) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17405,7 +17159,6 @@ func (m *awsRestjson1_deserializeOpGetListing) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17633,10 +17386,8 @@ func (m *awsRestjson1_deserializeOpGetMetadataGenerati
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17676,7 +17427,6 @@ func (m *awsRestjson1_deserializeOpGetMetadataGenerati
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -17870,10 +17620,8 @@ func (m *awsRestjson1_deserializeOpGetProject) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -17913,7 +17661,6 @@ func (m *awsRestjson1_deserializeOpGetProject) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18150,10 +17897,8 @@ func (m *awsRestjson1_deserializeOpGetProjectProfile) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18193,7 +17938,6 @@ func (m *awsRestjson1_deserializeOpGetProjectProfile) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18406,10 +18150,8 @@ func (m *awsRestjson1_deserializeOpGetRule) HandleDese
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18449,7 +18191,6 @@ func (m *awsRestjson1_deserializeOpGetRule) HandleDese
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18696,10 +18437,8 @@ func (m *awsRestjson1_deserializeOpGetSubscription) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -18739,7 +18478,6 @@ func (m *awsRestjson1_deserializeOpGetSubscription) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -18963,10 +18701,8 @@ func (m *awsRestjson1_deserializeOpGetSubscriptionGran
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19006,7 +18742,6 @@ func (m *awsRestjson1_deserializeOpGetSubscriptionGran
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19230,10 +18965,8 @@ func (m *awsRestjson1_deserializeOpGetSubscriptionRequ
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19273,7 +19006,6 @@ func (m *awsRestjson1_deserializeOpGetSubscriptionRequ
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19520,10 +19252,8 @@ func (m *awsRestjson1_deserializeOpGetSubscriptionTarg
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19563,7 +19293,6 @@ func (m *awsRestjson1_deserializeOpGetSubscriptionTarg
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -19819,10 +19548,8 @@ func (m *awsRestjson1_deserializeOpGetTimeSeriesDataPo
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -19862,7 +19589,6 @@ func (m *awsRestjson1_deserializeOpGetTimeSeriesDataPo
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20022,10 +19748,8 @@ func (m *awsRestjson1_deserializeOpGetUserProfile) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20065,7 +19789,6 @@ func (m *awsRestjson1_deserializeOpGetUserProfile) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20225,10 +19948,8 @@ func (m *awsRestjson1_deserializeOpListAssetFilters) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20268,7 +19989,6 @@ func (m *awsRestjson1_deserializeOpListAssetFilters) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20401,10 +20121,8 @@ func (m *awsRestjson1_deserializeOpListAssetRevisions)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20444,7 +20162,6 @@ func (m *awsRestjson1_deserializeOpListAssetRevisions)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20577,10 +20294,8 @@ func (m *awsRestjson1_deserializeOpListConnections) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20620,7 +20335,6 @@ func (m *awsRestjson1_deserializeOpListConnections) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20750,10 +20464,8 @@ func (m *awsRestjson1_deserializeOpListDataProductRevi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20793,7 +20505,6 @@ func (m *awsRestjson1_deserializeOpListDataProductRevi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -20926,10 +20637,8 @@ func (m *awsRestjson1_deserializeOpListDataSourceRunAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -20969,7 +20678,6 @@ func (m *awsRestjson1_deserializeOpListDataSourceRunAc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21108,10 +20816,8 @@ func (m *awsRestjson1_deserializeOpListDataSourceRuns)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21151,7 +20857,6 @@ func (m *awsRestjson1_deserializeOpListDataSourceRuns)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21290,10 +20995,8 @@ func (m *awsRestjson1_deserializeOpListDataSources) Ha
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21333,7 +21036,6 @@ func (m *awsRestjson1_deserializeOpListDataSources) Ha
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21472,10 +21174,8 @@ func (m *awsRestjson1_deserializeOpListDomains) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21515,7 +21215,6 @@ func (m *awsRestjson1_deserializeOpListDomains) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21654,10 +21353,8 @@ func (m *awsRestjson1_deserializeOpListDomainUnitsForP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21697,7 +21394,6 @@ func (m *awsRestjson1_deserializeOpListDomainUnitsForP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -21827,10 +21523,8 @@ func (m *awsRestjson1_deserializeOpListEntityOwners) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -21870,7 +21564,6 @@ func (m *awsRestjson1_deserializeOpListEntityOwners) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22000,10 +21693,8 @@ func (m *awsRestjson1_deserializeOpListEnvironmentActi
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22043,7 +21734,6 @@ func (m *awsRestjson1_deserializeOpListEnvironmentActi
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22173,10 +21863,8 @@ func (m *awsRestjson1_deserializeOpListEnvironmentBlue
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22216,7 +21904,6 @@ func (m *awsRestjson1_deserializeOpListEnvironmentBlue
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22349,10 +22036,8 @@ func (m *awsRestjson1_deserializeOpListEnvironmentBlue
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22392,7 +22077,6 @@ func (m *awsRestjson1_deserializeOpListEnvironmentBlue
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22525,10 +22209,8 @@ func (m *awsRestjson1_deserializeOpListEnvironmentProf
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22568,7 +22250,6 @@ func (m *awsRestjson1_deserializeOpListEnvironmentProf
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22698,10 +22379,8 @@ func (m *awsRestjson1_deserializeOpListEnvironments) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22741,7 +22420,6 @@ func (m *awsRestjson1_deserializeOpListEnvironments) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -22871,10 +22549,8 @@ func (m *awsRestjson1_deserializeOpListJobRuns) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -22914,7 +22590,6 @@ func (m *awsRestjson1_deserializeOpListJobRuns) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23047,10 +22722,8 @@ func (m *awsRestjson1_deserializeOpListLineageEvents) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23090,7 +22763,6 @@ func (m *awsRestjson1_deserializeOpListLineageEvents) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23220,10 +22892,8 @@ func (m *awsRestjson1_deserializeOpListLineageNodeHist
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23263,7 +22933,6 @@ func (m *awsRestjson1_deserializeOpListLineageNodeHist
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23396,10 +23065,8 @@ func (m *awsRestjson1_deserializeOpListMetadataGenerat
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23439,7 +23106,6 @@ func (m *awsRestjson1_deserializeOpListMetadataGenerat
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23572,10 +23238,8 @@ func (m *awsRestjson1_deserializeOpListNotifications) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23615,7 +23279,6 @@ func (m *awsRestjson1_deserializeOpListNotifications) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23748,10 +23411,8 @@ func (m *awsRestjson1_deserializeOpListPolicyGrants) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23791,7 +23452,6 @@ func (m *awsRestjson1_deserializeOpListPolicyGrants) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -23921,10 +23581,8 @@ func (m *awsRestjson1_deserializeOpListProjectMembersh
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -23964,7 +23622,6 @@ func (m *awsRestjson1_deserializeOpListProjectMembersh
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24097,10 +23754,8 @@ func (m *awsRestjson1_deserializeOpListProjectProfiles
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24140,7 +23795,6 @@ func (m *awsRestjson1_deserializeOpListProjectProfiles
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24270,10 +23924,8 @@ func (m *awsRestjson1_deserializeOpListProjects) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24313,7 +23965,6 @@ func (m *awsRestjson1_deserializeOpListProjects) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24443,10 +24094,8 @@ func (m *awsRestjson1_deserializeOpListRules) HandleDe
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24486,7 +24135,6 @@ func (m *awsRestjson1_deserializeOpListRules) HandleDe
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24619,10 +24267,8 @@ func (m *awsRestjson1_deserializeOpListSubscriptionGra
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24662,7 +24308,6 @@ func (m *awsRestjson1_deserializeOpListSubscriptionGra
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24795,10 +24440,8 @@ func (m *awsRestjson1_deserializeOpListSubscriptionReq
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -24838,7 +24481,6 @@ func (m *awsRestjson1_deserializeOpListSubscriptionReq
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -24971,10 +24613,8 @@ func (m *awsRestjson1_deserializeOpListSubscriptions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25014,7 +24654,6 @@ func (m *awsRestjson1_deserializeOpListSubscriptions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25147,10 +24786,8 @@ func (m *awsRestjson1_deserializeOpListSubscriptionTar
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25190,7 +24827,6 @@ func (m *awsRestjson1_deserializeOpListSubscriptionTar
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25323,10 +24959,8 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25366,7 +25000,6 @@ func (m *awsRestjson1_deserializeOpListTagsForResource
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25490,10 +25123,8 @@ func (m *awsRestjson1_deserializeOpListTimeSeriesDataP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25533,7 +25164,6 @@ func (m *awsRestjson1_deserializeOpListTimeSeriesDataP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25666,10 +25296,8 @@ func (m *awsRestjson1_deserializeOpPostLineageEvent) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25709,7 +25337,6 @@ func (m *awsRestjson1_deserializeOpPostLineageEvent) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -25852,10 +25479,8 @@ func (m *awsRestjson1_deserializeOpPostTimeSeriesDataP
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -25895,7 +25520,6 @@ func (m *awsRestjson1_deserializeOpPostTimeSeriesDataP
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26052,10 +25676,8 @@ func (m *awsRestjson1_deserializeOpPutEnvironmentBluep
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26095,7 +25717,6 @@ func (m *awsRestjson1_deserializeOpPutEnvironmentBluep
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26303,10 +25924,8 @@ func (m *awsRestjson1_deserializeOpRejectPredictions) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26346,7 +25965,6 @@ func (m *awsRestjson1_deserializeOpRejectPredictions) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26495,10 +26113,8 @@ func (m *awsRestjson1_deserializeOpRejectSubscriptionR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26538,7 +26154,6 @@ func (m *awsRestjson1_deserializeOpRejectSubscriptionR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26788,10 +26403,8 @@ func (m *awsRestjson1_deserializeOpRemoveEntityOwner) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26803,7 +26416,6 @@ func (m *awsRestjson1_deserializeOpRemoveEntityOwner) 
 	output := &RemoveEntityOwnerOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26891,10 +26503,8 @@ func (m *awsRestjson1_deserializeOpRemovePolicyGrant) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -26906,7 +26516,6 @@ func (m *awsRestjson1_deserializeOpRemovePolicyGrant) 
 	output := &RemovePolicyGrantOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -26991,10 +26600,8 @@ func (m *awsRestjson1_deserializeOpRevokeSubscription)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27034,7 +26641,6 @@ func (m *awsRestjson1_deserializeOpRevokeSubscription)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27261,10 +26867,8 @@ func (m *awsRestjson1_deserializeOpSearch) HandleDeser
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27304,7 +26908,6 @@ func (m *awsRestjson1_deserializeOpSearch) HandleDeser
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27447,10 +27050,8 @@ func (m *awsRestjson1_deserializeOpSearchGroupProfiles
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27490,7 +27091,6 @@ func (m *awsRestjson1_deserializeOpSearchGroupProfiles
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27623,10 +27223,8 @@ func (m *awsRestjson1_deserializeOpSearchListings) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27666,7 +27264,6 @@ func (m *awsRestjson1_deserializeOpSearchListings) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27809,10 +27406,8 @@ func (m *awsRestjson1_deserializeOpSearchTypes) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -27852,7 +27447,6 @@ func (m *awsRestjson1_deserializeOpSearchTypes) Handle
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -27995,10 +27589,8 @@ func (m *awsRestjson1_deserializeOpSearchUserProfiles)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28038,7 +27630,6 @@ func (m *awsRestjson1_deserializeOpSearchUserProfiles)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28171,10 +27762,8 @@ func (m *awsRestjson1_deserializeOpStartDataSourceRun)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28214,7 +27803,6 @@ func (m *awsRestjson1_deserializeOpStartDataSourceRun)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28464,10 +28052,8 @@ func (m *awsRestjson1_deserializeOpStartMetadataGenera
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28507,7 +28093,6 @@ func (m *awsRestjson1_deserializeOpStartMetadataGenera
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28702,10 +28287,8 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28717,7 +28300,6 @@ func (m *awsRestjson1_deserializeOpTagResource) Handle
 	output := &TagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28805,10 +28387,8 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28820,7 +28400,6 @@ func (m *awsRestjson1_deserializeOpUntagResource) Hand
 	output := &UntagResourceOutput{}
 	out.Result = output
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -28905,10 +28484,8 @@ func (m *awsRestjson1_deserializeOpUpdateAssetFilter) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -28948,7 +28525,6 @@ func (m *awsRestjson1_deserializeOpUpdateAssetFilter) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29168,10 +28744,8 @@ func (m *awsRestjson1_deserializeOpUpdateConnection) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29211,7 +28785,6 @@ func (m *awsRestjson1_deserializeOpUpdateConnection) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29418,10 +28991,8 @@ func (m *awsRestjson1_deserializeOpUpdateDataSource) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29461,7 +29032,6 @@ func (m *awsRestjson1_deserializeOpUpdateDataSource) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -29777,10 +29347,8 @@ func (m *awsRestjson1_deserializeOpUpdateDomain) Handl
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -29820,7 +29388,6 @@ func (m *awsRestjson1_deserializeOpUpdateDomain) Handl
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30020,10 +29587,8 @@ func (m *awsRestjson1_deserializeOpUpdateDomainUnit) H
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30063,7 +29628,6 @@ func (m *awsRestjson1_deserializeOpUpdateDomainUnit) H
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30285,10 +29849,8 @@ func (m *awsRestjson1_deserializeOpUpdateEnvironment) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30328,7 +29890,6 @@ func (m *awsRestjson1_deserializeOpUpdateEnvironment) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30628,10 +30189,8 @@ func (m *awsRestjson1_deserializeOpUpdateEnvironmentAc
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30671,7 +30230,6 @@ func (m *awsRestjson1_deserializeOpUpdateEnvironmentAc
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -30843,10 +30401,8 @@ func (m *awsRestjson1_deserializeOpUpdateEnvironmentPr
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -30886,7 +30442,6 @@ func (m *awsRestjson1_deserializeOpUpdateEnvironmentPr
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31123,10 +30678,8 @@ func (m *awsRestjson1_deserializeOpUpdateGlossary) Han
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31166,7 +30719,6 @@ func (m *awsRestjson1_deserializeOpUpdateGlossary) Han
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31342,10 +30894,8 @@ func (m *awsRestjson1_deserializeOpUpdateGlossaryTerm)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31385,7 +30935,6 @@ func (m *awsRestjson1_deserializeOpUpdateGlossaryTerm)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31575,10 +31124,8 @@ func (m *awsRestjson1_deserializeOpUpdateGroupProfile)
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31618,7 +31165,6 @@ func (m *awsRestjson1_deserializeOpUpdateGroupProfile)
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -31773,10 +31319,8 @@ func (m *awsRestjson1_deserializeOpUpdateProject) Hand
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -31816,7 +31360,6 @@ func (m *awsRestjson1_deserializeOpUpdateProject) Hand
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32059,10 +31602,8 @@ func (m *awsRestjson1_deserializeOpUpdateProjectProfil
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32102,7 +31643,6 @@ func (m *awsRestjson1_deserializeOpUpdateProjectProfil
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32321,10 +31861,8 @@ func (m *awsRestjson1_deserializeOpUpdateRule) HandleD
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32364,7 +31902,6 @@ func (m *awsRestjson1_deserializeOpUpdateRule) HandleD
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32608,10 +32145,8 @@ func (m *awsRestjson1_deserializeOpUpdateSubscriptionG
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32651,7 +32186,6 @@ func (m *awsRestjson1_deserializeOpUpdateSubscriptionG
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -32878,10 +32412,8 @@ func (m *awsRestjson1_deserializeOpUpdateSubscriptionR
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -32921,7 +32453,6 @@ func (m *awsRestjson1_deserializeOpUpdateSubscriptionR
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33171,10 +32702,8 @@ func (m *awsRestjson1_deserializeOpUpdateSubscriptionT
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -33214,7 +32743,6 @@ func (m *awsRestjson1_deserializeOpUpdateSubscriptionT
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
@@ -33473,10 +33001,8 @@ func (m *awsRestjson1_deserializeOpUpdateUserProfile) 
 		return out, metadata, err
 	}
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 	response, ok := out.RawResponse.(*smithyhttp.Response)
 	if !ok {
 		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
@@ -33516,7 +33042,6 @@ func (m *awsRestjson1_deserializeOpUpdateUserProfile) 
 		}
 	}
 
-	span.End()
 	return out, metadata, err
 }
 
