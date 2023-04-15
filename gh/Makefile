PORTNAME=	gh
PORTVERSION=	2.27.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel

MAINTAINER=	lcook@FreeBSD.org
COMMENT=	GitHub's official command line tool
WWW=		https://cli.github.com

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

RUN_DEPENDS=	git:devel/git

USES=		go:modules
USE_GITHUB=	yes
GH_ACCOUNT=	cli
GH_PROJECT=	${GH_ACCOUNT}

GO_TARGET=	./cmd/${PORTNAME}
GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/internal/build.Version=${PORTVERSION} \
	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/v2/internal/build.Date=${%FT%TZ:L:gmtime}"

_DOCS=		README.md docs/gh-vs-hub.md docs/releasing.md \
		docs/command-line-syntax.md docs/source.md \
		docs/install_linux.md docs/project-layout.md docs/triage.md
PORTDOCS=	${_DOCS:T}

OPTIONS_DEFINE=		COMPLETIONS DOCS MANPAGES
OPTIONS_DEFAULT=	COMPLETIONS MANPAGES
OPTIONS_SUB=		yes

COMPLETIONS_DESC=		Build and install bash, fish and zsh shell completions
COMPLETIONS_PLIST_FILES=	share/bash-completion/completions/${PORTNAME} \
				share/fish/vendor_completions.d/${PORTNAME}.fish \
				share/zsh/site-functions/_${PORTNAME}

post-install-COMPLETIONS-on:
.for _shell in bash fish zsh
	${GO_WRKDIR_BIN}/${PORTNAME} completion -s ${_shell} \
		>${STAGEDIR}${PREFIX}/${COMPLETIONS_PLIST_FILES:M*${_shell}*}
.endfor

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
		${STAGEDIR}${MAN1PREFIX}/man/man1

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
