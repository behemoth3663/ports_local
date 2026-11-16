--- vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/nonquerycompatiblerpcv2protocol/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/nonquerycompatiblerpcv2protocol/deserializers.go
@@ -9,7 +9,6 @@ import (
 	smithycbor "github.com/aws/smithy-go/encoding/cbor"
 	"github.com/aws/smithy-go/middleware"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -37,10 +36,8 @@ func (m *smithyRpcv2cbor_deserializeOpNonQueryCompatib
 ) {
 	out, metadata, err = next.HandleDeserialize(ctx, in)
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if err != nil {
 		return out, metadata, err
