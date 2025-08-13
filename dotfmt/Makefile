PORTNAME=	dotfmt
PORTVERSION=	20250119
DISTVERSIONPREFIX=	v
#PORTREVISION=	3
CATEGORIES=	local textproc

MAINTAINER=	none@FreeBSD.org
COMMENT=	Format DOT language files
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	GPLv3

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	teleivo
GH_PROJECT=	dot
GH_TAGNAME=	c0c9cfa

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION} | \
		${SED} -rn 's/(.{7}).*/\1/p'

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X 'main.version=${DISTVERSIONPREFIX}${PORTVERSION}' \
	-X 'main.commit=$$(${GIT_COMMIT})'"

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
