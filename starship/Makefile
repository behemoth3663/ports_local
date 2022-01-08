#

PORTNAME=	starship
PORTVERSION=	1.1.1
DISTVERSIONPREFIX=	v
CATEGORIES=	local misc

MAINTAINER=	none@freebsd.org
COMMENT=	Minimal fast customizable prompt for any shell

LICENSE=	ISCL
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		cargo ssl
USE_GITHUB=	yes
GH_ACCOUNT=	starship

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

post-install: .PHONY
	${STRIP_CMD} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
#make makesum
#make makesum cargo-crates | grep -Ev '^=' >Makefile.deps
#make makesum
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
