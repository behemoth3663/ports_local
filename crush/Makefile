PORTNAME=	crush
DISTVERSIONPREFIX=	v
DISTVERSION=	0.98.1
CATEGORIES=	local misc devel
PKGNAMESUFFIX=	-local

MAINTAINER=	tagattie@FreeBSD.org
COMMENT=	Glamourous AI coding agent for your favourite terminal
WWW=		https://${GO_MODULE}

LICENSE=	FSL-1.1-MIT
LICENSE_NAME=	Functional Source License, Version 1.1, MIT Future License
LICENSE_FILE=	${WRKSRC}/LICENSE.md
LICENSE_PERMS=	dist-mirror pkg-mirror auto-accept

USES=		go:1.27+,modules

GO_MODULE=	github.com/charmbracelet/crush

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-X ${GO_MODULE}/internal/version.Version=${PORTVERSION}
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}
PORTDOCS=	README.md

OPTIONS_DEFINE=		COMPLETIONS DISABLE_COMMAND_BLOCKING DOCS MANPAGES
OPTIONS_DEFAULT=	COMPLETIONS DISABLE_COMMAND_BLOCKING MANPAGES

COMPLETIONS_DESC=		Build and/or install shell completions
DISABLE_COMMAND_BLOCKING_DESC=	Support for disabling dangerous command blocking

COMPLETIONS_PLIST_FILES=etc/bash_completion.d/${PORTNAME} \
			share/fish/completions/${PORTNAME}.fish \
			share/zsh/site-functions/_${PORTNAME}

DISABLE_COMMAND_BLOCKING_EXTRA_PATCHES=	${PATCHDIR}/extra-patch-disable-command-blocking

MANPAGES_PLIST_FILES=	share/man/man1/${PORTNAME}.1.gz

pre-build: .PHONY
	${RM} -r ${WRKSRC}/vendor/go.opentelemetry.io
	${RM} ${WRKSRC}/vendor/github.com/googleapis/gax-go/v2/telemetry.go

post-build-COMPLETIONS-on:
	@${MKDIR} ${GO_WRKSRC}/completions
.for _sh in bash fish zsh
	${GO_WRKDIR_BIN}/${PORTNAME} completion ${_sh} > \
		${GO_WRKSRC}/completions/${PORTNAME}.${_sh}
.endfor

post-build-MANPAGES-on:
	@${MKDIR} ${GO_WRKSRC}/manpages
	${GO_WRKDIR_BIN}/${PORTNAME} man > ${GO_WRKSRC}/manpages/${PORTNAME}.1

post-install-COMPLETIONS-on:
	@${MKDIR} ${COMPLETIONS_PLIST_FILES:H:S|^|${STAGEDIR}${PREFIX}/|}
.for _sh in bash fish zsh
	${INSTALL_DATA} ${GO_WRKSRC}/completions/${PORTNAME}.${_sh} \
		${COMPLETIONS_PLIST_FILES:M*${_sh}*:S|^|${STAGEDIR}${PREFIX}/|}
.endfor

post-install-DOCS-on:
	@${MKDIR} ${STAGEDIR}${DOCSDIR}
	${INSTALL_MAN} ${PORTDOCS:S|^|${GO_WRKSRC}/|} ${STAGEDIR}${DOCSDIR}

post-install-MANPAGES-on:
	${INSTALL_MAN} ${GO_WRKSRC}/manpages/${PORTNAME}.1 \
		${STAGEDIR}/${PREFIX}/share/man/man1

.include <bsd.port.mk>
