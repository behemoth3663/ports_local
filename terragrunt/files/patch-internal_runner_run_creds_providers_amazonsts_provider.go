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
@@ -143,17 +142,7 @@ func (provider *Provider) assumeAndCache(
 
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
+		resp, collectErr := awshelper.AssumeIamRole(ctx, assumeVenv(v, sourceEnv), iamRoleOpts, "")
 		if collectErr != nil {
 			return nil, collectErr
 		}
@@ -321,21 +310,15 @@ func (s *stsCredentialsStore) get(ctx context.Context,
 	s.mu.Lock()
 	defer s.mu.Unlock()
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, s.name+"_cache_get", 1)
-
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
 
@@ -343,21 +326,15 @@ func (s *stsCredentialsStore) getBySession(ctx context
 	s.mu.Lock()
 	defer s.mu.Unlock()
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, s.name+"_cache_get", 1)
-
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
 
@@ -365,24 +342,18 @@ func (s *stsCredentialsStore) getPastRefreshBySession(
 	s.mu.Lock()
 	defer s.mu.Unlock()
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, s.name+"_cache_get", 1)
-
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
@@ -396,8 +367,6 @@ func (s *stsCredentialsStore) put(ctx context.Context,
 func (s *stsCredentialsStore) put(ctx context.Context, identityKey, roleKey string, entry *cacheEntry) {
 	s.mu.Lock()
 	defer s.mu.Unlock()
-
-	telemetry.TelemeterFromContext(ctx).Count(ctx, s.name+"_cache_put", 1)
 
 	previous := s.byID[identityKey]
 	if previous != nil && previous.sessionFP != "" && previous.sessionFP != entry.sessionFP {
