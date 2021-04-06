# $FreeBSD$

PORTNAME=	terraformer
PORTVERSION=	0.8.12
#DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Generate terraform files from existing infrastructure

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	GoogleCloudPlatform

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor &&  make gomod-vendor >Makefile.deps && make makesum && mkcal
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
