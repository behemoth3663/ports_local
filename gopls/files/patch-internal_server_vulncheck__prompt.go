--- internal/server/vulncheck_prompt.go.orig	2026-05-15 14:53:57 UTC
+++ internal/server/vulncheck_prompt.go
@@ -175,7 +175,6 @@ func (s *server) checkDependencyChanges(ctx context.Co
 				event.Error(ctx, "parsing vulncheck action failed", fmt.Errorf("unexpected action: %s", choice))
 				return
 			}
-			recordVulncheckAction(action)
 			if action == vulncheckActionAlways || action == vulncheckActionNever {
 				if err := setVulncheckPreference(action); err != nil {
 					event.Error(ctx, "writing vulncheck preference failed", err)
@@ -192,28 +191,6 @@ func (s *server) checkDependencyChanges(ctx context.Co
 	}
 }
 
-func recordVulncheckAction(action vulncheckAction) {
-	switch action {
-	case vulncheckActionYes:
-		countVulncheckPromptYes.Inc()
-	case vulncheckActionNo:
-		countVulncheckPromptNo.Inc()
-	case vulncheckActionAlways:
-		countVulncheckPromptAlways.Inc()
-	case vulncheckActionNever:
-		countVulncheckPromptNever.Inc()
-	}
-}
-
-func recordVulncheckupgradeAction(action vulnupgradeAction) {
-	switch action {
-	case vulnupgradeActionUpgradeAll:
-		countVulncheckUpgradeAll.Inc()
-	case vulnupgradeActionIgnore:
-		countVulncheckUpgradeIgnore.Inc()
-	}
-}
-
 func (s *server) handleVulncheck(ctx context.Context, uri protocol.DocumentURI) {
 	_, snapshot, release, err := s.session.FileOf(ctx, uri)
 	if err != nil {
@@ -285,7 +262,6 @@ func (s *server) handleVulncheck(ctx context.Context, 
 		event.Error(ctx, "parsing vulncheck remediation action failed", fmt.Errorf("unexpected action: %s", action))
 		return
 	}
-	recordVulncheckupgradeAction(upgradeAction)
 
 	if upgradeAction == vulnupgradeActionUpgradeAll {
 		if err := s.upgradeModules(ctx, snapshot, uri, modulesToUpgrade); err != nil {
