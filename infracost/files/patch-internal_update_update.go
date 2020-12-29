--- internal/update/update.go.orig	2020-12-29 14:28:32 UTC
+++ internal/update/update.go
@@ -13,11 +13,8 @@ import (
 	"time"
 
 	"github.com/pkg/errors"
-	log "github.com/sirupsen/logrus"
-	"golang.org/x/mod/semver"
 
 	"github.com/infracost/infracost/internal/config"
-	"github.com/infracost/infracost/internal/version"
 )
 
 type Info struct {
@@ -26,70 +23,7 @@ type Info struct {
 }
 
 func CheckForUpdate() (*Info, error) {
-	if skipUpdateCheck() {
-		return nil, nil
-	}
-
-	var cmd string
-
-	state, err := config.ReadStateFileIfNotExists()
-	if err != nil {
-		log.Debugf("error reading state file: %v", err)
-	}
-
-	cachedRelease, err := checkCachedLatestVersion(state)
-	if err != nil {
-		log.Debugf("error getting cached latest version: %v", err)
-	}
-
-	latestVersion := cachedRelease
-
-	isBrew, err := isBrewInstall()
-	if err != nil {
-		// don't fail if we can't detect brew, just fallback to other update method
-		log.Debugf("error checking if executable was installed via brew: %v", err)
-	}
-
-	if isBrew && err == nil {
-		if latestVersion == "" {
-			latestVersion, err = getLatestBrewVersion()
-			if err != nil {
-				return nil, err
-			}
-		}
-
-		cmd = "$ brew upgrade infracost"
-	} else {
-		if latestVersion == "" {
-			latestVersion, err = getLatestGitHubVersion()
-			if err != nil {
-				return nil, err
-			}
-		}
-
-		cmd = "Go to https://www.infracost.io/docs/update for instructions"
-		if runtime.GOOS == "linux" && runtime.GOARCH == "amd64" {
-			cmd = "$ curl -s -L https://github.com/infracost/infracost/releases/latest/download/infracost-linux-amd64.tar.gz | tar xz -C /tmp && \\\n  sudo mv /tmp/infracost-linux-amd64 /usr/local/bin/infracost"
-		} else if runtime.GOOS == "darwin" && runtime.GOARCH == "amd64" {
-			cmd = "$ curl -s -L https://github.com/infracost/infracost/releases/latest/download/infracost-darwin-amd64.tar.gz | tar xz -C /tmp && \\\n  sudo mv /tmp/infracost-darwin-amd64 /usr/local/bin/infracost"
-		}
-	}
-
-	if cachedRelease == "" {
-		err := saveCachedLatestVersion(state, latestVersion)
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
 
 func skipUpdateCheck() bool {
