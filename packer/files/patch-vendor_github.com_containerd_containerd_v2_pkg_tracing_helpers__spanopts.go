--- vendor/github.com/containerd/containerd/v2/pkg/tracing/helpers_spanopts.go.orig	2026-10-02 20:06:39 UTC
+++ vendor/github.com/containerd/containerd/v2/pkg/tracing/helpers_spanopts.go
@@ -20,7 +20,6 @@ import (
 	"context"
 
 	"github.com/containerd/containerd/v2/pkg/namespaces"
-	"go.opentelemetry.io/otel/trace"
 )
 
 // WithNamespace adds containerd namespace attribute to spans when available.
@@ -31,8 +30,6 @@ func WithNamespace(ctx context.Context) SpanOpt {
 		if err != nil {
 			return
 		}
-		config.spanOpts = append(config.spanOpts,
-			trace.WithAttributes(Attribute("namespace", ns)),
-		)
+		config.spanOpts = append(config.spanOpts, Attribute("namespace", ns))
 	}
 }
