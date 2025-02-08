PORTNAME=	terraform
PORTVERSION=	1.10.5
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils
PKGNAMESUFFIX?=	-${PORTVERSION:R}

MAINTAINER= 	dutra@FreeBSD.org
COMMENT=	Provides a common configuration to launch infrastructure
WWW=		https://terraform.io/

USES=		go:modules

LICENSE_NAME=   Business Source License 1.1
LICENSE_FILE=	${WRKSRC}/LICENSE
LICENSE_PERMS=	dist-mirror pkg-mirror pkg-sell auto-accept

USE_GITHUB=	yes
GH_ACCOUNT=	hashicorp

PLIST_FILES=	bin/${PORTNAME}${PKGNAMESUFFIX} \
		bin/${PORTNAME}-${PORTVERSION}

JQ_GETVER_EXPR=	[.[].name | gsub("^[a-z]+";"") | \
		select(test("^${PORTVERSION:C/[^.]+$/[0-9]+/:S/./\\\\./g}$$"; \
		"n"))][0]

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.Version=${PORTVERSION}' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.dev=no' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.VersionPrerelease='"

M2T_OFFLINE=	true
.export	M2T_OFFLINE

pre-extract: .PHONY .SILENT
	test -z '${WRKSRC_grpc_grpc_go_1}' || ${LN} -s ${WRKDIR}/grpc-go-cmd-protoc-gen-go-grpc-v1.1.0 ${WRKSRC_grpc_grpc_go_1}

post-install: .PHONY
	${MV} ${STAGEDIR}${PREFIX}/bin/${PORTNAME} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}-${PORTVERSION}
	${INSTALL} -l rs ${STAGEDIR}${PREFIX}/bin/${PORTNAME}-${PORTVERSION} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}${PKGNAMESUFFIX}


pre-deps post-deps-all: .PHONY
pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes-sigs GHT_PROJ=json GHT_PATH=sigs.k8s.io/json )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes-sigs GHT_PROJ=structured-merge-diff GHT_PATH=sigs.k8s.io/structured-merge-diff/v4 )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes GHT_PROJ=klog GHT_PATH=k8s.io/klog/v2 )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

# XXX: Work around ports framework inability to regenerate GH_TUPLE for Go
# modules for modules with embedded slashes in version numbers.
GH_TUPLE:=	${GH_TUPLE:S|v1.1.0:grpc_grpc_go_1|cmd%2Fprotoc-gen-go-grpc%2Fv1.1.0:grpc_grpc_go_1/|1}

# Azure:azure-sdk-for-go:v59.2.0:azure_azure_sdk_for_go/vendor/github.com/Azure/azure-sdk-for-go
#-Azure:go-autorest:v0.11.27:azure_go_autorest_autorest/vendor/github.com/Azure/go-autorest/autorest
#+Azure:go-autorest:autorest/v0.11.27:azure_go_autorest_autorest
#-Azure:go-autorest:v0.2.1:azure_go_autorest_logger/vendor/github.com/Azure/go-autorest/logger
#-Azure:go-autorest:v0.3.0:azure_go_autorest_date/vendor/github.com/Azure/go-autorest/autorest/date
#-Azure:go-autorest:v0.3.1:azure_go_autorest_validation/vendor/github.com/Azure/go-autorest/autorest/validation
#-Azure:go-autorest:v0.4.0:azure_go_autorest_to/vendor/github.com/Azure/go-autorest/autorest/to
#-Azure:go-autorest:v0.4.4:azure_go_autorest_cli/vendor/github.com/Azure/go-autorest/autorest/azure/cli
#-Azure:go-autorest:v0.6.0:azure_go_autorest_tracing/vendor/github.com/Azure/go-autorest/tracing
#-Azure:go-autorest:v0.9.20:azure_go_autorest_adal/vendor/github.com/Azure/go-autorest/autorest/adal
#+Azure:go-autorest:logger/v0.2.1:azure_go_autorest_logger
#+Azure:go-autorest:autorest/date/v0.3.0:azure_go_autorest_date
#+Azure:go-autorest:autorest/validation/v0.3.1:azure_go_autorest_validation
#+Azure:go-autorest:autorest/to/v0.4.0:azure_go_autorest_to
#+Azure:go-autorest:autorest/azure/cli/v0.4.4:azure_go_autorest_cli
#+Azure:go-autorest:tracing/v0.6.0:azure_go_autorest_tracing
#+Azure:go-autorest:autorest/adal/v0.9.20:azure_go_autorest_adal
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+)/vendor/github.com/Azure/go-autorest/(.*)|\1\3/\2|}

