--- pkg/cmd/auth/refresh/refresh.go.orig	2026-10-26 18:02:47 UTC
+++ pkg/cmd/auth/refresh/refresh.go
@@ -198,7 +198,6 @@ func refreshRun(opts *RefreshOptions) error {
 
 	// Begin the auth-state telemetry event. It is nil for disabled telemetry or enterprise hosts, and all event methods
 	// are nil-tolerant, so the calls below stay unconditional.
-	authEvent := shared.BeginAuthTelemetry(opts.Telemetry, "refresh", hostname, expiring)
 
 	additionalScopes := set.NewStringSet()
 
@@ -247,11 +246,9 @@ func refreshRun(opts *RefreshOptions) error {
 	gitProtocol := cfg.GitProtocol(hostname).Value
 
 	// Record the selected git operations protocol for every refresh, before any prompt.
-	authEvent.RecordGitProtocol(gitProtocol)
 
 	// Record the configured credential helper up front, before any prompt, so this telemetry is captured even for
 	// non-interactive refreshes that never reach the prompt.
-	authEvent.RecordHelperBefore(credentialFlow.DetectConfiguredHelper(hostname))
 
 	if opts.Interactive && gitProtocol == "https" {
 		// refresh keeps the existing prompt; the expiring cred-helper prompt is login-only and is revisited in the
@@ -274,29 +271,23 @@ func refreshRun(opts *RefreshOptions) error {
 
 	// Record the server's issuance outcome as soon as the auth flow returns. It is independent of local persistence, so
 	// a later storage failure must not hide that an expiring, refreshable token was issued.
-	authEvent.RecordExpiringReceived(result.Refreshable != nil)
 
 	activeUser, _ := authCfg.ActiveUser(hostname)
 	if activeUser != "" && activeUser != result.Username {
 		return fmt.Errorf("error refreshing credentials for %s, received credentials for %s, did you use the correct account in the browser?", activeUser, result.Username)
 	}
 	if result.Refreshable != nil {
-		if insecureStorageUsed, err := authCfg.LoginRefreshable(hostname, result.Username, *result.Refreshable, "", !opts.InsecureStorage); err != nil {
+		if _, err := authCfg.LoginRefreshable(hostname, result.Username, *result.Refreshable, "", !opts.InsecureStorage); err != nil {
 			return err
-		} else {
-			authEvent.RecordStorage(insecureStorageUsed)
 		}
 	} else {
-		if insecureStorageUsed, err := authCfg.Login(hostname, result.Username, result.Token, "", !opts.InsecureStorage); err != nil {
+		if _, err := authCfg.Login(hostname, result.Username, result.Token, "", !opts.InsecureStorage); err != nil {
 			return err
-		} else {
-			authEvent.RecordStorage(insecureStorageUsed)
 		}
 	}
 
 	// Record success as soon as the credential is persisted. The git credential-helper setup below is a separate side
 	// effect whose failure must not retract a credential that was already stored successfully.
-	authEvent.RecordSuccess()
 
 	fmt.Fprintf(opts.IO.ErrOut, "%s Authentication complete.\n", cs.SuccessIcon())
 
@@ -328,7 +319,6 @@ func refreshRun(opts *RefreshOptions) error {
 	// Record the final helper state for every refresh. For non-interactive refreshes no setup runs, so this equals the
 	// before value (gh observed the helper and changed nothing); for interactive HTTPS it reflects any setup. It is
 	// derived from the flow's decisions rather than re-reading git configuration.
-	authEvent.RecordHelperAfter(credentialFlow.ResultingHelperCategory())
 
 	return nil
 }
