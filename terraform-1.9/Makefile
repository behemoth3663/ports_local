PORTNAME=	terraform
PORTVERSION=	1.9.8
DISTVERSIONPREFIX=	v
PORTREVISION=	20250113
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

#pre-build: .PHONY
#	cd ${WRKSRC} && ${SETENV} ${MAKE_ENV} ${GO_ENV} ${GO_CMD} generate && ${SETENV} ${MAKE_ENV} ${GO_ENV} ${GO_CMD} mod vendor

pre-extract: .PHONY .SILENT
	test -z '${WRKSRC_grpc_grpc_go_1}' || ${LN} -s ${WRKDIR}/grpc-go-cmd-protoc-gen-go-grpc-v1.1.0 ${WRKSRC_grpc_grpc_go_1}

post-install: .PHONY
	${MV} ${STAGEDIR}${PREFIX}/bin/${PORTNAME} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}-${PORTVERSION}
	${INSTALL} -l rs ${STAGEDIR}${PREFIX}/bin/${PORTNAME}-${PORTVERSION} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}${PKGNAMESUFFIX}

'pre-deps: .PHONY
	if [ -d '${MASTERDIR}/files_' ]; then \
		${RM} -r ${PATCHDIR}; \
	else \
		${MV} ${PATCHDIR} ${PATCHDIR}_; \
	fi
	${RM} -r ${PATCHDIR}-${PORTVERSION}
	${MKDIR} ${PATCHDIR}-${PORTVERSION}

	cd ${WRKSRC} && \
		${CP} -a go.mod go.mod.orig && \
		${CP} -a go.sum go.sum.orig && \
		${GO_CMD} get -u=patch && \
		${GO_CMD} mod tidy && \
		${SETENV} TZ=UTC ${DIFF} -udp go.mod.orig go.mod | \
			${SED} -e '/^---/s|\.[0-9]* +0000$$| UTC|' -e '/^+++/s|\([[:blank:]][-0-9:.+]*\)*$$||' >${PATCHDIR}-${PORTVERSION}/patch-go.mod && \
		${SETENV} TZ=UTC ${DIFF} -udp go.sum.orig go.sum | \
			${SED} -e '/^---/s|\.[0-9]* +0000$$| UTC|' -e '/^+++/s|\([[:blank:]][-0-9:.+]*\)*$$||' >${PATCHDIR}-${PORTVERSION}/patch-go.sum

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}

pre-deps post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

# XXX: Work around ports framework inability to regenerate GH_TUPLE for Go
# modules for modules with embedded slashes in version numbers.
GH_TUPLE:=	${GH_TUPLE:S|v1.1.0:grpc_grpc_go_1|cmd%2Fprotoc-gen-go-grpc%2Fv1.1.0:grpc_grpc_go_1/|1}

#-Azure:go-autorest:v0.3.1:azure_go_autorest_validation/vendor/github.com/Azure/go-autorest/autorest/validation
# Azure:go-autorest:v14.2.0:azure_go_autorest/vendor/github.com/Azure/go-autorest
#+Azure:go-autorest:autorest/validation/v0.3.1:azure_go_autorest_validation
GH_TUPLE:=	${GH_TUPLE:C|(Azure:go-autorest:)([^/]+).*/go-autorest/(.+)$|\1\3/\2|}

#-aws:aws-sdk-go-v2:v1.3.15:aws_aws_sdk_go_v2_configsources/vendor/github.com/aws/aws-sdk-go-v2/internal/configsources
# aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2/vendor/github.com/aws/aws-sdk-go-v2
#-aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2_sts/vendor/github.com/aws/aws-sdk-go-v2/service/sts
GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*/vendor/github.com/aws/aws-sdk-go-v2/*}

# google:go-genproto:b8732ec3820d:google_go_genproto/vendor/google.golang.org/genproto
#-google:go-genproto:b8732ec3820d:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/api
#-google:go-genproto:b8732ec3820d:google_go_genproto_2/vendor/google.golang.org/genproto/googleapis/rpc
GH_TUPLE:=	${GH_TUPLE:Ngoogle\:go-genproto\:*/genproto/*}

#-		googleapis:gax-go:v2.13.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+		googleapis:gax-go:v2.13.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|}

# googleapis:google-cloud-go:v0.110.7:googleapis_google_cloud_go/vendor/cloud.google.com/go
#-googleapis:google-cloud-go:v1.1.1:googleapis_google_cloud_go_iam/vendor/cloud.google.com/go/iam
#-googleapis:google-cloud-go:v1.23.0:googleapis_google_cloud_go_compute/vendor/cloud.google.com/go/compute
#-googleapis:google-cloud-go:v1.30.1:googleapis_google_cloud_go_storage/vendor/cloud.google.com/go/storage
GH_TUPLE:=	${GH_TUPLE:Ngoogleapis\:google-cloud-go\:*/go/*}

#-hashicorp:consul:v1.13.0:hashicorp_consul_api/vendor/github.com/hashicorp/consul/api
#+hashicorp:consul:v1.13.0:hashicorp_consul_api/vendor/github.com/hashicorp/consul
GH_TUPLE:=	${GH_TUPLE:S|/consul/api$|/consul|}

#-hashicorp:terraform:000000000000:hashicorp_terraform_azure_1
#-hashicorp:terraform::hashicorp_terraform_azure/github.com/hashicorp/terraform/internal/backend/remote-state/azure
GH_TUPLE:=	${GH_TUPLE:C|(hashicorp:terraform:)[[:xdigit:]]{12}:|\1:|}
GH_TUPLE:=	${GH_TUPLE:Nhashicorp\:terraform\:\:*}

#-tencentcloud:tencentcloud-sdk-go:v1.0.233:tencentcloud_tencentcloud_sdk_go_tag/vendor/github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tag
#+tencentcloud:tencentcloud-sdk-go:v1.0.233:tencentcloud_tencentcloud_sdk_go_tag/vendor/github.com/tencentcloud/tencentcloud-sdk-go
GH_TUPLE:=	${GH_TUPLE:S|/tencentcloud/tag$||}

#-tencentcloud:tencentcloud-sdk-go:v1.0.588:tencentcloud_tencentcloud_sdk_go_common/vendor/github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common
#-tencentcloud:tencentcloud-sdk-go:v1.0.588:tencentcloud_tencentcloud_sdk_go_sts/vendor/github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sts
GH_TUPLE:=	${GH_TUPLE:Ntencentcloud\:tencentcloud-sdk-go\:*/tencentcloud-sdk-go/*}

.include <bsd.port.mk>
