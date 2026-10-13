PORTNAME=	copilot-proxy-go
PORTVERSION=	20260907
CATEGORIES=	local misc

MAINTAINER=	none@FreeBSD.org
COMMENT=	Turn GitHub Copilot into OpenAI/Anthropic API compatible server
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	tonghaoch
GH_TAGNAME=	4827e07

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-X main.version=${GH_TAGNAME}/${DISTVERSION}
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}

_OPTIONS_extract+=	750:post-post-extract
post-post-extract: .PHONY
	@${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps: .PHONY
	${TEST} -d ${PATCHDIR}_ && ${RM} -r ${PATCHDIR} || ${MV} ${PATCHDIR} ${PATCHDIR}_

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}_

post-deps-all: .PHONY
	${TEST} ! -d ${PATCHDIR}_ || ${MV} ${PATCHDIR}_ ${PATCHDIR}

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
