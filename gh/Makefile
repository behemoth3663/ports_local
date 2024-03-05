PORTNAME=	gh
PORTVERSION=	2.45.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel

MAINTAINER=	lcook@FreeBSD.org
COMMENT=	GitHub's official command line tool
WWW=		https://cli.github.com

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

RUN_DEPENDS=	git:devel/git

USES=		go:1.21,modules
USE_GITHUB=	yes
GH_ACCOUNT=	cli
GH_PROJECT=	${GH_ACCOUNT}

GO_TARGET=	./cmd/${PORTNAME}
GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/internal/build.Version=${PORTVERSION} \
	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/internal/build.Date=${%FT%TZ:L:gmtime}"

_DOCS=		README.md docs/[!README]*.md
PORTDOCS=	${_DOCS:T}

OPTIONS_DEFINE=		DOCS MANPAGES
OPTIONS_DEFAULT=	MANPAGES
OPTIONS_SUB=		yes

do-build-MANPAGES-on:
	(cd ${WRKSRC} && \
		${SETENV} ${MAKE_ENV} ${GO_ENV} ${GO_CMD} run ${GO_BUILDFLAGS} \
		./cmd/gen-docs --man-page --doc-path docs/man 2>/dev/null)

post-install-DOCS-on:
	@${MKDIR} ${STAGEDIR}${DOCSDIR}
	(cd ${WRKSRC} && \
		${INSTALL_DATA} ${_DOCS} \
		${STAGEDIR}${DOCSDIR})

post-install-MANPAGES-on:
	${INSTALL_MAN} ${WRKSRC}/docs/man/*.1 \
		${STAGEDIR}${PREFIX}/share/man/man1

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
