--- internal/codespaces/api/api.go.orig	2025-01-30 18:14:09 UTC
+++ internal/codespaces/api/api.go
@@ -43,7 +43,6 @@ import (
 	"github.com/cli/cli/v2/api"
 	"github.com/cli/cli/v2/internal/ghinstance"
 	"github.com/cli/cli/v2/pkg/cmdutil"
-	"github.com/opentracing/opentracing-go"
 )
 
 const (
@@ -1180,9 +1179,6 @@ func (a *API) do(ctx context.Context, req *http.Reques
 // do executes the given request and returns the response. It creates an
 // opentracing span to track the length of the request.
 func (a *API) do(ctx context.Context, req *http.Request, spanName string) (*http.Response, error) {
-	// TODO(adonovan): use NewRequestWithContext(ctx) and drop ctx parameter.
-	span, ctx := opentracing.StartSpanFromContext(ctx, spanName)
-	defer span.Finish()
 	req = req.WithContext(ctx)
 
 	httpClient, err := a.client()
