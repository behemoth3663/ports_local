--- vendor/github.com/elastic/go-elasticsearch/v8/elasticsearch.go.orig	2026-10-06 09:46:09 UTC
+++ vendor/github.com/elastic/go-elasticsearch/v8/elasticsearch.go
@@ -21,7 +21,6 @@ import (
 	"encoding/base64"
 	"errors"
 	"fmt"
-	"go.opentelemetry.io/otel/trace"
 	"net/http"
 	"net/url"
 	"os"
@@ -127,7 +126,7 @@ type Config struct {
 //	search_template
 //	msearch_template
 //	render_search_template
-func NewOpenTelemetryInstrumentation(provider trace.TracerProvider, captureSearchBody bool) elastictransport.Instrumentation {
+func NewOpenTelemetryInstrumentation(provider interface{}, captureSearchBody bool) elastictransport.Instrumentation {
 	return elastictransport.NewOtelInstrumentation(provider, captureSearchBody, Version)
 }
 
@@ -448,7 +447,7 @@ func addrFromCloudID(input string) (string, error) {
 // addrFromCloudID extracts the Elasticsearch URL from CloudID.
 // See: https://www.elastic.co/guide/en/cloud/current/ec-cloud-id.html
 func addrFromCloudID(input string) (string, error) {
-	var scheme = "https://"
+	scheme := "https://"
 
 	values := strings.Split(input, ":")
 	if len(values) != 2 {
@@ -505,7 +504,7 @@ func initMetaHeader(transport interface{}) string {
 		strippedTransportVersion = strippedEsVersion
 	}
 
-	var duos = [][]string{
+	duos := [][]string{
 		{
 			"es",
 			strippedEsVersion,
