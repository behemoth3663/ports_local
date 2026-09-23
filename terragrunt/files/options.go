package traceopts

// Options keeps CLI-parsed telemetry-related flags as inert configuration data.
type Options struct {
	TraceExporter                  string
	TraceExporterHTTPEndpoint      string
	TraceParent                    string
	MetricExporter                 string
	LogsExporter                   string
	TraceExporterInsecureEndpoint  bool
	MetricExporterInsecureEndpoint bool
	LogsExporterInsecureEndpoint   bool
}
