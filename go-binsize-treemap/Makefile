PORTNAME=	go-binsize-treemap
DISTVERSIONPREFIX=	v
DISTVERSION=	0.2.0
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

#GO_BUILDFLAGS+=	-ldflags="-s -w -X 'main.Version=${DISTVERSIONPREFIX}${PORTVERSION}'"

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
