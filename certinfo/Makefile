PORTNAME=	certinfo
PORTVERSION=	1.0.33
DISTVERSIONPREFIX=	v
CATEGORIES=	local security

MAINTAINER=	none@freebsd.org
COMMENT=	Print x509 certificate info
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	pete911

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'main.Version=${DISTVERSIONPREFIX}${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
