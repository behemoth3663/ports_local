package telemetry

import (
	"context"
	"fmt"
	"io"

	"github.com/gruntwork-io/terragrunt/options"
)

// TelemetryOptions - options for telemetry provider.
type TelemetryOptions struct {
	Vars       map[string]string
	AppName    string
	AppVersion string
	Writer     io.Writer
	ErrWriter  io.Writer
}

// InitTelemetry - initialize the telemetry provider.
func InitTelemetry(ctx context.Context, opts *TelemetryOptions) error {
	return nil
}

// ShutdownTelemetry - shutdown the telemetry provider.
func ShutdownTelemetry(ctx context.Context) error {
	return nil
}

// Telemetry - collect telemetry from function execution - metrics and traces.
func Telemetry(opts *options.TerragruntOptions, name string, attrs map[string]interface{}, fn func(childCtx context.Context) error) error {
	return nil
}

// ErrorMissingEnvVariable error for missing environment variable.
type ErrorMissingEnvVariable struct {
	Vars []string
}

func (e *ErrorMissingEnvVariable) Error() string {
	return fmt.Sprintf("missing environment variable: %v", e.Vars)
}
