PORTNAME=	golint
PORTVERSION=	g20210508
#PORTREVISION=	19
CATEGORIES=	local devel

MAINTAINER=	yuri@FreeBSD.org
COMMENT=	Linter for Go source code
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	BSD3CLAUSE
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	golang
GH_PROJECT=	lint
GH_TAGNAME=	6edffad

GO_TARGET=	./${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w  -extldflags '-static'"

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
