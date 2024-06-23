--- vendor/github.com/aws/session-manager-plugin/src/message/messageparser.go.orig	2022-10-12 15:59:18 UTC
+++ vendor/github.com/aws/session-manager-plugin/src/message/messageparser.go
@@ -167,31 +167,31 @@ func getUuid(log log.T, byteArray []byte, offset int) 
 	byteArrayLength := len(byteArray)
 	if offset > byteArrayLength-1 || offset+16-1 > byteArrayLength-1 || offset < 0 {
 		log.Error("getUuid failed: Offset is invalid.")
-		return nil, errors.New("Offset is outside the byte array.")
+		return uuid.UUID{}, errors.New("Offset is outside the byte array.")
 	}
 
 	leastSignificantLong, err := getLong(log, byteArray, offset)
 	if err != nil {
 		log.Error("getUuid failed: failed to get uuid LSBs Long value.")
-		return nil, errors.New("Failed to get uuid LSBs long value.")
+		return uuid.UUID{}, errors.New("Failed to get uuid LSBs long value.")
 	}
 
 	leastSignificantBytes, err := longToBytes(log, leastSignificantLong)
 	if err != nil {
 		log.Error("getUuid failed: failed to get uuid LSBs bytes value.")
-		return nil, errors.New("Failed to get uuid LSBs bytes value.")
+		return uuid.UUID{}, errors.New("Failed to get uuid LSBs bytes value.")
 	}
 
 	mostSignificantLong, err := getLong(log, byteArray, offset+8)
 	if err != nil {
 		log.Error("getUuid failed: failed to get uuid MSBs Long value.")
-		return nil, errors.New("Failed to get uuid MSBs long value.")
+		return uuid.UUID{}, errors.New("Failed to get uuid MSBs long value.")
 	}
 
 	mostSignificantBytes, err := longToBytes(log, mostSignificantLong)
 	if err != nil {
 		log.Error("getUuid failed: failed to get uuid MSBs bytes value.")
-		return nil, errors.New("Failed to get uuid MSBs bytes value.")
+		return uuid.UUID{}, errors.New("Failed to get uuid MSBs bytes value.")
 	}
 
 	uuidBytes := append(mostSignificantBytes, leastSignificantBytes...)
@@ -414,11 +414,6 @@ func putBytes(log log.T, byteArray []byte, offsetStart
 
 // putUuid puts the 128 bit uuid to an array of bytes starting from the offset.
 func putUuid(log log.T, byteArray []byte, offset int, input uuid.UUID) (err error) {
-	if input == nil {
-		log.Error("putUuid failed: input is null.")
-		return errors.New("putUuid failed: input is null.")
-	}
-
 	byteArrayLength := len(byteArray)
 	if offset > byteArrayLength-1 || offset+16-1 > byteArrayLength-1 || offset < 0 {
 		log.Error("putUuid failed: Offset is invalid.")
