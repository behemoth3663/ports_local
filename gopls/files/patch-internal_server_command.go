--- internal/server/command.go.orig	2025-12-05 20:12:17 UTC
+++ internal/server/command.go
@@ -25,7 +25,6 @@ import (
 
 	"github.com/fatih/gomodifytags/modifytags"
 	"golang.org/x/mod/modfile"
-	"golang.org/x/telemetry/counter"
 	"golang.org/x/tools/go/ast/astutil"
 	"golang.org/x/tools/gopls/internal/cache"
 	"golang.org/x/tools/gopls/internal/cache/metadata"
@@ -71,9 +70,6 @@ func (s *server) ExecuteCommand(ctx context.Context, p
 			Source string `json:"source"`
 		}
 		if err := json.Unmarshal(last, &meta); err == nil && meta.Source != "" {
-			commandName := strings.TrimPrefix(params.Command, "gopls.")
-			counterName := fmt.Sprintf("gopls/cmd/source:%s-%s", commandName, meta.Source)
-			counter.New(counterName).Inc()
 			params.Arguments = params.Arguments[:len(params.Arguments)-1]
 		}
 	}
@@ -292,7 +288,6 @@ func (*commandHandler) AddTelemetryCounters(_ context.
 		if n == "" || v < 0 {
 			continue
 		}
-		counter.Add("fwd/"+n, v)
 	}
 	return nil
 }
@@ -1641,7 +1636,6 @@ func (c *commandHandler) ChangeSignature(ctx context.C
 }
 
 func (c *commandHandler) ChangeSignature(ctx context.Context, args command.ChangeSignatureArgs) (*protocol.WorkspaceEdit, error) {
-	countChangeSignature.Inc()
 	var result *protocol.WorkspaceEdit
 	err := c.run(ctx, commandConfig{
 		forURI: args.Location.URI,
@@ -1838,10 +1832,8 @@ func (c *commandHandler) ModifyTags(ctx context.Contex
 		// Each command involves either adding or removing tags, depending on
 		// whether Add or Clear is set.
 		if args.Add != "" {
-			countAddStructTags.Inc()
 			m.Add = strings.Split(args.Add, ",")
 		} else if args.Clear {
-			countRemoveStructTags.Inc()
 		}
 		if args.AddOptions != "" {
 			if options, err := optionsStringToMap(args.AddOptions); err != nil {
