--- internal/runner/run/creds/providers/amazonsts/provider.go.orig	1979-11-29 21:00:00 UTC
+++ internal/runner/run/creds/providers/amazonsts/provider.go
@@ -17,7 +17,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/awshelper"
 	"github.com/gruntwork-io/terragrunt/internal/iam"
 	"github.com/gruntwork-io/terragrunt/internal/runner/run/creds/providers"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/venv"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"golang.org/x/sync/singleflight"
@@ -143,19 +142,9 @@ func (provider *Provider) assumeAndCache(
 
 		var resp *types.Credentials
 
-		collectErr := telemetry.TelemeterFromContext(ctx).Collect(ctx, l, "creds_assume_role", map[string]any{
-			"role_arn":     iamRoleOpts.RoleARN,
-			"session_name": iamRoleOpts.AssumeRoleSessionName,
-			"duration":     iamRoleOpts.AssumeRoleDuration,
-		}, func(ctx context.Context, l log.Logger) error {
-			var assumeErr error
-
-			resp, assumeErr = awshelper.AssumeIamRole(ctx, assumeVenv(v, sourceEnv), iamRoleOpts, "")
-
-			return assumeErr
-		})
-		if collectErr != nil {
-			return nil, collectErr
+		resp, err := awshelper.AssumeIamRole(ctx, assumeVenv(v, sourceEnv), iamRoleOpts, "")
+		if err != nil {
+			return nil, err
 		}
 
 		creds := &providers.Credentials{
@@ -321,21 +310,17 @@ func (s *stsCredentialsStore) get(ctx context.Context,
 	s.mu.Lock()
 	defer s.mu.Unlock()
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, s.name+"_cache_get", 1)
+	_ = ctx
 
 	entry, found := s.byID[identityKey]
 	if !found {
-		telemetry.TelemeterFromContext(ctx).Count(ctx, s.name+"_cache_miss", 1)
 		return nil
 	}
 
 	if time.Now().After(entry.refreshAt) {
-		telemetry.TelemeterFromContext(ctx).Count(ctx, s.name+"_cache_expiry", 1)
 		return nil
 	}
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, s.name+"_cache_hit", 1)
-
 	return entry
 }
 
@@ -343,21 +328,17 @@ func (s *stsCredentialsStore) getBySession(ctx context
 	s.mu.Lock()
 	defer s.mu.Unlock()
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, s.name+"_cache_get", 1)
+	_ = ctx
 
 	entry, found := s.bySession[sessionIndexKey(roleKey, sessionFP)]
 	if !found {
-		telemetry.TelemeterFromContext(ctx).Count(ctx, s.name+"_cache_miss", 1)
 		return nil
 	}
 
 	if time.Now().After(entry.refreshAt) {
-		telemetry.TelemeterFromContext(ctx).Count(ctx, s.name+"_cache_expiry", 1)
 		return nil
 	}
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, s.name+"_cache_hit", 1)
-
 	return entry
 }
 
@@ -365,24 +346,20 @@ func (s *stsCredentialsStore) getPastRefreshBySession(
 	s.mu.Lock()
 	defer s.mu.Unlock()
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, s.name+"_cache_get", 1)
+	_ = ctx
 
 	key := sessionIndexKey(roleKey, sessionFP)
 
 	entry, found := s.bySession[key]
 	if !found {
-		telemetry.TelemeterFromContext(ctx).Count(ctx, s.name+"_cache_miss", 1)
 		return nil
 	}
 
 	now := time.Now()
 	if !now.After(entry.refreshAt) {
-		telemetry.TelemeterFromContext(ctx).Count(ctx, s.name+"_cache_hit", 1)
 		return nil
 	}
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, s.name+"_cache_expiry", 1)
-
 	if now.After(entry.expiresAt) {
 		delete(s.bySession, key)
 		s.deleteIdentityLocked(entry)
@@ -397,7 +374,7 @@ func (s *stsCredentialsStore) put(ctx context.Context,
 	s.mu.Lock()
 	defer s.mu.Unlock()
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, s.name+"_cache_put", 1)
+	_ = ctx
 
 	previous := s.byID[identityKey]
 	if previous != nil && previous.sessionFP != "" && previous.sessionFP != entry.sessionFP {
