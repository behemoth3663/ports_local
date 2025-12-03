--- go.mod.orig	2020-10-24 23:50:18 UTC
+++ go.mod
@@ -1,8 +1,20 @@ module github.com/apparentlymart/terraform-clean-synta
 module github.com/apparentlymart/terraform-clean-syntax
 
-go 1.12
+go 1.24.0
 
 require (
-	github.com/hashicorp/hcl/v2 v2.5.1
-	github.com/spf13/pflag v1.0.5
+	github.com/hashicorp/hcl/v2 v2.24.0
+	github.com/spf13/pflag v1.0.10
+)
+
+require (
+	github.com/agext/levenshtein v1.2.3 // indirect
+	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
+	github.com/google/go-cmp v0.7.0 // indirect
+	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
+	github.com/zclconf/go-cty v1.17.0 // indirect
+	golang.org/x/mod v0.30.0 // indirect
+	golang.org/x/sync v0.18.0 // indirect
+	golang.org/x/text v0.31.0 // indirect
+	golang.org/x/tools v0.39.0 // indirect
 )
