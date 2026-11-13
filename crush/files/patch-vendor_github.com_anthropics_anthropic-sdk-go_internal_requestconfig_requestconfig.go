--- vendor/github.com/anthropics/anthropic-sdk-go/internal/requestconfig/requestconfig.go.orig	2026-10-15 19:41:18 UTC
+++ vendor/github.com/anthropics/anthropic-sdk-go/internal/requestconfig/requestconfig.go
@@ -20,7 +20,6 @@ import (
 	"github.com/anthropics/anthropic-sdk-go/internal/apierror"
 	"github.com/anthropics/anthropic-sdk-go/internal/apiform"
 	"github.com/anthropics/anthropic-sdk-go/internal/apiquery"
-	"github.com/anthropics/anthropic-sdk-go/internal/tracing"
 )
 
 func getDefaultHeaders() map[string]string {
@@ -80,8 +79,10 @@ type RequestOption interface {
 	Apply(*RequestConfig) error
 }
 
-type RequestOptionFunc func(*RequestConfig) error
-type PreRequestOptionFunc func(*RequestConfig) error
+type (
+	RequestOptionFunc    func(*RequestConfig) error
+	PreRequestOptionFunc func(*RequestConfig) error
+)
 
 func (s RequestOptionFunc) Apply(r *RequestConfig) error    { return s(r) }
 func (s PreRequestOptionFunc) Apply(r *RequestConfig) error { return s(r) }
@@ -234,6 +235,32 @@ type RequestConfig struct {
 	Body         io.Reader
 }
 
+type TracesConfig struct {
+	Enabled         *bool
+	ContentMode     string
+	MaxContentBytes *int64
+}
+
+type OpenTelemetryOptions struct {
+	ConversationID string
+}
+
+type OpenTelemetryConfig struct {
+	TracerProvider         any
+	Propagation            *bool
+	Traces                 TracesConfig
+	Platform               string
+	SpanQueryParams        bool
+	SpanContentQueryParams bool
+	Options                OpenTelemetryOptions
+}
+
+func (o OpenTelemetryConfig) clone() OpenTelemetryConfig { return o }
+
+type APICall struct {
+	Response *http.Response
+}
+
 // middleware is exactly the same type as the Middleware type found in the [option] package,
 // but it is redeclared here for circular dependency issues.
 type middleware = func(*http.Request, middlewareNext) (*http.Response, error)
@@ -503,7 +530,7 @@ func (cfg *RequestConfig) prepare() (err error) {
 
 // handler returns what sends one attempt: the HTTP client, wrapped in the middlewares.
 // The span of call, unless call is nil, records the request that reaches the client.
-func (cfg *RequestConfig) handler(call *tracing.APICall) middlewareNext {
+func (cfg *RequestConfig) handler(call *APICall) middlewareNext {
 	handler := cfg.HTTPClient.Do
 	if cfg.CustomHTTPDoer != nil {
 		handler = cfg.CustomHTTPDoer.Do
@@ -549,7 +576,7 @@ func (cfg *RequestConfig) Execute() (err error) {
 
 // send sends the request in ctx, with retries, and decodes the response into
 // cfg.ResponseBodyInto. It records each attempt on call, unless call is nil.
-func (cfg *RequestConfig) send(ctx context.Context, call *tracing.APICall) (err error) {
+func (cfg *RequestConfig) send(ctx context.Context, call *APICall) (err error) {
 	handler := cfg.handler(call)
 	if call != nil {
 		handler = tracedHandler(call, handler)
@@ -800,4 +827,48 @@ func WithDefaultBaseURL(baseURL string) RequestOption 
 		r.DefaultBaseURL = u
 		return nil
 	})
+}
+
+func (cfg *RequestConfig) sendInSpan() error {
+	return cfg.send(cfg.Context, nil)
+}
+
+func tracedDoer(call *APICall, spanQueryParams bool, spanContentQueryParams bool, next func(*http.Request) (*http.Response, error)) func(*http.Request) (*http.Response, error) {
+	_ = call
+	_ = spanQueryParams
+	_ = spanContentQueryParams
+	return next
+}
+
+func tracedHandler(call *APICall, next middlewareNext) middlewareNext {
+	_ = call
+	return next
+}
+
+func WithSpan(name string, path string, queryParams ...string) RequestOption {
+	_ = name
+	_ = path
+	_ = queryParams
+	return RequestOptionFunc(func(*RequestConfig) error { return nil })
+}
+
+func WithSpanAttribute(key string, value string) RequestOption {
+	_ = key
+	_ = value
+	return RequestOptionFunc(func(*RequestConfig) error { return nil })
+}
+
+func WithSpanContentQuery(queryParams ...string) RequestOption {
+	_ = queryParams
+	return RequestOptionFunc(func(r *RequestConfig) error {
+		r.OpenTelemetry.SpanContentQueryParams = true
+		return nil
+	})
+}
+
+func WithSpanCreatedID[T any](key string, created **T, getID func(*T) string) RequestOption {
+	_ = key
+	_ = created
+	_ = getID
+	return RequestOptionFunc(func(*RequestConfig) error { return nil })
 }
