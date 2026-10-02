--- internal/cache/cache.go.orig	1979-11-29 21:00:00 UTC
+++ internal/cache/cache.go
@@ -9,8 +9,6 @@ import (
 	"fmt"
 	"sync"
 	"time"
-
-	"github.com/gruntwork-io/terragrunt/internal/telemetry"
 )
 
 // Cache - generic cache implementation
@@ -38,12 +36,8 @@ func (c *Cache[V]) Get(ctx context.Context, key string
 	cacheKey := hex.EncodeToString(keyHash[:])
 	value, found := c.Cache[cacheKey]
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, c.Name+"_cache_get", 1)
-
 	if found {
-		telemetry.TelemeterFromContext(ctx).Count(ctx, c.Name+"_cache_hit", 1)
 	} else {
-		telemetry.TelemeterFromContext(ctx).Count(ctx, c.Name+"_cache_miss", 1)
 	}
 
 	return value, found
@@ -54,8 +48,6 @@ func (c *Cache[V]) Put(ctx context.Context, key string
 	c.Mutex.Lock()
 	defer c.Mutex.Unlock()
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, c.Name+"_cache_put", 1)
-
 	keyHash := sha256.Sum256([]byte(key))
 	cacheKey := hex.EncodeToString(keyHash[:])
 	c.Cache[cacheKey] = value
@@ -89,22 +81,17 @@ func (c *ExpiringCache[V]) Get(ctx context.Context, ke
 	defer c.Mutex.Unlock()
 
 	item, found := c.Cache[key]
-	telemetry.TelemeterFromContext(ctx).Count(ctx, c.Name+"_cache_get", 1)
 
 	if !found {
-		telemetry.TelemeterFromContext(ctx).Count(ctx, c.Name+"_cache_miss", 1)
 		return item.Value, false
 	}
 
 	if time.Now().After(item.Expiration) {
-		telemetry.TelemeterFromContext(ctx).Count(ctx, c.Name+"_cache_expiry", 1)
 		delete(c.Cache, key)
 
 		return item.Value, false
 	}
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, c.Name+"_cache_hit", 1)
-
 	return item.Value, true
 }
 
@@ -112,8 +99,6 @@ func (c *ExpiringCache[V]) Put(ctx context.Context, ke
 func (c *ExpiringCache[V]) Put(ctx context.Context, key string, value V, expiration time.Time) {
 	c.Mutex.Lock()
 	defer c.Mutex.Unlock()
-
-	telemetry.TelemeterFromContext(ctx).Count(ctx, c.Name+"_cache_put", 1)
 	c.Cache[key] = ExpiringItem[V]{Value: value, Expiration: expiration}
 }
 
@@ -145,17 +130,11 @@ func (c *RepoRootCache) Lookup(ctx context.Context, di
 	c.mu.RLock()
 	defer c.mu.RUnlock()
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, c.name+"_cache_get", 1)
-
 	root, ok := c.roots[dir]
 	if !ok {
-		telemetry.TelemeterFromContext(ctx).Count(ctx, c.name+"_cache_miss", 1)
-
 		return "", false
 	}
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, c.name+"_cache_hit", 1)
-
 	return root, true
 }
 
@@ -174,8 +153,6 @@ func (c *RepoRootCache) Add(ctx context.Context, root 
 		if dir == "" {
 			continue
 		}
-
-		telemetry.TelemeterFromContext(ctx).Count(ctx, c.name+"_cache_put", 1)
 
 		c.roots[dir] = root
 	}
