--- internal/state/provider_schema.go.orig	2024-10-11 13:19:56 UTC
+++ internal/state/provider_schema.go
@@ -20,10 +20,6 @@ import (
 	"github.com/hashicorp/terraform-ls/internal/schemas"
 	tfaddr "github.com/hashicorp/terraform-registry-address"
 	tfschema "github.com/hashicorp/terraform-schema/schema"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 )
 
 type ProviderSchema struct {
@@ -522,17 +518,8 @@ func PreloadSchemaForProviderAddr(ctx context.Context,
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
@@ -540,34 +527,16 @@ func PreloadSchemaForProviderAddr(ctx context.Context,
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
@@ -577,16 +546,8 @@ func PreloadSchemaForProviderAddr(ctx context.Context,
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
 		existsError := &AlreadyExistsError{}
 		if errors.As(err, &existsError) {
 			// This accounts for a possible race condition
@@ -597,12 +558,9 @@ func PreloadSchemaForProviderAddr(ctx context.Context,
 		}
 		return err
 	}
-	span.SetStatus(codes.Ok, "schema loaded successfully")
-	span.End()
 
 	elapsedTime := time.Since(startTime)
 	logger.Printf("preloaded schema for %s %s in %s", pAddr, pSchemaFile.Version, elapsedTime)
-	rootSpan.SetStatus(codes.Ok, "schema loaded successfully")
 
 	return nil
 }
