PORTNAME=	speedtest-go-server
PORTVERSION=	1.1.5
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local net

MAINTAINER=	none@freebsd.org
COMMENT=	Go backend for LibreSpeed
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	LGPL3

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	librespeed
GH_PROJECT=	speedtest-go

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
