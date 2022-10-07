#

PORTNAME=	s5cmd
PORTVERSION=	1.4.0
DISTVERSIONPREFIX=	v
PORTREVISION=	2
CATEGORIES=	local net

MAINTAINER=	none@freebsd.org
COMMENT=	Parallel S3 and local filesystem execution tool

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	peak

#GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/version.Version=${DISTVERSIONPREFIX}${PORTVERSION}'"

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
