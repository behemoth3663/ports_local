--- vendor/github.com/aws/aws-sdk-go-v2/service/drs/serializers.go.orig	2025-07-19 01:28:18 UTC
+++ vendor/github.com/aws/aws-sdk-go-v2/service/drs/serializers.go
@@ -11,7 +11,6 @@ import (
 	"github.com/aws/smithy-go/encoding/httpbinding"
 	smithyjson "github.com/aws/smithy-go/encoding/json"
 	"github.com/aws/smithy-go/middleware"
-	"github.com/aws/smithy-go/tracing"
 	smithyhttp "github.com/aws/smithy-go/transport/http"
 )
 
@@ -25,10 +24,8 @@ func (m *awsRestjson1_serializeOpAssociateSourceNetwor
 func (m *awsRestjson1_serializeOpAssociateSourceNetworkStack) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -73,7 +70,6 @@ func (m *awsRestjson1_serializeOpAssociateSourceNetwor
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsAssociateSourceNetworkStackInput(v *AssociateSourceNetworkStackInput, encoder *httpbinding.Encoder) error {
@@ -111,10 +107,8 @@ func (m *awsRestjson1_serializeOpCreateExtendedSourceS
 func (m *awsRestjson1_serializeOpCreateExtendedSourceServer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -159,7 +153,6 @@ func (m *awsRestjson1_serializeOpCreateExtendedSourceS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateExtendedSourceServerInput(v *CreateExtendedSourceServerInput, encoder *httpbinding.Encoder) error {
@@ -199,10 +192,8 @@ func (m *awsRestjson1_serializeOpCreateLaunchConfigura
 func (m *awsRestjson1_serializeOpCreateLaunchConfigurationTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -247,7 +238,6 @@ func (m *awsRestjson1_serializeOpCreateLaunchConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateLaunchConfigurationTemplateInput(v *CreateLaunchConfigurationTemplateInput, encoder *httpbinding.Encoder) error {
@@ -324,10 +314,8 @@ func (m *awsRestjson1_serializeOpCreateReplicationConf
 func (m *awsRestjson1_serializeOpCreateReplicationConfigurationTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -372,7 +360,6 @@ func (m *awsRestjson1_serializeOpCreateReplicationConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateReplicationConfigurationTemplateInput(v *CreateReplicationConfigurationTemplateInput, encoder *httpbinding.Encoder) error {
@@ -483,10 +470,8 @@ func (m *awsRestjson1_serializeOpCreateSourceNetwork) 
 func (m *awsRestjson1_serializeOpCreateSourceNetwork) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -531,7 +516,6 @@ func (m *awsRestjson1_serializeOpCreateSourceNetwork) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsCreateSourceNetworkInput(v *CreateSourceNetworkInput, encoder *httpbinding.Encoder) error {
@@ -581,10 +565,8 @@ func (m *awsRestjson1_serializeOpDeleteJob) HandleSeri
 func (m *awsRestjson1_serializeOpDeleteJob) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -629,7 +611,6 @@ func (m *awsRestjson1_serializeOpDeleteJob) HandleSeri
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteJobInput(v *DeleteJobInput, encoder *httpbinding.Encoder) error {
@@ -662,10 +643,8 @@ func (m *awsRestjson1_serializeOpDeleteLaunchAction) H
 func (m *awsRestjson1_serializeOpDeleteLaunchAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -710,7 +689,6 @@ func (m *awsRestjson1_serializeOpDeleteLaunchAction) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteLaunchActionInput(v *DeleteLaunchActionInput, encoder *httpbinding.Encoder) error {
@@ -748,10 +726,8 @@ func (m *awsRestjson1_serializeOpDeleteLaunchConfigura
 func (m *awsRestjson1_serializeOpDeleteLaunchConfigurationTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -796,7 +772,6 @@ func (m *awsRestjson1_serializeOpDeleteLaunchConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteLaunchConfigurationTemplateInput(v *DeleteLaunchConfigurationTemplateInput, encoder *httpbinding.Encoder) error {
@@ -829,10 +804,8 @@ func (m *awsRestjson1_serializeOpDeleteRecoveryInstanc
 func (m *awsRestjson1_serializeOpDeleteRecoveryInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -877,7 +850,6 @@ func (m *awsRestjson1_serializeOpDeleteRecoveryInstanc
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteRecoveryInstanceInput(v *DeleteRecoveryInstanceInput, encoder *httpbinding.Encoder) error {
@@ -910,10 +882,8 @@ func (m *awsRestjson1_serializeOpDeleteReplicationConf
 func (m *awsRestjson1_serializeOpDeleteReplicationConfigurationTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -958,7 +928,6 @@ func (m *awsRestjson1_serializeOpDeleteReplicationConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteReplicationConfigurationTemplateInput(v *DeleteReplicationConfigurationTemplateInput, encoder *httpbinding.Encoder) error {
@@ -991,10 +960,8 @@ func (m *awsRestjson1_serializeOpDeleteSourceNetwork) 
 func (m *awsRestjson1_serializeOpDeleteSourceNetwork) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1039,7 +1006,6 @@ func (m *awsRestjson1_serializeOpDeleteSourceNetwork) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSourceNetworkInput(v *DeleteSourceNetworkInput, encoder *httpbinding.Encoder) error {
@@ -1072,10 +1038,8 @@ func (m *awsRestjson1_serializeOpDeleteSourceServer) H
 func (m *awsRestjson1_serializeOpDeleteSourceServer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1120,7 +1084,6 @@ func (m *awsRestjson1_serializeOpDeleteSourceServer) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDeleteSourceServerInput(v *DeleteSourceServerInput, encoder *httpbinding.Encoder) error {
@@ -1153,10 +1116,8 @@ func (m *awsRestjson1_serializeOpDescribeJobLogItems) 
 func (m *awsRestjson1_serializeOpDescribeJobLogItems) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1201,7 +1162,6 @@ func (m *awsRestjson1_serializeOpDescribeJobLogItems) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeJobLogItemsInput(v *DescribeJobLogItemsInput, encoder *httpbinding.Encoder) error {
@@ -1244,10 +1204,8 @@ func (m *awsRestjson1_serializeOpDescribeJobs) HandleS
 func (m *awsRestjson1_serializeOpDescribeJobs) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1292,7 +1250,6 @@ func (m *awsRestjson1_serializeOpDescribeJobs) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeJobsInput(v *DescribeJobsInput, encoder *httpbinding.Encoder) error {
@@ -1337,10 +1294,8 @@ func (m *awsRestjson1_serializeOpDescribeLaunchConfigu
 func (m *awsRestjson1_serializeOpDescribeLaunchConfigurationTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1385,7 +1340,6 @@ func (m *awsRestjson1_serializeOpDescribeLaunchConfigu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeLaunchConfigurationTemplatesInput(v *DescribeLaunchConfigurationTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -1430,10 +1384,8 @@ func (m *awsRestjson1_serializeOpDescribeRecoveryInsta
 func (m *awsRestjson1_serializeOpDescribeRecoveryInstances) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1478,7 +1430,6 @@ func (m *awsRestjson1_serializeOpDescribeRecoveryInsta
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeRecoveryInstancesInput(v *DescribeRecoveryInstancesInput, encoder *httpbinding.Encoder) error {
@@ -1523,10 +1474,8 @@ func (m *awsRestjson1_serializeOpDescribeRecoverySnaps
 func (m *awsRestjson1_serializeOpDescribeRecoverySnapshots) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1571,7 +1520,6 @@ func (m *awsRestjson1_serializeOpDescribeRecoverySnaps
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeRecoverySnapshotsInput(v *DescribeRecoverySnapshotsInput, encoder *httpbinding.Encoder) error {
@@ -1626,10 +1574,8 @@ func (m *awsRestjson1_serializeOpDescribeReplicationCo
 func (m *awsRestjson1_serializeOpDescribeReplicationConfigurationTemplates) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1674,7 +1620,6 @@ func (m *awsRestjson1_serializeOpDescribeReplicationCo
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeReplicationConfigurationTemplatesInput(v *DescribeReplicationConfigurationTemplatesInput, encoder *httpbinding.Encoder) error {
@@ -1719,10 +1664,8 @@ func (m *awsRestjson1_serializeOpDescribeSourceNetwork
 func (m *awsRestjson1_serializeOpDescribeSourceNetworks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1767,7 +1710,6 @@ func (m *awsRestjson1_serializeOpDescribeSourceNetwork
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeSourceNetworksInput(v *DescribeSourceNetworksInput, encoder *httpbinding.Encoder) error {
@@ -1812,10 +1754,8 @@ func (m *awsRestjson1_serializeOpDescribeSourceServers
 func (m *awsRestjson1_serializeOpDescribeSourceServers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1860,7 +1800,6 @@ func (m *awsRestjson1_serializeOpDescribeSourceServers
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDescribeSourceServersInput(v *DescribeSourceServersInput, encoder *httpbinding.Encoder) error {
@@ -1905,10 +1844,8 @@ func (m *awsRestjson1_serializeOpDisconnectRecoveryIns
 func (m *awsRestjson1_serializeOpDisconnectRecoveryInstance) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -1953,7 +1890,6 @@ func (m *awsRestjson1_serializeOpDisconnectRecoveryIns
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisconnectRecoveryInstanceInput(v *DisconnectRecoveryInstanceInput, encoder *httpbinding.Encoder) error {
@@ -1986,10 +1922,8 @@ func (m *awsRestjson1_serializeOpDisconnectSourceServe
 func (m *awsRestjson1_serializeOpDisconnectSourceServer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2034,7 +1968,6 @@ func (m *awsRestjson1_serializeOpDisconnectSourceServe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsDisconnectSourceServerInput(v *DisconnectSourceServerInput, encoder *httpbinding.Encoder) error {
@@ -2067,10 +2000,8 @@ func (m *awsRestjson1_serializeOpExportSourceNetworkCf
 func (m *awsRestjson1_serializeOpExportSourceNetworkCfnTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2115,7 +2046,6 @@ func (m *awsRestjson1_serializeOpExportSourceNetworkCf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsExportSourceNetworkCfnTemplateInput(v *ExportSourceNetworkCfnTemplateInput, encoder *httpbinding.Encoder) error {
@@ -2148,10 +2078,8 @@ func (m *awsRestjson1_serializeOpGetFailbackReplicatio
 func (m *awsRestjson1_serializeOpGetFailbackReplicationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2196,7 +2124,6 @@ func (m *awsRestjson1_serializeOpGetFailbackReplicatio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetFailbackReplicationConfigurationInput(v *GetFailbackReplicationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2229,10 +2156,8 @@ func (m *awsRestjson1_serializeOpGetLaunchConfiguratio
 func (m *awsRestjson1_serializeOpGetLaunchConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2277,7 +2202,6 @@ func (m *awsRestjson1_serializeOpGetLaunchConfiguratio
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetLaunchConfigurationInput(v *GetLaunchConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2310,10 +2234,8 @@ func (m *awsRestjson1_serializeOpGetReplicationConfigu
 func (m *awsRestjson1_serializeOpGetReplicationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2358,7 +2280,6 @@ func (m *awsRestjson1_serializeOpGetReplicationConfigu
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsGetReplicationConfigurationInput(v *GetReplicationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -2391,10 +2312,8 @@ func (m *awsRestjson1_serializeOpInitializeService) Ha
 func (m *awsRestjson1_serializeOpInitializeService) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2428,7 +2347,6 @@ func (m *awsRestjson1_serializeOpInitializeService) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsInitializeServiceInput(v *InitializeServiceInput, encoder *httpbinding.Encoder) error {
@@ -2449,10 +2367,8 @@ func (m *awsRestjson1_serializeOpListExtensibleSourceS
 func (m *awsRestjson1_serializeOpListExtensibleSourceServers) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2497,7 +2413,6 @@ func (m *awsRestjson1_serializeOpListExtensibleSourceS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListExtensibleSourceServersInput(v *ListExtensibleSourceServersInput, encoder *httpbinding.Encoder) error {
@@ -2540,10 +2455,8 @@ func (m *awsRestjson1_serializeOpListLaunchActions) Ha
 func (m *awsRestjson1_serializeOpListLaunchActions) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2588,7 +2501,6 @@ func (m *awsRestjson1_serializeOpListLaunchActions) Ha
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListLaunchActionsInput(v *ListLaunchActionsInput, encoder *httpbinding.Encoder) error {
@@ -2638,10 +2550,8 @@ func (m *awsRestjson1_serializeOpListStagingAccounts) 
 func (m *awsRestjson1_serializeOpListStagingAccounts) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2679,7 +2589,6 @@ func (m *awsRestjson1_serializeOpListStagingAccounts) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListStagingAccountsInput(v *ListStagingAccountsInput, encoder *httpbinding.Encoder) error {
@@ -2708,10 +2617,8 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 func (m *awsRestjson1_serializeOpListTagsForResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2749,7 +2656,6 @@ func (m *awsRestjson1_serializeOpListTagsForResource) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsListTagsForResourceInput(v *ListTagsForResourceInput, encoder *httpbinding.Encoder) error {
@@ -2779,10 +2685,8 @@ func (m *awsRestjson1_serializeOpPutLaunchAction) Hand
 func (m *awsRestjson1_serializeOpPutLaunchAction) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2827,7 +2731,6 @@ func (m *awsRestjson1_serializeOpPutLaunchAction) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsPutLaunchActionInput(v *PutLaunchActionInput, encoder *httpbinding.Encoder) error {
@@ -2912,10 +2815,8 @@ func (m *awsRestjson1_serializeOpRetryDataReplication)
 func (m *awsRestjson1_serializeOpRetryDataReplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -2960,7 +2861,6 @@ func (m *awsRestjson1_serializeOpRetryDataReplication)
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsRetryDataReplicationInput(v *RetryDataReplicationInput, encoder *httpbinding.Encoder) error {
@@ -2993,10 +2893,8 @@ func (m *awsRestjson1_serializeOpReverseReplication) H
 func (m *awsRestjson1_serializeOpReverseReplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3041,7 +2939,6 @@ func (m *awsRestjson1_serializeOpReverseReplication) H
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsReverseReplicationInput(v *ReverseReplicationInput, encoder *httpbinding.Encoder) error {
@@ -3074,10 +2971,8 @@ func (m *awsRestjson1_serializeOpStartFailbackLaunch) 
 func (m *awsRestjson1_serializeOpStartFailbackLaunch) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3122,7 +3017,6 @@ func (m *awsRestjson1_serializeOpStartFailbackLaunch) 
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartFailbackLaunchInput(v *StartFailbackLaunchInput, encoder *httpbinding.Encoder) error {
@@ -3164,10 +3058,8 @@ func (m *awsRestjson1_serializeOpStartRecovery) Handle
 func (m *awsRestjson1_serializeOpStartRecovery) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3212,7 +3104,6 @@ func (m *awsRestjson1_serializeOpStartRecovery) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartRecoveryInput(v *StartRecoveryInput, encoder *httpbinding.Encoder) error {
@@ -3259,10 +3150,8 @@ func (m *awsRestjson1_serializeOpStartReplication) Han
 func (m *awsRestjson1_serializeOpStartReplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3307,7 +3196,6 @@ func (m *awsRestjson1_serializeOpStartReplication) Han
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartReplicationInput(v *StartReplicationInput, encoder *httpbinding.Encoder) error {
@@ -3340,10 +3228,8 @@ func (m *awsRestjson1_serializeOpStartSourceNetworkRec
 func (m *awsRestjson1_serializeOpStartSourceNetworkRecovery) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3388,7 +3274,6 @@ func (m *awsRestjson1_serializeOpStartSourceNetworkRec
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartSourceNetworkRecoveryInput(v *StartSourceNetworkRecoveryInput, encoder *httpbinding.Encoder) error {
@@ -3435,10 +3320,8 @@ func (m *awsRestjson1_serializeOpStartSourceNetworkRep
 func (m *awsRestjson1_serializeOpStartSourceNetworkReplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3483,7 +3366,6 @@ func (m *awsRestjson1_serializeOpStartSourceNetworkRep
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStartSourceNetworkReplicationInput(v *StartSourceNetworkReplicationInput, encoder *httpbinding.Encoder) error {
@@ -3516,10 +3398,8 @@ func (m *awsRestjson1_serializeOpStopFailback) HandleS
 func (m *awsRestjson1_serializeOpStopFailback) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3564,7 +3444,6 @@ func (m *awsRestjson1_serializeOpStopFailback) HandleS
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopFailbackInput(v *StopFailbackInput, encoder *httpbinding.Encoder) error {
@@ -3597,10 +3476,8 @@ func (m *awsRestjson1_serializeOpStopReplication) Hand
 func (m *awsRestjson1_serializeOpStopReplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3645,7 +3522,6 @@ func (m *awsRestjson1_serializeOpStopReplication) Hand
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopReplicationInput(v *StopReplicationInput, encoder *httpbinding.Encoder) error {
@@ -3678,10 +3554,8 @@ func (m *awsRestjson1_serializeOpStopSourceNetworkRepl
 func (m *awsRestjson1_serializeOpStopSourceNetworkReplication) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3726,7 +3600,6 @@ func (m *awsRestjson1_serializeOpStopSourceNetworkRepl
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsStopSourceNetworkReplicationInput(v *StopSourceNetworkReplicationInput, encoder *httpbinding.Encoder) error {
@@ -3759,10 +3632,8 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 func (m *awsRestjson1_serializeOpTagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3811,7 +3682,6 @@ func (m *awsRestjson1_serializeOpTagResource) HandleSe
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTagResourceInput(v *TagResourceInput, encoder *httpbinding.Encoder) error {
@@ -3855,10 +3725,8 @@ func (m *awsRestjson1_serializeOpTerminateRecoveryInst
 func (m *awsRestjson1_serializeOpTerminateRecoveryInstances) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3903,7 +3771,6 @@ func (m *awsRestjson1_serializeOpTerminateRecoveryInst
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsTerminateRecoveryInstancesInput(v *TerminateRecoveryInstancesInput, encoder *httpbinding.Encoder) error {
@@ -3938,10 +3805,8 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 func (m *awsRestjson1_serializeOpUntagResource) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -3979,7 +3844,6 @@ func (m *awsRestjson1_serializeOpUntagResource) Handle
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUntagResourceInput(v *UntagResourceInput, encoder *httpbinding.Encoder) error {
@@ -4015,10 +3879,8 @@ func (m *awsRestjson1_serializeOpUpdateFailbackReplica
 func (m *awsRestjson1_serializeOpUpdateFailbackReplicationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4063,7 +3925,6 @@ func (m *awsRestjson1_serializeOpUpdateFailbackReplica
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateFailbackReplicationConfigurationInput(v *UpdateFailbackReplicationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -4111,10 +3972,8 @@ func (m *awsRestjson1_serializeOpUpdateLaunchConfigura
 func (m *awsRestjson1_serializeOpUpdateLaunchConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4159,7 +4018,6 @@ func (m *awsRestjson1_serializeOpUpdateLaunchConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateLaunchConfigurationInput(v *UpdateLaunchConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -4236,10 +4094,8 @@ func (m *awsRestjson1_serializeOpUpdateLaunchConfigura
 func (m *awsRestjson1_serializeOpUpdateLaunchConfigurationTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4284,7 +4140,6 @@ func (m *awsRestjson1_serializeOpUpdateLaunchConfigura
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateLaunchConfigurationTemplateInput(v *UpdateLaunchConfigurationTemplateInput, encoder *httpbinding.Encoder) error {
@@ -4359,10 +4214,8 @@ func (m *awsRestjson1_serializeOpUpdateReplicationConf
 func (m *awsRestjson1_serializeOpUpdateReplicationConfiguration) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4407,7 +4260,6 @@ func (m *awsRestjson1_serializeOpUpdateReplicationConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateReplicationConfigurationInput(v *UpdateReplicationConfigurationInput, encoder *httpbinding.Encoder) error {
@@ -4528,10 +4380,8 @@ func (m *awsRestjson1_serializeOpUpdateReplicationConf
 func (m *awsRestjson1_serializeOpUpdateReplicationConfigurationTemplate) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
 	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
 ) {
-	_, span := tracing.StartSpan(ctx, "OperationSerializer")
 	endTimer := startMetricTimer(ctx, "client.call.serialization_duration")
 	defer endTimer()
-	defer span.End()
 	request, ok := in.Request.(*smithyhttp.Request)
 	if !ok {
 		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
@@ -4576,7 +4426,6 @@ func (m *awsRestjson1_serializeOpUpdateReplicationConf
 	in.Request = request
 
 	endTimer()
-	span.End()
 	return next.HandleSerialize(ctx, in)
 }
 func awsRestjson1_serializeOpHttpBindingsUpdateReplicationConfigurationTemplateInput(v *UpdateReplicationConfigurationTemplateInput, encoder *httpbinding.Encoder) error {
