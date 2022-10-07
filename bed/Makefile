#

PORTNAME=	bed
PORTVERSION=	0.2.2
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local editors

MAINTAINER=	none@freebsd.org
COMMENT=	Binary editor written in Go

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	itchyny

PLIST_FILES=	bin/${PORTNAME}

GO_TARGET=	./cmd/${PORTNAME}

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
