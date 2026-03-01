PORTNAME=	trdsql
PORTVERSION=	1.1.0
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local textproc
PKGNAMESUFFIX=	-local

MAINTAINER=	lcook@FreeBSD.org
COMMENT=	Execute SQL queries on CSV, LTSV, JSON, YAML and TBLN
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	noborus
GO_TARGET=	./cmd/${PORTNAME}

USE_GITLAB=	nodefault

# Building without CGO (CGO Free) can reduce issues related to cross-compiling, but it may result in slower execution times
GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}.Version=${PORTVERSION}'"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}
PORTDOCS=	README.md

OPTIONS_DEFINE=	DOCS

post-install-DOCS-on:
	@${MKDIR} ${STAGEDIR}${DOCSDIR}
	@${INSTALL_DATA} ${WRKSRC}/${PORTDOCS} ${STAGEDIR}${DOCSDIR}

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

post-extract: .PHONY
	@${TEST} -f ${WRKSRC}/vendor/modules.txt || ${TEST} ! -f ${PATCHDIR}/modules.txt || ${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor
	@${MKDIR} ${WRKSRC}/vendor/gonum.org/v1
	@${RLN} ${WRKSRC_gonum_gonum} ${WRKSRC}/vendor/gonum.org/v1/gonum

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

GH_TUPLE+=	gonum:gonum:v0.14.0:gonum_gonum/vendor/github.com/gonum/gonum

.include <bsd.port.mk>
