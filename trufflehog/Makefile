PORTNAME=	trufflehog
PORTVERSION=	3.54.4
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Find and verify credentials
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	AGPLv3
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	trufflesecurity

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/v3/pkg/version.BuildVersion=${DISTVERSIONPREFIX}${PORTVERSION}'"

_OPTIONS_extract=	600:pre-post-extract
pre-post-extract: .PHONY .SILENT
	test -z '${WRKSRC_go_redis_redis}' || ${MV} ${WRKSRC_go_redis_redis:H}/go-${WRKSRC_go_redis_redis:T} ${WRKSRC_go_redis_redis}

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
