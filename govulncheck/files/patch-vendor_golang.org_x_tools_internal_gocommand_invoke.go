--- vendor/golang.org/x/tools/internal/gocommand/invoke.go.orig	2026-10-10 19:29:17 UTC
+++ vendor/golang.org/x/tools/internal/gocommand/invoke.go
@@ -246,7 +246,7 @@ func (i *Invocation) run(ctx context.Context, stdout, 
 		appendOverlayFlag()
 		goArgs = append(goArgs, i.Args...)
 	}
-	cmd := exec.Command("go", goArgs...)
+	cmd := exec.Command("%%GO_CMD%%", goArgs...)
 	setProcessGroup(cmd)
 	cmd.Stdout = stdout
 	cmd.Stderr = stderr
