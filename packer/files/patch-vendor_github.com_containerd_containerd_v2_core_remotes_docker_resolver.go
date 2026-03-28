--- vendor/github.com/containerd/containerd/v2/core/remotes/docker/resolver.go.orig	2026-05-20 23:39:32 UTC
+++ vendor/github.com/containerd/containerd/v2/core/remotes/docker/resolver.go
@@ -40,7 +40,6 @@ import (
 	"github.com/containerd/containerd/v2/core/remotes"
 	"github.com/containerd/containerd/v2/core/transfer"
 	"github.com/containerd/containerd/v2/pkg/reference"
-	"github.com/containerd/containerd/v2/pkg/tracing"
 	"github.com/containerd/containerd/v2/version"
 )
 
@@ -637,8 +636,6 @@ func (r *request) do(ctx context.Context) (*http.Respo
 			return nil
 		}
 	}
-
-	tracing.UpdateHTTPClient(client, tracing.Name("remotes.docker.resolver", "HTTPRequest"))
 
 	resp, err := client.Do(req)
 	if err != nil {
