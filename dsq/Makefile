PORTNAME=	dsq
PORTVERSION=	0.23.0
DISTVERSIONPREFIX=	v
#ORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Run SQL queries against JSON, CSV, Excel, Parquet, and more
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE.md

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	multiprocessio

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'main.Version=${DISTVERSIONPREFIX}${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

post-post-extract: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
