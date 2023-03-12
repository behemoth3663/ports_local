#

PORTNAME=	hcledit
PORTVERSION=	0.2.6
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Command line editor for HCL
WWW=		https://github.com/minamijoyo/hcledit

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	minamijoyo

PLIST_FILES=	bin/${PORTNAME}

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
