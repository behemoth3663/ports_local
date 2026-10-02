--- internal/runner/run/options.go.orig	1979-11-29 21:00:00 UTC
+++ internal/runner/run/options.go
@@ -3,12 +3,11 @@ import (
 import (
 	"context"
 	"encoding/json"
+	"errors"
 	"fmt"
 	"path/filepath"
 	"time"
 
-	"errors"
-
 	"github.com/gruntwork-io/terragrunt/internal/cloner"
 	"github.com/gruntwork-io/terragrunt/internal/engine"
 	"github.com/gruntwork-io/terragrunt/internal/errorconfig"
@@ -19,10 +18,10 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/report"
 	"github.com/gruntwork-io/terragrunt/internal/shell"
 	"github.com/gruntwork-io/terragrunt/internal/strict"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/tf"
 	"github.com/gruntwork-io/terragrunt/internal/tfimpl"
 	"github.com/gruntwork-io/terragrunt/internal/tflint"
+	telemetryopts "github.com/gruntwork-io/terragrunt/internal/traceopts"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"github.com/gruntwork-io/terragrunt/pkg/log/format/placeholders"
@@ -46,7 +45,7 @@ type Options struct {
 	EngineOptions                *engine.EngineOptions
 	Errors                       *errorconfig.Config
 	FeatureFlags                 map[string]string
-	Telemetry                    *telemetry.Options
+	Telemetry                    *telemetryopts.Options
 	SourceMap                    map[string]string
 	TFPath                       string
 	TerraformCommand             string
@@ -366,7 +365,7 @@ func (o *Options) handleIgnoreSignals(l log.Logger, fs
 		return err
 	}
 
-	const ownerPerms = 0644
+	const ownerPerms = 0o644
 
 	l.Warnf("Writing error signals to %s", signalsFile)
 
