package telemetry

import (
	"context"

	"github.com/gruntwork-io/terragrunt/options"
)

// Time - collect time for function execution
func Time(opts *options.TerragruntOptions, name string, attrs map[string]interface{}, fn func(childCtx context.Context) error) error {
	return nil
}

// Count - add to counter provided value
func Count(opts *options.TerragruntOptions, name string, value int64) {
}
