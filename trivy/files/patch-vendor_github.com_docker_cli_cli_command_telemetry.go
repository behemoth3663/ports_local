--- vendor/github.com/docker/cli/cli/command/telemetry.go.orig	2026-10-02 12:18:10 UTC
+++ vendor/github.com/docker/cli/cli/command/telemetry.go
@@ -3,241 +3,62 @@ import (
 import (
 	"context"
 	"os"
-	"path/filepath"
 	"strings"
-	"sync"
-	"time"
-
-	"github.com/google/uuid"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/metric"
-	otelsdk "go.opentelemetry.io/otel/sdk"
-	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
-	"go.opentelemetry.io/otel/sdk/metric/metricdata"
-	"go.opentelemetry.io/otel/sdk/resource"
-	sdktrace "go.opentelemetry.io/otel/sdk/trace"
-	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
-	"go.opentelemetry.io/otel/trace"
 )
 
-const exportTimeout = 50 * time.Millisecond
-
-// TracerProvider is an extension of the trace.TracerProvider interface for CLI programs.
+// TracerProvider is retained for compatibility.
 type TracerProvider interface {
-	trace.TracerProvider
 	ForceFlush(ctx context.Context) error
 	Shutdown(ctx context.Context) error
 }
 
-// MeterProvider is an extension of the metric.MeterProvider interface for CLI programs.
+// MeterProvider is retained for compatibility.
 type MeterProvider interface {
-	metric.MeterProvider
 	ForceFlush(ctx context.Context) error
 	Shutdown(ctx context.Context) error
 }
 
-// TelemetryClient provides the methods for using OTEL tracing or metrics.
-type TelemetryClient interface {
-	// Resource returns the OTEL Resource configured with this TelemetryClient.
-	// This resource may be created lazily, but the resource should be the same
-	// each time this function is invoked.
-	Resource() *resource.Resource
+// TelemetryClient provides methods for telemetry integration.
+type TelemetryClient interface{}
 
-	// TracerProvider returns the currently initialized TracerProvider. This TracerProvider will be configured
-	// with the default tracing components for a CLI program
-	TracerProvider() trace.TracerProvider
+type telemetryResource struct{}
 
-	// MeterProvider returns the currently initialized MeterProvider. This MeterProvider will be configured
-	// with the default metric components for a CLI program
-	MeterProvider() metric.MeterProvider
+func (cli *DockerCli) Resource() *telemetryResource {
+	_ = cli
+	return &telemetryResource{}
 }
 
-func (cli *DockerCli) Resource() *resource.Resource {
-	return cli.res.Get()
-}
-
-func (*DockerCli) TracerProvider() trace.TracerProvider {
-	return otel.GetTracerProvider()
-}
-
-func (*DockerCli) MeterProvider() metric.MeterProvider {
-	return otel.GetMeterProvider()
-}
-
-// WithResourceOptions configures additional options for the default resource. The default
-// resource will continue to include its default options.
-func WithResourceOptions(opts ...resource.Option) CLIOption {
+func WithResourceOptions(_ ...interface{}) CLIOption {
 	return func(cli *DockerCli) error {
-		cli.res.AppendOptions(opts...)
+		_ = cli
 		return nil
 	}
 }
 
-// WithResource overwrites the default resource and prevents its creation.
-func WithResource(res *resource.Resource) CLIOption {
+func WithResource(_ interface{}) CLIOption {
 	return func(cli *DockerCli) error {
-		cli.res.Set(res)
+		_ = cli
 		return nil
 	}
 }
 
-type telemetryResource struct {
-	res  *resource.Resource
-	opts []resource.Option
-	once sync.Once
-}
+func (r *telemetryResource) Set(_ interface{})              { _ = r }
+func (r *telemetryResource) Get() *telemetryResource        { return r }
+func (r *telemetryResource) AppendOptions(_ ...interface{}) {}
 
-func (r *telemetryResource) Set(res *resource.Resource) {
-	r.res = res
+func (cli *DockerCli) createGlobalMeterProvider(_ context.Context, _ ...interface{}) {
+	_ = cli
 }
 
-func (r *telemetryResource) Get() *resource.Resource {
-	r.once.Do(r.init)
-	return r.res
+func (cli *DockerCli) createGlobalTracerProvider(_ context.Context, _ ...interface{}) {
+	_ = cli
 }
 
-func (r *telemetryResource) init() {
-	if r.res != nil {
-		r.opts = nil
-		return
-	}
-
-	opts := append(defaultResourceOptions(), r.opts...)
-	res, err := resource.New(context.Background(), opts...)
-	if err != nil {
-		otel.Handle(err)
-	}
-	r.res = res
-
-	// Clear the resource options since they'll never be used again and to allow
-	// the garbage collector to retrieve that memory.
-	r.opts = nil
-}
-
-// createGlobalMeterProvider creates a new MeterProvider from the initialized DockerCli struct
-// with the given options and sets it as the global meter provider
-func (cli *DockerCli) createGlobalMeterProvider(ctx context.Context, opts ...sdkmetric.Option) {
-	allOpts := make([]sdkmetric.Option, 0, len(opts)+2)
-	allOpts = append(allOpts, sdkmetric.WithResource(cli.Resource()))
-	allOpts = append(allOpts, dockerMetricExporter(ctx, cli)...)
-	allOpts = append(allOpts, opts...)
-	mp := sdkmetric.NewMeterProvider(allOpts...)
-	otel.SetMeterProvider(mp)
-}
-
-// createGlobalTracerProvider creates a new TracerProvider from the initialized DockerCli struct
-// with the given options and sets it as the global tracer provider
-func (cli *DockerCli) createGlobalTracerProvider(ctx context.Context, opts ...sdktrace.TracerProviderOption) {
-	allOpts := make([]sdktrace.TracerProviderOption, 0, len(opts)+2)
-	allOpts = append(allOpts, sdktrace.WithResource(cli.Resource()))
-	allOpts = append(allOpts, dockerSpanExporter(ctx, cli)...)
-	allOpts = append(allOpts, opts...)
-	tp := sdktrace.NewTracerProvider(allOpts...)
-	otel.SetTracerProvider(tp)
-}
-
-func defaultResourceOptions() []resource.Option {
-	return []resource.Option{
-		resource.WithDetectors(serviceNameDetector{}),
-		resource.WithAttributes(
-			// Use a unique instance id so OTEL knows that each invocation
-			// of the CLI is its own instance. Without this, downstream
-			// OTEL processors may think the same process is restarting
-			// continuously.
-			semconv.ServiceInstanceID(uuid.NewString()),
-		),
-		resource.WithFromEnv(),
-		resource.WithDetectors(telemetrySDK{}),
-	}
-}
-
-func (r *telemetryResource) AppendOptions(opts ...resource.Option) {
-	if r.res != nil {
-		return
-	}
-	r.opts = append(r.opts, opts...)
-}
-
-type (
-	serviceNameDetector struct{}
-	telemetrySDK        struct{}
+const (
+	resourceAttributesEnvVar = "OTEL_RESOURCE_ATTRIBUTES"
+	dockerCLIAttributePrefix = "docker.cli."
 )
 
-func (serviceNameDetector) Detect(ctx context.Context) (*resource.Resource, error) {
-	return resource.StringDetector(
-		semconv.SchemaURL,
-		semconv.ServiceNameKey,
-		func() (string, error) {
-			return filepath.Base(os.Args[0]), nil
-		},
-	).Detect(ctx)
-}
-
-// Detect returns a *Resource that describes the OpenTelemetry SDK used.
-func (telemetrySDK) Detect(context.Context) (*resource.Resource, error) {
-	return resource.NewWithAttributes(
-		semconv.SchemaURL,
-		semconv.TelemetrySDKName("opentelemetry"),
-		semconv.TelemetrySDKLanguageGo,
-		semconv.TelemetrySDKVersion(otelsdk.Version()),
-	), nil
-}
-
-// cliReader is an implementation of Reader that will automatically
-// report to a designated Exporter when Shutdown is called.
-type cliReader struct {
-	sdkmetric.Reader
-	exporter sdkmetric.Exporter
-}
-
-func newCLIReader(exp sdkmetric.Exporter) sdkmetric.Reader {
-	reader := sdkmetric.NewManualReader(
-		sdkmetric.WithTemporalitySelector(deltaTemporality),
-	)
-	return &cliReader{
-		Reader:   reader,
-		exporter: exp,
-	}
-}
-
-func (r *cliReader) Shutdown(ctx context.Context) error {
-	// Place a pretty tight constraint on the actual reporting.
-	// We don't want CLI metrics to prevent the CLI from exiting
-	// so if there's some kind of issue we need to abort pretty
-	// quickly.
-	ctx, cancel := context.WithTimeout(ctx, exportTimeout)
-	defer cancel()
-
-	return r.ForceFlush(ctx)
-}
-
-func (r *cliReader) ForceFlush(ctx context.Context) error {
-	var rm metricdata.ResourceMetrics
-	if err := r.Reader.Collect(ctx, &rm); err != nil {
-		return err
-	}
-
-	return r.exporter.Export(ctx, &rm)
-}
-
-// deltaTemporality sets the Temporality of every instrument to delta.
-//
-// This isn't really needed since we create a unique resource on each invocation,
-// but it can help with cardinality concerns for downstream processors since they can
-// perform aggregation for a time interval and then discard the data once that time
-// period has passed. Cumulative temporality would imply to the downstream processor
-// that they might receive a successive point and they may unnecessarily keep state
-// they really shouldn't.
-func deltaTemporality(_ sdkmetric.InstrumentKind) metricdata.Temporality {
-	return metricdata.DeltaTemporality
-}
-
-// resourceAttributesEnvVar is the name of the envvar that includes additional
-// resource attributes for OTEL as defined in the [OpenTelemetry specification].
-//
-// [OpenTelemetry specification]: https://opentelemetry.io/docs/specs/otel/configuration/sdk-environment-variables/#general-sdk-configuration
-const resourceAttributesEnvVar = "OTEL_RESOURCE_ATTRIBUTES"
-
 func filterResourceAttributesEnvvar() {
 	if v := os.Getenv(resourceAttributesEnvVar); v != "" {
 		if filtered := filterResourceAttributes(v); filtered != "" {
@@ -248,12 +69,6 @@ func filterResourceAttributesEnvvar() {
 	}
 }
 
-// dockerCLIAttributePrefix is the prefix for any docker cli OTEL attributes.
-// When updating, make sure to also update the copy in cli-plugins/manager.
-//
-// TODO(thaJeztah): move telemetry-related code to an (internal) package to reduce dependency on cli/command in cli-plugins, which has too many imports.
-const dockerCLIAttributePrefix = "docker.cli."
-
 func filterResourceAttributes(s string) string {
 	if trimmed := strings.TrimSpace(s); trimmed == "" {
 		return trimmed
@@ -264,12 +79,9 @@ func filterResourceAttributes(s string) string {
 	for _, p := range pairs {
 		k, _, found := strings.Cut(p, "=")
 		if !found {
-			// Do not interact with invalid otel resources.
 			elems = append(elems, p)
 			continue
 		}
-
-		// Skip attributes that have our docker.cli prefix.
 		if strings.HasPrefix(k, dockerCLIAttributePrefix) {
 			continue
 		}
