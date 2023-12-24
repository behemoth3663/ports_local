--- internal/update/update.go.orig	2023-12-05 22:53:36 UTC
+++ internal/update/update.go
@@ -14,11 +14,8 @@ import (
 	"time"
 
 	"github.com/pkg/errors"
-	"github.com/rs/zerolog/log"
-	"golang.org/x/mod/semver"
 
 	"github.com/infracost/infracost/internal/config"
-	"github.com/infracost/infracost/internal/version"
 )
 
 type Info struct {
@@ -27,66 +24,7 @@ func CheckForUpdate(ctx *config.RunContext) (*Info, er
 }
 
 func CheckForUpdate(ctx *config.RunContext) (*Info, error) {
-	if skipUpdateCheck(ctx) {
-		return nil, nil
-	}
-
-	// Check cache for the latest version
-	cachedLatestVersion, err := checkCachedLatestVersion(ctx)
-	if err != nil {
-		log.Debug().Msgf("error getting cached latest version: %v", err)
-	}
-
-	// Nothing to do if the current version is the latest cached version
-	if cachedLatestVersion != "" && semver.Compare(version.Version, cachedLatestVersion) >= 0 {
-		return nil, nil
-	}
-
-	isBrew, err := isBrewInstall()
-	if err != nil {
-		// don't fail if we can't detect brew, just fallback to other update method
-		log.Debug().Msgf("error checking if executable was installed via brew: %v", err)
-	}
-
-	var cmd string
-	if isBrew {
-		cmd = "$ brew upgrade infracost"
-	} else {
-		cmd = "Go to https://www.infracost.io/docs/update for instructions"
-		if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
-			cmd = "$ curl -fsSL https://raw.githubusercontent.com/infracost/infracost/master/scripts/install.sh | sh"
-		}
-	}
-
-	// Get the latest version
-	latestVersion := cachedLatestVersion
-	if latestVersion == "" {
-		if isBrew {
-			latestVersion, err = getLatestBrewVersion()
-		} else {
-			latestVersion, err = getLatestGitHubVersion()
-		}
-		if err != nil {
-			return nil, err
-		}
-	}
-
-	// Save the latest version in the cache
-	if latestVersion != cachedLatestVersion {
-		err := setCachedLatestVersion(ctx, latestVersion)
-		if err != nil {
-			log.Debug().Msgf("error saving cached latest version: %v", err)
-		}
-	}
-
-	if semver.Compare(version.Version, latestVersion) >= 0 {
-		return nil, nil
-	}
-
-	return &Info{
-		LatestVersion: latestVersion,
-		Cmd:           cmd,
-	}, nil
+	return nil, nil
 }
 
 func skipUpdateCheck(ctx *config.RunContext) bool {
