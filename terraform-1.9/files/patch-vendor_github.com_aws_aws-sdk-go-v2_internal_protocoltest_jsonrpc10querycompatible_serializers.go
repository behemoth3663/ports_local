--- vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/jsonrpc10querycompatible/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/jsonrpc10querycompatible/serializers.go
@@ -10,7 +10,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"path"
 )
@@ -25,10 +24,8 @@ func (m *awsAwsjson10_serializeOpGetItem) HandleSerial
 func (m *awsAwsjson10_serializeOpGetItem) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -73,7 +70,6 @@ func (m *awsAwsjson10_serializeOpGetItem) HandleSerial
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsAwsjson10_serializeOpDocumentGetItemInput(v *GetItemInput, value smithyjson.Value) error {
