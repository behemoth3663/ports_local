PORTNAME=	qrcp
PORTVERSION=	0.11.7
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils
PKGNAMESUFFIX=	-local

MAINTAINER=	none@freebsd.org
COMMENT=	Transfer files over Wi-Fi from computer to mobile device
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	claudiodangelis

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_LDFLAGS+=	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.version=${PORTVERSION}
GO_LDFLAGS+=	-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.date=${%FT%TZ:L:gmtime}
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
