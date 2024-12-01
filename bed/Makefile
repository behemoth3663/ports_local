PORTNAME=	bed
PORTVERSION=	0.2.8
DISTVERSIONPREFIX=	v
CATEGORIES=	local editors

MAINTAINER=	none@freebsd.org
COMMENT=	Binary editor written in Go
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	itchyny

PLIST_FILES=	bin/${PORTNAME}

GO_TARGET=	./cmd/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static"

post-extract: .PHONY
	test -f ${WRKSRC}/vendor/modules.txt || test ! -f ${PATCHDIR}/modules.txt || ${CP} ${PATCHDIR}/modules.txt ${WRKSRC}/vendor

pre-deps post-deps-all: .PHONY
pre-deps-all: .PHONY
	${CP} ${WRKSRC}/vendor/modules.txt ${PATCHDIR}

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
