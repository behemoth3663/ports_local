--- vendor/github.com/go-openapi/runtime/client/opentracing.go.orig	2023-12-09 18:15:16 UTC
+++ vendor/github.com/go-openapi/runtime/client/opentracing.go
@@ -2,12 +2,8 @@ import (
 
 import (
 	"fmt"
-	"net/http"
 
 	"github.com/go-openapi/strfmt"
-	"github.com/opentracing/opentracing-go"
-	"github.com/opentracing/opentracing-go/ext"
-	"github.com/opentracing/opentracing-go/log"
 
 	"github.com/go-openapi/runtime"
 )
@@ -15,18 +11,8 @@ type tracingTransport struct {
 type tracingTransport struct {
 	transport runtime.ClientTransport
 	host      string
-	opts      []opentracing.StartSpanOption
 }
 
-func newOpenTracingTransport(transport runtime.ClientTransport, host string, opts []opentracing.StartSpanOption,
-) runtime.ClientTransport {
-	return &tracingTransport{
-		transport: transport,
-		host:      host,
-		opts:      opts,
-	}
-}
-
 func (t *tracingTransport) Submit(op *runtime.ClientOperation) (interface{}, error) {
 	if op.Context == nil {
 		return t.transport.Submit(op)
@@ -35,60 +21,16 @@ func (t *tracingTransport) Submit(op *runtime.ClientOp
 	params := op.Params
 	reader := op.Reader
 
-	var span opentracing.Span
-	defer func() {
-		if span != nil {
-			span.Finish()
-		}
-	}()
-
 	op.Params = runtime.ClientRequestWriterFunc(func(req runtime.ClientRequest, reg strfmt.Registry) error {
-		span = createClientSpan(op, req.GetHeaderParams(), t.host, t.opts)
 		return params.WriteToRequest(req, reg)
 	})
 
 	op.Reader = runtime.ClientResponseReaderFunc(func(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
-		if span != nil {
-			code := response.Code()
-			ext.HTTPStatusCode.Set(span, uint16(code))
-			if code >= 400 {
-				ext.Error.Set(span, true)
-			}
-		}
 		return reader.ReadResponse(response, consumer)
 	})
 
 	submit, err := t.transport.Submit(op)
-	if err != nil && span != nil {
-		ext.Error.Set(span, true)
-		span.LogFields(log.Error(err))
-	}
 	return submit, err
-}
-
-func createClientSpan(op *runtime.ClientOperation, header http.Header, host string,
-	opts []opentracing.StartSpanOption) opentracing.Span {
-	ctx := op.Context
-	span := opentracing.SpanFromContext(ctx)
-
-	if span != nil {
-		opts = append(opts, ext.SpanKindRPCClient)
-		span, _ = opentracing.StartSpanFromContextWithTracer(
-			ctx, span.Tracer(), operationName(op), opts...)
-
-		ext.Component.Set(span, "go-openapi")
-		ext.PeerHostname.Set(span, host)
-		span.SetTag("http.path", op.PathPattern)
-		ext.HTTPMethod.Set(span, op.Method)
-
-		_ = span.Tracer().Inject(
-			span.Context(),
-			opentracing.HTTPHeaders,
-			opentracing.HTTPHeadersCarrier(header))
-
-		return span
-	}
-	return nil
 }
 
 func operationName(op *runtime.ClientOperation) string {
