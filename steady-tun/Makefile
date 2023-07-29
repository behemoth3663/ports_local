PORTNAME=	steady-tun
PORTVERSION=	1.2.1
DISTVERSIONPREFIX=	v
CATEGORIES=	local net

MAINTAINER=	none@freebsd.org
COMMENT=	Secure TLS tunnel with pool of prepared upstream connections
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	Snawoot

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static -X 'main.version=${PORTVERSION}'"

post-post-extract: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
