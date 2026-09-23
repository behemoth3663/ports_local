diff --git a/vendor/cloud.google.com/go/storage/storage.go b/vendor/cloud.google.com/go/storage/storage.go
index cb4d36802..426ea1492 100644
--- vendor/cloud.google.com/go/storage/storage.go.orig
+++ vendor/cloud.google.com/go/storage/storage.go
@@ -43,17 +43,12 @@ import (
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
@@ -294,50 +289,10 @@ func NewGRPCClient(ctx context.Context, opts ...option.ClientOption) (*Client, e
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
-	client, err := NewGRPCClient(ctx, combinedOpts...)
-	if err != nil {
-		return fmt.Errorf("storage.NewGRPCClient: %w", err)
-	}
-	defer client.Close()
-	if _, err = client.Bucket(bucket).Attrs(ctx); err != nil {
-		return fmt.Errorf("Bucket.Attrs: %w", err)
-	}
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
-						return nil
-					}
-				}
-			}
-		}
-	}
-	return errors.New("storage: direct connectivity not detected")
+	_ = ctx
+	_ = bucket
+	_ = opts
+	return nil
 }
 
 // Close closes the Client.
