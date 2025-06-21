PORTNAME=	gopls
PORTVERSION=	0.19.1
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel

MAINTAINER=	fuz@FreeBSD.org
COMMENT=	Go language LSP server
WWW=		https://pkg.go.dev/golang.org/x/tools/gopls

LICENSE=	BSD3CLAUSE
LICENSE_FILE=	${WRKSRC:H}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	golang
GH_PROJECT=	tools
GH_TAGNAME=	${PORTNAME}/${DISTVERSIONPREFIX}${PORTVERSION}

WRKSRC_SUBDIR=	${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'main.version=${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

M2T_OFFLINE=	true
.export	M2T_OFFLINE

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_
	${RM} ${.CURDIR}/generated.deps.mk
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=mvdan GHT_PROJ=xurls GHT_PATH=mvdan.cc/xurls/v2 )
	( make -C ${.CURDIR} make-from-deps GHT_ACNT=golang GHT_PROJ=exp GHT_PATH=golang.org/x/exp/typeparams )

make-from-deps: .PHONY
	@${SED} -rn \
		's|^[[:space:]]*#[[:space:]]*::([v0-9.a-f]+):.*/(vendor/${GHT_PATH})[[:space:]].*$$|GH_TUPLE+=${GHT_ACNT}:${GHT_PROJ}:\1:${GHT_ACNT:tl:S/-/_/g}_${GHT_PROJ:tl:S/-/_/g}/\2|p' \
	${.CURDIR}/Makefile.deps >>${.CURDIR}/generated.deps.mk

post-deps-all: .PHONY
	test -d ${PATCHDIR}_ && ${MV} ${PATCHDIR}_ ${PATCHDIR} || :

post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${RM} -r ${WRKSRC}/vendor/golang.org/x/tools/gopls
	@${RM} -r ${WRKSRC}/internal/telemetry
	@${RM} -r ${WRKSRC}/internal/server/counters.go

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-golang:telemetry:cc96b6e0d3d9:golang_telemetry/vendor/golang.org/x/telemetry
GH_TUPLE:=	${GH_TUPLE:Ngolang\:telemetry\:*}

.sinclude "${.CURDIR}/generated.deps.mk"

GH_TUPLE:=	${GH_TUPLE:S|/typeparams||1}

.include <bsd.port.mk>
