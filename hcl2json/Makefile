# $FreeBSD$

PORTNAME=	hcl2json
PORTVERSION=	0.3.2
DISTVERSIONPREFIX=	v
CATEGORIES=	textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Convert hcl2 to json

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	tmccombs

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

#rm -f -- distinfo Makefile.deps && make makesum && make go113mod-vendor && make makesum && mkcal
.sinclude "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
