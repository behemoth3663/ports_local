--- vendor/github.com/anthropics/anthropic-sdk-go/betasessiontoolrunner.go.orig	2026-10-15 19:41:18 UTC
+++ vendor/github.com/anthropics/anthropic-sdk-go/betasessiontoolrunner.go
@@ -15,11 +15,8 @@ import (
 
 	"github.com/anthropics/anthropic-sdk-go/internal/sendwindow"
 	"github.com/anthropics/anthropic-sdk-go/internal/stainlessheader"
-	"github.com/anthropics/anthropic-sdk-go/internal/telemetry"
-	"github.com/anthropics/anthropic-sdk-go/internal/tracing"
 	"github.com/anthropics/anthropic-sdk-go/option"
 	"github.com/anthropics/anthropic-sdk-go/packages/param"
-	"go.opentelemetry.io/otel/attribute"
 	"golang.org/x/sync/errgroup"
 )
 
@@ -281,8 +278,7 @@ type SessionToolRunner struct {
 	// x-stainless-helper telemetry header. Immutable after construction.
 	reqOpts []option.RequestOption
 
-	// telemetry configures the OpenTelemetry span of each tool run.
-	telemetry tracing.Options
+	telemetry toolRunnerTelemetry
 
 	// ctx is the runner's internal context — a child of the caller's ctx —
 	// used by all background loops. cancel terminates them.
@@ -1152,9 +1148,7 @@ func (r *SessionToolRunner) execute(ctx context.Contex
 		}
 	}
 	sendOpts := r.reqOpts
-	if threadConversationID != "" {
-		sendOpts = append(slices.Clip(sendOpts), option.WithOpenTelemetryConversationID(threadConversationID))
-	}
+	_ = threadConversationID
 
 	rawInput, err := json.Marshal(p.input())
 	if err != nil {
@@ -1190,16 +1184,7 @@ func (r *SessionToolRunner) execute(ctx context.Contex
 		// while teardown blocks on the drain.
 		toolTimeout := r.toolTimeout()
 		toolCtx, cancel := context.WithTimeout(ctx, toolTimeout)
-		resourceIDs := []attribute.KeyValue{attribute.String("anthropic.session.id", r.sessionID)}
-		if threadID := p.sessionThreadID(); threadID != "" {
-			resourceIDs = append(resourceIDs, attribute.String("anthropic.session.thread.id", threadID))
-		}
-		// The platform's spans have the thread ID as their conversation ID, so that the tool spans share their conversation.
-		if threadConversationID != "" {
-			resourceIDs = append(resourceIDs, telemetry.GenAIConversationIDKey.String(threadConversationID))
-		}
 		// A Session event has no toolset_name.
-		toolCtx, span := telemetry.StartToolUseSpan(toolCtx, r.telemetry, toolSpanName(tool, !p.custom), name, "", tool.Description, id, resourceIDs...)
 		out, runErr := tool.Execute(toolCtx, rawInput)
 		cancel()
 		switch {
@@ -1218,9 +1203,8 @@ func (r *SessionToolRunner) execute(ctx context.Contex
 		if len(blocks) == 0 {
 			blocks = textOnlyResult("(no output)")
 		}
-		// ctx, not toolCtx: the runner's timeout on toolCtx is a failure of the tool,
-		// and toolCtx is always canceled by now.
-		telemetry.EndToolUseSpan(ctx, span, r.telemetry, rawInput, sentToolResultContent(blocks), runErr)
+		_ = sentToolResultContent(blocks)
+		_ = runErr
 	}
 	return r.postCall(ctx, call, blocks, sendOpts)
 }
