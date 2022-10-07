#

PORTNAME=	smogcloud
PORTVERSION=	0.0.1
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local security

MAINTAINER=	none@freebsd.org
COMMENT=	Comprehensive asset inventory

LICENSE=	GPLv3

USES=		go

USE_GITHUB=	yes
GH_ACCOUNT=	BishopFox
GH_TAGNAME=	30ce52e

PLIST_FILES=	bin/${PORTNAME}

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
