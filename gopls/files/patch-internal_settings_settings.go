--- internal/settings/settings.go.orig	2025-02-24 17:39:34 UTC
+++ internal/settings/settings.go
@@ -14,7 +14,6 @@ import (
 	"golang.org/x/tools/gopls/internal/file"
 	"golang.org/x/tools/gopls/internal/protocol"
 	"golang.org/x/tools/gopls/internal/protocol/semtok"
-	"golang.org/x/tools/gopls/internal/telemetry"
 	"golang.org/x/tools/gopls/internal/util/frob"
 )
 
@@ -828,7 +827,7 @@ const (
 	// TODO: support "Manual"?
 )
 
-type CounterPath = telemetry.CounterPath
+type CounterPath = []string
 
 // Set updates *Options based on the provided JSON value:
 // null, bool, string, number, array, or object.
