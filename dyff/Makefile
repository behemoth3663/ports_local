PORTNAME=	dyff
DISTVERSIONPREFIX=	v
DISTVERSION=	1.8.1
CATEGORIES=	local textproc

MAINTAINER=	none@freebsd.org
COMMENT=	diff tool for YAML files, and sometimes JSON
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	homeport

GO_TARGET=	./cmd/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X 'github.com/${GH_ACCOUNT}/${GH_PROJECT}/internal/cmd.version=${DISTVERSIONPREFIX}${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
