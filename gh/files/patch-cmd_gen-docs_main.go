--- cmd/gen-docs/main.go.orig	2026-10-26 18:02:47 UTC
+++ cmd/gen-docs/main.go
@@ -10,8 +10,8 @@ import (
 	"github.com/cli/cli/v2/internal/config"
 	"github.com/cli/cli/v2/internal/docs"
 	"github.com/cli/cli/v2/internal/gh"
 	"github.com/cli/cli/v2/internal/ghrepo"
-	"github.com/cli/cli/v2/internal/telemetry"
+	"github.com/cli/cli/v2/internal/gh/ghtelemetry"
 	"github.com/cli/cli/v2/pkg/cmd/root"
 	"github.com/cli/cli/v2/pkg/cmdutil"
 	"github.com/cli/cli/v2/pkg/extensions"
@@ -54,7 +54,7 @@ func run(args []string) error {
 			return config.NewMockConfigFromString(""), nil
 		},
 		ExtensionManager: &em{},
-	}, &telemetry.NoOpService{}, "", "")
+	}, &ghtelemetry.NoOpService{}, "", "")
 	rootCmd.InitDefaultHelpCmd()
 
 	if err := os.MkdirAll(*dir, 0755); err != nil {
