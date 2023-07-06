PORTNAME=	dumbproxy
PORTVERSION=	1.10.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local net

MAINTAINER=	none@freebsd.org
COMMENT=	Dumbiest HTTP proxy ever
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	Snawoot

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static -X 'main.version=${PORTVERSION}'"

fetch-version: .PHONY get-version

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
