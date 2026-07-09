--- internal/server/general.go.orig	2026-07-07 18:53:34 UTC
+++ internal/server/general.go
@@ -21,7 +21,6 @@ import (
 	"strings"
 	"sync"
 
-	"golang.org/x/telemetry/counter"
 	"golang.org/x/tools/gopls/internal/cache"
 	"golang.org/x/tools/gopls/internal/debug"
 	debuglog "golang.org/x/tools/gopls/internal/debug/log"
@@ -31,7 +30,6 @@ import (
 	"golang.org/x/tools/gopls/internal/protocol"
 	"golang.org/x/tools/gopls/internal/protocol/semtok"
 	"golang.org/x/tools/gopls/internal/settings"
-	"golang.org/x/tools/gopls/internal/telemetry"
 	"golang.org/x/tools/gopls/internal/util/bug"
 	"golang.org/x/tools/gopls/internal/util/goversion"
 	"golang.org/x/tools/gopls/internal/util/moremaps"
@@ -44,12 +42,6 @@ func (s *server) Initialize(ctx context.Context, param
 	ctx, done := event.Start(ctx, "server.Initialize")
 	defer done()
 
-	var clientName string
-	if params != nil && params.ClientInfo != nil {
-		clientName = params.ClientInfo.Name
-	}
-	recordClientInfo(clientName)
-
 	s.stateMu.Lock()
 	if s.state >= serverInitializing {
 		defer s.stateMu.Unlock()
@@ -336,9 +328,6 @@ func (s *server) checkViewGoVersions() {
 		if oldestVersion == -1 || viewVersion < oldestVersion {
 			oldestVersion, fromBuild = viewVersion, false
 		}
-		if viewVersion >= 0 {
-			counter.Inc(fmt.Sprintf("gopls/goversion:1.%d", viewVersion))
-		}
 	}
 
 	if msg, isError := goversion.Message(oldestVersion, fromBuild); msg != "" {
@@ -649,18 +638,6 @@ func (s *server) newFolder(ctx context.Context, folder
 		return nil, err
 	}
 
-	// Increment folder counters.
-	switch {
-	case env.GOTOOLCHAIN == "auto" || strings.Contains(env.GOTOOLCHAIN, "+auto"):
-		counter.Inc("gopls/gotoolchain:auto")
-	case env.GOTOOLCHAIN == "path" || strings.Contains(env.GOTOOLCHAIN, "+path"):
-		counter.Inc("gopls/gotoolchain:path")
-	case env.GOTOOLCHAIN == "local": // local+auto and local+path handled above
-		counter.Inc("gopls/gotoolchain:local")
-	default:
-		counter.Inc("gopls/gotoolchain:other")
-	}
-
 	// Record whether a driver is in use so that it appears in the
 	// user's telemetry upload. Although we can't correlate the
 	// driver information with the crash or bug.Report at the
@@ -668,9 +645,6 @@ func (s *server) newFolder(ctx context.Context, folder
 	// driver tend to do so most of the time, so we'll get a
 	// strong clue. See #60890 for an example of an issue where
 	// this information would have been helpful.
-	if env.EffectiveGOPACKAGESDRIVER != "" {
-		counter.Inc("gopls/gopackagesdriver")
-	}
 
 	return &cache.Folder{
 		Dir:     folder,
@@ -722,10 +696,9 @@ func (s *server) eventuallyShowMessage(ctx context.Con
 	s.notifications = append(s.notifications, msg)
 }
 
-func (s *server) handleOptionResult(ctx context.Context, applied []telemetry.CounterPath, optionErrors []error) {
+func (s *server) handleOptionResult(ctx context.Context, applied [][]string, optionErrors []error) {
 	for _, path := range applied {
 		path = append(settings.CounterPath{"gopls", "setting"}, path...)
-		counter.Inc(path.FullName())
 	}
 
 	var warnings, errs []string
@@ -809,66 +782,4 @@ func (s *server) Exit(ctx context.Context) error {
 	// We don't terminate the process on a normal exit, we just allow it to
 	// close naturally if needed after the connection is closed.
 	return nil
-}
-
-// recordClientInfo records gopls client info.
-func recordClientInfo(clientName string) {
-	// This table maps LSP (not MCP) clientInfo.Name prefixes to Go telemetry counters.
-	// Where authoritative source is available, we link to it.
-	for _, cli := range [...]struct {
-		clientNamePrefix, telemetryKey string
-	}{
-		{"Visual Studio Code - Insiders", "gopls/client:vscode-insiders"},
-		{"Visual Studio Code", "gopls/client:vscode"},
-
-		{"VSCodium", "gopls/client:vscodium"},
-
-		// https://github.com/coder/code-server/blob/3cb92edc76ecc2cfa5809205897d93d4379b16a6/ci/build/build-vscode.sh#L19
-		{"code-server", "gopls/client:code-server"},
-
-		// https://lists.gnu.org/archive/html/bug-gnu-emacs/2023-03/msg00954.html
-		{"Eglot", "gopls/client:eglot"},
-
-		// https://github.com/govim/govim/pull/1189
-		{"govim", "gopls/client:govim"},
-
-		// https://github.com/helix-editor/helix/blob/d0218f7e78bc0c3af4b0995ab8bda66b9c542cf3/helix-lsp/src/client.rs#L714
-		{"helix", "gopls/client:helix"},
-
-		// https://github.com/neovim/neovim/blob/42333ea98dfcd2994ee128a3467dfe68205154cd/runtime/lua/vim/lsp.lua#L1361
-		// https://github.com/neovim/neovim/blob/fe6026825883b44b09a8d3a03f2d49bfc8ed4725/runtime/lua/vim/lsp/client.lua#564
-		{"Neovim", "gopls/client:neovim"},
-
-		// https://github.com/neoclide/coc.nvim/blob/3dc6153a85ed0f185abec1deb972a66af3fbbfb4/src/language-client/client.ts#L994
-		{"coc.nvim", "gopls/client:coc.nvim"},
-
-		// https://github.com/sublimelsp/LSP/blob/e608f878e7e9dd34aabe4ff0462540fadcd88fcc/plugin/core/sessions.py#L493
-		{"Sublime Text LSP", "gopls/client:sublimetext"},
-
-		{"Cursor", "gopls/client:cursor"},
-
-		// https: //github.com/zed-industries/zed/blob/0ac17526687bf11007f0fbb5c3b2ff463ce47293/crates/release_channel/src/lib.rs#L147
-		{"Zed", "gopls/client:zed"}, // incl. "Zed Dev", "Zed Nightly", "Zed Preview"
-
-		// (Observed empirically.)
-		{"Claude Code", "gopls/client:claude"},
-
-		// (Observed empirically.)
-		{"Antigravity", "gopls/client:antigravity"},
-		{"Jetski", "gopls/client:antigravity"},
-		{"Windsurf", "gopls/client:windsurf"},
-	} {
-		if strings.HasPrefix(clientName, cli.clientNamePrefix) {
-			counter.Inc(cli.telemetryKey)
-			return
-		}
-	}
-
-	// Accumulate at least a local counter for an unknown
-	// client name, but also fall through to count it as
-	// ":other" for collection.
-	if clientName != "" {
-		counter.New(fmt.Sprintf("gopls/client-other:%s", clientName)).Inc()
-	}
-	counter.Inc("gopls/client:other")
 }
