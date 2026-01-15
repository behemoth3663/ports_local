PORTNAME=	hcledit
PORTVERSION=	0.2.17
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Command line editor for HCL
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	minamijoyo

PLIST_FILES=	bin/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w"

post-extract: .PHONY
	${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}

pre-deps post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
