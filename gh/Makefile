PORTNAME=	gh
PORTVERSION=	2.59.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel

MAINTAINER=	lcook@FreeBSD.org
COMMENT=	GitHub's official command line tool
WWW=		https://cli.github.com

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

RUN_DEPENDS=	git:devel/git

USES=		go:modules,1.23
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

.include "/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-moby:moby:v0.7.0:moby_moby_1/vendor/github.com/docker/docker-credential-helpers
#+docker:docker-credential-helpers:v0.7.0:docker_docker_credential_helpers/vendor/github.com/docker/docker-credential-helpers
GH_TUPLE:=	${GH_TUPLE:C|^moby:moby(:.*/docker-credential-helpers)$|docker:docker-credential-helpers\1|}

#-google:go-genproto:c811ad7063a7:google_go_genproto/vendor/google.golang.org/genproto/googleapis/api
#+google:go-genproto:c811ad7063a7:google_go_genproto_api
#-google:go-genproto:94a12d6c2237:google_go_genproto_1/vendor/google.golang.org/genproto/googleapis/rpc
#+google:go-genproto:94a12d6c2237:google_go_genproto_rpc
GH_TUPLE:=	${GH_TUPLE:C|(:google_go_genproto).*/|\1_|}

post-extract: .PHONY
	@${MKDIR} ${WRKSRC}/vendor/google.golang.org/genproto/googleapis
	@test -z '${WRKSRC_google_go_genproto_api}' || ${RLN} ${WRKSRC_google_go_genproto_api}/googleapis/api ${WRKSRC}/vendor/google.golang.org/genproto/googleapis/api
	@test -z '${WRKSRC_google_go_genproto_rpc}' || ${RLN} ${WRKSRC_google_go_genproto_rpc}/googleapis/rpc ${WRKSRC}/vendor/google.golang.org/genproto/googleapis/rpc

.include <bsd.port.mk>
