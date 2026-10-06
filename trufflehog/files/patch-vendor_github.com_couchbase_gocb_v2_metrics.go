--- vendor/github.com/couchbase/gocb/v2/metrics.go.orig	2026-10-06 09:46:09 UTC
+++ vendor/github.com/couchbase/gocb/v2/metrics.go
@@ -3,10 +3,10 @@ import (
 import (
 	"errors"
 	"fmt"
-	"github.com/couchbase/gocbcore/v10"
-	"go.opentelemetry.io/otel/metric"
 	"sync"
 	"time"
+
+	"github.com/couchbase/gocbcore/v10"
 )
 
 // Meter handles metrics information for SDK operations.
@@ -16,8 +16,8 @@ type OtelAwareMeter interface {
 }
 
 type OtelAwareMeter interface {
-	Wrapped() metric.Meter
-	Provider() metric.MeterProvider
+	Wrapped() interface{}
+	Provider() interface{}
 }
 
 // Counter is used for incrementing a synchronous count metric.
@@ -31,8 +31,7 @@ type ValueRecorder interface {
 }
 
 // NoopMeter is a Meter implementation which performs no metrics operations.
-type NoopMeter struct {
-}
+type NoopMeter struct{}
 
 var (
 	defaultNoopCounter       = &noopCounter{}
