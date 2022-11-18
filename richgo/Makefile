#

PORTNAME=	richgo
PORTVERSION=	0.3.11
DISTVERSIONPREFIX=	v
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Enrich `go test` outputs with text decorations
WWW=		https://github.com/kyoh86/richgo

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	kyoh86

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
