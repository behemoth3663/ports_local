--- vendor/github.com/containerd/containerd/v2/core/mount/mount_windows.go.orig	2026-11-02 21:24:36 UTC
+++ vendor/github.com/containerd/containerd/v2/core/mount/mount_windows.go
@@ -27,7 +27,6 @@ import (
 	"strings"
 
 	"github.com/Microsoft/go-winio/pkg/bindfilter"
-	"github.com/Microsoft/hcsshim"
 	"github.com/containerd/log"
 	"golang.org/x/sys/windows"
 )
@@ -40,72 +39,8 @@ func (m *Mount) mount(target string) (retErr error) {
 		return fmt.Errorf("invalid windows mount type: '%s'", m.Type)
 	}
 
-	home, layerID := filepath.Split(m.Source)
-
-	parentLayerPaths, err := m.GetParentPaths()
-	if err != nil {
-		return err
-	}
-
-	var di = hcsshim.DriverInfo{
-		HomeDir: home,
-	}
-
-	if err := hcsshim.ActivateLayer(di, layerID); err != nil {
-		return fmt.Errorf("failed to activate layer %s: %w", m.Source, err)
-	}
-	defer func() {
-		if retErr != nil {
-			if layerErr := hcsshim.DeactivateLayer(di, layerID); layerErr != nil {
-				log.G(context.TODO()).WithError(layerErr).Error("failed to deactivate layer during mount failure cleanup")
-			}
-		}
-	}()
-
-	if err := hcsshim.PrepareLayer(di, layerID, parentLayerPaths); err != nil {
-		return fmt.Errorf("failed to prepare layer %s: %w", m.Source, err)
-	}
-
-	defer func() {
-		if retErr != nil {
-			if layerErr := hcsshim.UnprepareLayer(di, layerID); layerErr != nil {
-				log.G(context.TODO()).WithError(layerErr).Error("failed to unprepare layer during mount failure cleanup")
-			}
-		}
-	}()
-
-	volume, err := hcsshim.GetLayerMountPath(di, layerID)
-	if err != nil {
-		return fmt.Errorf("failed to get volume path for layer %s: %w", m.Source, err)
-	}
-
-	if len(parentLayerPaths) == 0 {
-		// this is a base layer. It gets mounted without going through WCIFS. We need to mount the Files
-		// folder, not the actual source, or the client may inadvertently remove metadata files.
-		volume = filepath.Join(volume, "Files")
-		if _, err := os.Stat(volume); err != nil {
-			return fmt.Errorf("no Files folder in layer %s", layerID)
-		}
-	}
-	if err := bindfilter.ApplyFileBinding(target, volume, m.ReadOnly()); err != nil {
-		return fmt.Errorf("failed to set volume mount path for layer %s: %w", m.Source, err)
-	}
-	defer func() {
-		if retErr != nil {
-			if bindErr := bindfilter.RemoveFileBinding(target); bindErr != nil {
-				log.G(context.TODO()).WithError(bindErr).Error("failed to remove binding during mount failure cleanup")
-			}
-		}
-	}()
-
-	// Add an Alternate Data Stream to record the layer source.
-	// See https://docs.microsoft.com/en-au/archive/blogs/askcore/alternate-data-streams-in-ntfs
-	// for details on Alternate Data Streams.
-	if err := os.WriteFile(filepath.Clean(target)+":"+sourceStreamName, []byte(m.Source), 0666); err != nil {
-		return fmt.Errorf("failed to record source for layer %s: %w", m.Source, err)
-	}
-
-	return nil
+	_, _ = target, retErr
+	return errors.New("windows-layer mount is not supported in this build")
 }
 
 // ParentLayerPathsFlag is the options flag used to represent the JSON encoded
@@ -148,22 +83,7 @@ func Unmount(mount string, flags int) error {
 		return fmt.Errorf("removing mount: %w", err)
 	}
 
-	if layerPath != "" {
-		var (
-			home, layerID = filepath.Split(layerPath)
-			di            = hcsshim.DriverInfo{
-				HomeDir: home,
-			}
-		)
-
-		if err := hcsshim.UnprepareLayer(di, layerID); err != nil {
-			return fmt.Errorf("failed to unprepare layer %s: %w", mount, err)
-		}
-
-		if err := hcsshim.DeactivateLayer(di, layerID); err != nil {
-			return fmt.Errorf("failed to deactivate layer %s: %w", mount, err)
-		}
-	}
+	_ = layerPath
 	return nil
 }
 
