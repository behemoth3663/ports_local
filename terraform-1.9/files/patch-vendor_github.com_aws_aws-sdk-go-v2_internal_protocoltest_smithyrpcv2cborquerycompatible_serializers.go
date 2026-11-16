--- vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/smithyrpcv2cborquerycompatible/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/smithyrpcv2cborquerycompatible/serializers.go
@@ -9,7 +9,6 @@ import (
 	smithy "github.com/aws/smithy-go"
 	smithycbor "github.com/aws/smithy-go/encoding/cbor"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net/http"
 )
@@ -24,10 +23,8 @@ func (m *smithyRpcv2cbor_serializeOpGetItem) HandleSer
 func (m *smithyRpcv2cbor_serializeOpGetItem) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*GetItemInput)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -60,7 +57,6 @@ func (m *smithyRpcv2cbor_serializeOpGetItem) HandleSer
 	in.Request = req
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
