#

PORTNAME=	topngx
#DISTVERSIONPREFIX=	v
DISTVERSION=	0.3.1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Top for www/nginx

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		cargo
USE_GITHUB=	yes
GH_ACCOUNT=	gsquire
#GH_TAGNAME=	66d24fc

fetch-version: .PHONY get-version

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
#make makesum
#make cargo-crates | grep -Ev '^=' >Makefile.deps
#make makesum
.-include "${.CURDIR}/Makefile.deps"

PLIST_FILES=	bin/${PORTNAME}

post-install:
	${STRIP_CMD} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}

.include <bsd.port.mk>
