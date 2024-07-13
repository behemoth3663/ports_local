--- go.mod.orig	2024-02-22 11:32:48 UTC
+++ go.mod
@@ -30,11 +30,11 @@ require (
 	github.com/go-test/deep v1.0.3
 	github.com/golang/mock v1.6.0
 	github.com/google/go-cmp v0.6.0
-	github.com/google/uuid v1.3.1
+	github.com/google/uuid v1.4.0
 	github.com/hashicorp/aws-sdk-go-base/v2 v2.0.0-beta.43
 	github.com/hashicorp/consul/api v1.13.0
 	github.com/hashicorp/consul/sdk v0.8.0
-	github.com/hashicorp/copywrite v0.16.3
+	github.com/hashicorp/copywrite v0.16.6
 	github.com/hashicorp/errwrap v1.1.0
 	github.com/hashicorp/go-azure-helpers v0.43.0
 	github.com/hashicorp/go-cleanhttp v0.5.2
@@ -55,7 +55,7 @@ require (
 	github.com/lib/pq v1.10.3
 	github.com/manicminer/hamilton v0.44.0
 	github.com/masterzen/winrm v0.0.0-20200615185753-c42b5136ff88
-	github.com/mattn/go-isatty v0.0.17
+	github.com/mattn/go-isatty v0.0.20
 	github.com/mattn/go-shellwords v1.0.4
 	github.com/mitchellh/cli v1.1.5
 	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db
@@ -72,7 +72,7 @@ require (
 	github.com/pkg/browser v0.0.0-20201207095918-0426ae3fba23
 	github.com/pkg/errors v0.9.1
 	github.com/posener/complete v1.2.3
-	github.com/spf13/afero v1.9.3
+	github.com/spf13/afero v1.9.5
 	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.588
 	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts v1.0.588
 	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tag v1.0.233
@@ -87,21 +87,21 @@ require (
 	go.opentelemetry.io/otel v1.21.0
 	go.opentelemetry.io/otel/sdk v1.21.0
 	go.opentelemetry.io/otel/trace v1.21.0
-	golang.org/x/crypto v0.16.0
-	golang.org/x/exp v0.0.0-20230905200255-921286631fa9
-	golang.org/x/mod v0.12.0
-	golang.org/x/net v0.19.0
+	golang.org/x/crypto v0.24.0
+	golang.org/x/exp v0.0.0-20240613232115-7f521ea00fb8
+	golang.org/x/mod v0.18.0
+	golang.org/x/net v0.26.0
 	golang.org/x/oauth2 v0.11.0
-	golang.org/x/sys v0.15.0
-	golang.org/x/term v0.15.0
-	golang.org/x/text v0.14.0
-	golang.org/x/tools v0.13.0
+	golang.org/x/sys v0.21.0
+	golang.org/x/term v0.21.0
+	golang.org/x/text v0.16.0
+	golang.org/x/tools v0.22.0
 	google.golang.org/api v0.126.0
 	google.golang.org/genproto v0.0.0-20230822172742-b8732ec3820d
 	google.golang.org/grpc v1.59.0
 	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
-	google.golang.org/protobuf v1.31.0
-	honnef.co/go/tools v0.4.2
+	google.golang.org/protobuf v1.33.0
+	honnef.co/go/tools v0.4.7
 	k8s.io/api v0.23.4
 	k8s.io/apimachinery v0.23.4
 	k8s.io/client-go v0.23.4
@@ -113,7 +113,7 @@ require (
 	cloud.google.com/go/compute v1.23.0 // indirect
 	cloud.google.com/go/compute/metadata v0.2.3 // indirect
 	cloud.google.com/go/iam v1.1.1 // indirect
-	github.com/AlecAivazis/survey/v2 v2.3.6 // indirect
+	github.com/AlecAivazis/survey/v2 v2.3.7 // indirect
 	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
 	github.com/Azure/go-autorest/autorest/adal v0.9.18 // indirect
 	github.com/Azure/go-autorest/autorest/azure/cli v0.4.4 // indirect
@@ -134,7 +134,7 @@ require (
 	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
 	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da // indirect
 	github.com/armon/go-radix v1.0.0 // indirect
-	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef // indirect
+	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
 	github.com/aws/aws-sdk-go v1.44.122 // indirect
 	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.5.1 // indirect
 	github.com/aws/aws-sdk-go-v2/config v1.25.8 // indirect
@@ -155,13 +155,13 @@ require (
 	github.com/aws/aws-sdk-go-v2/service/sts v1.25.6 // indirect
 	github.com/aws/smithy-go v1.17.0 // indirect
 	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d // indirect
-	github.com/bmatcuk/doublestar/v4 v4.6.0 // indirect
-	github.com/bradleyfalzon/ghinstallation/v2 v2.1.0 // indirect
+	github.com/bmatcuk/doublestar/v4 v4.6.1 // indirect
+	github.com/bradleyfalzon/ghinstallation/v2 v2.5.0 // indirect
 	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
-	github.com/cli/go-gh v1.0.0 // indirect
-	github.com/cli/safeexec v1.0.0 // indirect
-	github.com/cli/shurcooL-graphql v0.0.2 // indirect
-	github.com/cloudflare/circl v1.3.3 // indirect
+	github.com/cli/go-gh v1.2.1 // indirect
+	github.com/cli/safeexec v1.0.1 // indirect
+	github.com/cli/shurcooL-graphql v0.0.4 // indirect
+	github.com/cloudflare/circl v1.3.9 // indirect
 	github.com/coreos/go-systemd v0.0.0-20181012123002-c6f51f82210d // indirect
 	github.com/creack/pty v1.1.18 // indirect
 	github.com/dimchansky/utfbom v1.1.1 // indirect
@@ -170,14 +170,15 @@ require (
 	github.com/fsnotify/fsnotify v1.5.4 // indirect
 	github.com/go-logr/logr v1.3.0 // indirect
 	github.com/go-logr/stdr v1.2.2 // indirect
-	github.com/go-openapi/errors v0.20.2 // indirect
-	github.com/go-openapi/strfmt v0.21.3 // indirect
+	github.com/go-openapi/errors v0.21.1 // indirect
+	github.com/go-openapi/strfmt v0.21.10 // indirect
 	github.com/gofrs/uuid v4.0.0+incompatible // indirect
 	github.com/gogo/protobuf v1.3.2 // indirect
-	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
+	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
 	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
-	github.com/golang/protobuf v1.5.3 // indirect
+	github.com/golang/protobuf v1.5.4 // indirect
 	github.com/google/go-github/v45 v45.2.0 // indirect
+	github.com/google/go-github/v53 v53.0.0 // indirect
 	github.com/google/go-querystring v1.1.0 // indirect
 	github.com/google/gofuzz v1.1.0 // indirect
 	github.com/google/s2a-go v0.1.4 // indirect
@@ -199,7 +200,7 @@ require (
 	github.com/imdario/mergo v0.3.13 // indirect
 	github.com/inconshreveable/mousetrap v1.0.1 // indirect
 	github.com/jedib0t/go-pretty v4.3.0+incompatible // indirect
-	github.com/jedib0t/go-pretty/v6 v6.4.4 // indirect
+	github.com/jedib0t/go-pretty/v6 v6.4.9 // indirect
 	github.com/joho/godotenv v1.3.0 // indirect
 	github.com/json-iterator/go v1.1.12 // indirect
 	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
@@ -209,7 +210,7 @@ require (
 	github.com/manicminer/hamilton-autorest v0.2.0 // indirect
 	github.com/masterzen/simplexml v0.0.0-20190410153822-31eea3082786 // indirect
 	github.com/mattn/go-colorable v0.1.13 // indirect
-	github.com/mattn/go-runewidth v0.0.13 // indirect
+	github.com/mattn/go-runewidth v0.0.15 // indirect
 	github.com/mergestat/timediff v0.0.3 // indirect
 	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
 	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
@@ -228,12 +229,12 @@ require (
 	github.com/spf13/cast v1.5.0 // indirect
 	github.com/spf13/cobra v1.6.1 // indirect
 	github.com/spf13/pflag v1.0.5 // indirect
-	github.com/thanhpk/randstr v1.0.4 // indirect
-	github.com/thlib/go-timezone-local v0.0.0-20210907160436-ef149e42d28e // indirect
+	github.com/thanhpk/randstr v1.0.6 // indirect
+	github.com/thlib/go-timezone-local v0.0.2 // indirect
 	github.com/ulikunitz/xz v0.5.10 // indirect
 	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
 	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
-	go.mongodb.org/mongo-driver v1.11.6 // indirect
+	go.mongodb.org/mongo-driver v1.13.4 // indirect
 	go.opencensus.io v0.24.0 // indirect
 	go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws v0.46.1 // indirect
 	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.21.0 // indirect
@@ -241,11 +242,11 @@ require (
 	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.21.0 // indirect
 	go.opentelemetry.io/otel/metric v1.21.0 // indirect
 	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
-	golang.org/x/exp/typeparams v0.0.0-20221208152030-732eee02a75a // indirect
-	golang.org/x/sync v0.4.0 // indirect
+	golang.org/x/exp/typeparams v0.0.0-20240613232115-7f521ea00fb8 // indirect
+	golang.org/x/sync v0.7.0 // indirect
 	golang.org/x/time v0.3.0 // indirect
 	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
-	google.golang.org/appengine v1.6.7 // indirect
+	google.golang.org/appengine v1.6.8 // indirect
 	google.golang.org/genproto/googleapis/api v0.0.0-20230822172742-b8732ec3820d // indirect
 	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
 	gopkg.in/inf.v0 v0.9.1 // indirect
