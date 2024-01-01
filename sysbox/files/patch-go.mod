--- go.mod.orig	2023-07-12 03:50:15 UTC
+++ go.mod
@@ -1,13 +1,26 @@ module github.com/skx/sysbox
 module github.com/skx/sysbox
 
-go 1.16
+go 1.21
 
 require (
-	github.com/armon/go-metrics v0.5.1 // indirect
 	github.com/creack/pty v1.1.18
 	github.com/gdamore/tcell/v2 v2.6.0
-	github.com/google/btree v1.1.2 // indirect
 	github.com/google/goexpect v0.0.0-20210430020637-ab937bf7fd6f
+	github.com/hashicorp/memberlist v0.5.0
+	github.com/nightlyone/lockfile v1.0.0
+	github.com/peterh/liner v1.2.2
+	github.com/rivo/tview v0.0.0-20230621164836-6cc0565babaf
+	github.com/skx/subcommands v0.9.2
+	golang.org/x/net v0.12.0
+	golang.org/x/term v0.10.0
+	gopkg.in/yaml.v2 v2.4.0
+)
+
+require (
+	github.com/armon/go-metrics v0.5.1 // indirect
+	github.com/gdamore/encoding v1.0.0 // indirect
+	github.com/golang/protobuf v1.5.0 // indirect
+	github.com/google/btree v1.1.2 // indirect
 	github.com/google/goterm v0.0.0-20200907032337-555d40f16ae2 // indirect
 	github.com/hashicorp/errwrap v1.1.0 // indirect
 	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
@@ -16,18 +29,19 @@ require (
 	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
 	github.com/hashicorp/go-uuid v1.0.1 // indirect
 	github.com/hashicorp/golang-lru v0.5.4 // indirect
-	github.com/hashicorp/memberlist v0.5.0
+	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
+	github.com/mattn/go-runewidth v0.0.14 // indirect
 	github.com/miekg/dns v1.1.55 // indirect
-	github.com/nightlyone/lockfile v1.0.0
-	github.com/peterh/liner v1.2.2
-	github.com/rivo/tview v0.0.0-20230621164836-6cc0565babaf
 	github.com/rivo/uniseg v0.4.4 // indirect
-	github.com/skx/subcommands v0.9.2
-	golang.org/x/net v0.12.0
-	golang.org/x/term v0.10.0
+	github.com/sean-/seed v0.0.0-20170313163322-e2103e2c3529 // indirect
+	golang.org/x/crypto v0.11.0 // indirect
+	golang.org/x/mod v0.12.0 // indirect
+	golang.org/x/sys v0.10.0 // indirect
+	golang.org/x/text v0.11.0 // indirect
 	golang.org/x/tools v0.11.0 // indirect
 	google.golang.org/genproto/googleapis/rpc v0.0.0-20230711160842-782d3b101e98 // indirect
-	gopkg.in/yaml.v2 v2.4.0
+	google.golang.org/grpc v1.31.0 // indirect
+	google.golang.org/protobuf v1.31.0 // indirect
 )
 
 replace github.com/armon/go-metrics => github.com/hashicorp/go-metrics v0.5.1
