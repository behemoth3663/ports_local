--- vendor/cloud.google.com/go/auth/httptransport/transport.go.orig	2026-03-23 17:47:48 UTC
+++ vendor/cloud.google.com/go/auth/httptransport/transport.go
@@ -17,12 +17,9 @@ import (
 import (
 	"context"
 	"crypto/tls"
-	"errors"
-	"fmt"
 	"net"
 	"net/http"
 	"os"
-	"strconv"
 	"time"
 
 	"cloud.google.com/go/auth"
@@ -31,11 +28,6 @@ import (
 	"cloud.google.com/go/auth/internal/transport"
 	"cloud.google.com/go/auth/internal/transport/cert"
 	"cloud.google.com/go/auth/internal/transport/headers"
-	"github.com/googleapis/gax-go/v2"
-	"github.com/googleapis/gax-go/v2/callctx"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 	"golang.org/x/net/http2"
 )
 
@@ -177,22 +169,7 @@ func addOpenTelemetryTransport(trans http.RoundTripper
 }
 
 func addOpenTelemetryTransport(trans http.RoundTripper, opts *Options) http.RoundTripper {
-	if opts.DisableTelemetry {
 		return trans
-	}
-	if !gax.IsFeatureEnabled("TRACING") {
-		return otelhttp.NewTransport(trans)
-	}
-	var staticAttrs []attribute.KeyValue
-	if opts.InternalOptions != nil {
-		staticAttrs = transport.StaticTelemetryAttributes(opts.InternalOptions.TelemetryAttributes)
-	}
-	otelOpts := []otelhttp.Option{
-		otelhttp.WithSpanOptions(trace.WithAttributes(staticAttrs...)),
-	}
-	return otelhttp.NewTransport(&otelAttributeTransport{
-		base: trans,
-	}, otelOpts...)
 }
 
 // otelAttributeTransport is a wrapper around an http.RoundTripper that adds
@@ -205,55 +182,7 @@ func (t *otelAttributeTransport) RoundTrip(req *http.R
 // OpenTelemetry span with static and dynamic attributes, as well as detailed
 // error information.
 func (t *otelAttributeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
-	span := trace.SpanFromContext(req.Context())
-	if span.IsRecording() {
-		var attrs []attribute.KeyValue
-		attrs = append(attrs, attribute.String("rpc.system.name", "http"))
-		if resName, ok := callctx.TelemetryFromContext(req.Context(), "resource_name"); ok {
-			attrs = append(attrs, attribute.String("gcp.resource.destination.id", resName))
-		}
-		if resendCountStr, ok := callctx.TelemetryFromContext(req.Context(), "resend_count"); ok {
-			if count, err := strconv.Atoi(resendCountStr); err == nil {
-				attrs = append(attrs, attribute.Int("http.request.resend_count", count))
-			}
-		}
-		if urlTemplate, ok := callctx.TelemetryFromContext(req.Context(), "url_template"); ok {
-			attrs = append(attrs, attribute.String("url.template", urlTemplate))
-			span.SetName(fmt.Sprintf("%s %s", req.Method, urlTemplate))
-		}
-		span.SetAttributes(attrs...)
-	}
-
 	resp, err := t.base.RoundTrip(req)
-
-	if span.IsRecording() {
-		if err != nil {
-			var errorType string
-			switch {
-			case errors.Is(err, context.DeadlineExceeded):
-				errorType = "CLIENT_TIMEOUT"
-			case errors.Is(err, context.Canceled):
-				errorType = "CLIENT_CANCELLED"
-			default:
-				errorType = "CLIENT_CONNECTION_ERROR"
-			}
-			span.SetAttributes(
-				attribute.String("error.type", errorType),
-				attribute.String("status.message", err.Error()),
-				attribute.String("exception.type", fmt.Sprintf("%T", err)),
-			)
-		} else {
-			span.SetAttributes(attribute.Int("http.response.status_code", resp.StatusCode))
-			if resp.StatusCode >= 400 {
-				errorType := strconv.Itoa(resp.StatusCode)
-				span.SetAttributes(
-					attribute.String("error.type", errorType),
-					attribute.String("status.message", resp.Status),
-				)
-			}
-		}
-	}
-
 	return resp, err
 }
 
