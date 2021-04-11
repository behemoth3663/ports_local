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

_OPTIONS_extract=	501:pre-post-extract

pre-post-extract: .PHONY
	${MV} ${WRKDIR}/logzio_terraform_client-1.2.0 ${WRKDIR}/logzio_client-1.2.0

fetch-version: .PHONY get-version

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor &&  make gomod-vendor >Makefile.deps && make makesum && mkcal
.-include "${.CURDIR}/Makefile.deps"

GH_TUPLE+=	gonum:gonum:v0.7.0:gonum_gonum/vendor/gonum.org/v1/gonum

.include <bsd.port.mk>
