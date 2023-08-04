--- internal/usage/aws/dynamodb.go.orig	2023-07-21 14:54:00 UTC
+++ internal/usage/aws/dynamodb.go
@@ -45,7 +45,7 @@ func DynamoDBGetStorageBytes(ctx context.Context, regi
 	if err != nil {
 		return 0, err
 	}
-	return result.Table.TableSizeBytes, nil
+	return *result.Table.TableSizeBytes, nil
 }
 
 func DynamoDBGetRRU(ctx context.Context, region string, table string) (float64, error) {
