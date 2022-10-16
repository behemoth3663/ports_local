--- go.mod.orig	2022-09-01 05:27:20 UTC
+++ go.mod
@@ -1,13 +1,26 @@
 module github.com/skx/sysbox
 
-go 1.16
+go 1.18
 
 require (
-	github.com/armon/go-metrics v0.4.0 // indirect
 	github.com/creack/pty v1.1.18
 	github.com/gdamore/tcell/v2 v2.5.3
-	github.com/google/btree v1.1.2 // indirect
 	github.com/google/goexpect v0.0.0-20210430020637-ab937bf7fd6f
+	github.com/hashicorp/memberlist v0.4.0
+	github.com/nightlyone/lockfile v1.0.0
+	github.com/peterh/liner v1.2.2
+	github.com/rivo/tview v0.0.0-20220812085834-0e6b21a48e96
+	github.com/skx/subcommands v0.9.2
+	golang.org/x/net v0.0.0-20220826154423-83b083e8dc8b
+	golang.org/x/term v0.0.0-20220722155259-a9ba230a4035
+	gopkg.in/yaml.v2 v2.4.0
+)
+
+require (
+	github.com/armon/go-metrics v0.4.0 // indirect
+	github.com/gdamore/encoding v1.0.0 // indirect
+	github.com/golang/protobuf v1.5.2 // indirect
+	github.com/google/btree v1.1.2 // indirect
 	github.com/google/goterm v0.0.0-20200907032337-555d40f16ae2 // indirect
 	github.com/hashicorp/errwrap v1.1.0 // indirect
 	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
@@ -16,22 +29,21 @@ require (
 	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
 	github.com/hashicorp/go-uuid v1.0.1 // indirect
 	github.com/hashicorp/golang-lru v0.5.4 // indirect
-	github.com/hashicorp/memberlist v0.4.0
 	github.com/kr/pretty v0.3.0 // indirect
+	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
+	github.com/mattn/go-runewidth v0.0.13 // indirect
 	github.com/miekg/dns v1.1.50 // indirect
-	github.com/nightlyone/lockfile v1.0.0
-	github.com/peterh/liner v1.2.2
-	github.com/rivo/tview v0.0.0-20220812085834-0e6b21a48e96
+	github.com/rivo/uniseg v0.3.4 // indirect
 	github.com/rogpeppe/go-internal v1.8.0 // indirect
-	github.com/skx/subcommands v0.9.2
+	github.com/sean-/seed v0.0.0-20170313163322-e2103e2c3529 // indirect
 	github.com/stretchr/testify v1.7.1 // indirect
 	golang.org/x/crypto v0.0.0-20220829220503-c86fa9a7ed90 // indirect
-	golang.org/x/net v0.0.0-20220826154423-83b083e8dc8b
+	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
 	golang.org/x/sys v0.0.0-20220829200755-d48e67d00261 // indirect
-	golang.org/x/term v0.0.0-20220722155259-a9ba230a4035
+	golang.org/x/text v0.3.7 // indirect
 	golang.org/x/tools v0.1.12 // indirect
 	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
 	google.golang.org/genproto v0.0.0-20220829175752-36a9c930ecbf // indirect
 	google.golang.org/grpc v1.49.0 // indirect
-	gopkg.in/yaml.v2 v2.4.0
+	google.golang.org/protobuf v1.28.1 // indirect
 )
