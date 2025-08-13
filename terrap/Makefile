PORTNAME=	terrap
PORTVERSION=	0.0.4
DISTVERSIONPREFIX=	v
#ORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Scans infrastructure and identifies any required changes
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	sirrend
GH_PROJECT=	terrap-cli

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
		-X 'github.com/sirrend/terrap-cli/internal/version.Version=${DISTVERSIONPREFIX}${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
