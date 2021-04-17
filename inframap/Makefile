# $FreeBSD$

PORTNAME=	inframap
PORTVERSION=	0.6.2
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Generate a human-readable graph from tfstate or HCL

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	cycloidio

#GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

#_OPTIONS_extract=	501:pre-post-extract
#pre-post-extract: .PHONY
#	${MV} ${WRKDIR}/go-commons-0.7.0 ${WRKDIR}/gruntwork-cli-0.7.0

fetch-version: .PHONY get-version

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor && make gomod-vendor >Makefile.deps && make makesum && mkcal && instpc
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
