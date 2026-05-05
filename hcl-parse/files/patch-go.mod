--- go.mod.orig	2024-07-27 12:50:44 UTC
+++ go.mod
@@ -1,17 +1,16 @@ module github.com/terraform-linters/hcl-parse
 module github.com/terraform-linters/hcl-parse
 
-go 1.20
+go 1.25.0
 
-require github.com/hashicorp/hcl/v2 v2.21.0
+require github.com/hashicorp/hcl/v2 v2.24.0
 
 require (
-	github.com/agext/levenshtein v1.2.1 // indirect
-	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
+	github.com/agext/levenshtein v1.2.3 // indirect
 	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
-	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
-	github.com/zclconf/go-cty v1.13.0 // indirect
-	golang.org/x/mod v0.8.0 // indirect
-	golang.org/x/sys v0.5.0 // indirect
-	golang.org/x/text v0.11.0 // indirect
-	golang.org/x/tools v0.6.0 // indirect
+	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
+	github.com/zclconf/go-cty v1.18.1 // indirect
+	golang.org/x/mod v0.35.0 // indirect
+	golang.org/x/sync v0.20.0 // indirect
+	golang.org/x/text v0.36.0 // indirect
+	golang.org/x/tools v0.44.0 // indirect
 )
