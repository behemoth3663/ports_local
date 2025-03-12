--- internal/cache/cache.go.orig	2025-03-12 13:22:56 UTC
+++ internal/cache/cache.go
@@ -9,8 +9,6 @@ import (
 	"fmt"
 	"sync"
 	"time"
-
-	"github.com/gruntwork-io/terragrunt/telemetry"
 )
 
 // Cache - generic cache implementation
@@ -38,14 +36,6 @@ func (c *Cache[V]) Get(ctx context.Context, key string
 	cacheKey := hex.EncodeToString(keyHash[:])
 	value, found := c.Cache[cacheKey]
 
-	telemetry.Count(ctx, c.Name+"_cache_get", 1)
-
-	if found {
-		telemetry.Count(ctx, c.Name+"_cache_hit", 1)
-	} else {
-		telemetry.Count(ctx, c.Name+"_cache_miss", 1)
-	}
-
 	return value, found
 }
 
@@ -53,7 +43,6 @@ func (c *Cache[V]) Put(ctx context.Context, key string
 func (c *Cache[V]) Put(ctx context.Context, key string, value V) {
 	c.Mutex.Lock()
 	defer c.Mutex.Unlock()
-	telemetry.Count(ctx, c.Name+"_cache_put", 1)
 
 	keyHash := sha256.Sum256([]byte(key))
 	cacheKey := hex.EncodeToString(keyHash[:])
@@ -87,22 +76,17 @@ func (c *ExpiringCache[V]) Get(ctx context.Context, ke
 	c.Mutex.Lock()
 	defer c.Mutex.Unlock()
 	item, found := c.Cache[key]
-	telemetry.Count(ctx, c.Name+"_cache_get", 1)
 
 	if !found {
-		telemetry.Count(ctx, c.Name+"_cache_miss", 1)
 		return item.Value, false
 	}
 
 	if time.Now().After(item.Expiration) {
-		telemetry.Count(ctx, c.Name+"_cache_expiry", 1)
 		delete(c.Cache, key)
 
 		return item.Value, false
 	}
 
-	telemetry.Count(ctx, c.Name+"_cache_hit", 1)
-
 	return item.Value, true
 }
 
@@ -111,7 +95,6 @@ func (c *ExpiringCache[V]) Put(ctx context.Context, ke
 	c.Mutex.Lock()
 	defer c.Mutex.Unlock()
 
-	telemetry.Count(ctx, c.Name+"_cache_put", 1)
 	c.Cache[key] = ExpiringItem[V]{Value: value, Expiration: expiration}
 }
 
