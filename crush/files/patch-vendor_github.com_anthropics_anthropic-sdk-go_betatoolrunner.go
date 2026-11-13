--- vendor/github.com/anthropics/anthropic-sdk-go/betatoolrunner.go.orig	2026-10-15 19:41:18 UTC
+++ vendor/github.com/anthropics/anthropic-sdk-go/betatoolrunner.go
@@ -6,20 +6,13 @@ import (
 	"errors"
 	"fmt"
 	"iter"
-	"net/http"
 	"os"
 	"runtime"
 	"slices"
-	"sync"
 
-	"github.com/anthropics/anthropic-sdk-go/internal/requestconfig"
 	"github.com/anthropics/anthropic-sdk-go/internal/stainlessheader"
-	"github.com/anthropics/anthropic-sdk-go/internal/telemetry"
-	"github.com/anthropics/anthropic-sdk-go/internal/tracing"
 	"github.com/anthropics/anthropic-sdk-go/option"
 	"github.com/anthropics/anthropic-sdk-go/packages/param"
-	"go.opentelemetry.io/otel/codes"
-	"go.opentelemetry.io/otel/trace"
 	"golang.org/x/sync/errgroup"
 )
 
@@ -64,65 +57,19 @@ type betaToolRunnerBase struct {
 	compacting         bool
 	pendingToolChanges []BetaContentBlockParamUnion
 
-	// telemetry configures the OpenTelemetry spans of the current loop.
-	telemetry tracing.Options
+	telemetry toolRunnerTelemetry
 	loopSpan  *toolRunnerLoopSpan
 }
 
-// toolRunnerLoopSpan is the span of the current loop. It is the argument of the
-// runner's cleanup, which runs on another goroutine, hence mu, and never runs if
-// this type points back to the runner.
-type toolRunnerLoopSpan struct {
-	mu sync.Mutex
-	// span is anthropic.messages.tool_runner, or nil when no loop is running.
-	span trace.Span
-	// off is whether the current loop found the spans off: its steps do not look again.
-	off bool
-}
+type (
+	toolRunnerTelemetry struct{ ConversationID string }
+	toolRunnerLoopSpan  struct{}
+)
 
-// take returns the open span, or nil, and leaves the loop without one. mu is not held
-// while the span ends: a span processor may block, and the cleanup would wait for mu.
-func (s *toolRunnerLoopSpan) take() trace.Span {
-	s.mu.Lock()
-	defer s.mu.Unlock()
-	s.off = false
-	span := s.span
-	s.span = nil
-	return span
-}
+func (s *toolRunnerLoopSpan) end(err error) { _ = err }
+func (s *toolRunnerLoopSpan) stop()         {}
+func (s *toolRunnerLoopSpan) stopDropped()  {}
 
