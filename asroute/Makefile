# $FreeBSD$

PORTNAME=	asroute
#DISTVERSIONPREFIX=	v
DISTVERSION=	0.1.0
CATEGORIES=	net

MAINTAINER=	none@freebsd.org
COMMENT=	Interpret traceroute output to show names of ASN traversed 

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		cargo
USE_GITHUB=	yes
GH_ACCOUNT=	stevenpack
GH_TAGNAME=	97d4f25

fetch-version: .PHONY get-version

#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-cargo
#make makesum
#make cargo-crates | grep -Ev '^=' >Makefile.deps
#make makesum
.sinclude "${.CURDIR}/Makefile.deps"

PLIST_FILES=	bin/${PORTNAME}

post-install:
	${STRIP_CMD} ${STAGEDIR}${PREFIX}/bin/${PORTNAME}

.include <bsd.port.mk>
