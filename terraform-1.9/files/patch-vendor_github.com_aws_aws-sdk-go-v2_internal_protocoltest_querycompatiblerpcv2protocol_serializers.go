--- vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/querycompatiblerpcv2protocol/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/querycompatiblerpcv2protocol/serializers.go
@@ -7,7 +7,6 @@ import (
 	"fmt"
 	smithycbor "github.com/aws/smithy-go/encoding/cbor"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"net/http"
 )
@@ -22,10 +21,8 @@ func (m *smithyRpcv2cbor_serializeOpQueryCompatibleOpe
 func (m *smithyRpcv2cbor_serializeOpQueryCompatibleOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	input, ok := in.Parameters.(*QueryCompatibleOperationInput)
 	if !ok {
 		return out, metadata, fmt.Errorf("unexpected input type %T", in.Parameters)
@@ -45,7 +42,6 @@ func (m *smithyRpcv2cbor_serializeOpQueryCompatibleOpe
 	req.Header.Set("X-Amzn-Query-Mode", "true")
 
 	endTimer()
-	span.End()
 
 	return next.HandleSerialize(ctx, in)
 }
