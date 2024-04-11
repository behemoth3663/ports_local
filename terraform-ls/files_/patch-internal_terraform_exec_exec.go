--- internal/terraform/exec/exec.go.orig	2023-10-12 15:35:06 UTC
+++ internal/terraform/exec/exec.go
@@ -15,10 +15,6 @@ import (
 	"github.com/hashicorp/terraform-exec/tfexec"
 	tfjson "github.com/hashicorp/terraform-json"
 	"github.com/hashicorp/terraform-ls/internal/logging"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 )
 
 var defaultExecTimeout = 30 * time.Second
@@ -77,15 +73,6 @@ func (e *Executor) contextfulError(ctx context.Context
 	return e.enrichCtxErr(method, err)
 }
 
-func (e *Executor) setSpanStatus(span trace.Span, err error) {
-	if err != nil {
-		span.RecordError(err)
-		span.SetStatus(codes.Error, "execution returned error")
-		return
-	}
-	span.SetStatus(codes.Ok, "execution successful")
-}
-
 func (e *Executor) enrichCtxErr(method string, err error) error {
 	if err == nil {
 		return nil
@@ -114,11 +101,8 @@ func (e *Executor) Init(ctx context.Context, opts ...t
 	if err != nil {
 		return err
 	}
-	ctx, span := otel.Tracer(tracerName).Start(ctx, "terraform-exec:Init")
-	defer span.End()
 
 	err = e.tf.Init(ctx, opts...)
-	e.setSpanStatus(span, err)
 
 	return e.contextfulError(ctx, "Init", err)
 }
@@ -130,11 +114,8 @@ func (e *Executor) Get(ctx context.Context, opts ...tf
 	if err != nil {
 		return err
 	}
-	ctx, span := otel.Tracer(tracerName).Start(ctx, "terraform-exec:Get")
-	defer span.End()
 
 	err = e.tf.Get(ctx, opts...)
-	e.setSpanStatus(span, err)
 
 	return e.contextfulError(ctx, "Get", err)
 }
@@ -147,18 +128,10 @@ func (e *Executor) Format(ctx context.Context, input [
 		return nil, err
 	}
 
-	ctx, span := otel.Tracer(tracerName).Start(ctx, "terraform-exec:Format",
-		trace.WithAttributes(attribute.KeyValue{
-			Key:   attribute.Key("stdinLength"),
-			Value: attribute.IntValue(len(input)),
-		}))
-	defer span.End()
-
 	br := bytes.NewReader(input)
 	buf := bytes.NewBuffer([]byte{})
 
 	err = e.tf.Format(ctx, br, buf)
-	e.setSpanStatus(span, err)
 
 	return buf.Bytes(), e.contextfulError(ctx, "Format", err)
 }
@@ -171,11 +144,7 @@ func (e *Executor) Validate(ctx context.Context) ([]tf
 		return []tfjson.Diagnostic{}, err
 	}
 
-	ctx, span := otel.Tracer(tracerName).Start(ctx, "terraform-exec:Validate")
-	defer span.End()
-
 	validation, err := e.tf.Validate(ctx)
-	e.setSpanStatus(span, err)
 	if err != nil {
 		return []tfjson.Diagnostic{}, e.contextfulError(ctx, "Validate", err)
 	}
@@ -191,11 +160,7 @@ func (e *Executor) Version(ctx context.Context) (*vers
 		return nil, nil, err
 	}
 
-	ctx, span := otel.Tracer(tracerName).Start(ctx, "terraform-exec:Version")
-	defer span.End()
-
 	ver, pv, err := e.tf.Version(ctx, true)
-	e.setSpanStatus(span, err)
 
 	return ver, pv, e.contextfulError(ctx, "Version", err)
 }
@@ -208,11 +173,7 @@ func (e *Executor) ProviderSchemas(ctx context.Context
 		return nil, err
 	}
 
-	ctx, span := otel.Tracer(tracerName).Start(ctx, "terraform-exec:ProviderSchemas")
-	defer span.End()
-
 	ps, err := e.tf.ProvidersSchema(ctx)
-	e.setSpanStatus(span, err)
 
 	return ps, e.contextfulError(ctx, "ProviderSchemas", err)
 }
