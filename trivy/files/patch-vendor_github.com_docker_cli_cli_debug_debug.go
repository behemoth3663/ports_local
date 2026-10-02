--- vendor/github.com/docker/cli/cli/debug/debug.go.orig	2026-10-02 12:18:10 UTC
+++ vendor/github.com/docker/cli/cli/debug/debug.go
@@ -4,7 +4,6 @@ import (
 	"os"
 
 	"github.com/sirupsen/logrus"
-	"go.opentelemetry.io/otel"
 )
 
 // Enable sets the DEBUG env var to true
@@ -32,9 +31,9 @@ func IsEnabled() bool {
 //
 // The default is to log to the debug level which is only
 // enabled when debugging is enabled.
-var OTELErrorHandler otel.ErrorHandler = otel.ErrorHandlerFunc(func(err error) {
+var OTELErrorHandler = func(err error) {
 	if err == nil {
 		return
 	}
 	logrus.WithError(err).Debug("otel error")
-})
+}
