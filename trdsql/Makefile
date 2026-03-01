PORTNAME=	trdsql
PORTVERSION=	1.1.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local textproc
PKGNAMESUFFIX?=	-local

MAINTAINER=	lcook@FreeBSD.org
COMMENT=	Execute SQL queries on CSV, LTSV, JSON, YAML and TBLN
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	noborus
GO_TARGET=	./cmd/${PORTNAME}

USE_GITLAB=	nodefault

# Building without CGO (CGO Free) can reduce issues related to cross-compiling, but it may result in slower execution times.
CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}.Version=${PORTVERSION}'"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}
PORTDOCS=	README.md

OPTIONS_DEFINE=	DOCS

post-install-DOCS-on:
	@${MKDIR} ${STAGEDIR}${DOCSDIR}
	${INSTALL_DATA} ${WRKSRC}/${PORTDOCS} ${STAGEDIR}${DOCSDIR}

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

post-extract: .PHONY
	${TEST} -f ${WRKSRC}/vendor/modules.txt || ${TEST} ! -f ${PATCHDIR}/modules.txt || ${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
