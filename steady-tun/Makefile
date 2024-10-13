PORTNAME=	steady-tun
PORTVERSION=	1.3.0
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

pre-deps post-deps-all: .PHONY

pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}

_OPTIONS_extract+=	750:post-post-extract
post-post-extract: .PHONY
	test -f ${WRKSRC}/vendor/modules.txt || test ! -f ${PATCHDIR}/modules.txt || ${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

.include "/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
