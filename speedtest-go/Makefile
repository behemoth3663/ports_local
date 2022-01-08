#

PORTNAME=	speedtest-go
PORTVERSION=	1.1.4
DISTVERSIONPREFIX=	v
CATEGORIES=	local net

MAINTAINER=	none@freebsd.org
COMMENT=	Go backend for LibreSpeed

LICENSE=	LGPL3
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	librespeed

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
