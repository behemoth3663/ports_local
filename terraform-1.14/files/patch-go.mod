--- go.mod.orig	2026-04-20 17:54:01 UTC
+++ go.mod
@@ -1,21 +1,21 @@
 module github.com/hashicorp/terraform
 
-go 1.25.8
+go 1.26.0
 
 godebug winsymlink=0
 
 require (
 	github.com/Netflix/go-expect v0.0.0-20220104043353-73e0943537d2
-	github.com/ProtonMail/go-crypto v1.1.3
+	github.com/ProtonMail/go-crypto v1.1.6
 	github.com/agext/levenshtein v1.2.3
-	github.com/apparentlymart/go-cidr v1.1.0
+	github.com/apparentlymart/go-cidr v1.1.1
 	github.com/apparentlymart/go-shquot v0.0.1
 	github.com/apparentlymart/go-userdirs v0.0.0-20200915174352-b0c018a67c13
-	github.com/apparentlymart/go-versions v1.0.2
+	github.com/apparentlymart/go-versions v1.0.3
 	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2
 	github.com/bgentry/speakeasy v0.1.0
 	github.com/bmatcuk/doublestar v1.1.5
-	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e
+	github.com/chzyer/readline v1.5.1
 	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc
 	github.com/dylanmei/winrmtest v0.0.0-20210303004826-fbc9ae56efb6
 	github.com/go-test/deep v1.0.3
@@ -24,33 +24,33 @@
 	github.com/hashicorp/cli v1.1.7
 	github.com/hashicorp/go-checkpoint v0.5.0
 	github.com/hashicorp/go-cleanhttp v0.5.2
-	github.com/hashicorp/go-getter v1.8.6
+	github.com/hashicorp/go-getter v1.8.9
 	github.com/hashicorp/go-hclog v1.6.3
 	github.com/hashicorp/go-plugin v1.7.0
 	github.com/hashicorp/go-retryablehttp v0.7.8
-	github.com/hashicorp/go-slug v0.18.1
+	github.com/hashicorp/go-slug v0.18.4
 	github.com/hashicorp/go-tfe v1.94.0
-	github.com/hashicorp/go-uuid v1.0.3
-	github.com/hashicorp/go-version v1.8.0
+	github.com/hashicorp/go-uuid v1.0.4
+	github.com/hashicorp/go-version v1.9.0
 	github.com/hashicorp/hcl v1.0.0
 	github.com/hashicorp/hcl/v2 v2.24.0
 	github.com/hashicorp/jsonapi v1.4.3-0.20250220162346-81a76b606f3e
 	github.com/hashicorp/terraform-registry-address v0.4.0
 	github.com/hashicorp/terraform-svchost v0.1.1
-	github.com/hashicorp/terraform/internal/backend/remote-state/azure v0.0.0-00010101000000-000000000000
-	github.com/hashicorp/terraform/internal/backend/remote-state/consul v0.0.0-00010101000000-000000000000
-	github.com/hashicorp/terraform/internal/backend/remote-state/cos v0.0.0-00010101000000-000000000000
-	github.com/hashicorp/terraform/internal/backend/remote-state/gcs v0.0.0-00010101000000-000000000000
-	github.com/hashicorp/terraform/internal/backend/remote-state/kubernetes v0.0.0-00010101000000-000000000000
-	github.com/hashicorp/terraform/internal/backend/remote-state/oci v0.0.0-00010101000000-000000000000
-	github.com/hashicorp/terraform/internal/backend/remote-state/oss v0.0.0-00010101000000-000000000000
-	github.com/hashicorp/terraform/internal/backend/remote-state/pg v0.0.0-00010101000000-000000000000
-	github.com/hashicorp/terraform/internal/backend/remote-state/s3 v0.0.0-00010101000000-000000000000
-	github.com/hashicorp/terraform/internal/legacy v0.0.0-00010101000000-000000000000
+	github.com/hashicorp/terraform/internal/backend/remote-state/azure v0.0.0-20260924130052-55d258b232a6
+	github.com/hashicorp/terraform/internal/backend/remote-state/consul v0.0.0-20260924130052-55d258b232a6
+	github.com/hashicorp/terraform/internal/backend/remote-state/cos v0.0.0-20260924130052-55d258b232a6
+	github.com/hashicorp/terraform/internal/backend/remote-state/gcs v0.0.0-20260924130052-55d258b232a6
+	github.com/hashicorp/terraform/internal/backend/remote-state/kubernetes v0.0.0-20260924130052-55d258b232a6
+	github.com/hashicorp/terraform/internal/backend/remote-state/oci v0.0.0-20260924130052-55d258b232a6
+	github.com/hashicorp/terraform/internal/backend/remote-state/oss v0.0.0-20260924130052-55d258b232a6
+	github.com/hashicorp/terraform/internal/backend/remote-state/pg v0.0.0-20260924130052-55d258b232a6
+	github.com/hashicorp/terraform/internal/backend/remote-state/s3 v0.0.0-20260924130052-55d258b232a6
+	github.com/hashicorp/terraform/internal/legacy v0.0.0-20260924130052-55d258b232a6
 	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0
-	github.com/masterzen/winrm v0.0.0-20200615185753-c42b5136ff88
-	github.com/mattn/go-isatty v0.0.20
-	github.com/mattn/go-shellwords v1.0.12
+	github.com/masterzen/winrm v0.0.0-20260407182533-5570be7f80cf
+	github.com/mattn/go-isatty v0.0.24
+	github.com/mattn/go-shellwords v1.0.15
 	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db
 	github.com/mitchellh/go-homedir v1.1.0
 	github.com/mitchellh/go-linereader v0.0.0-20190213213312-1b945b3263eb
@@ -60,132 +60,134 @@
 	github.com/packer-community/winrmcp v0.0.0-20221126162354-6e900dd2c68f
 	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c
 	github.com/posener/complete v1.2.3
-	github.com/spf13/afero v1.10.0
+	github.com/spf13/afero v1.15.0
 	github.com/xanzy/ssh-agent v0.3.3
 	github.com/xlab/treeprint v0.0.0-20161029104018-1d6e34225557
-	github.com/zclconf/go-cty v1.16.3
+	github.com/zclconf/go-cty v1.16.4
 	github.com/zclconf/go-cty-debug v0.0.0-20240509010212-0d6042c53940
 	github.com/zclconf/go-cty-yaml v1.1.0
 	go.opentelemetry.io/contrib/exporters/autoexport v0.68.0
-	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.63.0
-	go.opentelemetry.io/otel v1.43.0
-	go.opentelemetry.io/otel/sdk v1.43.0
-	go.opentelemetry.io/otel/trace v1.43.0
+	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.68.0
+	go.opentelemetry.io/otel v1.44.0
+	go.opentelemetry.io/otel/sdk v1.44.0
+	go.opentelemetry.io/otel/trace v1.44.0
 	go.uber.org/mock v0.6.0
-	golang.org/x/crypto v0.49.0
-	golang.org/x/mod v0.33.0
-	golang.org/x/net v0.52.0
+	golang.org/x/crypto v0.57.0
+	golang.org/x/mod v0.41.0
+	golang.org/x/net v0.59.0
 	golang.org/x/oauth2 v0.36.0
-	golang.org/x/sys v0.42.0
-	golang.org/x/term v0.41.0
-	golang.org/x/text v0.35.0
-	golang.org/x/tools v0.42.0
+	golang.org/x/sys v0.48.0
+	golang.org/x/term v0.46.0
+	golang.org/x/text v0.42.0
+	golang.org/x/tools v0.50.0
 	golang.org/x/tools/cmd/cover v0.1.0-deprecated
-	google.golang.org/grpc v1.80.0
-	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0
-	google.golang.org/protobuf v1.36.11
+	google.golang.org/grpc v1.83.2
+	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.3.0
+	google.golang.org/protobuf v1.36.12
 	honnef.co/go/tools v0.6.0
 )
 
 require (
-	cel.dev/expr v0.25.1 // indirect
+	cel.dev/expr v0.25.3 // indirect
 	cloud.google.com/go v0.123.0 // indirect
-	cloud.google.com/go/auth v0.18.2 // indirect
+	cloud.google.com/go/auth v0.23.3 // indirect
 	cloud.google.com/go/auth/oauth2adapt v0.2.8 // indirect
-	cloud.google.com/go/compute/metadata v0.9.0 // indirect
-	cloud.google.com/go/iam v1.5.3 // indirect
-	cloud.google.com/go/monitoring v1.24.3 // indirect
-	cloud.google.com/go/storage v1.61.3 // indirect
-	dario.cat/mergo v1.0.1 // indirect
+	cloud.google.com/go/compute/metadata v0.9.1 // indirect
+	cloud.google.com/go/iam v1.12.0 // indirect
+	cloud.google.com/go/monitoring v1.30.0 // indirect
+	cloud.google.com/go/storage v1.65.1 // indirect
+	dario.cat/mergo v1.0.2 // indirect
 	github.com/AlecAivazis/survey/v2 v2.3.7 // indirect
-	github.com/Azure/go-ntlmssp v0.0.0-20200615164410-66371956d46c // indirect
+	github.com/Azure/go-ntlmssp v0.0.1 // indirect
 	github.com/BurntSushi/toml v1.4.1-0.20240526193622-a339e1f7089c // indirect
-	github.com/ChrisTrenkamp/goxpath v0.0.0-20190607011252-c5096ec8773d // indirect
-	github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v1.31.0 // indirect
-	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v0.55.0 // indirect
-	github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping v0.55.0 // indirect
+	github.com/ChrisTrenkamp/goxpath v0.0.0-20210404020558-97928f7e12b6 // indirect
+	github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v1.33.0 // indirect
+	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v0.57.0 // indirect
+	github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping v0.57.0 // indirect
 	github.com/Masterminds/goutils v1.1.1 // indirect
-	github.com/Masterminds/semver/v3 v3.3.0 // indirect
+	github.com/Masterminds/semver/v3 v3.3.1 // indirect
 	github.com/Masterminds/sprig/v3 v3.3.0 // indirect
 	github.com/Microsoft/go-winio v0.6.2 // indirect
-	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1501 // indirect
+	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1843 // indirect
 	github.com/aliyun/aliyun-oss-go-sdk v0.0.0-20190103054945-8205d1f41e70 // indirect
-	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.2+incompatible // indirect
+	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.3+incompatible // indirect
 	github.com/antchfx/xmlquery v1.3.5 // indirect
 	github.com/antchfx/xpath v1.3.6 // indirect
 	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
 	github.com/armon/go-metrics v0.4.1 // indirect
 	github.com/armon/go-radix v1.0.0 // indirect
 	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
-	github.com/aws/aws-sdk-go-v2 v1.41.5 // indirect
-	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.7.8 // indirect
-	github.com/aws/aws-sdk-go-v2/config v1.32.12 // indirect
-	github.com/aws/aws-sdk-go-v2/credentials v1.19.12 // indirect
-	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.18.20 // indirect
-	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.17.22 // indirect
-	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.21 // indirect
-	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.21 // indirect
-	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.6 // indirect
-	github.com/aws/aws-sdk-go-v2/internal/v4a v1.4.22 // indirect
-	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.56.2 // indirect
-	github.com/aws/aws-sdk-go-v2/service/iam v1.53.6 // indirect
-	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.7 // indirect
-	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.9.13 // indirect
-	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.11.20 // indirect
-	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.21 // indirect
-	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.19.21 // indirect
-	github.com/aws/aws-sdk-go-v2/service/s3 v1.97.3 // indirect
-	github.com/aws/aws-sdk-go-v2/service/signin v1.0.8 // indirect
-	github.com/aws/aws-sdk-go-v2/service/sns v1.39.13 // indirect
-	github.com/aws/aws-sdk-go-v2/service/sqs v1.42.24 // indirect
-	github.com/aws/aws-sdk-go-v2/service/sso v1.30.13 // indirect
-	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.35.17 // indirect
-	github.com/aws/aws-sdk-go-v2/service/sts v1.41.9 // indirect
-	github.com/aws/smithy-go v1.24.2 // indirect
+	github.com/aws/aws-sdk-go-v2 v1.44.0 // indirect
+	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.7.20 // indirect
+	github.com/aws/aws-sdk-go-v2/config v1.32.40 // indirect
+	github.com/aws/aws-sdk-go-v2/credentials v1.19.39 // indirect
+	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.18.40 // indirect
+	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.17.85 // indirect
+	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.40 // indirect
+	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.40 // indirect
+	github.com/aws/aws-sdk-go-v2/internal/v4a v1.4.41 // indirect
+	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.62.3 // indirect
+	github.com/aws/aws-sdk-go-v2/service/iam v1.56.2 // indirect
+	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.19 // indirect
+	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.10.0 // indirect
+	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.12.17 // indirect
+	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.40 // indirect
+	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.19.41 // indirect
+	github.com/aws/aws-sdk-go-v2/service/s3 v1.108.0 // indirect
+	github.com/aws/aws-sdk-go-v2/service/signin v1.6.0 // indirect
+	github.com/aws/aws-sdk-go-v2/service/sns v1.39.21 // indirect
+	github.com/aws/aws-sdk-go-v2/service/sqs v1.46.8 // indirect
+	github.com/aws/aws-sdk-go-v2/service/sso v1.34.0 // indirect
+	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.39.0 // indirect
+	github.com/aws/aws-sdk-go-v2/service/sts v1.46.0 // indirect
+	github.com/aws/smithy-go v1.28.2 // indirect
 	github.com/beorn7/perks v1.0.1 // indirect
 	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d // indirect
 	github.com/bmatcuk/doublestar/v4 v4.6.1 // indirect
+	github.com/bodgit/ntlmssp v0.0.0-20240506230425-31973bb52d9b // indirect
+	github.com/bodgit/windows v1.0.1 // indirect
 	github.com/bradleyfalzon/ghinstallation/v2 v2.11.0 // indirect
 	github.com/cenkalti/backoff/v5 v5.0.3 // indirect
 	github.com/cespare/xxhash/v2 v2.3.0 // indirect
 	github.com/clbanning/mxj v1.8.4 // indirect
 	github.com/cli/go-gh/v2 v2.12.1 // indirect
 	github.com/cli/safeexec v1.0.1 // indirect
-	github.com/cloudflare/circl v1.6.3 // indirect
-	github.com/cncf/xds/go v0.0.0-20251210132809-ee656c7534f5 // indirect
+	github.com/cloudflare/circl v1.6.5 // indirect
+	github.com/cncf/xds/go v0.0.0-20260202195803-dba9d589def2 // indirect
 	github.com/creack/pty v1.1.17 // indirect
 	github.com/dylanmei/iso8601 v0.1.0 // indirect
-	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
-	github.com/envoyproxy/go-control-plane/envoy v1.36.0 // indirect
-	github.com/envoyproxy/protoc-gen-validate v1.3.0 // indirect
-	github.com/fatih/color v1.18.0 // indirect
+	github.com/emicklei/go-restful/v3 v3.11.3 // indirect
+	github.com/envoyproxy/go-control-plane/envoy v1.37.0 // indirect
+	github.com/envoyproxy/protoc-gen-validate v1.3.3 // indirect
+	github.com/fatih/color v1.19.0 // indirect
 	github.com/felixge/httpsnoop v1.0.4 // indirect
 	github.com/fsnotify/fsnotify v1.7.0 // indirect
-	github.com/fxamacker/cbor/v2 v2.7.0 // indirect
-	github.com/go-jose/go-jose/v4 v4.1.4 // indirect
-	github.com/go-logr/logr v1.4.3 // indirect
+	github.com/fxamacker/cbor/v2 v2.7.1 // indirect
+	github.com/go-jose/go-jose/v4 v4.1.5 // indirect
+	github.com/go-logr/logr v1.4.4 // indirect
 	github.com/go-logr/stdr v1.2.2 // indirect
 	github.com/go-openapi/errors v0.22.0 // indirect
-	github.com/go-openapi/jsonpointer v0.21.0 // indirect
-	github.com/go-openapi/jsonreference v0.20.2 // indirect
+	github.com/go-openapi/jsonpointer v0.21.2 // indirect
+	github.com/go-openapi/jsonreference v0.20.5 // indirect
 	github.com/go-openapi/strfmt v0.23.0 // indirect
-	github.com/go-openapi/swag v0.23.0 // indirect
+	github.com/go-openapi/swag v0.23.1 // indirect
+	github.com/go-viper/mapstructure/v2 v2.5.0 // indirect
 	github.com/gofrs/flock v0.10.0 // indirect
-	github.com/gofrs/uuid v4.0.0+incompatible // indirect
+	github.com/gofrs/uuid v4.4.0+incompatible // indirect
 	github.com/gogo/protobuf v1.3.2 // indirect
 	github.com/golang-jwt/jwt/v4 v4.5.2 // indirect
-	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
+	github.com/golang/groupcache v0.0.0-20241129210726-2c02b8208cf8 // indirect
 	github.com/golang/protobuf v1.5.4 // indirect
 	github.com/google/gnostic-models v0.6.9 // indirect
 	github.com/google/go-github/v45 v45.2.0 // indirect
 	github.com/google/go-github/v62 v62.0.0 // indirect
 	github.com/google/go-querystring v1.1.0 // indirect
-	github.com/google/s2a-go v0.1.9 // indirect
-	github.com/googleapis/enterprise-certificate-proxy v0.3.14 // indirect
-	github.com/googleapis/gax-go/v2 v2.17.0 // indirect
+	github.com/google/s2a-go v0.1.10 // indirect
+	github.com/googleapis/enterprise-certificate-proxy v0.3.22 // indirect
+	github.com/googleapis/gax-go/v2 v2.24.1 // indirect
 	github.com/grpc-ecosystem/grpc-gateway/v2 v2.28.0 // indirect
-	github.com/hashicorp/aws-sdk-go-base/v2 v2.0.0-beta.72 // indirect
-	github.com/hashicorp/consul/api v1.32.1 // indirect
+	github.com/hashicorp/aws-sdk-go-base/v2 v2.0.0-beta.74 // indirect
+	github.com/hashicorp/consul/api v1.32.4 // indirect
 	github.com/hashicorp/copywrite v0.20.0 // indirect
 	github.com/hashicorp/errwrap v1.1.0 // indirect
 	github.com/hashicorp/go-azure-helpers v0.72.0 // indirect
@@ -193,19 +195,25 @@
 	github.com/hashicorp/go-azure-sdk/sdk v0.20250131.1134653 // indirect
 	github.com/hashicorp/go-cty v1.4.1 // indirect
 	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
-	github.com/hashicorp/go-metrics v0.5.4 // indirect
+	github.com/hashicorp/go-metrics v0.6.1 // indirect
 	github.com/hashicorp/go-multierror v1.1.1 // indirect
 	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
 	github.com/hashicorp/golang-lru v1.0.2 // indirect
 	github.com/hashicorp/logutils v1.0.0 // indirect
-	github.com/hashicorp/serf v0.10.2 // indirect
+	github.com/hashicorp/serf v0.10.4 // indirect
 	github.com/hashicorp/terraform-plugin-go v0.26.0 // indirect
-	github.com/hashicorp/terraform-plugin-log v0.10.0 // indirect
+	github.com/hashicorp/terraform-plugin-log v0.10.1-0.20260721164921-2f73dc013b2c // indirect
 	github.com/hashicorp/terraform-plugin-sdk/v2 v2.36.1 // indirect
 	github.com/hashicorp/yamux v0.1.2 // indirect
 	github.com/huandu/xstrings v1.5.0 // indirect
 	github.com/inconshreveable/mousetrap v1.1.0 // indirect
 	github.com/jackofallops/giovanni v0.28.0 // indirect
+	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
+	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
+	github.com/jcmturner/gofork v1.7.6 // indirect
+	github.com/jcmturner/goidentity/v6 v6.0.1 // indirect
+	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
+	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
 	github.com/jedib0t/go-pretty v4.3.0+incompatible // indirect
 	github.com/jedib0t/go-pretty/v6 v6.5.9 // indirect
 	github.com/jmespath/go-jmespath v0.4.0 // indirect
@@ -213,12 +221,12 @@
 	github.com/josharian/intern v1.0.0 // indirect
 	github.com/json-iterator/go v1.1.12 // indirect
 	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
-	github.com/klauspost/compress v1.18.5 // indirect
+	github.com/klauspost/compress v1.19.2 // indirect
 	github.com/knadh/koanf v1.5.0 // indirect
-	github.com/lib/pq v1.10.3 // indirect
-	github.com/mailru/easyjson v0.7.7 // indirect
+	github.com/lib/pq v1.10.9 // indirect
+	github.com/mailru/easyjson v0.9.2 // indirect
 	github.com/masterzen/simplexml v0.0.0-20190410153822-31eea3082786 // indirect
-	github.com/mattn/go-colorable v0.1.14 // indirect
+	github.com/mattn/go-colorable v0.1.15 // indirect
 	github.com/mattn/go-runewidth v0.0.16 // indirect
 	github.com/mergestat/timediff v0.0.3 // indirect
 	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
@@ -229,33 +237,34 @@
 	github.com/mitchellh/reflectwalk v1.0.2 // indirect
 	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
 	github.com/modern-go/reflect2 v1.0.2 // indirect
-	github.com/mozillazg/go-httpheader v0.3.0 // indirect
+	github.com/mozillazg/go-httpheader v0.3.1 // indirect
 	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
 	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
 	github.com/oklog/run v1.1.0 // indirect
 	github.com/oklog/ulid v1.3.1 // indirect
-	github.com/oracle/oci-go-sdk/v65 v65.89.1 // indirect
+	github.com/oracle/oci-go-sdk/v65 v65.89.3 // indirect
 	github.com/pkg/errors v0.9.1 // indirect
 	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
-	github.com/prometheus/client_golang v1.23.2 // indirect
-	github.com/prometheus/client_model v0.6.2 // indirect
-	github.com/prometheus/common v0.67.5 // indirect
+	github.com/prometheus/client_golang v1.24.1 // indirect
+	github.com/prometheus/client_model v0.6.3 // indirect
+	github.com/prometheus/common v0.70.1 // indirect
 	github.com/prometheus/otlptranslator v1.0.0 // indirect
-	github.com/prometheus/procfs v0.20.1 // indirect
+	github.com/prometheus/procfs v0.21.1 // indirect
 	github.com/rivo/uniseg v0.4.7 // indirect
 	github.com/samber/lo v1.47.0 // indirect
 	github.com/shopspring/decimal v1.4.0 // indirect
 	github.com/sony/gobreaker v0.5.0 // indirect
-	github.com/spf13/cast v1.7.0 // indirect
+	github.com/spf13/cast v1.7.1 // indirect
 	github.com/spf13/cobra v1.8.1 // indirect
-	github.com/spf13/pflag v1.0.5 // indirect
-	github.com/spiffe/go-spiffe/v2 v2.6.0 // indirect
-	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.588 // indirect
-	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts v1.0.588 // indirect
-	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tag v1.0.233 // indirect
-	github.com/tencentyun/cos-go-sdk-v5 v0.7.42 // indirect
+	github.com/spf13/pflag v1.0.10 // indirect
+	github.com/spiffe/go-spiffe/v2 v2.7.0 // indirect
+	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.1217 // indirect
+	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts v1.0.1200 // indirect
+	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tag v1.0.1200 // indirect
+	github.com/tencentyun/cos-go-sdk-v5 v0.7.75 // indirect
 	github.com/thanhpk/randstr v1.0.6 // indirect
-	github.com/ulikunitz/xz v0.5.15 // indirect
+	github.com/tidwall/transform v0.0.0-20201103190739-32f242e2dbde // indirect
+	github.com/ulikunitz/xz v0.5.17 // indirect
 	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
 	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
 	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
@@ -264,9 +273,9 @@
 	go.mongodb.org/mongo-driver v1.16.1 // indirect
 	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
 	go.opentelemetry.io/contrib/bridges/prometheus v0.68.0 // indirect
-	go.opentelemetry.io/contrib/detectors/gcp v1.39.0 // indirect
-	go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws v0.67.0 // indirect
-	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.61.0 // indirect
+	go.opentelemetry.io/contrib/detectors/gcp v1.44.0 // indirect
+	go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws v0.69.0 // indirect
+	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.67.0 // indirect
 	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc v0.19.0 // indirect
 	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.19.0 // indirect
 	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.43.0 // indirect
@@ -276,36 +285,34 @@
 	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.43.0 // indirect
 	go.opentelemetry.io/otel/exporters/prometheus v0.65.0 // indirect
 	go.opentelemetry.io/otel/exporters/stdout/stdoutlog v0.19.0 // indirect
-	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.43.0 // indirect
+	go.opentelemetry.io/otel/exporters/stdout/stdoutmetric v1.44.0 // indirect
 	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.43.0 // indirect
 	go.opentelemetry.io/otel/log v0.19.0 // indirect
-	go.opentelemetry.io/otel/metric v1.43.0 // indirect
+	go.opentelemetry.io/otel/metric v1.44.0 // indirect
 	go.opentelemetry.io/otel/sdk/log v0.19.0 // indirect
-	go.opentelemetry.io/otel/sdk/metric v1.43.0 // indirect
+	go.opentelemetry.io/otel/sdk/metric v1.44.0 // indirect
 	go.opentelemetry.io/proto/otlp v1.10.0 // indirect
-	go.yaml.in/yaml/v2 v2.4.4 // indirect
-	golang.org/x/exp v0.0.0-20250506013437-ce4c2cf36ca6 // indirect
+	golang.org/x/exp v0.0.0-20260908205506-85c1c2202aba // indirect
 	golang.org/x/exp/typeparams v0.0.0-20231108232855-2478ac86f678 // indirect
-	golang.org/x/sync v0.20.0 // indirect
-	golang.org/x/telemetry v0.0.0-20260209163413-e7419c687ee4 // indirect
+	golang.org/x/sync v0.23.0 // indirect
+	golang.org/x/telemetry v0.0.0-20260908163034-4bcc4b2ee518 // indirect
 	golang.org/x/time v0.15.0 // indirect
-	golang.org/x/tools/go/expect v0.1.1-deprecated // indirect
-	google.golang.org/api v0.271.0 // indirect
+	google.golang.org/api v0.294.0 // indirect
 	google.golang.org/appengine v1.6.8 // indirect
-	google.golang.org/genproto v0.0.0-20260128011058-8636f8732409 // indirect
-	google.golang.org/genproto/googleapis/api v0.0.0-20260406210006-6f92a3bedf2d // indirect
-	google.golang.org/genproto/googleapis/rpc v0.0.0-20260406210006-6f92a3bedf2d // indirect
+	google.golang.org/genproto v0.0.0-20260921155816-b14227669459 // indirect
+	google.golang.org/genproto/googleapis/api v0.0.0-20260921155816-b14227669459 // indirect
+	google.golang.org/genproto/googleapis/rpc v0.0.0-20260921155816-b14227669459 // indirect
 	gopkg.in/evanphx/json-patch.v4 v4.12.0 // indirect
 	gopkg.in/inf.v0 v0.9.1 // indirect
-	gopkg.in/ini.v1 v1.66.2 // indirect
+	gopkg.in/ini.v1 v1.66.6 // indirect
 	gopkg.in/yaml.v3 v3.0.1 // indirect
-	k8s.io/api v0.33.0 // indirect
-	k8s.io/apimachinery v0.33.0 // indirect
-	k8s.io/client-go v0.33.0 // indirect
+	k8s.io/api v0.33.13 // indirect
+	k8s.io/apimachinery v0.33.13 // indirect
+	k8s.io/client-go v0.33.13 // indirect
 	k8s.io/klog/v2 v2.130.1 // indirect
-	k8s.io/kube-openapi v0.0.0-20250318190949-c8a335a9a2ff // indirect
-	k8s.io/utils v0.0.0-20241104100929-3ea5e8cea738 // indirect
-	sigs.k8s.io/json v0.0.0-20241010143419-9aa6b5e7a4b3 // indirect
+	k8s.io/kube-openapi v0.0.0-20260911184034-7970a1e230da // indirect
+	k8s.io/utils v0.0.0-20260707023825-cf1189d6abe3 // indirect
+	sigs.k8s.io/json v0.0.0-20260909141634-11ed52e25bc5 // indirect
 	sigs.k8s.io/randfill v1.0.0 // indirect
 	sigs.k8s.io/structured-merge-diff/v4 v4.6.0 // indirect
 	sigs.k8s.io/yaml v1.4.0 // indirect
@@ -364,3 +371,7 @@
 	golang.org/x/tools/cmd/stringer
 	honnef.co/go/tools/cmd/staticcheck
 )
+
+replace k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20250318190949-c8a335a9a2ff
+
+replace google.golang.org/genproto => google.golang.org/genproto v0.0.0-20260803160001-6ac0973c030d
