PORTNAME=	dasel
PORTVERSION=	2.8.1
DISTVERSIONPREFIX=	v
PORTREVISION=	3
CATEGORIES=	textproc local

MAINTAINER=	lcook@FreeBSD.org
COMMENT=	Query and update data structures from the command line
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:1.24,modules

USE_GITHUB=	yes
GH_ACCOUNT=	tomwright

GO_TARGET=	./cmd/${PORTNAME}

GO_BUILDFLAGS+=	-tags noupdater -ldflags="-s -w -extldflags -static \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/internal.Version=${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}
PORTDOCS=	CHANGELOG.md CODE_OF_CONDUCT.md README.md

OPTIONS_DEFINE=	DOCS

post-install-DOCS-on:
	@${MKDIR} ${STAGEDIR}${DOCSDIR}
	${INSTALL_DATA} ${PORTDOCS:S|^|${WRKSRC}/|} ${STAGEDIR}${DOCSDIR}

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
