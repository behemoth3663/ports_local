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
@@ -38,14 +36,8 @@ func (c *Cache[V]) Get(ctx context.Context, key string
 	cacheKey := hex.EncodeToString(keyHash[:])
 	value, found := c.Cache[cacheKey]
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, c.Name+"_cache_get", 1)
+	_ = ctx
 
-	if found {
-		telemetry.TelemeterFromContext(ctx).Count(ctx, c.Name+"_cache_hit", 1)
-	} else {
-		telemetry.TelemeterFromContext(ctx).Count(ctx, c.Name+"_cache_miss", 1)
-	}
-
 	return value, found
 }
 
@@ -54,7 +46,7 @@ func (c *Cache[V]) Put(ctx context.Context, key string
 	c.Mutex.Lock()
 	defer c.Mutex.Unlock()
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, c.Name+"_cache_put", 1)
+	_ = ctx
 
 	keyHash := sha256.Sum256([]byte(key))
 	cacheKey := hex.EncodeToString(keyHash[:])
@@ -89,22 +81,18 @@ func (c *ExpiringCache[V]) Get(ctx context.Context, ke
 	defer c.Mutex.Unlock()
 
 	item, found := c.Cache[key]
-	telemetry.TelemeterFromContext(ctx).Count(ctx, c.Name+"_cache_get", 1)
+	_ = ctx
 
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
 
@@ -113,7 +101,7 @@ func (c *ExpiringCache[V]) Put(ctx context.Context, ke
 	c.Mutex.Lock()
 	defer c.Mutex.Unlock()
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, c.Name+"_cache_put", 1)
+	_ = ctx
 	c.Cache[key] = ExpiringItem[V]{Value: value, Expiration: expiration}
 }
 
@@ -146,17 +134,13 @@ func (c *RepoRootCache) Lookup(ctx context.Context, di
 	c.mu.RLock()
 	defer c.mu.RUnlock()
 
-	telemetry.TelemeterFromContext(ctx).Count(ctx, c.name+"_cache_get", 1)
+	_ = ctx
 
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
 
@@ -175,8 +159,6 @@ func (c *RepoRootCache) Add(ctx context.Context, root 
 		if dir == "" {
 			continue
 		}
-
-		telemetry.TelemeterFromContext(ctx).Count(ctx, c.name+"_cache_put", 1)
 
 		c.roots[dir] = root
 	}
