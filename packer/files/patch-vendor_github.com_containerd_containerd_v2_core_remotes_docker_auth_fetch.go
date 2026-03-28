--- vendor/github.com/containerd/containerd/v2/core/remotes/docker/auth/fetch.go.orig	2025-12-18 16:59:11 UTC
+++ vendor/github.com/containerd/containerd/v2/core/remotes/docker/auth/fetch.go
@@ -27,7 +27,6 @@ import (
 	"time"
 
 	remoteserrors "github.com/containerd/containerd/v2/core/remotes/errors"
-	"github.com/containerd/containerd/v2/pkg/tracing"
 	"github.com/containerd/containerd/v2/version"
 	"github.com/containerd/log"
 )
@@ -98,7 +97,6 @@ func FetchTokenWithOAuth(ctx context.Context, client *
 func FetchTokenWithOAuth(ctx context.Context, client *http.Client, headers http.Header, clientID string, to TokenOptions) (*OAuthTokenResponse, error) {
 	c := *client
 	client = &c
-	tracing.UpdateHTTPClient(client, tracing.Name("remotes.docker.resolver", "FetchTokenWithOAuth"))
 
 	form := url.Values{}
 	if len(to.Scopes) > 0 {
@@ -168,7 +166,6 @@ func FetchToken(ctx context.Context, client *http.Clie
 func FetchToken(ctx context.Context, client *http.Client, headers http.Header, to TokenOptions) (*FetchTokenResponse, error) {
 	c := *client
 	client = &c
-	tracing.UpdateHTTPClient(client, tracing.Name("remotes.docker.resolver", "FetchToken"))
 
 	req, err := http.NewRequestWithContext(ctx, http.MethodGet, to.Realm, nil)
 	if err != nil {
