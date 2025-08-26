PORTNAME=	pike
PORTVERSION=	0.3.77
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Determine permissions or policy required for IAC code
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	jamesWoolfenden

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static"

M2T_OFFLINE=	true
.export	M2T_OFFLINE

post-patch: .PHONY
	${REINPLACE_CMD} -r \
		-e 's|^(const Version).*$$|\1 = "${PORTVERSION}"|' \
	${WRKSRC}/src/version.go

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=imdario GHT_PROJ=mergo GHT_PATH=dario.cat/mergo )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	test -d ${PATCHDIR}_ && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-		aws:aws-sdk-go-v2:v1.3.15:aws_aws_sdk_go_v2_configsources/vendor/github.com/aws/aws-sdk-go-v2/internal/configsources
# 		aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2/vendor/github.com/aws/aws-sdk-go-v2
#-		aws:aws-sdk-go-v2:v1.30.3:aws_aws_sdk_go_v2_sts/vendor/github.com/aws/aws-sdk-go-v2/service/sts
GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*/vendor/github.com/aws/aws-sdk-go-v2/*}

#-googleapis:google-cloud-go:v0.16.2:googleapis_google_cloud_go_auth/vendor/cloud.google.com/go/auth
#+googleapis:google-cloud-go:auth/v0.16.2:googleapis_google_cloud_go_auth/vendor/cloud.google.com/go
GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)(.*/cloud.google.com/go)/(auth)$|\1\3/\2|}

#-googleapis:gax-go:v2.13.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go/v2
#+googleapis:gax-go:v2.13.0:googleapis_gax_go_v2/vendor/github.com/googleapis/gax-go
GH_TUPLE:=	${GH_TUPLE:S|/gax-go/v2|/gax-go|1}

#-google:go-genproto:513f23925822:google_go_genproto/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:513f23925822:google_go_genproto_rpc
GH_TUPLE:=	${GH_TUPLE:C,/vendor/google\.golang\.org/genproto/googleapis/(rpc),_\1,}

.sinclude "${.CURDIR}/generated.deps.mk"

GH_TUPLE+=	hashicorp:go-version:v1.7.0:hashicorp_go_version/vendor/github.com/hashicorp/go-version

.include <bsd.port.mk>
