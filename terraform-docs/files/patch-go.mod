--- go.mod.orig	2023-12-19 18:31:44 UTC
+++ go.mod
@@ -1,6 +1,6 @@ module github.com/terraform-docs/terraform-docs
 module github.com/terraform-docs/terraform-docs
 
-go 1.16
+go 1.21
 
 require (
 	github.com/BurntSushi/toml v1.3.2
@@ -23,18 +23,50 @@ require (
 )
 
 require (
+	github.com/Masterminds/goutils v1.1.1 // indirect
 	github.com/Masterminds/semver/v3 v3.2.1 // indirect
 	github.com/agext/levenshtein v1.2.3 // indirect
+	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
+	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
 	github.com/fatih/color v1.16.0 // indirect
+	github.com/fsnotify/fsnotify v1.7.0 // indirect
+	github.com/golang/protobuf v1.5.3 // indirect
+	github.com/google/go-cmp v0.5.9 // indirect
+	github.com/google/uuid v1.4.0 // indirect
+	github.com/hashicorp/hcl v1.0.0 // indirect
+	github.com/hashicorp/yamux v0.1.1 // indirect
 	github.com/huandu/xstrings v1.4.0 // indirect
+	github.com/inconshreveable/mousetrap v1.1.0 // indirect
+	github.com/magiconair/properties v1.8.7 // indirect
+	github.com/mattn/go-colorable v0.1.13 // indirect
+	github.com/mattn/go-isatty v0.0.20 // indirect
 	github.com/mitchellh/copystructure v1.2.0 // indirect
 	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
 	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
+	github.com/mitchellh/mapstructure v1.5.0 // indirect
+	github.com/mitchellh/reflectwalk v1.0.2 // indirect
 	github.com/oklog/run v1.1.0 // indirect
+	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
+	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
+	github.com/sagikazarmark/locafero v0.4.0 // indirect
+	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
 	github.com/shopspring/decimal v1.3.1 // indirect
+	github.com/sourcegraph/conc v0.3.0 // indirect
+	github.com/spf13/afero v1.11.0 // indirect
+	github.com/spf13/cast v1.6.0 // indirect
+	github.com/subosito/gotenv v1.6.0 // indirect
 	github.com/zclconf/go-cty v1.14.1 // indirect
 	go.uber.org/multierr v1.11.0 // indirect
+	golang.org/x/crypto v0.16.0 // indirect
 	golang.org/x/exp v0.0.0-20231206192017-f3f8817b8deb // indirect
 	golang.org/x/exp/typeparams v0.0.0-20220722155223-a9213eeb770e // indirect
+	golang.org/x/mod v0.14.0 // indirect
+	golang.org/x/net v0.19.0 // indirect
+	golang.org/x/sys v0.15.0 // indirect
+	golang.org/x/text v0.14.0 // indirect
+	golang.org/x/tools v0.16.0 // indirect
 	google.golang.org/genproto/googleapis/rpc v0.0.0-20231127180814-3a041ad873d4 // indirect
+	google.golang.org/grpc v1.59.0 // indirect
+	google.golang.org/protobuf v1.31.0 // indirect
+	gopkg.in/ini.v1 v1.67.0 // indirect
 )
