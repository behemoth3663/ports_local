PORTNAME=	terrascan
PORTVERSION=	1.19.9
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local security

MAINTAINER=	none@freebsd.org
COMMENT=	Detect compliance and security violations across IaC
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

USES=		go:modules

USE_GITHUB=	yes
USE_GITLAB=	nodefault

GH_ACCOUNT=	accurics

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

CGO_ENABLED=	0
# https://protobuf.dev/reference/go/faq/#namespace-conflict
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=ignore"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

#M2T_OFFLINE=	true
#.export	M2T_OFFLINE

_OPTIONS_extract=	750:post-post-extract
post-post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_
	cd ${WRKSRC} && \
		${CP} -a go.mod go.mod.orig && \
		${CP} -a go.sum go.sum.orig && \
		${GO_CMD} mod edit -replace=github.com/hashicorp/terraform=github.com/hashicorp/terraform@v0.15.3 && \
		${GO_CMD} mod edit -replace=modernc.org/libc=modernc.org/libc@v1.11.105 && \
		${GO_CMD} mod edit -replace=modernc.org/sqlite=modernc.org/sqlite@v1.13.3
	cd ${WRKSRC}/${GO_TARGET} && \
		${GO_CMD} mod tidy
#		${GO_CMD} get -u=patch &&
	cd ${WRKSRC} && \
		${SETENV} TZ=UTC ${DIFF} --minimal -u go.mod.orig go.mod | \
			${SED} -e '/^---/s|\.[0-9]* +0000$$| UTC|' -e '/^+++/s|\([[:blank:]][-0-9:.+]*\)*$$||' >${PATCHDIR}_/patch-go.mod && \
		${SETENV} TZ=UTC ${DIFF} --minimal -u go.sum.orig go.sum | \
			${SED} -e '/^---/s|\.[0-9]* +0000$$| UTC|' -e '/^+++/s|\([[:blank:]][-0-9:.+]*\)*$$||' >${PATCHDIR}_/patch-go.sum

post-deps-all: .PHONY
	${TEST} ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

#	${RM} ${.CURDIR}/generated.deps.mk
#	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=google GHT_PROJ=cel-spec GHT_PATH=cel.dev/expr )
#	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes-sigs GHT_PROJ=json GHT_PATH=sigs.k8s.io/json )
#	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes-sigs GHT_PROJ=structured-merge-diff GHT_PATH=sigs.k8s.io/structured-merge-diff/v4 )
#	( ${MAKE} -C ${.CURDIR} make-from-deps GHT_ACNT=kubernetes GHT_PROJ=klog GHT_PATH=k8s.io/klog/v2 )

#make-from-deps: .PHONY
#	@${SED} -rn \
#		's|^[[:blank:]]*#[[:blank:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:blank:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
#	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+)/vendor/github.com/Azure/go-autorest/(.*)|\1\3/\2|}
#GH_TUPLE:=	${GH_TUPLE:NGoogleCloudPlatform\:*}
#GH_TUPLE:=	${GH_TUPLE:Naws*aws_aws_sdk_go_v2_*/vendor/github.com/aws/aws-sdk-go-v2/*}
#GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+.*/envoyproxy/go-control-plane)/(envoy)$|\1\3/\2|}

# google:go-genproto:6ac0973c030d:google_go_genproto/vendor/google.golang.org/genproto
#-google:go-genproto:6ac0973c030d:google_go_genproto_api/vendor/google.golang.org/genproto/googleapis/api
#-google:go-genproto:6ac0973c030d:google_go_genproto_rpc/vendor/google.golang.org/genproto/googleapis/rpc
#GH_TUPLE:=	${GH_TUPLE:N*/vendor/google.golang.org/genproto/googleapis/*}

#GH_TUPLE:=	${GH_TUPLE:C|([^:]+:[^:]+:)([^:]+[^/]+)/vendor/cloud.google.com/go/(storage)|\1\3/\2|}
#GH_TUPLE:=	${GH_TUPLE:N*/cloud.google.com/go/*}

.include <bsd.port.mk>
