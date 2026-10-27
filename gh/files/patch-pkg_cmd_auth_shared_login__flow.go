--- pkg/cmd/auth/shared/login_flow.go.orig	2026-10-26 18:02:47 UTC
+++ pkg/cmd/auth/shared/login_flow.go
@@ -106,7 +106,6 @@ func Login(opts *LoginOptions) error {
 
 	// Begin the auth-state telemetry event. It is nil for disabled telemetry or enterprise hosts, and all event methods
 	// are nil-tolerant, so the calls below stay unconditional.
-	authEvent := BeginAuthTelemetry(opts.Telemetry, "login", hostname, expiring)
 
 	// A GitHub Enterprise Server host may not support expiring tokens. Surface this once the hostname is known, from
 	// interactive selection or the --hostname flag, so the user understands they may receive a non-expiring token.
@@ -134,16 +133,13 @@ func Login(opts *LoginOptions) error {
 	var additionalScopes []string
 
 	// Record the selected git operations protocol for every sign-in, before any prompt.
-	authEvent.RecordGitProtocol(gitProtocol)
 
 	// Record the configured credential helper up front, before any prompt, so this telemetry is captured for every
 	// sign-in, including non-interactive and SSH flows that never reach the prompt.
-	authEvent.RecordHelperBefore(opts.CredentialFlow.DetectConfiguredHelper(hostname))
 
 	if opts.Interactive && gitProtocol == "https" {
 		err := opts.CredentialFlow.Prompt(hostname, expiring)
 		if errors.Is(err, ErrExpiringSetupCredHelperCancelled) {
-			authEvent.RecordExpiringSetupCancelled(true)
 			return cmdutil.CancelError
 		}
 		if err != nil {
@@ -241,7 +237,6 @@ func Login(opts *LoginOptions) error {
 
 	// Record the server's issuance outcome as soon as the auth flow returns. It is independent of local persistence, so
 	// a later storage failure must not hide that an expiring, refreshable token was issued.
-	authEvent.RecordExpiringReceived(refreshableCredential != nil)
 
 	if username == "" {
 		var err error
@@ -279,11 +274,9 @@ func Login(opts *LoginOptions) error {
 		fmt.Fprintf(opts.IO.ErrOut, "%s Authentication credentials saved in plain text\n", cs.Yellow("!"))
 	}
 
-	authEvent.RecordStorage(insecureStorageUsed)
 
 	// Record success as soon as the credential is persisted. The git credential-helper setup and SSH key upload below
 	// are separate side effects whose failure must not retract a credential that was already stored successfully.
-	authEvent.RecordSuccess()
 
 	// Expiring tokens are a preference, not a guarantee: the server decides based on its own and the OAuth app
 	// configuration, so it may issue a refreshable token without being asked or a non-expiring one despite the request.
@@ -309,7 +302,6 @@ func Login(opts *LoginOptions) error {
 	// Record the final helper state for every sign-in. For non-interactive and SSH flows no setup runs, so this equals
 	// the before value (gh observed the helper and changed nothing); for interactive HTTPS it reflects any setup. It is
 	// derived from the flow's decisions rather than re-reading git configuration.
-	authEvent.RecordHelperAfter(opts.CredentialFlow.ResultingHelperCategory())
 
 	if keyToUpload != "" {
 		uploaded, err := sshKeyUpload(httpClient, hostname, keyToUpload, keyTitle)
