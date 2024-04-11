--- internal/langserver/handlers/service.go.orig	2023-10-12 15:35:06 UTC
+++ internal/langserver/handlers/service.go
@@ -34,11 +34,6 @@ import (
 	"github.com/hashicorp/terraform-ls/internal/terraform/discovery"
 	"github.com/hashicorp/terraform-ls/internal/terraform/exec"
 	"github.com/hashicorp/terraform-ls/internal/walker"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
-	"go.opentelemetry.io/otel/trace"
 )
 
 type service struct {
@@ -600,17 +595,6 @@ const tracerName = "github.com/hashicorp/terraform-ls/
 
 // handle calls a jrpc2.Func compatible function
 func handle(ctx context.Context, req *jrpc2.Request, fn interface{}) (interface{}, error) {
-	attrs := []attribute.KeyValue{
-		{
-			Key:   semconv.RPCMethodKey,
-			Value: attribute.StringValue(req.Method()),
-		},
-		{
-			Key:   semconv.RPCJsonrpcRequestIDKey,
-			Value: attribute.StringValue(req.ID()),
-		},
-	}
-
 	// We could capture all parameters here but for now we just
 	// opportunistically track the most important ones only.
 	type t struct {
@@ -632,16 +616,6 @@ func handle(ctx context.Context, req *jrpc2.Request, f
 		uri = params.RootURI
 	}
 
-	attrs = append(attrs, attribute.KeyValue{
-		Key:   attribute.Key("URI"),
-		Value: attribute.StringValue(uri),
-	})
-
-	tracer := otel.Tracer(tracerName)
-	ctx, span := tracer.Start(ctx, "rpc:"+req.Method(),
-		trace.WithAttributes(attrs...))
-	defer span.End()
-
 	ctx = lsctx.WithDocumentContext(ctx, lsctx.Document{
 		Method:     req.Method(),
 		LanguageID: params.TextDocument.LanguageID,
@@ -651,12 +625,6 @@ func handle(ctx context.Context, req *jrpc2.Request, f
 	result, err := rpch.New(fn)(ctx, req)
 	if ctx.Err() != nil && errors.Is(ctx.Err(), context.Canceled) {
 		err = fmt.Errorf("%w: %s", requestCancelled.Err(), err)
-	}
-	if err != nil {
-		span.RecordError(err)
-		span.SetStatus(codes.Error, "request failed")
-	} else {
-		span.SetStatus(codes.Ok, "ok")
 	}
 
 	return result, err
