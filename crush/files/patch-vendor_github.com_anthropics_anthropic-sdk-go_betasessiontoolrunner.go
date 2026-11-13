--- vendor/github.com/anthropics/anthropic-sdk-go/betasessiontoolrunner.go.orig	2026-11-13 20:45:47 UTC
+++ vendor/github.com/anthropics/anthropic-sdk-go/betasessiontoolrunner.go
@@ -8,18 +8,14 @@ import (
 	"iter"
 	"log/slog"
 	"math/rand/v2"
-	"slices"
 	"strings"
 	"sync"
 	"time"
 
 	"github.com/anthropics/anthropic-sdk-go/internal/sendwindow"
 	"github.com/anthropics/anthropic-sdk-go/internal/stainlessheader"
-	"github.com/anthropics/anthropic-sdk-go/internal/telemetry"
-	"github.com/anthropics/anthropic-sdk-go/internal/tracing"
 	"github.com/anthropics/anthropic-sdk-go/option"
 	"github.com/anthropics/anthropic-sdk-go/packages/param"
-	"go.opentelemetry.io/otel/attribute"
 	"golang.org/x/sync/errgroup"
 )
 
@@ -284,9 +280,6 @@ type SessionToolRunner struct {
 	// x-stainless-helper telemetry header. Immutable after construction.
 	reqOpts []option.RequestOption
 
-	// telemetry configures the OpenTelemetry span of each tool run.
-	telemetry tracing.Options
-
 	// ctx is the runner's internal context — a child of the caller's ctx —
 	// used by all background loops. cancel terminates them.
 	ctx    context.Context
@@ -397,7 +390,6 @@ func (r *BetaSessionEventService) NewToolRunner(ctx co
 		log:          log,
 		byName:       byName,
 		reqOpts:      reqOpts,
-		telemetry:    openTelemetryOptions(ctx, slices.Concat(r.Options, reqOpts)),
 		ctx:          internalCtx,
 		cancel:       cancel,
 		sendWindow:   sendwindow.FromContext(ctx),
@@ -1196,16 +1188,7 @@ func (r *SessionToolRunner) execute(ctx context.Contex
 	// of the thread that the tool call belongs to: the event's thread, or the
 	// primary thread for an event without one. The span of the request that sends
 	// the tool's result has it, as the tool's span does.
-	threadConversationID := ""
-	if r.telemetry.ConversationID == "" {
-		if threadConversationID = p.sessionThreadID(); threadConversationID == "" {
-			threadConversationID = r.primaryThread()
-		}
-	}
 	sendOpts := r.reqOpts
-	if threadConversationID != "" {
-		sendOpts = append(slices.Clip(sendOpts), option.WithOpenTelemetryConversationID(threadConversationID))
-	}
 
 	rawInput, err := json.Marshal(p.input())
 	if err != nil {
@@ -1241,16 +1224,6 @@ func (r *SessionToolRunner) execute(ctx context.Contex
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
-		// A Session event has no toolset_name.
-		toolCtx, span := telemetry.StartToolUseSpan(toolCtx, r.telemetry, toolSpanName(tool, !p.custom), name, "", tool.Description, id, resourceIDs...)
 		r.setRunning(p.processedAt(), cancel)
 		out, runErr := tool.Execute(toolCtx, rawInput)
 		r.setRunning(time.Time{}, nil)
@@ -1260,7 +1233,6 @@ func (r *SessionToolRunner) execute(ctx context.Contex
 			// would reject a result with a permanent 4xx, so post nothing. The
 			// run was stopped, not failed: its span ends as a canceled run's.
 			log.Info("tool call interrupted; not posting a result")
-			telemetry.EndToolUseSpan(ctx, span, r.telemetry, rawInput, nil, context.Canceled)
 			return call
 		}
 		switch {
@@ -1279,9 +1251,8 @@ func (r *SessionToolRunner) execute(ctx context.Contex
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
