--- vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/smithyrpcv2cborquerycompatible/deserializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/internal/protocoltest/smithyrpcv2cborquerycompatible/deserializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/middleware"
 	"github.com/aws/smithy-go/ptr"
 	smithytime "github.com/aws/smithy-go/time"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"io"
 	"io/ioutil"
@@ -39,10 +38,8 @@ func (m *smithyRpcv2cbor_deserializeOpGetItem) HandleD
 ) {
 	out, metadata, err = next.HandleDeserialize(ctx, in)
 
-	_, span := tracing.StartSpan(ctx, "OperationDeserializer")
 	endTimer := startMetricTimer(ctx, "client.call.deserialization_duration")
 	defer endTimer()
-	defer span.End()
 
 	if err != nil {
 		return out, metadata, err
