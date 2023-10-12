--- internal/registry/module.go.orig	2023-10-12 15:35:06 UTC
+++ internal/registry/module.go
@@ -9,16 +9,11 @@ import (
 	"fmt"
 	"io/ioutil"
 	"net/http"
-	"net/http/httptrace"
 	"sort"
 	"time"
 
 	"github.com/hashicorp/go-version"
 	tfaddr "github.com/hashicorp/terraform-registry-address"
-	"go.opentelemetry.io/contrib/instrumentation/net/http/httptrace/otelhttptrace"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/trace"
 )
 
 type ModuleResponse struct {
@@ -67,8 +62,6 @@ func (rce ClientError) Error() string {
 }
 
 func (c Client) GetModuleData(ctx context.Context, addr tfaddr.Module, cons version.Constraints) (*ModuleResponse, error) {
-	ctx, span := otel.Tracer(tracerName).Start(ctx, "registry:GetModuleData")
-	defer span.End()
 	var response ModuleResponse
 
 	v, err := c.GetMatchingModuleVersion(ctx, addr, cons)
@@ -76,8 +69,6 @@ func (c Client) GetModuleData(ctx context.Context, add
 		return nil, err
 	}
 
-	ctx = httptrace.WithClientTrace(ctx, otelhttptrace.NewClientTrace(ctx, otelhttptrace.WithoutSubSpans()))
-
 	url := fmt.Sprintf("%s/v1/modules/%s/%s/%s/%s", c.BaseURL,
 		addr.Package.Namespace,
 		addr.Package.Name,
@@ -113,8 +104,6 @@ func (c Client) GetModuleData(ctx context.Context, add
 }
 
 func (c Client) GetMatchingModuleVersion(ctx context.Context, addr tfaddr.Module, con version.Constraints) (*version.Version, error) {
-	ctx, span := otel.Tracer(tracerName).Start(ctx, "registry:GetMatchingModuleVersion")
-	defer span.End()
 	foundVersions, err := c.GetModuleVersions(ctx, addr)
 	if err != nil {
 		return nil, err
@@ -130,16 +119,11 @@ func (c Client) GetMatchingModuleVersion(ctx context.C
 }
 
 func (c Client) GetModuleVersions(ctx context.Context, addr tfaddr.Module) (version.Collection, error) {
-	ctx, span := otel.Tracer(tracerName).Start(ctx, "registry:GetModuleVersions")
-	defer span.End()
-
 	url := fmt.Sprintf("%s/v1/modules/%s/%s/%s/versions", c.BaseURL,
 		addr.Package.Namespace,
 		addr.Package.Name,
 		addr.Package.TargetSystem)
 
-	ctx = httptrace.WithClientTrace(ctx, otelhttptrace.NewClientTrace(ctx, otelhttptrace.WithoutSubSpans()))
-
 	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
 	if err != nil {
 		return nil, err
@@ -160,13 +144,11 @@ func (c Client) GetModuleVersions(ctx context.Context,
 		return nil, ClientError{StatusCode: resp.StatusCode, Body: string(bodyBytes)}
 	}
 
-	_, decodeSpan := otel.Tracer(tracerName).Start(ctx, "registry:GetModuleVersions:decodeJson")
 	var response ModuleVersionsResponse
 	err = json.NewDecoder(resp.Body).Decode(&response)
 	if err != nil {
 		return nil, err
 	}
-	decodeSpan.End()
 
 	var foundVersions version.Collection
 	for _, module := range response.Modules {
@@ -177,12 +159,6 @@ func (c Client) GetModuleVersions(ctx context.Context,
 			}
 		}
 	}
-	span.AddEvent("registry:foundModuleVersions",
-		trace.WithAttributes(attribute.KeyValue{
-			Key:   attribute.Key("moduleVersionCount"),
-			Value: attribute.IntValue(len(foundVersions)),
-		}))
-
 	sort.Sort(sort.Reverse(foundVersions))
 
 	return foundVersions, nil
