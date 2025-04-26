--- vendor/github.com/hashicorp/aws-sdk-go-base/v2/logger.go.orig	2025-04-04 12:31:04 UTC
+++ vendor/github.com/hashicorp/aws-sdk-go-base/v2/logger.go
@@ -4,31 +4,20 @@ import (
 package awsbase
 
 import (
-	"bufio"
-	"bytes"
 	"context"
 	"fmt"
-	"io"
 	"log"
 	"net/http"
-	"net/textproto"
 	"strings"
 	"time"
 
 	"github.com/aws/aws-sdk-go-v2/aws"
 	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
-	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
-	"github.com/aws/aws-sdk-go-v2/service/s3"
 	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
-	"github.com/aws/aws-sdk-go-v2/service/sqs"
 	smithylogging "github.com/aws/smithy-go/logging"
 	"github.com/aws/smithy-go/middleware"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 	"github.com/hashicorp/aws-sdk-go-base/v2/logging"
-	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/semconv/v1.17.0/httpconv"
-	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
 )
 
 type debugLogger struct {
@@ -47,7 +36,7 @@ func (l debugLogger) Logf(classification smithylogging
 		}
 	} else {
 		s = strings.ReplaceAll(s, "\r", "") // Works around https://github.com/jen20/teamcity-go-test/pull/2
-		log.Printf("[%s] missing_context: %s "+string(logging.AwsSdkKey)+"="+awsSdkGoV2Val, classification, s)
+		log.Printf("[%s] missing_context: %s tf_aws.sdk="+awsSdkGoV2Val, classification, s)
 	}
 }
 
@@ -59,10 +48,6 @@ const awsSdkGoV2Val = "aws-sdk-go-v2"
 
 const awsSdkGoV2Val = "aws-sdk-go-v2"
 
