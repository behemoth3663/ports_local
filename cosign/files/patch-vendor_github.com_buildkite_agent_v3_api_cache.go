--- vendor/github.com/buildkite/agent/v3/api/cache.go.orig	2026-10-02 15:35:42 UTC
+++ vendor/github.com/buildkite/agent/v3/api/cache.go
@@ -12,9 +12,6 @@ import (
 	"time"
 
 	"github.com/buildkite/agent/v3/internal/agenthttp"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/codes"
-	oteltrace "go.opentelemetry.io/otel/trace"
 )
 
 // Cache API "not found" messages. The cache service returns HTTP 404 with one
@@ -29,8 +26,6 @@ var ErrCacheEntryNotFound = errors.New("cache entry no
 // (resp, exists, err) return shape; this value is exported for parity.
 var ErrCacheEntryNotFound = errors.New("cache entry not found")
 
-var cacheTracer = otel.Tracer("github.com/buildkite/agent/v3/api/cache")
-
 // CacheKeyPart is one element of the structured cache_key sent on the wire.
 // Mandatory parts must match for a cache hit.
 type CacheKeyPart struct {
@@ -129,22 +124,19 @@ func (c *Client) CacheRegistry(ctx context.Context, re
 
 // CacheRegistry retrieves information about a cache registry.
 func (c *Client) CacheRegistry(ctx context.Context, registry string) (CacheRegistryResp, *Response, error) {
-	ctx, span := cacheTracer.Start(ctx, "Client.CacheRegistry")
-	defer span.End()
-
 	var cacheResp CacheRegistryResp
 
 	req, err := c.newRequest(ctx, http.MethodGet, cachePath("/cache_registries/%s", registry), nil)
 	if err != nil {
-		return cacheResp, nil, cacheSpanErr(span, "failed to create request: %w", err)
+		return cacheResp, nil, cacheSpanErr("failed to create request: %w", err)
 	}
 
 	apiResp, err := c.cacheDo(req, &cacheResp)
 	if err != nil {
-		return cacheResp, apiResp, cacheSpanErr(span, "%w", err)
+		return cacheResp, apiResp, cacheSpanErr("%w", err)
 	}
 	if apiResp.StatusCode != http.StatusOK {
-		return cacheResp, apiResp, cacheSpanErr(span, "failed to get cache registry: %s", apiResp.Status)
+		return cacheResp, apiResp, cacheSpanErr("failed to get cache registry: %s", apiResp.Status)
 	}
 	return cacheResp, apiResp, nil
 }
@@ -153,64 +145,55 @@ func (c *Client) CacheEntryPeekExists(ctx context.Cont
 // Returns (resp, true, _, nil) on hit, (resp, false, _, nil) on miss (HTTP 404
 // with CacheEntryNotFound), or (resp, false, _, err) on any other failure.
 func (c *Client) CacheEntryPeekExists(ctx context.Context, registry string, peek CacheEntryPeekReq) (CacheEntryPeekResp, bool, *Response, error) {
-	ctx, span := cacheTracer.Start(ctx, "Client.CacheEntryPeekExists")
-	defer span.End()
-
 	var cacheResp CacheEntryPeekResp
 
 	req, err := c.newRequest(ctx, http.MethodPost, cachePath("/cache_registries/%s/peek", registry), &peek)
 	if err != nil {
-		return cacheResp, false, nil, cacheSpanErr(span, "failed to create request: %w", err)
+		return cacheResp, false, nil, cacheSpanErr("failed to create request: %w", err)
 	}
 
 	apiResp, err := c.cacheDo(req, &cacheResp)
 	if err != nil {
-		return cacheResp, false, apiResp, cacheSpanErr(span, "%w", err)
+		return cacheResp, false, apiResp, cacheSpanErr("%w", err)
 	}
-	cacheResp, exists, err := interpretCacheResponse(span, apiResp, cacheResp)
+	cacheResp, exists, err := interpretCacheResponse(apiResp, cacheResp)
 	return cacheResp, exists, apiResp, err
 }
 
 // CacheEntryCreate creates a new cache entry and returns upload instructions.
 func (c *Client) CacheEntryCreate(ctx context.Context, registry string, create CacheEntryCreateReq) (CacheEntryCreateResp, *Response, error) {
-	ctx, span := cacheTracer.Start(ctx, "Client.CacheEntryCreate")
-	defer span.End()
-
 	var cacheResp CacheEntryCreateResp
 
 	req, err := c.newRequest(ctx, http.MethodPut, cachePath("/cache_registries/%s/store", registry), &create)
 	if err != nil {
-		return cacheResp, nil, cacheSpanErr(span, "failed to create request: %w", err)
+		return cacheResp, nil, cacheSpanErr("failed to create request: %w", err)
 	}
 
 	apiResp, err := c.cacheDo(req, &cacheResp)
 	if err != nil {
-		return cacheResp, apiResp, cacheSpanErr(span, "%w", err)
+		return cacheResp, apiResp, cacheSpanErr("%w", err)
 	}
 	if apiResp.StatusCode != http.StatusOK {
-		return cacheResp, apiResp, cacheSpanErr(span, "failed to save: %s", apiResp.Status)
+		return cacheResp, apiResp, cacheSpanErr("failed to save: %s", apiResp.Status)
 	}
 	return cacheResp, apiResp, nil
 }
 
 // CacheEntryCommit marks a previously created cache entry as committed.
 func (c *Client) CacheEntryCommit(ctx context.Context, registry string, commit CacheEntryCommitReq) (CacheEntryCommitResp, *Response, error) {
-	ctx, span := cacheTracer.Start(ctx, "Client.CacheEntryCommit")
-	defer span.End()
-
 	var cacheResp CacheEntryCommitResp
 
 	req, err := c.newRequest(ctx, http.MethodPut, cachePath("/cache_registries/%s/commit", registry), &commit)
 	if err != nil {
-		return cacheResp, nil, cacheSpanErr(span, "failed to create request: %w", err)
+		return cacheResp, nil, cacheSpanErr("failed to create request: %w", err)
 	}
 
 	apiResp, err := c.cacheDo(req, &cacheResp)
 	if err != nil {
-		return cacheResp, apiResp, cacheSpanErr(span, "%w", err)
+		return cacheResp, apiResp, cacheSpanErr("%w", err)
 	}
 	if apiResp.StatusCode != http.StatusOK {
-		return cacheResp, apiResp, cacheSpanErr(span, "failed to commit: %s", apiResp.Status)
+		return cacheResp, apiResp, cacheSpanErr("failed to commit: %s", apiResp.Status)
 	}
 	return cacheResp, apiResp, nil
 }
@@ -219,22 +202,19 @@ func (c *Client) CacheEntryRetrieve(ctx context.Contex
 // Returns (resp, true, _, nil) on hit (possibly via a fallback key),
 // (resp, false, _, nil) on miss, or (resp, false, _, err) on any other failure.
 func (c *Client) CacheEntryRetrieve(ctx context.Context, registry string, retrieve CacheEntryRetrieveReq) (CacheEntryRetrieveResp, bool, *Response, error) {
-	ctx, span := cacheTracer.Start(ctx, "Client.CacheEntryRetrieve")
-	defer span.End()
-
 	var cacheResp CacheEntryRetrieveResp
 
 	req, err := c.newRequest(ctx, http.MethodPost, cachePath("/cache_registries/%s/retrieve", registry), &retrieve)
 	if err != nil {
-		return cacheResp, false, nil, cacheSpanErr(span, "failed to create request: %w", err)
+		return cacheResp, false, nil, cacheSpanErr("failed to create request: %w", err)
 	}
 
 	apiResp, err := c.cacheDo(req, &cacheResp)
 	if err != nil {
-		return cacheResp, false, apiResp, cacheSpanErr(span, "%w", err)
+		return cacheResp, false, apiResp, cacheSpanErr("%w", err)
 	}
 
-	cacheResp, exists, err := interpretCacheResponse(span, apiResp, cacheResp)
+	cacheResp, exists, err := interpretCacheResponse(apiResp, cacheResp)
 	return cacheResp, exists, apiResp, err
 }
 
@@ -299,7 +279,7 @@ func (r CacheEntryRetrieveResp) cacheMessage() string 
 
 // interpretCacheResponse maps the dual "200 = hit, 404 + message = miss"
 // convention into the (resp, exists, err) return shape used by peek/retrieve.
-func interpretCacheResponse[T cacheMessage](span oteltrace.Span, apiResp *Response, cacheResp T) (T, bool, error) {
+func interpretCacheResponse[T cacheMessage](apiResp *Response, cacheResp T) (T, bool, error) {
 	if apiResp.StatusCode == http.StatusOK {
 		return cacheResp, true, nil
 	}
@@ -309,21 +289,18 @@ func interpretCacheResponse[T cacheMessage](span otelt
 		case CacheEntryNotFound:
 			return cacheResp, false, nil
 		case CacheRegistryNotFound:
-			return cacheResp, false, cacheSpanErr(span, "cache registry not found: %s", apiResp.Status)
+			return cacheResp, false, cacheSpanErr("cache registry not found: %s", apiResp.Status)
 		}
-		return cacheResp, false, cacheSpanErr(span, "not found: %s", apiResp.Status)
+		return cacheResp, false, cacheSpanErr("not found: %s", apiResp.Status)
 	case http.StatusBadRequest:
-		return cacheResp, false, cacheSpanErr(span, "bad request: %s", apiResp.Status)
+		return cacheResp, false, cacheSpanErr("bad request: %s", apiResp.Status)
 	default:
-		return cacheResp, false, cacheSpanErr(span, "request failed with status: %s", apiResp.Status)
+		return cacheResp, false, cacheSpanErr("request failed with status: %s", apiResp.Status)
 	}
 }
 
-func cacheSpanErr(span oteltrace.Span, format string, args ...any) error {
-	err := fmt.Errorf(format, args...)
-	span.RecordError(err)
-	span.SetStatus(codes.Error, err.Error())
-	return err
+func cacheSpanErr(format string, args ...any) error {
+	return fmt.Errorf(format, args...)
 }
 
 // isJSONContent reports whether contentType represents JSON, including media
