--- vendor/cloud.google.com/go/storage/storage.go.orig	2025-05-30 06:48:19 UTC
+++ vendor/cloud.google.com/go/storage/storage.go
@@ -40,21 +40,15 @@ import (
 
 	"cloud.google.com/go/auth"
 	"cloud.google.com/go/internal/optional"
-	"cloud.google.com/go/internal/trace"
 	"cloud.google.com/go/storage/internal"
 	"cloud.google.com/go/storage/internal/apiv2/storagepb"
 	"github.com/googleapis/gax-go/v2"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/sdk/metric"
-	"go.opentelemetry.io/otel/sdk/metric/metricdata"
 	"google.golang.org/api/googleapi"
 	"google.golang.org/api/option"
 	"google.golang.org/api/option/internaloption"
 	raw "google.golang.org/api/storage/v1"
 	htransport "google.golang.org/api/transport/http"
 	"google.golang.org/grpc/codes"
-	"google.golang.org/grpc/experimental/stats"
-	"google.golang.org/grpc/stats/opentelemetry"
 	"google.golang.org/grpc/status"
 	"google.golang.org/protobuf/proto"
 	"google.golang.org/protobuf/reflect/protoreflect"
@@ -266,23 +260,7 @@ func CheckDirectConnectivitySupported(ctx context.Cont
 //
 // You can pass in [option.ClientOption] you plan on passing to [NewGRPCClient]
 func CheckDirectConnectivitySupported(ctx context.Context, bucket string, opts ...option.ClientOption) error {
-	view := metric.NewView(
-		metric.Instrument{
-			Name: "grpc.client.attempt.duration",
-			Kind: metric.InstrumentKindHistogram,
-		},
-		metric.Stream{AttributeFilter: attribute.NewAllowKeysFilter("grpc.lb.locality")},
-	)
-	mr := metric.NewManualReader()
-	provider := metric.NewMeterProvider(metric.WithReader(mr), metric.WithView(view))
-	// Provider handles shutting down ManualReader
-	defer provider.Shutdown(ctx)
-	mo := opentelemetry.MetricsOptions{
-		MeterProvider:  provider,
-		Metrics:        stats.NewMetrics("grpc.client.attempt.duration"),
-		OptionalLabels: []string{"grpc.lb.locality"},
-	}
-	combinedOpts := append(opts, WithDisabledClientMetrics(), option.WithGRPCDialOption(opentelemetry.DialOption(opentelemetry.Options{MetricsOptions: mo})))
+	combinedOpts := append(opts, WithDisabledClientMetrics())
 	client, err := NewGRPCClient(ctx, combinedOpts...)
 	if err != nil {
 		return fmt.Errorf("storage.NewGRPCClient: %w", err)
@@ -291,25 +269,7 @@ func CheckDirectConnectivitySupported(ctx context.Cont
 	if _, err = client.Bucket(bucket).Attrs(ctx); err != nil {
 		return fmt.Errorf("Bucket.Attrs: %w", err)
 	}
-	// Call manual reader to collect metric
-	rm := metricdata.ResourceMetrics{}
-	if err = mr.Collect(context.Background(), &rm); err != nil {
-		return fmt.Errorf("ManualReader.Collect: %w", err)
-	}
-	for _, sm := range rm.ScopeMetrics {
-		for _, m := range sm.Metrics {
-			if m.Name == "grpc.client.attempt.duration" {
-				hist := m.Data.(metricdata.Histogram[float64])
-				for _, d := range hist.DataPoints {
-					v, present := d.Attributes.Value("grpc.lb.locality")
-					if present && v.AsString() != "" && v.AsString() != "{}" {
 						return nil
-					}
-				}
-			}
-		}
-	}
-	return errors.New("storage: direct connectivity not detected")
 }
 
 // Close closes the Client.
@@ -1037,8 +997,6 @@ func (o *ObjectHandle) Attrs(ctx context.Context) (att
 // Attrs returns meta information about the object.
 // ErrObjectNotExist will be returned if the object is not found.
 func (o *ObjectHandle) Attrs(ctx context.Context) (attrs *ObjectAttrs, err error) {
-	ctx, _ = startSpan(ctx, "Object.Attrs")
-	defer func() { endSpan(ctx, err) }()
 
 	if err := o.validate(); err != nil {
 		return nil, err
@@ -1051,8 +1009,6 @@ func (o *ObjectHandle) Update(ctx context.Context, uat
 // ObjectAttrsToUpdate docs for details on treatment of zero values.
 // ErrObjectNotExist will be returned if the object is not found.
 func (o *ObjectHandle) Update(ctx context.Context, uattrs ObjectAttrsToUpdate) (oa *ObjectAttrs, err error) {
-	ctx, _ = startSpan(ctx, "Object.Update")
-	defer func() { endSpan(ctx, err) }()
 
 	if err := o.validate(); err != nil {
 		return nil, err
@@ -1122,8 +1078,6 @@ func (o *ObjectHandle) Delete(ctx context.Context) (er
 
 // Delete deletes the single specified object.
 func (o *ObjectHandle) Delete(ctx context.Context) (err error) {
-	ctx, _ = startSpan(ctx, "Object.Delete")
-	defer func() { endSpan(ctx, err) }()
 	if err := o.validate(); err != nil {
 		return err
 	}
@@ -1247,7 +1201,6 @@ func (o *ObjectHandle) NewWriter(ctx context.Context) 
 // It is the caller's responsibility to call Close when writing is done. To
 // stop writing without saving the data, cancel the context.
 func (o *ObjectHandle) NewWriter(ctx context.Context) *Writer {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.Object.Writer")
 	return &Writer{
 		ctx:         ctx,
 		o:           o,
@@ -1283,7 +1236,6 @@ func (o *ObjectHandle) NewWriterFromAppendableObject(c
 // objects which were created append semantics and not finalized.
 // This feature is in preview and is not yet available for general use.
 func (o *ObjectHandle) NewWriterFromAppendableObject(ctx context.Context, opts *AppendableWriterOpts) (*Writer, int64, error) {
-	ctx = trace.StartSpan(ctx, "cloud.google.com/go/storage.Object.Writer")
 	if o.gen < 0 {
 		return nil, 0, errors.New("storage: ObjectHandle.Generation must be set to use NewWriterFromAppendableObject")
 	}
