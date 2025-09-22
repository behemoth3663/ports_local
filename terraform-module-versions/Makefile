PORTNAME=	terraform-module-versions
PORTVERSION=	3.3.15
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	CLI tool that checks Terraform code for module updates
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	keilerkonzept

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'main.version=${PORTVERSION}'"

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${RM} -r ${WRKSRC}/vendor/cloud.google.com/go/compute/metadata
	@${RLN} ${WRKSRC_googleapis_google_cloud_go_compute_metadata}/compute/metadata ${WRKSRC}/vendor/cloud.google.com/go/compute

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=darccio GHT_PROJ=mergo GHT_PATH=dario.cat/mergo )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=googleapis GHT_PROJ=google-cloud-go GHT_PATH=cloud.google.com/go/compute/metadata )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	test ! -d '${PATCHDIR}_' || ${MV} ${PATCHDIR}_ ${PATCHDIR}

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-aws:aws-sdk-go-v2:v1.3.15:aws_aws_sdk_go_v2_configsources/vendor/github.com/aws/aws-sdk-go-v2/internal/configsources
# aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2/vendor/github.com/aws/aws-sdk-go-v2
#-aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2_sts/vendor/github.com/aws/aws-sdk-go-v2/service/sts
GH_TUPLE:=	${GH_TUPLE:N*/vendor/github.com/aws/aws-sdk-go-v2/*}

# google:go-genproto:b8732ec3820d:google_go_genproto/vendor/google.golang.org/genproto
#-google:go-genproto:b8732ec3820d:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/api
#-google:go-genproto:b8732ec3820d:google_go_genproto_2/vendor/google.golang.org/genproto/googleapis/rpc
GH_TUPLE:=	${GH_TUPLE:N*/vendor/google.golang.org/genproto/googleapis/*}

#-googleapis:gax-go:v2.11.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.11.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S,/gax-go/v2,/gax-go,1}

# googleapis:google-cloud-go:v0.110.7:googleapis_google_cloud_go/vendor/cloud.google.com/go
GH_TUPLE:=	${GH_TUPLE:N*/vendor/cloud.google.com/go/*}

#-hashicorp:aws-sdk-go-base:v2.0.0-beta.65:hashicorp_aws_sdk_go_base_v2/vendor/github.com/hashicorp/aws-sdk-go-base/v2
#+hashicorp:aws-sdk-go-base:v2.0.0-beta.65:hashicorp_aws_sdk_go_base_v2/vendor/github.com/hashicorp/aws-sdk-go-base
GH_TUPLE:=	${GH_TUPLE:S,/hashicorp/aws-sdk-go-base/v2,/hashicorp/aws-sdk-go-base,1}

.sinclude "${.CURDIR}/generated.deps.mk"

GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+).*/cloud.google.com/go/(compute)/(metadata)$|\1\3/\4/\2_\3_\4|}

.include <bsd.port.mk>
