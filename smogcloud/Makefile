# $FreeBSD$

PORTNAME=	smogcloud
PORTVERSION=	0.0.1
DISTVERSIONPREFIX=	v
CATEGORIES=	security

MAINTAINER=	none@freebsd.org
COMMENT=	Comprehensive asset inventory

LICENSE=	GPLv3

USES=		go

USE_GITHUB=	yes
GH_ACCOUNT=	BishopFox
GH_TAGNAME=	30ce52e

PLIST_FILES=	bin/${PORTNAME}

.sinclude "${.CURDIR}/Makefile.deps"
.include "${.CURDIR}/../go113mod-vendor.mk"

.include <bsd.port.mk>
