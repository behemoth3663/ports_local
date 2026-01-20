--- internal/discovery/worktreediscovery.go.orig	2026-01-19 21:45:47 UTC
+++ internal/discovery/worktreediscovery.go
@@ -18,7 +18,6 @@ import (
 	"github.com/gruntwork-io/terragrunt/internal/component"
 	"github.com/gruntwork-io/terragrunt/internal/errors"
 	"github.com/gruntwork-io/terragrunt/internal/filter"
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 	"github.com/gruntwork-io/terragrunt/internal/worktrees"
 	"github.com/gruntwork-io/terragrunt/pkg/log"
 	"github.com/gruntwork-io/terragrunt/pkg/options"
@@ -174,13 +173,6 @@ func recordWorktreeDiscoveryMetrics(ctx context.Contex
 
 // recordWorktreeDiscoveryMetrics records telemetry metrics for worktree discovery.
 func recordWorktreeDiscoveryMetrics(ctx context.Context, pairCount, componentCount int) {
-	telemeter := telemetry.TelemeterFromContext(ctx)
-	if telemeter == nil || telemeter.Meter == nil {
-		return
-	}
-
-	telemeter.Count(ctx, "git_worktree_discovery_worktree_pairs", int64(pairCount))
-	telemeter.Count(ctx, "git_worktree_discovery_components", int64(componentCount))
 }
 
 // discoverInWorktree discovers components in a single worktree.
