--- pkg/dwarf/line/line_parser.go.orig	2026-09-09 15:48:17 UTC
+++ pkg/dwarf/line/line_parser.go
@@ -8,7 +8,6 @@ import (
 
 	"github.com/go-delve/delve/pkg/dwarf"
 	"github.com/go-delve/delve/pkg/dwarf/leb128"
-	"github.com/go-delve/delve/pkg/logflags"
 )
 
 // DebugLinePrologue prologue of .debug_line data.
@@ -195,7 +194,6 @@ func parseIncludeDirs5(info *DebugLineInfo, buf *bytes
 						info.IncludeDirs = append(info.IncludeDirs, "<DW_FORM_line_strp without debug_line_str section>")
 						if first_DW_FORM_line_strp_bug {
 							first_DW_FORM_line_strp_bug = false
-							logflags.Bug.Inc()
 						}
 					} else {
 						buf := bytes.NewBuffer(info.debugLineStr[dirEntryFormReader.u64:])
