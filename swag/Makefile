PORTNAME=	swag
PORTVERSION=	1.16.6
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Generate RESTful API documentation with Swagger 2.0 for Go
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/license

USES=		go:modules
GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	swaggo

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static"

PLIST_FILES=	bin/${PORTNAME}

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
