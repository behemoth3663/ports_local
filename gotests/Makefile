PORTNAME=	gotests
PORTVERSION=	1.6.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	19
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Automatically generate Go test boilerplate from source code
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	cweill

GO_TARGET=	./${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w  -extldflags '-static'"

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
