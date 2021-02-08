--- internal/update/update.go.orig	2021-02-07 19:16:38 UTC
+++ internal/update/update.go
@@ -14,11 +14,8 @@ import (
 	"time"
 
 	"github.com/pkg/errors"
-	log "github.com/sirupsen/logrus"
-	"golang.org/x/mod/semver"
 
 	"github.com/infracost/infracost/internal/config"
-	"github.com/infracost/infracost/internal/version"
 )
 
 type Info struct {
@@ -27,68 +24,7 @@ type Info struct {
 }
 
 func CheckForUpdate(cfg *config.Config) (*Info, error) {
-	if skipUpdateCheck(cfg) {
-		return nil, nil
-	}
-
-	// Check cache for the latest version
-	cachedLatestVersion, err := checkCachedLatestVersion(cfg)
-	if err != nil {
-		log.Debugf("error getting cached latest version: %v", err)
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
-		log.Debugf("error checking if executable was installed via brew: %v", err)
-	}
-
-	var cmd string
-	if isBrew {
-		cmd = "$ brew upgrade infracost"
-	} else {
-		cmd = "Go to https://www.infracost.io/docs/update for instructions"
-		if runtime.GOOS == "linux" && runtime.GOARCH == "amd64" {
-			cmd = "$ curl -s -L https://github.com/infracost/infracost/releases/latest/download/infracost-linux-amd64.tar.gz | tar xz -C /tmp && \\\n  sudo mv /tmp/infracost-linux-amd64 /usr/local/bin/infracost"
-		} else if runtime.GOOS == "darwin" && runtime.GOARCH == "amd64" {
-			cmd = "$ curl -s -L https://github.com/infracost/infracost/releases/latest/download/infracost-darwin-amd64.tar.gz | tar xz -C /tmp && \\\n  sudo mv /tmp/infracost-darwin-amd64 /usr/local/bin/infracost"
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
-		err := setCachedLatestVersion(cfg, latestVersion)
-		if err != nil {
-			log.Debugf("error saving cached latest version: %v", err)
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
 
 func skipUpdateCheck(cfg *config.Config) bool {
