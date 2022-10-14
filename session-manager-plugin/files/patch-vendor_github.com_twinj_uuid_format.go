--- vendor/github.com/twinj/uuid/format.go.orig	2017-05-27 12:27:00 UTC
+++ vendor/github.com/twinj/uuid/format.go
@@ -16,6 +16,7 @@ const (
 
 	// FormatCanonical is the default format.
 	FormatCanonical Format = "%x-%x-%x-%x-%x"
+	CleanHyphen     Format = "%x-%x-%x-%x-%x"
 
 	FormatCanonicalCurly   Format = "{%x-%x-%x-%x-%x}"
 	FormatCanonicalBracket Format = "(%x-%x-%x-%x-%x)"
