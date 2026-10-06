--- vendor/github.com/elastic/elastic-transport-go/v8/elastictransport/instrumentation.go.orig	2026-10-06 09:46:09 UTC
+++ vendor/github.com/elastic/elastic-transport-go/v8/elastictransport/instrumentation.go
@@ -18,33 +18,11 @@ import (
 package elastictransport
 
 import (
-	"bytes"
 	"context"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 	"io"
 	"net/http"
-	"strconv"
 )
 
-const schemaUrl = "https://opentelemetry.io/schemas/1.21.0"
-const tracerName = "elasticsearch-api"
-
-// Constants for Semantic Convention
-// see https://opentelemetry.io/docs/specs/semconv/database/elasticsearch/ for details.
-const attrDbSystem = "db.system"
-const attrDbStatement = "db.statement"
-const attrDbOperation = "db.operation"
-const attrDbElasticsearchClusterName = "db.elasticsearch.cluster.name"
-const attrDbElasticsearchNodeName = "db.elasticsearch.node.name"
-const attrHttpRequestMethod = "http.request.method"
-const attrUrlFull = "url.full"
-const attrServerAddress = "server.address"
-const attrServerPort = "server.port"
-const attrPathParts = "db.elasticsearch.path_parts."
-
 // Instrumentation defines the interface the client uses to propagate information about the requests.
 // Each method is called with the current context or request for propagation.
 type Instrumentation interface {
@@ -74,142 +52,29 @@ type Instrumentation interface {
 	AfterResponse(ctx context.Context, res *http.Response)
 }
 
-type ElasticsearchOpenTelemetry struct {
-	tracer     trace.Tracer
-	recordBody bool
-}
+type ElasticsearchOpenTelemetry struct{}
 
-// NewOtelInstrumentation returns a new instrument for Open Telemetry traces
-// If no provider is passed, the instrumentation will fall back to the global otel provider.
-// captureSearchBody sets the query capture behavior for search endpoints.
-// version should be set to the version provided by the caller.
-func NewOtelInstrumentation(provider trace.TracerProvider, captureSearchBody bool, version string) *ElasticsearchOpenTelemetry {
-	if provider == nil {
-		provider = otel.GetTracerProvider()
-	}
-	return &ElasticsearchOpenTelemetry{
-		tracer: provider.Tracer(
-			tracerName,
-			trace.WithInstrumentationVersion(version),
-			trace.WithSchemaURL(schemaUrl),
-		),
-		recordBody: captureSearchBody,
-	}
+// NewOtelInstrumentation returns a no-op instrument for compatibility.
+func NewOtelInstrumentation(provider interface{}, captureSearchBody bool, version string) *ElasticsearchOpenTelemetry {
+	return &ElasticsearchOpenTelemetry{}
 }
 
-// Start begins a new span in the given context with the provided name.
-// Span will always have a kind set to trace.SpanKindClient.
-// The context span aware is returned for use within the client.
 func (i ElasticsearchOpenTelemetry) Start(ctx context.Context, name string) context.Context {
-	newCtx, _ := i.tracer.Start(ctx, name, trace.WithSpanKind(trace.SpanKindClient))
-	return newCtx
+	return ctx
 }
 
-// Close call for the end of the span, preferably defered by the client once started.
-func (i ElasticsearchOpenTelemetry) Close(ctx context.Context) {
-	span := trace.SpanFromContext(ctx)
-	if span.IsRecording() {
-		span.End()
-	}
-}
+func (i ElasticsearchOpenTelemetry) Close(ctx context.Context) {}
 
-// shouldRecordRequestBody filters for search endpoints.
-func (i ElasticsearchOpenTelemetry) shouldRecordRequestBody(endpoint string) bool {
-	// allow list of endpoints that will propagate query to OpenTelemetry.
-	// see https://opentelemetry.io/docs/specs/semconv/database/elasticsearch/#call-level-attributes
-	var searchEndpoints = map[string]struct{}{
-		"search":                 {},
-		"async_search.submit":    {},
-		"msearch":                {},
-		"eql.search":             {},
-		"terms_enum":             {},
-		"search_template":        {},
-		"msearch_template":       {},
-		"render_search_template": {},
-	}
-
-	if i.recordBody {
-		if _, ok := searchEndpoints[endpoint]; ok {
-			return true
-		}
-	}
-	return false
-}
-
-// RecordRequestBody add the db.statement attributes only for search endpoints.
-// Returns a new reader if the query has been recorded, nil otherwise.
 func (i ElasticsearchOpenTelemetry) RecordRequestBody(ctx context.Context, endpoint string, query io.Reader) io.ReadCloser {
-	if i.shouldRecordRequestBody(endpoint) == false {
-		return nil
-	}
-
-	span := trace.SpanFromContext(ctx)
-	if span.IsRecording() {
-		buf := bytes.Buffer{}
-		buf.ReadFrom(query)
-		span.SetAttributes(attribute.String(attrDbStatement, buf.String()))
-		getBody := func() (io.ReadCloser, error) {
-			reader := buf
-			return io.NopCloser(&reader), nil
-		}
-		reader, _ := getBody()
-		return reader
-	}
-
 	return nil
 }
 
-// RecordError sets any provided error as an OTel error in the active span.
-func (i ElasticsearchOpenTelemetry) RecordError(ctx context.Context, err error) {
-	span := trace.SpanFromContext(ctx)
-	if span.IsRecording() {
-		span.SetStatus(codes.Error, "an error happened while executing a request")
-		span.RecordError(err)
-	}
-}
+func (i ElasticsearchOpenTelemetry) RecordError(ctx context.Context, err error) {}
 
-// RecordPathPart sets the couple for a specific path part.
-// An index placed in the path would translate to `db.elasticsearch.path_parts.index`.
-func (i ElasticsearchOpenTelemetry) RecordPathPart(ctx context.Context, pathPart, value string) {
-	span := trace.SpanFromContext(ctx)
-	if span.IsRecording() {
-		span.SetAttributes(attribute.String(attrPathParts+pathPart, value))
-	}
-}
+func (i ElasticsearchOpenTelemetry) RecordPathPart(ctx context.Context, pathPart, value string) {}
 
-// BeforeRequest noop for interface.
 func (i ElasticsearchOpenTelemetry) BeforeRequest(req *http.Request, endpoint string) {}
 
-// AfterRequest enrich the span with the available data from the request.
-func (i ElasticsearchOpenTelemetry) AfterRequest(req *http.Request, system, endpoint string) {
-	span := trace.SpanFromContext(req.Context())
-	if span.IsRecording() {
-		span.SetAttributes(
-			attribute.String(attrDbSystem, system),
-			attribute.String(attrDbOperation, endpoint),
-			attribute.String(attrHttpRequestMethod, req.Method),
-			attribute.String(attrUrlFull, req.URL.String()),
-			attribute.String(attrServerAddress, req.URL.Hostname()),
-		)
-		if value, err := strconv.ParseInt(req.URL.Port(), 10, 32); err == nil {
-			span.SetAttributes(attribute.Int64(attrServerPort, value))
-		}
-	}
-}
+func (i ElasticsearchOpenTelemetry) AfterRequest(req *http.Request, system, endpoint string) {}
 
-// AfterResponse enric the span with the cluster id and node name if the query was executed on Elastic Cloud.
-func (i ElasticsearchOpenTelemetry) AfterResponse(ctx context.Context, res *http.Response) {
-	span := trace.SpanFromContext(ctx)
-	if span.IsRecording() {
-		if id := res.Header.Get("X-Found-Handling-Cluster"); id != "" {
-			span.SetAttributes(
-				attribute.String(attrDbElasticsearchClusterName, id),
-			)
-		}
-		if name := res.Header.Get("X-Found-Handling-Instance"); name != "" {
-			span.SetAttributes(
-				attribute.String(attrDbElasticsearchNodeName, name),
-			)
-		}
-	}
-}
+func (i ElasticsearchOpenTelemetry) AfterResponse(ctx context.Context, res *http.Response) {}
