PORTNAME=	zek
PORTVERSION=	0.1.21
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Create a Go struct from an XML document
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	GPLv3
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules
GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	miku

PLIST_FILES=	bin/${PORTNAME}

post-post-extract: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
