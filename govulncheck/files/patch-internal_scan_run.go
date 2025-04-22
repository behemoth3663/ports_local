--- internal/scan/run.go.orig	2025-01-06 19:26:26 UTC
+++ internal/scan/run.go
@@ -15,7 +15,6 @@ import (
 	"strings"
 	"time"
 
-	"golang.org/x/telemetry/counter"
 	"golang.org/x/vuln/internal/client"
 	"golang.org/x/vuln/internal/govulncheck"
 	"golang.org/x/vuln/internal/openvex"
@@ -55,8 +54,6 @@ func RunGovulncheck(ctx context.Context, env []string,
 		return err
 	}
 
-	incTelemetryFlagCounters(cfg)
-
 	switch cfg.ScanMode {
 	case govulncheck.ScanModeSource:
 		dir := filepath.FromSlash(cfg.dir)
@@ -137,19 +134,6 @@ func scannerVersion(cfg *config, bi *debug.BuildInfo) 
 		}
 	}
 	cfg.ScannerVersion = buf.String()
-}
-
-func incTelemetryFlagCounters(cfg *config) {
-	counter.Inc(fmt.Sprintf("govulncheck/mode:%s", cfg.ScanMode))
-	counter.Inc(fmt.Sprintf("govulncheck/scan:%s", cfg.ScanLevel))
-	counter.Inc(fmt.Sprintf("govulncheck/format:%s", cfg.format))
-
-	if len(cfg.show) == 0 {
-		counter.Inc("govulncheck/show:none")
-	}
-	for _, s := range cfg.show {
-		counter.Inc(fmt.Sprintf("govulncheck/show:%s", s))
-	}
 }
 
 func Flush(h govulncheck.Handler) error {
