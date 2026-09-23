--- vendor/cloud.google.com/go/longrunning/autogen/operations_client.go.orig	2026-09-23 21:57:03 UTC
+++ vendor/cloud.google.com/go/longrunning/autogen/operations_client.go
@@ -28,7 +28,6 @@ import (
 
 	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
 	gax "github.com/googleapis/gax-go/v2"
-	"github.com/googleapis/gax-go/v2/callctx"
 	"google.golang.org/api/iterator"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
@@ -307,16 +306,7 @@ func NewOperationsClient(ctx context.Context, opts ...
 // developers can have a consistent client experience.
 func NewOperationsClient(ctx context.Context, opts ...option.ClientOption) (*OperationsClient, error) {
 	clientOpts := defaultOperationsGRPCClientOptions()
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		clientOpts = append(clientOpts, internaloption.WithTelemetryAttributes(map[string]string{
-			"gcp.client.service":  "longrunning",
-			"gcp.client.version":  getVersionClient(),
-			"gcp.client.repo":     "googleapis/google-cloud-go",
-			"gcp.client.artifact": "cloud.google.com/go/longrunning/autogen",
-			"gcp.client.language": "go",
-			"url.domain":          "longrunning.googleapis.com",
-		}))
-	}
+
 	if newOperationsClientHook != nil {
 		hookOpts, err := newOperationsClientHook(ctx, clientHookParams{})
 		if err != nil {
@@ -338,25 +328,7 @@ func NewOperationsClient(ctx context.Context, opts ...
 		logger:           internaloption.GetLogger(opts),
 	}
 	c.setGoogleClientInfo()
-	if gax.IsFeatureEnabled("METRICS") {
-		metrics := gax.NewClientMetrics(
-			gax.WithTelemetryLogger(c.logger),
-			gax.WithTelemetryAttributes(map[string]string{
-				gax.ClientService:  "longrunning",
-				gax.ClientVersion:  getVersionClient(),
-				gax.ClientArtifact: "cloud.google.com/go/longrunning/autogen",
-				gax.RPCSystem:      "grpc",
-				gax.URLDomain:      "longrunning.googleapis.com",
-			}),
-		)
 
-		client.CallOptions.ListOperations = append(client.CallOptions.ListOperations, gax.WithClientMetrics(metrics))
-		client.CallOptions.GetOperation = append(client.CallOptions.GetOperation, gax.WithClientMetrics(metrics))
-		client.CallOptions.DeleteOperation = append(client.CallOptions.DeleteOperation, gax.WithClientMetrics(metrics))
-		client.CallOptions.CancelOperation = append(client.CallOptions.CancelOperation, gax.WithClientMetrics(metrics))
-		client.CallOptions.WaitOperation = append(client.CallOptions.WaitOperation, gax.WithClientMetrics(metrics))
-	}
-
 	client.internalClient = c
 
 	return &client, nil
@@ -417,16 +389,7 @@ func NewOperationsRESTClient(ctx context.Context, opts
 // developers can have a consistent client experience.
 func NewOperationsRESTClient(ctx context.Context, opts ...option.ClientOption) (*OperationsClient, error) {
 	clientOpts := append(defaultOperationsRESTClientOptions(), opts...)
-	if gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		clientOpts = append(clientOpts, internaloption.WithTelemetryAttributes(map[string]string{
-			"gcp.client.service":  "longrunning",
-			"gcp.client.version":  getVersionClient(),
-			"gcp.client.repo":     "googleapis/google-cloud-go",
-			"gcp.client.artifact": "cloud.google.com/go/longrunning/autogen",
-			"gcp.client.language": "go",
-			"url.domain":          "longrunning.googleapis.com",
-		}))
-	}
+
 	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
 	if err != nil {
 		return nil, err
@@ -441,25 +404,6 @@ func NewOperationsRESTClient(ctx context.Context, opts
 	}
 	c.setGoogleClientInfo()
 
-	if gax.IsFeatureEnabled("METRICS") {
-		metrics := gax.NewClientMetrics(
-			gax.WithTelemetryLogger(c.logger),
-			gax.WithTelemetryAttributes(map[string]string{
-				gax.ClientService:  "longrunning",
-				gax.ClientVersion:  getVersionClient(),
-				gax.ClientArtifact: "cloud.google.com/go/longrunning/autogen",
-				gax.RPCSystem:      "http",
-				gax.URLDomain:      "longrunning.googleapis.com",
-			}),
-		)
-
-		callOpts.ListOperations = append(callOpts.ListOperations, gax.WithClientMetrics(metrics))
-		callOpts.GetOperation = append(callOpts.GetOperation, gax.WithClientMetrics(metrics))
-		callOpts.DeleteOperation = append(callOpts.DeleteOperation, gax.WithClientMetrics(metrics))
-		callOpts.CancelOperation = append(callOpts.CancelOperation, gax.WithClientMetrics(metrics))
-		callOpts.WaitOperation = append(callOpts.WaitOperation, gax.WithClientMetrics(metrics))
-	}
-
 	return &OperationsClient{internalClient: c, CallOptions: callOpts}, nil
 }
 
@@ -505,9 +449,7 @@ func (c *operationsGRPCClient) ListOperations(ctx cont
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.longrunning.Operations/ListOperations")
-	}
+
 	opts = append((*c.CallOptions).ListOperations[0:len((*c.CallOptions).ListOperations):len((*c.CallOptions).ListOperations)], opts...)
 	it := &OperationIterator{}
 	req = proto.CloneOf(req)
@@ -554,9 +496,7 @@ func (c *operationsGRPCClient) GetOperation(ctx contex
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.longrunning.Operations/GetOperation")
-	}
+
 	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
 	var resp *longrunningpb.Operation
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -575,9 +515,7 @@ func (c *operationsGRPCClient) DeleteOperation(ctx con
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.longrunning.Operations/DeleteOperation")
-	}
+
 	opts = append((*c.CallOptions).DeleteOperation[0:len((*c.CallOptions).DeleteOperation):len((*c.CallOptions).DeleteOperation)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		var err error
@@ -592,9 +530,7 @@ func (c *operationsGRPCClient) CancelOperation(ctx con
 
 	hds = append(c.xGoogHeaders, hds...)
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.longrunning.Operations/CancelOperation")
-	}
+
 	opts = append((*c.CallOptions).CancelOperation[0:len((*c.CallOptions).CancelOperation):len((*c.CallOptions).CancelOperation)], opts...)
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		var err error
@@ -606,9 +542,7 @@ func (c *operationsGRPCClient) WaitOperation(ctx conte
 
 func (c *operationsGRPCClient) WaitOperation(ctx context.Context, req *longrunningpb.WaitOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
 	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, c.xGoogHeaders...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.longrunning.Operations/WaitOperation")
-	}
+
 	opts = append((*c.CallOptions).WaitOperation[0:len((*c.CallOptions).WaitOperation):len((*c.CallOptions).WaitOperation)], opts...)
 	var resp *longrunningpb.Operation
 	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
@@ -722,10 +656,7 @@ func (c *operationsRESTClient) GetOperation(ctx contex
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.longrunning.Operations/GetOperation")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=operations/**}")
-	}
+
 	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &longrunningpb.Operation{}
@@ -774,10 +705,7 @@ func (c *operationsRESTClient) DeleteOperation(ctx con
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.longrunning.Operations/DeleteOperation")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=operations/**}")
-	}
+
 	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		if settings.Path != "" {
 			baseUrl.Path = settings.Path
@@ -824,10 +752,7 @@ func (c *operationsRESTClient) CancelOperation(ctx con
 	hds = append(c.xGoogHeaders, hds...)
 	hds = append(hds, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.longrunning.Operations/CancelOperation")
-		ctx = callctx.WithTelemetryContext(ctx, "url_template", "/v1/{name=operations/**}:cancel")
-	}
+
 	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
 		if settings.Path != "" {
 			baseUrl.Path = settings.Path
@@ -877,9 +802,7 @@ func (c *operationsRESTClient) WaitOperation(ctx conte
 	// Build HTTP headers from client and context metadata.
 	hds := append(c.xGoogHeaders, "Content-Type", "application/json")
 	headers := gax.BuildHeaders(ctx, hds...)
-	if gax.IsFeatureEnabled("METRICS") || gax.IsFeatureEnabled("TRACING") || gax.IsFeatureEnabled("LOGGING") {
-		ctx = callctx.WithTelemetryContext(ctx, "rpc_method", "google.longrunning.Operations/WaitOperation")
-	}
+
 	opts = append((*c.CallOptions).WaitOperation[0:len((*c.CallOptions).WaitOperation):len((*c.CallOptions).WaitOperation)], opts...)
 	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
 	resp := &longrunningpb.Operation{}