-// end ends the span of a loop that ended with err, or with no error. A canceled
-// context is not a failure, so it leaves the status unset.
-func (s *toolRunnerLoopSpan) end(err error) {
-	span := s.take()
-	if span == nil {
-		return
-	}
-	tracing.Guard(func() {
-		defer span.End()
-		if err != nil {
-			tracing.SetErrorStatus(span, err, "")
-		} else {
-			span.SetStatus(codes.Ok, "")
-		}
-	})
-}
-
-// stop ends the span of a loop that the application stopped. The status stays unset.
-func (s *toolRunnerLoopSpan) stop() {
-	if span := s.take(); span != nil {
-		tracing.Guard(func() { span.End() })
-	}
-}
-
-// stopDropped is the runner's cleanup. An open span ends on a goroutine of its own: a span
-// processor may block, and other cleanups would wait for it.
-func (s *toolRunnerLoopSpan) stopDropped() {
-	if span := s.take(); span != nil {
-		go tracing.Guard(func() { span.End() })
-	}
-}
-
 func newBetaToolRunnerBase(messageService *BetaMessageService, tools []BetaTool, params BetaToolRunnerParams, opts []option.RequestOption) betaToolRunnerBase {
 	toolMap := make(map[string]BetaTool)
 	apiTools := make([]BetaToolUnionParam, len(tools))
@@ -344,14 +291,10 @@ func determineNextStepFromStopReason(reason BetaStopRe
 
 // openTelemetryOptions returns the OpenTelemetry options of the requests that a
 // tool runner sends with opts, for the runner's own spans.
-func openTelemetryOptions(ctx context.Context, opts []option.RequestOption) tracing.Options {
-	// The options apply to an empty JSON request as they do to the runner's requests.
-	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, "", []byte("{}"), nil, opts...)
-	if err != nil {
-		off := false
-		return tracing.Options{Traces: tracing.TracesOptions{Enabled: &off}}
-	}
-	return cfg.OpenTelemetry.Options
+func openTelemetryOptions(ctx context.Context, opts []option.RequestOption) toolRunnerTelemetry {
+	_ = ctx
+	_ = opts
+	return toolRunnerTelemetry{}
 }
 
 // startToolRunnerSpan starts the tool runner span of a loop, when no loop is
@@ -365,21 +308,8 @@ func (b *betaToolRunnerBase) startToolRunnerSpan(ctx c
 // inside the tool runner span, whichever context the step was called with. The
 // application's own context is never changed.
 func (b *betaToolRunnerBase) startToolRunnerSpan(ctx context.Context) context.Context {
-	b.loopSpan.mu.Lock()
-	defer b.loopSpan.mu.Unlock()
-	if b.completed || b.loopSpan.off {
-		return ctx
-	}
-	if b.loopSpan.span == nil {
-		b.telemetry = openTelemetryOptions(ctx, slices.Concat(b.messageService.Options, b.opts))
-		hasOutputFormat := !param.IsOmitted(b.Params.OutputConfig.Format) || !param.IsOmitted(b.Params.OutputFormat)
-		b.loopSpan.span = telemetry.StartToolRunnerSpan(ctx, b.telemetry, string(b.Params.Model), hasOutputFormat)
-		b.loopSpan.off = b.loopSpan.span == nil
-	}
-	if b.loopSpan.span == nil {
-		return ctx
-	}
-	return trace.ContextWithSpan(ctx, b.loopSpan.span)
+	b.telemetry = openTelemetryOptions(ctx, slices.Concat(b.messageService.Options, b.opts))
+	return ctx
 }
 
 // executeTools processes any tool use blocks in the given message and returns a tool result message.
@@ -564,13 +494,12 @@ func (b *betaToolRunnerBase) executeToolUse(ctx contex
 		)
 	}
 
-	ctx, span := telemetry.StartToolUseSpan(ctx, b.telemetry, toolSpanName(tool, false), toolUse.Name, toolUse.ToolsetName, tool.Description, toolUse.ID)
 	content, err := tool.Execute(ctx, inputBytes)
 	result := BetaToolResultBlockParam{ToolUseID: toolUse.ID, Content: content}
 	if err != nil {
 		result = newBetaToolResultErrorBlockParam(toolUse.ID, fmt.Sprintf("Error: %v", err))
 	}
-	telemetry.EndToolUseSpan(ctx, span, b.telemetry, inputBytes, result.Content, err)
+	_ = inputBytes
 	return result
 }
 
@@ -578,11 +507,9 @@ func toolSpanName(tool BetaTool, agentTool bool) strin
 // starts with, which says where tool comes from. agentTool is whether an
 // agent.tool_use Session event calls it.
 func toolSpanName(tool BetaTool, agentTool bool) string {
+	_ = tool
 	if agentTool {
 		return "anthropic.tool_use"
-	}
-	if _, ok := tool.(telemetry.MCPHelperTool); ok {
-		return "anthropic.mcp_tool_use"
 	}
 	return "anthropic.custom_tool_use"
 }
