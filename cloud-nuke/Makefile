# $FreeBSD$

PORTNAME=	cloud-nuke
PORTVERSION=	0.1.30
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	CLI tool to delete all resources in an AWS account

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE.txt

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	gruntwork-io

#GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

_OPTIONS_extract=	501:pre-post-extract

pre-post-extract: .PHONY
	${MV} ${WRKDIR}/go-commons-0.7.0 ${WRKDIR}/gruntwork-cli-0.7.0

fetch-version: .PHONY get-version

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor && make gomod-vendor >Makefile.deps && make makesum && mkcal && instpc
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
