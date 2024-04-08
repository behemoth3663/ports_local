PORTNAME=	go-binsize-treemap
DISTVERSIONPREFIX=	v
DISTVERSION=	0.2.3
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Go binary size SVG treemap
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	nikolaydubina

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
