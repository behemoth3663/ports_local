--- pkg/terminal/starbind/starlark.go.orig	2026-09-09 15:48:17 UTC
+++ pkg/terminal/starbind/starlark.go
@@ -15,7 +15,6 @@ import (
 	"go.starlark.net/starlark"
 	"go.starlark.net/syntax"
 
-	"github.com/go-delve/delve/pkg/logflags"
 	"github.com/go-delve/delve/service"
 	"github.com/go-delve/delve/service/api"
 )
@@ -226,7 +225,6 @@ func (env *Env) Execute(path string, source any, mainF
 		if err == nil {
 			return
 		}
-		logflags.Bug.Inc()
 		_err = fmt.Errorf("panic executing starlark script: %v", err)
 		fmt.Fprintf(env.out, "panic executing starlark script: %v\n", err)
 		for i := 0; ; i++ {
