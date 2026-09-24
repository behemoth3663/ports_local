--- internal/policy/client.go.orig	2026-09-23 11:57:56 UTC
+++ internal/policy/client.go
@@ -16,10 +16,6 @@ import (
 	"github.com/hashicorp/go-hclog"
 	"github.com/hashicorp/go-plugin"
 	"github.com/zclconf/go-cty/cty"
-	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 	"google.golang.org/grpc"
 
 	"github.com/hashicorp/terraform/internal/policy/callback"
@@ -116,9 +112,6 @@ func Connect(ctx context.Context, policyPluginPath str
 // Connect creates a connection to tfpolicy-plugin. If policyPluginPath is empty, the command lookup
 // will default to the executable "tfpolicy-plugin" in the $PATH.
 func Connect(ctx context.Context, policyPluginPath string) (Client, error) {
-	ctx, span := tracer().Start(ctx, "policy.client.connect")
-	defer span.End()
-
 	pgm := "tfpolicy-plugin" // by default, just use this if it's in the path
 
 	if policyPluginPath != "" {
@@ -142,7 +135,6 @@ func Connect(ctx context.Context, policyPluginPath str
 		// This propagates the active trace context via gRPC metadata  so the plugin's handler
 		// spans become children of the terraform-side span that invoked them.
 		GRPCDialOptions: []grpc.DialOption{
-			grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
 		},
 		Logger: hclog.New(&hclog.LoggerOptions{
 			Level: func() hclog.Level {
@@ -159,8 +151,6 @@ func Connect(ctx context.Context, policyPluginPath str
 	if err != nil {
 		plugin.Kill()
 		err = fmt.Errorf("failed to connect to plugin: %v", err)
-		span.RecordError(err)
-		span.SetStatus(codes.Error, err.Error())
 		return nil, err
 	}
 
@@ -168,8 +158,6 @@ func Connect(ctx context.Context, policyPluginPath str
 	if err != nil {
 		plugin.Kill()
 		err = fmt.Errorf("failed to dispense plugin: %v", err)
-		span.RecordError(err)
-		span.SetStatus(codes.Error, err.Error())
 		return nil, err
 	}
 
@@ -206,7 +194,6 @@ func (c *client) RegisterCallbackService(ctx context.C
 		// trace context the plugin client injects into outgoing callback RPC
 		// metadata is extracted and made the parent of the callback handler
 		// spans.
-		opts = append(opts, grpc.StatsHandler(otelgrpc.NewServerHandler()))
 		server := grpc.NewServer(opts...)
 		proto.RegisterCallbackServiceServer(server, c.cbServer)
 		serverCh <- server
@@ -240,9 +227,6 @@ func (c *client) Setup(ctx context.Context, req SetupR
 }
 
 func (c *client) Setup(ctx context.Context, req SetupRequest) SetupResponse {
-	ctx, span := tracer().Start(ctx, "policy.client.setup")
-	defer span.End()
-
 	log.Printf("[DEBUG] Setting up Terraform Policy connection")
 	protoReq := &proto.PolicySetupRequest{
 		ClientCapabilities: new(proto.PolicySetupRequest_ClientCapabilities),
@@ -273,11 +257,6 @@ func (c *client) EvaluateResource(ctx context.Context,
 }
 
 func (c *client) EvaluateResource(ctx context.Context, req EvaluationRequest[*proto.PolicyEvaluateResourceRequest_ResourceMetadata]) EvaluationResponse {
-	ctx, span := tracer().Start(ctx, "policy.client.evaluate_resource",
-		trace.WithAttributes(attribute.String("policy.resource.type", req.Target)),
-	)
-	defer span.End()
-
 	log.Printf("[DEBUG] Evaluating policy for resource %s", req.Target)
 	var diags []*proto.Diagnostic
 
@@ -330,11 +309,6 @@ func (c *client) EvaluateProvider(ctx context.Context,
 }
 
 func (c *client) EvaluateProvider(ctx context.Context, req EvaluationRequest[*proto.PolicyEvaluateProviderRequest_ProviderMetadata]) EvaluationResponse {
-	ctx, span := tracer().Start(ctx, "policy.client.evaluate_provider",
-		trace.WithAttributes(attribute.String("policy.provider.type", req.Target)),
-	)
-	defer span.End()
-
 	log.Printf("[DEBUG] Evaluating policy for provider %s", req.Target)
 	var diags []*proto.Diagnostic
 	req = normalizeRequest(req)
@@ -367,11 +341,6 @@ func (c *client) EvaluateModule(ctx context.Context, r
 }
 
 func (c *client) EvaluateModule(ctx context.Context, req EvaluationRequest[*proto.PolicyEvaluateModuleRequest_ModuleMetadata]) EvaluationResponse {
-	ctx, span := tracer().Start(ctx, "policy.client.evaluate_module",
-		trace.WithAttributes(attribute.String("policy.module.source", req.Target)),
-	)
-	defer span.End()
-
 	log.Printf("[DEBUG] Evaluating policy for module %s", req.Target)
 	var diags []*proto.Diagnostic
 
