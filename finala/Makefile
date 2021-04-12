# $FreeBSD$

PORTNAME=	finala
PORTVERSION=	0.4.0
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils local

MAINTAINER=	none@freebsd.org
COMMENT=	Resource cloud scanner that reports wasteful and unused resources

LICENSE=	BSD3CLAUSE
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	similarweb

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

.-include "${.CURDIR}/Makefile.deps"
#.include "${.CURDIR}/../go113mod-vendor.mk"

.include <bsd.port.mk>
