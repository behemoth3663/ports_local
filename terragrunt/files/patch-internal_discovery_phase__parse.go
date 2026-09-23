--- internal/discovery/phase_parse.go.orig	1979-11-29 21:00:00 UTC
+++ internal/discovery/phase_parse.go
@@ -2,6 +2,7 @@ import (
 
 import (
 	"context"
+	"errors"
 	"fmt"
 	"io"
 	"path/filepath"
@@ -9,13 +10,10 @@ import (
 	"strings"
 	"sync"
 
-	"errors"
-
 	"github.com/gruntwork-io/terragrunt/internal/component"
 	"github.com/gruntwork-io/terragrunt/internal/configbridge"
 	"github.com/gruntwork-io/terragrunt/internal/filter"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/internal/vfs"
 	"github.com/gruntwork-io/terragrunt/pkg/config"
@@ -117,7 +115,6 @@ func ensureParsed(
 		// Cache hits are emitted as a counter rather than a span. A graph or
 		// relationship traversal can re-visit the same component many times, and
 		// a span per visit drowns out the parse misses that are the actual cost.
-		telemetry.TelemeterFromContext(ctx).Count(ctx, "discovery_parse_cache_hit", 1)
 		l.Debugf(
 			"Discovery: parse cache hit for %s (phase=%s, depth=%d)",
 			c.Path(), parsePhaseFromContext(ctx), parseDepthFromContext(ctx),
@@ -320,132 +317,124 @@ func parseComponent(
 	phase := parsePhaseFromContext(ctx)
 	depth := parseDepthFromContext(ctx)
 
-	return telemetry.TelemeterFromContext(ctx).
-		Collect(ctx, l, "discovery_parse_component", map[string]any{
-			"path":      c.Path(),
-			"phase":     phase,
-			"depth":     depth,
-			"cache_hit": false,
-		}, func(ctx context.Context, l log.Logger) error {
-			l.Debugf("Discovery: parsing %s (phase=%s, depth=%d)", c.Path(), phase, depth)
+		l.Debugf("Discovery: parsing %s (phase=%s, depth=%d)", c.Path(), phase, depth)
 
-			parseOpts := opts.Clone()
+		parseOpts := opts.Clone()
 
-			componentPath := c.Path()
-			workingDir := componentPath
+		componentPath := c.Path()
+		workingDir := componentPath
 
-			if vfs.Exists(v.FS, componentPath) && !vfs.IsDir(v.FS, componentPath) {
-				workingDir = filepath.Dir(componentPath)
-			}
+		if vfs.Exists(v.FS, componentPath) && !vfs.IsDir(v.FS, componentPath) {
+			workingDir = filepath.Dir(componentPath)
+		}
 
-			configFilename := config.DefaultTerragruntConfigPath
+		configFilename := config.DefaultTerragruntConfigPath
 
-			switch c.(type) {
-			case *component.Stack:
-				configFilename = config.DefaultStackFile
-			default:
-				if unit, ok := c.(*component.Unit); ok && unit.ConfigFile() != "" {
-					configFilename = unit.ConfigFile()
-					break
-				}
+		switch c.(type) {
+		case *component.Stack:
+			configFilename = config.DefaultStackFile
+		default:
+			if unit, ok := c.(*component.Unit); ok && unit.ConfigFile() != "" {
+				configFilename = unit.ConfigFile()
+				break
+			}
 
-				if opts.TerragruntConfigPath != "" && !vfs.IsDir(v.FS, opts.TerragruntConfigPath) {
-					configFilename = filepath.Base(opts.TerragruntConfigPath)
-				}
+			if opts.TerragruntConfigPath != "" && !vfs.IsDir(v.FS, opts.TerragruntConfigPath) {
+				configFilename = filepath.Base(opts.TerragruntConfigPath)
 			}
+		}
 
-			parseOpts.WorkingDir = workingDir
-			parseOpts.SkipOutput = true
-			parseOpts.TerragruntConfigPath = filepath.Join(parseOpts.WorkingDir, configFilename)
-			parseOpts.OriginalTerragruntConfigPath = parseOpts.TerragruntConfigPath
+		parseOpts.WorkingDir = workingDir
+		parseOpts.SkipOutput = true
+		parseOpts.TerragruntConfigPath = filepath.Join(parseOpts.WorkingDir, configFilename)
+		parseOpts.OriginalTerragruntConfigPath = parseOpts.TerragruntConfigPath
 
-			// Clone v.Env so concurrent parseComponent goroutines launched by
-			// ParsePhase and RelationshipPhase don't race on the shared map when
-			// ObtainCredsForParsing writes auth-provider-cmd output into it.
-			parseV := v.WithEnvCloned().WithWriter(io.Discard).WithErrWriter(io.Discard)
+		// Clone v.Env so concurrent parseComponent goroutines launched by
+		// ParsePhase and RelationshipPhase don't race on the shared map when
+		// ObtainCredsForParsing writes auth-provider-cmd output into it.
+		parseV := v.WithEnvCloned().WithWriter(io.Discard).WithErrWriter(io.Discard)
 
-			shellOpts := configbridge.ShellRunOptsFromOpts(v.Env, parseOpts)
+		shellOpts := configbridge.ShellRunOptsFromOpts(v.Env, parseOpts)
 
-			if parseOpts.DiscoveryAuthProviderCmd {
-				if _, err := creds.ObtainCredsForParsing(
-					ctx, l, parseV, parseOpts.AuthProviderCmd, shellOpts,
-				); err != nil {
-					return fmt.Errorf(
-						"obtaining auth provider credentials for %s: %w",
-						parseOpts.TerragruntConfigPath,
-						err,
-					)
-				}
+		if parseOpts.DiscoveryAuthProviderCmd {
+			if _, err := creds.ObtainCredsForParsing(
+				ctx, l, parseV, parseOpts.AuthProviderCmd, shellOpts,
+			); err != nil {
+				return fmt.Errorf(
+					"obtaining auth provider credentials for %s: %w",
+					parseOpts.TerragruntConfigPath,
+					err,
+				)
 			}
+		}
 
-			ctx, parsingCtx := configbridge.NewParsingContext(ctx, l, parseV, parseOpts)
-			parsingCtx = parsingCtx.WithDecodeList(
-				config.TerraformSource,
-				config.DependenciesBlock,
-				config.DependencyBlock,
-				config.TerragruntFlags,
-				config.FeatureFlagsBlock,
-				config.ExcludeBlock,
-				config.ErrorsBlock,
-				config.RemoteStateBlock,
-				config.TerragruntVersionConstraints,
-			).WithSkipOutputsResolution()
+		ctx, parsingCtx := configbridge.NewParsingContext(ctx, l, parseV, parseOpts)
+		parsingCtx = parsingCtx.WithDecodeList(
+			config.TerraformSource,
+			config.DependenciesBlock,
+			config.DependencyBlock,
+			config.TerragruntFlags,
+			config.FeatureFlagsBlock,
+			config.ExcludeBlock,
+			config.ErrorsBlock,
+			config.RemoteStateBlock,
+			config.TerragruntVersionConstraints,
+		).WithSkipOutputsResolution()
 
-			if len(discovery.parserOptions) > 0 {
-				parsingCtx = parsingCtx.WithParseOption(discovery.parserOptions)
-			}
+		if len(discovery.parserOptions) > 0 {
+			parsingCtx = parsingCtx.WithParseOption(discovery.parserOptions)
+		}
 
-			if discovery.trackReads {
-				parsingCtx = parsingCtx.WithFileReadTracking()
-			}
+		if discovery.trackReads {
+			parsingCtx = parsingCtx.WithFileReadTracking()
+		}
 
+		if discovery.suppressParseErrors {
+			parserOpts := parsingCtx.ParserOptions
+			parserOpts = append(parserOpts, hclparse.WithDiagnosticsHandler(func(
+				file *hcl.File,
+				hclDiags hcl.Diagnostics,
+			) (hcl.Diagnostics, error) {
+				l.Debugf("Suppressed parsing errors %v", hclDiags)
+				return nil, nil
+			}))
+			parsingCtx = parsingCtx.WithParseOption(parserOpts)
+		}
+
+		cfg, err := config.PartialParseConfigFile(
+			ctx,
+			parsingCtx,
+			l,
+			parseOpts.TerragruntConfigPath,
+			nil,
+		)
+		if err != nil {
 			if discovery.suppressParseErrors {
-				parserOpts := parsingCtx.ParserOptions
-				parserOpts = append(parserOpts, hclparse.WithDiagnosticsHandler(func(
-					file *hcl.File,
-					hclDiags hcl.Diagnostics,
-				) (hcl.Diagnostics, error) {
-					l.Debugf("Suppressed parsing errors %v", hclDiags)
-					return nil, nil
-				}))
-				parsingCtx = parsingCtx.WithParseOption(parserOpts)
-			}
+				if _, ok := errors.AsType[config.TerragruntConfigNotFoundError](err); ok {
+					l.Debugf(
+						"Skipping missing config during discovery: %s",
+						parseOpts.TerragruntConfigPath,
+					)
 
-			cfg, err := config.PartialParseConfigFile(
-				ctx,
-				parsingCtx,
-				l,
-				parseOpts.TerragruntConfigPath,
-				nil,
-			)
-			if err != nil {
-				if discovery.suppressParseErrors {
-					if _, ok := errors.AsType[config.TerragruntConfigNotFoundError](err); ok {
-						l.Debugf(
-							"Skipping missing config during discovery: %s",
-							parseOpts.TerragruntConfigPath,
-						)
-
-						return nil
-					}
+					return nil
 				}
-
-				if !discovery.suppressParseErrors || cfg == nil {
-					return err
-				}
-
-				l.Debugf("Suppressing parse error for %s: %s", parseOpts.TerragruntConfigPath, err)
 			}
 
-			if unit, ok := c.(*component.Unit); ok {
-				unit.StoreConfig(cfg)
+			if !discovery.suppressParseErrors || cfg == nil {
+				return err
 			}
 
-			if parsingCtx.FilesRead.Tracking() {
-				readFiles := sanitizeReadFiles(parsingCtx.FilesRead.Paths())
-				c.SetReading(readFiles...)
-			}
+			l.Debugf("Suppressing parse error for %s: %s", parseOpts.TerragruntConfigPath, err)
+		}
 
-			return nil
-		})
+		if unit, ok := c.(*component.Unit); ok {
+			unit.StoreConfig(cfg)
+		}
+
+		if parsingCtx.FilesRead.Tracking() {
+			readFiles := sanitizeReadFiles(parsingCtx.FilesRead.Paths())
+			c.SetReading(readFiles...)
+		}
+
+		return nil
 }
