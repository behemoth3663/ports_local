PORTNAME=	qrcp
PORTVERSION=	0.11.6
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Transfer files over Wi-Fi from computer to mobile device
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	claudiodangelis

GO_BUILDFLAGS+=-ldflags="-s -w -extldflags -static \
	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.version=${PORTVERSION} \
	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.date=${%FT%TZ:L:gmtime}'"

PLIST_FILES=	bin/${PORTNAME}

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
