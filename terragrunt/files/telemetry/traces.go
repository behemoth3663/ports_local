package telemetry

import (
	"context"

	"github.com/gruntwork-io/terragrunt/options"
)

// Trace - collect traces for method execution
func Trace(opts *options.TerragruntOptions, name string, attrs map[string]interface{}, fn func(childCtx context.Context) error) error {
	return nil
}
