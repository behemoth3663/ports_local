--- go.mod.orig	2024-10-16 12:28:59 UTC
+++ go.mod
@@ -1,6 +1,6 @@ module github.com/hashicorp/terraform
 module github.com/hashicorp/terraform
 
-go 1.22.7
+go 1.24.0
 
 require (
 	github.com/Netflix/go-expect v0.0.0-20220104043353-73e0943537d2
@@ -9,44 +9,44 @@ require (
 	github.com/apparentlymart/go-cidr v1.1.0
 	github.com/apparentlymart/go-shquot v0.0.1
 	github.com/apparentlymart/go-userdirs v0.0.0-20200915174352-b0c018a67c13
-	github.com/apparentlymart/go-versions v1.0.2
+	github.com/apparentlymart/go-versions v1.0.3
 	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2
 	github.com/bgentry/speakeasy v0.1.0
 	github.com/bmatcuk/doublestar v1.1.5
 	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e
-	github.com/davecgh/go-spew v1.1.1
+	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc
 	github.com/dylanmei/winrmtest v0.0.0-20210303004826-fbc9ae56efb6
 	github.com/go-test/deep v1.0.3
-	github.com/google/go-cmp v0.6.0
-	github.com/google/uuid v1.3.1
-	github.com/hashicorp/cli v1.1.6
-	github.com/hashicorp/copywrite v0.16.3
+	github.com/google/go-cmp v0.7.0
+	github.com/google/uuid v1.6.0
+	github.com/hashicorp/cli v1.1.7
+	github.com/hashicorp/copywrite v0.16.6
 	github.com/hashicorp/go-checkpoint v0.5.0
 	github.com/hashicorp/go-cleanhttp v0.5.2
-	github.com/hashicorp/go-getter v1.7.5
+	github.com/hashicorp/go-getter v1.7.10
 	github.com/hashicorp/go-hclog v1.6.3
-	github.com/hashicorp/go-plugin v1.6.0
-	github.com/hashicorp/go-retryablehttp v0.7.7
-	github.com/hashicorp/go-slug v0.15.0
+	github.com/hashicorp/go-plugin v1.6.3
+	github.com/hashicorp/go-retryablehttp v0.7.8
+	github.com/hashicorp/go-slug v0.15.2
 	github.com/hashicorp/go-tfe v1.58.0
 	github.com/hashicorp/go-uuid v1.0.3
 	github.com/hashicorp/go-version v1.6.0
 	github.com/hashicorp/hcl v1.0.0
-	github.com/hashicorp/hcl/v2 v2.20.0
-	github.com/hashicorp/jsonapi v1.3.1
-	github.com/hashicorp/terraform-registry-address v0.2.3
+	github.com/hashicorp/hcl/v2 v2.20.1
+	github.com/hashicorp/jsonapi v1.3.2
+	github.com/hashicorp/terraform-registry-address v0.2.5
 	github.com/hashicorp/terraform-svchost v0.1.1
-	github.com/hashicorp/terraform/internal/backend/remote-state/azure v0.0.0-00010101000000-000000000000
-	github.com/hashicorp/terraform/internal/backend/remote-state/consul v0.0.0-00010101000000-000000000000
-	github.com/hashicorp/terraform/internal/backend/remote-state/cos v0.0.0-00010101000000-000000000000
-	github.com/hashicorp/terraform/internal/backend/remote-state/gcs v0.0.0-00010101000000-000000000000
-	github.com/hashicorp/terraform/internal/backend/remote-state/kubernetes v0.0.0-00010101000000-000000000000
-	github.com/hashicorp/terraform/internal/backend/remote-state/oss v0.0.0-00010101000000-000000000000
-	github.com/hashicorp/terraform/internal/backend/remote-state/pg v0.0.0-00010101000000-000000000000
-	github.com/hashicorp/terraform/internal/backend/remote-state/s3 v0.0.0-00010101000000-000000000000
-	github.com/hashicorp/terraform/internal/legacy v0.0.0-00010101000000-000000000000
+	github.com/hashicorp/terraform/internal/backend/remote-state/azure v0.0.0-20251030193032-f8ae45cfc8df
+	github.com/hashicorp/terraform/internal/backend/remote-state/consul v0.0.0-20251030193032-f8ae45cfc8df
+	github.com/hashicorp/terraform/internal/backend/remote-state/cos v0.0.0-20251030193032-f8ae45cfc8df
+	github.com/hashicorp/terraform/internal/backend/remote-state/gcs v0.0.0-20251030193032-f8ae45cfc8df
+	github.com/hashicorp/terraform/internal/backend/remote-state/kubernetes v0.0.0-20251030193032-f8ae45cfc8df
+	github.com/hashicorp/terraform/internal/backend/remote-state/oss v0.0.0-20251030193032-f8ae45cfc8df
+	github.com/hashicorp/terraform/internal/backend/remote-state/pg v0.0.0-20251030193032-f8ae45cfc8df
+	github.com/hashicorp/terraform/internal/backend/remote-state/s3 v0.0.0-20251030193032-f8ae45cfc8df
+	github.com/hashicorp/terraform/internal/legacy v0.0.0-20251030193032-f8ae45cfc8df
 	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0
-	github.com/masterzen/winrm v0.0.0-20200615185753-c42b5136ff88
+	github.com/masterzen/winrm v0.0.0-20250927112105-5f8e6c707321
 	github.com/mattn/go-isatty v0.0.20
 	github.com/mattn/go-shellwords v1.0.12
 	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db
@@ -58,220 +58,245 @@ require (
 	github.com/packer-community/winrmcp v0.0.0-20221126162354-6e900dd2c68f
 	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c
 	github.com/posener/complete v1.2.3
-	github.com/spf13/afero v1.9.3
+	github.com/spf13/afero v1.10.0
 	github.com/xanzy/ssh-agent v0.3.3
 	github.com/xlab/treeprint v0.0.0-20161029104018-1d6e34225557
 	github.com/zclconf/go-cty v1.14.4
 	github.com/zclconf/go-cty-debug v0.0.0-20240417160409-8c45e122ae1a
 	github.com/zclconf/go-cty-yaml v1.0.3
 	go.opentelemetry.io/contrib/exporters/autoexport v0.45.0
-	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.46.0
-	go.opentelemetry.io/otel v1.24.0
-	go.opentelemetry.io/otel/sdk v1.24.0
-	go.opentelemetry.io/otel/trace v1.24.0
+	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.61.0
+	go.opentelemetry.io/otel v1.37.0
+	go.opentelemetry.io/otel/sdk v1.37.0
+	go.opentelemetry.io/otel/trace v1.37.0
 	go.uber.org/mock v0.4.0
-	golang.org/x/crypto v0.21.0
-	golang.org/x/mod v0.16.0
-	golang.org/x/net v0.23.0
-	golang.org/x/oauth2 v0.18.0
-	golang.org/x/sys v0.20.0
-	golang.org/x/term v0.18.0
-	golang.org/x/text v0.14.0
-	golang.org/x/tools v0.19.0
+	golang.org/x/crypto v0.43.0
+	golang.org/x/mod v0.29.0
+	golang.org/x/net v0.46.0
+	golang.org/x/oauth2 v0.30.0
+	golang.org/x/sys v0.37.0
+	golang.org/x/term v0.36.0
+	golang.org/x/text v0.30.0
+	golang.org/x/tools v0.38.0
 	golang.org/x/tools/cmd/cover v0.1.0-deprecated
-	google.golang.org/grpc v1.59.0
+	google.golang.org/grpc v1.75.1
 	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
-	google.golang.org/protobuf v1.33.0
-	honnef.co/go/tools v0.5.0-0.dev.0.20230826160118-ad5ca31ff221
+	google.golang.org/protobuf v1.36.10
+	honnef.co/go/tools v0.5.1
 )
 
 require (
-	cloud.google.com/go v0.110.7 // indirect
-	cloud.google.com/go/compute v1.23.0 // indirect
-	cloud.google.com/go/compute/metadata v0.2.3 // indirect
-	cloud.google.com/go/iam v1.1.1 // indirect
-	cloud.google.com/go/storage v1.30.1 // indirect
-	github.com/AlecAivazis/survey/v2 v2.3.6 // indirect
+	cel.dev/expr v0.24.0 // indirect
+	cloud.google.com/go v0.121.6 // indirect
+	cloud.google.com/go/auth v0.16.5 // indirect
+	cloud.google.com/go/auth/oauth2adapt v0.2.8 // indirect
+	cloud.google.com/go/compute/metadata v0.8.4 // indirect
+	cloud.google.com/go/iam v1.5.3 // indirect
+	cloud.google.com/go/monitoring v1.24.3 // indirect
+	cloud.google.com/go/storage v1.56.3 // indirect
+	github.com/AlecAivazis/survey/v2 v2.3.7 // indirect
 	github.com/Azure/azure-sdk-for-go v59.2.0+incompatible // indirect
 	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
-	github.com/Azure/go-autorest/autorest v0.11.27 // indirect
-	github.com/Azure/go-autorest/autorest/adal v0.9.20 // indirect
-	github.com/Azure/go-autorest/autorest/azure/cli v0.4.4 // indirect
-	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
-	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
-	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
-	github.com/Azure/go-autorest/logger v0.2.1 // indirect
-	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
-	github.com/Azure/go-ntlmssp v0.0.0-20200615164410-66371956d46c // indirect
-	github.com/BurntSushi/toml v1.2.1 // indirect
-	github.com/ChrisTrenkamp/goxpath v0.0.0-20190607011252-c5096ec8773d // indirect
+	github.com/Azure/go-autorest/autorest v0.11.30 // indirect
+	github.com/Azure/go-autorest/autorest/adal v0.9.24 // indirect
+	github.com/Azure/go-autorest/autorest/azure/cli v0.4.7 // indirect
+	github.com/Azure/go-autorest/autorest/date v0.3.1 // indirect
+	github.com/Azure/go-autorest/autorest/to v0.4.1 // indirect
+	github.com/Azure/go-autorest/autorest/validation v0.3.2 // indirect
+	github.com/Azure/go-autorest/logger v0.2.2 // indirect
+	github.com/Azure/go-autorest/tracing v0.6.1 // indirect
+	github.com/Azure/go-ntlmssp v0.0.0-20221128193559-754e69321358 // indirect
+	github.com/BurntSushi/toml v1.4.1-0.20240526193622-a339e1f7089c // indirect
+	github.com/ChrisTrenkamp/goxpath v0.0.0-20210404020558-97928f7e12b6 // indirect
+	github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v1.29.0 // indirect
+	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v0.53.0 // indirect
+	github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping v0.53.0 // indirect
 	github.com/Masterminds/goutils v1.1.1 // indirect
-	github.com/Masterminds/semver/v3 v3.2.0 // indirect
+	github.com/Masterminds/semver/v3 v3.2.1 // indirect
 	github.com/Masterminds/sprig/v3 v3.2.3 // indirect
-	github.com/Microsoft/go-winio v0.5.2 // indirect
+	github.com/Microsoft/go-winio v0.6.2 // indirect
 	github.com/PuerkitoBio/purell v1.1.1 // indirect
 	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
-	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1501 // indirect
+	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1843 // indirect
 	github.com/aliyun/aliyun-oss-go-sdk v0.0.0-20190103054945-8205d1f41e70 // indirect
-	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.2+incompatible // indirect
+	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.3+incompatible // indirect
 	github.com/antchfx/xmlquery v1.3.5 // indirect
 	github.com/antchfx/xpath v1.1.10 // indirect
 	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
 	github.com/armon/go-metrics v0.0.0-20180917152333-f0300d1749da // indirect
 	github.com/armon/go-radix v1.0.0 // indirect
-	github.com/asaskevich/govalidator v0.0.0-20200907205600-7a23bdc65eef // indirect
-	github.com/aws/aws-sdk-go v1.44.122 // indirect
-	github.com/aws/aws-sdk-go-v2 v1.25.3 // indirect
-	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.1 // indirect
-	github.com/aws/aws-sdk-go-v2/config v1.27.7 // indirect
-	github.com/aws/aws-sdk-go-v2/credentials v1.17.7 // indirect
-	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.15.3 // indirect
-	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.16.9 // indirect
-	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.3 // indirect
-	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.3 // indirect
-	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.0 // indirect
-	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.3 // indirect
-	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.30.4 // indirect
-	github.com/aws/aws-sdk-go-v2/service/iam v1.28.5 // indirect
-	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.11.1 // indirect
-	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.3.5 // indirect
-	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.9.4 // indirect
-	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.11.5 // indirect
-	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.17.3 // indirect
-	github.com/aws/aws-sdk-go-v2/service/s3 v1.51.4 // indirect
-	github.com/aws/aws-sdk-go-v2/service/sqs v1.29.5 // indirect
-	github.com/aws/aws-sdk-go-v2/service/sso v1.20.2 // indirect
-	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.23.2 // indirect
-	github.com/aws/aws-sdk-go-v2/service/sts v1.28.4 // indirect
-	github.com/aws/smithy-go v1.20.1 // indirect
+	github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2 // indirect
+	github.com/aws/aws-sdk-go v1.44.334 // indirect
+	github.com/aws/aws-sdk-go-v2 v1.36.6 // indirect
+	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.6.11 // indirect
+	github.com/aws/aws-sdk-go-v2/config v1.27.43 // indirect
+	github.com/aws/aws-sdk-go-v2/credentials v1.17.71 // indirect
+	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.33 // indirect
+	github.com/aws/aws-sdk-go-v2/feature/s3/manager v1.16.25 // indirect
+	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.37 // indirect
+	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.37 // indirect
+	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.4 // indirect
+	github.com/aws/aws-sdk-go-v2/internal/v4a v1.3.37 // indirect
+	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.34.10 // indirect
+	github.com/aws/aws-sdk-go-v2/service/iam v1.35.3 // indirect
+	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.4 // indirect
+	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.3.20 // indirect
+	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.9.19 // indirect
+	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.18 // indirect
+	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.17.18 // indirect
+	github.com/aws/aws-sdk-go-v2/service/s3 v1.60.1 // indirect
+	github.com/aws/aws-sdk-go-v2/service/sqs v1.34.9 // indirect
+	github.com/aws/aws-sdk-go-v2/service/sso v1.25.6 // indirect
+	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.30.4 // indirect
+	github.com/aws/aws-sdk-go-v2/service/sts v1.34.1 // indirect
+	github.com/aws/smithy-go v1.22.5 // indirect
 	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d // indirect
-	github.com/bmatcuk/doublestar/v4 v4.6.0 // indirect
-	github.com/bradleyfalzon/ghinstallation/v2 v2.1.0 // indirect
+	github.com/bmatcuk/doublestar/v4 v4.6.1 // indirect
+	github.com/bodgit/ntlmssp v0.0.0-20240506230425-31973bb52d9b // indirect
+	github.com/bodgit/windows v1.0.1 // indirect
+	github.com/bradleyfalzon/ghinstallation/v2 v2.5.0 // indirect
 	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
+	github.com/cespare/xxhash/v2 v2.3.0 // indirect
 	github.com/clbanning/mxj v1.8.4 // indirect
-	github.com/cli/go-gh v1.0.0 // indirect
-	github.com/cli/safeexec v1.0.0 // indirect
-	github.com/cli/shurcooL-graphql v0.0.2 // indirect
-	github.com/cloudflare/circl v1.3.7 // indirect
+	github.com/cli/go-gh v1.2.1 // indirect
+	github.com/cli/safeexec v1.0.1 // indirect
+	github.com/cli/shurcooL-graphql v0.0.4 // indirect
+	github.com/clipperhouse/uax29/v2 v2.2.0 // indirect
+	github.com/cloudflare/circl v1.3.9 // indirect
+	github.com/cncf/xds/go v0.0.0-20251022180443-0feb69152e9f // indirect
 	github.com/creack/pty v1.1.17 // indirect
 	github.com/dimchansky/utfbom v1.1.1 // indirect
 	github.com/dylanmei/iso8601 v0.1.0 // indirect
 	github.com/emicklei/go-restful/v3 v3.8.0 // indirect
-	github.com/fatih/color v1.16.0 // indirect
+	github.com/envoyproxy/go-control-plane/envoy v1.32.4 // indirect
+	github.com/envoyproxy/protoc-gen-validate v1.2.1 // indirect
+	github.com/fatih/color v1.17.0 // indirect
+	github.com/felixge/httpsnoop v1.0.4 // indirect
 	github.com/fsnotify/fsnotify v1.5.4 // indirect
-	github.com/go-logr/logr v1.4.1 // indirect
+	github.com/go-jose/go-jose/v4 v4.1.3 // indirect
+	github.com/go-logr/logr v1.4.3 // indirect
 	github.com/go-logr/stdr v1.2.2 // indirect
-	github.com/go-openapi/errors v0.20.2 // indirect
-	github.com/go-openapi/jsonpointer v0.19.5 // indirect
-	github.com/go-openapi/jsonreference v0.19.5 // indirect
-	github.com/go-openapi/strfmt v0.21.3 // indirect
-	github.com/go-openapi/swag v0.19.14 // indirect
-	github.com/gofrs/uuid v4.0.0+incompatible // indirect
+	github.com/go-openapi/errors v0.21.1 // indirect
+	github.com/go-openapi/jsonpointer v0.19.6 // indirect
+	github.com/go-openapi/jsonreference v0.19.6 // indirect
+	github.com/go-openapi/strfmt v0.21.10 // indirect
+	github.com/go-openapi/swag v0.22.10 // indirect
+	github.com/gofrs/uuid v4.4.0+incompatible // indirect
 	github.com/gogo/protobuf v1.3.2 // indirect
-	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
-	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
-	github.com/golang/protobuf v1.5.3 // indirect
-	github.com/google/gnostic v0.5.7-v3refs // indirect
+	github.com/golang-jwt/jwt/v4 v4.5.2 // indirect
+	github.com/golang/groupcache v0.0.0-20241129210726-2c02b8208cf8 // indirect
+	github.com/golang/protobuf v1.5.4 // indirect
+	github.com/google/gnostic v0.5.7 // indirect
 	github.com/google/go-github/v45 v45.2.0 // indirect
+	github.com/google/go-github/v53 v53.0.0 // indirect
 	github.com/google/go-querystring v1.1.0 // indirect
 	github.com/google/gofuzz v1.1.0 // indirect
-	github.com/google/s2a-go v0.1.4 // indirect
-	github.com/googleapis/enterprise-certificate-proxy v0.2.3 // indirect
-	github.com/googleapis/gax-go/v2 v2.11.0 // indirect
-	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.0 // indirect
-	github.com/hashicorp/aws-sdk-go-base/v2 v2.0.0-beta.45 // indirect
-	github.com/hashicorp/consul/api v1.13.0 // indirect
+	github.com/google/s2a-go v0.1.9 // indirect
+	github.com/googleapis/enterprise-certificate-proxy v0.3.6 // indirect
+	github.com/googleapis/gax-go/v2 v2.15.0 // indirect
+	github.com/grpc-ecosystem/grpc-gateway/v2 v2.16.2 // indirect
+	github.com/hashicorp/aws-sdk-go-base/v2 v2.0.0-beta.67 // indirect
+	github.com/hashicorp/consul/api v1.13.1 // indirect
 	github.com/hashicorp/errwrap v1.1.0 // indirect
 	github.com/hashicorp/go-azure-helpers v0.43.0 // indirect
 	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
 	github.com/hashicorp/go-multierror v1.1.1 // indirect
 	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
 	github.com/hashicorp/go-safetemp v1.0.0 // indirect
-	github.com/hashicorp/golang-lru v0.5.1 // indirect
-	github.com/hashicorp/serf v0.9.6 // indirect
+	github.com/hashicorp/golang-lru v0.5.4 // indirect
+	github.com/hashicorp/serf v0.9.8 // indirect
 	github.com/hashicorp/terraform-plugin-log v0.9.0 // indirect
-	github.com/hashicorp/yamux v0.1.1 // indirect
+	github.com/hashicorp/yamux v0.1.2 // indirect
 	github.com/henvic/httpretty v0.0.6 // indirect
 	github.com/huandu/xstrings v1.3.3 // indirect
-	github.com/imdario/mergo v0.3.13 // indirect
+	github.com/imdario/mergo v0.3.16 // indirect
 	github.com/inconshreveable/mousetrap v1.0.1 // indirect
+	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
+	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
+	github.com/jcmturner/gofork v1.7.6 // indirect
+	github.com/jcmturner/goidentity/v6 v6.0.1 // indirect
+	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
+	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
 	github.com/jedib0t/go-pretty v4.3.0+incompatible // indirect
-	github.com/jedib0t/go-pretty/v6 v6.4.4 // indirect
+	github.com/jedib0t/go-pretty/v6 v6.4.9 // indirect
 	github.com/jmespath/go-jmespath v0.4.0 // indirect
 	github.com/joho/godotenv v1.3.0 // indirect
 	github.com/josharian/intern v1.0.0 // indirect
 	github.com/json-iterator/go v1.1.12 // indirect
 	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
-	github.com/klauspost/compress v1.15.11 // indirect
+	github.com/klauspost/compress v1.15.15 // indirect
 	github.com/knadh/koanf v1.5.0 // indirect
-	github.com/lib/pq v1.10.3 // indirect
+	github.com/lib/pq v1.10.9 // indirect
 	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
-	github.com/mailru/easyjson v0.7.6 // indirect
+	github.com/mailru/easyjson v0.7.7 // indirect
 	github.com/manicminer/hamilton v0.44.0 // indirect
 	github.com/manicminer/hamilton-autorest v0.2.0 // indirect
 	github.com/masterzen/simplexml v0.0.0-20190410153822-31eea3082786 // indirect
-	github.com/mattn/go-colorable v0.1.13 // indirect
-	github.com/mattn/go-runewidth v0.0.13 // indirect
-	github.com/mergestat/timediff v0.0.3 // indirect
+	github.com/mattn/go-colorable v0.1.14 // indirect
+	github.com/mattn/go-runewidth v0.0.19 // indirect
+	github.com/mergestat/timediff v0.0.4 // indirect
 	github.com/mgutz/ansi v0.0.0-20200706080929-d51e80ef957d // indirect
 	github.com/mitchellh/copystructure v1.2.0 // indirect
-	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
 	github.com/mitchellh/iochan v1.0.0 // indirect
 	github.com/mitchellh/mapstructure v1.5.0 // indirect
 	github.com/mitchellh/reflectwalk v1.0.2 // indirect
 	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
 	github.com/modern-go/reflect2 v1.0.2 // indirect
-	github.com/mozillazg/go-httpheader v0.3.0 // indirect
+	github.com/mozillazg/go-httpheader v0.3.1 // indirect
 	github.com/muesli/termenv v0.12.0 // indirect
 	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
 	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
 	github.com/oklog/run v1.0.0 // indirect
 	github.com/oklog/ulid v1.3.1 // indirect
-	github.com/rivo/uniseg v0.2.0 // indirect
+	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
 	github.com/samber/lo v1.37.0 // indirect
 	github.com/shopspring/decimal v1.3.1 // indirect
-	github.com/spf13/cast v1.5.0 // indirect
+	github.com/spf13/cast v1.5.1 // indirect
 	github.com/spf13/cobra v1.6.1 // indirect
-	github.com/spf13/pflag v1.0.5 // indirect
-	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.588 // indirect
-	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts v1.0.588 // indirect
-	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tag v1.0.233 // indirect
-	github.com/tencentyun/cos-go-sdk-v5 v0.7.42 // indirect
-	github.com/thanhpk/randstr v1.0.4 // indirect
-	github.com/thlib/go-timezone-local v0.0.0-20210907160436-ef149e42d28e // indirect
+	github.com/spf13/pflag v1.0.10 // indirect
+	github.com/spiffe/go-spiffe/v2 v2.5.0 // indirect
+	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.1217 // indirect
+	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts v1.0.1200 // indirect
+	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tag v1.0.1200 // indirect
+	github.com/tencentyun/cos-go-sdk-v5 v0.7.71 // indirect
+	github.com/thanhpk/randstr v1.0.6 // indirect
+	github.com/thlib/go-timezone-local v0.0.7 // indirect
+	github.com/tidwall/transform v0.0.0-20201103190739-32f242e2dbde // indirect
 	github.com/tombuildsstuff/giovanni v0.15.1 // indirect
-	github.com/ulikunitz/xz v0.5.10 // indirect
-	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
+	github.com/ulikunitz/xz v0.5.15 // indirect
+	github.com/vmihailenco/msgpack/v5 v5.3.6 // indirect
 	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
-	go.mongodb.org/mongo-driver v1.10.0 // indirect
-	go.opencensus.io v0.24.0 // indirect
-	go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws v0.46.1 // indirect
+	github.com/zeebo/errs v1.4.0 // indirect
+	go.mongodb.org/mongo-driver v1.13.4 // indirect
+	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
+	go.opentelemetry.io/contrib/detectors/gcp v1.36.0 // indirect
+	go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws v0.53.0 // indirect
+	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.61.0 // indirect
 	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.19.0 // indirect
 	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.19.0 // indirect
 	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.19.0 // indirect
-	go.opentelemetry.io/otel/metric v1.24.0 // indirect
+	go.opentelemetry.io/otel/metric v1.37.0 // indirect
+	go.opentelemetry.io/otel/sdk/metric v1.37.0 // indirect
 	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
-	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
-	golang.org/x/exp/typeparams v0.0.0-20221208152030-732eee02a75a // indirect
-	golang.org/x/sync v0.7.0 // indirect
-	golang.org/x/time v0.5.0 // indirect
-	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
-	google.golang.org/api v0.126.0 // indirect
-	google.golang.org/appengine v1.6.7 // indirect
-	google.golang.org/genproto v0.0.0-20230822172742-b8732ec3820d // indirect
-	google.golang.org/genproto/googleapis/api v0.0.0-20230822172742-b8732ec3820d // indirect
-	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
+	golang.org/x/exp v0.0.0-20251023183803-a4bb9ffd2546 // indirect
+	golang.org/x/exp/typeparams v0.0.0-20251023183803-a4bb9ffd2546 // indirect
+	golang.org/x/sync v0.17.0 // indirect
+	golang.org/x/time v0.12.0 // indirect
+	google.golang.org/api v0.247.0 // indirect
+	google.golang.org/genproto v0.0.0-20251029180050-ab9386a59fda // indirect
+	google.golang.org/genproto/googleapis/api v0.0.0-20251029180050-ab9386a59fda // indirect
+	google.golang.org/genproto/googleapis/rpc v0.0.0-20251029180050-ab9386a59fda // indirect
 	gopkg.in/inf.v0 v0.9.1 // indirect
-	gopkg.in/ini.v1 v1.66.2 // indirect
+	gopkg.in/ini.v1 v1.66.6 // indirect
 	gopkg.in/yaml.v2 v2.4.0 // indirect
 	gopkg.in/yaml.v3 v3.0.1 // indirect
-	k8s.io/api v0.25.5 // indirect
-	k8s.io/apimachinery v0.25.5 // indirect
-	k8s.io/client-go v0.25.5 // indirect
-	k8s.io/klog/v2 v2.70.1 // indirect
-	k8s.io/kube-openapi v0.0.0-20220803162953-67bda5d908f1 // indirect
-	k8s.io/utils v0.0.0-20220728103510-ee6ede2d64ed // indirect
-	sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2 // indirect
+	k8s.io/api v0.25.16 // indirect
+	k8s.io/apimachinery v0.25.16 // indirect
+	k8s.io/client-go v0.25.16 // indirect
+	k8s.io/klog/v2 v2.80.1 // indirect
+	k8s.io/kube-openapi v0.0.0-20250910181357-589584f1c912 // indirect
+	k8s.io/utils v0.0.0-20251002143259-bc988d571ff4 // indirect
+	sigs.k8s.io/json v0.0.0-20250730193827-2d320260d730 // indirect
 	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
 	sigs.k8s.io/yaml v1.2.0 // indirect
 )
@@ -316,3 +341,11 @@ replace github.com/hashicorp/terraform/internal/legacy
 replace github.com/hashicorp/terraform/internal/backend/remote-state/s3 => ./internal/backend/remote-state/s3
 
 replace github.com/hashicorp/terraform/internal/legacy => ./internal/legacy
+
+replace k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20220803162953-67bda5d908f1
+
+replace github.com/google/gnostic => github.com/google/gnostic v0.5.7-v3refs
+
+replace github.com/hashicorp/aws-sdk-go-base/v2 => github.com/hashicorp/aws-sdk-go-base/v2 v2.0.0-beta.55
+
+replace github.com/hashicorp/go-getter => github.com/hashicorp/go-getter v1.7.9
