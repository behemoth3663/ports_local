#

PORTNAME=	checkip
PORTVERSION=	0.17.3
DISTVERSIONPREFIX=	v
CATEGORIES=	local net

MAINTAINER=	none@FreeBSD.org
COMMENT=	Check an IP address using various public services 

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	jreisinger

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
