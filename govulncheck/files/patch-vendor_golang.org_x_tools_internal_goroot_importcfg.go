--- vendor/golang.org/x/tools/internal/goroot/importcfg.go.orig	2026-04-09 21:00:41 UTC
+++ vendor/golang.org/x/tools/internal/goroot/importcfg.go
@@ -25,7 +25,7 @@ var pkgFileMapOnce = sync.OnceValues(func() (map[strin
 
 var pkgFileMapOnce = sync.OnceValues(func() (map[string]string, error) {
 	m := make(map[string]string)
-	output, err := exec.Command("go", "list", "-export", "-e", "-f", "{{.ImportPath}} {{.Export}}", "std", "cmd").Output()
+	output, err := exec.Command("%%GO_CMD%%", "list", "-export", "-e", "-f", "{{.ImportPath}} {{.Export}}", "std", "cmd").Output()
 	if err != nil {
 		return m, err
 	}
