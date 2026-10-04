--- cli/debug/debug.go.orig	2026-09-23 09:45:10 UTC
+++ cli/debug/debug.go
@@ -4,7 +4,6 @@ import (
 	"os"
 
 	"github.com/sirupsen/logrus"
-	"go.opentelemetry.io/otel"
 )
 
 // Enable sets the DEBUG env var to true
@@ -26,15 +25,10 @@ func IsEnabled() bool {
 	return os.Getenv("DEBUG") != ""
 }
 
-// OTELErrorHandler is an error handler for OTEL that
-// uses the CLI debug package to log messages when an error
-// occurs.
-//
-// The default is to log to the debug level which is only
-// enabled when debugging is enabled.
-var OTELErrorHandler otel.ErrorHandler = otel.ErrorHandlerFunc(func(err error) {
+// LogTelemetryError logs telemetry-related errors at debug-level.
+func LogTelemetryError(err error) {
 	if err == nil {
 		return
 	}
 	logrus.WithError(err).Debug("otel error")
-})
+}