-func awsSDKv2Attr() attribute.KeyValue {
-	return logging.AwsSdkKey.String(awsSdkGoV2Val)
-}
-
 type logAttributeExtractor struct{}
 
 func (l *logAttributeExtractor) ID() string {
@@ -71,32 +56,6 @@ func (l *logAttributeExtractor) HandleInitialize(ctx c
 
 func (l *logAttributeExtractor) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
 	out middleware.InitializeOutput, metadata middleware.Metadata, err error) {
-	logger := logging.RetrieveLogger(ctx)
-
-	region := awsmiddleware.GetRegion(ctx)
-	serviceID := awsmiddleware.GetServiceID(ctx)
-
-	attributes := []attribute.KeyValue{
-		otelaws.SystemAttr(),
-		otelaws.ServiceAttr(serviceID),
-		otelaws.RegionAttr(region),
-		otelaws.OperationAttr(awsmiddleware.GetOperationName(ctx)),
-		awsSDKv2Attr(),
-	}
-
-	setters := map[string]otelaws.AttributeBuilder{
-		dynamodb.ServiceID: otelaws.DynamoDBAttributeBuilder,
-		s3.ServiceID:       s3AttributeBuilder,
-		sqs.ServiceID:      otelaws.SQSAttributeBuilder,
-	}
-	if setter, ok := setters[serviceID]; ok {
-		attributes = append(attributes, setter(ctx, in, out)...)
-	}
-
-	for _, attribute := range attributes {
-		ctx = logger.SetField(ctx, string(attribute.Key), attribute.Value.AsInterface())
-	}
-
 	return next.HandleInitialize(ctx, in)
 }
 
@@ -120,10 +79,10 @@ func (r *requestResponseLogger) HandleDeserialize(ctx 
 	region := awsmiddleware.GetRegion(ctx)
 
 	if signingRegion := awsmiddleware.GetSigningRegion(ctx); signingRegion != region { //nolint:staticcheck // Not retrievable elsewhere
-		ctx = logger.SetField(ctx, string(logging.SigningRegionKey), signingRegion)
+		ctx = logger.SetField(ctx, "tf_aws.signing_region", signingRegion)
 	}
 	if awsmiddleware.GetEndpointSource(ctx) == aws.EndpointSourceCustom {
-		ctx = logger.SetField(ctx, string(logging.CustomEndpointKey), true)
+		ctx = logger.SetField(ctx, "tf_aws.custom_endpoint", true)
 	}
 
 	smithyRequest, ok := in.Request.(*smithyhttp.Request)
@@ -168,25 +127,7 @@ func decomposeHTTPResponse(ctx context.Context, resp *
 }
 
 func decomposeHTTPResponse(ctx context.Context, resp *http.Response, elapsed time.Duration) (map[string]any, error) {
-	var attributes []attribute.KeyValue
-
-	attributes = append(attributes, attribute.Int64("http.duration", elapsed.Milliseconds()))
-
-	attributes = append(attributes, httpconv.ClientResponse(resp)...)
-
-	attributes = append(attributes, logging.DecomposeResponseHeaders(resp)...)
-
-	bodyLogger := responseBodyLogger(ctx)
-	err := bodyLogger.Log(ctx, resp, &attributes)
-	if err != nil {
-		return nil, err
-	}
-
-	result := make(map[string]any, len(attributes))
-	for _, attribute := range attributes {
-		result[string(attribute.Key)] = attribute.Value.AsInterface()
-	}
-
+	result := make(map[string]any, 0)
 	return result, nil
 }
 
@@ -203,108 +144,6 @@ type defaultResponseBodyLogger struct{}
 var _ logging.ResponseBodyLogger = &defaultResponseBodyLogger{}
 
 type defaultResponseBodyLogger struct{}
-
-func (l *defaultResponseBodyLogger) Log(ctx context.Context, resp *http.Response, attrs *[]attribute.KeyValue) error {
-	content, err := io.ReadAll(resp.Body)
-	if err != nil {
-		return err
-	}
-
-	// Restore the body reader
-	resp.Body = io.NopCloser(bytes.NewBuffer(content))
-
-	reader := textproto.NewReader(bufio.NewReader(bytes.NewReader(content)))
-
-	body, err := logging.ReadTruncatedBody(reader, logging.MaxResponseBodyLen)
-	if err != nil {
-		return err
-	}
-
-	*attrs = append(*attrs, attribute.String("http.response.body", body))
-
-	return nil
-}
-
-// May be contributed to go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws
-// See: https://github.com/open-telemetry/opentelemetry-go-contrib/issues/4321
-func s3AttributeBuilder(ctx context.Context, in middleware.InitializeInput, out middleware.InitializeOutput) []attribute.KeyValue {
-	s3Attributes := []attribute.KeyValue{}
-
-	switch v := in.Parameters.(type) {
-	case *s3.AbortMultipartUploadInput:
-		s3Attributes = append(s3Attributes, semconv.AWSS3Bucket(aws.ToString(v.Bucket)))
-
-		s3Attributes = append(s3Attributes, semconv.AWSS3Key(aws.ToString(v.Key)))
-
-		s3Attributes = append(s3Attributes, semconv.AWSS3UploadID(aws.ToString(v.UploadId)))
-
-	case *s3.CompleteMultipartUploadInput:
-		s3Attributes = append(s3Attributes, semconv.AWSS3Bucket(aws.ToString(v.Bucket)))
-
-		s3Attributes = append(s3Attributes, semconv.AWSS3Key(aws.ToString(v.Key)))
-
-		s3Attributes = append(s3Attributes, semconv.AWSS3UploadID(aws.ToString(v.UploadId)))
-
-	case *s3.CreateBucketInput:
-		s3Attributes = append(s3Attributes, semconv.AWSS3Bucket(aws.ToString(v.Bucket)))
-
-	case *s3.CreateMultipartUploadInput:
-		s3Attributes = append(s3Attributes, semconv.AWSS3Bucket(aws.ToString(v.Bucket)))
-
-		s3Attributes = append(s3Attributes, semconv.AWSS3Key(aws.ToString(v.Key)))
-
-	case *s3.DeleteBucketInput:
-		s3Attributes = append(s3Attributes, semconv.AWSS3Bucket(aws.ToString(v.Bucket)))
-
-	case *s3.DeleteObjectInput:
-		s3Attributes = append(s3Attributes, semconv.AWSS3Bucket(aws.ToString(v.Bucket)))
-
-		s3Attributes = append(s3Attributes, semconv.AWSS3Key(aws.ToString(v.Key)))
-
-	case *s3.DeleteObjectsInput:
-		s3Attributes = append(s3Attributes, semconv.AWSS3Bucket(aws.ToString(v.Bucket)))
-
-		s3Attributes = append(s3Attributes, semconv.AWSS3Delete(serializeDeleteShorthand(v.Delete)))
-
-	case *s3.GetObjectInput:
-		s3Attributes = append(s3Attributes, semconv.AWSS3Bucket(aws.ToString(v.Bucket)))
-
-		s3Attributes = append(s3Attributes, semconv.AWSS3Key(aws.ToString(v.Key)))
-
-	case *s3.HeadBucketInput:
-		s3Attributes = append(s3Attributes, semconv.AWSS3Bucket(aws.ToString(v.Bucket)))
-
-	case *s3.HeadObjectInput:
-		s3Attributes = append(s3Attributes, semconv.AWSS3Bucket(aws.ToString(v.Bucket)))
-
-		s3Attributes = append(s3Attributes, semconv.AWSS3Key(aws.ToString(v.Key)))
-
-	case *s3.ListBucketsInput:
-		// ListBucketsInput defines no attributes
-
-	case *s3.ListObjectsInput:
-		s3Attributes = append(s3Attributes, semconv.AWSS3Bucket(aws.ToString(v.Bucket)))
-
-	case *s3.ListObjectsV2Input:
-		s3Attributes = append(s3Attributes, semconv.AWSS3Bucket(aws.ToString(v.Bucket)))
-
-	case *s3.PutObjectInput:
-		s3Attributes = append(s3Attributes, semconv.AWSS3Bucket(aws.ToString(v.Bucket)))
-
-		s3Attributes = append(s3Attributes, semconv.AWSS3Key(aws.ToString(v.Key)))
-
-	case *s3.UploadPartInput:
-		s3Attributes = append(s3Attributes, semconv.AWSS3Bucket(aws.ToString(v.Bucket)))
-
-		s3Attributes = append(s3Attributes, semconv.AWSS3Key(aws.ToString(v.Key)))
-
-		s3Attributes = append(s3Attributes, semconv.AWSS3PartNumber(int(aws.ToInt32(v.PartNumber))))
-
-		s3Attributes = append(s3Attributes, semconv.AWSS3UploadID(aws.ToString(v.UploadId)))
-	}
-
-	return s3Attributes
-}
 
 func serializeDeleteShorthand(d *s3types.Delete) string {
 	var builder strings.Builder
