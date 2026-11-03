--- vendor/github.com/containerd/containerd/v2/pkg/archive/tar_opts_windows.go.orig	2026-11-02 21:24:36 UTC
+++ vendor/github.com/containerd/containerd/v2/pkg/archive/tar_opts_windows.go
@@ -18,22 +18,21 @@ import (
 
 import (
 	"context"
+	"errors"
 	"io"
 
 	"github.com/Microsoft/go-winio"
-	"github.com/Microsoft/hcsshim/pkg/ociwclayer"
 )
 
 // applyWindowsLayer applies a tar stream of an OCI style diff tar of a Windows layer
 // See https://github.com/opencontainers/image-spec/blob/main/layer.md#applying-changesets
 func applyWindowsLayer(ctx context.Context, root string, r io.Reader, options ApplyOptions) (size int64, err error) {
+	_, _, _, _ = ctx, root, r, options
 	// It seems that in certain situations, like having the containerd root and state on a file system hosted on a
 	// mounted VHDX, we need SeSecurityPrivilege when opening a file with winio.ACCESS_SYSTEM_SECURITY. This happens
 	// in the base layer writer in hcsshim when adding a new file.
 	err = winio.RunWithPrivileges([]string{winio.SeSecurityPrivilege}, func() error {
-		var innerErr error
-		size, innerErr = ociwclayer.ImportLayerFromTar(ctx, r, root, options.Parents)
-		return innerErr
+		return errors.New("windows container layer import is not supported in this build")
 	})
 	return
 }
@@ -56,7 +55,8 @@ func writeDiffWindowsLayers(ctx context.Context, w io.
 // based off AUFS whiteouts.
 // See https://github.com/opencontainers/image-spec/blob/main/layer.md
 func writeDiffWindowsLayers(ctx context.Context, w io.Writer, _, layer string, options WriteDiffOptions) error {
-	return ociwclayer.ExportLayerToTar(ctx, w, layer, options.ParentLayers)
+	_, _, _, _ = ctx, w, layer, options
+	return errors.New("windows container layer diff is not supported in this build")
 }
 
 // AsWindowsContainerLayerPair indicates that the paths to diff are a pair of
