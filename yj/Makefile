PORTNAME=	yj
PORTVERSION=	5.1.0
DISTVERSIONPREFIX=	v
PORTREVISION=	26
CATEGORIES=	local converters

MAINTAINER=	ehaupt@FreeBSD.org
COMMENT=	Convert between YAML, TOML, JSON, and HCL
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	sclevine

CONFLICTS_INSTALL=	yj-bruceadams

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static -X 'main.Version=${PORTVERSION}'"

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
