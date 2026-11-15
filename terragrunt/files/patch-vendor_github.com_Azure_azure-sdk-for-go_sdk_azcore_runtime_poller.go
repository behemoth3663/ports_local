--- vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime/poller.go.orig	2026-11-15 15:04:53 UTC
+++ vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime/poller.go
@@ -23,7 +23,6 @@ import (
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/pollers/loc"
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/pollers/op"
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/internal/shared"
-	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
 	"github.com/Azure/azure-sdk-for-go/sdk/internal/poller"
 )
 
@@ -63,7 +62,7 @@ type NewPollerOptions[T any] struct {
 	Handler PollingHandler[T]
 
 	// Tracer contains the Tracer from the client that's creating the Poller.
-	Tracer tracing.Tracer
+	Tracer any
 }
 
 // NewPoller creates a Poller based on the provided initial response.
@@ -140,7 +139,7 @@ type NewPollerFromResumeTokenOptions[T any] struct {
 	Handler PollingHandler[T]
 
 	// Tracer contains the Tracer from the client that's creating the Poller.
-	Tracer tracing.Tracer
+	Tracer any
 }
 
 // NewPollerFromResumeToken creates a Poller from a resume token string.
@@ -212,7 +211,7 @@ type Poller[T any] struct {
 	resp   *http.Response
 	err    error
 	result *T
-	tracer tracing.Tracer
+	tracer any
 	done   bool
 }
 
