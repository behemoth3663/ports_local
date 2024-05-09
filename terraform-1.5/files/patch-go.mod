--- go.mod.orig	2023-09-07 18:19:50 UTC
+++ go.mod
@@ -1,32 +1,32 @@ require (
 module github.com/hashicorp/terraform
 
 require (
-	cloud.google.com/go/kms v1.6.0
-	cloud.google.com/go/storage v1.28.0
+	cloud.google.com/go/kms v1.15.9
+	cloud.google.com/go/storage v1.39.1
 	github.com/Azure/azure-sdk-for-go v59.2.0+incompatible
-	github.com/Azure/go-autorest/autorest v0.11.24
+	github.com/Azure/go-autorest/autorest v0.11.29
 	github.com/Netflix/go-expect v0.0.0-20220104043353-73e0943537d2
 	github.com/agext/levenshtein v1.2.3
-	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1501
+	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1843
 	github.com/aliyun/aliyun-oss-go-sdk v0.0.0-20190103054945-8205d1f41e70
-	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.2+incompatible
+	github.com/aliyun/aliyun-tablestore-go-sdk v4.1.3+incompatible
 	github.com/apparentlymart/go-cidr v1.1.0
 	github.com/apparentlymart/go-dump v0.0.0-20190214190832-042adf3cf4a0
 	github.com/apparentlymart/go-shquot v0.0.1
 	github.com/apparentlymart/go-userdirs v0.0.0-20200915174352-b0c018a67c13
-	github.com/apparentlymart/go-versions v1.0.1
+	github.com/apparentlymart/go-versions v1.0.2
 	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2
-	github.com/aws/aws-sdk-go v1.44.122
+	github.com/aws/aws-sdk-go v1.44.334
 	github.com/bgentry/speakeasy v0.1.0
 	github.com/bmatcuk/doublestar v1.1.5
 	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e
-	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f
+	github.com/coreos/pkg v0.0.0-20240122114842-bbd7aa9bf6fb
 	github.com/davecgh/go-spew v1.1.1
 	github.com/dylanmei/winrmtest v0.0.0-20210303004826-fbc9ae56efb6
 	github.com/go-test/deep v1.0.3
 	github.com/golang/mock v1.6.0
-	github.com/google/go-cmp v0.5.9
-	github.com/google/uuid v1.3.0
+	github.com/google/go-cmp v0.6.0
+	github.com/google/uuid v1.6.0
 	github.com/hashicorp/aws-sdk-go-base v0.7.1
 	github.com/hashicorp/consul/api v1.9.1
 	github.com/hashicorp/consul/sdk v0.8.0
@@ -34,11 +34,11 @@ require (
 	github.com/hashicorp/go-azure-helpers v0.43.0
 	github.com/hashicorp/go-checkpoint v0.5.0
 	github.com/hashicorp/go-cleanhttp v0.5.2
-	github.com/hashicorp/go-getter v1.7.0
-	github.com/hashicorp/go-hclog v0.15.0
+	github.com/hashicorp/go-getter v1.7.4
+	github.com/hashicorp/go-hclog v1.6.3
 	github.com/hashicorp/go-multierror v1.1.1
-	github.com/hashicorp/go-plugin v1.4.3
-	github.com/hashicorp/go-retryablehttp v0.7.2
+	github.com/hashicorp/go-plugin v1.4.10
+	github.com/hashicorp/go-retryablehttp v0.7.6
 	github.com/hashicorp/go-tfe v1.26.0
 	github.com/hashicorp/go-uuid v1.0.3
 	github.com/hashicorp/go-version v1.6.0
@@ -46,14 +46,14 @@ require (
 	github.com/hashicorp/hcl/v2 v2.16.2
 	github.com/hashicorp/jsonapi v0.0.0-20210826224640-ee7dae0fb22d
 	github.com/hashicorp/terraform-registry-address v0.0.0-20220623143253-7d51757b572c
-	github.com/hashicorp/terraform-svchost v0.1.0
+	github.com/hashicorp/terraform-svchost v0.1.1
 	github.com/jmespath/go-jmespath v0.4.0
 	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0
-	github.com/lib/pq v1.10.3
+	github.com/lib/pq v1.10.9
 	github.com/manicminer/hamilton v0.44.0
-	github.com/masterzen/winrm v0.0.0-20200615185753-c42b5136ff88
-	github.com/mattn/go-isatty v0.0.16
-	github.com/mattn/go-shellwords v1.0.4
+	github.com/masterzen/winrm v0.0.0-20231227165926-e811dad5ac77
+	github.com/mattn/go-isatty v0.0.20
+	github.com/mattn/go-shellwords v1.0.12
 	github.com/mitchellh/cli v1.1.5
 	github.com/mitchellh/colorstring v0.0.0-20190213212951-d06e56a500db
 	github.com/mitchellh/copystructure v1.2.0
@@ -61,65 +61,66 @@ require (
 	github.com/mitchellh/go-linereader v0.0.0-20190213213312-1b945b3263eb
 	github.com/mitchellh/go-wordwrap v1.0.1
 	github.com/mitchellh/gox v1.0.1
-	github.com/mitchellh/mapstructure v1.1.2
+	github.com/mitchellh/mapstructure v1.4.3
 	github.com/mitchellh/reflectwalk v1.0.2
 	github.com/nishanths/exhaustive v0.7.11
-	github.com/packer-community/winrmcp v0.0.0-20180921211025-c76d91c1e7db
-	github.com/pkg/browser v0.0.0-20201207095918-0426ae3fba23
+	github.com/packer-community/winrmcp v0.0.0-20221126162354-6e900dd2c68f
+	github.com/pkg/browser v0.0.0-20240102092130-5ac0b6a4141c
 	github.com/pkg/errors v0.9.1
 	github.com/posener/complete v1.2.3
 	github.com/spf13/afero v1.2.2
-	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.588
-	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts v1.0.588
-	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tag v1.0.233
-	github.com/tencentyun/cos-go-sdk-v5 v0.7.29
+	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common v1.0.916
+	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts v1.0.916
+	github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tag v1.0.916
+	github.com/tencentyun/cos-go-sdk-v5 v0.7.47
 	github.com/tombuildsstuff/giovanni v0.15.1
-	github.com/xanzy/ssh-agent v0.3.1
+	github.com/xanzy/ssh-agent v0.3.3
 	github.com/xlab/treeprint v0.0.0-20161029104018-1d6e34225557
-	github.com/zclconf/go-cty v1.12.2
+	github.com/zclconf/go-cty v1.13.3
 	github.com/zclconf/go-cty-debug v0.0.0-20191215020915-b22d67c1ba0b
 	github.com/zclconf/go-cty-yaml v1.0.3
-	golang.org/x/crypto v0.1.0
-	golang.org/x/mod v0.8.0
-	golang.org/x/net v0.7.0
-	golang.org/x/oauth2 v0.4.0
-	golang.org/x/sys v0.5.0
-	golang.org/x/term v0.5.0
-	golang.org/x/text v0.8.0
-	golang.org/x/tools v0.6.0
+	golang.org/x/crypto v0.22.0
+	golang.org/x/mod v0.17.0
+	golang.org/x/net v0.24.0
+	golang.org/x/oauth2 v0.19.0
+	golang.org/x/sys v0.20.0
+	golang.org/x/term v0.19.0
+	golang.org/x/text v0.15.0
+	golang.org/x/tools v0.20.0
 	golang.org/x/tools/cmd/cover v0.1.0-deprecated
-	google.golang.org/api v0.103.0
-	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f
-	google.golang.org/grpc v1.53.0
+	google.golang.org/api v0.177.0
+	google.golang.org/genproto v0.0.0-20240509183442-62759503f434
+	google.golang.org/grpc v1.63.2
 	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.1.0
-	google.golang.org/protobuf v1.28.1
-	honnef.co/go/tools v0.4.2
-	k8s.io/api v0.23.4
-	k8s.io/apimachinery v0.23.4
-	k8s.io/client-go v0.23.4
-	k8s.io/utils v0.0.0-20211116205334-6203023598ed
+	google.golang.org/protobuf v1.34.1
+	honnef.co/go/tools v0.4.7
+	k8s.io/api v0.23.17
+	k8s.io/apimachinery v0.23.17
+	k8s.io/client-go v0.23.17
+	k8s.io/utils v0.0.0-20240502163921-fe8a2dddb1d0
 )
 
 require (
-	cloud.google.com/go v0.107.0 // indirect
-	cloud.google.com/go/compute v1.15.1 // indirect
-	cloud.google.com/go/compute/metadata v0.2.3 // indirect
-	cloud.google.com/go/iam v0.8.0 // indirect
+	cloud.google.com/go v0.112.2 // indirect
+	cloud.google.com/go/auth v0.3.0 // indirect
+	cloud.google.com/go/auth/oauth2adapt v0.2.2 // indirect
+	cloud.google.com/go/compute/metadata v0.3.0 // indirect
+	cloud.google.com/go/iam v1.1.8 // indirect
 	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
-	github.com/Azure/go-autorest/autorest/adal v0.9.18 // indirect
-	github.com/Azure/go-autorest/autorest/azure/cli v0.4.4 // indirect
+	github.com/Azure/go-autorest/autorest/adal v0.9.23 // indirect
+	github.com/Azure/go-autorest/autorest/azure/cli v0.4.6 // indirect
 	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
 	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
 	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
 	github.com/Azure/go-autorest/logger v0.2.1 // indirect
 	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
-	github.com/Azure/go-ntlmssp v0.0.0-20200615164410-66371956d46c // indirect
+	github.com/Azure/go-ntlmssp v0.0.0-20221128193559-754e69321358 // indirect
 	github.com/BurntSushi/toml v1.2.1 // indirect
-	github.com/ChrisTrenkamp/goxpath v0.0.0-20190607011252-c5096ec8773d // indirect
+	github.com/ChrisTrenkamp/goxpath v0.0.0-20210404020558-97928f7e12b6 // indirect
 	github.com/Masterminds/goutils v1.1.1 // indirect
-	github.com/Masterminds/semver/v3 v3.1.1 // indirect
-	github.com/Masterminds/sprig/v3 v3.2.2 // indirect
-	github.com/Microsoft/go-winio v0.5.0 // indirect
+	github.com/Masterminds/semver/v3 v3.2.1 // indirect
+	github.com/Masterminds/sprig/v3 v3.2.3 // indirect
+	github.com/Microsoft/go-winio v0.5.2 // indirect
 	github.com/antchfx/xmlquery v1.3.5 // indirect
 	github.com/antchfx/xpath v1.1.10 // indirect
 	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
@@ -127,34 +128,52 @@ require (
 	github.com/armon/go-radix v1.0.0 // indirect
 	github.com/baiyubin/aliyun-sts-go-sdk v0.0.0-20180326062324-cfa1a18b161f // indirect
 	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d // indirect
-	github.com/coreos/go-systemd v0.0.0-20181012123002-c6f51f82210d // indirect
+	github.com/bodgit/ntlmssp v0.0.0-20240506230425-31973bb52d9b // indirect
+	github.com/bodgit/windows v1.0.1 // indirect
+	github.com/clbanning/mxj v1.8.4 // indirect
+	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
 	github.com/creack/pty v1.1.18 // indirect
 	github.com/dimchansky/utfbom v1.1.1 // indirect
 	github.com/dylanmei/iso8601 v0.1.0 // indirect
-	github.com/fatih/color v1.13.0 // indirect
-	github.com/go-logr/logr v1.2.0 // indirect
-	github.com/gofrs/uuid v4.0.0+incompatible // indirect
+	github.com/fatih/color v1.16.0 // indirect
+	github.com/felixge/httpsnoop v1.0.4 // indirect
+	github.com/go-logr/logr v1.4.1 // indirect
+	github.com/go-logr/stdr v1.2.2 // indirect
+	github.com/go-openapi/jsonpointer v0.20.3 // indirect
+	github.com/go-openapi/jsonreference v0.20.5 // indirect
+	github.com/go-openapi/swag v0.22.10 // indirect
+	github.com/gofrs/uuid v4.4.0+incompatible // indirect
 	github.com/gogo/protobuf v1.3.2 // indirect
-	github.com/golang-jwt/jwt/v4 v4.2.0 // indirect
+	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
 	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
-	github.com/golang/protobuf v1.5.2 // indirect
+	github.com/golang/protobuf v1.5.4 // indirect
+	github.com/google/gnostic-models v0.6.8 // indirect
 	github.com/google/go-querystring v1.1.0 // indirect
 	github.com/google/gofuzz v1.1.0 // indirect
-	github.com/googleapis/enterprise-certificate-proxy v0.2.0 // indirect
-	github.com/googleapis/gax-go/v2 v2.7.0 // indirect
+	github.com/google/s2a-go v0.1.7 // indirect
+	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
+	github.com/googleapis/gax-go/v2 v2.12.4 // indirect
 	github.com/googleapis/gnostic v0.5.5 // indirect
 	github.com/hashicorp/go-immutable-radix v1.0.0 // indirect
 	github.com/hashicorp/go-msgpack v0.5.4 // indirect
 	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
 	github.com/hashicorp/go-safetemp v1.0.0 // indirect
 	github.com/hashicorp/go-slug v0.11.1 // indirect
-	github.com/hashicorp/golang-lru v0.5.1 // indirect
-	github.com/hashicorp/serf v0.9.5 // indirect
+	github.com/hashicorp/golang-lru v0.5.4 // indirect
+	github.com/hashicorp/serf v0.9.8 // indirect
 	github.com/hashicorp/yamux v0.0.0-20181012175058-2f1d1f20f75d // indirect
 	github.com/huandu/xstrings v1.3.3 // indirect
-	github.com/imdario/mergo v0.3.13 // indirect
+	github.com/imdario/mergo v0.3.16 // indirect
+	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
+	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
+	github.com/jcmturner/gofork v1.7.6 // indirect
+	github.com/jcmturner/goidentity/v6 v6.0.1 // indirect
+	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
+	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
+	github.com/josharian/intern v1.0.0 // indirect
 	github.com/json-iterator/go v1.1.12 // indirect
-	github.com/klauspost/compress v1.15.11 // indirect
+	github.com/klauspost/compress v1.15.15 // indirect
+	github.com/mailru/easyjson v0.7.7 // indirect
 	github.com/manicminer/hamilton-autorest v0.2.0 // indirect
 	github.com/masterzen/simplexml v0.0.0-20190410153822-31eea3082786 // indirect
 	github.com/mattn/go-colorable v0.1.13 // indirect
@@ -162,32 +181,40 @@ require (
 	github.com/mitchellh/iochan v1.0.0 // indirect
 	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
 	github.com/modern-go/reflect2 v1.0.2 // indirect
-	github.com/mozillazg/go-httpheader v0.3.0 // indirect
+	github.com/mozillazg/go-httpheader v0.3.1 // indirect
 	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
 	github.com/oklog/run v1.0.0 // indirect
 	github.com/satori/go.uuid v1.2.0 // indirect
 	github.com/sergi/go-diff v1.2.0 // indirect
 	github.com/shopspring/decimal v1.3.1 // indirect
-	github.com/spf13/cast v1.5.0 // indirect
+	github.com/spf13/cast v1.5.1 // indirect
 	github.com/spf13/pflag v1.0.5 // indirect
-	github.com/ulikunitz/xz v0.5.10 // indirect
-	github.com/vmihailenco/msgpack/v4 v4.3.12 // indirect
-	github.com/vmihailenco/tagparser v0.1.1 // indirect
+	github.com/tidwall/transform v0.0.0-20201103190739-32f242e2dbde // indirect
+	github.com/ulikunitz/xz v0.5.12 // indirect
+	github.com/vmihailenco/msgpack/v5 v5.3.6 // indirect
+	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
 	go.opencensus.io v0.24.0 // indirect
-	golang.org/x/exp/typeparams v0.0.0-20221208152030-732eee02a75a // indirect
-	golang.org/x/time v0.3.0 // indirect
-	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
-	google.golang.org/appengine v1.6.7 // indirect
-	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
+	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.49.0 // indirect
+	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.49.0 // indirect
+	go.opentelemetry.io/otel v1.24.0 // indirect
+	go.opentelemetry.io/otel/metric v1.24.0 // indirect
+	go.opentelemetry.io/otel/trace v1.24.0 // indirect
+	golang.org/x/exp/typeparams v0.0.0-20240506185415-9bf2ced13842 // indirect
+	golang.org/x/sync v0.7.0 // indirect
+	golang.org/x/time v0.5.0 // indirect
+	google.golang.org/genproto/googleapis/api v0.0.0-20240509183442-62759503f434 // indirect
+	google.golang.org/genproto/googleapis/rpc v0.0.0-20240509183442-62759503f434 // indirect
 	gopkg.in/inf.v0 v0.9.1 // indirect
-	gopkg.in/ini.v1 v1.66.2 // indirect
+	gopkg.in/ini.v1 v1.66.6 // indirect
 	gopkg.in/yaml.v2 v2.4.0 // indirect
 	gopkg.in/yaml.v3 v3.0.1 // indirect
-	k8s.io/klog/v2 v2.30.0 // indirect
-	k8s.io/kube-openapi v0.0.0-20211115234752-e816edb12b65 // indirect
-	sigs.k8s.io/json v0.0.0-20211020170558-c049b76a60c6 // indirect
-	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
+	k8s.io/klog/v2 v2.120.1 // indirect
+	k8s.io/kube-openapi v0.0.0-20240430033511-f0e62f92d13f // indirect
+	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
+	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
 	sigs.k8s.io/yaml v1.2.0 // indirect
 )
 
-go 1.18
+go 1.21
+
+toolchain go1.22.3
