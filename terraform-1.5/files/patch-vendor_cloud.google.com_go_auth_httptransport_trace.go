--- vendor/cloud.google.com/go/auth/httptransport/trace.go.orig	2024-03-27 20:20:16 UTC
+++ vendor/cloud.google.com/go/auth/httptransport/trace.go
@@ -14,80 +14,11 @@ package httptransport
 
 package httptransport
 
-import (
-	"encoding/binary"
-	"encoding/hex"
-	"fmt"
-	"net/http"
-	"strconv"
-	"strings"
-
-	"go.opencensus.io/trace"
-	"go.opencensus.io/trace/propagation"
-)
-
 const (
 	httpHeaderMaxSize = 200
 	cloudTraceHeader  = `X-Cloud-Trace-Context`
 )
 
-// asserts the httpFormat fulfills this foreign interface
-var _ propagation.HTTPFormat = (*httpFormat)(nil)
-
 // httpFormat implements propagation.httpFormat to propagate
 // traces in HTTP headers for Google Cloud Platform and Cloud Trace.
 type httpFormat struct{}
-
-// SpanContextFromRequest extracts a Cloud Trace span context from incoming requests.
-func (f *httpFormat) SpanContextFromRequest(req *http.Request) (sc trace.SpanContext, ok bool) {
-	h := req.Header.Get(cloudTraceHeader)
-	// See https://cloud.google.com/trace/docs/faq for the header HTTPFormat.
-	// Return if the header is empty or missing, or if the header is unreasonably
-	// large, to avoid making unnecessary copies of a large string.
-	if h == "" || len(h) > httpHeaderMaxSize {
-		return trace.SpanContext{}, false
-	}
-
-	// Parse the trace id field.
-	slash := strings.Index(h, `/`)
-	if slash == -1 {
-		return trace.SpanContext{}, false
-	}
-	tid, h := h[:slash], h[slash+1:]
-
-	buf, err := hex.DecodeString(tid)
-	if err != nil {
-		return trace.SpanContext{}, false
-	}
-	copy(sc.TraceID[:], buf)
-
-	// Parse the span id field.
-	spanstr := h
-	semicolon := strings.Index(h, `;`)
-	if semicolon != -1 {
-		spanstr, h = h[:semicolon], h[semicolon+1:]
-	}
-	sid, err := strconv.ParseUint(spanstr, 10, 64)
-	if err != nil {
-		return trace.SpanContext{}, false
-	}
-	binary.BigEndian.PutUint64(sc.SpanID[:], sid)
-
-	// Parse the options field, options field is optional.
-	if !strings.HasPrefix(h, "o=") {
-		return sc, true
-	}
-	o, err := strconv.ParseUint(h[2:], 10, 32)
-	if err != nil {
-		return trace.SpanContext{}, false
-	}
-	sc.TraceOptions = trace.TraceOptions(o)
-	return sc, true
-}
-
-// SpanContextToRequest modifies the given request to include a Cloud Trace header.
-func (f *httpFormat) SpanContextToRequest(sc trace.SpanContext, req *http.Request) {
-	sid := binary.BigEndian.Uint64(sc.SpanID[:])
-	header := fmt.Sprintf("%s/%d;o=%d", hex.EncodeToString(sc.TraceID[:]), sid, int64(sc.TraceOptions))
-	req.Header.Set(cloudTraceHeader, header)
-}