# asaskevich:govalidator:a9d515a09cc2:asaskevich_govalidator/vendor/github.com/asaskevich/govalidator
#-aws:aws-sdk-go-v2:v1.11.20:aws_aws_sdk_go_v2_presigned_url/vendor/github.com/aws/aws-sdk-go-v2/service/internal/presigned-url
#-aws:aws-sdk-go-v2:v1.11.5:aws_aws_sdk_go_v2_accept_encoding/vendor/github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding
#-aws:aws-sdk-go-v2:v1.16.14:aws_aws_sdk_go_v2_imds/vendor/github.com/aws/aws-sdk-go-v2/feature/ec2/imds
#-aws:aws-sdk-go-v2:v1.17.18:aws_aws_sdk_go_v2_s3shared/vendor/github.com/aws/aws-sdk-go-v2/service/internal/s3shared
#-aws:aws-sdk-go-v2:v1.17.22:aws_aws_sdk_go_v2_manager/vendor/github.com/aws/aws-sdk-go-v2/feature/s3/manager
#-aws:aws-sdk-go-v2:v1.17.34:aws_aws_sdk_go_v2_credentials/vendor/github.com/aws/aws-sdk-go-v2/credentials
#-aws:aws-sdk-go-v2:v1.23.0:aws_aws_sdk_go_v2_sso/vendor/github.com/aws/aws-sdk-go-v2/service/sso
#-aws:aws-sdk-go-v2:v1.27.0:aws_aws_sdk_go_v2_ssooidc/vendor/github.com/aws/aws-sdk-go-v2/service/ssooidc
#-aws:aws-sdk-go-v2:v1.27.36:aws_aws_sdk_go_v2_config/vendor/github.com/aws/aws-sdk-go-v2/config
#-aws:aws-sdk-go-v2:v1.3.18:aws_aws_sdk_go_v2_configsources/vendor/github.com/aws/aws-sdk-go-v2/internal/configsources
#-aws:aws-sdk-go-v2:v1.3.18:aws_aws_sdk_go_v2_v4a/vendor/github.com/aws/aws-sdk-go-v2/internal/v4a
#-aws:aws-sdk-go-v2:v1.3.20:aws_aws_sdk_go_v2_checksum/vendor/github.com/aws/aws-sdk-go-v2/service/internal/checksum
# aws:aws-sdk-go-v2:v1.31.0:aws_aws_sdk_go_v2/vendor/github.com/aws/aws-sdk-go-v2
#-aws:aws-sdk-go-v2:v1.31.0:aws_aws_sdk_go_v2_sts/vendor/github.com/aws/aws-sdk-go-v2/service/sts
#-aws:aws-sdk-go-v2:v1.35.0:aws_aws_sdk_go_v2_dynamodb/vendor/github.com/aws/aws-sdk-go-v2/service/dynamodb
#-aws:aws-sdk-go-v2:v1.35.0:aws_aws_sdk_go_v2_sqs/vendor/github.com/aws/aws-sdk-go-v2/service/sqs
#-aws:aws-sdk-go-v2:v1.36.0:aws_aws_sdk_go_v2_iam/vendor/github.com/aws/aws-sdk-go-v2/service/iam
#-aws:aws-sdk-go-v2:v1.6.5:aws_aws_sdk_go_v2_eventstream/vendor/github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream
#-aws:aws-sdk-go-v2:v1.63.0:aws_aws_sdk_go_v2_s3/vendor/github.com/aws/aws-sdk-go-v2/service/s3
#-aws:aws-sdk-go-v2:v1.8.1:aws_aws_sdk_go_v2_ini/vendor/github.com/aws/aws-sdk-go-v2/internal/ini
#-aws:aws-sdk-go-v2:v1.9.19:aws_aws_sdk_go_v2_endpoint_discovery/vendor/github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery
#-aws:aws-sdk-go-v2:v2.6.18:aws_aws_sdk_go_v2_v2/vendor/github.com/aws/aws-sdk-go-v2/internal/endpoints/v2
# aws:aws-sdk-go:v1.44.122:aws_aws_sdk_go/vendor/github.com/aws/aws-sdk-go
GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*/vendor/github.com/aws/aws-sdk-go-v2/*}

# google:go-genproto:b8732ec3820d:google_go_genproto/vendor/google.golang.org/genproto
#-google:go-genproto:b8732ec3820d:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/api
#-google:go-genproto:b8732ec3820d:google_go_genproto_2/vendor/google.golang.org/genproto/googleapis/rpc
# google:go-github:v45.2.0:google_go_github_v45/vendor/github.com/google/go-github/v45
GH_TUPLE:=	${GH_TUPLE:N*/google.golang.org/genproto/*}

#-googleapis:gax-go:v2.11.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.11.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|}

# googleapis:google-cloud-go:v0.110.7:googleapis_google_cloud_go/vendor/cloud.google.com/go
#-googleapis:google-cloud-go:v1.1.1:googleapis_google_cloud_go_iam/vendor/cloud.google.com/go/iam
#-googleapis:google-cloud-go:v1.30.1:googleapis_google_cloud_go_storage/vendor/cloud.google.com/go/storage
#+googleapis:google-cloud-go:storage/v1.30.1:googleapis_google_cloud_go_storage
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+)/vendor/cloud.google.com/go/(storage)|\1\3/\2|}
GH_TUPLE:=	${GH_TUPLE:N*/cloud.google.com/go/*}

#-hashicorp:consul:v1.13.0:hashicorp_consul_api/vendor/github.com/hashicorp/consul/api
#+hashicorp:consul:v1.13.0:hashicorp_consul_api/vendor/github.com/hashicorp/consul
GH_TUPLE:=	${GH_TUPLE:S|/consul/api|/consul|}

#+hashicorp:go-version:v1.7.0:hashicorp_go_version/vendor/github.com/hashicorp/go-version
GH_TUPLE+=	hashicorp:go-version:v1.7.0:hashicorp_go_version/vendor/github.com/hashicorp/go-version

# hashicorp:terraform-svchost:v0.1.1:hashicorp_terraform_svchost/vendor/github.com/hashicorp/terraform-svchost
#-hashicorp:terraform:000000000000:hashicorp_terraform_azure_1
#-hashicorp:terraform:000000000000:hashicorp_terraform_consul_1
#-hashicorp:terraform:000000000000:hashicorp_terraform_cos_1
#-hashicorp:terraform:000000000000:hashicorp_terraform_gcs_1
#-hashicorp:terraform:000000000000:hashicorp_terraform_kubernetes_1
#-hashicorp:terraform:000000000000:hashicorp_terraform_oss_1
#-hashicorp:terraform:000000000000:hashicorp_terraform_pg
#-hashicorp:terraform:000000000000:hashicorp_terraform_s3
#-hashicorp:terraform:000000000000:hashicorp_terraform_legacy
GH_TUPLE:=	${GH_TUPLE:Nhashicorp\:terraform\:000000000000\:*}

#-hashicorp:terraform::hashicorp_terraform_azure/github.com/hashicorp/terraform/internal/backend/remote-state/azure
#-hashicorp:terraform::hashicorp_terraform_consul/github.com/hashicorp/terraform/internal/backend/remote-state/consul
#-hashicorp:terraform::hashicorp_terraform_cos/github.com/hashicorp/terraform/internal/backend/remote-state/cos
#-hashicorp:terraform::hashicorp_terraform_gcs/github.com/hashicorp/terraform/internal/backend/remote-state/gcs
#-hashicorp:terraform::hashicorp_terraform_kubernetes/github.com/hashicorp/terraform/internal/backend/remote-state/kubernetes
#-hashicorp:terraform::hashicorp_terraform_oss/github.com/hashicorp/terraform/internal/backend/remote-state/oss
#-hashicorp:terraform::hashicorp_terraform_pg_1/github.com/hashicorp/terraform/internal/backend/remote-state/pg
#-hashicorp:terraform::hashicorp_terraform_s3_1/github.com/hashicorp/terraform/internal/backend/remote-state/s3
#-hashicorp:terraform::hashicorp_terraform_legacy_1/github.com/hashicorp/terraform/internal/legacy
# hashicorp:yamux:v0.1.1:hashicorp_yamux/vendor/github.com/hashicorp/yamux
GH_TUPLE:=	${GH_TUPLE:Nhashicorp\:terraform\:\:*}

#+open-telemetry:opentelemetry-go-contrib:instrumentation/github.com/aws/aws-sdk-go-v2/otelaws/v0.55.0:open_telemetry_opentelemetry_go_contrib_otelaws
#+open-telemetry:opentelemetry-go-contrib:instrumentation/google.golang.org/grpc/otelgrpc/v0.46.0:open_telemetry_opentelemetry_go_contrib_otelgrpc
#+open-telemetry:opentelemetry-go-contrib:exporters/autoexport/v0.45.0:open_telemetry_opentelemetry_go_contrib_autoexport
#+open-telemetry:opentelemetry-go:exporters/otlp/otlptrace/v1.19.0:open_telemetry_opentelemetry_go_otlptrace
#+open-telemetry:opentelemetry-go:v1.30.0:open_telemetry_opentelemetry_go/vendor/go.opentelemetry.io/otel
#+open-telemetry:opentelemetry-proto-go:otlp/v1.0.0:open_telemetry_opentelemetry_proto_go/vendor/go.opentelemetry.io/proto
GH_TUPLE+=	open-telemetry:opentelemetry-go-contrib:instrumentation/github.com/aws/aws-sdk-go-v2/otelaws/v0.55.0:open_telemetry_opentelemetry_go_contrib_otelaws
GH_TUPLE+=	open-telemetry:opentelemetry-go-contrib:instrumentation/google.golang.org/grpc/otelgrpc/v0.46.0:open_telemetry_opentelemetry_go_contrib_otelgrpc
GH_TUPLE+=	open-telemetry:opentelemetry-go-contrib:exporters/autoexport/v0.45.0:open_telemetry_opentelemetry_go_contrib_autoexport
GH_TUPLE+=	open-telemetry:opentelemetry-go:exporters/otlp/otlptrace/v1.19.0:open_telemetry_opentelemetry_go_otlptrace
GH_TUPLE+=	open-telemetry:opentelemetry-go:v1.30.0:open_telemetry_opentelemetry_go/vendor/go.opentelemetry.io/otel
GH_TUPLE+=	open-telemetry:opentelemetry-proto-go:otlp/v1.0.0:open_telemetry_opentelemetry_proto_go/vendor/go.opentelemetry.io/proto

# spf13:pflag:v1.0.5:spf13_pflag/vendor/github.com/spf13/pflag
#-tencentcloud:tencentcloud-sdk-go:v1.0.233:tencentcloud_tencentcloud_sdk_go_tag/vendor/github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tag
#+tencentcloud:tencentcloud-sdk-go:v1.0.233:tencentcloud_tencentcloud_sdk_go_tag/vendor/github.com/tencentcloud/tencentcloud-sdk-go
GH_TUPLE:=	${GH_TUPLE:S|/tencentcloud-sdk-go/tencentcloud/tag|/tencentcloud-sdk-go|}

#-tencentcloud:tencentcloud-sdk-go:v1.0.588:tencentcloud_tencentcloud_sdk_go_common/vendor/github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common
#+tencentcloud:tencentcloud-sdk-go:v1.0.588:tencentcloud_tencentcloud_sdk_go_common
GH_TUPLE:=	${GH_TUPLE:C|/vendor/github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common||}

#-tencentcloud:tencentcloud-sdk-go:v1.0.588:tencentcloud_tencentcloud_sdk_go_sts/vendor/github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts
GH_TUPLE:=	${GH_TUPLE:N*/tencentcloud-sdk-go/*}

.sinclude "${.CURDIR}/generated.deps.mk"
.include <bsd.port.mk>
