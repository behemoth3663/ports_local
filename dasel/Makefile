PORTNAME=	dasel
PORTVERSION=	3.0.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	textproc local

MAINTAINER=	lcook@FreeBSD.org
COMMENT=	Query and update data structures from the command line
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	tomwright

GO_TARGET=	./cmd/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-tags noupdater -ldflags="-s -w \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/internal.Version=${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}
PORTDOCS=	CHANGELOG.md CODE_OF_CONDUCT.md README.md

OPTIONS_DEFINE=	DOCS

post-install-DOCS-on:
	@${MKDIR} ${STAGEDIR}${DOCSDIR}
	${INSTALL_DATA} ${PORTDOCS:S|^|${WRKSRC}/|} ${STAGEDIR}${DOCSDIR}

pre-deps: .PHONY
	test -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	test ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

_OPTIONS_extract+=	750:post-post-extract
post-post-extract: .PHONY
	test -f ${WRKSRC}/vendor/modules.txt || test ! -f ${PATCHDIR}/modules.txt || ${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
