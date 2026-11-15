--- vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime/pager.go.orig	2026-11-15 15:04:53 UTC
+++ vendor/github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime/pager.go
@@ -12,7 +12,6 @@ import (
 	"reflect"
 
 	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
-	"github.com/Azure/azure-sdk-for-go/sdk/azcore/tracing"
 )
 
 // PagingHandler contains the required data for constructing a Pager.
@@ -25,7 +24,7 @@ type PagingHandler[T any] struct {
 	Fetcher func(context.Context, *T) (T, error)
 
 	// Tracer contains the Tracer from the client that's creating the Pager.
-	Tracer tracing.Tracer
+	Tracer any
 }
 
 // Pager provides operations for iterating over paged responses.
@@ -33,7 +32,7 @@ type Pager[T any] struct {
 type Pager[T any] struct {
 	current   *T
 	handler   PagingHandler[T]
-	tracer    tracing.Tracer
+	tracer    any
 	firstPage bool
 	// fetchErr is set when Fetcher returns an error, placing the Pager in a
 	// terminal state so that More returns false.
