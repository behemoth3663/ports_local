--- internal/terraform/module/module_ops.go.orig	2023-10-12 15:35:06 UTC
+++ internal/terraform/module/module_ops.go
@@ -43,10 +43,6 @@ import (
 	tfschema "github.com/hashicorp/terraform-schema/schema"
 	"github.com/zclconf/go-cty/cty"
 	ctyjson "github.com/zclconf/go-cty/cty/json"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 )
 
 const tracerName = "github.com/hashicorp/terraform-ls/internal/terraform/module"
@@ -286,17 +282,8 @@ func preloadSchemaForProviderAddr(ctx context.Context,
 			originalAddr.ForDisplay(), pAddr.ForDisplay())
 	}
 
-	ctx, rootSpan := otel.Tracer(tracerName).Start(ctx, "preloadProviderSchema",
-		trace.WithAttributes(attribute.KeyValue{
-			Key:   attribute.Key("ProviderAddress"),
-			Value: attribute.StringValue(pAddr.String()),
-		}))
-	defer rootSpan.End()
-
 	pSchemaFile, err := schemas.FindProviderSchemaFile(fs, pAddr)
 	if err != nil {
-		rootSpan.RecordError(err)
-		rootSpan.SetStatus(codes.Error, "schema file not found")
 		if errors.Is(err, schemas.SchemaNotAvailable{Addr: pAddr}) {
 			logger.Printf("preloaded schema not available for %s", pAddr)
 			return nil
@@ -304,34 +291,16 @@ func preloadSchemaForProviderAddr(ctx context.Context,
 		return err
 	}
 
-	_, span := otel.Tracer(tracerName).Start(ctx, "readProviderSchemaFile",
-		trace.WithAttributes(attribute.KeyValue{
-			Key:   attribute.Key("ProviderAddress"),
-			Value: attribute.StringValue(pAddr.String()),
-		}))
 	b, err := io.ReadAll(pSchemaFile.File)
 	if err != nil {
-		span.RecordError(err)
-		span.SetStatus(codes.Error, "schema file not readable")
 		return err
 	}
-	span.SetStatus(codes.Ok, "schema file read successfully")
-	span.End()
 
-	_, span = otel.Tracer(tracerName).Start(ctx, "decodeProviderSchemaData",
-		trace.WithAttributes(attribute.KeyValue{
-			Key:   attribute.Key("ProviderAddress"),
-			Value: attribute.StringValue(pAddr.String()),
-		}))
 	jsonSchemas := tfjson.ProviderSchemas{}
 	err = json.Unmarshal(b, &jsonSchemas)
 	if err != nil {
-		span.RecordError(err)
-		span.SetStatus(codes.Error, "schema file not decodable")
 		return err
 	}
-	span.SetStatus(codes.Ok, "schema data decoded successfully")
-	span.End()
 
 	ps, ok := jsonSchemas.Schemas[pAddr.String()]
 	if !ok {
@@ -341,16 +310,8 @@ func preloadSchemaForProviderAddr(ctx context.Context,
 	pSchema := tfschema.ProviderSchemaFromJson(ps, pAddr)
 	pSchema.SetProviderVersion(pAddr, pSchemaFile.Version)
 
-	_, span = otel.Tracer(tracerName).Start(ctx, "loadProviderSchemaDataIntoMemDb",
-		trace.WithAttributes(attribute.KeyValue{
-			Key:   attribute.Key("ProviderAddress"),
-			Value: attribute.StringValue(pAddr.String()),
-		}))
 	err = schemaStore.AddPreloadedSchema(pAddr, pSchemaFile.Version, pSchema)
 	if err != nil {
-		span.RecordError(err)
-		span.SetStatus(codes.Error, "loading schema into mem-db failed")
-		span.End()
 		existsError := &state.AlreadyExistsError{}
 		if errors.As(err, &existsError) {
 			// This accounts for a possible race condition
@@ -361,12 +322,9 @@ func preloadSchemaForProviderAddr(ctx context.Context,
 		}
 		return err
 	}
-	span.SetStatus(codes.Ok, "schema loaded successfully")
-	span.End()
 
 	elapsedTime := time.Now().Sub(startTime)
 	logger.Printf("preloaded schema for %s %s in %s", pAddr, pSchemaFile.Version, elapsedTime)
-	rootSpan.SetStatus(codes.Ok, "schema loaded successfully")
 
 	return nil
 }
