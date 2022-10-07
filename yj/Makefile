#

PORTNAME=	yj
PORTVERSION=	5.1.0
DISTVERSIONPREFIX=	v
PORTREVISION=	2
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	Convert between YAML, TOML, JSON, and HCL

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	sclevine

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -X 'main.Version=${PORTVERSION}'"
GO_BUILDFLAGS+=	-buildvcs=false

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
